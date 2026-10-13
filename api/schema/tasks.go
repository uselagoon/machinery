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
	Files       []Files     `json:"files,omitempty"`
}

type Files struct {
	Filename string `json:"filename,omitempty"`
}

type FileUploadForm struct {
	PostUrl    string            `json:"postUrl,omitempty"`
	FormFields map[string]string `json:"formFields,omitempty"`
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

type TaskPermission string

const (
	Guest      TaskPermission = "GUEST"
	Developer  TaskPermission = "DEVELOPER"
	Maintainer TaskPermission = "MAINTAINER"
)

type AdvancedTaskDefinitionArgumentType string

const (
	String                           AdvancedTaskDefinitionArgumentType = "STRING"
	Numeric                          AdvancedTaskDefinitionArgumentType = "NUMERIC"
	EnvironmentSourceName            AdvancedTaskDefinitionArgumentType = "ENVIRONMENT_SOURCE_NAME"
	EnvironmentSourceNameExcludeSelf AdvancedTaskDefinitionArgumentType = "ENVIRONMENT_SOURCE_NAME_EXCLUDE_SELF"
)

type AdvancedTaskDefinition struct {
	ID                   uint                             `yaml:"id,omitempty"`
	Name                 string                           `yaml:"name"`
	Description          string                           `yaml:"description,omitempty"`
	ConfirmationText     string                           `yaml:"confirmationText,omitempty"`
	Type                 string                           `yaml:"type"`
	Environment          int                              `yaml:"environment,omitempty"`
	Project              int                              `yaml:"project,omitempty"`
	Service              string                           `yaml:"service"`
	Command              string                           `yaml:"command,omitempty"`
	Arguments            []AdvancedTaskDefinitionArgument `yaml:"arguments,omitempty"`
	Permission           TaskPermission                   `yaml:"permission,omitempty"`
	SystemWide           bool                             `yaml:"systemWide,omitempty"`
	DeployTokenInjection bool                             `yaml:"deployTokenInjection,omitempty"`
	ProjectKeyInjection  bool                             `yaml:"projectKeyInjection,omitempty"`
	AdminOnlyView        bool                             `yaml:"adminOnlyView,omitempty"`
}

type AdvancedTaskDefinitionArgument struct {
	Name         string                             `json:"name" yaml:"name"`
	DisplayName  string                             `json:"displayName" yaml:"displayName"`
	Type         AdvancedTaskDefinitionArgumentType `json:"type" yaml:"type"`
	DefaultValue string                             `json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`
	Optional     bool                               `json:"optional,omitempty" yaml:"optional,omitempty"`
}
