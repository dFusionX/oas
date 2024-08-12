package oas

// Schema represents a schema object in OpenAPI
type Schema struct {
	Title                string                 `json:"title,omitempty" yaml:"title"`                               // The title of the schema.
	MultipleOf           *float64               `json:"multipleOf,omitempty" yaml:"multipleOf"`                     // Constrains the value to be a multiple of a given number.
	Maximum              *float64               `json:"maximum,omitempty" yaml:"maximum"`                           // Constrains the value to be at most a maximum.
	ExclusiveMaximum     bool                   `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum"`         // Whether `maximum` is exclusive.
	Minimum              *float64               `json:"minimum,omitempty" yaml:"minimum"`                           // Constrains the value to be at least a minimum.
	ExclusiveMinimum     bool                   `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum"`         // Whether `minimum` is exclusive.
	MaxLength            *int                   `json:"maxLength,omitempty" yaml:"maxLength"`                       // Constrains the length of a string.
	MinLength            *int                   `json:"minLength,omitempty" yaml:"minLength"`                       // Constrains the length of a string.
	Pattern              *string                `json:"pattern,omitempty" yaml:"pattern"`                           // A regular expression a string value must conform to.
	MaxItems             *int                   `json:"maxItems,omitempty" yaml:"maxItems"`                         // Constrains the number of items in an array.
	MinItems             *int                   `json:"minItems,omitempty" yaml:"minItems"`                         // Constrains the number of items in an array.
	UniqueItems          bool                   `json:"uniqueItems,omitempty" yaml:"uniqueItems"`                   // Ensures that items in an array are unique.
	MaxProperties        *int                   `json:"maxProperties,omitempty" yaml:"maxProperties"`               // Constrains the number of properties in an object.
	MinProperties        *int                   `json:"minProperties,omitempty" yaml:"minProperties"`               // Constrains the number of properties in an object.
	Required             []string               `json:"required,omitempty" yaml:"required"`                         // Lists the required properties.
	Enum                 []interface{}          `json:"enum,omitempty" yaml:"enum"`                                 // Specifies the allowed values.
	Type                 string                 `json:"type,omitempty" yaml:"type"`                                 // The type of the schema (e.g., string, integer, etc.).
	AllOf                []*Schema              `json:"allOf,omitempty" yaml:"allOf"`                               // Combines subschemas into a single schema.
	OneOf                []*Schema              `json:"oneOf,omitempty" yaml:"oneOf"`                               // Combines subschemas into a single schema, only one of which should validate.
	AnyOf                []*Schema              `json:"anyOf,omitempty" yaml:"anyOf"`                               // Combines subschemas into a single schema, any of which can validate.
	Not                  *Schema                `json:"not,omitempty" yaml:"not"`                                   // Inverts a schema, ensuring it does not validate.
	Items                *Schema                `json:"items,omitempty" yaml:"items"`                               // Defines the type of items in an array schema.
	Properties           map[string]*Schema     `json:"properties,omitempty" yaml:"properties"`                     // Defines the properties for an object schema.
	AdditionalProperties *Schema                `json:"additionalProperties,omitempty" yaml:"additionalProperties"` // Allows for additional properties in an object schema.
	Description          string                 `json:"description,omitempty" yaml:"description"`                   // A description of the schema.
	Format               string                 `json:"format,omitempty" yaml:"format"`                             // Provides additional data type information.
	Default              interface{}            `json:"default,omitempty" yaml:"default"`                           // The default value for the schema.
	Nullable             bool                   `json:"nullable,omitempty" yaml:"nullable"`                         // Allows the schema to be null.
	ReadOnly             bool                   `json:"readOnly,omitempty" yaml:"readOnly"`                         // Marks the schema as read-only.
	WriteOnly            bool                   `json:"writeOnly,omitempty" yaml:"writeOnly"`                       // Marks the schema as write-only.
	XML                  *XML                   `json:"xml,omitempty" yaml:"xml"`                                   // XML modeling information.
	ExternalDocs         *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs"`                 // Additional external documentation.
	Example              interface{}            `json:"example,omitempty" yaml:"example"`                           // An example of the schema's potential value.
	Deprecated           bool                   `json:"deprecated,omitempty" yaml:"deprecated"`                     // Marks the schema as deprecated.
	Ref                  string                 `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

// XML represents XML modeling information for an object in OpenAPI
type XML struct {
	Name      string `json:"name,omitempty" yaml:"name"`           // Replaces the name of the element/attribute used for the described schema property.
	Namespace string `json:"namespace,omitempty" yaml:"namespace"` // The URL of the namespace definition.
	Prefix    string `json:"prefix,omitempty" yaml:"prefix"`       // The prefix to be used for the name.
	Attribute bool   `json:"attribute,omitempty" yaml:"attribute"` // Declares whether the property definition translates to an attribute instead of an element.
	Wrapped   bool   `json:"wrapped,omitempty" yaml:"wrapped"`     // Applies to an array schema; adds a wrapping element.
}
