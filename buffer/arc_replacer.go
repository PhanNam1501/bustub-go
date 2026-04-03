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

// ArcReplacer is the Go counterpart of bustub::ArcReplacer.
type ArcReplacer struct {
	replacerSize int
	// TODO: add the internal lists and bookkeeping fields corresponding to mru_, mfu_, ghost lists, etc.
}

// NewArcReplacer creates a new ArcReplacer with the given maximum number of frames.
//
// TODO: P1 - Add full implementation as in the C++ version.
func NewArcReplacer(numFrames int) *ArcReplacer {
	// TODO: initialize lists to be empty and target size to 0.
	return &ArcReplacer{
		replacerSize: numFrames,
	}
}

// Evict performs the ARC replacement logic and returns the victim frame ID, if any.
//
// TODO: P1 - Implement the full ARC algorithm, including handling of ghost lists and evictable flags.
func (ar *ArcReplacer) Evict() (FrameID, bool) {
	// TODO: implement ARC eviction.
	return 0, false
}

// RecordAccess records an access to a frame and updates ARC bookkeeping.
//
// TODO: P1 - Implement handling of hits/misses in all four lists as in the C++ comments.
func (ar *ArcReplacer) RecordAccess(frameID FrameID, pageID PageID, accessType AccessType) {
	// TODO: implement RecordAccess.
}

// SetEvictable toggles whether a frame is evictable, adjusting replacer size accordingly.
//
// TODO: P1 - Track per-frame evictable state and update size.
func (ar *ArcReplacer) SetEvictable(frameID FrameID, setEvictable bool) {
	// TODO: implement SetEvictable.
}

// Remove removes an evictable frame from the replacer.
//
// TODO: P1 - Implement Remove semantics (different from Evict).
func (ar *ArcReplacer) Remove(frameID FrameID) {
	// TODO: implement Remove.
}

// Size returns the number of evictable frames tracked by the replacer.
//
// TODO: P1 - Return the actual size instead of 0.
func (ar *ArcReplacer) Size() int {
	return 0
}
