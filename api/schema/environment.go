package schema

// EnvType .
type EnvType string

// . .
const (
	ProductionEnv  EnvType = "PRODUCTION"
	DevelopmentEnv EnvType = "DEVELOPMENT"
)

// NotificationType .
type NotificationType string

// . .
const (
	SlackNotification          NotificationType = "SLACK"
	RocketChatNotification     NotificationType = "ROCKETCHAT"
	EmailNotification          NotificationType = "EMAIL"
	MicrosoftTeamsNotification NotificationType = "MICROSOFTTEAMS"
)

// AddEnvironmentInput is based on the input to
// addOrUpdateEnvironment.
type AddEnvironmentInput struct {
	ID                   uint       `json:"id,omitempty"`
	Name                 string     `json:"name"`
	ProjectID            uint       `json:"project"`
	DeployType           DeployType `json:"deployType"`
	DeployBaseRef        string     `json:"deployBaseRef"`
	DeployHeadRef        string     `json:"deployHeadRef,omitempty"`
	DeployTitle          string     `json:"deployTitle,omitempty"`
	EnvironmentType      EnvType    `json:"environmentType"`
	OpenshiftProjectName string     `json:"openshiftProjectName"`
}

// Environment is the Lagoon API Environment object.
type Environment struct {
	AddEnvironmentInput
	AutoIdle     uint                 `json:"autoIdle"`
	EnvVariables []EnvKeyValue        `json:"envVariables,omitempty"`
	Route        string               `json:"route,omitempty"`
	Routes       string               `json:"routes,omitempty"`
	Backups      []Backup             `json:"backups,omitempty"`
	Deployments  []Deployment         `json:"deployments,omitempty"`
	Services     []EnvironmentService `json:"services,omitempty"`
	// TODO use a unixtime type
	Updated string `json:"updated,omitempty"`
	Created string `json:"created,omitempty"`
	Deleted string `json:"deleted,omitempty"`
}

// EnvironmentConfig contains Environment configuration.
type EnvironmentConfig struct {
	Environment
	// override embedded AddEnvironmentInput.ProjectID to omitempty
	ProjectID uint `json:"project,omitempty"`
}

// EnvironmentInput  is based on the Lagoon API type.
type EnvironmentInput struct {
	ID      int          `json:"id,omitempty"`
	Name    string       `json:"name,omitempty"`
	Project ProjectInput `json:"project"`
}

// UpdateEnvironmentStorageInput is used as the input for updating an environments storage.
type UpdateEnvironmentStorageInput struct {
	Environment          int    `json:"environment"`
	PersisteStorageClaim string `json:"persistentStorageClaim"`
	BytesUsed            int    `json:"bytesUsed"`
}

// UpdateEnvironmentStorage is the response.
type UpdateEnvironmentStorage struct {
	ID int `json:"id"`
}

// DeleteEnvironment is the response.
type DeleteEnvironment struct {
	DeleteEnvironment string `json:"deleteEnvironment"`
}

type UpdateEnvironmentPatchInput struct {
	ID                   uint       `json:"id,omitempty"`
	Name                 string     `json:"name,omitempty"`
	ProjectID            uint       `json:"project,omitempty"`
	DeployType           DeployType `json:"deployType,omitempty"`
	DeployBaseRef        string     `json:"deployBaseRef,omitempty"`
	DeployHeadRef        string     `json:"deployHeadRef,omitempty"`
	DeployTitle          string     `json:"deployTitle,omitempty"`
	EnvironmentType      EnvType    `json:"environmentType,omitempty"`
	OpenshiftProjectName string     `json:"openshiftProjectName,omitempty"`
	Route                string     `json:"route,omitempty"`
	Routes               string     `json:"routes,omitempty"`
	MonitoringURLs       string     `json:"monitoringUrls,omitempty"`
	AutoIdle             uint       `json:"autoIdle"`
	Openshift            uint       `json:"openshift,omitempty"`
	Created              string     `json:"created,omitempty"`
}

// EnvironmentService  is based on the Lagoon API type.
type EnvironmentService struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
