package utils

import (
	"encoding/binary"
	"math"

	"github.com/PhanNam1501/bustub-go/include/types"
)

type HashT uint64

const PrimeFactor HashT = 10000019

func HashBytes(bytes []byte) HashT {
	hash := HashT(len(bytes))
	for _, b := range bytes {
		val := HashT(int8(b))
		hash = ((hash << 5) ^ (hash >> 27)) ^ val
	}
	return hash
}

func CombineHashes(l, r HashT) HashT {
	buf := make([]byte, 16)
	binary.LittleEndian.PutUint64(buf[0:8], uint64(l))
	binary.LittleEndian.PutUint64(buf[8:16], uint64(r))
	return HashBytes(buf)
}

func SumHashes(l, r HashT) HashT {
	return (l%PrimeFactor + r%PrimeFactor) % PrimeFactor
}

func hashInt64(val int64) HashT {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(val))
	return HashBytes(buf)
}

func hashFloat64(val float64) HashT {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, math.Float64bits(val))
	return HashBytes(buf)
}

func HashValue(val types.Value) HashT {
	if val.IsNull() {
		return 0
	}

	switch val.GetTypeId() {
	case types.TINYINT, types.SMALLINT, types.INTEGER, types.BIGINT:
		var raw int64

		switch v := val.GetValue().(type) {
		case int8:
			raw = int64(v)
		case int16:
			raw = int64(v)
		case int32:
			raw = int64(v)
		case int64:
			raw = v
		case int:
			raw = int64(v)
		}
		return hashInt64(raw)

	case types.BOOLEAN:
		buf := make([]byte, 1)
		if val.GetValue().(bool) {
			buf[0] = 1
		}
		return HashBytes(buf)

	case types.DECIMAL:
		return hashFloat64(val.GetValue().(float64))

	case types.VARCHAR:
		str := val.GetValue().(string)
		return HashBytes([]byte(str))

	case types.TIMESTAMP:
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, val.GetValue().(uint64))
		return HashBytes(buf)

	default:
		panic("Unsupported type for hashing")
	}
}
