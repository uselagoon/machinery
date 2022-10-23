package client

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

// DeployTargetConfigsByProjectID queries the Lagoon API for a projects deploytarget configs by its id, and
// unmarshals the response into deploytargetconfigs.
func (c *Client) DeployTargetConfigsByProjectID(
	ctx context.Context, project int, deploytargetconfigs *[]schema.DeployTargetConfig) error {
	req, err := c.newRequest("_lgraphql/deploytargetconfigs/deployTargetConfigsByProjectId.graphql",
		map[string]interface{}{
			"project": project,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.DeployTargetConfig `json:"deployTargetConfigsByProjectId"`
	}{
		Response: deploytargetconfigs,
	})
}

// AddDeployTargetConfiguration adds a deploytarget configuration to a project.
func (c *Client) AddDeployTargetConfiguration(ctx context.Context,
	in *schema.AddDeployTargetConfigInput, out *schema.DeployTargetConfig) error {
	req, err := c.newRequest("_lgraphql/deploytargetconfigs/addDeployTargetConfig.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.DeployTargetConfig `json:"addDeployTargetConfig"`
	}{
		Response: out,
	})
}

// UpdateDeployTargetConfiguration adds a deploytarget configuration to a project.
func (c *Client) UpdateDeployTargetConfiguration(ctx context.Context,
	in *schema.UpdateDeployTargetConfigInput, out *schema.DeployTargetConfig) error {
	req, err := c.newRequest("_lgraphql/deploytargetconfigs/updateDeployTargetConfig.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.DeployTargetConfig `json:"updateDeployTargetConfig"`
	}{
		Response: out,
	})
}

// DeleteDeployTargetConfig deletes a deploytarget config from a project.
func (c *Client) DeleteDeployTargetConfiguration(ctx context.Context,
	id int, project int, out *schema.DeleteDeployTargetConfig) error {
	req, err := c.newRequest("_lgraphql/deploytargetconfigs/deleteDeployTargetConfig.graphql",
		map[string]interface{}{
			"id":      id,
			"project": project,
		})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}
