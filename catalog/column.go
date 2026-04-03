// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// column.go
//
// ===----------------------------------------------------------------------===//

package catalog

// Column describes one column in a schema.
type Column struct {
	Name string
	// TODO: replace with proper type system from `type/`.
	Type any
}

