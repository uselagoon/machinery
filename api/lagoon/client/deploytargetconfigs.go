package client

import (
	"context"
	"fmt"

	"github.com/uselagoon/machinery/api/schema"
)

// DeployTargetConfigsByProjectID queries the Lagoon API for a projects deploytarget configs by its id, and
// unmarshals the response into deploytargetconfigs.
func (c *Client) DeployTargetConfigsByProjectID(
	ctx context.Context, project uint, deploytargetconfigs *[]schema.DeployTargetConfig) error {
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

// DeployTargetConfigsByProjectName is the same as DeployTargetConfigsByProjectID but does a project lookup before the deployments call
func (c *Client) DeployTargetConfigsByProjectName(
	ctx context.Context, projectName string, deploytargetconfigs *[]schema.DeployTargetConfig) error {
	project := &schema.Project{}
	if err := c.veryMinimalProjectByName(ctx, projectName, project); err != nil {
		return err
	}
	if project.Name == "" {
		//lint:ignore ST1005 return a generic Lagoon API unauthorized error based on the permission called
		// this is because organizationbyname will return null instead of an error, the api should probably return an error
		return fmt.Errorf(`Unauthorized: You don't have permission to "view" on "project"`)
	}
	return c.DeployTargetConfigsByProjectID(ctx, project.ID, deploytargetconfigs)
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

// AddDeployTargetConfiguration adds a deploytarget configuration to a project.
func (c *Client) AddDeployTargetConfigurationByProjectName(ctx context.Context,
	projectName string,
	in *schema.AddDeployTargetConfigInput, out *schema.DeployTargetConfig) error {
	project := &schema.Project{}
	if err := c.veryMinimalProjectByName(ctx, projectName, project); err != nil {
		return err
	}
	if project.Name == "" {
		//lint:ignore ST1005 return a generic Lagoon API unauthorized error based on the permission called
		// this is because organizationbyname will return null instead of an error, the api should probably return an error
		return fmt.Errorf(`Unauthorized: You don't have permission to "view" on "project"`)
	}
	in.Project = project.ID
	return c.AddDeployTargetConfiguration(ctx, in, out)
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
	id, project uint, out *schema.DeleteDeployTargetConfig) error {
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

// DeleteDeployTargetConfigurationByIDAndProjectName is the same as DeleteDeployTargetConfig but does a project lookup before the deployments call
func (c *Client) DeleteDeployTargetConfigurationByIDAndProjectName(ctx context.Context,
	id uint, projectName string, out *schema.DeleteDeployTargetConfig) error {
	project := &schema.Project{}
	if err := c.veryMinimalProjectByName(ctx, projectName, project); err != nil {
		return err
	}
	if project.Name == "" {
		//lint:ignore ST1005 return a generic Lagoon API unauthorized error based on the permission called
		// this is because organizationbyname will return null instead of an error, the api should probably return an error
		return fmt.Errorf(`Unauthorized: You don't have permission to "view" on "project"`)
	}
	return c.DeleteDeployTargetConfiguration(ctx, id, project.ID, out)
}
