// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// catalog.go
//
// Catalog stores schema/tables/index metadata.
//
// ===----------------------------------------------------------------------===//

package catalog

// Catalog is a placeholder catalog implementation.
// TODO: implement CreateTable, GetTable, CreateIndex, etc.
type Catalog struct {
	// TODO: store schemas, tables and indexes.
}

// NewCatalog creates a new Catalog.
func NewCatalog() *Catalog { return &Catalog{} }

// CreateTable creates a new table with the given name and schema.
//
// TODO: CreateTable
func (c *Catalog) CreateTable(name string, schema *Schema) error { return nil }

// GetTable returns a table schema by name.
//
// TODO: GetTable
func (c *Catalog) GetTable(name string) (*Schema, bool) { return nil, false }

// CreateIndex creates an index on a table.
//
// TODO: CreateIndex
func (c *Catalog) CreateIndex(name string) error { return nil }

