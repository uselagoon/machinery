package client

import (
	"context"
	"encoding/json"
	"fmt"

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

// DeployTargetsByOrganizationId queries the Lagoon API for deploy targets by the given organization id
// and unmarshals the response into organization.
func (c *Client) DeployTargetsByOrganizationId(ctx context.Context, id uint, out *[]schema.DeployTarget) error {

	req, err := c.newRequest("_lgraphql/deploytargets/deployTargetsByOrganizationId.graphql",
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
	if len(o.DeployTargets) == 0 {
		return fmt.Errorf("no deploy targets found for organization %s", o.Name)
	}
	data, err := json.Marshal(o.DeployTargets)
	if err != nil {
		return err
	}
	json.Unmarshal(data, out)
	return nil
}

// DeployTargetsByOrganizationName queries the Lagoon API for deploy targets by the given organization name
// and unmarshals the response into organization.
func (c *Client) DeployTargetsByOrganizationName(ctx context.Context, name string, out *[]schema.DeployTarget) error {

	req, err := c.newRequest("_lgraphql/deploytargets/deployTargetsByOrganizationName.graphql",
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
	if len(o.DeployTargets) == 0 {
		return fmt.Errorf("no deploy targets found for organization %s", o.Name)
	}
	data, err := json.Marshal(o.DeployTargets)
	if err != nil {
		return err
	}
	json.Unmarshal(data, out)
	return nil
}

// AddDeployTargetToOrganization adds an existing deploytarget to an organization.
func (c *Client) AddDeployTargetToOrganization(ctx context.Context, in *schema.AddDeployTargetToOrganizationInput, out *schema.AddDeployTargetResponse) error {
	req, err := c.newRequest("_lgraphql/deploytargets/addDeployTargetToOrganization.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.AddDeployTargetResponse `json:"addDeployTargetToOrganization"`
	}{
		Response: out,
	})
}

// RemoveDeployTargetFromOrganization removes a deploytarget from an organization.
func (c *Client) RemoveDeployTargetFromOrganization(ctx context.Context, in *schema.RemoveDeployTargetFromOrganizationInput, out *schema.DeleteDeployTargetResponse) error {
	req, err := c.newRequest("_lgraphql/deploytargets/removeDeployTargetFromOrganization.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.DeleteDeployTargetResponse `json:"removeDeployTargetFromOrganization"`
	}{
		Response: out,
	})
}
