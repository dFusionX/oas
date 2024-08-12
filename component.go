package oas

// RequestBody represents a request body object in OpenAPI
type RequestBody struct {
	Description string                `json:"description,omitempty" yaml:"description"`
	Content     map[string]*MediaType `json:"content" yaml:"content"`
	Required    bool                  `json:"required,omitempty" yaml:"required"`
	Ref         string                `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

// Components represent the component object in OpenAPI
type Components struct {
	Schemas         map[string]*Schema         `json:"schemas,omitempty" yaml:"schemas"`
	Responses       map[string]*Response       `json:"responses,omitempty" yaml:"responses"`
	Parameters      map[string]*Parameter      `json:"parameters,omitempty" yaml:"parameters"`
	Examples        map[string]*Example        `json:"examples,omitempty" yaml:"examples"`
	RequestBodies   map[string]*RequestBody    `json:"requestBodies,omitempty" yaml:"requestBodies"`
	Headers         map[string]*Header         `json:"headers,omitempty" yaml:"headers"`
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes,omitempty" yaml:"securitySchemes"`
	Links           map[string]*Link           `json:"links,omitempty" yaml:"links"`
	Callbacks       map[string]*Callback       `json:"callbacks,omitempty" yaml:"callbacks"`
}
