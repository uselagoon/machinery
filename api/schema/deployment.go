package schema

// DeployType .
type DeployType string

// . .
const (
	Branch      DeployType = "BRANCH"
	PullRequest DeployType = "PULLREQUEST"
	Promote     DeployType = "PROMOTE"
)

// DeployEnvironmentLatest is the response.
type DeployEnvironmentLatest struct {
	DeployEnvironmentLatest string `json:"deployEnvironmentLatest"`
}

// DeployEnvironmentPullrequestInput is used as the input for deploying a pull request.
type DeployEnvironmentPullrequestInput struct {
	Project        ProjectInput `json:"project"`
	Number         uint         `json:"number"`
	Title          string       `json:"title"`
	BaseBranchName string       `json:"baseBranchName"`
	BaseBranchRef  string       `json:"baseBranchRef"`
	HeadBranchName string       `json:"headBranchName"`
	HeadBranchRef  string       `json:"headBranchRef"`
	ReturnData     bool         `json:"returnData"`
}

// DeployEnvironmentPullrequest is the response.
type DeployEnvironmentPullrequest struct {
	DeployEnvironmentPullrequest string `json:"deployEnvironmentPullrequest"`
}

// DeployEnvironmentBranchInput is used as the input for deploying a branch.
type DeployEnvironmentBranchInput struct {
	Project    string `json:"project"`
	Branch     string `json:"branch"`
	BranchRef  string `json:"branchRef"`
	ReturnData bool   `json:"returnData"`
}

// DeployEnvironmentBranch is the response.
type DeployEnvironmentBranch struct {
	DeployEnvironmentBranch string `json:"deployEnvironmentBranch"`
}

// DeployEnvironmentPromoteInput is used as the input for promoting one environment to another.
type DeployEnvironmentPromoteInput struct {
	Project                string `json:"project"`
	SourceEnvironment      string `json:"sourceEnvironment"`
	DestinationEnvironment string `json:"destinationEnvironment"`
	ReturnData             bool   `json:"returnData"`
}

// DeployEnvironmentPromote is the response.
type DeployEnvironmentPromote struct {
	DeployEnvironmentPromote string `json:"deployEnvironmentPromote"`
}

// DeployEnvironmentLatestInput is used as the input for deploying an environment.
type DeployEnvironmentLatestInput struct {
	Environment    EnvironmentInput   `json:"environment,omitempty"`
	BulkID         string             `json:"bulkId,omitempty"`
	BulkName       string             `json:"bulkName,omitempty"`
	Priority       int                `json:"priority,omitempty"`
	BuildVariables []EnvKeyValueInput `json:"buildVariables,omitempty"`
	ReturnData     bool               `json:"returnData"`
}

type Deployment struct {
	ID          int              `json:"id,omitempty"`
	Name        string           `json:"name,omitempty"`
	Status      string           `json:"status,omitempty"`
	Created     string           `json:"created,omitempty"`
	Started     string           `json:"started,omitempty"`
	Completed   string           `json:"completed,omitempty"`
	RemoteID    string           `json:"remoteId,omitempty"`
	Logs        string           `json:"logs,omitempty"`
	Environment EnvironmentInput `json:"environment"`
}

type UpdateDeploymentPatchInput struct {
	Name        *string      `json:"name,omitempty"`
	Status      *StatusTypes `json:"status,omitempty"`
	Created     *string      `json:"created,omitempty"`
	Started     *string      `json:"started,omitempty"`
	Completed   *string      `json:"completed,omitempty"`
	Environment *uint        `json:"environment,omitempty"`
	RemoteID    *string      `json:"remoteId,omitempty"`
	Priority    *uint        `json:"priority,omitempty"`
	BulkId      *string      `json:"bulkId,omitempty"`
	BulkName    *string      `json:"bulkName,omitempty"`
}
