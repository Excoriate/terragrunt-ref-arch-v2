package utils

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// ConvertToBoolean is a method on ConfigProcessor to convert various inputs to boolean
// ToBoolean converts an input value to a boolean representation.
// It supports multiple input types and formats:
//
// For boolean inputs:
//   - Directly returns the boolean value if input is already a bool
//
// For string inputs (case-insensitive):
//   - Converts to true: "true", "1", "yes", "y"
//   - Converts to false: "false", "0", "no", "n"
//
// For other input types:
//   - Returns an error indicating the input cannot be converted
//
// Parameters:
//   - value: An interface{} that can be a bool or string
//
// Returns:
//   - bool: The converted boolean value
//   - error: An error if the conversion is not possible
func ToBoolean(value interface{}) (bool, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	case string:
		v = strings.ToLower(v)
		switch v {
		case "true", "1", "yes", "y":
			return true, nil
		case "false", "0", "no", "n":
			return false, nil
		default:
			return false, fmt.Errorf("cannot convert %s to boolean", v)
		}
	default:
		return false, fmt.Errorf("cannot convert %v to boolean", value)
	}
}

// SafeStringConvert attempts to convert interface{} to string with type safety
// SafeStringConvert attempts to convert an interface{} value to a string representation.
// It provides a type-safe way to handle various input types and returns a default value
// if the input is nil or cannot be converted.
//
// Parameters:
//   - value: An interface{} that can be of various types including string, bool, int, int64, or float64.
//   - defaultValue: A string that will be returned if the input value is nil or cannot be converted.
//
// Returns:
//   - string: The string representation of the input value if conversion is successful,
//     otherwise returns the provided defaultValue.
//
// The function handles the following types:
//   - If the value is a string, it returns the string directly.
//   - If the value is a boolean, it converts it to a string using strconv.FormatBool.
//   - If the value is an integer (int, int64) or a float (float64), it converts it to a string
//     using fmt.Sprintf.
//   - For any other types or if the value is nil, it returns the defaultValue.
func SafeStringConvert(value interface{}, defaultValue string) string {
	if value == nil {
		return defaultValue
	}

	switch v := value.(type) {
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	case int, int64, float64:
		return fmt.Sprintf("%v", v)
	default:
		return defaultValue
	}
}

// Helper function to convert map to struct for various types
// MapToStruct converts a map of string keys to a struct using YAML encoding/decoding.
//
// This function takes an input map where the keys are strings and the values are of type
// interface{}. It marshals the map into a YAML format and then unmarshals it into the
// provided output struct. This is useful for converting dynamic data structures into
// strongly typed Go structs, allowing for easier manipulation and access to the data.
//
// Parameters:
//   - input: A map[string]interface{} that contains the data to be converted. The keys
//     represent the field names of the struct, and the values represent the corresponding
//     field values.
//   - output: An interface{} that should be a pointer to a struct where the data from
//     the input map will be stored. The struct must have fields that correspond to the
//     keys in the input map.
//
// Returns:
//   - error: Returns nil if the conversion is successful. If there is an error during
//     marshalling or unmarshalling, it returns an error that wraps the original error
//     with a descriptive message.
//
// Example usage:
//
//	type MyStruct struct {
//	    Name  string `yaml:"name"`
//	    Age   int    `yaml:"age"`
//	}
//	input := map[string]interface{}{"name": "John", "age": 30}
//	var output MyStruct
//	err := MapToStruct(input, &output)
//	if err != nil {
//	    log.Fatalf("error converting map to struct: %v", err)
//	}
func MapToStruct(input map[string]interface{}, output interface{}) error {
	// Use yaml.Marshal to convert map to struct
	data, err := yaml.Marshal(input)
	if err != nil {
		return fmt.Errorf("marshalling input map: %w", err)
	}

	if err := yaml.Unmarshal(data, output); err != nil {
		return fmt.Errorf("unmarshalling to struct: %w", err)
	}

	return nil
}
