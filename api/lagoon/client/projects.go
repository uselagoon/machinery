package client

import (
	"context"
	"encoding/json"
	"github.com/uselagoon/machinery/api/schema"
)

// ProjectByName queries the Lagoon API for a project by its name, and
// unmarshals the response into project.
func (c *Client) ProjectByName(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newRequest("_lgraphql/projects/projectByName.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"projectByName"`
	}{
		Response: project,
	})
}

// ProjectByNameExtended queries the Lagoon API for a project by its name, and
// unmarshals the response into project. Built for use with the Lagoon CLI - get project command
func (c *Client) ProjectByNameExtended(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newRequest("_lgraphql/projects/projectByNameExtended.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"projectByName"`
	}{
		Response: project,
	})
}

// MinimalProjectByName queries the Lagoon API for a project by its name, and
// unmarshals the response into project.
func (c *Client) MinimalProjectByName(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newRequest("_lgraphql/projects/minimalProjectByName.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"projectByName"`
	}{
		Response: project,
	})
}

// ProjectByNameMetadata queries the Lagoon API for a project by its name, and
// unmarshals the response into project.
func (c *Client) ProjectByNameMetadata(
	ctx context.Context, name string, project *schema.ProjectMetadata) error {

	req, err := c.newRequest("_lgraphql/projects/projectByNameMetadata.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.ProjectMetadata `json:"projectByName"`
	}{
		Response: project,
	})
}

// ProjectsByMetadata queries the Lagoon API for a project by its name, and
// unmarshals the response into project.
func (c *Client) ProjectsByMetadata(
	ctx context.Context, key string, value string, projects *[]schema.ProjectMetadata) error {

	req, err := c.newRequest("_lgraphql/projects/projectsByMetadata.graphql",
		map[string]interface{}{
			"key":   key,
			"value": value,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.ProjectMetadata `json:"projectsByMetadata"`
	}{
		Response: projects,
	})
}

// AddProject adds a project.
func (c *Client) AddProject(
	ctx context.Context, in *schema.AddProjectInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/projects/addProject.graphql", in)
	if err != nil {
		return err
	}
	return wrapErr(c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addProject"`
	}{
		Response: out,
	}))
}

// UpdateProjectMetadata updates a projects metadata.
func (c *Client) UpdateProjectMetadata(
	ctx context.Context, id int, key string, value string, projects *schema.ProjectMetadata) error {

	req, err := c.newRequest("_lgraphql/projects/updateProjectMetadata.graphql",
		map[string]interface{}{
			"id":    id,
			"key":   key,
			"value": value,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.ProjectMetadata `json:"updateProjectMetadata"`
	}{
		Response: projects,
	})
}

// RemoveProjectMetadataByKey removes metadata from a project for given key.
func (c *Client) RemoveProjectMetadataByKey(
	ctx context.Context, id int, key string, projects *schema.ProjectMetadata) error {

	req, err := c.newRequest("_lgraphql/projects/removeProjectMetadataByKey.graphql",
		map[string]interface{}{
			"id":  id,
			"key": key,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.ProjectMetadata `json:"removeProjectMetadataByKey"`
	}{
		Response: projects,
	})
}

// UpdateProject updates a project.
func (c *Client) UpdateProject(
	ctx context.Context, id int, patch schema.UpdateProjectPatchInput, projects *schema.Project) error {

	req, err := c.newRequest("_lgraphql/projects/updateProject.graphql",
		map[string]interface{}{
			"id":    id,
			"patch": patch,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"updateProject"`
	}{
		Response: projects,
	})
}

// ProjectGroups queries the Lagoon API for a project by its name, returning all groups within the project.
func (c *Client) ProjectGroups(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newRequest("_lgraphql/projects/projectGroups.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"projectByName"`
	}{
		Response: project,
	})
}

// ProjectsByOrganizationID queries the Lagoon API for projects by the given organization id
// and unmarshals the response into organization.
func (c *Client) ProjectsByOrganizationID(ctx context.Context, id uint, projects *[]schema.OrgProject) error {

	req, err := c.newRequest("_lgraphql/projects/projectsByOrganizationId.graphql",
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
	data, err := json.Marshal(o.Projects)
	if err != nil {
		return err
	}
	json.Unmarshal(data, projects)
	return nil
}

// RemoveProjectFromOrganization removes a project from an organization.
func (c *Client) RemoveProjectFromOrganization(ctx context.Context, in *schema.RemoveProjectFromOrganizationInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/projects/removeProjectFromOrganization.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"removeProjectFromOrganization"`
	}{
		Response: out,
	})
}
