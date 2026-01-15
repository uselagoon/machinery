package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/uselagoon/machinery/api/schema"
)

// GetTaskByID gets a task by ID.
func (c *Client) GetTaskByID(
	ctx context.Context, id int, task *schema.Task) error {

	req, err := c.newRequest("_lgraphql/tasks/taskByID.graphql",
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Task `json:"taskById"`
	}{
		Response: task,
	})
}

// UpdateTask updates a task.
func (c *Client) UpdateTask(
	ctx context.Context, id int, patch schema.UpdateTaskPatchInput, task *schema.Task) error {

	req, err := c.newRequest("_lgraphql/tasks/updateTask.graphql",
		map[string]interface{}{
			"id":    id,
			"patch": patch,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Task `json:"updateTask"`
	}{
		Response: task,
	})
}

// RunActiveStandbySwitch runs Active/Standby Switch task.
func (c *Client) RunActiveStandbySwitch(
	ctx context.Context, project string, out *schema.Task) error {

	req, err := c.newRequest("_lgraphql/tasks/switchActiveStandby.graphql",
		map[string]interface{}{
			"project": project,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Task `json:"switchActiveStandby"`
	}{
		Response: out,
	})
}

// TaskFileUploadForm gets form data needed to upload task files directly to S3 compatible endpoint.
func (c *Client) TaskFileUploadForm(
	ctx context.Context, task int, filename string, out *schema.FileUploadForm) error {

	req, err := c.newRequest("_lgraphql/tasks/getTaskFileUploadForm.graphql",
		map[string]interface{}{
			"task":     task,
			"filename": filename,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.FileUploadForm `json:"getTaskFileUploadForm"`
	}{
		Response: out,
	})
}

// UploadFileForTask uploads a single file to a task.
func (c *Client) UploadFileForTask(
	ctx context.Context, task int, file string) error {

	filename := filepath.Base(file)

	// Get direct S3 upload info.
	formData := schema.FileUploadForm{}
	err := c.TaskFileUploadForm(ctx, task, filename, &formData)
	if err != nil {
		return fmt.Errorf("couldn't get upload form data: %w", err)
	}

	// Prepare multipart/form-data HTTP request.
	requestForm := new(bytes.Buffer)
	writer := multipart.NewWriter(requestForm)

	// Set generic S3 form fields.
	for name, value := range formData.FormFields {
		err := writer.WriteField(name, value)
		if err != nil {
			return fmt.Errorf("couldn't write upload form field %s: %w", name, err)
		}
	}

	// Set file field.
	fileField, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return fmt.Errorf("couldn't create upload file field: %w", err)
	}

	fd, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("couldn't read file %s: %w", file, err)
	}
	defer fd.Close()

	_, err = io.Copy(fileField, fd)
	if err != nil {
		return fmt.Errorf("couldn't copy file %s: %w", file, err)
	}

	writer.Close()

	// Upload the file.
	client := &http.Client{}
	req, err := http.NewRequest("POST", formData.PostUrl, requestForm)
	if err != nil {
		return fmt.Errorf("couldn't create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", fmt.Sprintf("lagoon-client: %s", c.userAgent))
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("couldn't upload file: %w", err)
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("couldn't read response: %w", err)
	}

	// Some S3 compatible services can respond with `204 No Content` for successful upload.
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("couldn't upload file (HTTP %d): %s", resp.StatusCode, bodyText)
	}

	return nil
}

// UploadFilesForTask uploads multiple files to a task.
func (c *Client) UploadFilesForTask(
	ctx context.Context, task int, files []string, out *schema.Task) error {

	for _, file := range files {
		err := c.UploadFileForTask(ctx, task, file)
		if err != nil {
			return fmt.Errorf("couldn't upload file %s: %w", filepath.Base(file), err)
		}
	}

	// Load return data.
	err := c.GetTaskByID(ctx, task, out)
	if err != nil {
		return fmt.Errorf("couldn't get task: %w", err)
	}

	return nil
}

