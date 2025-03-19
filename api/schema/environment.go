package schema

import (
	"fmt"
	"strings"
)

// EnvType .
type EnvType string

func (e EnvType) validateType() error {
	envTypes := map[EnvType]struct{}{
		ProductionEnv:  {},
		DevelopmentEnv: {},
	}

	envT := strings.ToUpper(string(e))
	envType := EnvType(envT)
	_, ok := envTypes[envType]
	if !ok {
		return fmt.Errorf(`cannot parse:[%s] as EnvType`, e)
	}
	return nil
}

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
	WebhookNotification        NotificationType = "WEBHOOK"
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
	AutoIdle      *uint                `json:"autoIdle,omitempty"`
	EnvVariables  []EnvKeyValue        `json:"envVariables,omitempty"`
	Route         string               `json:"route,omitempty"`
	Routes        string               `json:"routes,omitempty"`
	Backups       []Backup             `json:"backups,omitempty"`
	Deployments   []Deployment         `json:"deployments,omitempty"`
	Services      []EnvironmentService `json:"services,omitempty"`
	DeployTarget  DeployTarget         `json:"openshift,omitempty"`
	Tasks         []Task               `json:"tasks,omitempty"`
	AdvancedTasks []AdvancedTask       `json:"advancedTasks,omitempty"`
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
// @DEPRECATED
type UpdateEnvironmentStorageInput struct {
	Environment          int    `json:"environment"`
	PersisteStorageClaim string `json:"persistentStorageClaim"`
	BytesUsed            int    `json:"bytesUsed"`
}

// UpdateEnvironmentStorage is the response.
// @DEPRECATED
type UpdateEnvironmentStorage struct {
	ID int `json:"id"`
}

// UpdateEnvironmentStorageInput is used as the input for updating an environments storage.
type UpdateStorageOnEnvironmentInput struct {
	Environment          int    `json:"environment"`
	PersisteStorageClaim string `json:"persistentStorageClaim"`
	KiBUsed              uint64 `json:"kibUsed"`
}

// UpdateEnvironmentStorage is the response.
type UpdateStorageOnEnvironment struct {
	ID int `json:"id"`
}

// DeleteEnvironment is the response.
type DeleteEnvironment struct {
	DeleteEnvironment string `json:"deleteEnvironment"`
}

type UpdateEnvironmentPatchInput struct {
	ID                   *uint       `json:"id,omitempty"`
	Name                 *string     `json:"name,omitempty"`
	ProjectID            *uint       `json:"project,omitempty"`
	DeployType           *DeployType `json:"deployType,omitempty"`
	DeployBaseRef        *string     `json:"deployBaseRef,omitempty"`
	DeployHeadRef        *string     `json:"deployHeadRef,omitempty"`
	DeployTitle          *string     `json:"deployTitle,omitempty"`
	EnvironmentType      *EnvType    `json:"environmentType,omitempty"`
	OpenshiftProjectName *string     `json:"openshiftProjectName,omitempty"`
	Route                *string     `json:"route,omitempty"`
	Routes               *string     `json:"routes,omitempty"`
	AutoIdle             *uint       `json:"autoIdle,omitempty"`
	Openshift            *uint       `json:"openshift,omitempty"`
	Created              *string     `json:"created,omitempty"`
}

// EnvironmentService  is based on the Lagoon API type.
type EnvironmentService struct {
	ID         int                `json:"id,omitempty"`
	Name       string             `json:"name,omitempty"`
	Type       string             `json:"type,omitempty"`
	Updated    string             `json:"updated,omitempty"`
	Containers []ServiceContainer `json:"containers,omitempty"`
	Created    string             `json:"created,omitempty"`
}

// ServiceContainer  is based on the Lagoon API type.
type ServiceContainer struct {
	Name    string                   `json:"name,omitempty"`
	Ports   []ServiceContainerPort   `json:"ports,omitempty"`
	Volumes []ServiceContainerVolume `json:"volumes,omitempty"`
}

type ServiceContainerVolume struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

type ServiceContainerPort struct {
	Port     int32  `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

// AddEnvironmentServiceInput is based on the input to
// addOrUpdateEnvironmentService.
type AddEnvironmentServiceInput struct {
	ID            uint                    `json:"id,omitempty"`
	Name          string                  `json:"name"`
	Type          string                  `json:"type"`
	Containers    []ServiceContainerInput `json:"containers,omitempty"`
	EnvironmentID uint                    `json:"environment"`
}

// ServiceContainerInput  is based on the Lagoon API type.
type ServiceContainerInput struct {
	Name string `json:"name"`
}

// DeleteEnvironmentServiceInput is based on the input to
// deleteEnvironmentService.
type DeleteEnvironmentServiceInput struct {
	Name          string `json:"name"`
	EnvironmentID uint   `json:"environment"`
}

// DeleteEnvironmentService is the response.
type DeleteEnvironmentService struct {
	DeleteEnvironmentService string `json:"deleteEnvironmentService"`
}
