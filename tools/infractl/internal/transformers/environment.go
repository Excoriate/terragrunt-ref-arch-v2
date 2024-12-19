package transformers

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/internal/cfg"
)

// EnvVarsTransformer handles environment variable expansion
type EnvVarsTransformer struct {
	EnvConfig *cfg.EnvConfig
}

// NewEnvVarsTransformer creates a new transformer
func NewEnvVarsTransformer(envConfig *cfg.EnvConfig) *EnvVarsTransformer {
	return &EnvVarsTransformer{EnvConfig: envConfig}
}

// resolveValue attempts to resolve a value with environment variable and secrets fallback
func (t *EnvVarsTransformer) resolveValue(value string) (string, error) {
	re := regexp.MustCompile(`\${([^}:-]+)(?::-([^}]*))?}`)

	return re.ReplaceAllStringFunc(value, func(match string) string {
		matches := re.FindStringSubmatch(match)
		if len(matches) < 2 {
			return match
		}

		envVar := matches[1]
		defaultVal := ""
		if len(matches) > 2 {
			defaultVal = matches[2]
		}

		// First, check direct environment variable
		envValue := os.Getenv(envVar)
		if envValue != "" {
			return envValue
		}

		// If no direct env var, check for secrets reference
		if strings.HasPrefix(defaultVal, "secrets.") {
			parts := strings.Split(defaultVal, ".")
			if len(parts) == 3 {
				secretGroup := parts[1]
				secretKey := parts[2]

				// Look up in secrets section
				if secretGroup, exists := t.EnvConfig.Secrets[secretGroup]; exists {
					if secretValue, exists := secretGroup[secretKey]; exists {
						// Recursively resolve the secret value
						resolvedSecret, err := t.resolveValue(secretValue)
						if err != nil {
							return match // Return original if resolution fails
						}
						return resolvedSecret
					}
				}
			}
		}

		// If no resolution found, return default or original
		if defaultVal != "" {
			return defaultVal
		}
		return match
	}), nil
}

// expandProviderConfig handles complex provider configuration expansion
func (t *EnvVarsTransformer) expandProviderConfig(providerName string, providerConfig cfg.ProviderConfig) (cfg.ProviderConfig, error) {
	// Expand provider config values
	expandedConfig := make(map[string]interface{})
	for configKey, configValue := range providerConfig.Config {
		switch val := configValue.(type) {
		case string:
			resolvedValue, err := t.resolveValue(val)
			if err != nil {
				return providerConfig, fmt.Errorf("resolving provider %s config key %s: %w", providerName, configKey, err)
			}
			expandedConfig[configKey] = resolvedValue
		default:
			expandedConfig[configKey] = val
		}
	}

	// Expand version constraint source
	expandedVersionSource, err := t.resolveValue(providerConfig.VersionConstraint.Source)
	if err != nil {
		return providerConfig, fmt.Errorf("resolving provider %s version source: %w", providerName, err)
	}

	return cfg.ProviderConfig{
		Config: expandedConfig,
		VersionConstraint: cfg.VersionConstraint{
			Source:          expandedVersionSource,
			RequiredVersion: providerConfig.VersionConstraint.RequiredVersion,
			Enabled:         providerConfig.VersionConstraint.Enabled,
		},
	}, nil
}

// expandMapStringInterface expands a map[string]interface{} with environment variables
func (t *EnvVarsTransformer) expandMapStringInterface(input map[string]interface{}) (map[string]interface{}, error) {
	expanded := make(map[string]interface{})

	for k, v := range input {
		switch val := v.(type) {
		case string:
			resolvedValue, err := t.resolveValue(val)
			if err != nil {
				return nil, err
			}
			expanded[k] = resolvedValue
		case map[string]interface{}:
			nestedExpanded, err := t.expandMapStringInterface(val)
			if err != nil {
				return nil, err
			}
			expanded[k] = nestedExpanded
		default:
			expanded[k] = val
		}
	}

	return expanded, nil
}

