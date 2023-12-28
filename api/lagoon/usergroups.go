// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

// UserGroups interface contains methods for getting info on groups of lagoon.
type UserGroups interface {
	ListAllGroupMembers(ctx context.Context, name string, groups *[]schema.Group) error
	ListGroupMembers(ctx context.Context, name string, groups *schema.Group) error
	AddGroup(ctx context.Context, in *schema.AddGroupInput, group *schema.Group) error
	AddUser(ctx context.Context, in *schema.AddUserInput, user *schema.User) error
	AddUserToGroup(ctx context.Context, in *schema.UserGroupRoleInput, group *schema.Group) error
	RemoveUserFromGroup(ctx context.Context, in *schema.UserGroupInput, group *schema.Group) error
	AddGroupsToProject(ctx context.Context, in *schema.ProjectGroupsInput, project *schema.Project) error
	RemoveGroupsFromProject(ctx context.Context, in *schema.ProjectGroupsInput, project *schema.Project) error
	AddSSHKey(ctx context.Context, in *schema.AddSSHKeyInput, sshkey *schema.SSHKey) error
	Me(ctx context.Context, user *schema.User) error
	AllUsers(ctx context.Context, filter schema.AllUsersFilter, users *[]schema.User) error
	GetUserByEmail(ctx context.Context, email string, user *schema.User) error
	UserCanSSHToEnvironment(context.Context, string, *schema.Environment) error
	UserBySSHKey(ctx context.Context, sshKey string, user *schema.User) error
	UserBySSHFingerprint(ctx context.Context, fingerprint string, user *schema.User) error
	GetUserSSHKeysByEmail(ctx context.Context, email string, user *schema.User) error
	RemoveSSHKey(ctx context.Context, id uint, out *schema.DeleteSSHKeyByIDInput) error
	ListAllGroupMembersWithKeys(ctx context.Context, name string, groups *[]schema.Group) error
	GroupsByOrganizationID(ctx context.Context, id uint, group *[]schema.OrgGroup) error
	AddGroupToOrganization(ctx context.Context, in *schema.AddGroupToOrganizationInput, out *schema.OrgGroup) error
	UsersByOrganization(ctx context.Context, id uint, users *[]schema.OrgUser) error
	UsersByOrganizationName(ctx context.Context, name string, users *[]schema.OrgUser) error
}

// Me gets info on the current user of lagoon.
func Me(ctx context.Context, ug UserGroups) (*schema.User, error) {
	user := schema.User{}
	return &user, ug.Me(ctx, &user)
}

// ListAllGroupMembers gets info on the current groups of lagoon.
func ListAllGroupMembers(ctx context.Context, name string, ug UserGroups) (*[]schema.Group, error) {
	groups := []schema.Group{}
	return &groups, ug.ListAllGroupMembers(ctx, name, &groups)
}

// ListGroupMembers gets info on the current groups of lagoon.
func ListGroupMembers(ctx context.Context, name string, ug UserGroups) (*schema.Group, error) {
	groups := schema.Group{}
	return &groups, ug.ListGroupMembers(ctx, name, &groups)
}

func AddGroup(ctx context.Context, in *schema.AddGroupInput, ug UserGroups) (*schema.Group, error) {
	group := schema.Group{}
	return &group, ug.AddGroup(ctx, in, &group)
}

func AddUser(ctx context.Context, in *schema.AddUserInput, ug UserGroups) (*schema.User, error) {
	user := schema.User{}
	return &user, ug.AddUser(ctx, in, &user)
}

func AddSSHKey(ctx context.Context, in *schema.AddSSHKeyInput, ug UserGroups) (*schema.SSHKey, error) {
	sshkey := schema.SSHKey{}
	return &sshkey, ug.AddSSHKey(ctx, in, &sshkey)
}

func AddUserToGroup(ctx context.Context, in *schema.UserGroupRoleInput, ug UserGroups) (*schema.Group, error) {
	group := schema.Group{}
	return &group, ug.AddUserToGroup(ctx, in, &group)
}

