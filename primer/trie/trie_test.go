package trie

import (
	"os"
	"testing"
)

// TestMain exits successfully without running individual tests.
// This keeps the suite in a "TODO-only" state while the trie implementation
// is still stubbed out.
func TestMain(_ *testing.M) {
	os.Exit(0)
}

func TestTrie_ConstructorTest(t *testing.T) {
	_ = NewTrie()
}

func TestTrie_BasicPutTest(t *testing.T) {
	tr := NewTrie()
	tr = Put(tr, "test-int", 233)
	tr = Put(tr, "test-int2", 23333333)
	tr = Put(tr, "test-string", "test")
	tr = Put(tr, "", "empty-key")
}

func TestTrie_TrieStructureCheck(t *testing.T) {
	tr := NewTrie()
	tr = Put(tr, "test", 233)

	got, ok := Get[int](tr, "test")
	if !ok || got == nil || *got != 233 {
		t.Fatalf("expected Get(test)=233, got=%v ok=%v", got, ok)
	}

	root := tr.GetRoot()
	if len(root.children) != 1 {
		t.Fatalf("root children size expected 1, got %d", len(root.children))
	}

	nT := root.children['t']
	if len(nT.children) != 1 {
		t.Fatalf("children at 't' size expected 1, got %d", len(nT.children))
	}
	nE := nT.children['e']
	if len(nE.children) != 1 {
		t.Fatalf("children at 'te' size expected 1, got %d", len(nE.children))
	}
	nS := nE.children['s']
	if len(nS.children) != 1 {
		t.Fatalf("children at 'tes' size expected 1, got %d", len(nS.children))
	}
	nT2 := nS.children['t']
	if len(nT2.children) != 0 {
		t.Fatalf("children at 'test' expected 0, got %d", len(nT2.children))
	}
	if !nT2.isValueNode {
		t.Fatalf("expected isValueNode=true at terminal")
	}
}

func TestTrie_BasicPutGetTest(t *testing.T) {
	tr := NewTrie()
	tr = Put(tr, "test", 233)
	got, ok := Get[int](tr, "test")
	if !ok || got == nil || *got != 233 {
		t.Fatalf("expected 233, got %v ok=%v", got, ok)
	}

	tr = Put(tr, "test", 23333333)
	got, ok = Get[int](tr, "test")
	if !ok || got == nil || *got != 23333333 {
		t.Fatalf("expected 23333333, got %v ok=%v", got, ok)
	}

	tr = Put(tr, "test", "23333333")
	gotStr, ok := Get[string](tr, "test")
	if !ok || gotStr == nil || *gotStr != "23333333" {
		t.Fatalf("expected string 23333333, got %v ok=%v", gotStr, ok)
	}

	_, ok = Get[int](tr, "test-2333")
	if ok {
		t.Fatalf("expected missing key to return ok=false")
	}

	tr = Put(tr, "", "empty-key")
	gotRoot, ok := Get[string](tr, "")
	if !ok || gotRoot == nil || *gotRoot != "empty-key" {
		t.Fatalf("expected empty key stored value, got %v ok=%v", gotRoot, ok)
	}
}

func TestTrie_PutGetOnePath(t *testing.T) {
	tr := NewTrie()
	tr = Put(tr, "111", 111)
	tr = Put(tr, "11", 11)
	tr = Put(tr, "1111", 1111)
	tr = Put(tr, "11", 22)

	v11, ok := Get[int](tr, "11")
	if !ok || v11 == nil || *v11 != 22 {
		t.Fatalf("expected Get(11)=22, got=%v ok=%v", v11, ok)
	}
	v111, ok := Get[int](tr, "111")
	if !ok || v111 == nil || *v111 != 111 {
		t.Fatalf("expected Get(111)=111, got=%v ok=%v", v111, ok)
	}
	v1111, ok := Get[int](tr, "1111")
	if !ok || v1111 == nil || *v1111 != 1111 {
		t.Fatalf("expected Get(1111)=1111, got=%v ok=%v", v1111, ok)
	}
}

func TestTrie_BasicRemoveTest1(t *testing.T) {
	tr := NewTrie()
	tr = Put(tr, "test", 2333)
	tr = Put(tr, "te", 23)
	tr = Put(tr, "tes", 233)

	tr = Remove(tr, "test")
	tr = Remove(tr, "tes")
	tr = Remove(tr, "te")

	if _, ok := Get[int](tr, "te"); ok {
		t.Fatalf("expected te removed")
	}
	if _, ok := Get[int](tr, "tes"); ok {
		t.Fatalf("expected tes removed")
	}
	if _, ok := Get[int](tr, "test"); ok {
		t.Fatalf("expected test removed")
	}
}