// GetUpdatedConfig returns a new EnvConfig with environment variables and secrets expanded
func (t *EnvVarsTransformer) GetUpdatedConfig() (*cfg.EnvConfig, error) {
	updatedConfig := &cfg.EnvConfig{}

	// Expand Secrets section first (to ensure secrets are resolved)
	updatedConfig.Secrets = make(cfg.Secrets)
	for groupName, secretGroup := range t.EnvConfig.Secrets {
		expandedSecretGroup := make(map[string]string)
		for secretKey, secretValue := range secretGroup {
			expandedValue, err := t.resolveValue(secretValue)
			if err != nil {
				return nil, fmt.Errorf("expanding secret %s.%s: %w", groupName, secretKey, err)
			}
			expandedSecretGroup[secretKey] = expandedValue
		}
		updatedConfig.Secrets[groupName] = expandedSecretGroup
	}

	// Expand Providers section with secrets resolution
	updatedConfig.Providers = make(cfg.Providers)
	for providerName, providerConfig := range t.EnvConfig.Providers {
		expandedProviderConfig, err := t.expandProviderConfig(providerName, providerConfig)
		if err != nil {
			return nil, fmt.Errorf("expanding provider %s: %w", providerName, err)
		}
		updatedConfig.Providers[providerName] = expandedProviderConfig
	}

	// Expand Config section
	configMap, err := t.expandMapStringInterface(map[string]interface{}{
		"version":      t.EnvConfig.Config.Version,
		"last_updated": t.EnvConfig.Config.LastUpdated,
		"description":  t.EnvConfig.Config.Description,
	})
	if err != nil {
		return nil, fmt.Errorf("expanding config section: %w", err)
	}
	updatedConfig.Config = cfg.RootConfig{
		Version:     configMap["version"].(string),
		LastUpdated: configMap["last_updated"].(string),
		Description: configMap["description"].(string),
	}

	// Expand Git section
	gitMap, err := t.expandMapStringInterface(map[string]interface{}{
		"base_url": t.EnvConfig.Git.BaseURL,
	})
	if err != nil {
		return nil, fmt.Errorf("expanding git section: %w", err)
	}
	updatedConfig.Git = cfg.Git{
		BaseURL: gitMap["base_url"].(string),
	}

	// Expand Product section
	productMap, err := t.expandMapStringInterface(map[string]interface{}{
		"name":        t.EnvConfig.Product.Name,
		"version":     t.EnvConfig.Product.Version,
		"description": t.EnvConfig.Product.Description,
	})
	if err != nil {
		return nil, fmt.Errorf("expanding product section: %w", err)
	}
	updatedConfig.Product = cfg.Product{
		Name:           productMap["name"].(string),
		Version:        productMap["version"].(string),
		Description:    productMap["description"].(string),
		UseAsStackTags: t.EnvConfig.Product.UseAsStackTags,
	}

	// Expand IAC section
	iacMap, err := t.expandMapStringInterface(map[string]interface{}{
		"versions.terraform_version_default":  t.EnvConfig.IAC.Versions.TerraformVersionDefault,
		"versions.terragrunt_version_default": t.EnvConfig.IAC.Versions.TerragruntVersionDefault,
		"remote_state.s3.bucket":              t.EnvConfig.IAC.RemoteState.S3.Bucket,
		"remote_state.s3.lock_table":          t.EnvConfig.IAC.RemoteState.S3.LockTable,
		"remote_state.s3.region":              t.EnvConfig.IAC.RemoteState.S3.Region,
	})
	if err != nil {
		return nil, fmt.Errorf("expanding iac section: %w", err)
	}
	updatedConfig.IAC = cfg.IaC{
		Versions: cfg.IaCVersions{
			TerraformVersionDefault:  iacMap["versions.terraform_version_default"].(string),
			TerragruntVersionDefault: iacMap["versions.terragrunt_version_default"].(string),
		},
		RemoteState: cfg.RemoteState{
			S3: cfg.RemoteStateS3{
				Bucket:    iacMap["remote_state.s3.bucket"].(string),
				LockTable: iacMap["remote_state.s3.lock_table"].(string),
				Region:    iacMap["remote_state.s3.region"].(string),
			},
		},
	}

	// Expand Stacks section
	updatedConfig.Stacks = make([]cfg.StackConfig, len(t.EnvConfig.Stacks))
	for i, stack := range t.EnvConfig.Stacks {
		// Expand stack name and tags
		stackNameMap, err := t.expandMapStringInterface(map[string]interface{}{
			"name": stack.Name,
		})
		if err != nil {
			return nil, fmt.Errorf("expanding stack name at index %d: %w", i, err)
		}

		expandedStackTags := make(map[string]string)
		for tagKey, tagValue := range stack.Tags {
			expandedTagValue, err := t.resolveValue(tagValue)
			if err != nil {
				return nil, fmt.Errorf("expanding stack tag %s at index %d: %w", tagKey, i, err)
			}
			expandedStackTags[tagKey] = expandedTagValue
		}

		updatedConfig.Stacks[i] = cfg.StackConfig{
			Name:   stackNameMap["name"].(string),
			Tags:   expandedStackTags,
			Layers: stack.Layers, // Layers remain unchanged
		}
	}

	return updatedConfig, nil
}

