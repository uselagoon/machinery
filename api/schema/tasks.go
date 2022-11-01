package schema

// ProjectAvailability determines the number of pods used to run a project.
type StatusTypes string

// High tells Lagoon to load balance across multiple pods.
// Standard tells Lagoon to use a single pod for the site.
const (
	Active    StatusTypes = "ACTIVE"
	Succeeded StatusTypes = "SUCCEEDED"
	Failed    StatusTypes = "FAILED"
	New       StatusTypes = "NEW"
	Pending   StatusTypes = "PENDING"
	Running   StatusTypes = "RUNNING"
	Cancelled StatusTypes = "CANCELLED"
	Error     StatusTypes = "ERROR"
	Complete  StatusTypes = "COMPLETE"
)

// Task is based on the Lagoon API type.
type Task struct {
	ID          uint        `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Status      string      `json:"status,omitempty"`
	Created     string      `json:"created,omitempty"`
	Started     string      `json:"started,omitempty"`
	Completed   string      `json:"completed,omitempty"`
	Service     string      `json:"service,omitempty"`
	Command     string      `json:"command,omitempty"`
	RemoteID    string      `json:"remoteId,omitempty"`
	Logs        string      `json:"logs,omitempty"`
	Environment Environment `json:"environment,omitempty"`
}

type ActiveStandbyResult struct {
	Status                       string `json:"status"`
	ProductionEnvironment        string `json:"productionEnvironment"`
	StandbyProductionEnvironment string `json:"standbyProductionEnvironment"`
	ProductionRoutes             string `json:"productionRoutes"`
	StandbyRoutes                string `json:"standbyRoutes"`
}

type UpdateTaskPatchInput struct {
	Name        string      `json:"name,omitempty"`
	TaskName    string      `json:"taskName,omitempty"`
	Status      StatusTypes `json:"status,omitempty"`
	Created     string      `json:"created,omitempty"`
	Started     string      `json:"started,omitempty"`
	Completed   string      `json:"completed,omitempty"`
	Environment *uint       `json:"environment,omitempty"`
	Service     string      `json:"service,omitempty"`
	Command     string      `json:"command,omitempty"`
	RemoteID    string      `json:"remoteId,omitempty"`
}
