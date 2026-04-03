// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// value.go
//
// Typed values used by expressions/executors.
//
// ===----------------------------------------------------------------------===//

package typ

// Value represents one typed SQL value.
// TODO: implement storage and serialization precisely.
type Value struct {
	Type TypeID
	Data any
}

// CompareEquals checks equality.
//
// TODO: implement CompareEquals
func (v Value) CompareEquals(other Value) bool { return false }

// CompareLessThan checks ordering.
//
// TODO: implement CompareLessThan
func (v Value) CompareLessThan(other Value) bool { return false }

// Serialize serializes the value.
//
// TODO: implement Serialize
func (v Value) Serialize() []byte { return nil }

