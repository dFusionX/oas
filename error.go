package oas

// UnmarshalError is an error type for unmarshalling errors.
// Choose the error based on input type(JSON or YAML) you passed.
type UnmarshalError struct {
	JSONErr error
	YAMLErr error
}

func (e *UnmarshalError) Error() string {
	return "error marshalling to JSON: " + e.JSONErr.Error() + ", error marshalling to YAML: " + e.YAMLErr.Error()
}

type ValidationError struct {
	Err   error
	Field string
}

type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	var err string
	for _, v := range e {
		err += v.Field + ": " + v.Err.Error() + "\n"
	}
	return err
}
