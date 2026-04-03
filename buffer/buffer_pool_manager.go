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

	"github.com/PhanNam1501/bustub-go/common"
)

// AccessType describes the kind of access for buffer pool operations.
// This mirrors bustub::AccessType but is simplified for now.
type AccessType int

const (
	AccessTypeUnknown AccessType = iota
	AccessTypeRead
	AccessTypeWrite
)

// PageID is an alias for common.PageID to keep names close to the C++ version.
type PageID = common.PageID

// FrameID is an alias for common.FrameID.
type FrameID = common.FrameID

// FrameHeader is the Go counterpart of bustub::FrameHeader.
// It stores metadata and the actual page bytes.
type FrameHeader struct {
	frameID  FrameID
	data     []byte
	pinCount int32
	isDirty  bool
	mu       sync.RWMutex
}

// NewFrameHeader creates a new FrameHeader with a given frame ID.
func NewFrameHeader(frameID FrameID) *FrameHeader {
	h := &FrameHeader{
		frameID: frameID,
		data:    make([]byte, common.BustubPageSize),
	}
	h.Reset()
	return h
}

// GetData returns an immutable view of the frame data.
func (h *FrameHeader) GetData() []byte {
	return h.data
}

// GetDataMut returns a mutable view of the frame data.
func (h *FrameHeader) GetDataMut() []byte {
	return h.data
}

// Reset clears the frame contents and metadata.
func (h *FrameHeader) Reset() {
	for i := range h.data {
		h.data[i] = 0
	}
	h.pinCount = 0
	h.isDirty = false
}

// DiskManager is a minimal interface placeholder for the disk manager used by the buffer pool.
// This will be filled in when storage/disk is ported.
type DiskManager interface{}

// LogManager is a minimal interface placeholder for the log manager.
type LogManager interface{}

// Replacer is the interface implemented by replacement policies (ARC, LRU-K, ...).
type Replacer interface {
	// TODO: define the methods when we port the replacers.
}

// ReadPageGuard provides safe shared read access to a page's data.
// TODO: implement proper guard semantics (pin/unpin + lock/latch).
type ReadPageGuard struct {
	bpm    *BufferPoolManager
	pageID PageID
	frame  *FrameHeader
}

// GetData returns immutable page bytes.
func (g *ReadPageGuard) GetData() []byte {
	if g == nil || g.frame == nil {
		return nil
	}
	return g.frame.GetData()
}

// Release frees access (stub).
func (g *ReadPageGuard) Release() {
	// TODO: implement unpin / unlock semantics
}

// WritePageGuard provides safe exclusive write access to a page's data.
// TODO: implement proper guard semantics (pin/unpin + lock/latch).
type WritePageGuard struct {
	bpm    *BufferPoolManager
	pageID PageID
	frame  *FrameHeader
}

// GetData returns immutable page bytes.
func (g *WritePageGuard) GetData() []byte {
	if g == nil || g.frame == nil {
		return nil
	}
	return g.frame.GetData()
}

// GetDataMut returns mutable page bytes.
func (g *WritePageGuard) GetDataMut() []byte {
	if g == nil || g.frame == nil {
		return nil
	}
	return g.frame.GetDataMut()
}

// Release frees access (stub).
func (g *WritePageGuard) Release() {
	// TODO: implement unpin / unlock semantics
}

// BufferPoolManager is the Go counterpart of bustub::BufferPoolManager.
type BufferPoolManager struct {
	numFrames   int
	nextPageID  int32
	bpmLatch    *sync.Mutex
	replacer    Replacer
	diskManager DiskManager
	logManager  LogManager

	frames    []*FrameHeader
	pageTable map[PageID]FrameID
	freeList  []FrameID
}

// NewBufferPoolManager creates a new BufferPoolManager.
func NewBufferPoolManager(numFrames int, diskManager DiskManager, logManager LogManager) *BufferPoolManager {
	// NOTE: This mirrors the C++ constructor logic at a high level,
	// but leaves detailed behavior for later.

	bpm := &BufferPoolManager{
		numFrames:   numFrames,
		nextPageID:  0,
		bpmLatch:    &sync.Mutex{},
		replacer:    nil, // will be set once ArcReplacer is ported
		diskManager: diskManager,
		logManager:  logManager,
		frames:      make([]*FrameHeader, 0, numFrames),
		pageTable:   make(map[PageID]FrameID, numFrames),
		freeList:    make([]FrameID, 0, numFrames),
	}

	bpm.bpmLatch.Lock()
	defer bpm.bpmLatch.Unlock()

	for i := 0; i < int(numFrames); i++ {
		fid := FrameID(i)
		bpm.frames = append(bpm.frames, NewFrameHeader(fid))
		bpm.freeList = append(bpm.freeList, fid)
	}

	return bpm
}