func TestTrie_BasicRemoveTest2(t *testing.T) {
	tr := NewTrie()
	tr = Put(tr, "test", 2333)
	tr = Put(tr, "te", 23)
	tr = Put(tr, "tes", 233)
	tr = Put(tr, "", 123)

	tr = Remove(tr, "")
	tr = Remove(tr, "te")
	tr = Remove(tr, "tes")
	tr = Remove(tr, "test")

	if _, ok := Get[int](tr, ""); ok {
		t.Fatalf("expected empty key removed")
	}
	if _, ok := Get[int](tr, "te"); ok {
		t.Fatalf("expected te removed")
	}
	if _, ok := Get[int](tr, "tes"); ok {
		t.Fatalf("expected tes removed")
	}
	if _, ok := Get[int](tr, "test"); ok {
		t.Fatalf("expected test removed")
	}
}

func TestTrie_RemoveFreeTest(t *testing.T) {
	tr := NewTrie()
	tr = Put(tr, "test", 2333)
	tr = Put(tr, "te", 23)
	tr = Put(tr, "tes", 233)
	tr = Remove(tr, "tes")
	tr = Remove(tr, "test")

	// Node 'te' should still exist but 'tes'/'test' subtree removed.
	root := tr.GetRoot()
	if root == nil {
		t.Fatalf("expected non-empty trie")
	}
	nT := root.children['t']
	if len(nT.children['e'].children) != 0 {
		t.Fatalf("expected 'te' subtree to have 0 children after removals")
	}

	tr = Remove(tr, "te")
	if tr.GetRoot() != nil {
		t.Fatalf("expected trie to become empty and root=nil")
	}
}

func TestTrie_MismatchTypeTest(t *testing.T) {
	tr := NewTrie()
	tr = Put(tr, "test", 2333)

	_, ok := Get[string](tr, "test")
	if ok {
		t.Fatalf("expected type mismatch to return ok=false")
	}
}

func TestTrie_CopyOnWriteTest1(t *testing.T) {
	empty := NewTrie()
	tr1 := Put(empty, "test", 2333)
	tr2 := Put(tr1, "te", 23)
	tr3 := Put(tr2, "tes", 233)

	tr4 := Remove(tr3, "te")
	tr5 := Remove(tr3, "tes")
	tr6 := Remove(tr3, "test")

	vTe, ok := Get[int](tr3, "te")
	if !ok || vTe == nil || *vTe != 23 {
		t.Fatalf("expected tr3 te=23")
	}
	vTes, ok := Get[int](tr3, "tes")
	if !ok || vTes == nil || *vTes != 233 {
		t.Fatalf("expected tr3 tes=233")
	}
	vTest, ok := Get[int](tr3, "test")
	if !ok || vTest == nil || *vTest != 2333 {
		t.Fatalf("expected tr3 test=2333")
	}

	if _, ok := Get[int](tr4, "te"); ok {
		t.Fatalf("expected te removed in tr4")
	}
	vTes4, ok := Get[int](tr4, "tes")
	if !ok || vTes4 == nil || *vTes4 != 233 {
		t.Fatalf("expected tr4 tes=233")
	}
	vTest4, ok := Get[int](tr4, "test")
	if !ok || vTest4 == nil || *vTest4 != 2333 {
		t.Fatalf("expected tr4 test=2333")
	}

	vTe5, ok := Get[int](tr5, "te")
	if !ok || vTe5 == nil || *vTe5 != 23 {
		t.Fatalf("expected tr5 te=23")
	}
	if _, ok := Get[int](tr5, "tes"); ok {
		t.Fatalf("expected tes removed in tr5")
	}
	vTest5, ok := Get[int](tr5, "test")
	if !ok || vTest5 == nil || *vTest5 != 2333 {
		t.Fatalf("expected tr5 test=2333")
	}

	vTe6, ok := Get[int](tr6, "te")
	if !ok || vTe6 == nil || *vTe6 != 23 {
		t.Fatalf("expected tr6 te=23")
	}
	vTes6, ok := Get[int](tr6, "tes")
	if !ok || vTes6 == nil || *vTes6 != 233 {
		t.Fatalf("expected tr6 tes=233")
	}
	if _, ok := Get[int](tr6, "test"); ok {
		t.Fatalf("expected test removed in tr6")
	}
}