// TasksByEnvironment gets tasks for an environment by project ID.
func (c *Client) TasksByEnvironment(
	ctx context.Context,
	projectID uint,
	environmentName string,
	environment *schema.Environment,
) error {
	req, err := c.newRequest("_lgraphql/tasks/getTasksForEnvironment.graphql",
		map[string]interface{}{
			"project":     projectID,
			"environment": environmentName,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"environmentByName"`
	}{
		Response: environment,
	})
}

// TasksByEnvironmentAndProjectName gets tasks for an environment by project name.
func (c *Client) TasksByEnvironmentAndProjectName(ctx context.Context,
	environmentName string,
	projectName string,
	environment *schema.Environment,
) error {
	project := &schema.Project{}
	if err := c.veryMinimalProjectByName(ctx, projectName, project); err != nil {
		return err
	}
	if project.Name == "" {
		//lint:ignore ST1005 return a generic Lagoon API unauthorized error based on the permission called
		// this is because organizationbyname will return null instead of an error, the api should probably return an error
		return fmt.Errorf(`Unauthorized: You don't have permission to "view" on "project"`) //nolint:staticcheck
	}
	return c.TasksByEnvironment(ctx, project.ID, environmentName, environment)
}

// InvokableAdvancedTaskDefinitionsByEnvironment gets advanced task definitions by project ID.
func (c *Client) InvokableAdvancedTaskDefinitionsByEnvironment(
	ctx context.Context,
	projectID uint,
	environmentName string,
	environment *schema.Environment,
) error {
	req, err := c.newRequest("_lgraphql/tasks/getInvokableAdvancedTaskDefinitionsByEnvironment.graphql",
		map[string]interface{}{
			"project":     projectID,
			"environment": environmentName,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"environmentByName"`
	}{
		Response: environment,
	})
}

// InvokableAdvancedTaskDefinitionsByEnvironmentAndProjectName gets advanced task definitions by
// project name.
func (c *Client) InvokableAdvancedTaskDefinitionsByEnvironmentAndProjectName(
	ctx context.Context,
	projectName string,
	environmentName string,
	environment *schema.Environment,
) error {
	project := &schema.Project{}
	if err := c.veryMinimalProjectByName(ctx, projectName, project); err != nil {
		return err
	}
	if project.Name == "" {
		//lint:ignore ST1005 return a generic Lagoon API unauthorized error based on the permission called
		// this is because organizationbyname will return null instead of an error, the api should probably return an error
		return fmt.Errorf(`Unauthorized: You don't have permission to "view" on "project"`) //nolint:staticcheck
	}
	return c.InvokableAdvancedTaskDefinitionsByEnvironment(ctx, project.ID, environmentName, environment)
}

// InvokeAdvancedTaskDefinition invokes an advanced task definition.
func (c *Client) InvokeAdvancedTaskDefinition(
	ctx context.Context, environmentID uint, taskID uint, task *schema.Task) error {

	req, err := c.newRequest("_lgraphql/tasks/invokeRegisteredTask.graphql",
		map[string]interface{}{
			"environment":            environmentID,
			"advancedTaskDefinition": taskID,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Task `json:"invokeRegisteredTask"`
	}{
		Response: task,
	})
}

// AdvancedTasksByEnvironment gets advanced tasks for an environment.
func (c *Client) AdvancedTasksByEnvironment(
	ctx context.Context,
	projectID uint,
	environmentName string,
	environment *schema.Environment,
) error {
	req, err := c.newRequest("_lgraphql/tasks/getAdvancedTasksByEnvironment.graphql",
		map[string]interface{}{
			"project": projectID,
			"name":    environmentName,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"environmentByName"`
	}{
		Response: environment,
	})
}

// AddTask adds a task.
func (c *Client) AddTask(
	ctx context.Context, environmentID uint, task schema.Task, out *schema.Task) error {

	req, err := c.newRequest("_lgraphql/tasks/addTask.graphql",
		map[string]interface{}{
			"environment": environmentID,
			"name":        task.Name,
			"command":     task.Command,
			"service":     task.Service,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Task `json:"addTask"`
	}{
		Response: out,
	})
}
