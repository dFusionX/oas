package oas

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

// Validate validates the given interface against the schema.
func (s *Schema) Validate(i interface{}, strict bool, stopOnFailure bool) error {
	if s == nil {
		return errors.New("schema is nil")
	}

	if i == nil {
		return errors.New("input is nil")
	}

	// Get the type of the provided interface
	value := reflect.ValueOf(i)

	switch s.Type {
	case "string":
		if err := s.validateString(value); err != nil {
			return err
		}
	case "integer":
		if err := s.validateInteger(value); err != nil {
			return err
		}
	case "number":
		if err := s.validateNumber(value); err != nil {
			return err
		}
	case "boolean":
		if value.Kind() != reflect.Bool {
			return fmt.Errorf("expected boolean, got %s", value.Kind().String())
		}
	case "array":
		if err := s.validateArray(value, strict, stopOnFailure); err != nil {
			return err
		}
	case "object":
		if err := s.validateObject(value, strict, stopOnFailure); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported schema type: %s", s.Type)
	}

	return nil
}

func (s *Schema) validateString(value reflect.Value) error {
	if value.Kind() != reflect.String {
		return fmt.Errorf("expected string, got %s", value.Kind().String())
	}
	if s.MinLength != nil && len(value.String()) < *s.MinLength {
		return fmt.Errorf("string length is less than minimum length of %d", *s.MinLength)
	}
	if s.MaxLength != nil && len(value.String()) > *s.MaxLength {
		return fmt.Errorf("string length exceeds maximum length of %d", *s.MaxLength)
	}
	if s.Pattern != nil {
		patternRegexp, err := regexp.Compile(*s.Pattern)
		if err != nil {
			return fmt.Errorf("error compiling pattern: %s", *s.Pattern)
		}
		if !patternRegexp.MatchString(value.String()) {
			return fmt.Errorf("string does not match pattern: %s", *s.Pattern)
		}
	}
	if s.Enum != nil {
		found := false
		for _, e := range s.Enum {
			if e == value.String() {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("string is not one of the allowed values: %v", s.Enum)
		}
	}
	if s.Format != "" {
		// TODO: Further string format validations can be added here (e.g., email, date, etc.)
	}
	return nil
}

func (s *Schema) validateInteger(value reflect.Value) error {
	if value.Kind() != reflect.Int && value.Kind() != reflect.Int64 && value.Kind() != reflect.Int32 {
		return fmt.Errorf("expected integer, got %s", value.Kind().String())
	}
	if s.Minimum != nil && value.Int() < int64(*s.Minimum) {
		return fmt.Errorf("integer value is less than minimum value of %d", s.Minimum)
	}
	if s.Maximum != nil && value.Int() > int64(*s.Maximum) {
		return fmt.Errorf("integer value exceeds maximum value of %d", s.Maximum)
	}
	if s.MultipleOf != nil && value.Int()%int64(*s.MultipleOf) != 0 {
		return fmt.Errorf("integer value is not a multiple of %d", s.MultipleOf)
	}
	if s.ExclusiveMinimum && s.Minimum != nil && value.Int() == int64(*s.Minimum) {
		return fmt.Errorf("integer value is equal to exclusive minimum value of %d", s.Minimum)
	}
	if s.ExclusiveMaximum && s.Maximum != nil && value.Int() == int64(*s.Maximum) {
		return fmt.Errorf("integer value is equal to exclusive maximum value of %d", s.Maximum)
	}
	if s.Enum != nil {
		found := false
		for _, e := range s.Enum {
			if e == value.Int() {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("integer value is not one of the allowed values: %v", s.Enum)
		}
	}
	return nil
}

func (s *Schema) validateNumber(value reflect.Value) error {
	if value.Kind() != reflect.Float64 && value.Kind() != reflect.Float32 {
		return fmt.Errorf("expected number, got %s", value.Kind().String())
	}
	if s.Minimum != nil && value.Float() < *s.Minimum {
		return fmt.Errorf("number value is less than minimum value of %f", *s.Minimum)
	}
	if s.Maximum != nil && value.Float() > *s.Maximum {
		return fmt.Errorf("number value exceeds maximum value of %f", *s.Maximum)
	}
	if s.MultipleOf != nil && float64(int(value.Float()/(*s.MultipleOf))) != value.Float()/(*s.MultipleOf) {
		return fmt.Errorf("number value is not a multiple of %f", *s.MultipleOf)
	}
	if s.ExclusiveMinimum && s.Minimum != nil && value.Float() == *s.Minimum {
		return fmt.Errorf("number value is equal to exclusive minimum value of %f", *s.Minimum)
	}
	if s.ExclusiveMaximum && s.Maximum != nil && value.Float() == *s.Maximum {
		return fmt.Errorf("number value is equal to exclusive maximum value of %f", *s.Maximum)
	}
	if s.Enum != nil {
		found := false
		for _, e := range s.Enum {
			if e == value.Float() {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("number value is not one of the allowed values: %v", s.Enum)
		}
	}
	return nil
}

func (s *Schema) validateArray(value reflect.Value, strict, stopOnFailure bool) error {
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return fmt.Errorf("expected array, got %s", value.Kind().String())
	}
	// Array item validation
	for i := 0; i < value.Len(); i++ {
		item := value.Index(i).Interface()
		if s.Items != nil {
			if err := s.Items.Validate(item, strict, stopOnFailure); err != nil {
				return fmt.Errorf("array item at index %d failed validation: %v", i, err)
			}
		}
	}
	return nil
}

func (s *Schema) validateObject(value reflect.Value, strict, stopOnFailure bool) error {
	if value.Kind() != reflect.Map {
		return fmt.Errorf("expected object, got %s", value.Kind().String())
	}

	requiredMap := make(map[string]bool, len(s.Required))
	for _, r := range s.Required {
		requiredMap[r] = false
	}

	for propName, propSchema := range s.Properties {
		propValue := value.MapIndex(reflect.ValueOf(propName))
		if propValue.IsValid() {
			if err := propSchema.Validate(propValue.Interface(), strict, stopOnFailure); err != nil {
				return fmt.Errorf("property '%s' failed validation: %v", propName, err)
			}
		}
		if _, ok := requiredMap[propName]; ok {
			requiredMap[propName] = true
		} else if strict {
			return fmt.Errorf("property '%s' is not allowed", propName)
		}
	}

	for propName, required := range requiredMap {
		if !required {
			return fmt.Errorf("required property '%s' is missing", propName)
		}
	}

	if s.MinProperties != nil && value.Len() < *s.MinProperties {
		return fmt.Errorf("object has fewer properties than minimum of %d", *s.MinProperties)
	}
	if s.MaxProperties != nil && value.Len() > *s.MaxProperties {
		return fmt.Errorf("object has more properties than maximum of %d", *s.MaxProperties)
	}
	return nil
}
