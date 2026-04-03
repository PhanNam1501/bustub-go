// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// exception.go
//
// Error/exception types.
//
// ===----------------------------------------------------------------------===//

package common

import "errors"

// ErrNotImplemented indicates a feature is not implemented yet.
var ErrNotImplemented = errors.New("not implemented")

// ErrInvalidType indicates a value/type mismatch.
var ErrInvalidType = errors.New("invalid type")

