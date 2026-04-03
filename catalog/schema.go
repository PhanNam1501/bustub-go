// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// schema.go
//
// ===----------------------------------------------------------------------===//

package catalog

// Schema is a collection of columns.
type Schema struct {
	Columns []Column
}

// NewSchema constructs a schema.
func NewSchema(cols []Column) *Schema {
	// Copy slice to avoid caller mutation surprises.
	out := make([]Column, len(cols))
	copy(out, cols)
	return &Schema{Columns: out}
}

