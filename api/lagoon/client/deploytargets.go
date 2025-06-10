package client

import (
	"context"
	"encoding/json"

	"github.com/machinebox/graphql"
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

// DeployTargetsByOrganizationNameOrID queries the Lagoon API for deploy targets by the given organization name
// and unmarshals the response into organization.
// Need to rework the argument types here
func (c *Client) DeployTargetsByOrganizationNameOrID(ctx context.Context, name *string, id *uint, out *[]schema.DeployTarget) error {
	var req *graphql.Request
	var err error
	o := &schema.Organization{}

	if name != nil {
		req, err = c.newRequest("_lgraphql/deploytargets/deployTargetsByOrganizationName.graphql",
			map[string]interface{}{
				"name": name,
			})
		if err != nil {
			return err
		}
		err = c.client.Run(ctx, req, &struct {
			Response *schema.Organization `json:"organizationByName"`
		}{
			Response: o,
		})
		if err != nil {
			return err
		}
	} else if id != nil {
		req, err = c.newRequest("_lgraphql/deploytargets/deployTargetsByOrganizationId.graphql",
			map[string]interface{}{
				"id": id,
			})
		if err != nil {
			return err
		}
		err = c.client.Run(ctx, req, &struct {
			Response *schema.Organization `json:"organizationByID"`
		}{
			Response: o,
		})
		if err != nil {
			return err
		}
	}

	data, err := json.Marshal(o.DeployTargets)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, out)
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
