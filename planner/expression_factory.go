// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// expression_factory.go
//
// ===----------------------------------------------------------------------===//

package planner

// ExpressionFactory builds expression trees.
// TODO: implement expression nodes and wiring.
type ExpressionFactory struct{}

// NewExpressionFactory creates a new ExpressionFactory.
func NewExpressionFactory() *ExpressionFactory { return &ExpressionFactory{} }

