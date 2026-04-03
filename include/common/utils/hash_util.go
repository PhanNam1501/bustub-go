package utils

import (
	"encoding/binary"
	"math"

	"github.com/PhanNam1501/bustub-go/include/types"
)

// Dùng uint64 thay cho std::size_t (hash_t) trên C++
type HashT uint64

const PrimeFactor HashT = 10000019

// HashBytes là thuật toán băm cốt lõi (dựa trên Greenplum)
func HashBytes(bytes []byte) HashT {
	hash := HashT(len(bytes))
	for _, b := range bytes {
		// Ép kiểu byte sang int8 trước để giữ nguyên logic bitwise của C++
		val := HashT(int8(b))
		hash = ((hash << 5) ^ (hash >> 27)) ^ val
	}
	return hash
}

// CombineHashes gộp 2 mã băm lại với nhau
func CombineHashes(l, r HashT) HashT {
	buf := make([]byte, 16)
	binary.LittleEndian.PutUint64(buf[0:8], uint64(l))
	binary.LittleEndian.PutUint64(buf[8:16], uint64(r))
	return HashBytes(buf)
}

// SumHashes cộng 2 mã băm (có chia lấy dư)
func SumHashes(l, r HashT) HashT {
	return (l%PrimeFactor + r%PrimeFactor) % PrimeFactor
}

// hashInt64 là hàm hỗ trợ chuyển int64 thành byte để băm
func hashInt64(val int64) HashT {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(val))
	return HashBytes(buf)
}

// hashFloat64 là hàm hỗ trợ băm số thập phân (double)
func hashFloat64(val float64) HashT {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, math.Float64bits(val))
	return HashBytes(buf)
}

// HashValue băm một đối tượng Value của BusTub
func HashValue(val types.Value) HashT { // <--- Thêm types.
	if val.IsNull() {
		return 0
	}

	switch val.GetTypeId() {
	// <--- Thêm types. cho toàn bộ enum
	case types.TINYINT, types.SMALLINT, types.INTEGER, types.BIGINT:
		var raw int64

		// Giả sử bạn sử dụng getter GetValue() hoặc trường Data viết hoa
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
