// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"
	"github.com/uselagoon/machinery/api/schema"
)

type Variables interface {
	EnvVariablesByProjectEnvironmentName(ctx context.Context, in *schema.EnvVariableByProjectEnvironmentNameInput, envvar *[]schema.EnvKeyValue) error
}

// GetEnvVariablesByProjectEnvironmentName lists variables given a project/environment.
func GetEnvVariablesByProjectEnvironmentName(ctx context.Context, in *schema.EnvVariableByProjectEnvironmentNameInput, v Variables) (*[]schema.EnvKeyValue, error) {
	envvar := []schema.EnvKeyValue{}
	return &envvar, v.EnvVariablesByProjectEnvironmentName(ctx, in, &envvar)
}
