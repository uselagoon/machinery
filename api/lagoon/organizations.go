// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

// Projects interface contains methods for getting info on projects.
type Organizations interface {
	OrganizationByID(ctx context.Context, id int, organization *schema.Organization) error
}

// GetOrganizationByID gets info of an organization in lagoon that matches the provided ID.
func GetOrganizationByID(ctx context.Context, id int, o Organizations) (*schema.Organization, error) {
	organization := schema.Organization{}
	return &organization, o.OrganizationByID(ctx, id, &organization)
}