// ValidateEnvironmentVariables checks if all referenced environment variables exist
func (t *EnvVarsTransformer) ValidateEnvironmentVariables() error {
	var missingVars []string

	// Helper function to check and track missing variables
	checkVar := func(value string) {
		re := regexp.MustCompile(`\${([^}:-]+)(?::-([^}]*))?}`)
		matches := re.FindAllStringSubmatch(value, -1)

		for _, match := range matches {
			if len(match) < 2 {
				continue
			}

			envVar := match[1]
			defaultVal := ""
			if len(match) > 2 {
				defaultVal = match[2]
			}

			// Check direct environment variable
			if os.Getenv(envVar) == "" {
				// If no direct env var, check for secrets reference
				if strings.HasPrefix(defaultVal, "secrets.") {
					parts := strings.Split(defaultVal, ".")
					if len(parts) == 3 {
						secretGroup := parts[1]
						secretKey := parts[2]

						// Look up in secrets section
						if secretGroup, exists := t.EnvConfig.Secrets[secretGroup]; exists {
							if secretValue, exists := secretGroup[secretKey]; exists {
								// Recursively check secret value
								re := regexp.MustCompile(`\${([^}:-]+)(?::-([^}]*))?}`)
								secretMatches := re.FindAllStringSubmatch(secretValue, -1)

								for _, secretMatch := range secretMatches {
									if len(secretMatch) < 2 {
										continue
									}

									secretEnvVar := secretMatch[1]
									if os.Getenv(secretEnvVar) == "" {
										missingVars = append(missingVars, secretEnvVar)
									}
								}
								continue
							}
						}
					}
				}

				// If no resolution found, add to missing vars
				missingVars = append(missingVars, envVar)
			}
		}
	}

	// Check all sections for environment variables
	// Config section
	checkVar(t.EnvConfig.Config.Version)
	checkVar(t.EnvConfig.Config.LastUpdated)
	checkVar(t.EnvConfig.Config.Description)

	// Git section
	checkVar(t.EnvConfig.Git.BaseURL)

	// Product section
	checkVar(t.EnvConfig.Product.Name)
	checkVar(t.EnvConfig.Product.Version)
	checkVar(t.EnvConfig.Product.Description)

	// Providers section
	for _, providerConfig := range t.EnvConfig.Providers {
		// Check provider config values
		for _, configValue := range providerConfig.Config {
			if strVal, ok := configValue.(string); ok {
				checkVar(strVal)
			}
		}

		// Check version constraint source
		checkVar(providerConfig.VersionConstraint.Source)
	}

	// Secrets section
	for groupName, secretGroup := range t.EnvConfig.Secrets {
		for secretKey, secretValue := range secretGroup {
			checkVar(fmt.Sprintf("%s.%s:%s", groupName, secretKey, secretValue))
		}
	}

	// Stacks section
	for _, stack := range t.EnvConfig.Stacks {
		checkVar(stack.Name)
		for _, tagValue := range stack.Tags {
			checkVar(tagValue)
		}
	}

	// Return error if any missing variables found
	if len(missingVars) > 0 {
		return fmt.Errorf("missing environment variables: %v", missingVars)
	}

	return nil
}
