// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// checkpoint_manager.go
//
// ===----------------------------------------------------------------------===//

package recovery

// CheckpointManager manages checkpoints.
// TODO: BeginCheckpoint, EndCheckpoint.
type CheckpointManager struct{}

func NewCheckpointManager() *CheckpointManager { return &CheckpointManager{} }

func (cm *CheckpointManager) BeginCheckpoint() {}
func (cm *CheckpointManager) EndCheckpoint() {}

