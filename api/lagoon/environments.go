// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

// Environments interface contains methods for getting info on environments.
type Environments interface {
	BackupsForEnvironmentByName(context.Context, string, uint, *schema.Environment) error
	AddRestore(context.Context, string, *schema.Restore) error
	UpdateEnvironmentStorage(ctx context.Context, storage *schema.UpdateEnvironmentStorageInput, result *schema.UpdateEnvironmentStorage) error
	DeleteEnvironment(ctx context.Context, name, project string, execute bool, result *schema.DeleteEnvironment) error
	UpdateEnvironment(ctx context.Context, id uint, patch schema.UpdateEnvironmentPatchInput, result *schema.Environment) error
	SetEnvironmentServices(ctx context.Context, id uint, services []string, result *[]schema.EnvironmentService) error
	EnvironmentByName(ctx context.Context, name string, project uint, result *schema.Environment) error
	EnvironmentByNamespace(ctx context.Context, namespace string, result *schema.Environment) error
	SSHEndpointByNamespace(ctx context.Context, namespace string, result *schema.Environment) error
}

// GetBackupsForEnvironmentByName gets backup info in lagoon for specific environment.
func GetBackupsForEnvironmentByName(ctx context.Context, name string, project uint, e Environments) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, e.BackupsForEnvironmentByName(ctx, name, project, &environment)
}

// AddBackupRestore adds a backup restore based on backup ID.
func AddBackupRestore(ctx context.Context, backupID string, e Environments) (*schema.Restore, error) {
	restore := schema.Restore{}
	return &restore, e.AddRestore(ctx, backupID, &restore)
}

// UpdateStorage updates environment storage.
func UpdateStorage(ctx context.Context, storage *schema.UpdateEnvironmentStorageInput, e Environments) (*schema.UpdateEnvironmentStorage, error) {
	result := schema.UpdateEnvironmentStorage{}
	return &result, e.UpdateEnvironmentStorage(ctx, storage, &result)
}

// DeleteEnvironment updates environment storage.
func DeleteEnvironment(ctx context.Context, name, project string, execute bool, e Environments) (*schema.DeleteEnvironment, error) {
	result := schema.DeleteEnvironment{}
	return &result, e.DeleteEnvironment(ctx, name, project, execute, &result)
}

// UpdateEnvironment updates an environment.
func UpdateEnvironment(ctx context.Context, id uint, patch schema.UpdateEnvironmentPatchInput, e Environments) (*schema.Environment, error) {
	result := schema.Environment{}
	return &result, e.UpdateEnvironment(ctx, id, patch, &result)
}

// SetEnvironmentServices updates an environments services list.
func SetEnvironmentServices(ctx context.Context, id uint, services []string, e Environments) (*[]schema.EnvironmentService, error) {
	result := []schema.EnvironmentService{}
	return &result, e.SetEnvironmentServices(ctx, id, services, &result)
}

// GetMinimalProjectByName gets info of projects in lagoon that have matching metadata.
func GetEnvironmentByName(ctx context.Context, name string, project uint, e Environments) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, e.EnvironmentByName(ctx, name, project, &environment)
}

// GetMinimalProjectByName gets info of projects in lagoon that have matching metadata.
func GetEnvironmentByNamespace(ctx context.Context, namespace string, e Environments) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, e.EnvironmentByNamespace(ctx, namespace, &environment)
}

// SSHEndpointByNamespace gets info of projects in lagoon that have matching metadata.
func SSHEndpointByNamespace(ctx context.Context, namespace string, e Environments) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, e.SSHEndpointByNamespace(ctx, namespace, &environment)
}
