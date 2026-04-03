package buffer

import "testing"

func TestLRUKReplacer_DISABLED_SampleTest(t *testing.T) {
	t.Skip("TODO: LRU-K replacer logic not implemented yet")

	replacer := NewLRUKReplacer(7, 2)

	replacer.RecordAccess(1, AccessTypeUnknown)
	replacer.RecordAccess(2, AccessTypeUnknown)
	replacer.RecordAccess(3, AccessTypeUnknown)
	replacer.RecordAccess(4, AccessTypeUnknown)
	replacer.RecordAccess(5, AccessTypeUnknown)
	replacer.RecordAccess(6, AccessTypeUnknown)
	replacer.SetEvictable(1, true)
	replacer.SetEvictable(2, true)
	replacer.SetEvictable(3, true)
	replacer.SetEvictable(4, true)
	replacer.SetEvictable(5, true)
	replacer.SetEvictable(6, false)

	_ = replacer.Size()

	replacer.RecordAccess(1, AccessTypeUnknown)
	_, _ = replacer.Evict()
}

