//===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// page_guard.go
//
//===----------------------------------------------------------------------===//

package buffer

import (
	"github.com/PhanNam1501/bustub-go/common"
	typ "github.com/PhanNam1501/bustub-go/type"
)

// --- Interfaces ---

type IReadPageGuard interface {
	MoveAssign(that *ReadPageGuard)
	GetPageId() common.PageID
	GetData() []byte
	IsDirty() bool
	Flush()
	Drop()
}

type IWritePageGuard interface {
	MoveAssign(that *WritePageGuard)
	GetPageId() common.PageID
	GetData() []byte
	GetDataMut() []byte
	IsDirty() bool
	Flush()
	Drop()
}

// ============================================================================
// --- ReadPageGuard ---
// ============================================================================

type ReadPageGuard struct {
	bpm     *BufferPoolManager
	frame   *typ.FrameHeader
	isValid bool
}

// NewReadPageGuard khởi tạo một Read Guard.
// Chú ý: Ta truyền thẳng BufferPoolManager vào thay vì truyền lẻ tẻ từng thành phần.
func NewReadPageGuard(bpm *BufferPoolManager, frame *typ.FrameHeader) *ReadPageGuard {
	return &ReadPageGuard{
		bpm:     bpm,
		frame:   frame,
		isValid: true,
	}
}

func (g *ReadPageGuard) invalidate() {
	g.isValid = false
	g.frame = nil
	g.bpm = nil
}

func (g *ReadPageGuard) MoveAssign(that *ReadPageGuard) {
	if g == that {
		return
	}
	if g.isValid {
		g.Drop()
	}
	if that != nil && that.isValid {
		g.bpm = that.bpm
		g.frame = that.frame
		g.isValid = true
		that.invalidate()
	}
}

func (g *ReadPageGuard) GetPageId() common.PageID {
	if !g.isValid {
		panic("BUSTUB_ENSURE: tried to use an invalid read guard")
	}
	return g.frame.PageID
}

func (g *ReadPageGuard) GetData() []byte {
	if !g.isValid {
		panic("BUSTUB_ENSURE: tried to use an invalid read guard")
	}
	return g.frame.Data
}

func (g *ReadPageGuard) IsDirty() bool {
	if !g.isValid {
		panic("BUSTUB_ENSURE: tried to use an invalid read guard")
	}
	return g.frame.IsDirty
}

func (g *ReadPageGuard) Flush() {
	if !g.isValid {
		return
	}
	// Thay vì Guard tự lập lịch đĩa, hãy để BPM làm việc của nó!
	g.bpm.FlushPage(g.frame.PageID)
}

func (g *ReadPageGuard) Drop() {
	if !g.isValid {
		return
	}

	// Giải phóng Read Lock của Page
	g.frame.Mu.RUnlock()

	// Gọi hàm Helper unpin của BPM (Đã bao gồm lock bpmLatch và cập nhật replacer)
	g.bpm.unpinPage(g.frame.PageID, false)

	g.invalidate()
}

// ============================================================================
// --- WritePageGuard ---
// ============================================================================

type WritePageGuard struct {
	bpm     *BufferPoolManager
	frame   *typ.FrameHeader
	isValid bool
}

func NewWritePageGuard(bpm *BufferPoolManager, frame *typ.FrameHeader) *WritePageGuard {
	return &WritePageGuard{
		bpm:     bpm,
		frame:   frame,
		isValid: true,
	}
}

func (g *WritePageGuard) invalidate() {
	g.isValid = false
	g.frame = nil
	g.bpm = nil
}

func (g *WritePageGuard) MoveAssign(that *WritePageGuard) {
	if g == that {
		return
	}
	if g.isValid {
		g.Drop()
	}
	if that != nil && that.isValid {
		g.bpm = that.bpm
		g.frame = that.frame
		g.isValid = true
		that.invalidate()
	}
}

func (g *WritePageGuard) GetPageId() common.PageID {
	if !g.isValid {
		panic("BUSTUB_ENSURE: tried to use an invalid write guard")
	}
	return g.frame.PageID
}

func (g *WritePageGuard) GetData() []byte {
	if !g.isValid {
		panic("BUSTUB_ENSURE: tried to use an invalid write guard")
	}
	return g.frame.Data
}

func (g *WritePageGuard) GetDataMut() []byte {
	if !g.isValid {
		panic("BUSTUB_ENSURE: tried to use an invalid write guard")
	}
	g.frame.IsDirty = true
	return g.frame.Data
}

func (g *WritePageGuard) IsDirty() bool {
	if !g.isValid {
		panic("BUSTUB_ENSURE: tried to use an invalid write guard")
	}
	return g.frame.IsDirty
}

func (g *WritePageGuard) Flush() {
	if !g.isValid {
		return
	}
	g.bpm.FlushPage(g.frame.PageID)
}

func (g *WritePageGuard) Drop() {
	if !g.isValid {
		return
	}

	g.frame.Mu.Unlock()

	g.bpm.unpinPage(g.frame.PageID, g.frame.IsDirty)

	g.invalidate()
}
