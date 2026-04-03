// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// optimizer.go
//
// ===----------------------------------------------------------------------===//

package optimizer

import "github.com/PhanNam1501/bustub-go/execution/plans"

// Optimizer optimizes plans.
type Optimizer struct {
	// TODO: store rules or configuration.
}

func NewOptimizer() *Optimizer { return &Optimizer{} }

// Optimize applies optimization rules to the given plan.
//
// TODO: Optimize
func (o *Optimizer) Optimize(inputPlan plans.Plan) (plans.Plan, error) {
	return inputPlan, nil
}

