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
