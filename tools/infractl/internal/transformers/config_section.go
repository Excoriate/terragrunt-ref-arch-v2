package transformers

import (
	"fmt"
	"time"
)

// ConfigSectionValidationResult represents the result of config section validation
type ConfigSectionValidationResult struct {
	Valid       bool
	Errors      []string
	ProcessedAt time.Time
}

// ValidateConfigSection validates the configuration section with comprehensive error tracking
func ValidateConfigSection(config map[string]interface{}) (map[string]interface{}, error) {
	result := ConfigSectionValidationResult{
		Valid:       true,
		Errors:      []string{},
		ProcessedAt: time.Now(),
	}

	// Extract config section
	configSection, ok := config["config"].(map[string]interface{})
	if !ok {
		result.Valid = false
		result.Errors = append(result.Errors, "Missing or invalid 'config' section")
		return nil, fmt.Errorf("config section validation failed: %v", result.Errors)
	}

	// Validate version
	version, ok := configSection["version"].(string)
	if !ok {
		result.Valid = false
		result.Errors = append(result.Errors, "Missing or invalid 'version' field")
	} else if version != "1.0.0" {
		result.Valid = false
		result.Errors = append(result.Errors, fmt.Sprintf("Invalid version. Expected '1.0.0', got '%s'", version))
	}

	// Validate last_updated (optional, but should be a valid timestamp if present)
	lastUpdated, ok := configSection["last_updated"].(string)
	if ok {
		_, err := time.Parse(time.RFC3339, lastUpdated)
		if err != nil {
			result.Valid = false
			result.Errors = append(result.Errors, "Invalid 'last_updated' timestamp format")
		}
	}

	if !result.Valid {
		return nil, fmt.Errorf("config section validation failed: %v", result.Errors)
	}

	return config, nil
}
