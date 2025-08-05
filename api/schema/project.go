package schema

import "github.com/guregu/null"

// ProjectAvailability determines the number of pods used to run a project.
type ProjectAvailability string

// High tells Lagoon to load balance across multiple pods.
// Standard tells Lagoon to use a single pod for the site.
const (
	High     ProjectAvailability = "HIGH"
	Standard ProjectAvailability = "STANDARD"
)

// AddProjectInput is based on the Lagoon API type.
type AddProjectInput struct {
	ID                           uint                `json:"id,omitempty"`
	Name                         string              `json:"name"`
	GitURL                       string              `json:"gitUrl"`
	Subfolder                    string              `json:"subfolder,omitempty"`
	Openshift                    uint                `json:"openshift"`
	OpenshiftProjectPattern      string              `json:"openshiftProjectPattern,omitempty"`
	Branches                     string              `json:"branches,omitempty"`
	PullRequests                 string              `json:"pullrequests,omitempty"`
	ProductionEnvironment        string              `json:"productionEnvironment"`
	StandbyProductionEnvironment string              `json:"standbyProductionEnvironment,omitempty"`
	Availability                 ProjectAvailability `json:"availability,omitempty"`
	// AutoIdle and StorageCalc can't be omitempty because their zero-values are
	// significant - Lagoon uses it as a boolean (0/1).
	AutoIdle                     *uint         `json:"autoIdle,omitempty"`
	StorageCalc                  *uint         `json:"storageCalc,omitempty"`
	DevelopmentEnvironmentsLimit *uint         `json:"developmentEnvironmentsLimit,omitempty"`
	PrivateKey                   string        `json:"privateKey,omitempty"`
	PublicKey                    string        `json:"publicKey,omitempty"`
	BuildImage                   string        `json:"buildImage,omitempty"`
	Organization                 *uint         `json:"organization,omitempty"`
	OrganizationDetails          *Organization `json:"organizationDetails,omitempty"`
	AddOrgOwner                  *bool         `json:"addOrgOwner,omitempty"`
	RouterPattern                string        `json:"routerPattern,omitempty"`
	ProblemsUI                   *uint         `json:"problemsUi,omitempty"`
	FactsUI                      *uint         `json:"factsUi,omitempty"`
	ProductionBuildPriority      *uint         `json:"productionBuildPriority,omitempty"`
	DevelopmentBuildPriority     *uint         `json:"developmentBuildPriority,omitempty"`
	DeploymentsDisabled          *uint         `json:"deploymentsDisabled,omitempty"`
}

// Project is the Lagoon API Project object.
type Project struct {
	AddProjectInput
	Environments []EnvironmentConfig `json:"environments,omitempty"`
	EnvVariables []EnvKeyValue       `json:"envVariables,omitempty"`
	// Notifications is unmarshalled during export.
	Notifications *Notifications `json:"notifications,omitempty"`
	// Openshift is unmarshalled during export.
	OpenshiftID *OpenshiftID `json:"openshift,omitempty"`
	// Groups are unmarshalled during export.
	Groups              []Group              `json:"groups,omitempty"`
	DeployTargetConfigs []DeployTargetConfig `json:"deployTargetConfigs,omitempty"`
}

// ProjectConfig contains project configuration.
type ProjectConfig struct {
	Project
	// ProjectNotifications are (un)marshalled during import.
	Notifications *ProjectNotifications `json:"notifications,omitempty"`
	// Group are (un)marshalled during import.
	Groups []string `json:"groups,omitempty"`
	// Users are members of the project.
	// Note that in Lagoon this is implemented as being a member of the
	// project-<projectname> group.
	Users []UserRoleConfig `json:"users,omitempty"`
}

// ProjectNotifications lists the notifications for a project within a
// ProjectConfig.
type ProjectNotifications struct {
	Slack          []string `json:"slack,omitempty"`
	RocketChat     []string `json:"rocketChat,omitempty"`
	Email          []string `json:"email,omitempty"`
	MicrosoftTeams []string `json:"microsoftTeams,omitempty"`
}

// OpenshiftID is unmarshalled during export.
type OpenshiftID struct {
	ID uint `json:"id,omitempty"`
}

// ProjectGroupsInput is based on the input to
// addGroupsToProject.
type ProjectGroupsInput struct {
	Project ProjectInput `json:"project"`
	Groups  []GroupInput `json:"groups"`
}

// ProjectInput is based on the Lagoon API type.
type ProjectInput struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ProjectMetadata .
type ProjectMetadata struct {
	Project
	Metadata map[string]string `json:"metadata"`
}

type UpdateProjectPatchInput struct {
	Name                         *string              `json:"name,omitempty"`
	GitURL                       *string              `json:"gitUrl,omitempty"`
	Availability                 *ProjectAvailability `json:"availability,omitempty"`
	PrivateKey                   *string              `json:"privateKey,omitempty"`
	Subfolder                    *string              `json:"subfolder,omitempty"`
	RouterPattern                *string              `json:"routerPattern,omitempty"`
	Branches                     *string              `json:"branches,omitempty"`
	ProductionEnvironment        *string              `json:"productionEnvironment,omitempty"`
	ProductionRoutes             *string              `json:"productionRoutes,omitempty"`
	ProductionAlias              *string              `json:"productionAlias,omitempty"`
	StandbyProductionEnvironment *string              `json:"standbyProductionEnvironment,omitempty"`
	StandbyRoutes                *string              `json:"standbyRoutes,omitempty"`
	StandbyAlias                 *string              `json:"standbyAlias,omitempty"`
	AutoIdle                     *uint                `json:"autoIdle,omitempty"`
	StorageCalc                  *uint                `json:"storageCalc,omitempty"`
	Pullrequests                 *string              `json:"pullrequests,omitempty"`
	Kubernetes                   *uint                `json:"kubernetes,omitempty"`
	KubernetesNamespacePattern   *string              `json:"kubernetesNamespacePattern,omitempty"`
	DevelopmentEnvironmentsLimit *uint                `json:"developmentEnvironmentsLimit,omitempty"`
	ProblemsUI                   *uint                `json:"problemsUi,omitempty"`
	FactsUI                      *uint                `json:"factsUi,omitempty"`
	ProductionBuildPriority      *uint                `json:"productionBuildPriority,omitempty"`
	DevelopmentBuildPriority     *uint                `json:"developmentBuildPriority,omitempty"`
	DeploymentsDisabled          *uint                `json:"deploymentsDisabled,omitempty"`
	// `null` is valid graphql, use a pointer to allow `nil` to be empty
	BuildImage              *null.String `json:"buildImage,omitempty"`
	Openshift               *uint        `json:"openshift,omitempty"`
	OpenshiftProjectPattern *string      `json:"openshiftProjectPattern,omitempty"`
}

// RemoveProjectFromOrganizationInput is based on the Lagoon API type.
type RemoveProjectFromOrganizationInput struct {
	Project      uint `json:"project"`
	Organization uint `json:"organization"`
}

type DeleteProject struct {
	Project string `json:"name"`
}
