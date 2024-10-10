// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

// DeployTargetConfigs interface contains methods for getting info on deploytarget configs.
type DeployTargetConfigs interface {
	DeployTargetConfigsByProjectID(ctx context.Context, project uint, deployTargets *[]schema.DeployTargetConfig) error
	AddDeployTargetConfiguration(ctx context.Context, in *schema.AddDeployTargetConfigInput, deployTargets *schema.DeployTargetConfig) error
	UpdateDeployTargetConfiguration(ctx context.Context, in *schema.UpdateDeployTargetConfigInput, deployTargets *schema.DeployTargetConfig) error
	DeleteDeployTargetConfiguration(ctx context.Context, id, project uint, deployTargets *schema.DeleteDeployTargetConfig) error

	DeployTargetConfigsByProjectName(ctx context.Context, project string, deployTargets *[]schema.DeployTargetConfig) error
	AddDeployTargetConfigurationByProjectName(ctx context.Context, project string, in *schema.AddDeployTargetConfigInput, deployTargets *schema.DeployTargetConfig) error
	DeleteDeployTargetConfigurationByIDAndProjectName(ctx context.Context, id uint, project string, deployTargets *schema.DeleteDeployTargetConfig) error
}

// GetDeployTargetConfigs gets deploytarget configs for a specific project.
func GetDeployTargetConfigs(ctx context.Context, project uint, dtc DeployTargetConfigs) (*[]schema.DeployTargetConfig, error) {
	deployTargets := []schema.DeployTargetConfig{}
	return &deployTargets, dtc.DeployTargetConfigsByProjectID(ctx, project, &deployTargets)
}

// GetDeployTargetConfigs gets deploytarget configs for a specific project.
func GetDeployTargetConfigsByProjectName(ctx context.Context, project string, dtc DeployTargetConfigs) (*[]schema.DeployTargetConfig, error) {
	deployTargets := []schema.DeployTargetConfig{}
	return &deployTargets, dtc.DeployTargetConfigsByProjectName(ctx, project, &deployTargets)
}

// AddDeployTargetConfiguration adds a deploytarget config to a specific project.
func AddDeployTargetConfiguration(ctx context.Context, in *schema.AddDeployTargetConfigInput, dtc DeployTargetConfigs) (*schema.DeployTargetConfig, error) {
	deployTarget := schema.DeployTargetConfig{}
	return &deployTarget, dtc.AddDeployTargetConfiguration(ctx, in, &deployTarget)
}

// AddDeployTargetConfiguration adds a deploytarget config to a specific project.
func AddDeployTargetConfigurationByProjectName(ctx context.Context, project string, in *schema.AddDeployTargetConfigInput, dtc DeployTargetConfigs) (*schema.DeployTargetConfig, error) {
	deployTarget := schema.DeployTargetConfig{}
	return &deployTarget, dtc.AddDeployTargetConfigurationByProjectName(ctx, project, in, &deployTarget)
}

// UpdateDeployTargetConfiguration adds a deploytarget config to a specific project.
func UpdateDeployTargetConfiguration(ctx context.Context, in *schema.UpdateDeployTargetConfigInput, dtc DeployTargetConfigs) (*schema.DeployTargetConfig, error) {
	deployTarget := schema.DeployTargetConfig{}
	return &deployTarget, dtc.UpdateDeployTargetConfiguration(ctx, in, &deployTarget)
}

// DeleteDeployTargetConfiguration deletes a deploytarget config from a specific project.
func DeleteDeployTargetConfiguration(ctx context.Context, id uint, project uint, dtc DeployTargetConfigs) (*schema.DeleteDeployTargetConfig, error) {
	deployTarget := schema.DeleteDeployTargetConfig{}
	return &deployTarget, dtc.DeleteDeployTargetConfiguration(ctx, id, project, &deployTarget)
}

// DeleteDeployTargetConfigurationByIDAndProjectName deletes a deploytarget config from a specific project.
func DeleteDeployTargetConfigurationByIDAndProjectName(ctx context.Context, id uint, project string, dtc DeployTargetConfigs) (*schema.DeleteDeployTargetConfig, error) {
	deployTarget := schema.DeleteDeployTargetConfig{}
	return &deployTarget, dtc.DeleteDeployTargetConfigurationByIDAndProjectName(ctx, id, project, &deployTarget)
}
