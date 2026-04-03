package types

type CmpBool int

const (
	CmpFalse CmpBool = iota
	CmpTrue
	CmpNull
)

// Type định nghĩa hành vi cho từng loại dữ liệu
type Type interface {
	GetTypeId() TypeId
	CompareEquals(left, right Value) CmpBool
	CompareLessThan(left, right Value) CmpBool
	Add(left, right Value) Value
	SerializeTo(v Value) []byte
	DeserializeFrom(data []byte) Value
	ToString(v Value) string
}

// Registry để quản lý các singleton types (giống k_types trong C++)
var kTypes = make(map[TypeId]Type)

func GetInstance(id TypeId) Type {
	return kTypes[id]
}
