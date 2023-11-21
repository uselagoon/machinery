package schema

import (
	"github.com/google/uuid"
)

// GroupRole .
type GroupRole string

// Guest .
const (
	GuestRole      GroupRole = "GUEST"
	ReporterRole   GroupRole = "REPORTER"
	DeveloperRole  GroupRole = "DEVELOPER"
	MaintainerRole GroupRole = "MAINTAINER"
	OwnerRole      GroupRole = "OWNER"
)

// AddGroupInput is based on the input to addGroup.
type AddGroupInput struct {
	Name        string      `json:"name"`
	ParentGroup *GroupInput `json:"parentGroup,omitempty"`
}

// Group provides for unmarshalling the groups contained with a Project.
type Group struct {
	AddGroupInput
	ID           *uuid.UUID `json:"id,omitempty"`
	Organization int        `json:"organization,omitempty"`
	Members      []struct {
		User User      `json:"user"`
		Role GroupRole `json:"role"`
	} `json:"members,omitempty"`
}

// GroupConfig embeds AddGroupInput as well as a list of members.
type GroupConfig struct {
	AddGroupInput
	Users []UserRoleConfig `json:"users,omitempty"`
}

// GroupInput is based on the Lagoon API type.
type GroupInput ProjectInput

// UserGroupRoleInput is based on the Lagoon API type.
type UserGroupRoleInput struct {
	UserEmail string    `json:"userEmail"`
	GroupName string    `json:"groupName"`
	GroupRole GroupRole `json:"groupRole"`
}

// UserGroupInput is based on the Lagoon API type.
type UserGroupInput struct {
	UserEmail string `json:"userEmail"`
	GroupName string `json:"groupName"`
}

// UserRoleConfig stores a user/role config within a group.
type UserRoleConfig struct {
	Email string    `json:"email"`
	Role  GroupRole `json:"role"`
}

// Groups represents possible Lagoon group types.
// These are unmarshalled from a projectByName query response.
type Groups struct {
	Groups []Group
}
