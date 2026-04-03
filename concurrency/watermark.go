// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// watermark.go
//
// Watermark helps with MVCC/snapshot management.
//
// ===----------------------------------------------------------------------===//

package concurrency

// Watermark tracks transaction visibility.
// TODO: implement AddTxn/RemoveTxn/GetWatermark.
type Watermark struct {
	// TODO: track active transaction ids or timestamps.
}

func NewWatermark() *Watermark { return &Watermark{} }

func (w *Watermark) AddTxn(txnID int32) {}
func (w *Watermark) RemoveTxn(txnID int32) {}
func (w *Watermark) GetWatermark() int32 { return 0 }

