package asyncprocessmgr

type Process struct {
	ID          string
	Name        string
	Description string
	Attempts    int
	Status      string
}

type ProcessStatus string

const (
	ProcessStatusPending   ProcessStatus = "PENDING"
	ProcessStatusRunning   ProcessStatus = "RUNNING"
	ProcessStatusCompleted ProcessStatus = "COMPLETED"
	ProcessStatusFailed    ProcessStatus = "FAILED"
)
