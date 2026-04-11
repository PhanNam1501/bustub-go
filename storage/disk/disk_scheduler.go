// ===----------------------------------------------------------------------===//
//
//                        BusTub (Go port)
//
// disk_scheduler.go
//
// ===----------------------------------------------------------------------===//

package disk

import (
	"sync"

	"github.com/PhanNam1501/bustub-go/common"
)

type PageID = common.PageID

type DiskRequest struct {
	IsWrite bool
	Data    []byte
	PageID  PageID

	Callback chan bool
}

type DiskScheduler struct {
	diskManager *DiskManager

	requestQueue chan *DiskRequest

	wg sync.WaitGroup
}

func NewDiskScheduler(diskManager *DiskManager) *DiskScheduler {
	ds := &DiskScheduler{
		diskManager:  diskManager,
		requestQueue: make(chan *DiskRequest, 1024),
	}

	ds.wg.Add(1)
	go ds.StartWorkerThread()

	return ds
}

func (ds *DiskScheduler) Close() {
	close(ds.requestQueue)
	ds.wg.Wait()
}

func (ds *DiskScheduler) Schedule(requests []*DiskRequest) {
	for _, rq := range requests {
		ds.requestQueue <- rq
	}
}

func (ds *DiskScheduler) StartWorkerThread() {
	defer ds.wg.Done()

	for rq := range ds.requestQueue {
		if rq.IsWrite {
			ds.diskManager.WritePage(rq.PageID, rq.Data)
		} else {
			ds.diskManager.ReadPage(rq.PageID, rq.Data)
		}

		if rq.Callback != nil {
			rq.Callback <- true
		}
	}
}

func (ds *DiskScheduler) CreatePromise() chan bool {
	return make(chan bool, 1)
}

func (ds *DiskScheduler) DeallocatePage(pageID PageID) {
	ds.diskManager.DeletePage(pageID)
}
