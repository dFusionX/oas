package oas

// SecurityScheme represents a security scheme in OpenAPI
type SecurityScheme struct {
	Type             string      `json:"type" yaml:"type"`                                   // REQUIRED. The type of the security scheme.
	Description      string      `json:"description,omitempty" yaml:"description"`           // A short description for security scheme.
	Name             string      `json:"name,omitempty" yaml:"name"`                         // REQUIRED. The name of the header, query, cookie, or path parameter to be used.
	In               string      `json:"in,omitempty" yaml:"in"`                             // REQUIRED. The location of the API key.
	Scheme           string      `json:"scheme,omitempty" yaml:"scheme"`                     // REQUIRED. The name of the HTTP Authorization scheme to be used.
	BearerFormat     string      `json:"bearerFormat,omitempty" yaml:"bearerFormat"`         // A hint to the client to identify how the bearer token is formatted.
	Flows            *OAuthFlows `json:"flows,omitempty" yaml:"flows"`                       // REQUIRED. An object containing configuration information for the flow types supported.
	OpenIDConnectURL string      `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl"` // REQUIRED. OpenID Connect URL to discover OAuth2 configuration values.
	Ref              string      `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

// OAuthFlows represents OAuth flows in OpenAPI
type OAuthFlows struct {
	Implicit          *OAuthFlow `json:"implicit,omitempty" yaml:"implicit"`
	Password          *OAuthFlow `json:"password,omitempty" yaml:"password"`
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty" yaml:"clientCredentials"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty" yaml:"authorizationCode"`
}

// OAuthFlow represents a single OAuth flow in OpenAPI
type OAuthFlow struct {
	AuthorizationURL string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl"`
	TokenURL         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl"`
	RefreshURL       string            `json:"refreshUrl,omitempty" yaml:"refreshUrl"`
	Scopes           map[string]string `json:"scopes" yaml:"scopes"`
}
