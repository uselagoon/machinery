package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/uselagoon/machinery/api/schema"
)

// Me queries the Lagoon API for me, and
// unmarshals the response into project.
func (c *Client) Me(
	ctx context.Context, user *schema.User) error {

	req, err := c.newRequest("_lgraphql/usergroups/me.graphql",
		nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.User `json:"me"`
	}{
		Response: user,
	})
}

// AddGroupsToProject adds Groups to a Project.
func (c *Client) AddGroupsToProject(ctx context.Context,
	in *schema.ProjectGroupsInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/usergroups/addGroupsToProject.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addGroupsToProject"`
	}{
		Response: out,
	})
}

// RemoveGroupsFromProject adds Groups to a Project.
func (c *Client) RemoveGroupsFromProject(ctx context.Context,
	in *schema.ProjectGroupsInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/usergroups/removeGroupsFromProject.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"removeGroupsFromProject"`
	}{
		Response: out,
	})
}

// AddGroup adds a group.
func (c *Client) AddGroup(
	ctx context.Context, in *schema.AddGroupInput, out *schema.Group) error {
	req, err := c.newRequest("_lgraphql/usergroups/addGroup.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Group `json:"addGroup"`
	}{
		Response: out,
	})
}

// AddUser adds a user.
func (c *Client) AddUser(
	ctx context.Context, in *schema.AddUserInput, out *schema.User) error {
	req, err := c.newRequest("_lgraphql/usergroups/addUser.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.User `json:"addUser"`
	}{
		Response: out,
	})
}

// AddSSHKey adds an SSH key to a user.
func (c *Client) AddSSHKey(ctx context.Context, in *schema.AddSSHKeyInput, out *schema.SSHKey) error {
	req, err := c.newRequest("_lgraphql/usergroups/addSshKey.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.SSHKey `json:"addSshKey"`
	}{
		Response: out,
	})
}

// AddUserToGroup adds a user to a group.
func (c *Client) AddUserToGroup(
	ctx context.Context, in *schema.UserGroupRoleInput, out *schema.Group) error {
	req, err := c.newRequest("_lgraphql/usergroups/addUserToGroup.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Group `json:"addUserToGroup"`
	}{
		Response: out,
	})
}

// RemoveUserFromGroup removes a user from a group.
func (c *Client) RemoveUserFromGroup(
	ctx context.Context, in *schema.UserGroupInput, out *schema.Group) error {
	req, err := c.newRequest("_lgraphql/usergroups/removeUserFromGroup.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Group `json:"removeUserFromGroup"`
	}{
		Response: out,
	})
}

// ListAllGroupMembers queries the Lagoon API for all groups and members, and
// unmarshals the response into project.
func (c *Client) ListAllGroupMembers(
	ctx context.Context, name string, groups *[]schema.Group) error {
	req, err := c.newRequest("_lgraphql/usergroups/allGroupMembers.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.Group `json:"allGroupMembers"`
	}{
		Response: groups,
	})
}

// ListGroupMembers queries the Lagoon API for groups and members, and
// unmarshals the response into project.
func (c *Client) ListGroupMembers(
	ctx context.Context, name string, groups *schema.Group) error {
	req, err := c.newRequest("_lgraphql/usergroups/getGroupMembers.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Group `json:"groupByName"`
	}{
		Response: groups,
	})
}

// AllUsers queries the Lagoon API for allUsers, and
// unmarshals the response into project.
func (c *Client) AllUsers(
	ctx context.Context, filter schema.AllUsersFilter, user *[]schema.User) error {
	fB, _ := json.Marshal(filter)
	var variables map[string]interface{}
	json.Unmarshal(fB, &variables)
	req, err := c.newRequest("_lgraphql/usergroups/allUsers.graphql", variables)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.User `json:"allUsers"`
	}{
		Response: user,
	})
}

func (c *Client) GetUserByEmail(
	ctx context.Context, email string, user *schema.User) error {

	req, err := c.newRequest("_lgraphql/usergroups/userByEmail.graphql",
		map[string]interface{}{
			"email": email,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.User `json:"userByEmail"`
	}{
		Response: user,
	})
}

// UserCanSSHToEnvironment queries the Lagoon API as a user to check if the user has access to the environment, and
// unmarshals the response.
func (c *Client) UserCanSSHToEnvironment(
	ctx context.Context, namespace string, environment *schema.Environment) error {

	req, err := c.newRequest("_lgraphql/usergroups/userCanSSHToEnvironment.graphql",
		map[string]string{"namespace": namespace})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &environment)
}

// UserBySSHKey queries the Lagoon API to find user by ssh key, and
// unmarshals the response.
func (c *Client) UserBySSHKey(
	ctx context.Context, sshKey string, user *schema.User) error {

	req, err := c.newRequest("_lgraphql/usergroups/userBySSHKey.graphql",
		map[string]string{"sshKey": sshKey})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &user)
}

// UserBySSHFingerprint queries the Lagoon API to find user by ssh key fingerprint, and
// unmarshals the response.
func (c *Client) UserBySSHFingerprint(
	ctx context.Context, fingerprint string, user *schema.User) error {

	req, err := c.newRequest("_lgraphql/usergroups/userBySSHFingerprint.graphql",
		map[string]string{"fingerprint": fingerprint})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &user)
}

// GetUserSSHKeysByEmail queries the Lagoon API to find user by email and return their associated ssh keys
func (c *Client) GetUserSSHKeysByEmail(ctx context.Context, email string, user *schema.User) error {
	req, err := c.newRequest("_lgraphql/usergroups/userByEmailSSHKeys.graphql",
		map[string]interface{}{
			"email": email,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.User `json:"userByEmail"`
	}{
		Response: user,
	})
}

// RemoveSSHKey removes an SSH key.
func (c *Client) RemoveSSHKey(ctx context.Context, id uint, out *schema.DeleteSSHKeyByIDInput) error {
	req, err := c.newRequest("_lgraphql/usergroups/removeSSHKeyById.graphql",
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// ListAllGroupMembersWithKeys queries the Lagoon API for all groups and members, and
// unmarshals the response into project.
func (c *Client) ListAllGroupMembersWithKeys(ctx context.Context, name string, groups *[]schema.Group) error {
	req, err := c.newRequest("_lgraphql/usergroups/allGroupMembersWithKeys.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.Group `json:"allGroups"`
	}{
		Response: groups,
	})
}

// GroupsByOrganizationID queries the Lagoon API for groups by the given organization id
// and unmarshals the response into organization.
func (c *Client) GroupsByOrganizationID(ctx context.Context, id uint, groups *[]schema.OrgGroup) error {

	req, err := c.newRequest("_lgraphql/usergroups/groupsByOrganizationId.graphql",
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return err
	}

	o := &schema.Organization{}
	err = c.client.Run(ctx, req, &struct {
		Response *schema.Organization `json:"organizationById"`
	}{
		Response: o,
	})
	if err != nil {
		return err
	}
	data, err := json.Marshal(o.Groups)
	if err != nil {
		return err
	}
	json.Unmarshal(data, groups)
	return nil
}

// AddGroupToOrganization adds a Group to an Organization.
func (c *Client) AddGroupToOrganization(ctx context.Context, in *schema.AddGroupToOrganizationInput, out *schema.OrgGroup) error {
	req, err := c.newRequest("_lgraphql/usergroups/addGroupToOrganization.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.OrgGroup `json:"addGroupToOrganization"`
	}{
		Response: out,
	})
}

// AddAdminToOrganization adds a User to an Organization.
func (c *Client) AddAdminToOrganization(ctx context.Context, in *schema.AddUserToOrganizationInput, out *schema.Organization) error {
	req, err := c.newRequest("_lgraphql/usergroups/addAdminToOrganization.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Organization `json:"addUserToOrganization"`
	}{
		Response: out,
	})
}

// RemoveAdminFromOrganization removes a User from an Organization. TODO - Create new input type for this
func (c *Client) RemoveAdminFromOrganization(ctx context.Context, in *schema.AddUserToOrganizationInput, out *schema.Organization) error {
	req, err := c.newRequest("_lgraphql/usergroups/removeAdminFromOrganization.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Organization `json:"removeUserFromOrganization"`
	}{
		Response: out,
	})
}

// UsersByOrganizationName queries the Lagoon API for users by the given organization name
func (c *Client) UsersByOrganizationName(ctx context.Context, name string, users *[]schema.OrgUser) error {
	org := &schema.Organization{}
	if err := c.OrganizationByName(ctx, name, org); err != nil {
		return err
	}
	if org.Name == "" {
		//lint:ignore ST1005 return a generic Lagoon API unauthorized error based on the permission called
		// this is because organizationbyname will return null instead of an error, the api should probably return an error
		return fmt.Errorf(`Unauthorized: You don't have permission to "viewUsers" on "organization"`)
	}

	return c.UsersByOrganizationID(ctx, org.ID, users)
}

