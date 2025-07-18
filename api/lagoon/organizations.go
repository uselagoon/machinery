// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

// Organizations interface contains methods for getting info on projects.
type Organizations interface {
	OrganizationByID(ctx context.Context, id uint, organization *schema.Organization) error
	OrganizationByName(ctx context.Context, name string, organization *schema.Organization) error
	DeleteOrg(ctx context.Context, id uint, organization *schema.DeleteOrganizationInput) error
	UpdateOrganization(ctx context.Context, id uint, patch schema.UpdateOrganizationPatchInput, result *schema.Organization) error
	AllOrganizations(ctx context.Context, organizations *[]schema.Organization) error
	AllOrganizationsExtended(ctx context.Context, organizations *[]schema.Organization) error
}

// GetOrganizationByID gets info of an organization in lagoon that matches the provided ID.
func GetOrganizationByID(ctx context.Context, id uint, o Organizations) (*schema.Organization, error) {
	organization := schema.Organization{}
	return &organization, o.OrganizationByID(ctx, id, &organization)
}

// GetOrganizationByName gets info of an organization in lagoon that matches the provided name.
func GetOrganizationByName(ctx context.Context, name string, o Organizations) (*schema.Organization, error) {
	organization := schema.Organization{}
	return &organization, o.OrganizationByName(ctx, name, &organization)
}

// DeleteOrganization deletes an organization.
func DeleteOrganization(ctx context.Context, id uint, o Organizations) (*schema.DeleteOrganizationInput, error) {
	result := schema.DeleteOrganizationInput{}
	return &result, o.DeleteOrg(ctx, id, &result)
}

// UpdateOrganization updates an organization.
func UpdateOrganization(ctx context.Context, id uint, patch schema.UpdateOrganizationPatchInput, o Organizations) (*schema.Organization, error) {
	result := schema.Organization{}
	return &result, o.UpdateOrganization(ctx, id, patch, &result)
}

// AllOrganizations lists all organizations.
func AllOrganizations(ctx context.Context, o Organizations) (*[]schema.Organization, error) {
	organization := []schema.Organization{}
	return &organization, o.AllOrganizations(ctx, &organization)
}

// AllOrganizationsExtended lists all organizations with associated projects.
func AllOrganizationsExtended(ctx context.Context, o Organizations) (*[]schema.Organization, error) {
	organization := []schema.Organization{}
	return &organization, o.AllOrganizationsExtended(ctx, &organization)
}
