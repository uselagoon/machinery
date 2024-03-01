// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"
	"github.com/uselagoon/machinery/api/schema"
)

type Variables interface {
	AddOrUpdateEnvVariableByName(ctx context.Context, in *schema.EnvVariableByNameInput, envvar *schema.UpdateEnvVarResponse) error
	DeleteEnvVariableByName(ctx context.Context, in *schema.DeleteEnvVariableByNameInput, envvar *schema.DeleteEnvVarResponse) error
	GetEnvVariablesByProjectEnvironmentName(ctx context.Context, in *schema.EnvVariableByProjectEnvironmentNameInput, envvar *[]schema.EnvKeyValue) error
}

func AddOrUpdateEnvVariableByName(ctx context.Context, in *schema.EnvVariableByNameInput, v Variables) (*schema.UpdateEnvVarResponse, error) {
	envvar := schema.UpdateEnvVarResponse{}
	return &envvar, v.AddOrUpdateEnvVariableByName(ctx, in, &envvar)
}

func DeleteEnvVariableByName(ctx context.Context, in *schema.DeleteEnvVariableByNameInput, v Variables) (*schema.DeleteEnvVarResponse, error) {
	envvar := schema.DeleteEnvVarResponse{}
	return &envvar, v.DeleteEnvVariableByName(ctx, in, &envvar)
}

// GetEnvVariablesByProjectEnvironmentName lists variables given a project/environment.
func GetEnvVariablesByProjectEnvironmentName(ctx context.Context, in *schema.EnvVariableByProjectEnvironmentNameInput, v Variables) (*[]schema.EnvKeyValue, error) {
	envvar := []schema.EnvKeyValue{}
	return &envvar, v.GetEnvVariablesByProjectEnvironmentName(ctx, in, &envvar)
}
