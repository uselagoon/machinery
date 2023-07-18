package schema

type ParsedType interface {
	validateType() error
}

// ValidateType validates that the given type is valid. EnvironmentType & DeployType Validation currently supported
func ValidateType(p ParsedType) error {
	return p.validateType()
}
