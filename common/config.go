// ===----------------------------------------------------------------------===//
//
//                         BusTub
//
// config.go
//
// Identification: common/config.go
//
// Copyright (c) 2015-2025, Carnegie Mellon University Database Group
//
// ===----------------------------------------------------------------------===//

package common

const (
	// InvalidFrameID is the sentinel value for an invalid frame id.
	InvalidFrameID = -1
	// InvalidPageID is the sentinel value for an invalid page id.
	InvalidPageID = -1
	// InvalidTxnID is the sentinel value for an invalid transaction id.
	InvalidTxnID = -1
	// InvalidLSN is the sentinel value for an invalid log sequence number.
	InvalidLSN = -1

	// BustubPageSize is the size of a data page in bytes.
	BustubPageSize = 8192
	// BufferPoolSize is the default size of the buffer pool.
	BufferPoolSize = 128
	// DefaultDBIOSize is the starting size of the file on disk.
	DefaultDBIOSize = 16
)

// PageID is the type for page IDs.
type PageID int32

// FrameID is the type for frame IDs (index into the buffer pool).
type FrameID int32

// TxnID is the type for transaction IDs.
type TxnID int32

// LSN is the type for log sequence numbers.
type LSN int32
