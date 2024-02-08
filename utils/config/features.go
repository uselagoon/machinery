package config

import (
	"fmt"
)

func featurePrefix(prefix, feature string) string {
	return fmt.Sprintf("%s-%s", prefix, feature)
}

// SetGlobalFeature updates a feature within the global config
func (c *Config) SetGlobalFeature(prefix, feature string, state bool) error {
	if prefix == "" || feature == "" {
		return fmt.Errorf("prefix or feature must not be empty")
	}
	if _, ok := c.Features[featurePrefix(prefix, feature)]; !ok {
		c.Features = map[string]bool{}
	}
	c.Features[featurePrefix(prefix, feature)] = state
	return nil
}

// SetContextFeature updates a feature within a contexts config
func (c *Config) SetContextFeature(context, prefix, feature string, state bool) error {
	if prefix == "" || feature == "" {
		return fmt.Errorf("prefix or feature must not be empty")
	}
	if i, ok := c.checkContextExists(context); ok {
		if _, ok := c.Contexts[i].ContextConfig.Features[featurePrefix(prefix, feature)]; !ok {
			c.Contexts[i].ContextConfig.Features = map[string]bool{}
		}
		c.Contexts[i].ContextConfig.Features[featurePrefix(prefix, feature)] = state
		return nil
	} else {
		return fmt.Errorf("context %s doesn't exist", context)
	}
}

func (c *Config) GetFeature(context, prefix, feature string) (bool, error) {
	if prefix == "" || feature == "" {
		return false, fmt.Errorf("prefix or feature must not be empty")
	}
	if i, ok := c.checkContextExists(context); ok {
		if state, ok := c.Contexts[i].ContextConfig.Features[featurePrefix(prefix, feature)]; ok {
			return state, nil
		} else {
			if state, ok := c.Features[featurePrefix(prefix, feature)]; ok {
				return state, nil
			}
			return false, nil
		}
	} else {
		return false, fmt.Errorf("context %s doesn't exist", context)
	}
}