// Size returns the number of frames managed by this buffer pool.
func (bpm *BufferPoolManager) Size() int {
	return int(bpm.numFrames)
}

// NewPage allocates a new page ID.
//
// TODO: Implement full logic to coordinate with DiskManager and page table.
func (bpm *BufferPoolManager) NewPage() PageID {
	// TODO: P1 - implement NewPage properly.
	// For now, just bump the counter so code compiles and tests can be added.
	bpm.bpmLatch.Lock()
	defer bpm.bpmLatch.Unlock()
	id := PageID(bpm.nextPageID)
	bpm.nextPageID++
	return id
}

// DeletePage removes a page from disk and the buffer pool, if possible.
//
// TODO: Implement full logic from the C++ version.
func (bpm *BufferPoolManager) DeletePage(pageID PageID) bool {
	// TODO: P1 - implement DeletePage.
	return false
}

// CheckedWritePage attempts to bring a page into memory and returns exclusive access.
//
// TODO: Implement full logic including replacement and I/O.
func (bpm *BufferPoolManager) CheckedWritePage(pageID PageID, accessType AccessType) (*WritePageGuard, bool) {
	// TODO: P1 - implement CheckedWritePage.
	return nil, false
}

// CheckedReadPage attempts to bring a page into memory and returns shared access.
//
// TODO: Implement full logic including replacement and I/O.
func (bpm *BufferPoolManager) CheckedReadPage(pageID PageID, accessType AccessType) (*ReadPageGuard, bool) {
	// TODO: P1 - implement CheckedReadPage.
	return nil, false
}

// WritePage is a test/ergonomic wrapper around CheckedWritePage.
//
// TODO: implement proper abort behavior once IO/replacement is ported.
func (bpm *BufferPoolManager) WritePage(pageID PageID, accessType AccessType) *WritePageGuard {
	guard, ok := bpm.CheckedWritePage(pageID, accessType)
	if !ok {
		panic("CheckedWritePage failed to bring in page")
	}
	return guard
}

// ReadPage is a test/ergonomic wrapper around CheckedReadPage.
//
// TODO: implement proper abort behavior once IO/replacement is ported.
func (bpm *BufferPoolManager) ReadPage(pageID PageID, accessType AccessType) *ReadPageGuard {
	guard, ok := bpm.CheckedReadPage(pageID, accessType)
	if !ok {
		panic("CheckedReadPage failed to bring in page")
	}
	return guard
}

// FlushPageUnsafe flushes a single page to disk without taking page-level locks.
//
// TODO: Implement full logic after storage and disk are available.
func (bpm *BufferPoolManager) FlushPageUnsafe(pageID PageID) bool {
	// TODO: P1 - implement FlushPageUnsafe.
	return false
}

// FlushPage flushes a single page to disk with page-level locking.
//
// TODO: Implement full logic after storage and disk are available.
func (bpm *BufferPoolManager) FlushPage(pageID PageID) bool {
	// TODO: P1 - implement FlushPage.
	return false
}

// FlushAllPagesUnsafe flushes all dirty pages to disk without taking page-level locks.
//
// TODO: Implement full logic after storage and disk are available.
func (bpm *BufferPoolManager) FlushAllPagesUnsafe() {
	// TODO: P1 - implement FlushAllPagesUnsafe.
}

// FlushAllPages flushes all dirty pages to disk with page-level locks.
//
// TODO: Implement full logic after storage and disk are available.
func (bpm *BufferPoolManager) FlushAllPages() {
	// TODO: P1 - implement FlushAllPages.
}

// GetPinCount returns the pin count of a page if it exists.
//
// TODO: Implement full logic mirroring the C++ behavior.
func (bpm *BufferPoolManager) GetPinCount(pageID PageID) (int32, bool) {
	// TODO: P1 - implement GetPinCount.
	return 0, false
}
