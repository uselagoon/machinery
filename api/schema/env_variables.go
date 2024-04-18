package schema

// EnvVariableType .
type EnvVariableType string

// . .
const (
	ProjectVar     EnvVariableType = "PROJECT"
	EnvironmentVar EnvVariableType = "ENVIRONMENT"
)

// EnvVariableScope .
type EnvVariableScope string

// . .
const (
	BuildVar                     EnvVariableScope = "BUILD"
	RuntimeVar                   EnvVariableScope = "RUNTIME"
	GlobalVar                    EnvVariableScope = "GLOBAL"
	InternalContainerRegistryVar EnvVariableScope = "INTERNAL_CONTAINER_REGISTRY"
	ContainerRegistryVar         EnvVariableScope = "CONTAINER_REGISTRY"
)

// EnvKeyValue is the base type of Environment variable.
type EnvKeyValue struct {
	ID    uint             `json:"id,omitempty"`
	Scope EnvVariableScope `json:"scope"`
	Name  string           `json:"name"`
	Value string           `json:"value"`
}

// EnvVariableInput is based on the Lagoon API type.
type EnvVariableInput struct {
	EnvKeyValue
	Type   EnvVariableType `json:"type"`
	TypeID uint            `json:"typeId"`
}

// EnvKeyValueInput  is based on the Lagoon API type.
type EnvKeyValueInput struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type EnvVariableByProjectEnvironmentNameInput struct {
	Environment string `json:"environment,omitempty"`
	Project     string `json:"project"`
}
