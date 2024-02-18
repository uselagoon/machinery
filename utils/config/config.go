package config

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
)

const (
	// this is the directory within the XDG path that configuration files will be stored
	configDir = "lagoon"

	// this is the name of the base lagoon config file
	configFile = "config.yaml"
)

// Config is the basic config structure, it is where contexts, users, and the default context are stored.
type Config struct {
	Contexts       []Context       `yaml:"contexts"`
	Users          []User          `yaml:"users"`
	DefaultContext string          `yaml:"default-context"`
	Features       map[string]bool `yaml:"features,omitempty"`
}

// Context is the where the context configuration and user reference are joined
type Context struct {
	Name          string        `yaml:"name"`
	ContextConfig ContextConfig `yaml:"context"`
	User          string        `yaml:"user"`
}

// ContextConfig contains information about the context itself
// this is where configuration overrides should be added if there is a reason to extend in the future
type ContextConfig struct {
	Name                   string          `yaml:"name"`
	APIHostname            string          `yaml:"apihostname"`
	AuthenticationEndpoint string          `yaml:"authenticationendpoint"`
	TokenHost              string          `yaml:"tokenhost"`
	TokenPort              int             `yaml:"tokenport"`
	UIHostname             string          `yaml:"uihostname,omitempty"`
	WebhookEndpoint        string          `yaml:"webhookendpoint,omitempty"`
	Version                string          `yaml:"version,omitempty"`
	Features               map[string]bool `yaml:"features,omitempty"`
}

// User is the where the user configuration is linked to the name
type User struct {
	Name       string     `yaml:"name"`
	UserConfig UserConfig `yaml:"user"`
}

// UserConfig contains any user information. In the future, if there are reasons to extend or override configurations for users
// this is where those overrides would be configured
type UserConfig struct {
	Name string `yaml:"name"`
	// sshkey is optional
	// if defined, then a client should use the `tokenhost` method of authentication
	// otherwise, client will use `authenticationendpoint` authentication method
	SSHKey string        `yaml:"sshkey"`
	Grant  *oauth2.Token `yaml:"grant"`
}

// LoadConfig will load the configuration if it exists
// if create is defined as true, it will attempt to create the file, otherwise it will error that the configuration does not exist
func LoadConfig(create bool) (*Config, error) {
	c := &Config{}
	if !create {
		_, err := xdg.SearchConfigFile(fmt.Sprintf("%s/%s", configDir, configFile))
		if err != nil {
			// if the config can't be created, the read will fail
			return c, err
		}
	}
	configFilePath, err := xdg.ConfigFile(fmt.Sprintf("%s/%s", configDir, configFile))
	if err != nil {
		return c, err
	}
	err = c.readConfig(configFilePath)
	if err != nil {
		return c, err
	}
	return c, nil
}

// GetConfigLocation will return the config file location using XDG
func (c *Config) GetConfigLocation() (string, error) {
	configFilePath, err := xdg.ConfigFile(fmt.Sprintf("%s/%s", configDir, configFile))
	if err != nil {
		return "", err
	}
	return configFilePath, nil
}

// ReadConfig will read the config file using XDG into the config
func (c *Config) ReadConfig() error {
	configFilePath, err := xdg.ConfigFile(fmt.Sprintf("%s/%s", configDir, configFile))
	if err != nil {
		return err
	}
	return c.readConfig(configFilePath)
}

// WriteConfig will write the config into the config file using XDG
func (c *Config) WriteConfig() error {
	configFilePath, err := xdg.ConfigFile(fmt.Sprintf("%s/%s", configDir, configFile))
	if err != nil {
		return err
	}
	return c.writeConfig(configFilePath)
}

// GetUserContext will return a user configuration and context configuration from the requested context name from the config
func (c *Config) GetUserContext(context string) (*UserConfig, *ContextConfig, error) {
	if i, ok := c.checkContextExists(context); ok {
		if i, ok := c.checkUserExists(c.Contexts[i].User); ok {
			return &c.Users[i].UserConfig, &c.Contexts[i].ContextConfig, nil
		}
		return nil, nil, fmt.Errorf("user %s doesn't exist for context %s", c.Contexts[i].User, context)
	}
	return nil, nil, fmt.Errorf("context %s doesn't exist", context)
}

// helper for reading the config file from a defined path
func (c *Config) readConfig(file string) error {
	rawYAML, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("couldn't read %v: %v", file, err)
	}
	err = yaml.Unmarshal(rawYAML, c)
	if err != nil {
		return err
	}
	return nil
}

// helper for writing the config file from a defined path
func (c *Config) writeConfig(file string) error {
	yamlB, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(file, yamlB, 0644)
}

// helper for checking if a user exists within the config
func (c *Config) checkUserExists(user string) (int, bool) {
	for idx, u := range c.Users {
		if u.Name == user {
			return idx, true
		}
	}
	return 0, false
}

// helper for checking if a context exists within the config
func (c *Config) checkContextExists(context string) (int, bool) {
	for idx, con := range c.Contexts {
		if con.Name == context {
			return idx, true
		}
	}
	return 0, false
}
