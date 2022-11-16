package client

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

// Me queries the Lagoon API for me, and
// unmarshals the response into project.
func (c *Client) Me(
	ctx context.Context, user *schema.User) error {

	req, err := c.newRequest("_lgraphql/me.graphql",
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

// LagoonAPIVersion queries the Lagoon API for its version, and
// unmarshals the response.
func (c *Client) LagoonAPIVersion(
	ctx context.Context, lagoonAPIVersion *schema.LagoonVersion) error {

	req, err := c.newRequest("_lgraphql/lagoonVersion.graphql",
		nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &lagoonAPIVersion)
}

// LagoonSchema queries the Lagoon API for its schema, and
// unmarshals the response.
func (c *Client) LagoonSchema(
	ctx context.Context, lagoonSchema *schema.LagoonSchema) error {

	req, err := c.newRequest("_lgraphql/lagoonSchema.graphql",
		nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.LagoonSchema `json:"__schema"`
	}{
		Response: lagoonSchema,
	})
}

// GetTaskByID queries the Lagoon API for a task by its ID, and
// unmarshals the response.
func (c *Client) GetTaskByID(
	ctx context.Context, id int, task *schema.Task) error {

	req, err := c.newRequest("_lgraphql/taskByID.graphql",
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Task `json:"taskById"`
	}{
		Response: task,
	})
}

// NotificationsForProjectByName gets all notifications for a project
func (c *Client) NotificationsForProjectByName(
	ctx context.Context, name string, project *schema.Project) error {
	req, err := c.newRequest("_lgraphql/projectNotifications.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"projectByName"`
	}{
		Response: project,
	})
}

// SSHEndpointsByProject queries the Lagoon API for a project by its name, and
// unmarshals the response into project.
func (c *Client) SSHEndpointsByProject(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newRequest("_lgraphql/sshEndpointsByProject.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"projectByName"`
	}{
		Response: project,
	})
}
