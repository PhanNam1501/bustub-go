// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// table_iterator.go
//
// ===----------------------------------------------------------------------===//

package table

// TableIterator iterates over tuples in a table.
// TODO: implement Next() and Get().
type TableIterator struct {
	// TODO: store current page and slot.
}

func (it *TableIterator) Next() bool { return false }

func (it *TableIterator) Get() *Tuple { return nil }

