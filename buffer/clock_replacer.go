// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// clock_replacer.go
//
// Optional Clock policy replacer (stub).
//
// ===----------------------------------------------------------------------===//

package buffer

// ClockReplacer is a placeholder for the optional Clock replacement policy.
// TODO: implement ClockReplacer per BusTub.
type ClockReplacer struct {
	// TODO: add bookkeeping fields.
}

// NewClockReplacer creates a new ClockReplacer.
func NewClockReplacer(numFrames int) *ClockReplacer {
	return &ClockReplacer{}
}

// Victim chooses a victim frame, if any.
func (cr *ClockReplacer) Victim(frameID *FrameID) bool {
	if frameID == nil {
		return false
	}
	*frameID = 0
	return false
}

// Pin marks a frame as pinned.
func (cr *ClockReplacer) Pin(frameID FrameID) {}

// Unpin marks a frame as unpinned.
func (cr *ClockReplacer) Unpin(frameID FrameID) {}

// Size returns the number of evictable frames.
func (cr *ClockReplacer) Size() int { return 0 }

