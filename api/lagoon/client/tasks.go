package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/uselagoon/machinery/api/schema"
)

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

// RunActiveStandbySwitch deploys a branch.
func (c *Client) RunActiveStandbySwitch(ctx context.Context,
	project string, out *schema.Task) error {
	req, err := c.newRequest("_lgraphql/tasks/switchActiveStandby.graphql", map[string]interface{}{
		"project": project,
	})
	if err != nil {
		return err
	}

	// return c.client.Run(ctx, req, &out)
	return c.client.Run(ctx, req, &struct {
		Response *schema.Task `json:"switchActiveStandby"`
	}{
		Response: out,
	})
}

// UploadFilesForTask uploads a file to a task.
func (c *Client) UploadFilesForTask(ctx context.Context,
	task int, files []string, out *schema.Task) error {

	// uploading files via graphql is different, so this entire function is crafted specifically for uploading files
	// it was built using the curl example from lagoon. this one supports multiple file uploads though
	// machinebox graphql client doesn't support file uploads that conform to graphql spec, so this is custom compared to other query/mutations
	// curl -sS "$TASK_API_HOST"/graphql \
	// -H "Authorization: Bearer $TOKEN" \
	// -F operations='{ "query": "mutation ($task: Int!, $files: [Upload!]!) { uploadFilesForTask(input:{task:$task, files:$files}) { id files { filename } } }", "variables": { "task": '"$TASK_DATA_ID"', "files": [null] } }' \
	// -F map='{ "0": ["variables.files.0"] }' \
	// -F 0=@$file.gz

	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)
	formField, err := writer.CreateFormField("operations")
	if err != nil {
		return fmt.Errorf("couldn't create operations form field: %w", err)
	}

	// this asset is specific for this, it has to be one line if changes are made
	q, err := lgraphql.ReadFile("_lgraphql/tasks/uploadFilesForTask.graphql")
	if err != nil {
		return fmt.Errorf("couldn't get graphql asset from assets: %w", err)
	}
	_, err = fmt.Fprintf(formField, "{ \"query\": \"%s\", \"variables\": { \"task\": %d, \"files\": [null] } }", string(q), task)
	if err != nil {
		return fmt.Errorf("couldn't create upload map form field: %w", err)
	}

	formField, err = writer.CreateFormField("map")
	if err != nil {
		return fmt.Errorf("couldn't create upload map form field: %w", err)
	}

	// adding the files is done *after* the map formfield is written
	fileMap := make(map[string][]string)
	for idx := range files {
		fileMap[fmt.Sprintf("%d", idx)] = append(fileMap[fmt.Sprintf("%d", idx)], fmt.Sprintf("variables.files.%d", idx))
	}
	ffmBytes, err := json.Marshal(fileMap)
	if err != nil {
		return fmt.Errorf("couldn't create upload map form field: %w", err)
	}
	_, err = formField.Write(ffmBytes)
	if err != nil {
		return fmt.Errorf("couldn't create upload map form field: %w", err)
	}
	for idx, file := range files {
		fileMap[fmt.Sprintf("%d", idx)] = append(fileMap[fmt.Sprintf("%d", idx)], fmt.Sprintf("variables.files.%d", idx))

		fw, err := writer.CreateFormFile(fmt.Sprintf("%d", idx), filepath.Base(file))
		if err != nil {
			return fmt.Errorf("couldn't read file %s: %w", file, err)
		}
		fd, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("couldn't read file %s: %w", file, err)
		}
		defer fd.Close()
		_, err = io.Copy(fw, fd)
		if err != nil {
			return fmt.Errorf("couldn't copy file %s: %w", file, err)
		}
	}

	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", c.endpoint, form)
	if err != nil {
		return fmt.Errorf("couldn't create API request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+*c.token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", fmt.Sprintf("lagoon-client: %s", c.userAgent))
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("couldn't post file(s) to API: %w", err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("couldn't read response from API: %w", err)
	}

	// unmarshal the response into a map
	var respDat map[string]map[string]interface{}
	err = json.Unmarshal([]byte(bodyText), &respDat)
	if err != nil {
		return fmt.Errorf("couldn't unmarshal response from API: %w", err)
	}

	// then extract our specific uploadFilesForTask data
	b, _ := json.Marshal(respDat["data"]["uploadFilesForTask"])
	err = json.Unmarshal([]byte(b), out)
	if err != nil {
		return fmt.Errorf("couldn't unmarshal response from API: %w", err)
	}

	return nil
}

// TasksByEnvironment gets tasks for an environment.
func (c *Client) TasksByEnvironment(ctx context.Context, projectID uint, environmentName string, environment *schema.Environment) error {
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

// TasksByEnvironmentAndProjectName is the same as TasksByEnvironment but does a project lookup first
func (c *Client) TasksByEnvironmentAndProjectName(ctx context.Context, environmentName string,
	projectName string, environment *schema.Environment) error {
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

// InvokableAdvancedTaskDefinitionsByEnvironment gets tasks for an environment.
func (c *Client) InvokableAdvancedTaskDefinitionsByEnvironment(ctx context.Context, projectID uint, environmentName string, environment *schema.Environment) error {
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

// InvokableAdvancedTaskDefinitionsByEnvironmentAndProjectName is the same as InvokableAdvancedTaskDefinitionsByEnvironment but does a project lookup first
func (c *Client) InvokableAdvancedTaskDefinitionsByEnvironmentAndProjectName(ctx context.Context, projectName string,
	environmentName string, environment *schema.Environment) error {
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
func (c *Client) InvokeAdvancedTaskDefinition(ctx context.Context, environmentID uint, taskID uint, task *schema.Task) error {
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
func (c *Client) AdvancedTasksByEnvironment(ctx context.Context, projectID uint, environmentName string, environment *schema.Environment) error {
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
func (c *Client) AddTask(ctx context.Context, environmentID uint, task schema.Task, out *schema.Task) error {
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
