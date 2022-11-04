package schema

type LagoonMessage struct {
	Type      string         `json:"type,omitempty"`
	Namespace string         `json:"namespace,omitempty"`
	Meta      *LagoonLogMeta `json:"meta,omitempty"`
	// BuildInfo *LagoonBuildInfo `json:"buildInfo,omitempty"`
}

// LagoonLogMeta is the metadata that is used by logging in Lagoon.
type LagoonLogMeta struct {
	BranchName     string          `json:"branchName,omitempty"`
	BuildName      string          `json:"buildName,omitempty"`
	BuildStatus    string          `json:"buildStatus,omitempty"`
	BuildPhase     string          `json:"buildPhase,omitempty"` // legacy but still used until `buildStatus` is more available
	BuildStep      string          `json:"buildStep,omitempty"`
	EndTime        string          `json:"endTime,omitempty"`
	Environment    string          `json:"environment,omitempty"`
	EnvironmentID  *uint           `json:"environmentId,omitempty"`
	JobName        string          `json:"jobName,omitempty"`   // used by tasks/jobs
	JobStatus      string          `json:"jobStatus,omitempty"` // used by tasks/jobs
	JobStep        string          `json:"jobStep,omitempty"`   // used by tasks/jobs
	LogLink        string          `json:"logLink,omitempty"`
	MonitoringURLs []string        `json:"monitoringUrls,omitempty"`
	Project        string          `json:"project,omitempty"`
	ProjectID      *uint           `json:"projectId,omitempty"`
	ProjectName    string          `json:"projectName,omitempty"`
	RemoteID       string          `json:"remoteId,omitempty"`
	Route          string          `json:"route,omitempty"`
	Routes         []string        `json:"routes,omitempty"`
	StartTime      string          `json:"startTime,omitempty"`
	Services       []string        `json:"services,omitempty"`
	Task           *LagoonTaskInfo `json:"task,omitempty"`
	Key            string          `json:"key,omitempty"`
	AdvancedData   string          `json:"advancedData,omitempty"`
	Cluster        string          `json:"clusterName,omitempty"`
}

// LagoonTaskInfo defines what a task can use to communicate with Lagoon via SSH/API.
type LagoonTaskInfo struct {
	ID       string `json:"id"` // should be int, but the api sends it as a string :\
	Name     string `json:"name,omitempty"`
	TaskName string `json:"taskName,omitempty"`
	Service  string `json:"service,omitempty"`
	Command  string `json:"command,omitempty"`
	SSHHost  string `json:"sshHost,omitempty"`
	SSHPort  string `json:"sshPort,omitempty"`
	APIHost  string `json:"apiHost,omitempty"`
}
