package types

import (
	"fmt"
)

type Value struct {
	typeId TypeId
	value  interface{} // Thay thế cho union Val trong C++
	size   uint32      // Độ dài cho Varlen (VARCHAR)
}

// Constructor cho Integer
func NewIntValue(val int32) Value {
	return Value{
		typeId: INTEGER,
		value:  val,
		size:   0,
	}
}

// Constructor cho Varchar
func NewVarcharValue(val string) Value {
	return Value{
		typeId: VARCHAR,
		value:  val,
		size:   uint32(len(val)),
	}
}

func (v Value) IsNull() bool {
	return v.size == BUSTUB_VALUE_NULL
}

func (v Value) GetTypeId() TypeId {
	return v.typeId
}

// THÊM HÀM NÀY VÀO ĐÂY:
func (v Value) GetValue() any {
	return v.value
}

// Ủy quyền xử lý cho Type tương ứng (giống logic C++)
func (v Value) CompareEquals(other Value) CmpBool {
	t := GetInstance(v.typeId)
	return t.CompareEquals(v, other)
}

func (v Value) Add(other Value) Value {
	t := GetInstance(v.typeId)
	return t.Add(v, other)
}

func (v Value) ToString() string {
	if v.IsNull() {
		return "NULL"
	}
	return fmt.Sprintf("%v", v.value)
}
