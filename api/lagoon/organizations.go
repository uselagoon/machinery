// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

// Organizations interface contains methods for getting info on projects.
type Organizations interface {
	OrganizationByID(ctx context.Context, id int, organization *schema.Organization) error
	DeleteOrg(ctx context.Context, id int, organization *schema.DeleteOrganizationInput) error
}

// GetOrganizationByID gets info of an organization in lagoon that matches the provided ID.
func GetOrganizationByID(ctx context.Context, id int, o Organizations) (*schema.Organization, error) {
	organization := schema.Organization{}
	return &organization, o.OrganizationByID(ctx, id, &organization)
}

// DeleteOrganization updates environment storage.
func DeleteOrganization(ctx context.Context, id int, o Organizations) (*schema.DeleteOrganizationInput, error) {
	result := schema.DeleteOrganizationInput{}
	return &result, o.DeleteOrg(ctx, id, &result)
}
