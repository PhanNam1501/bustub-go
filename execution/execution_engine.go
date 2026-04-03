// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// execution_engine.go
//
// ExecutionEngine runs a physical plan.
//
// ===----------------------------------------------------------------------===//

package execution

import "github.com/PhanNam1501/bustub-go/execution/plans"

// ExecutionEngine executes plans.
type ExecutionEngine struct {
	// TODO: store executor factories / dependencies.
}

// Execute executes the given plan using the provided context.
//
// TODO: implement Execute
func (ee *ExecutionEngine) Execute(ctx *ExecutorContext, plan plans.Plan) (any, error) {
	return nil, nil
}

