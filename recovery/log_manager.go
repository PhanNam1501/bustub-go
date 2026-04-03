// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// log_manager.go
//
// ===----------------------------------------------------------------------===//

package recovery

// LogManager manages WAL log records.
// TODO: implement RunFlushThread, AppendLogRecord, etc.
type LogManager struct {
	// TODO: store log state.
}

func NewLogManager() *LogManager { return &LogManager{} }

// AppendLogRecord appends a log record and returns assigned LSN.
//
// TODO: AppendLogRecord
func (lm *LogManager) AppendLogRecord(rec *LogRecord) int32 { return 0 }

