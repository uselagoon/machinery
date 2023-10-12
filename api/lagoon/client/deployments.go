package client

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

func (c *Client) DeploymentsByBulkID(
	ctx context.Context, bulkID string, deployments *[]schema.Deployment) error {
	req, err := c.newRequest("_lgraphql/deployments/getDeploymentsByBulkID.graphql",
		map[string]interface{}{
			"bulkId": bulkID,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.Deployment `json:"deploymentsByBulkId"`
	}{
		Response: deployments,
	})
}

func (c *Client) DeploymentsByEnvironment(
	ctx context.Context, projectID uint, environmentName string, environment *schema.Environment) error {
	req, err := c.newRequest("_lgraphql/deployments/getDeploymentsForEnvironment.graphql",
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

// DeployEnvironmentLatest deploys a latest environment.
func (c *Client) DeployEnvironmentLatest(ctx context.Context,
	in *schema.DeployEnvironmentLatestInput, out *schema.DeployEnvironmentLatest) error {
	req, err := c.newRequest("_lgraphql/deployments/deployEnvironmentLatest.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeployEnvironmentPullrequest deploys a pullreguest.
func (c *Client) DeployEnvironmentPullrequest(ctx context.Context,
	in *schema.DeployEnvironmentPullrequestInput, out *schema.DeployEnvironmentPullrequest) error {
	req, err := c.newRequest("_lgraphql/deployments/deployEnvironmentPullrequest.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeployEnvironmentPromote promotes one environment into a new environment.
func (c *Client) DeployEnvironmentPromote(ctx context.Context,
	in *schema.DeployEnvironmentPromoteInput, out *schema.DeployEnvironmentPromote) error {
	req, err := c.newRequest("_lgraphql/deployments/deployEnvironmentPromote.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeployEnvironmentBranch deploys a branch.
func (c *Client) DeployEnvironmentBranch(ctx context.Context,
	in *schema.DeployEnvironmentBranchInput, out *schema.DeployEnvironmentBranch) error {
	req, err := c.newRequest("_lgraphql/deployments/deployEnvironmentBranch.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

func (c *Client) DeploymentByRemoteID(
	ctx context.Context, remoteID string, deployment *schema.Deployment) error {
	req, err := c.newRequest("_lgraphql/deployments/deploymentByRemoteID.graphql",
		map[string]interface{}{
			"remoteId": remoteID,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Deployment `json:"deploymentByRemoteID"`
	}{
		Response: deployment,
	})
}

func (c *Client) DeploymentByName(
	ctx context.Context, projectName, environmentName, deploymentName string, logs bool, deployment *schema.Deployment) error {
	// use template request
	tmpl := "_lgraphql/deployments/deploymentByName.graphql"
	if logs {
		tmpl = "_lgraphql/deployments/deploymentByNameWithLog.graphql"
	}
	req, err := c.newRequest(tmpl,
		map[string]interface{}{
			"projectName":     projectName,
			"environmentName": environmentName,
			"name":            deploymentName,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Deployment `json:"deploymentByName"`
	}{
		Response: deployment,
	})
}

// UpdateDeployment updates a deployment.
func (c *Client) UpdateDeployment(
	ctx context.Context, id int, patch schema.UpdateDeploymentPatchInput, deployment *schema.Deployment) error {

	req, err := c.newRequest("_lgraphql/deployments/updateDeployment.graphql",
		map[string]interface{}{
			"id":    id,
			"patch": patch,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Deployment `json:"updateDeployment"`
	}{
		Response: deployment,
	})
}
