package oas

import (
	"fmt"
	"regexp"
	"strings"
)

// Dereference replaces all $ref fields in the OpenAPI struct with the actual referenced objects.
func (o OpenAPI) Dereference() (*OpenAPI, error) {
	// Create a deep copy of the OpenAPI struct to avoid modifying the original
	dereferenced := o

	// Dereference Paths
	for pathKey, pathItem := range dereferenced.Paths {
		if pathItem.Ref != "" {
			refPath, err := resolveRefPath(pathItem.Ref, dereferenced.Components)
			if err != nil {
				return nil, err
			}
			dereferenced.Paths[pathKey] = refPath
		} else {
			dereferenced.Paths[pathKey] = pathItem.dereference(dereferenced.Components)
		}
	}

	// Dereference Components (schemas, responses, parameters, etc.)
	if dereferenced.Components != nil {
		dereferenced.Components = dereferenced.Components.dereference()
	}

	return &dereferenced, nil
}

// Dereference method for Path struct
func (p *Path) dereference(components *Components) *Path {
	if p == nil {
		return nil
	}
	// Create a copy of the Path to avoid modifying the original
	dereferenced := *p

	// Dereference each operation (GET, POST, etc.)
	if p.Get != nil {
		dereferenced.Get = p.Get.dereference(components)
	}
	if p.Put != nil {
		dereferenced.Put = p.Put.dereference(components)
	}
	if p.Post != nil {
		dereferenced.Post = p.Post.dereference(components)
	}
	if p.Delete != nil {
		dereferenced.Delete = p.Delete.dereference(components)
	}
	if p.Options != nil {
		dereferenced.Options = p.Options.dereference(components)
	}
	if p.Head != nil {
		dereferenced.Head = p.Head.dereference(components)
	}
	if p.Patch != nil {
		dereferenced.Patch = p.Patch.dereference(components)
	}
	if p.Trace != nil {
		dereferenced.Trace = p.Trace.dereference(components)
	}

	return &dereferenced
}

// Dereference method for Operation struct
func (o *Operation) dereference(components *Components) *Operation {
	if o == nil {
		return nil
	}
	// Create a copy of the Operation to avoid modifying the original
	dereferenced := *o

	// Dereference RequestBody
	if o.RequestBody != nil {
		dereferenced.RequestBody = o.RequestBody.dereference(components)
	}

	// Dereference Responses
	for statusCode, response := range dereferenced.Responses {
		dereferenced.Responses[statusCode] = response.dereference(components)
	}

	// Dereference Parameters
	for i, parameter := range dereferenced.Parameters {
		dereferenced.Parameters[i] = parameter.dereference(components)
	}

	return &dereferenced
}

// Dereference method for Components struct
func (c *Components) dereference() *Components {
	if c == nil {
		return nil
	}
	// Create a copy of the Components to avoid modifying the original
	dereferenced := *c

	// Dereference Schemas
	for schemaKey, schema := range dereferenced.Schemas {
		if schema.Ref != "" {
			dereferenced.Schemas[schemaKey] = schema.dereference(&dereferenced)
		}
	}

	// Dereference Responses
	for responseKey, response := range dereferenced.Responses {
		if response.Ref != "" {
			dereferenced.Responses[responseKey] = response.dereference(&dereferenced)
		}
	}

	// Dereference Parameters
	for parameterKey, parameter := range dereferenced.Parameters {
		if parameter.Ref != "" {
			dereferenced.Parameters[parameterKey] = parameter.dereference(&dereferenced)
		}
	}

	// Dereference RequestBodies
	for requestBodyKey, requestBody := range dereferenced.RequestBodies {
		if requestBody.Ref != "" {
			dereferenced.RequestBodies[requestBodyKey] = requestBody.dereference(&dereferenced)
		}
	}

	// Dereference Headers
	for headerKey, header := range dereferenced.Headers {
		if header.Ref != "" {
			dereferenced.Headers[headerKey] = header.dereference(&dereferenced)
		}
	}

	// Dereference SecuritySchemes
	for securitySchemeKey, securityScheme := range dereferenced.SecuritySchemes {
		if securityScheme.Ref != "" {
			dereferenced.SecuritySchemes[securitySchemeKey] = securityScheme.dereference(&dereferenced)
		}
	}

	return &dereferenced
}

