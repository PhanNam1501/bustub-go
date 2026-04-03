// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// transaction.go
//
// ===----------------------------------------------------------------------===//

package concurrency

import "github.com/PhanNam1501/bustub-go/common"

// IsolationLevel is the transaction isolation level.
// TODO: extend with all levels required by BusTub.
type IsolationLevel int

const (
	IsolationLevelSerializable IsolationLevel = iota
	IsolationLevelReadCommitted
)

// Transaction represents a transaction.
// TODO: add fields used by lock manager / MVCC.
type Transaction struct {
	ID common.TxnID

	// TODO: status, timestamps, etc.
}

