// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// binder.go
//
// Binder transforms parsed SQL AST into bound plan/expressions.
//
// ===----------------------------------------------------------------------===//

package binder

// BoundQuery is a placeholder for a bound representation of a query.
// TODO: replace `any` with proper bound AST types once the SQL parser is ported.
type BoundQuery = any

// BindQuery binds a parsed query AST into a bound representation.
//
// TODO: BindQuery
func BindQuery(parsedQuery any) (BoundQuery, error) {
	// Stub: keep compilation green.
	return nil, nil
}

