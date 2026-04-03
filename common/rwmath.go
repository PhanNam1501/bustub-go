// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// rwmath.go
//
// Small math helpers used across the system.
//
// ===----------------------------------------------------------------------===//

package common

// Min returns the minimum of two ordered values.
// TODO: switch to generic implementation when appropriate for the course.
func Min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two ordered values.
func Max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

