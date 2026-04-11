package typ

import (
	"sync"

	"github.com/PhanNam1501/bustub-go/common"
)

type FrameHeader struct {
	FrameID  common.FrameID
	PageID   common.PageID
	Data     []byte
	PinCount int32
	IsDirty  bool
	Mu       sync.RWMutex
}

func NewFrameHeader(frameID common.FrameID) *FrameHeader {
	h := &FrameHeader{
		FrameID: frameID,
		PageID:  common.InvalidPageID,
		Data:    make([]byte, common.BustubPageSize),
	}
	//h.Reset()
	return h
}

func (h *FrameHeader) Reset() {
	for i := range h.Data {
		h.Data[i] = 0
	}
	h.PageID = common.InvalidPageID
	h.PinCount = 0
	h.IsDirty = false
}

func (h *FrameHeader) GetData() []byte {
	return h.Data
}
