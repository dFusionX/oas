package oas

// Response represents a response object in OpenAPI
type Response struct {
	Description string                `json:"description" yaml:"description"`   // REQUIRED. A short description of the response.
	Headers     map[string]*Header    `json:"headers,omitempty" yaml:"headers"` // Maps a header name to its definition.
	Content     map[string]*MediaType `json:"content,omitempty" yaml:"content"` // A map containing descriptions of potential response payloads.
	Links       map[string]*Link      `json:"links,omitempty" yaml:"links"`     // A map of operations links that can be followed from the response.
	Ref         string                `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

// Header represents a header object in OpenAPI
type Header struct {
	Description     string              `json:"description,omitempty" yaml:"description"`
	Required        bool                `json:"required,omitempty" yaml:"required"`
	Deprecated      bool                `json:"deprecated,omitempty" yaml:"deprecated"`
	AllowEmptyValue bool                `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue"`
	Schema          *Schema             `json:"schema,omitempty" yaml:"schema"`
	Example         interface{}         `json:"example,omitempty" yaml:"example"`
	Examples        map[string]*Example `json:"examples,omitempty" yaml:"examples"`
	Ref             string              `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

// MediaType represents a media type object in OpenAPI
type MediaType struct {
	Schema   *Schema              `json:"schema,omitempty" yaml:"schema"`
	Example  interface{}          `json:"example,omitempty" yaml:"example"`
	Examples map[string]*Example  `json:"examples,omitempty" yaml:"examples"`
	Encoding map[string]*Encoding `json:"encoding,omitempty" yaml:"encoding"`
}

// Link represents a link object in OpenAPI
type Link struct {
	OperationRef string                 `json:"operationRef,omitempty" yaml:"operationRef"`
	OperationID  string                 `json:"operationId,omitempty" yaml:"operationId"`
	Parameters   map[string]interface{} `json:"parameters,omitempty" yaml:"parameters"`
	RequestBody  interface{}            `json:"requestBody,omitempty" yaml:"requestBody"`
	Description  string                 `json:"description,omitempty" yaml:"description"`
	Server       *Server                `json:"server,omitempty" yaml:"server"`
	Ref          string                 `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

// Encoding represents an encoding object in OpenAPI
type Encoding struct {
	ContentType   string                 `json:"contentType,omitempty" yaml:"contentType"`
	Headers       map[string]*Header     `json:"headers,omitempty" yaml:"headers"`
	Style         string                 `json:"style,omitempty" yaml:"style"`
	Explode       bool                   `json:"explode,omitempty" yaml:"explode"`
	AllowReserved bool                   `json:"allowReserved,omitempty" yaml:"allowReserved"`
	Extensions    map[string]interface{} `json:"-" yaml:"-"`
}

// ServerVariable represents a server variable object in OpenAPI
type ServerVariable struct {
	Enum        []string               `json:"enum,omitempty" yaml:"enum"`
	Default     string                 `json:"default" yaml:"default"`
	Description string                 `json:"description,omitempty" yaml:"description"`
	Extensions  map[string]interface{} `json:"-" yaml:"-"`
}
