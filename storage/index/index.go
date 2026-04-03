// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// index.go
//
// Common Index interface.
//
// ===----------------------------------------------------------------------===//

package index

// Index defines operations over (key -> RID) mappings.
// TODO: replace key/value types with typed SQL `type/`.
type Index interface {
	Insert(key any, rid any) error
	Remove(key any, rid any) error
	GetValue(key any) ([]any, bool)
}

// TODO: later replace `any` with typed key/value/RID types.

