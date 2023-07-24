// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

// Tasks interface contains methods for running tasks in projects and environments in lagoon.
type Tasks interface {
	RunActiveStandbySwitch(ctx context.Context, project string, result *schema.Task) error
	GetTaskByID(ctx context.Context, id int, result *schema.Task) error
	UpdateTask(ctx context.Context, id int, patch schema.UpdateTaskPatchInput, result *schema.Task) error
	UploadFilesForTask(ctx context.Context, id int, files []string, result *schema.Task) error
}

// ActiveStandbySwitch runs the activestandby switch.
func ActiveStandbySwitch(ctx context.Context, project string, t Tasks) (*schema.Task, error) {
	result := schema.Task{}
	return &result, t.RunActiveStandbySwitch(ctx, project, &result)
}

// TaskByID returns a task by the associated id
func TaskByID(ctx context.Context, id int, t Tasks) (*schema.Task, error) {
	result := schema.Task{}
	return &result, t.GetTaskByID(ctx, id, &result)
}

// UpdateTask updates a task.
func UpdateTask(ctx context.Context, id int, patch schema.UpdateTaskPatchInput, t Tasks) (*schema.Task, error) {
	result := schema.Task{}
	return &result, t.UpdateTask(ctx, id, patch, &result)
}

// UploadFilesForTask updates a task.
func UploadFilesForTask(ctx context.Context, id int, files []string, t Tasks) (*schema.Task, error) {
	result := schema.Task{}
	return &result, t.UploadFilesForTask(ctx, id, files, &result)
}
