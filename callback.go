package oas

// Callback represents a callback object in OpenAPI
type Callback struct {
	Expression map[string]*PathItem `json:"expression" yaml:"expression"` // A map of expressions to PathItems, representing the callback endpoints.
	Ref        string               `json:"$ref,omitempty" yaml:"$ref"`   // Allows for an external definition of this callback object.
}

// PathItem represents a path item object in OpenAPI
type PathItem struct {
	Ref         string       `json:"$ref,omitempty" yaml:"$ref"`               // Allows for an external definition of this path item.
	Summary     string       `json:"summary,omitempty" yaml:"summary"`         // An optional, string summary, intended to apply to all operations in this path.
	Description string       `json:"description,omitempty" yaml:"description"` // An optional, string description, intended to apply to all operations in this path.
	Get         *Operation   `json:"get,omitempty" yaml:"get"`                 // A definition of a GET operation on this path.
	Put         *Operation   `json:"put,omitempty" yaml:"put"`                 // A definition of a PUT operation on this path.
	Post        *Operation   `json:"post,omitempty" yaml:"post"`               // A definition of a POST operation on this path.
	Delete      *Operation   `json:"delete,omitempty" yaml:"delete"`           // A definition of a DELETE operation on this path.
	Options     *Operation   `json:"options,omitempty" yaml:"options"`         // A definition of an OPTIONS operation on this path.
	Head        *Operation   `json:"head,omitempty" yaml:"head"`               // A definition of a HEAD operation on this path.
	Patch       *Operation   `json:"patch,omitempty" yaml:"patch"`             // A definition of a PATCH operation on this path.
	Trace       *Operation   `json:"trace,omitempty" yaml:"trace"`             // A definition of a TRACE operation on this path.
	Servers     []*Server    `json:"servers,omitempty" yaml:"servers"`         // An alternative server array to service all operations in this path.
	Parameters  []*Parameter `json:"parameters,omitempty" yaml:"parameters"`   // A list of parameters that are applicable for all the operations described under this path.
}

// Operation represents an operation object in OpenAPI
type Operation struct {
	Tags         []string               `json:"tags,omitempty" yaml:"tags"`                 // A list of tags for API documentation control.
	Summary      string                 `json:"summary,omitempty" yaml:"summary"`           // A short summary of what the operation does.
	Description  string                 `json:"description,omitempty" yaml:"description"`   // A verbose explanation of the operation behavior.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs"` // Additional external documentation for this operation.
	OperationID  string                 `json:"operationId,omitempty" yaml:"operationId"`   // Unique string used to identify the operation.
	Parameters   []*Parameter           `json:"parameters,omitempty" yaml:"parameters"`     // A list of parameters applicable for this operation.
	RequestBody  *RequestBody           `json:"requestBody,omitempty" yaml:"requestBody"`   // The request body applicable for this operation.
	Responses    map[string]*Response   `json:"responses" yaml:"responses"`                 // REQUIRED. The list of possible responses as they are returned from executing this operation.
	Callbacks    map[string]*Callback   `json:"callbacks,omitempty" yaml:"callbacks"`       // A map of possible out-of band callbacks related to the parent operation.
	Deprecated   bool                   `json:"deprecated,omitempty" yaml:"deprecated"`     // Declares this operation to be deprecated.
	Security     []*SecurityRequirement `json:"security,omitempty" yaml:"security"`         // A declaration of which security mechanisms can be used for this operation.
	Servers      []*Server              `json:"servers,omitempty" yaml:"servers"`           // An alternative server array to service this operation.
}

// ExternalDocumentation represents an external documentation object in OpenAPI
type ExternalDocumentation struct {
	Description string `json:"description,omitempty" yaml:"description"`
	URL         string `json:"url" yaml:"url"`
}

// SecurityRequirement represents a security requirement object in OpenAPI
type SecurityRequirement map[string][]string
