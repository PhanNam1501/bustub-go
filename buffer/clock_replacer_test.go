package buffer

import "testing"

func TestClockReplacer_DISABLED_SampleTest(t *testing.T) {
	t.Skip("TODO: Clock replacer logic not implemented yet")

	clock := NewClockReplacer(7)

	clock.Unpin(1)
	clock.Unpin(2)
	clock.Unpin(3)
	clock.Unpin(4)
	clock.Unpin(5)
	clock.Unpin(6)
	clock.Unpin(1)

	var victim FrameID
	_ = clock.Victim(&victim)
}

