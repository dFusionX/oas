package oas

// Parameter represents a parameter object in OpenAPI
type Parameter struct {
	Name            string                `json:"name" yaml:"name"`                                 // REQUIRED. The name of the parameter.
	In              string                `json:"in" yaml:"in"`                                     // REQUIRED. The location of the parameter.
	Description     string                `json:"description,omitempty" yaml:"description"`         // A brief description of the parameter.
	Required        bool                  `json:"required,omitempty" yaml:"required"`               // Determines whether this parameter is mandatory.
	Deprecated      bool                  `json:"deprecated,omitempty" yaml:"deprecated"`           // Specifies that a parameter is deprecated and should be transitioned out of usage.
	AllowEmptyValue bool                  `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue"` // Sets the ability to pass empty-valued parameters.
	Style           string                `json:"style,omitempty" yaml:"style"`                     // Describes how the parameter value will be serialized.
	Explode         bool                  `json:"explode,omitempty" yaml:"explode"`                 // When this is true, parameter values of type array or object generate separate parameters for each value of the array or key-value pair of the map.
	AllowReserved   bool                  `json:"allowReserved,omitempty" yaml:"allowReserved"`     // Determines whether the parameter value should allow reserved characters.
	Schema          *Schema               `json:"schema,omitempty" yaml:"schema"`                   // The schema defining the type used for the parameter.
	Example         interface{}           `json:"example,omitempty" yaml:"example"`                 // Example of the parameter's potential value.
	Examples        map[string]*Example   `json:"examples,omitempty" yaml:"examples"`               // Examples of the parameter's potential value.
	Content         map[string]*MediaType `json:"content,omitempty" yaml:"content"`                 // A map containing the representations for the parameter.
	Ref             string                `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
