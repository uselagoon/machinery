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
	TaskFileUploadForm(ctx context.Context, id int, filename string, result *schema.FileUploadForm) error
	UploadFileForTask(ctx context.Context, id int, file string) error
	UploadFilesForTask(ctx context.Context, id int, files []string, result *schema.Task) error
	TasksByEnvironment(ctx context.Context, projectID uint, environmentName string, environment *schema.Environment) error
	InvokableAdvancedTaskDefinitionsByEnvironment(ctx context.Context, projectID uint, environmentName string, environment *schema.Environment) error
	InvokeAdvancedTaskDefinition(ctx context.Context, environmentID uint, taskID uint, result *schema.Task) error
	AdvancedTasksByEnvironment(ctx context.Context, projectID uint, environmentName string, environment *schema.Environment) error
	AddTask(ctx context.Context, environmentID uint, task schema.Task, result *schema.Task) error

	TasksByEnvironmentAndProjectName(ctx context.Context, environmentName string, projectName string, environment *schema.Environment) error
	InvokableAdvancedTaskDefinitionsByEnvironmentAndProjectName(ctx context.Context, projectName, environmentName string, environment *schema.Environment) error
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

// TaskFileUploadForm gets form fields required to upload a file to a task.
func TaskFileUploadForm(ctx context.Context, id int, filename string, t Tasks) (*schema.FileUploadForm, error) {
	result := schema.FileUploadForm{}
	return &result, t.TaskFileUploadForm(ctx, id, filename, &result)
}

// UploadFileForTask directly uploads a file to a task.
func UploadFileForTask(ctx context.Context, id int, file string, t Tasks) error {
	return t.UploadFileForTask(ctx, id, file)
}

// UploadFilesForTask directly uploads files to a task.
func UploadFilesForTask(ctx context.Context, id int, files []string, t Tasks) (*schema.Task, error) {
	result := schema.Task{}
	return &result, t.UploadFilesForTask(ctx, id, files, &result)
}

// GetTasksByEnvironment gets tasks for an environment.
func GetTasksByEnvironment(ctx context.Context, projectID uint, environmentName string, t Tasks) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, t.TasksByEnvironment(ctx, projectID, environmentName, &environment)
}

// GetTasksByEnvironmentAndProjectName gets tasks for an environment.
func GetTasksByEnvironmentAndProjectName(ctx context.Context, projectName, environmentName string, t Tasks) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, t.TasksByEnvironmentAndProjectName(ctx, environmentName, projectName, &environment)
}

// GetInvokableAdvancedTaskDefinitionsByEnvironment gets a list of tasks invokable against an environment.
func GetInvokableAdvancedTaskDefinitionsByEnvironment(ctx context.Context, projectID uint, environmentName string, t Tasks) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, t.InvokableAdvancedTaskDefinitionsByEnvironment(ctx, projectID, environmentName, &environment)
}

// GetInvokableAdvancedTaskDefinitionsByEnvironment gets a list of tasks invokable against an environment.
func GetInvokableAdvancedTaskDefinitionsByEnvironmentAndProjectName(ctx context.Context, projectName, environmentName string, t Tasks) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, t.InvokableAdvancedTaskDefinitionsByEnvironmentAndProjectName(ctx, projectName, environmentName, &environment)
}

// InvokeAdvancedTaskDefinition invokes an advanced task definition.
func InvokeAdvancedTaskDefinition(ctx context.Context, environmentID uint, taskID uint, t Tasks) (*schema.Task, error) {
	result := schema.Task{}
	return &result, t.InvokeAdvancedTaskDefinition(ctx, environmentID, taskID, &result)
}

// GetAdvancedTasksByEnvironment gets advanced tasks for an environment.
func GetAdvancedTasksByEnvironment(ctx context.Context, projectID uint, environmentName string, t Tasks) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, t.AdvancedTasksByEnvironment(ctx, projectID, environmentName, &environment)
}

// AddTask adds a task to an environment.
func AddTask(ctx context.Context, environmentID uint, task schema.Task, t Tasks) (*schema.Task, error) {
	result := schema.Task{}
	return &result, t.AddTask(ctx, environmentID, task, &result)
}
