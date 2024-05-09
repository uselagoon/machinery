package client

import (
	"context"
	"github.com/uselagoon/machinery/api/schema"
)

// GetEnvVariablesByProjectEnvironmentName queries the Lagoon API for a envvars by project environment and unmarshals the response.
func (c *Client) GetEnvVariablesByProjectEnvironmentName(
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

// AddOrUpdateEnvVariableByName queries the Lagoon API to add or update an envvar by project environment and unmarshals the response.
func (c *Client) AddOrUpdateEnvVariableByName(ctx context.Context, in *schema.EnvVariableByNameInput, envvar *schema.UpdateEnvVarResponse) error {
	req, err := c.newRequest("_lgraphql/variables/addOrUpdateEnvVariableByName.graphql",
		map[string]interface{}{
			"input": in,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.UpdateEnvVarResponse `json:"addOrUpdateEnvVariableByName"`
	}{
		Response: envvar,
	})
}

// DeleteEnvVariableByName queries the Lagoon API to delete an envvar by project environment and unmarshals the response.
func (c *Client) DeleteEnvVariableByName(ctx context.Context, in *schema.DeleteEnvVariableByNameInput, out *schema.DeleteEnvVarResponse) error {
	req, err := c.newRequest("_lgraphql/variables/deleteEnvVariableByName.graphql",
		map[string]interface{}{
			"input": in,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &out)
}
