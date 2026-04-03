// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// lru_k_replacer.go
//
// Go translation of src/buffer/lru_k_replacer.cpp.
// This file preserves the TODOs but provides stub implementations.
//
// ===----------------------------------------------------------------------===//

package buffer

// LRUKReplacer is the Go counterpart of bustub::LRUKReplacer.
type LRUKReplacer struct {
	replacerSize int
	k            int
	// TODO: add access history structures and evictable bookkeeping.
}

// NewLRUKReplacer creates a new LRUKReplacer.
//
// TODO: P1 - Add full implementation and data structures for k-distance tracking.
func NewLRUKReplacer(numFrames int, k int) *LRUKReplacer {
	return &LRUKReplacer{
		replacerSize: numFrames,
		k:            k,
	}
}

// Evict finds a victim frame ID based on backward k-distance.
//
// TODO: P1 - Implement eviction semantics exactly as described in the C++ comments.
func (r *LRUKReplacer) Evict() (FrameID, bool) {
	// TODO: implement LRUK eviction.
	return 0, false
}

// RecordAccess records an access to a frame at the current timestamp.
//
// TODO: P1 - Track history, handle invalid frame IDs, and maintain timestamps.
func (r *LRUKReplacer) RecordAccess(frameID FrameID, accessType AccessType) {
	// TODO: implement RecordAccess.
}

// SetEvictable toggles whether a frame is evictable and adjusts size.
//
// TODO: P1 - Track per-frame evictable state and update size.
func (r *LRUKReplacer) SetEvictable(frameID FrameID, setEvictable bool) {
	// TODO: implement SetEvictable.
}

// Remove removes an evictable frame from the replacer, along with its history.
//
// TODO: P1 - Implement Remove semantics (different from Evict).
func (r *LRUKReplacer) Remove(frameID FrameID) {
	// TODO: implement Remove.
}

// Size returns the number of evictable frames currently tracked.
//
// TODO: P1 - Return actual size instead of 0.
func (r *LRUKReplacer) Size() int {
	return 0
}
