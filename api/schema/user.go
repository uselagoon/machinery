package schema

import "github.com/google/uuid"

// AddUserInput is based on the Lagoon API type.
type AddUserInput struct {
	Email         string `json:"email"`
	FirstName     string `json:"firstName,omitempty"`
	LastName      string `json:"lastName,omitempty"`
	Comment       string `json:"comment,omitempty"`
	GitlabID      uint   `json:"gitlabId,omitempty"`
	ResetPassword bool   `json:"resetPassword,omitempty"`
}

// User provides for unmarshalling the users contained withing a Group.
type User struct {
	AddUserInput
	ID         *uuid.UUID   `json:"id,omitempty"`
	SSHKeys    []SSHKey     `json:"sshKeys,omitempty"`
	GroupRoles []GroupRoles `json:"groupRoles,omitempty"`
}

type UserInput struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

type GroupRoles struct {
	Name string `json:"name,omitempty"`
	Role string `json:"role,omitempty"`
}

type AllUsersFilter struct {
	ID       string `json:"id,omitempty"`
	GitlabID uint   `json:"gitlabId,omitempty"`
	Email    string `json:"email,omitempty"`
}

type AddUserToOrganizationInput struct {
	User         UserInput `json:"user"`
	Organization uint      `json:"organization"`
	Owner        bool      `json:"owner,omitempty"`
	Admin        bool      `json:"admin,omitempty"`
}

type DeleteUserInput struct {
	User UserInput `json:"user"`
}

type UpdateUserInput struct {
	User  UserInput            `json:"user"`
	Patch UpdateUserPatchInput `json:"patch"`
}

type UpdateUserPatchInput struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Comment   *string `json:"comment,omitempty"`
	GitlabID  *uint   `json:"gitlabId,omitempty"`
}

type ResetUserPasswordInput struct {
	User UserInput `json:"user"`
}