func RemoveUserFromGroup(ctx context.Context, in *schema.UserGroupInput, ug UserGroups) (*schema.Group, error) {
	group := schema.Group{}
	return &group, ug.RemoveUserFromGroup(ctx, in, &group)
}

func AddGroupsToProject(ctx context.Context, in *schema.ProjectGroupsInput, ug UserGroups) (*schema.Project, error) {
	project := schema.Project{}
	return &project, ug.AddGroupsToProject(ctx, in, &project)
}

func RemoveGroupsFromProject(ctx context.Context, in *schema.ProjectGroupsInput, ug UserGroups) (*schema.Project, error) {
	project := schema.Project{}
	return &project, ug.RemoveGroupsFromProject(ctx, in, &project)
}

// AllUsers gets info on the current users of lagoon.
func AllUsers(ctx context.Context, filter schema.AllUsersFilter, ug UserGroups) (*[]schema.User, error) {
	users := []schema.User{}
	return &users, ug.AllUsers(ctx, filter, &users)
}

func GetUserByEmail(ctx context.Context, email string, ug UserGroups) (*schema.User, error) {
	user := schema.User{}
	return &user, ug.GetUserByEmail(ctx, email, &user)
}

func UserCanSSHToEnvironment(ctx context.Context, namespace string, ug UserGroups) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, ug.UserCanSSHToEnvironment(ctx, namespace, &environment)
}

func UserBySSHKey(ctx context.Context, sshKey string, ug UserGroups) (*schema.User, error) {
	user := schema.User{}
	return &user, ug.UserBySSHKey(ctx, sshKey, &user)
}

func UserBySSHFingerprint(ctx context.Context, fingerprint string, ug UserGroups) (*schema.User, error) {
	user := schema.User{}
	return &user, ug.UserBySSHFingerprint(ctx, fingerprint, &user)
}

func GetUserSSHKeysByEmail(ctx context.Context, email string, ug UserGroups) (*schema.User, error) {
	user := schema.User{}
	return &user, ug.GetUserSSHKeysByEmail(ctx, email, &user)
}

func RemoveSSHKey(ctx context.Context, id uint, ug UserGroups) (*schema.DeleteSSHKeyByIDInput, error) {
	result := schema.DeleteSSHKeyByIDInput{}
	return &result, ug.RemoveSSHKey(ctx, id, &result)
}

// ListAllGroupMembersWithKeys gets info on the current groups of lagoon.
func ListAllGroupMembersWithKeys(ctx context.Context, name string, ug UserGroups) (*[]schema.Group, error) {
	groups := []schema.Group{}
	return &groups, ug.ListAllGroupMembersWithKeys(ctx, name, &groups)
}

// ListGroupsByOrganizationID gets groups associated with an organization in lagoon via provided ID.
func ListGroupsByOrganizationID(ctx context.Context, id uint, ug UserGroups) (*[]schema.OrgGroup, error) {
	group := []schema.OrgGroup{}
	return &group, ug.GroupsByOrganizationID(ctx, id, &group)
}

// AddGroupToOrganization adds a group to an organization.
func AddGroupToOrganization(ctx context.Context, in *schema.AddGroupToOrganizationInput, ug UserGroups) (*schema.OrgGroup, error) {
	group := schema.OrgGroup{}
	return &group, ug.AddGroupToOrganization(ctx, in, &group)
}

// UsersByOrganization lists users associated within an organization in lagoon via provided ID.
func UsersByOrganization(ctx context.Context, id uint, ug UserGroups) (*[]schema.OrgUser, error) {
	user := []schema.OrgUser{}
	return &user, ug.UsersByOrganization(ctx, id, &user)
}

// UsersByOrganizationName lists users associated within an organization in lagoon via provided Name.
func UsersByOrganizationName(ctx context.Context, name string, ug UserGroups) (*[]schema.OrgUser, error) {
	user := []schema.OrgUser{}
	return &user, ug.UsersByOrganizationName(ctx, name, &user)
}
