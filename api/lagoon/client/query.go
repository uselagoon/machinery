package client

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

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
