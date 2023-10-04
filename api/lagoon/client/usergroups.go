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
func (c *Client) AddSSHKey(
	ctx context.Context, in *schema.AddSSHKeyInput, out *schema.SSHKey) error {
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
	ctx context.Context, groups *[]schema.Group) error {

	req, err := c.newRequest("_lgraphql/usergroups/allGroupMembers.graphql", nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.Group `json:"allGroupMembers"`
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

// GroupsByOrganizationID queries the Lagoon API for groups by the given organization id
// and unmarshals the response into environment.
func (c *Client) GroupsByOrganizationID(ctx context.Context, id uint, groups *[]schema.Group) error {

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
	if len(o.Groups) == 0 {
		return fmt.Errorf("no associated groups found for organization %s", o.Name)
	}
	data, err := json.Marshal(o.Groups)
	if err != nil {
		return err
	}
	json.Unmarshal(data, groups)
	return nil
}
