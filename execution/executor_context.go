// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// executor_context.go
//
// ===----------------------------------------------------------------------===//

package execution

// ExecutorContext contains shared context for executors.
//
// TODO: strongly type fields once catalog/txn/buffer are fully implemented.
type ExecutorContext struct {
	Txn     any
	Catalog any
	BPM     any
}

