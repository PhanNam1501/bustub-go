// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// bound_statement.go
//
// ===----------------------------------------------------------------------===//

package binder

// BoundStatement is the bound form of a SQL statement.
// TODO: refine with specific statement types.
type BoundStatement interface {
	isBoundStatement()
}

type BoundSelect struct {
	// TODO: add select-specific bound fields.
}

func (*BoundSelect) isBoundStatement() {}

type BoundInsert struct {
	// TODO: add insert-specific bound fields.
}

func (*BoundInsert) isBoundStatement() {}

type BoundUpdate struct {
	// TODO: add update-specific bound fields.
}

func (*BoundUpdate) isBoundStatement() {}

type BoundDelete struct {
	// TODO: add delete-specific bound fields.
}

func (*BoundDelete) isBoundStatement() {}

