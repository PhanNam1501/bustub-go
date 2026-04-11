//===----------------------------------------------------------------------===//
//
//                        BusTub (Go port)
//
// disk_manager.go
//
//===----------------------------------------------------------------------===//

package disk

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Basic constants and data types
const PageSize = 4096 // Equivalent to BUSTUB_PAGE_SIZE

// DiskManager manages reading and writing pages to physical files.
type DiskManager struct {
	mu sync.Mutex // Equivalent to std::scoped_lock

	dbFileName  string
	logFileName string

	dbFile  *os.File
	logFile *os.File

	pages        map[PageID]int64 // Map from PageID to Offset (in bytes)
	freeSlots    []int64          // List of free offsets (from deleted pages)
	pageCapacity int64            // Current capacity (number of pages)

	numWrites  int
	numDeletes int
	numFlushes int
	flushLog   bool
}

// NewDiskManager creates a new DiskManager and initializes the necessary files.
func NewDiskManager(dbFilePath string) (*DiskManager, error) {
	// Get log file name from db file name (remove extension)
	ext := filepath.Ext(dbFilePath)
	base := strings.TrimSuffix(filepath.Base(dbFilePath), ext)
	logFilePath := filepath.Join(filepath.Dir(dbFilePath), base+".log")

	// Open or create the log file
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	// Open or create the database file
	dbFile, err := os.OpenFile(dbFilePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		logFile.Close()
		return nil, err
	}

	dm := &DiskManager{
		dbFileName:   dbFilePath,
		logFileName:  logFilePath,
		dbFile:       dbFile,
		logFile:      logFile,
		pages:        make(map[PageID]int64),
		freeSlots:    make([]int64, 0),
		pageCapacity: 1024, // Default initial capacity
	}

	// Initialize the DB file size
	targetSize := (dm.pageCapacity + 1) * PageSize
	if err := dbFile.Truncate(targetSize); err != nil {
		return nil, err
	}

	return dm, nil
}

// ShutDown closes all file resources safely.
func (dm *DiskManager) ShutDown() {
	dm.mu.Lock()
	defer dm.mu.Unlock() // Hold lock while closing files for safety

	if dm.dbFile != nil {
		dm.dbFile.Close()
	}
	if dm.logFile != nil {
		dm.logFile.Close()
	}
}

// WritePage writes the contents of pageData to the disk file.
// pageData must be exactly PageSize (4096 bytes) long.
func (dm *DiskManager) WritePage(pageID PageID, pageData []byte) {
	if len(pageData) != PageSize {
		log.Printf("I/O error: Write data does not match PageSize")
		return
	}

	dm.mu.Lock()
	defer dm.mu.Unlock()

	var offset int64
	if off, exists := dm.pages[pageID]; exists {
		offset = off // Page already exists, overwrite
	} else {
		offset = dm.allocatePage() // Page does not exist, allocate new space
	}

	// Write directly at offset using WriteAt (no seek needed)
	_, err := dm.dbFile.WriteAt(pageData, offset)
	if err != nil {
		log.Printf("I/O error while writing page %d: %v", pageID, err)
		return
	}

	dm.numWrites++
	dm.pages[pageID] = offset

	// Flush to disk (Equivalent to db_io_.flush())
	dm.dbFile.Sync()
}

// ReadPage reads the contents from disk into pageData.
func (dm *DiskManager) ReadPage(pageID PageID, pageData []byte) {
	if len(pageData) != PageSize {
		log.Printf("I/O error: Read buffer does not match PageSize")
		return
	}

	dm.mu.Lock()
	defer dm.mu.Unlock()

	var offset int64
	if off, exists := dm.pages[pageID]; exists {
		offset = off
	} else {
		offset = dm.allocatePage()
	}

	// Check if we are reading past the end of the file
	info, err := dm.dbFile.Stat()
	if err != nil {
		log.Printf("I/O error: Cannot get db file size")
		return
	}
	if offset > info.Size() {
		log.Printf("I/O error: Read page %d past the end of file at offset %d", pageID, offset)
		return
	}

	dm.pages[pageID] = offset

	// Read data using ReadAt
	n, err := dm.dbFile.ReadAt(pageData, offset)
	if err != nil && err != io.EOF {
		log.Printf("I/O error while reading page %d: %v", pageID, err)
		return
	}

	// If the file ends before a full page, fill the rest with zeros (Zero-padding)
	if n < PageSize {
		log.Printf("I/O error: Read page %d hit EOF, missing %d bytes", pageID, PageSize-n)
		for i := n; i < PageSize; i++ {
			pageData[i] = 0
		}
	}
}

// DeletePage cleans up the page and adds its offset to the free list.
func (dm *DiskManager) DeletePage(pageID PageID) {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	offset, exists := dm.pages[pageID]
	if !exists {
		return
	}

	dm.freeSlots = append(dm.freeSlots, offset)
	delete(dm.pages, pageID)
	dm.numDeletes++
}

// WriteLog writes the log sequentially to the file.
func (dm *DiskManager) WriteLog(logData []byte) {
	if len(logData) == 0 {
		return
	}

	dm.flushLog = true

	// Write sequentially to the end of the file
	_, err := dm.logFile.Write(logData)
	if err != nil {
		log.Printf("I/O error while writing log: %v", err)
		return
	}

	dm.numFlushes++
	dm.logFile.Sync()
	dm.flushLog = false
}

// ReadLog reads the log from a specific offset.
func (dm *DiskManager) ReadLog(logData []byte, offset int64) bool {
	info, err := dm.logFile.Stat()
	if err != nil || offset >= info.Size() {
		return false
	}

	n, err := dm.logFile.ReadAt(logData, offset)
	if err != nil && err != io.EOF {
		log.Printf("I/O error while reading log: %v", err)
		return false
	}

	// Zero-padding if the file is shorter than the requested size
	if n < len(logData) {
		for i := n; i < len(logData); i++ {
			logData[i] = 0
		}
	}

	return true
}

// Getters (Retrieve statistics)
func (dm *DiskManager) GetNumFlushes() int  { return dm.numFlushes }
func (dm *DiskManager) GetFlushState() bool { return dm.flushLog }
func (dm *DiskManager) GetNumWrites() int   { return dm.numWrites }
func (dm *DiskManager) GetNumDeletes() int  { return dm.numDeletes }

// allocatePage allocates space on the disk file.
// Note: Do not use Mutex here as the caller (WritePage/ReadPage) already holds the Lock.
func (dm *DiskManager) allocatePage() int64 {
	// Prioritize reusing free slots
	if len(dm.freeSlots) > 0 {
		lastIdx := len(dm.freeSlots) - 1
		offset := dm.freeSlots[lastIdx]
		dm.freeSlots = dm.freeSlots[:lastIdx]
		return offset
	}

	// If out of space, expand the disk file (Truncate)
	if int64(len(dm.pages))+1 >= dm.pageCapacity {
		dm.pageCapacity *= 2
		targetSize := (dm.pageCapacity + 1) * PageSize
		dm.dbFile.Truncate(targetSize)
	}

	// Based on C++ logic: offset = current number of pages * PageSize
	return int64(len(dm.pages)) * PageSize
}
