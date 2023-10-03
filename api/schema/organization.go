package schema

// Organization is based on the Lagoon API type.
type Organization struct {
	AddOrganizationInput
	DeployTarget  []DeployTarget   `json:"openshift,omitempty"`
	Projects      []OrgProject     `json:"projects,omitempty"`
	Environments  []OrgEnvironment `json:"environments,omitempty"`
	Groups        []*Groups        `json:"groups,omitempty"`
	Owners        []OrgUser        `json:"owners,omitempty"`
	Notifications *Notifications   `json:"notifications,omitempty"`
}

type AddOrganizationInput struct {
	ID                uint   `json:"id,omitempty"`
	Name              string `json:"name"`
	FriendlyName      string `json:"friendlyName"`
	Description       string `json:"description,omitempty"`
	QuotaProject      uint   `json:"quotaProject"`
	QuotaGroup        uint   `json:"quotaGroup,omitempty"`
	QuotaNotification uint   `json:"quotaNotification,omitempty"`
	QuotaEnvironment  uint   `json:"quotaEnvironment,omitempty"`
	QuotaRoute        uint   `json:"quotaRoute"`
}

type OrgProject struct {
	ID            uint                       `json:"id,omitempty"`
	Name          string                     `json:"nameOrgProject,omitempty"`
	Organization  uint                       `json:"organization,omitempty"`
	Groups        *Groups                    `json:"groups,omitempty"`
	Notifications []OrganizationNotification `json:"notifications,omitempty"`
}

type OrgEnvironment struct {
	ID                         uint                 `json:"id,omitempty"`
	Name                       string               `json:"name,omitempty"`
	ProjectID                  uint                 `json:"project,omitempty"`
	DeployType                 DeployType           `json:"deployType,omitempty"`
	DeployHeadRef              string               `json:"deployHeadRef,omitempty"`
	DeployTitle                string               `json:"deployTitle,omitempty"`
	EnvironmentType            EnvType              `json:"environmentType,omitempty"`
	OpenshiftProjectName       string               `json:"openshiftProjectName,omitempty"`
	AutoIdle                   uint                 `json:"autoIdle,omitempty"`
	KubernetesNamespacePattern *string              `json:"kubernetesNamespacePattern,omitempty"`
	Updated                    string               `json:"updated,omitempty"`
	Created                    string               `json:"created,omitempty"`
	Deleted                    string               `json:"deleted,omitempty"`
	Route                      string               `json:"route,omitempty"`
	Routes                     string               `json:"routes,omitempty"`
	Services                   []EnvironmentService `json:"services,omitempty"`
	DeployTarget               DeployTarget         `json:"openshift,omitempty"`
	Kubernetes                 *uint                `json:"kubernetes,omitempty"`
}

type OrganizationNotification struct {
	Name string           `json:"name"`
	Type NotificationType `json:"type"`
}

type OrgUser struct {
	ID         string       `json:"id,omitempty"`
	Email      string       `json:"email,omitempty"`
	FirstName  string       `json:"firstName,omitempty"`
	LastName   string       `json:"lastName,omitempty"`
	Comment    string       `json:"comment,omitempty"`
	Owner      bool         `json:"owner,omitempty"`
	GroupRoles []GroupRoles `json:"groupRoles,omitempty"`
}

type DeleteOrganizationInput struct {
	ID uint `json:"id"`
}

type UpdateOrganizationPatchInput struct {
	Name              *string `json:"name,omitempty"`
	FriendlyName      *string `json:"friendlyName,omitempty"`
	Description       *string `json:"description,omitempty"`
	QuotaProject      *uint   `json:"quotaProject,omitempty"`
	QuotaGroup        *uint   `json:"quotaGroup,omitempty"`
	QuotaNotification *uint   `json:"quotaNotification,omitempty"`
	QuotaEnvironment  *uint   `json:"quotaEnvironment,omitempty"`
	QuotaRoute        *uint   `json:"quotaRoute,omitempty"`
}

type UpdateOrganizationInput struct {
	ID    uint                         `json:"id"`
	Patch UpdateOrganizationPatchInput `json:"patch"`
}
