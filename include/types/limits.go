package types

import "math"

const (
	BUSTUB_INT8_MIN  = math.MinInt8 + 1
	BUSTUB_INT16_MIN = math.MinInt16 + 1
	BUSTUB_INT32_MIN = math.MinInt32 + 1
	BUSTUB_INT64_MIN = math.MinInt64 + 1

	BUSTUB_INT32_MAX = math.MaxInt32
	BUSTUB_INT64_MAX = math.MaxInt64

	BUSTUB_VALUE_NULL = math.MaxUint32

	// Sentinel NULL values
	BUSTUB_INT32_NULL = math.MinInt32
	BUSTUB_INT64_NULL = math.MinInt64
)
