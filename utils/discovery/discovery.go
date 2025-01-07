package discovery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type DiscoveryEndpoint struct {
	LagoonVersion         string           `json:"lagoon_version"`
	AuthorizationEndpoint string           `json:"authorization_endpoint"`
	SSHTokenExchange      SSHTokenExchange `json:"ssh_token_exchange"`
	WebhookEndpoint       string           `json:"webhook_endpoint"`
	UIHostname            string           `json:"ui_url"`
}

type SSHTokenExchange struct {
	TokenHost string `json:"token_endpoint_host"`
	TokenPort int    `json:"token_endpoint_port"`
}

// the wellknown path of the lagoon discovery endpoint
const wellKnown = ".well-known/appspecific/sh.lagoon.discovery.json"

func Discover(hostname string) (*DiscoveryEndpoint, error) {
	// check if the hostname provided is actually valid-ish
	u, err := url.Parse(strings.TrimRight(hostname, "/"))
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case "":
		u.Scheme = "https"
	case "http", "https":
	default:
		return nil, fmt.Errorf("url scheme must be http or https not %s", u.Scheme)

	}
	// check the discovery endpoint
	resp, err := http.Get(fmt.Sprintf("%s/%s", u, wellKnown))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// decode the response into the discovery struct
	d := &DiscoveryEndpoint{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(d)
	if err != nil {
		return nil, err
	}
	return d, nil
}
