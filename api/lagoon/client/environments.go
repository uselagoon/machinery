package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/uselagoon/machinery/api/schema"
)

// UpdateStorageOnEnvironment updates an environments storage values.
func (c *Client) UpdateStorageOnEnvironment(ctx context.Context,
	in *schema.UpdateStorageOnEnvironmentInput, out *schema.UpdateStorageOnEnvironment) error {
	req, err := c.newRequest("_lgraphql/environments/addOrUpdateStorageOnEnvironment.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.UpdateStorageOnEnvironment `json:"addOrUpdateStorageOnEnvironment"`
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

func (c *Client) EnvironmentByNameAndProjectName(ctx context.Context, name string,
	projectName string, environment *schema.Environment) error {
	project := &schema.Project{}
	if err := c.veryMinimalProjectByName(ctx, projectName, project); err != nil {
		return err
	}
	if project.Name == "" {
		//lint:ignore ST1005 return a generic Lagoon API unauthorized error based on the permission called
		// this is because organizationbyname will return null instead of an error, the api should probably return an error
		return fmt.Errorf(`Unauthorized: You don't have permission to "view" on "project"`) //nolint:staticcheck
	}
	return c.EnvironmentByName(ctx, name, project.ID, environment)
}

// EnvironmentByID queries the Lagoon API for an environment by its ID and unmarshals the response into environment.
func (c *Client) EnvironmentByID(ctx context.Context, environmentID uint, environment *schema.Environment) error {

	req, err := c.newRequest("_lgraphql/environments/environmentById.graphql",
		map[string]interface{}{
			"id": environmentID,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"environmentById"`
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

// BackupsByEnvironmentNamespace queries the Lagoon API for an environment by its namespace
// and unmarshals the response into environment.
func (c *Client) BackupsByEnvironmentNamespace(ctx context.Context, namespace string, environment *schema.Environment) error {

	req, err := c.newRequest("_lgraphql/environments/backupsByEnvironmentNamespace.graphql",
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

// EnvironmentsByProjectName queries the Lagoon API for environments by the given project name
// and unmarshals the response into environment.
func (c *Client) EnvironmentsByProjectName(ctx context.Context, project string, environments *[]schema.Environment) error {

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
	db, err := json.Marshal(p.Environments)
	if err != nil {
		return err
	}
	return json.Unmarshal(db, environments)
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

func (c *Client) BackupsForEnvironmentByNameAndProjectName(ctx context.Context, name string,
	projectName string, environment *schema.Environment) error {
	project := &schema.Project{}
	if err := c.veryMinimalProjectByName(ctx, projectName, project); err != nil {
		return err
	}
	if project.Name == "" {
		//lint:ignore ST1005 return a generic Lagoon API unauthorized error based on the permission called
		// this is because organizationbyname will return null instead of an error, the api should probably return an error
		return fmt.Errorf(`Unauthorized: You don't have permission to "view" on "project"`) //nolint:staticcheck
	}
	return c.BackupsForEnvironmentByName(ctx, name, project.ID, environment)
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

// AddBackup adds a backup.
func (c *Client) AddBackup(
	ctx context.Context, environment uint, source, backupID, created string, out *schema.Backup) error {
	req, err := c.newRequest("_lgraphql/environments/addBackup.graphql",
		map[string]interface{}{
			"environment": environment,
			"source":      source,
			"backupId":    backupID,
			"created":     created,
		})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Backup `json:"addBackup"`
	}{
		Response: out,
	})
}

// UpdateRestore updates a restore.
func (c *Client) UpdateRestore(
	ctx context.Context, input schema.UpdateRestoreInput, out *schema.Restore) error {
	req, err := c.newRequest("_lgraphql/environments/updateRestore.graphql",
		map[string]interface{}{
			"backupId": input.BackupID,
			"patch":    input.Patch,
		})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Restore `json:"updateRestore"`
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

// SSHEndpointByNamespace queries the Lagoon API for an environment by its namespace
// and unmarshals the response into environment.
func (c *Client) SSHEndpointByNamespace(ctx context.Context, namespace string, environment *schema.Environment) error {
	req, err := c.newRequest("_lgraphql/environments/sshEndpointByEnvironment.graphql",
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

// AddOrUpdateEnvironmentService updates an environment service.
func (c *Client) AddOrUpdateEnvironmentService(
	ctx context.Context, service schema.AddEnvironmentServiceInput, result *schema.EnvironmentService) error {

	req, err := c.newRequest("_lgraphql/environments/addOrUpdateEnvironmentService.graphql",
		map[string]interface{}{
			"service": service,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.EnvironmentService `json:"addOrUpdateEnvironmentService"`
	}{
		Response: result,
	})
}

// DeleteEnvironmentService deletes an environment service.
func (c *Client) DeleteEnvironmentService(
	ctx context.Context, service schema.DeleteEnvironmentServiceInput, result *schema.DeleteEnvironmentService) error {

	req, err := c.newRequest("_lgraphql/environments/deleteEnvironmentService.graphql",
		map[string]interface{}{
			"environment": service.EnvironmentID,
			"name":        service.Name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.DeleteEnvironmentService `json:"deleteEnvironmentService"`
	}{
		Response: result,
	})
}

// DeleteBackup deletes an environment backup.
func (c *Client) DeleteBackup(ctx context.Context,
	backupID string, result *schema.DeleteBackup) error {
	req, err := c.newRequest("_lgraphql/environments/deleteBackup.graphql",
		map[string]interface{}{
			"backupId": backupID,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &result)
}
