// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// arc_replacer.go
//
// Go translation of src/buffer/arc_replacer.cpp.
// This keeps the high-level API and TODOs but provides stubs so code compiles.
//
// ===----------------------------------------------------------------------===//

package buffer

import (
	"container/list"
	"sync"
)

// ArcReplacer is the Go counterpart of bustub::ArcReplacer.
type ArcReplacer struct {
	mu           sync.Mutex
	replacerSize int32
	// TODO: add the internal lists and bookkeeping fields corresponding to mru_, mfu_, ghost lists, etc.
	mruTargetSize int32
	currSize      int32
	mru           *list.List
	mfu           *list.List

	mruGhost *list.List
	mfuGhost *list.List

	mruMap      map[FrameID]*list.Element
	mfuMap      map[FrameID]*list.Element
	mruGhostMap map[PageID]*list.Element
	mfuGhostMap map[PageID]*list.Element

	frameToPage map[FrameID]PageID
	evictable   map[FrameID]bool
}

// NewArcReplacer creates a new ArcReplacer with the given maximum number of frames.
//
// TODO: P1 - Add full implementation as in the C++ version.
func NewArcReplacer(numFrames int32) *ArcReplacer {
	// TODO: initialize lists to be empty and target size to 0.
	return &ArcReplacer{
		replacerSize:  numFrames,
		mruTargetSize: 0,
		currSize:      0,

		mru:      list.New(),
		mfu:      list.New(),
		mruGhost: list.New(),
		mfuGhost: list.New(),

		mruMap:      make(map[FrameID]*list.Element),
		mfuMap:      make(map[FrameID]*list.Element),
		mruGhostMap: make(map[PageID]*list.Element),
		mfuGhostMap: make(map[PageID]*list.Element),

		frameToPage: make(map[FrameID]PageID),
		evictable:   make(map[FrameID]bool),
	}
}

func (ar *ArcReplacer) findEvictable(l *list.List) *list.Element {
	for e := l.Back(); e != nil; e = e.Prev() {
		frameID := e.Value.(FrameID)
		if ar.evictable[frameID] {
			return e
		}
	}
	return nil
}

// Evict performs the ARC replacement logic and returns the victim frame ID, if any.
//
// TODO: P1 - Implement the full ARC algorithm, including handling of ghost lists and evictable flags.
func (ar *ArcReplacer) Evict() (FrameID, bool) {
	// TODO: implement ARC eviction.
	ar.mu.Lock()
	defer ar.mu.Unlock()

	if ar.currSize == 0 {
		return 0, false
	}

	var victimElem *list.Element
	fromMru := false

	//If the MRU list size is smaller than the target size, we try to evict from the MFU list.
	if ar.mru.Len() < int(ar.mruTargetSize) {
		victimElem = ar.findEvictable(ar.mfu)
		if victimElem == nil {
			victimElem = ar.findEvictable(ar.mru)
			fromMru = true
		}
		//If the MRU list size is greater than or equal to the target size, we try to evict from the MRU list.
	} else {
		victimElem = ar.findEvictable(ar.mru)
		fromMru = true
		if victimElem == nil {
			victimElem = ar.findEvictable(ar.mfu)
			fromMru = false
		}
	}

	if victimElem == nil {
		return 0, false
	}

	frameID := victimElem.Value.(FrameID)
	pageID := ar.frameToPage[frameID]

	if fromMru {
		ar.mru.Remove(victimElem)
		delete(ar.mruMap, frameID)
		e := ar.mruGhost.PushFront(pageID)
		ar.mruGhostMap[pageID] = e
	} else {
		ar.mfu.Remove(victimElem)
		delete(ar.mfuMap, frameID)
		e := ar.mfuGhost.PushFront(pageID)
		ar.mruGhostMap[pageID] = e
	}

	ar.evictable[frameID] = false
	delete(ar.frameToPage, frameID)
	ar.currSize--

	return frameID, true
}

