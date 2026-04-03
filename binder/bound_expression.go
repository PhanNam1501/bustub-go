// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// bound_expression.go
//
// ===----------------------------------------------------------------------===//

package binder

// BoundExpression is the bound form of an expression.
// TODO: refine with concrete expression node types.
type BoundExpression interface {
	isBoundExpression()
}

// BoundColumnRef is a bound column reference expression.
type BoundColumnRef struct {
	// TODO: store column identity (table/column ids) as in Bustub.
}

func (*BoundColumnRef) isBoundExpression() {}

// BoundConstant is a bound constant literal expression.
type BoundConstant struct {
	// TODO: store typed literal value.
}

func (*BoundConstant) isBoundExpression() {}

