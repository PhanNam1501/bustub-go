// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// planner.go
//
// ===----------------------------------------------------------------------===//

package planner

import "github.com/PhanNam1501/bustub-go/execution/plans"

// Planner turns a bound statement into an executable plan.
type Planner struct {
	// TODO: dependency injection (binder output, catalog, etc.)
}

func NewPlanner() *Planner { return &Planner{} }

// PlanQuery plans a query and returns an executable plan.
//
// TODO: PlanQuery / PlanSelect / PlanInsert ...
func (p *Planner) PlanQuery(boundQuery any) (plans.Plan, error) {
	return nil, nil
}

