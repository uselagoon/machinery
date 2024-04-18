package client

import (
	"context"
	"github.com/uselagoon/machinery/api/schema"
)

// EnvVariablesByProjectEnvironmentName queries the Lagoon API for a envvars by project environment and unmarshals the response.
func (c *Client) EnvVariablesByProjectEnvironmentName(
	ctx context.Context, in *schema.EnvVariableByProjectEnvironmentNameInput, envkeyvalue *[]schema.EnvKeyValue) error {

	req, err := c.newRequest("_lgraphql/variables/getEnvVariablesByProjectEnvironmentName.graphql",
		map[string]interface{}{
			"input": in,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.EnvKeyValue `json:"getEnvVariablesByProjectEnvironmentName"`
	}{
		Response: envkeyvalue,
	})
}
