// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// bound_table_ref.go
//
// ===----------------------------------------------------------------------===//

package binder

// BoundTableRef is the bound representation of a table reference.
// TODO: refine with specific table/join semantics.
type BoundTableRef interface {
	isBoundTableRef()
}

type BoundBaseTableRef struct {
	// TODO: store table identity / alias / schema.
}

func (*BoundBaseTableRef) isBoundTableRef() {}

type BoundJoinRef struct {
	// TODO: store join type, left/right refs, join predicate.
}

func (*BoundJoinRef) isBoundTableRef() {}

