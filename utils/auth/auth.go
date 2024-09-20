package auth

import (
	"context"
	"embed"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

// store the html using embed
var (
	//go:embed html/response.html
	content embed.FS
)

const (
	defaultRealm = "lagoon"
	clientID     = "lagoon-cli"
	clientSecret = ""
	authPath     = "protocol/openid-connect/auth"
	tokenPath    = "protocol/openid-connect/token"
)

// TokenRequest will perform a request against keycloak
// it will prompt the user to authenticate and direct them to keycloak if the token is expired or cannot be refreshed
// otherwise it will perform a token refresh
func TokenRequest(keycloakURL, realm, idpHint string, token *oauth2.Token) error {
	ctx := context.Background()
	// create the oauth configuration from known paths
	if realm == "" {
		realm = defaultRealm
	}
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("%s/realms/%s/%s?kc_idp_hint=%s", keycloakURL, realm, authPath, idpHint),
			TokenURL: fmt.Sprintf("%s/realms/%s/%s", keycloakURL, realm, tokenPath),
		},
	}

	var src oauth2.TokenSource
	// if old token hasn't got a refresh token, its probably an old or invalid token
	if token.RefreshToken == "" {
		src = conf.TokenSource(ctx, nil)
	} else {
		// otherwise try with the provided full oauth token
		src = conf.TokenSource(ctx, token)
	}
	newToken, err := src.Token() // this actually goes and renews the token
	if err != nil {
		// if no token, or there is an error. start the local auth listener and direct the user to
		// keycloak to sign in and get a new token
		err := authListener(ctx, conf, token)
		if err != nil {
			return err
		}
	} else {
		if newToken.AccessToken != token.AccessToken {
			// set token to the new token
			*token = *newToken
		}
	}
	_ = conf.Client(ctx, token)
	return nil
}

func authListener(ctx context.Context, conf *oauth2.Config, token *oauth2.Token) error {
	// create a local listener on a random free port
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return fmt.Errorf("lagoon-auth: can't open listener on port: %s", err)
	}
	// create the redirect url with the random port assigned
	conf.RedirectURL = fmt.Sprintf("http://%s", l.Addr().String())

	// generate a verifier for this request
	verifier := oauth2.GenerateVerifier()

	// start a local server on the random local port
	server := &http.Server{Addr: l.Addr().String()}

	// define a handler that will get the authorization code, call the token endpoint, and close the HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// get the authorization code from the request
		code := r.URL.Query().Get("code")
		if code == "" {
			log.Fatalf("lagoon-auth: URL param 'code' is missing")
			io.WriteString(w, "Error: could not find 'code' URL parameter\n")
			go server.Close()
			return
		}

		// perform the token auth exchange with the auth endpoint
		newToken, err := conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
		if err != nil {
			log.Fatalf("lagoon-auth: unable to refresh token: %v", err)
			io.WriteString(w, "Error: unable to refresh token\n")
			go server.Close()
			return
		}
		// set token to the new token
		*token = *newToken

		// print the page to the user to say it has completed
		// retrieve the html using embed
		q, err := content.ReadFile("html/response.html")
		if err != nil {
			log.Fatalf("lagoon-auth: couldn't load response asset: %v", err)
			io.WriteString(w, "Error: unable to refresh token\n")
			go server.Close()
			return
		}
		io.WriteString(w, string(q))
		// print to stderr so redirects of stdout still work when the command exits with a successful token request
		fmt.Fprintln(os.Stderr, "Successfully logged in.")

		// close the server
		go server.Close()
	})

	// prompt user to visit URL to continue authentication
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOnline, oauth2.S256ChallengeOption(verifier))
	// print to stderr so redirects of stdout still work when the command exits with a successful token request
	fmt.Fprintln(os.Stderr, "\nLogin to Keycloak at", url)
	// this will exit when the handler gets fired and calls server.Close()
	if err := server.Serve(l); err != nil {
		return err
	}
	return nil
}
