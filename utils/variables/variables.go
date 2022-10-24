package variables

import (
	"os"
	"strconv"
	"strings"
)

// LagoonEnvironmentVariable is used to define Lagoon environment variables.
type LagoonEnvironmentVariable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Scope string `json:"scope"`
}

// ReplaceOrAddVariable will replace or add an environment variable to a slice of environment variables
func ReplaceOrAddVariable(vars *[]LagoonEnvironmentVariable, name, value, scope string) {
	exists := false
	existsIdx := 0
	for idx, v := range *vars {
		if v.Name == name {
			exists = true
			existsIdx = idx
		}
	}
	if exists {
		(*vars)[existsIdx].Value = value
	} else {
		(*vars) = append((*vars), LagoonEnvironmentVariable{
			Name:  name,
			Value: value,
			Scope: scope})
	}
}

// VariableExists checks if a variable exists in a slice of environment variables
func VariableExists(vars *[]LagoonEnvironmentVariable, name, value string) bool {
	exists := false
	for _, v := range *vars {
		if v.Name == name && v.Value == value {
			exists = true
		}
	}
	return exists
}

// GetLagoonFeatureFlags pull all the environment variables and check if there are any with the flag prefix
func GetLagoonFeatureFlags() map[string]string {
	flags := make(map[string]string)
	for _, envVar := range os.Environ() {
		envVarSplit := strings.SplitN(envVar, "=", 2)
		// if the variable name contains the flag prefix, add it to the map
		if strings.Contains(envVarSplit[0], "LAGOON_FEATURE_FLAG_") {
			flags[envVarSplit[0]] = envVarSplit[1]
		}
	}
	return flags
}

// GetEnv get environment variable
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetEnvInt get environment variable as int
func GetEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		valueInt, e := strconv.Atoi(value)
		if e == nil {
			return valueInt
		}
	}
	return fallback
}

// GetEnvBool get environment variable as bool
// accepts fallback values 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False
// anything else is false.
func GetEnvBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		rVal, _ := strconv.ParseBool(value)
		return rVal
	}
	return fallback
}
