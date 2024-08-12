package oas

// Path represents a path object in OpenAPI
type Path struct {
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
