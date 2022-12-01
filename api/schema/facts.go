package schema

type Fact struct {
	Id          int      `json:"id,omitempty"`
	Environment int      `json:"environment,omitempty"`
	Name        string   `json:"name,omitempty"`
	Value       string   `json:"value,omitempty"`
	Source      string   `json:"source,omitempty"`
	Description string   `json:"description,omitempty"`
	KeyFact     bool     `json:"keyFact,omitempty"`
	Type        FactType `json:"type,omitempty"`
	Category    string   `json:"category,omitempty"`
}

type Facts []Fact

type AddFactInput struct {
	Id          int      `json:"id,omitempty"`
	Environment int      `json:"environment,omitempty"`
	Name        string   `json:"name,omitempty"`
	Value       string   `json:"value,omitempty"`
	Source      string   `json:"source,omitempty"`
	Description string   `json:"description,omitempty"`
	KeyFact     bool     `json:"keyFact,omitempty"`
	Type        FactType `json:"type,omitempty"`
	Category    string   `json:"category,omitempty"`
}

type FactType string

const (
	FactTypeText   FactType = "TEXT"
	FactTypeUrl    FactType = "URL"
	FactTypeSemver FactType = "SEMVER"
)
