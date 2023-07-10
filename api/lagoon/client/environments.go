package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/uselagoon/machinery/api/schema"
)

// UpdateEnvironmentStorage updates an environments storage values.
func (c *Client) UpdateEnvironmentStorage(ctx context.Context,
	in *schema.UpdateEnvironmentStorageInput, out *schema.UpdateEnvironmentStorage) error {
	req, err := c.newRequest("_lgraphql/environments/addOrUpdateEnvironmentStorage.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.UpdateEnvironmentStorage `json:"addOrUpdateEnvironmentStorage"`
	}{
		Response: out,
	})
}

// DeleteEnvironment deletes an environment.
func (c *Client) DeleteEnvironment(ctx context.Context,
	name, project string, execute bool, out *schema.DeleteEnvironment) error {
	req, err := c.newRequest("_lgraphql/environments/deleteEnvironment.graphql",
		map[string]interface{}{
			"name":    name,
			"project": project,
			"execute": execute,
		})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// EnvironmentByName queries the Lagoon API for an environment by its name and
// parent projectID, and unmarshals the response into environment.
func (c *Client) EnvironmentByName(ctx context.Context, name string,
	projectID uint, environment *schema.Environment) error {

	req, err := c.newRequest("_lgraphql/environments/environmentByName.graphql",
		map[string]interface{}{
			"name":    name,
			"project": projectID,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"environmentByName"`
	}{
		Response: environment,
	})
}

// EnvironmentByName queries the Lagoon API for an environment by its namespace
// and unmarshals the response into environment.
func (c *Client) EnvironmentByNamespace(ctx context.Context, namespace string, environment *schema.Environment) error {

	req, err := c.newRequest("_lgraphql/environments/environmentByNamespace.graphql",
		map[string]interface{}{
			"namespace": namespace,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"environmentByNamespace"`
	}{
		Response: environment,
	})
}

// EnvironmentsByProject queries the Lagoon API for environments by the project name
// and unmarshals the response into environment.
func (c *Client) EnvironmentsByProject(ctx context.Context, project string, environments *[]schema.Environment) error {

	req, err := c.newRequest("_lgraphql/environments/environmentsByProjectName.graphql",
		map[string]interface{}{
			"project": project,
		})
	if err != nil {
		return err
	}

	p := &schema.Project{}
	err = c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"projectByName"`
	}{
		Response: p,
	})
	if err != nil {
		return err
	}
	if len(p.Environments) == 0 {
		return fmt.Errorf("no environments found for project")
	}
	db, err := json.Marshal(p.Environments)
	if err != nil {
		return err
	}
	json.Unmarshal(db, environments)
	return nil
}

// BackupsForEnvironmentByName queries the Lagoon API for an environment by its name and
// parent projectID, and unmarshals the response into environment.
func (c *Client) BackupsForEnvironmentByName(ctx context.Context, name string,
	projectID uint, environment *schema.Environment) error {

	req, err := c.newRequest("_lgraphql/environments/backupsForEnvironmentByName.graphql",
		map[string]interface{}{
			"name":    name,
			"project": projectID,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"environmentByName"`
	}{
		Response: environment,
	})
}

// AddRestore adds a restore.
func (c *Client) AddRestore(
	ctx context.Context, backupID string, out *schema.Restore) error {
	req, err := c.newRequest("_lgraphql/environments/addRestore.graphql",
		map[string]interface{}{
			"backupid": backupID,
		})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Restore `json:"addRestore"`
	}{
		Response: out,
	})
}

// AddOrUpdateEnvironment adds or updates a Project Environment.
func (c *Client) AddOrUpdateEnvironment(ctx context.Context,
	in *schema.AddEnvironmentInput, out *schema.Environment) error {
	req, err := c.newRequest("_lgraphql/environments/addOrUpdateEnvironment.graphql", in)
	if err != nil {
		return err
	}
	return wrapErr(c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"addOrUpdateEnvironment"`
	}{
		Response: out,
	}))
}

// UpdateEnvironment updates an environment.
func (c *Client) UpdateEnvironment(
	ctx context.Context, id uint, patch schema.UpdateEnvironmentPatchInput, environment *schema.Environment) error {

	req, err := c.newRequest("_lgraphql/environments/updateEnvironment.graphql",
		map[string]interface{}{
			"id":    id,
			"patch": patch,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"updateEnvironment"`
	}{
		Response: environment,
	})
}

// SetEnvironmentServices updates an environment.
func (c *Client) SetEnvironmentServices(
	ctx context.Context, id uint, services []string, result *[]schema.EnvironmentService) error {

	req, err := c.newRequest("_lgraphql/environments/setEnvironmentServices.graphql",
		map[string]interface{}{
			"id":       id,
			"services": services,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.EnvironmentService `json:"setEnvironmentServices"`
	}{
		Response: result,
	})
}