// UsersByOrganizationID queries the Lagoon API for users by the given organization id
func (c *Client) UsersByOrganizationID(ctx context.Context, id uint, users *[]schema.OrgUser) error {
	req, err := c.newRequest("_lgraphql/usergroups/usersByOrganization.graphql",
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.OrgUser `json:"usersByOrganization"`
	}{
		Response: users,
	})
}

// ListOrganizationAdminsByName queries the Lagoon API for users by the given organization name
// and unmarshals the response into organization.
func (c *Client) ListOrganizationAdminsByName(ctx context.Context, name string, users *[]schema.OrgUser) error {
	req, err := c.newRequest("_lgraphql/organizations/listOrganizationAdminsByName.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	o := &schema.Organization{}
	err = c.client.Run(ctx, req, &struct {
		Response *schema.Organization `json:"organizationByName"`
	}{
		Response: o,
	})
	if err != nil {
		return err
	}
	data, err := json.Marshal(o.Owners)
	if err != nil {
		return err
	}
	json.Unmarshal(data, users)
	return nil
}

// AddProjectToGroup adds a group to a project.
func (c *Client) AddProjectToGroup(
	ctx context.Context, in *schema.ProjectGroupsInput, out *schema.Group) error {
	req, err := c.newRequest("_lgraphql/usergroups/addProjectToGroup.graphql", map[string]interface{}{
		"project": in.Project.Name,
		"groups":  in.Groups,
	})

	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Group `json:"addGroupsToProject"`
	}{
		Response: out,
	})
}

