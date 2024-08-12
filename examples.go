package oas

// Example represents an example object in OpenAPI
type Example struct {
	Summary       string      `json:"summary,omitempty" yaml:"summary"`
	Description   string      `json:"description,omitempty" yaml:"description"`
	Value         interface{} `json:"value,omitempty" yaml:"value"`
	ExternalValue string      `json:"externalValue,omitempty" yaml:"externalValue"`
	Ref           string      `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
