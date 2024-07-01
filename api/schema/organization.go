package schema

import "github.com/google/uuid"

// Organization is based on the Lagoon API type.
type Organization struct {
	AddOrganizationInput
	DeployTargets []DeployTarget   `json:"deployTargets,omitempty"`
	Projects      []OrgProject     `json:"projects,omitempty"`
	Environments  []OrgEnvironment `json:"environments,omitempty"`
	Groups        []OrgGroup       `json:"groups,omitempty"`
	Owners        []OrgUser        `json:"owners,omitempty"`
	Notifications *Notifications   `json:"notifications,omitempty"`
}

type AddOrganizationInput struct {
	ID                uint   `json:"id,omitempty"`
	Name              string `json:"name"`
	FriendlyName      string `json:"friendlyName,omitempty"`
	Description       string `json:"description,omitempty"`
	QuotaProject      int    `json:"quotaProject,omitempty"`
	QuotaGroup        int    `json:"quotaGroup,omitempty"`
	QuotaNotification int    `json:"quotaNotification,omitempty"`
	QuotaEnvironment  int    `json:"quotaEnvironment,omitempty"`
	QuotaRoute        int    `json:"quotaRoute,omitempty"`
}

type OrgProject struct {
	ID            uint                       `json:"id,omitempty"`
	Name          string                     `json:"name,omitempty"`
	Organization  uint                       `json:"organization,omitempty"`
	GroupCount    uint                       `json:"groupCount,omitempty"`
	Groups        *OrgGroups                 `json:"groups,omitempty"`
	Notifications []OrganizationNotification `json:"notifications,omitempty"`
}

type OrgEnvironment struct {
	ID                         uint                 `json:"id,omitempty"`
	Name                       string               `json:"name,omitempty"`
	Project                    OrgProject           `json:"project,omitempty"`
	DeployType                 DeployType           `json:"deployType,omitempty"`
	DeployHeadRef              string               `json:"deployHeadRef,omitempty"`
	DeployTitle                string               `json:"deployTitle,omitempty"`
	EnvironmentType            EnvType              `json:"environmentType,omitempty"`
	OpenshiftProjectName       string               `json:"openshiftProjectName,omitempty"`
	AutoIdle                   uint                 `json:"autoIdle,omitempty"`
	KubernetesNamespacePattern string               `json:"kubernetesNamespacePattern,omitempty"`
	Updated                    string               `json:"updated,omitempty"`
	Created                    string               `json:"created,omitempty"`
	Deleted                    string               `json:"deleted,omitempty"`
	Route                      string               `json:"route,omitempty"`
	Routes                     string               `json:"routes,omitempty"`
	Services                   []EnvironmentService `json:"services,omitempty"`
	Openshift                  DeployTarget         `json:"openshift,omitempty"`
	Kubernetes                 DeployTarget         `json:"kubernetes,omitempty"`
}

// AddGroupToOrganizationInput is based on the input to AddGroupToOrganization.
type AddGroupToOrganizationInput struct {
	Name         string      `json:"name"`
	Organization uint        `json:"organization"`
	ParentGroup  *GroupInput `json:"parentGroup,omitempty"`
	AddOrgOwner  bool        `json:"addOrgOwner,omitempty"`
}

// OrgGroup provides for unmarshalling the groups contained with an Organization.
type OrgGroup struct {
	AddGroupToOrganizationInput
	ID      *uuid.UUID `json:"id,omitempty"`
	Type    string     `json:"type,omitempty"`
	Members []struct {
		User User      `json:"user"`
		Role GroupRole `json:"role"`
	} `json:"members,omitempty"`
	MemberCount int          `json:"memberCount,omitempty"`
	Projects    []OrgProject `json:"projects,omitempty"`
}

type OrgGroups struct {
	OrgGroups []OrgGroup
}

type OrganizationNotification struct {
	Name string           `json:"name"`
	Type NotificationType `json:"type"`
}

type OrgUser struct {
	ID         *uuid.UUID   `json:"id,omitempty"`
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
	QuotaProject      *int    `json:"quotaProject,omitempty"`
	QuotaGroup        *int    `json:"quotaGroup,omitempty"`
	QuotaNotification *int    `json:"quotaNotification,omitempty"`
	QuotaEnvironment  *int    `json:"quotaEnvironment,omitempty"`
	QuotaRoute        *int    `json:"quotaRoute,omitempty"`
}

type UpdateOrganizationInput struct {
	ID    uint                         `json:"id"`
	Patch UpdateOrganizationPatchInput `json:"patch"`
}
