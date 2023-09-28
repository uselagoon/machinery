package client

import (
	"context"
	"github.com/uselagoon/machinery/api/schema"
)

// OrganizationByID queries the Lagoon API for an organization by its ID, and
// unmarshals the response into organization.
func (c *Client) OrganizationByID(ctx context.Context, id int, organization *schema.Organization) error {
	req, err := c.newRequest("_lgraphql/organizations/organizationById.graphql",
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Organization `json:"organizationById"`
	}{
		Response: organization,
	})
}

// AddOrganization adds an organization.
func (c *Client) AddOrganization(
	ctx context.Context, in *schema.AddOrganizationInput, org *schema.Organization) error {
	req, err := c.newRequest("_lgraphql/organizations/addOrganization.graphql", in)
	if err != nil {
		return err
	}

	return wrapErr(c.client.Run(ctx, req, &struct {
		Response *schema.Organization `json:"addOrganization"`
	}{
		Response: org,
	}))
}

// DeleteOrganization deletes an organization.
func (c *Client) DeleteOrg(ctx context.Context,
	id int, out *schema.DeleteOrganizationInput) error {
	req, err := c.newRequest("_lgraphql/organizations/deleteOrganization.graphql",
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}