// Dereference method for Schema struct
func (s *Schema) dereference(components *Components) *Schema {
	if s == nil || components == nil {
		return s
	}
	if s.Ref != "" {
		referencedSchema, found := components.Schemas[getComponentName(s.Ref)]
		if found {
			return referencedSchema.dereference(components)
		}
		fmt.Printf("Error: Schema ref '%s' not found.\n", s.Ref)
	}
	return s
}

// Dereference method for Response struct
func (r *Response) dereference(components *Components) *Response {
	if r == nil || components == nil {
		return r
	}
	if r.Ref != "" {
		referencedResponse, found := components.Responses[getComponentName(r.Ref)]
		if found {
			return referencedResponse.dereference(components)
		}
		fmt.Printf("Error: Response ref '%s' not found.\n", r.Ref)
	}
	return r
}

// Dereference method for Parameter struct
func (p *Parameter) dereference(components *Components) *Parameter {
	if p == nil || components == nil {
		return p
	}
	if p.Ref != "" {
		referencedParameter, found := components.Parameters[getComponentName(p.Ref)]
		if found {
			return referencedParameter.dereference(components)
		}
		fmt.Printf("Error: Parameter ref '%s' not found.\n", p.Ref)
	}
	return p
}

// Dereference method for RequestBody struct
func (rb *RequestBody) dereference(components *Components) *RequestBody {
	if rb == nil || components == nil {
		return rb
	}
	if rb.Ref != "" {
		referencedRequestBody, found := components.RequestBodies[getComponentName(rb.Ref)]
		if found {
			return referencedRequestBody.dereference(components)
		}
		fmt.Printf("Error: RequestBody ref '%s' not found.\n", rb.Ref)
	}
	return rb
}

// Dereference method for Header struct
func (h *Header) dereference(components *Components) *Header {
	if h == nil || components == nil {
		return h
	}
	if h.Ref != "" {
		referencedHeader, found := components.Headers[getComponentName(h.Ref)]
		if found {
			return referencedHeader.dereference(components)
		}
		fmt.Printf("Error: Header ref '%s' not found.\n", h.Ref)
	}
	return h
}

// Dereference method for SecurityScheme struct
func (ss *SecurityScheme) dereference(components *Components) *SecurityScheme {
	if ss == nil || components == nil {
		return ss
	}
	if ss.Ref != "" {
		referencedSecurityScheme, found := components.SecuritySchemes[getComponentName(ss.Ref)]
		if found {
			return referencedSecurityScheme.dereference(components)
		}
		fmt.Printf("Error: SecurityScheme ref '%s' not found.\n", ss.Ref)
	}
	return ss
}

// resolveRefPath is a helper function to resolve $ref strings within components for Path objects
func resolveRefPath(ref string, components *Components) (*Path, error) {
	refComponent := getComponentType(ref)
	refName := getComponentName(ref)

	switch strings.ToLower(refComponent) {
	case "schemas":
		components.Schemas[refName].dereference(components)
	case "responses":
		components.Responses[refName].dereference(components)
	case "parameters":
		components.Parameters[refName].dereference(components)
	case "requestbodies":
		components.RequestBodies[refName].dereference(components)
	case "headers":
		components.Headers[refName].dereference(components)
	case "securityschemes":
		components.SecuritySchemes[refName].dereference(components)
	default:
		return nil, fmt.Errorf("reference '%s' not found", ref)
	}
	return nil, fmt.Errorf("reference '%s' not found", ref)
}

var (
	componentRegex = regexp.MustCompile(`^#/components/(.*)/(.*)$`)
)

func getComponentType(ref string) string {
	matches := componentRegex.FindStringSubmatch(ref)
	if len(matches) != 3 {
		return ""
	}
	return matches[1]
}

func getComponentName(ref string) string {
	matches := componentRegex.FindStringSubmatch(ref)
	if len(matches) != 3 {
		return ""
	}
	return matches[2]
}
