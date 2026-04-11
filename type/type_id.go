// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// type_id.go
//
// ===----------------------------------------------------------------------===//

package typ

// TypeID is a data type identifier used in the BusTub SQL type system.
type TypeID int32

const (
	TypeIDInvalid TypeID = -1
	TypeIDInteger TypeID = iota
	TypeIDBoolean
	TypeIDVarchar
)
