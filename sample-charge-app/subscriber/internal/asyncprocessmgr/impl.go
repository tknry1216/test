package asyncprocessmgr

type AsyncProcessManager interface {
	RegisterProcess(processID string, process Process) error
	GetProcess(processID string) (*Process, error)
	UpdateProcessStatus(processID string, status ProcessStatus) error
}

type asyncProcessManagerImpl struct {
	asyncProcessManager AsyncProcessManager
}

func NewAsyncProcessManagerImpl(asyncProcessManager AsyncProcessManager) *asyncProcessManagerImpl {
	return &asyncProcessManagerImpl{asyncProcessManager: asyncProcessManager}
}

func (s *asyncProcessManagerImpl) RegisterProcess(processID string, process Process) error {
	return s.asyncProcessManager.RegisterProcess(processID, process)
}

func (s *asyncProcessManagerImpl) GetProcess(processID string) (*Process, error) {
	return s.asyncProcessManager.GetProcess(processID)
}

func (s *asyncProcessManagerImpl) UpdateProcessStatus(processID string, status ProcessStatus) error {
	return s.asyncProcessManager.UpdateProcessStatus(processID, status)
}
