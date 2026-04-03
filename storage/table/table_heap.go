// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// table_heap.go
//
// ===----------------------------------------------------------------------===//

package table

// TableHeap stores table tuples across pages.
// TODO: implement linkage with TablePages and iterator.
type TableHeap struct {
	// TODO: keep track of first/last page ids.
}

// InsertTuple inserts a tuple into the table.
//
// TODO: InsertTuple
func (th *TableHeap) InsertTuple(t *Tuple) (/*rid*/ any, error) {
	return nil, nil
}

// DeleteTuple deletes a tuple (mark delete).
//
// TODO: MarkDelete
func (th *TableHeap) DeleteTuple(rid any) error { return nil }

