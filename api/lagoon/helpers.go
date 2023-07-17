package lagoon

import (
	"errors"
	"fmt"
	"strings"
)

// ErrExist indicates that an attempt was made to create an object that already
// exists.
var ErrExist = errors.New("object already exists")

// ParsedType validates that the given type is valid.
type ParsedType interface {
	validateType() error
}

type EnvType string
type DeployType string

const (
	ProductionEnv  EnvType = "PRODUCTION"
	DevelopmentEnv EnvType = "DEVELOPMENT"

	Branch  DeployType = "BRANCH"
	PR      DeployType = "PR"
	Promote DeployType = "PROMOTE"
)

func (e EnvType) validateType() error {
	envTypes := map[EnvType]struct{}{
		ProductionEnv:  {},
		DevelopmentEnv: {},
	}

	envT := strings.ToUpper(string(e))
	envType := EnvType(envT)
	_, ok := envTypes[envType]
	if !ok {
		return fmt.Errorf(`cannot parse:[%s] as EnvType`, e)
	}
	return nil
}

func (d DeployType) validateType() error {
	deployTypes := map[DeployType]struct{}{
		Branch:  {},
		Promote: {},
		PR:      {},
	}

	depT := strings.ToUpper(string(d))
	depType := DeployType(depT)
	_, ok := deployTypes[depType]
	if !ok {
		return fmt.Errorf(`cannot parse:[%s] as DeployType`, d)
	}
	return nil
}

// ValidateType validates that the given type is valid.
func ValidateType(p ParsedType) error {
	return p.validateType()
}
