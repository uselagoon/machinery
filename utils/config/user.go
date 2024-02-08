package config

import (
	"fmt"

	"golang.org/x/oauth2"
)

// GetUser will return a user from the config by name
func (c *Config) GetUser(user string) (*User, error) {
	if i, ok := c.checkUserExists(user); ok {
		return &c.Users[i], nil
	}
	return nil, fmt.Errorf("no user found")
}

// NewUser will add a new user to the config
func (c *Config) NewUser(user UserConfig) error {
	if user.Name == "" {
		return fmt.Errorf("user name not provided")
	}
	if user.Grant == nil {
		user.Grant = &oauth2.Token{}
	}
	if _, ok := c.checkUserExists(user.Name); !ok {
		c.Users = append(c.Users, User{
			Name:       user.Name,
			UserConfig: user,
		})
		return nil
	} else {
		return fmt.Errorf("user %s already exists", user.Name)
	}
}

// UpdateUser updates a user within the config
func (c *Config) UpdateUser(user UserConfig) error {
	if i, ok := c.checkUserExists(user.Name); ok {
		c.Users[i] = User{
			Name:       user.Name,
			UserConfig: user,
		}
		return nil
	} else {
		return fmt.Errorf("user %s doesn't exist", user.Name)
	}
}

// DeleteUser removes a user from the config
func (c *Config) DeleteUser(user string) error {
	if i, ok := c.checkUserExists(user); ok {
		for _, con := range c.Contexts {
			if con.User == user {
				return fmt.Errorf("user %s attached to context %s, change the user for this context or remove the context first", user, con.Name)
			}
		}
		c.Users[i] = c.Users[len(c.Users)-1]
		c.Users = append(c.Users[:i], c.Users[i+1:]...)
		return nil
	}
	return fmt.Errorf("user %s doesn't exist", user)
}
