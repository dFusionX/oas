package oas

type Contact struct {
	Email string `json:"email"`
}

type License struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// Info provides metadata about the API
type Info struct {
	Title          string   `json:"title" yaml:"title"`                             // REQUIRED. The title of the API.
	Description    string   `json:"description,omitempty" yaml:"description"`       // A short description of the API.
	TermsOfService string   `json:"termsOfService,omitempty" yaml:"termsOfService"` // A URL to the Terms of Service for the API.
	Contact        *Contact `json:"contact,omitempty" yaml:"contact"`               // The contact information for the exposed API.
	License        *License `json:"license,omitempty" yaml:"license"`               // The license information for the exposed API.
	Version        string   `json:"version" yaml:"version"`                         // REQUIRED. The version of the OpenAPI document.
}
