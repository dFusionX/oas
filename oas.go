package oas

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// Server represents a server object in OpenAPI
type Server struct {
	URL         string                     `json:"url" yaml:"url"`
	Description string                     `json:"description,omitempty" yaml:"description"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty" yaml:"variables"`
}

// Tag adds metadata to a single tag that is used by the Operation object
type Tag struct {
	Name         string                 `json:"name" yaml:"name"`                           // REQUIRED. The name of the tag.
	Description  string                 `json:"description,omitempty" yaml:"description"`   // A short description for the tag.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs"` // Additional external documentation for this tag.
}

// OpenAPI represents the root document object of the OpenAPI specification
type OpenAPI struct {
	OpenAPIVersion string                 `json:"openapi" yaml:"openapi"`                     // REQUIRED. The semantic version number of the OpenAPI specification.
	Info           *Info                  `json:"info" yaml:"info"`                           // REQUIRED. Provides metadata about the API.
	ExternalDocs   *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs"` // Additional external documentation.
	Servers        []*Server              `json:"servers,omitempty" yaml:"servers"`           // An array of Server objects, which provide connectivity information to a target server.
	Tags           []*Tag                 `json:"tags,omitempty" yaml:"tags"`                 // A list of tags used by the specification with additional metadata.
	Paths          map[string]*Path       `json:"paths" yaml:"paths"`                         // REQUIRED. The available paths and operations for the API.
	Components     *Components            `json:"components,omitempty" yaml:"components"`     // An element to hold various schemas for the specification.
}

func NewOpenAPI(bytes []byte) (*OpenAPI, error) {
	openAPI := &OpenAPI{}
	if err := json.Unmarshal(bytes, openAPI); err != nil {
		if err := yaml.Unmarshal(bytes, openAPI); err != nil {
			return nil, err
		}
	}
	return openAPI, nil
}