// DeleteGroup deletes a group.
func (c *Client) DeleteGroup(ctx context.Context, name string, out *schema.DeleteGroupInput) error {
	req, err := c.newRequest("_lgraphql/usergroups/deleteGroup.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// AllGroups queries the Lagoon API and returns all groups a user has access to.
func (c *Client) AllGroups(ctx context.Context, groups *[]schema.Group) error {
	req, err := c.newRequest("_lgraphql/usergroups/allGroups.graphql", nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.Group `json:"allGroups"`
	}{
		Response: groups,
	})
}

// GroupProjects queries the Lagoon API for a group by its name, returning all projects within the group.
func (c *Client) GroupProjects(
	ctx context.Context, name string, group *[]schema.Group) error {

	req, err := c.newRequest("_lgraphql/usergroups/groupProjects.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.Group `json:"allGroups"`
	}{
		Response: group,
	})
}

// DeleteUser removes a user from Lagoon.
func (c *Client) DeleteUser(ctx context.Context, in *schema.DeleteUserInput, out *schema.User) error {
	req, err := c.newRequest("_lgraphql/usergroups/deleteUser.graphql",
		map[string]interface{}{
			"email": in.User.Email,
		})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// UpdateUser updates a user in Lagoon.
func (c *Client) UpdateUser(ctx context.Context, in *schema.UpdateUserInput, out *schema.User) error {
	req, err := c.newRequest("_lgraphql/usergroups/updateUser.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.User `json:"updateUser"`
	}{
		Response: out,
	})
}

// ResetPassword resets a user's password.
func (c *Client) ResetPassword(ctx context.Context, in *schema.ResetUserPasswordInput, out *schema.User) error {
	req, err := c.newRequest("_lgraphql/usergroups/resetPassword.graphql",
		map[string]interface{}{
			"email": in.User.Email,
		})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}
