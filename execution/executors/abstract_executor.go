// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// abstract_executor.go
//
// ===----------------------------------------------------------------------===//

package executors

import "github.com/PhanNam1501/bustub-go/execution/plans"

// Executor executes a plan and yields results.
// TODO: define a concrete row/tuple type once storage is ported.
type Executor interface {
	// Init initializes the executor with a context.
	// TODO: include typed context.
	Init(ctx any, plan plans.Plan)

	// Next returns (row, ok). ok=false indicates no more results.
	// TODO: define row type.
	Next() (any, bool)
}

// ---- Concrete executor stubs ----

type SeqScanExecutor struct{}

func (*SeqScanExecutor) Init(ctx any, plan plans.Plan) {}
func (*SeqScanExecutor) Next() (any, bool) { return nil, false }

type InsertExecutor struct{}

func (*InsertExecutor) Init(ctx any, plan plans.Plan) {}
func (*InsertExecutor) Next() (any, bool) { return nil, false }

type UpdateExecutor struct{}

func (*UpdateExecutor) Init(ctx any, plan plans.Plan) {}
func (*UpdateExecutor) Next() (any, bool) { return nil, false }

type DeleteExecutor struct{}

func (*DeleteExecutor) Init(ctx any, plan plans.Plan) {}
func (*DeleteExecutor) Next() (any, bool) { return nil, false }

type HashJoinExecutor struct{}

func (*HashJoinExecutor) Init(ctx any, plan plans.Plan) {}
func (*HashJoinExecutor) Next() (any, bool) { return nil, false }

type NestedLoopJoinExecutor struct{}

func (*NestedLoopJoinExecutor) Init(ctx any, plan plans.Plan) {}
func (*NestedLoopJoinExecutor) Next() (any, bool) { return nil, false }

type AggregationExecutor struct{}

func (*AggregationExecutor) Init(ctx any, plan plans.Plan) {}
func (*AggregationExecutor) Next() (any, bool) { return nil, false }

type LimitExecutor struct{}

func (*LimitExecutor) Init(ctx any, plan plans.Plan) {}
func (*LimitExecutor) Next() (any, bool) { return nil, false }

type SortExecutor struct{}

func (*SortExecutor) Init(ctx any, plan plans.Plan) {}
func (*SortExecutor) Next() (any, bool) { return nil, false }

type TopNExecutor struct{}

func (*TopNExecutor) Init(ctx any, plan plans.Plan) {}
func (*TopNExecutor) Next() (any, bool) { return nil, false }

