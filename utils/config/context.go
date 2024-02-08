package config

import (
	"fmt"
)

// GetContext will get a context from the config by name and return it
func (c *Config) GetContext(context string) (*Context, error) {
	if context == "" {
		return nil, fmt.Errorf("context name not provided")
	}
	if i, ok := c.checkContextExists(context); ok {
		return &c.Contexts[i], nil
	}
	return nil, fmt.Errorf("context %s doesn't exist", context)
}

// GetDefaultContext will return whichever context is defined as the default in the config
func (c *Config) GetDefaultContext() (*Context, error) {
	if i, ok := c.checkContextExists(c.DefaultContext); ok {
		return &c.Contexts[i], nil
	}
	return nil, fmt.Errorf("no context found")
}

// SetDefaultContext will set the context by name as the default context in the config
func (c *Config) SetDefaultContext(context string) error {
	if context == "" {
		return fmt.Errorf("context name not provided")
	}
	if _, ok := c.checkContextExists(context); ok {
		c.DefaultContext = context
		return nil
	}
	return fmt.Errorf("context %s doesn't exist", context)
}

// NewContext will add a context to the config if it doesn't exist
func (c *Config) NewContext(context ContextConfig, user string) error {
	if context.Name == "" {
		return fmt.Errorf("context name not provided")
	}
	if context.APIHostname == "" {
		return fmt.Errorf("context api hostname not provided")
	}
	if context.TokenHost == "" {
		return fmt.Errorf("context token hostname not provided")
	}
	if context.TokenPort == 0 {
		return fmt.Errorf("context token port not provided")
	}
	if context.AuthenticationEndpoint == "" {
		return fmt.Errorf("context authentication hostname not provided")
	}
	if user == "" {
		return fmt.Errorf("user name not provided")
	}
	if _, ok := c.checkUserExists(user); ok {
		if _, ok := c.checkContextExists(context.Name); !ok {
			c.Contexts = append(c.Contexts, Context{
				Name:          context.Name,
				ContextConfig: context,
				User:          user,
			})
			return nil
		} else {
			return fmt.Errorf("context %s already exists", context.Name)
		}
	} else {
		return fmt.Errorf("user %s not found, add user first", user)
	}
}

// UpdateContext will update a context in the config if it exists
func (c *Config) UpdateContext(context ContextConfig, user string) error {
	if _, ok := c.checkUserExists(user); ok {
		if i, ok := c.checkContextExists(context.Name); ok {
			c.Contexts[i] = Context{
				Name:          context.Name,
				ContextConfig: context,
				User:          user,
			}
			return nil
		} else {
			return fmt.Errorf("context %s doesn't exist", context.Name)
		}
	} else {
		return fmt.Errorf("user %s not found, add user first", user)
	}
}

// DeleteContext will remove a context from the config
func (c *Config) DeleteContext(context string) error {
	if i, ok := c.checkContextExists(context); ok {
		c.Contexts[i] = c.Contexts[len(c.Contexts)-1]
		c.Contexts = append(c.Contexts[:i], c.Contexts[i+1:]...)
		return nil
	}
	return fmt.Errorf("context %s doesn't exist", context)
}
