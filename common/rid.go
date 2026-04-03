// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// rid.go
//
// Record Identifier (RID): (PageID, SlotNum)
//
// ===----------------------------------------------------------------------===//

package common

// RID uniquely identifies a tuple within a page.
type RID struct {
	PageID PageID
	SlotNum int32
}

