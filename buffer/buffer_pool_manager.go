// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// buffer_pool_manager.go
//
// Go translation of src/buffer/buffer_pool_manager.cpp.
// This file intentionally keeps high-level structure and TODOs,
// but provides stub implementations so the package compiles.
//
// ===----------------------------------------------------------------------===//

package buffer

import (
	"sync"
	"sync/atomic"

	"github.com/PhanNam1501/bustub-go/common"
	"github.com/PhanNam1501/bustub-go/storage/disk"
	typ "github.com/PhanNam1501/bustub-go/type"
)

type AccessType int

const (
	AccessTypeUnknown AccessType = iota
	AccessTypeRead
	AccessTypeWrite
)

type PageID = common.PageID

type FrameID = common.FrameID

type DiskManager = disk.DiskManager

type DiskScheduler = disk.DiskScheduler

type LogManager interface{}

type BufferPoolManager struct {
	NumFrames     int
	NextPageID    int32
	BPMLatch      sync.Mutex
	Replacer      *ArcReplacer
	DiskManager   *DiskManager
	DiskScheduler *DiskScheduler
	LogManager    LogManager

	Frames    []*typ.FrameHeader
	PageTable map[PageID]FrameID
	FreeList  []FrameID
}

func NewBufferPoolManager(numFrames int, diskManager *DiskManager, logManager LogManager) *BufferPoolManager {

	bpm := &BufferPoolManager{
		NumFrames:     numFrames,
		NextPageID:    0,
		Replacer:      NewArcReplacer(numFrames),
		DiskManager:   diskManager,
		DiskScheduler: disk.NewDiskScheduler(diskManager),
		LogManager:    logManager,
		Frames:        make([]*typ.FrameHeader, 0, numFrames),
		PageTable:     make(map[PageID]FrameID, numFrames),
		FreeList:      make([]FrameID, 0, numFrames),
	}

	bpm.BPMLatch.Lock()
	defer bpm.BPMLatch.Unlock()

	for i := 0; i < int(numFrames); i++ {
		fid := FrameID(i)
		bpm.Frames = append(bpm.Frames, typ.NewFrameHeader(fid))
		bpm.FreeList = append(bpm.FreeList, fid)
	}

	return bpm
}

func (bpm *BufferPoolManager) findAvailableFrame() (FrameID, bool) {
	if len(bpm.FreeList) > 0 {
		fid := bpm.FreeList[0]
		bpm.FreeList = bpm.FreeList[1:]
		return fid, true
	}

	fid, ok := bpm.Replacer.Evict()
	if ok {
		frame := bpm.Frames[fid]
		if frame.IsDirty {
			promise := bpm.DiskScheduler.CreatePromise()
			bpm.DiskScheduler.Schedule([]*disk.DiskRequest{{
				IsWrite: true, Data: frame.Data, PageID: frame.PageID, Callback: promise,
			}})
			<-promise

		}
		delete(bpm.PageTable, frame.PageID)
		return fid, true
	}
	return 0, false
}

func (bpm *BufferPoolManager) flushFrameToDiskUnsafe(frame *typ.FrameHeader) {
	promise := bpm.DiskScheduler.CreatePromise()
	bpm.DiskScheduler.Schedule([]*disk.DiskRequest{{
		IsWrite:  true,
		Data:     frame.Data,
		PageID:   frame.PageID,
		Callback: promise,
	}})
	<-promise
	frame.IsDirty = false
}

func (bpm *BufferPoolManager) Size() int {
	return bpm.NumFrames
}

func (bpm *BufferPoolManager) unpinPage(pageID PageID, isDirty bool) {
	bpm.BPMLatch.Lock()
	defer bpm.BPMLatch.Unlock()
	if fid, ok := bpm.PageTable[pageID]; ok {
		frame := bpm.Frames[fid]
		if frame.PinCount > 0 {
			frame.PinCount--
			if isDirty {
				frame.IsDirty = true
			}
			if frame.PinCount == 0 {
				bpm.Replacer.SetEvictable(fid, true)
			}
		}
	}
}

func (bpm *BufferPoolManager) NewPage() PageID {
	bpm.BPMLatch.Lock()
	defer bpm.BPMLatch.Unlock()

	fid, ok := bpm.findAvailableFrame()
	if !ok {
		return common.InvalidPageID
	}

	pID := PageID(atomic.AddInt32(&bpm.NextPageID, 1) - 1)
	frame := bpm.Frames[fid]
	frame.Reset()
	frame.PageID = pID
	frame.PinCount = 1

	bpm.PageTable[pID] = fid
	bpm.Replacer.SetEvictable(fid, false)
	bpm.Replacer.RecordAccess(fid, pID, AccessTypeWrite)

	return pID
}

func (bpm *BufferPoolManager) DeletePage(pageID PageID) bool {
	bpm.BPMLatch.Lock()
	defer bpm.BPMLatch.Unlock()

	fid, ok := bpm.PageTable[pageID]
	if !ok {
		bpm.DiskManager.DeletePage(pageID)
		return true
	}
	frame := bpm.Frames[fid]
	if frame.PinCount > 0 {
		return false
	}

	bpm.DiskManager.DeletePage(pageID)

	delete(bpm.PageTable, pageID)
	bpm.Replacer.Remove(fid)
	frame.Reset()
	bpm.FreeList = append(bpm.FreeList, fid)

	return true
}

