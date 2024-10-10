// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

// Projects interface contains methods for getting info on projects.
type Projects interface {
	MinimalProjectByName(ctx context.Context, name string, project *schema.Project) error
	ProjectByNameMetadata(ctx context.Context, name string, project *schema.ProjectMetadata) error
	ProjectsByMetadata(ctx context.Context, key string, value string, project *[]schema.ProjectMetadata) error
	UpdateProjectMetadata(ctx context.Context, id uint, key string, value string, project *schema.ProjectMetadata) error
	RemoveProjectMetadataByKey(ctx context.Context, id uint, key string, project *schema.ProjectMetadata) error
	NotificationsForProjectByName(ctx context.Context, name string, result *schema.Project) error
	UpdateProject(ctx context.Context, id uint, patch schema.UpdateProjectPatchInput, project *schema.Project) error
	SSHEndpointsByProject(ctx context.Context, name string, project *schema.Project) error
	ProjectGroups(ctx context.Context, name string, project *schema.Project) error
	ProjectByNameExtended(ctx context.Context, name string, project *schema.Project) error
	ProjectsByOrganizationID(ctx context.Context, id uint, project *[]schema.OrgProject) error
	ProjectsByOrganizationName(ctx context.Context, name string, project *[]schema.OrgProject) error
	RemoveProjectFromOrganization(ctx context.Context, in *schema.RemoveProjectFromOrganizationInput, out *schema.Project) error
	ProjectKeyByName(ctx context.Context, name string, revealKey bool, project *schema.Project) error
	AllProjects(ctx context.Context, project *[]schema.Project) error
	DeleteProject(ctx context.Context, project string, result *schema.DeleteProject) error

	UpdateProjectByName(ctx context.Context, name string, patch schema.UpdateProjectPatchInput, project *schema.Project) error
	UpdateProjectMetadataByName(ctx context.Context, name string, key string, value string, project *schema.ProjectMetadata) error
	RemoveProjectMetadataByKeyByName(ctx context.Context, name string, key string, project *schema.ProjectMetadata) error
}

// GetMinimalProjectByName gets basic information about the project including the openshift id.
func GetMinimalProjectByName(ctx context.Context, name string, p Projects) (*schema.Project, error) {
	project := schema.Project{}
	return &project, p.MinimalProjectByName(ctx, name, &project)
}

// GetProjectsByMetadata gets info of projects in lagoon that have matching metadata.
func GetProjectsByMetadata(ctx context.Context, key string, value string, p Projects) (*[]schema.ProjectMetadata, error) {
	project := []schema.ProjectMetadata{}
	return &project, p.ProjectsByMetadata(ctx, key, value, &project)
}

// GetProjectMetadata gets the metadata key:values for a lagoon project.
func GetProjectMetadata(ctx context.Context, name string, p Projects) (*schema.ProjectMetadata, error) {
	project := schema.ProjectMetadata{}
	return &project, p.ProjectByNameMetadata(ctx, name, &project)
}

// UpdateProjectMetadata updates a project with provided metadata.
func UpdateProjectMetadata(ctx context.Context, id uint, key string, value string, p Projects) (*schema.ProjectMetadata, error) {
	project := schema.ProjectMetadata{}
	return &project, p.UpdateProjectMetadata(ctx, id, key, value, &project)
}

// UpdateProjectMetadataByName updates a project with provided metadata.
func UpdateProjectMetadataByName(ctx context.Context, name string, key string, value string, p Projects) (*schema.ProjectMetadata, error) {
	project := schema.ProjectMetadata{}
	return &project, p.UpdateProjectMetadataByName(ctx, name, key, value, &project)
}

// RemoveProjectMetadataByKey remove metadata from a project by key.
func RemoveProjectMetadataByKey(ctx context.Context, id uint, key string, p Projects) (*schema.ProjectMetadata, error) {
	project := schema.ProjectMetadata{}
	return &project, p.RemoveProjectMetadataByKey(ctx, id, key, &project)
}

// RemoveProjectMetadataByKey remove metadata from a project by key.
func RemoveProjectMetadataByKeyByName(ctx context.Context, name string, key string, p Projects) (*schema.ProjectMetadata, error) {
	project := schema.ProjectMetadata{}
	return &project, p.RemoveProjectMetadataByKeyByName(ctx, name, key, &project)
}

// NotificationsForProject gets notifications for a project.
func NotificationsForProject(ctx context.Context, name string, p Projects) (*schema.Project, error) {
	result := schema.Project{}
	return &result, p.NotificationsForProjectByName(ctx, name, &result)
}

// UpdateProject updates a project with provided patch data.
func UpdateProject(ctx context.Context, id uint, patch schema.UpdateProjectPatchInput, p Projects) (*schema.Project, error) {
	project := schema.Project{}
	return &project, p.UpdateProject(ctx, id, patch, &project)
}

// UpdateProjectByName updates a project with provided patch data.
func UpdateProjectByName(ctx context.Context, name string, patch schema.UpdateProjectPatchInput, p Projects) (*schema.Project, error) {
	project := schema.Project{}
	return &project, p.UpdateProjectByName(ctx, name, patch, &project)
}

// GetSSHEndpointsByProject gets info of projects in lagoon that have matching metadata.
func GetSSHEndpointsByProject(ctx context.Context, name string, p Projects) (*schema.Project, error) {
	project := schema.Project{}
	return &project, p.SSHEndpointsByProject(ctx, name, &project)
}

// GetProjectGroups gets groups in a specified project.
func GetProjectGroups(ctx context.Context, name string, p Projects) (*schema.Project, error) {
	project := schema.Project{}
	return &project, p.ProjectGroups(ctx, name, &project)
}

// GetProjectByName gets info of projects in lagoon that have matching metadata.
func GetProjectByName(ctx context.Context, name string, p Projects) (*schema.Project, error) {
	project := schema.Project{}
	return &project, p.ProjectByNameExtended(ctx, name, &project)
}

// ListProjectsByOrganizationID gets projects associated with an organization in lagoon via provided ID.
func ListProjectsByOrganizationID(ctx context.Context, id uint, p Projects) (*[]schema.OrgProject, error) {
	project := []schema.OrgProject{}
	return &project, p.ProjectsByOrganizationID(ctx, id, &project)
}

// ListProjectsByOrganizationName gets projects associated with an organization in lagoon via provided name.
func ListProjectsByOrganizationName(ctx context.Context, name string, p Projects) (*[]schema.OrgProject, error) {
	project := []schema.OrgProject{}
	return &project, p.ProjectsByOrganizationName(ctx, name, &project)
}

// RemoveProjectFromOrganization removes a project from an organization.
func RemoveProjectFromOrganization(ctx context.Context, in *schema.RemoveProjectFromOrganizationInput, out Projects) (*schema.Project, error) {
	response := schema.Project{}
	return &response, out.RemoveProjectFromOrganization(ctx, in, &response)
}

// GetProjectKeyByName gets the project keys of a project in Lagoon.
func GetProjectKeyByName(ctx context.Context, name string, revealKey bool, p Projects) (*schema.Project, error) {
	project := schema.Project{}
	return &project, p.ProjectKeyByName(ctx, name, revealKey, &project)
}

// ListAllProjects lists all projects.
func ListAllProjects(ctx context.Context, p Projects) (*[]schema.Project, error) {
	project := []schema.Project{}
	return &project, p.AllProjects(ctx, &project)
}

// DeleteProject updates environment storage.
func DeleteProject(ctx context.Context, project string, p Projects) (*schema.DeleteProject, error) {
	result := schema.DeleteProject{}
	return &result, p.DeleteProject(ctx, project, &result)
}
