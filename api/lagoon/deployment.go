// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

// Deploy interface contains methods for deploying branches and environments in lagoon.
type Deploy interface {
	DeployEnvironmentPromote(ctx context.Context, deploy *schema.DeployEnvironmentPromoteInput, result *schema.DeployEnvironmentPromote) error
	DeployEnvironmentLatest(ctx context.Context, deploy *schema.DeployEnvironmentLatestInput, result *schema.DeployEnvironmentLatest) error
	DeployEnvironmentPullrequest(ctx context.Context, deploy *schema.DeployEnvironmentPullrequestInput, result *schema.DeployEnvironmentPullrequest) error
	DeployEnvironmentBranch(ctx context.Context, deploy *schema.DeployEnvironmentBranchInput, result *schema.DeployEnvironmentBranch) error
	DeploymentsByBulkID(ctx context.Context, bulkID string, deployments *[]schema.Deployment) error
	DeploymentsByEnvironment(ctx context.Context, projectID uint, environmentName string, environment *schema.Environment) error
	DeploymentByRemoteID(ctx context.Context, remoteID string, deployments *schema.Deployment) error
	DeploymentByName(ctx context.Context, projectName, environmentName, deploymentName string, logs bool, deployments *schema.Deployment) error
	UpdateDeployment(ctx context.Context, id int, patch schema.UpdateDeploymentPatchInput, result *schema.Deployment) error
}

// DeployLatest deploys the latest environment.
func DeployLatest(ctx context.Context, deploy *schema.DeployEnvironmentLatestInput, m Deploy) (*schema.DeployEnvironmentLatest, error) {
	result := schema.DeployEnvironmentLatest{}
	return &result, m.DeployEnvironmentLatest(ctx, deploy, &result)
}

// DeployPullRequest deploys a pull request.
func DeployPullRequest(ctx context.Context, deploy *schema.DeployEnvironmentPullrequestInput, m Deploy) (*schema.DeployEnvironmentPullrequest, error) {
	result := schema.DeployEnvironmentPullrequest{}
	return &result, m.DeployEnvironmentPullrequest(ctx, deploy, &result)
}

// DeployPromote promotes one environment into a new environment.
func DeployPromote(ctx context.Context, deploy *schema.DeployEnvironmentPromoteInput, m Deploy) (*schema.DeployEnvironmentPromote, error) {
	result := schema.DeployEnvironmentPromote{}
	return &result, m.DeployEnvironmentPromote(ctx, deploy, &result)
}

// DeployBranch deploys a branch.
func DeployBranch(ctx context.Context, deploy *schema.DeployEnvironmentBranchInput, m Deploy) (*schema.DeployEnvironmentBranch, error) {
	result := schema.DeployEnvironmentBranch{}
	return &result, m.DeployEnvironmentBranch(ctx, deploy, &result)
}

// GetDeploymentsByBulkID gets deployments for a given bulkID.
func GetDeploymentsByBulkID(ctx context.Context, bulkID string, d Deploy) (*[]schema.Deployment, error) {
	deployments := []schema.Deployment{}
	return &deployments, d.DeploymentsByBulkID(ctx, bulkID, &deployments)
}

// GetDeploymentsByEnvironment gets deployments for an envronment.
func GetDeploymentsByEnvironment(ctx context.Context, projectID uint, environmentName string, d Deploy) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, d.DeploymentsByEnvironment(ctx, projectID, environmentName, &environment)
}

// GetDeploymentsByBulkID gets deployments for a given bulkID.
func GetDeploymentByRemoteID(ctx context.Context, remoteID string, d Deploy) (*schema.Deployment, error) {
	deployment := schema.Deployment{}
	return &deployment, d.DeploymentByRemoteID(ctx, remoteID, &deployment)
}

// GetDeploymentByName gets deployment for a given build name.
func GetDeploymentByName(ctx context.Context, projectName, environmentName, deploymentName string, logs bool, d Deploy) (*schema.Deployment, error) {
	deployment := schema.Deployment{}
	return &deployment, d.DeploymentByName(ctx, projectName, environmentName, deploymentName, logs, &deployment)
}

// UpdateTask updates a task.
func UpdateDeployment(ctx context.Context, id int, patch schema.UpdateDeploymentPatchInput, d Deploy) (*schema.Deployment, error) {
	result := schema.Deployment{}
	return &result, d.UpdateDeployment(ctx, id, patch, &result)
}