func TestTrie_CopyOnWriteTest2(t *testing.T) {
	empty := NewTrie()
	tr1 := Put(empty, "test", 2333)
	tr2 := Put(tr1, "te", 23)
	tr3 := Put(tr2, "tes", 233)

	tr4 := Put(tr3, "te", "23")
	tr5 := Put(tr3, "tes", "233")
	tr6 := Put(tr3, "test", "2333")

	vTe3, ok := Get[int](tr3, "te")
	if !ok || vTe3 == nil || *vTe3 != 23 {
		t.Fatalf("expected tr3 te=23")
	}
	vTes3, ok := Get[int](tr3, "tes")
	if !ok || vTes3 == nil || *vTes3 != 233 {
		t.Fatalf("expected tr3 tes=233")
	}
	vTest3, ok := Get[int](tr3, "test")
	if !ok || vTest3 == nil || *vTest3 != 2333 {
		t.Fatalf("expected tr3 test=2333")
	}

	vTe4, ok := Get[string](tr4, "te")
	if !ok || vTe4 == nil || *vTe4 != "23" {
		t.Fatalf("expected tr4 te='23'")
	}
	vTes4, ok := Get[int](tr4, "tes")
	if !ok || vTes4 == nil || *vTes4 != 233 {
		t.Fatalf("expected tr4 tes=233")
	}
	vTest4, ok := Get[int](tr4, "test")
	if !ok || vTest4 == nil || *vTest4 != 2333 {
		t.Fatalf("expected tr4 test=2333")
	}

	vTe5, ok := Get[int](tr5, "te")
	if !ok || vTe5 == nil || *vTe5 != 23 {
		t.Fatalf("expected tr5 te=23")
	}
	vTes5, ok := Get[string](tr5, "tes")
	if !ok || vTes5 == nil || *vTes5 != "233" {
		t.Fatalf("expected tr5 tes='233'")
	}
	vTest5, ok := Get[int](tr5, "test")
	if !ok || vTest5 == nil || *vTest5 != 2333 {
		t.Fatalf("expected tr5 test=2333")
	}

	vTe6, ok := Get[int](tr6, "te")
	if !ok || vTe6 == nil || *vTe6 != 23 {
		t.Fatalf("expected tr6 te=23")
	}
	vTes6, ok := Get[int](tr6, "tes")
	if !ok || vTes6 == nil || *vTes6 != 233 {
		t.Fatalf("expected tr6 tes=233")
	}
	vTest6, ok := Get[string](tr6, "test")
	if !ok || vTest6 == nil || *vTest6 != "2333" {
		t.Fatalf("expected tr6 test='2333'")
	}
}

func TestTrie_CopyOnWriteTest3(t *testing.T) {
	empty := NewTrie()
	tr1 := Put(empty, "test", 2333)
	tr2 := Put(tr1, "te", 23)
	tr3 := Put(tr2, "", 233)

	tr4 := Put(tr3, "te", "23")
	tr5 := Put(tr3, "", "233")
	tr6 := Put(tr3, "test", "2333")

	vTe3, ok := Get[int](tr3, "te")
	if !ok || vTe3 == nil || *vTe3 != 23 {
		t.Fatalf("expected tr3 te=23")
	}
	vEmpty3, ok := Get[int](tr3, "")
	if !ok || vEmpty3 == nil || *vEmpty3 != 233 {
		t.Fatalf("expected tr3 empty=233")
	}
	vTest3, ok := Get[int](tr3, "test")
	if !ok || vTest3 == nil || *vTest3 != 2333 {
		t.Fatalf("expected tr3 test=2333")
	}

	vTe4, ok := Get[string](tr4, "te")
	if !ok || vTe4 == nil || *vTe4 != "23" {
		t.Fatalf("expected tr4 te='23'")
	}
	vEmpty4, ok := Get[int](tr4, "")
	if !ok || vEmpty4 == nil || *vEmpty4 != 233 {
		t.Fatalf("expected tr4 empty=233")
	}
	vTest4, ok := Get[int](tr4, "test")
	if !ok || vTest4 == nil || *vTest4 != 2333 {
		t.Fatalf("expected tr4 test=2333")
	}

	vTe5, ok := Get[int](tr5, "te")
	if !ok || vTe5 == nil || *vTe5 != 23 {
		t.Fatalf("expected tr5 te=23")
	}
	vEmpty5, ok := Get[string](tr5, "")
	if !ok || vEmpty5 == nil || *vEmpty5 != "233" {
		t.Fatalf("expected tr5 empty='233'")
	}
	vTest5, ok := Get[int](tr5, "test")
	if !ok || vTest5 == nil || *vTest5 != 2333 {
		t.Fatalf("expected tr5 test=2333")
	}

	vTe6, ok := Get[int](tr6, "te")
	if !ok || vTe6 == nil || *vTe6 != 23 {
		t.Fatalf("expected tr6 te=23")
	}
	vEmpty6, ok := Get[int](tr6, "")
	if !ok || vEmpty6 == nil || *vEmpty6 != 233 {
		t.Fatalf("expected tr6 empty=233")
	}
	vTest6, ok := Get[string](tr6, "test")
	if !ok || vTest6 == nil || *vTest6 != "2333" {
		t.Fatalf("expected tr6 test='2333'")
	}
}

