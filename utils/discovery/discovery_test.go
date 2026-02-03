package discovery

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestDiscover(t *testing.T) {
	type args struct {
		testserverpayload string
	}
	tests := []struct {
		name    string
		args    args
		want    *DiscoveryEndpoint
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				testserverpayload: `{"lagoon_version":"v2.17.0","authorization_endpoint":"https://keycloak.lagoon.sh","ssh_token_exchange":{"token_endpoint_host":"token.lagoon.sh","token_endpoint_port":22},"webhook_endpoint":"https://webhookhandler.lagoon.sh","ui_url":"https://ui.lagoon.sh"}`,
			},
			want: &DiscoveryEndpoint{
				LagoonVersion:         "v2.17.0",
				AuthorizationEndpoint: "https://keycloak.lagoon.sh",
				WebhookEndpoint:       "https://webhookhandler.lagoon.sh",
				UIHostname:            "https://ui.lagoon.sh",
				SSHTokenExchange: SSHTokenExchange{
					TokenPort: 22,
					TokenHost: "token.lagoon.sh",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create a test server with the wellknown path
			mux := http.NewServeMux()
			mux.HandleFunc(fmt.Sprintf("/%s", wellKnown), func(res http.ResponseWriter, req *http.Request) {
				// return the payload
				_, err := res.Write([]byte(tt.args.testserverpayload))
				if err != nil {
					t.Errorf("Discover() write error = %v", err)
					return
				}
			})
			// start the test server
			svr := httptest.NewServer(mux)
			defer svr.Close()
			// perform the request
			got, err := Discover(svr.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Discover() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// check the result is as expected
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Discover() = %v, want %v", got, tt.want)
			}
		})
	}
}