// RecordAccess records an access to a frame and updates ARC bookkeeping.
//
// TODO: P1 - Implement handling of hits/misses in all four lists as in the C++ comments.
func (ar *ArcReplacer) RecordAccess(frameID FrameID, pageID PageID, accessType AccessType) {
	// TODO: implement RecordAccess.
	ar.mu.Lock()
	defer ar.mu.Unlock()

	ar.frameToPage[frameID] = pageID

	_, inMru := ar.mruMap[frameID]
	_, inMfu := ar.mfuMap[frameID]
	if inMru || inMfu {
		if inMru {
			ar.mru.Remove(ar.mruMap[frameID])
			delete(ar.mruMap, frameID)
		} else {
			ar.mfu.Remove(ar.mfuMap[frameID])
			delete(ar.mfuMap, frameID)
		}
		e := ar.mfu.PushFront(frameID)
		ar.mfuMap[frameID] = e
		return
	}

	if elem, ok := ar.mruGhostMap[pageID]; ok {
		delta := 1
		if ar.mruGhost.Len() < ar.mfuGhost.Len() {
			delta = ar.mfuGhost.Len() / ar.mruGhost.Len()
		}
		ar.mruTargetSize += int32(delta)
		if ar.mruTargetSize > ar.replacerSize {
			ar.mruTargetSize = ar.replacerSize
		}

		ar.mruGhost.Remove(elem)
		delete(ar.mruGhostMap, pageID)

		e := ar.mfu.PushFront(frameID)
		ar.mfuMap[frameID] = e
		return
	}

	if elem, ok := ar.mfuGhostMap[pageID]; ok {
		delta := 1
		if ar.mfuGhost.Len() > ar.mruGhost.Len() {
			delta = ar.mruGhost.Len() / ar.mfuGhost.Len()
		}

		ar.mruTargetSize -= int32(delta)
		if ar.mruTargetSize < 0 {
			ar.mruTargetSize = 0
		}

		ar.mfuGhost.Remove(elem)
		delete(ar.mfuGhostMap, pageID)

		e := ar.mru.PushFront(frameID)
		ar.mruMap[frameID] = e
		return
	}

	if ar.mru.Len()+ar.mfu.Len() == int(ar.replacerSize) {
		if ar.mruGhost.Len() > 0 {
			last := ar.mruGhost.Back()
			pID := last.Value.(PageID)
			ar.mruGhost.Remove(last)
			delete(ar.mruGhostMap, pID)
		}
	} else if ar.mru.Len()+ar.mruGhost.Len() < int(ar.replacerSize) {
		totalSize := ar.mru.Len() + ar.mruGhost.Len() + ar.mfu.Len() + ar.mfuGhost.Len()
		if totalSize >= 2*int(ar.replacerSize) {
			if ar.mfuGhost.Len() > 0 {
				last := ar.mfuGhost.Back()
				pID := last.Value.(PageID)
				ar.mfuGhost.Remove(last)
				delete(ar.mfuGhostMap, pID)
			}
		}
	}

	e := ar.mru.PushFront(frameID)
	ar.mruMap[frameID] = e
}

// SetEvictable toggles whether a frame is evictable, adjusting replacer size accordingly.
//
// TODO: P1 - Track per-frame evictable state and update size.
func (ar *ArcReplacer) SetEvictable(frameID FrameID, setEvictable bool) {
	// TODO: implement SetEvictable.
	ar.mu.Lock()
	defer ar.mu.Unlock()

	currentStatus, exists := ar.evictable[frameID]
	if !exists {
		currentStatus = false
	}

	if currentStatus != setEvictable {
		ar.evictable[frameID] = setEvictable
		if setEvictable {
			ar.currSize++
		} else {
			ar.currSize--
		}
	}
}

// Remove removes an evictable frame from the replacer.
//
// TODO: P1 - Implement Remove semantics (different from Evict).
func (ar *ArcReplacer) Remove(frameID FrameID) {
	// TODO: implement Remove.
	ar.mu.Lock()
	defer ar.mu.Unlock()

	if !ar.evictable[frameID] {
		return
	}

	if elem, ok := ar.mruMap[frameID]; ok {
		ar.mru.Remove(elem)
		delete(ar.mruMap, frameID)
	} else if elem, ok := ar.mfuMap[frameID]; ok {
		ar.mfu.Remove(elem)
		delete(ar.mfuMap, frameID)
	} else {
		return
	}

	delete(ar.evictable, frameID)
	delete(ar.frameToPage, frameID)
	ar.currSize--
}

// Size returns the number of evictable frames tracked by the replacer.
//
// TODO: P1 - Return the actual size instead of 0.
func (ar *ArcReplacer) Size() int32 {
	ar.mu.Lock()
	defer ar.mu.Unlock()
	return ar.currSize
}
