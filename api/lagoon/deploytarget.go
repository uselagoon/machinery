package lagoon

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

type DeployTargets interface {
	AddDeployTarget(ctx context.Context, in *schema.AddDeployTargetInput, out *schema.AddDeployTargetResponse) error
	UpdateDeployTarget(ctx context.Context, in *schema.UpdateDeployTargetInput, out *schema.UpdateDeployTargetResponse) error
	DeleteDeployTarget(ctx context.Context, in *schema.DeleteDeployTargetInput, out *schema.DeleteDeployTargetResponse) error
	ListDeployTargets(ctx context.Context, out *[]schema.DeployTarget) error
	DeployTargetsByOrganizationNameOrID(ctx context.Context, name *string, id *uint, out *[]schema.DeployTarget) error
	AddDeployTargetToOrganization(ctx context.Context, in *schema.AddDeployTargetToOrganizationInput, out *schema.AddDeployTargetResponse) error
	RemoveDeployTargetFromOrganization(ctx context.Context, in *schema.RemoveDeployTargetFromOrganizationInput, out *schema.DeleteDeployTargetResponse) error
}

func AddDeployTarget(ctx context.Context, in *schema.AddDeployTargetInput, out DeployTargets) (*schema.AddDeployTargetResponse, error) {
	response := schema.AddDeployTargetResponse{}
	return &response, out.AddDeployTarget(ctx, in, &response)
}

func UpdateDeployTarget(ctx context.Context, in *schema.UpdateDeployTargetInput, out DeployTargets) (*schema.UpdateDeployTargetResponse, error) {
	response := schema.UpdateDeployTargetResponse{}
	return &response, out.UpdateDeployTarget(ctx, in, &response)
}

func DeleteDeployTarget(ctx context.Context, in *schema.DeleteDeployTargetInput, out DeployTargets) (*schema.DeleteDeployTargetResponse, error) {
	response := schema.DeleteDeployTargetResponse{}
	return &response, out.DeleteDeployTarget(ctx, in, &response)
}

// ListDeployTargets gets info of deploytargets in lagoon.
func ListDeployTargets(ctx context.Context, out DeployTargets) (*[]schema.DeployTarget, error) {
	deploytargets := []schema.DeployTarget{}
	return &deploytargets, out.ListDeployTargets(ctx, &deploytargets)
}

// ListDeployTargetsByOrganizationNameOrID gets deploy targets associated with an organization in lagoon via provided Name & ID.
func ListDeployTargetsByOrganizationNameOrID(ctx context.Context, name *string, id *uint, d DeployTargets) (*[]schema.DeployTarget, error) {
	deploytargets := []schema.DeployTarget{}
	return &deploytargets, d.DeployTargetsByOrganizationNameOrID(ctx, name, id, &deploytargets)
}

// AddDeployTargetToOrganization adds a deploy target to an organization.
func AddDeployTargetToOrganization(ctx context.Context, in *schema.AddDeployTargetToOrganizationInput, out DeployTargets) (*schema.AddDeployTargetResponse, error) {
	response := schema.AddDeployTargetResponse{}
	return &response, out.AddDeployTargetToOrganization(ctx, in, &response)
}

// RemoveDeployTargetFromOrganization removes a deploy target from an organization.
func RemoveDeployTargetFromOrganization(ctx context.Context, in *schema.RemoveDeployTargetFromOrganizationInput, out DeployTargets) (*schema.DeleteDeployTargetResponse, error) {
	response := schema.DeleteDeployTargetResponse{}
	return &response, out.RemoveDeployTargetFromOrganization(ctx, in, &response)
}
