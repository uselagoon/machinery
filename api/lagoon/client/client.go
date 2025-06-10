// Package client implements the interfaces required by the parent lagoon
// package.
package client

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

//go:embed _lgraphql/*
var lgraphql embed.FS

// Client implements the lagoon package interfaces for the Lagoon GraphQL API.
type Client struct {
	userAgent string
	token     *string
	endpoint  string
	client    *graphql.Client
	version   string // version of the lagoon api you're connecting to
}

// New creates a new Client for the given endpoint.
func New(endpoint, userAgent, version string, token *string, debug bool) *Client {
	if debug {
		return &Client{
			version:   version,
			userAgent: userAgent,
			token:     token,
			endpoint:  endpoint,
			client: graphql.NewClient(endpoint,
				// enable debug logging to stderr
				func(c *graphql.Client) {
					l := log.New(os.Stderr, "graphql", 0)
					c.Log = func(s string) {
						l.Println(s)
					}
				}),
		}
	}
	return &Client{
		version:   version,
		userAgent: userAgent,
		token:     token,
		endpoint:  endpoint,
		client:    graphql.NewClient(endpoint),
	}
}

// newRequest constructs a graphql request.
// assetName is the name of the graphql query template in _graphql/.
// varStruct is converted to a map of variables for the template.
func (c *Client) newRequest(
	assetName string, varStruct interface{}) (*graphql.Request, error) {

	q, err := lgraphql.ReadFile(assetName)
	if err != nil {
		return nil, fmt.Errorf("couldn't get asset: %w", err)
	}

	return c.doRequest(string(q), varStruct)
}

func (c *Client) doRequest(query string, varStruct interface{}) (*graphql.Request, error) {
	vars, err := structToVarMap(varStruct)
	if err != nil {
		return nil, fmt.Errorf("couldn't convert struct to map: %w", err)
	}

	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}

	headers := map[string]string{
		"User-Agent":    fmt.Sprintf("lagoon-client: %s", c.userAgent),
		"Authorization": fmt.Sprintf("Bearer %s", *c.token),
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

// structToVarMap encodes the given struct to a map. The idea is that by
// round-tripping through Marshal/Unmarshal, omitempty is applied to the
// zero-valued fields.
func structToVarMap(
	varStruct interface{}) (vars map[string]interface{}, err error) {
	data, err := json.Marshal(varStruct)
	if err != nil {
		return vars, err
	}
	return vars, json.Unmarshal(data, &vars)
}

// ProcessRaw runs a custom query/mutation and returns the response from the API.
func (c *Client) ProcessRaw(ctx context.Context, query string, variables map[string]interface{}) (interface{}, error) {
	req, err := c.doRequest(query, variables)
	if err != nil {
		return nil, err
	}

	var response interface{}
	if runErr := c.client.Run(ctx, req, &response); runErr != nil {
		return nil, runErr
	}
	return response, nil
}
