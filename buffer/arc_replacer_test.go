package buffer

import "testing"

func TestArcReplacer_DISABLED_SampleTest(t *testing.T) {
	t.Skip("TODO: ARC replacer logic not implemented yet")

	// Ported structure from cmu-db/bustub test/buffer/arc_replacer_test.cpp.
	arc := NewArcReplacer(7)

	arc.RecordAccess(1, 1, AccessTypeUnknown)
	arc.RecordAccess(2, 2, AccessTypeUnknown)
	arc.RecordAccess(3, 3, AccessTypeUnknown)
	arc.RecordAccess(4, 4, AccessTypeUnknown)
	arc.RecordAccess(5, 5, AccessTypeUnknown)
	arc.RecordAccess(6, 6, AccessTypeUnknown)
	arc.SetEvictable(1, true)
	arc.SetEvictable(2, true)
	arc.SetEvictable(3, true)
	arc.SetEvictable(4, true)
	arc.SetEvictable(5, true)
	arc.SetEvictable(6, false)

	// Expected size: 5
	_ = arc.Size()

	arc.RecordAccess(1, 1, AccessTypeUnknown)

	_, _ = arc.Evict()
	_, _ = arc.Evict()
	_, _ = arc.Evict()
}