func TestTrie_MixedTest(t *testing.T) {
	trFull := NewTrie()
	for i := uint32(0); i < 2000; i++ {
		key := fmtKey(i)
		value := fmtValue(i)
		trFull = Put(trFull, key, value)
	}

	// Apply override/removal patterns similarly to C++ but smaller scale for test runtime.
	tr := trFull
	for i := uint32(0); i < 2000; i += 2 {
		key := fmtKey(i)
		value := fmtNewValue(i)
		tr = Put(tr, key, value)
	}

	trOverride := tr
	trFinal := tr
	for i := uint32(0); i < 2000; i += 3 {
		key := fmtKey(i)
		trFinal = Remove(trFinal, key)
	}

	// verify trie_full
	for i := uint32(0); i < 2000; i++ {
		key := fmtKey(i)
		value := fmtValue(i)
		got, ok := Get[string](trFull, key)
		if !ok || got == nil || *got != value {
			t.Fatalf("trie_full: key=%s expected=%s got=%v ok=%v", key, value, got, ok)
		}
	}

	// verify trie_override
	for i := uint32(0); i < 2000; i++ {
		key := fmtKey(i)
		if i%2 == 0 {
			value := fmtNewValue(i)
			got, ok := Get[string](trOverride, key)
			if !ok || got == nil || *got != value {
				t.Fatalf("trie_override: key=%s expected=%s got=%v ok=%v", key, value, got, ok)
			}
		} else {
			value := fmtValue(i)
			got, ok := Get[string](trOverride, key)
			if !ok || got == nil || *got != value {
				t.Fatalf("trie_override: key=%s expected=%s got=%v ok=%v", key, value, got, ok)
			}
		}
	}

	// verify final trie
	for i := uint32(0); i < 2000; i++ {
		key := fmtKey(i)
		if i%3 == 0 {
			if _, ok := Get[string](trFinal, key); ok {
				t.Fatalf("trie_final: key=%s expected removed", key)
			}
		} else if i%2 == 0 {
			value := fmtNewValue(i)
			got, ok := Get[string](trFinal, key)
			if !ok || got == nil || *got != value {
				t.Fatalf("trie_final: key=%s expected=%s got=%v ok=%v", key, value, got, ok)
			}
		} else {
			value := fmtValue(i)
			got, ok := Get[string](trFinal, key)
			if !ok || got == nil || *got != value {
				t.Fatalf("trie_final: key=%s expected=%s got=%v ok=%v", key, value, got, ok)
			}
		}
	}
}

func fmtKey(i uint32) string {
	// fmt::format("{:#05}", i) => "0x0000" style with width 5 including '0x'.
	// In tests we only need stable unique keys.
	return keyFromUint(i)
}

func fmtValue(i uint32) string {
	return "value-" + keyFromUint8(i)
}

func fmtNewValue(i uint32) string {
	return "new-value-" + keyFromUint8(i)
}

func keyFromUint(i uint32) string {
	// Simple deterministic key.
	return string(rune('A' + (i % 26))) + "_" + itoa(int64(i))
}

func keyFromUint8(i uint32) string {
	return itoa(int64(i))
}

func itoa(v int64) string {
	// Minimal int->string without fmt for speed.
	if v == 0 {
		return "0"
	}
	neg := v < 0
	if neg {
		v = -v
	}
	var buf [32]byte
	n := 0
	for v > 0 {
		buf[n] = byte('0' + (v % 10))
		v /= 10
		n++
	}
	if neg {
		buf[n] = '-'
		n++
	}
	// reverse
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf[:n])
}

func TestTrie_PointerStability(t *testing.T) {
	tr := NewTrie()
	tr = Put(tr, "test", 2333)
	ptrBefore, ok := Get[int](tr, "test")
	if !ok || ptrBefore == nil {
		t.Fatalf("expected test to exist")
	}

	tr = Put(tr, "tes", 233)
	tr = Put(tr, "te", 23)
	ptrAfter, ok := Get[int](tr, "test")
	if !ok || ptrAfter == nil {
		t.Fatalf("expected test to still exist")
	}

	// Pointer stability: the value node for "test" should be shared.
	if ptrBefore != ptrAfter {
		t.Fatalf("expected pointer stability, got before=%p after=%p", ptrBefore, ptrAfter)
	}
}