func (bpm *BufferPoolManager) CheckedWritePage(pageID PageID, accessType AccessType) (*WritePageGuard, bool) {
	bpm.BPMLatch.Lock()

	fid, ok := bpm.PageTable[pageID]
	if !ok {
		fid, ok = bpm.findAvailableFrame()
		if !ok {
			bpm.BPMLatch.Unlock()
			return nil, false
		}

		frame := bpm.Frames[fid]
		frame.PageID = pageID
		frame.PinCount = 0

		promise := bpm.DiskScheduler.CreatePromise()
		bpm.DiskScheduler.Schedule([]*disk.DiskRequest{{
			IsWrite:  false,
			Data:     frame.Data,
			PageID:   pageID,
			Callback: promise,
		}})
		<-promise
		bpm.PageTable[pageID] = fid
	}

	frame := bpm.Frames[fid]
	frame.PinCount++
	bpm.Replacer.SetEvictable(fid, false)
	bpm.Replacer.RecordAccess(fid, pageID, accessType)

	frame.Mu.Lock()
	bpm.BPMLatch.Unlock()
	return &WritePageGuard{bpm: bpm, frame: frame}, true
}

func (bpm *BufferPoolManager) CheckedReadPage(pageID PageID, accessType AccessType) (*ReadPageGuard, bool) {
	bpm.BPMLatch.Lock()

	fid, ok := bpm.PageTable[pageID]
	if !ok {
		fid, ok = bpm.findAvailableFrame()
		if !ok {
			bpm.BPMLatch.Unlock()
			return nil, false
		}

		frame := bpm.Frames[fid]
		frame.PageID = pageID
		frame.PinCount = 0

		promise := bpm.DiskScheduler.CreatePromise()
		bpm.DiskScheduler.Schedule([]*disk.DiskRequest{{
			IsWrite:  false,
			Data:     frame.Data,
			PageID:   pageID,
			Callback: promise,
		}})
		<-promise
		bpm.PageTable[pageID] = fid
	}

	frame := bpm.Frames[fid]
	frame.PinCount++
	bpm.Replacer.SetEvictable(fid, false)
	bpm.Replacer.RecordAccess(fid, pageID, accessType)

	// Khóa Đọc
	frame.Mu.RLock()
	bpm.BPMLatch.Unlock()

	return &ReadPageGuard{bpm: bpm, frame: frame}, true
}

func (bpm *BufferPoolManager) WritePage(pageID PageID, accessType AccessType) *WritePageGuard {
	guard, ok := bpm.CheckedWritePage(pageID, accessType)
	if !ok {
		panic("CheckedWritePage failed to bring in page")
	}
	return guard
}

func (bpm *BufferPoolManager) ReadPage(pageID PageID, accessType AccessType) *ReadPageGuard {
	guard, ok := bpm.CheckedReadPage(pageID, accessType)
	if !ok {
		panic("CheckedReadPage failed to bring in page")
	}
	return guard
}

func (bpm *BufferPoolManager) FlushPageUnsafe(pageID PageID) bool {
	bpm.BPMLatch.Lock()
	defer bpm.BPMLatch.Unlock()

	fid, ok := bpm.PageTable[pageID]
	if !ok {
		return false
	}
	bpm.flushFrameToDiskUnsafe(bpm.Frames[fid])
	return true
}

func (bpm *BufferPoolManager) FlushPage(pageID PageID) bool {
	bpm.BPMLatch.Lock()
	fid, ok := bpm.PageTable[pageID]
	if !ok {
		bpm.BPMLatch.Unlock()
		return false
	}
	frame := bpm.Frames[fid]
	bpm.BPMLatch.Unlock()

	frame.Mu.RLock()
	defer frame.Mu.RUnlock()

	bpm.BPMLatch.Lock()
	defer bpm.BPMLatch.Unlock()
	if frame.IsDirty {
		bpm.flushFrameToDiskUnsafe(frame)
	}
	return true
}

func (bpm *BufferPoolManager) FlushAllPagesUnsafe() {
	bpm.BPMLatch.Lock()

	pageIDs := make([]PageID, 0, len(bpm.PageTable))
	for pageID := range bpm.PageTable {
		pageIDs = append(pageIDs, pageID)
	}
	bpm.BPMLatch.Unlock()

	for _, p := range pageIDs {
		bpm.FlushPageUnsafe(p)
	}
}

func (bpm *BufferPoolManager) FlushAllPages() {
	bpm.BPMLatch.Lock()

	pageIDs := make([]PageID, 0, len(bpm.PageTable))
	for pageID := range bpm.PageTable {
		pageIDs = append(pageIDs, pageID)
	}
	bpm.BPMLatch.Unlock()

	for _, p := range pageIDs {
		bpm.FlushPage(p)
	}
}

func (bpm *BufferPoolManager) GetPinCount(pageID PageID) (int32, bool) {
	bpm.BPMLatch.Lock()
	defer bpm.BPMLatch.Unlock()
	if fid, ok := bpm.PageTable[pageID]; ok {
		return bpm.Frames[fid].PinCount, true
	}
	return 0, false
}
