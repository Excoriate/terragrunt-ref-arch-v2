package envars

import (
	"fmt"
	"os"
	"regexp"
)

// ExpandEnvironmentVariable expands environment variables with optional default values.
// It supports two variable expansion formats:
//   - ${VAR}: Expands to the value of VAR
//   - ${VAR:-default}: Expands to the value of VAR, or 'default' if VAR is unset
//
// Parameters:
//   - template: A string containing environment variable placeholders to be expanded
//
// Returns:
//   - The expanded string with environment variables replaced
//   - An error if expansion fails (currently always returns nil)
func ExpandEnvironmentVariable(template string) (string, error) {
	re := regexp.MustCompile(`\${([^}:-]+)(?::-([^}]*))?}`)

	expanded := re.ReplaceAllStringFunc(template, func(match string) string {
		matches := re.FindStringSubmatch(match)
		if len(matches) < 2 {
			return match
		}

		envVar := matches[1]
		defaultVal := ""
		if len(matches) > 2 {
			defaultVal = matches[2]
		}

		value := os.Getenv(envVar)

		if value == "" {
			if defaultVal != "" {
				return defaultVal
			}
			return defaultVal
		}

		return value
	})

	return expanded, nil
}

// ExtractEnvVarName extracts the environment variable name from a secret template.
// It supports both ${VAR} and ${VAR:-default} patterns.
//
// Parameters:
//   - secretTemplate: An interface{} representing the secret template (expected to be a string)
//
// Returns:
//   - The extracted environment variable name
//   - An empty string if no valid environment variable is found or the template is invalid
func ExtractEnvVarName(secretTemplate interface{}) string {
	strTemplate, ok := secretTemplate.(string)
	if !ok {
		return ""
	}

	// Handle both ${VAR} and ${VAR:-default} patterns
	re := regexp.MustCompile(`\${([^}:-]+)(?::-[^}]*)?}`)
	matches := re.FindStringSubmatch(strTemplate)

	if len(matches) < 2 {
		return ""
	}

	return matches[1]
}

// CleanEnvVarsByKeys removes the specified environment variables from the current process.
//
// Parameters:
//   - keys: A slice of strings representing the names of the environment variables to be removed.
//
// This function iterates over the provided keys and calls os.Unsetenv for each key,
// effectively removing the environment variable from the process's environment.
//
// It is important to note that this operation only affects the environment variables
// of the current process and does not impact the system-wide environment variables or
// those of other processes. Additionally, if a key does not exist in the environment,
// os.Unsetenv will simply do nothing for that key, ensuring that no error is returned
// for non-existent variables.
func CleanEnvVarsByKeys(keys []string) error {
	for _, key := range keys {
		if _, exists := os.LookupEnv(key); !exists {
			return fmt.Errorf("environment variable '%s' does not exist", key)
		}
		os.Unsetenv(key)
	}

	return nil
}
