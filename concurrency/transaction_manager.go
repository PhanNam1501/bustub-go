// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// transaction_manager.go
//
// ===----------------------------------------------------------------------===//

package concurrency

// TransactionManager manages transaction lifecycle.
// TODO: implement Begin, Commit, Abort.
type TransactionManager struct {
	// TODO: keep track of active transactions, etc.
}

func NewTransactionManager() *TransactionManager { return &TransactionManager{} }

// Begin starts a new transaction.
//
// TODO: Begin
func (tm *TransactionManager) Begin() *Transaction { return &Transaction{} }

// Commit commits a transaction.
//
// TODO: Commit
func (tm *TransactionManager) Commit(txn *Transaction) error { return nil }

// Abort aborts a transaction.
//
// TODO: Abort
func (tm *TransactionManager) Abort(txn *Transaction) error { return nil }

