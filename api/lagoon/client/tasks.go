package client

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

// UpdateTask updates a task.
func (c *Client) UpdateTask(
	ctx context.Context, id int, patch schema.UpdateTaskPatchInput, task *schema.Task) error {

	req, err := c.newRequest("_lgraphql/tasks/updateTask.graphql",
		map[string]interface{}{
			"id":    id,
			"patch": patch,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Task `json:"updateTask"`
	}{
		Response: task,
	})
}

// RunActiveStandbySwitch deploys a branch.
func (c *Client) RunActiveStandbySwitch(ctx context.Context,
	project string, out *schema.Task) error {
	req, err := c.newRequest("_lgraphql/tasks/switchActiveStandby.graphql", map[string]interface{}{
		"project": project,
	})
	if err != nil {
		return err
	}

	// return c.client.Run(ctx, req, &out)
	return c.client.Run(ctx, req, &struct {
		Response *schema.Task `json:"switchActiveStandby"`
	}{
		Response: out,
	})
}
