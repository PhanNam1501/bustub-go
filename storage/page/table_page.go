// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// table_page.go
//
// Slotted page layout for storing tuples.
//
// ===----------------------------------------------------------------------===//

package page

import "github.com/PhanNam1501/bustub-go/storage/table"

// TablePage stores tuples in a slotted page format.
// TODO: implement InsertTuple and MarkDelete.
type TablePage struct {
	// TODO
}

func (tp *TablePage) InsertTuple(tuple *table.Tuple) error { return nil }

func (tp *TablePage) MarkDelete(slotNum int32) error { return nil }

