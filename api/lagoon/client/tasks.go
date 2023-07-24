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

	"github.com/uselagoon/machinery/api/lagoon/client/lgraphql"
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
	q, err := lgraphql.Asset("_lgraphql/tasks/uploadFilesForTask.graphql")
	if err != nil {
		return fmt.Errorf("couldn't get graphql asset from assets: %w", err)
	}
	_, err = formField.Write([]byte(fmt.Sprintf("{ \"query\": \"%s\", \"variables\": { \"task\": %d, \"files\": [null] } }", string(q), task)))

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
