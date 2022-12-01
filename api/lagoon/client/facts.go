package client

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

func (c *Client) AddFacts(ctx context.Context, in *schema.AddFactInput, out *[]schema.Facts) error {

	req, err := c.newRequest("_lgraphql/facts/addFactsByName.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *[]schema.Facts `json:"addFactsByName"`
	}{
		Response: out,
	})
}
