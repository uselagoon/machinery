package client

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

// ListDeployTargets queries the Lagoon API for a deploytargets and unmarshals the response into deploytargets.
func (c *Client) ListDeployTargets(
	ctx context.Context, deploytargets *[]schema.DeployTarget) error {
	req, err := c.newRequest("_lgraphql/deploytargets/listDeployTargets.graphql", nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.DeployTarget `json:"listDeployTargets"`
	}{
		Response: deploytargets,
	})
}

// AddDeployTarget adds a deploytarget (kubernetes/openshift).
func (c *Client) AddDeployTarget(ctx context.Context, in *schema.AddDeployTargetInput, out *schema.AddDeployTargetResponse) error {
	req, err := c.newRequest("_lgraphql/deploytargets/addDeployTarget.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.AddDeployTargetResponse `json:"addDeployTarget"`
	}{
		Response: out,
	})
}

// UpdateDeployTarget updates a deploytarget (kubernetes/openshift).
func (c *Client) UpdateDeployTarget(ctx context.Context, in *schema.UpdateDeployTargetInput, out *schema.UpdateDeployTargetResponse) error {
	req, err := c.newRequest("_lgraphql/deploytargets/updateDeployTarget.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.UpdateDeployTargetResponse `json:"updateDeployTarget"`
	}{
		Response: out,
	})
}

// DeleteDeployTarget deletes a deploytarget (kubernetes/openshift).
func (c *Client) DeleteDeployTarget(ctx context.Context, in *schema.DeleteDeployTargetInput, out *schema.DeleteDeployTargetResponse) error {
	req, err := c.newRequest("_lgraphql/deploytargets/deleteDeployTarget.graphql", in)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &out)
}
