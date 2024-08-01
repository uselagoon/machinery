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
	UpdateStorageOnEnvironment(ctx context.Context, storage *schema.UpdateStorageOnEnvironmentInput, result *schema.UpdateStorageOnEnvironment) error
	DeleteEnvironment(ctx context.Context, name, project string, execute bool, result *schema.DeleteEnvironment) error
	UpdateEnvironment(ctx context.Context, id uint, patch schema.UpdateEnvironmentPatchInput, result *schema.Environment) error
	SetEnvironmentServices(ctx context.Context, id uint, services []string, result *[]schema.EnvironmentService) error
	EnvironmentByName(ctx context.Context, name string, project uint, result *schema.Environment) error
	EnvironmentByID(ctx context.Context, environment uint, result *schema.Environment) error
	EnvironmentByNamespace(ctx context.Context, namespace string, result *schema.Environment) error
	BackupsByEnvironmentNamespace(ctx context.Context, namespace string, result *schema.Environment) error
	EnvironmentsByProjectName(ctx context.Context, project string, result *[]schema.Environment) error
	SSHEndpointByNamespace(ctx context.Context, namespace string, result *schema.Environment) error
	AddOrUpdateEnvironmentService(ctx context.Context, service schema.AddEnvironmentServiceInput, result *schema.EnvironmentService) error
	DeleteEnvironmentService(ctx context.Context, service schema.DeleteEnvironmentServiceInput, result *schema.DeleteEnvironmentService) error
	DeleteBackup(context.Context, string, *schema.DeleteBackup) error
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

// UpdateStorage updates environment storage.
func UpdateStorageOnEnvironment(ctx context.Context, storage *schema.UpdateStorageOnEnvironmentInput, e Environments) (*schema.UpdateStorageOnEnvironment, error) {
	result := schema.UpdateStorageOnEnvironment{}
	return &result, e.UpdateStorageOnEnvironment(ctx, storage, &result)
}

// DeleteEnvironment deletes an environment.
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
func GetEnvironmentByID(ctx context.Context, env uint, e Environments) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, e.EnvironmentByID(ctx, env, &environment)
}

// GetMinimalProjectByName gets info of projects in lagoon that have matching metadata.
func GetEnvironmentByNamespace(ctx context.Context, namespace string, e Environments) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, e.EnvironmentByNamespace(ctx, namespace, &environment)
}

// GetMinimalProjectByName gets info of projects in lagoon that have matching metadata.
func GetBackupsByEnvironmentNamespace(ctx context.Context, namespace string, e Environments) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, e.BackupsByEnvironmentNamespace(ctx, namespace, &environment)
}

// GetEnvironmentsByProjectName lists info of all environments for a given project.
func GetEnvironmentsByProjectName(ctx context.Context, project string, e Environments) (*[]schema.Environment, error) {
	environment := []schema.Environment{}
	return &environment, e.EnvironmentsByProjectName(ctx, project, &environment)
}

// SSHEndpointByNamespace gets info of projects in lagoon that have matching metadata.
func SSHEndpointByNamespace(ctx context.Context, namespace string, e Environments) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, e.SSHEndpointByNamespace(ctx, namespace, &environment)
}

// AddOrUpdateEnvironmentService updates an environments service.
func AddOrUpdateEnvironmentService(ctx context.Context, service schema.AddEnvironmentServiceInput, e Environments) (*schema.EnvironmentService, error) {
	result := schema.EnvironmentService{}
	return &result, e.AddOrUpdateEnvironmentService(ctx, service, &result)
}

// DeleteEnvironmentService deletes an environment service.
func DeleteEnvironmentService(ctx context.Context, service schema.DeleteEnvironmentServiceInput, e Environments) (*schema.DeleteEnvironmentService, error) {
	result := schema.DeleteEnvironmentService{}
	return &result, e.DeleteEnvironmentService(ctx, service, &result)
}

// DeleteBackup deletes an environment backup.
func DeleteBackup(ctx context.Context, backupID string, e Environments) (*schema.DeleteBackup, error) {
	result := schema.DeleteBackup{}
	return &result, e.DeleteBackup(ctx, backupID, &result)
}
