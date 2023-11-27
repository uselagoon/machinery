package client

import (
	"context"
	"fmt"
	"regexp"

	"github.com/uselagoon/machinery/api/lagoon"
	"github.com/uselagoon/machinery/api/schema"
)

var duplicate = regexp.MustCompile("(?s).*graphql.*Duplicate entry.* ")

// wrapErr wraps a response error with a lagoon.ErrExist type if the response
// is due to an object already existing
func wrapErr(err error) error {
	if err != nil && duplicate.MatchString(err.Error()) {
		return fmt.Errorf("couldn't create object: %w: %v", lagoon.ErrExist, err)
	}
	return err
}

// AddEnvVariable adds an EnvVariable to an Environment or Project.
func (c *Client) AddEnvVariable(ctx context.Context,
	in *schema.EnvVariableInput, out *schema.EnvKeyValue) error {
	req, err := c.newRequest("_lgraphql/addEnvVariable.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.EnvKeyValue `json:"addEnvVariable"`
	}{
		Response: out,
	})
}
