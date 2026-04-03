// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// lock_manager.go
//
// ===----------------------------------------------------------------------===//

package concurrency

import "github.com/PhanNam1501/bustub-go/common"

// LockManager provides concurrency control primitives.
// TODO: implement LockTable/UnlockTable/LockRow/UnlockRow.
type LockManager struct {
	// TODO: store lock state.
}

func NewLockManager() *LockManager { return &LockManager{} }

func (lm *LockManager) LockTable(txn *Transaction, tableID int32) error { return nil }
func (lm *LockManager) UnlockTable(txn *Transaction, tableID int32) error { return nil }

func (lm *LockManager) LockRow(txn *Transaction, rid common.RID) error { return nil }
func (lm *LockManager) UnlockRow(txn *Transaction, rid common.RID) error { return nil }

