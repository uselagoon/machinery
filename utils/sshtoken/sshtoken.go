package sshtoken

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	terminal "golang.org/x/term"
)

func GetPublicKeyAuthMethod(path string, publicKeyOverride *string, publicKeyIdentities []string, skipAgent, verboseOutput bool) (ssh.AuthMethod, func() error) {
	noopCloseFunc := func() error { return nil }

	if !skipAgent {
		// Connect to SSH agent to ask for unencrypted private keys
		if sshAgentConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
			sshAgent := agent.NewClient(sshAgentConn)
			agentSigners, err := sshAgent.Signers()
			if err != nil {
				return nil, func() error {
					return err
				}
			}
			// There are key(s) in the agent
			if len(agentSigners) > 0 {
				identities := make(map[string]ssh.PublicKey)
				if publicKeyOverride == nil {
					// check for identify files in the current lagoon config context
					for _, identityFile := range publicKeyIdentities {
						// append to identityfiles
						keybytes, err := os.ReadFile(identityFile)
						if err != nil {
							// if identity file doesn't exist, skip it and use all agent keys
							if verboseOutput {
								fmt.Fprintf(os.Stderr, "ssh: identity file %s not found, will use all keys from agent\n", identityFile)
							}
							continue
						}
						pubkey, _, _, _, err := ssh.ParseAuthorizedKey(keybytes)
						if err != nil {
							if verboseOutput {
								fmt.Fprintf(os.Stderr, "ssh: failed to parse identity file %s, will use all keys from agent\n", identityFile)
							}
							continue
						}
						identities[identityFile] = pubkey
					}
				} else {
					// append to identityfiles
					keybytes, err := os.ReadFile(*publicKeyOverride)
					if err != nil {
						// if override file doesn't exist but agent has keys, use all agent keys
						if verboseOutput {
							fmt.Fprintf(os.Stderr, "ssh: identity file %s not found, will use all keys from agent\n", *publicKeyOverride)
						}
					} else {
						pubkey, _, _, _, parseErr := ssh.ParseAuthorizedKey(keybytes)
						if parseErr != nil {
							if verboseOutput {
								fmt.Fprintf(os.Stderr, "ssh: failed to parse identity file %s, will use all keys from agent\n", *publicKeyOverride)
							}
						} else {
							identities[*publicKeyOverride] = pubkey
						}
					}
				}
				// check all keys in the agent to see if there is a matching identity file
				for _, signer := range agentSigners {
					for file, identity := range identities {
						if bytes.Equal(signer.PublicKey().Marshal(), identity.Marshal()) {
							if verboseOutput {
								fmt.Fprintf(os.Stderr, "ssh: attempting connection using identity file public key: %s\n", file)
							}
							// only provide this matching key back to the ssh client to use
							return ssh.PublicKeys(signer), noopCloseFunc
						}
					}
				}
				// If publicKeyOverride was specified but no identities were loaded successfully,
				// we still fall through to use all agent keys instead of failing
				// if no matching identity files, just return all agent keys like previous behaviour
				if verboseOutput {
					fmt.Fprintf(os.Stderr, "ssh: attempting connection using any keys in ssh-agent\n")
				}
				return ssh.PublicKeysCallback(sshAgent.Signers), noopCloseFunc
			}
		}
	}

	// if no keys in the agent, and a specific private key has been defined, then check the key and use it if possible
	if verboseOutput {
		fmt.Fprintf(os.Stderr, "ssh: attempting connection using private key: %s\n", path)
	}
	key, err := os.ReadFile(path)
	if err != nil {
		// provide helpful error message when no keys available
		if os.IsNotExist(err) {
			return nil, func() error {
				return fmt.Errorf("ssh: no keys found in ssh-agent and private key file %s does not exist. Please add keys to your ssh-agent with 'ssh-add', specify a key file path in config, or use --ssh-key flag", path)
			}
		}
		return nil, func() error {
			return err
		}
	}

	// Try to look for an unencrypted private key
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		// if encrypted, prompt for passphrase or error and ask user to add to their agent
		fmt.Printf("Enter passphrase for %s:", path)
		bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return nil, func() error {
				return fmt.Errorf("cannot decode private keys, you will need to add your private key to your ssh-agent: %v", err)
			}
		}
		fmt.Println()
		signer, err = ssh.ParsePrivateKeyWithPassphrase(key, bytePassword)
		if err != nil {
			return nil, func() error {
				return fmt.Errorf("cannot decode private keys, you will need to add your private key to your ssh-agent: %v", err)
			}
		}
	}
	return ssh.PublicKeys(signer), noopCloseFunc
}

// retrieve a token from the ssh token service of Lagoon
func RetrieveToken(sshKey, sshHost, sshPort string, publicKeyOverride *string, publicKeyIdentities []string, verbose bool) (string, error) {
	skipAgent := false
	authMethod, closeSSHAgent := GetPublicKeyAuthMethod(sshKey, publicKeyOverride, publicKeyIdentities, skipAgent, verbose)
	if authMethod == nil {
		return "", errors.New("couldn't load ssh key")
	}
	config := &ssh.ClientConfig{
		User: "lagoon",
		Auth: []ssh.AuthMethod{
			authMethod,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	defer closeSSHAgent() //nolint:errcheck

	sshHostString := fmt.Sprintf("%s:%s",
		sshHost,
		sshPort,
	)
	conn, err := ssh.Dial("tcp", sshHostString, config)
	if err != nil {
		return "", fmt.Errorf("couldn't connect to %s: %v", sshHostString, err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return "", fmt.Errorf("couldn't open session: %v", err)
	}

	out, err := session.CombinedOutput("token")
	if err != nil {
		return "", fmt.Errorf("couldn't get token: %v", err)
	}
	return strings.TrimSpace(string(out)), err
}

// ValidateOrRefreshToken validates if token is valid or not, and will attempt to refresh
func ValidateOrRefreshToken(sshKey, sshHost, sshPort string, publicKeyOverride *string, publicKeyIdentities []string, token *string, verbose bool) error {
	var err error
	if !tokenValidOrExpired(*token) {
		// token is valid, continue
		return nil
	}
	// token not valide, retrieve a new oken
	*token, err = RetrieveToken(sshKey, sshHost, sshPort, publicKeyOverride, publicKeyIdentities, verbose)
	if err != nil {
		return fmt.Errorf("couldn't refresh token: %w", err)
	}
	// token refreshed
	return nil
}

// check the token to see if it has expired/is valid
func tokenValidOrExpired(token string) bool {
	var p jwt.Parser
	t, _, err := p.ParseUnverified(
		token, &jwt.StandardClaims{})
	if err != nil {
		return true
	}
	if t.Claims.Valid() != nil {
		return true
	}
	return false
}
