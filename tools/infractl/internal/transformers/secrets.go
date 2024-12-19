package transformers

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/internal/cfg"
)

// SecretsTransformer provides methods for managing and validating secrets
type SecretsTransformer struct {
	EnvConfig *cfg.EnvConfig
}

// NewSecretsTransformer creates a new SecretsTransformer
func NewSecretsTransformer(envConfig *cfg.EnvConfig) *SecretsTransformer {
	return &SecretsTransformer{
		EnvConfig: envConfig,
	}
}

// ValidateSecrets performs comprehensive validation on secrets
func (st *SecretsTransformer) ValidateSecrets() error {
	// Create an environment transformer to leverage its resolution logic
	envTransformer := NewEnvVarsTransformer(st.EnvConfig)
	var missingSecretVars []string

	// Validate secrets section
	for groupName, secretGroup := range st.EnvConfig.Secrets {
		for secretKey, secretValue := range secretGroup {
			// Use environment transformer's resolution logic
			resolvedValue, err := envTransformer.resolveValue(secretValue)

			// If resolution fails or returns the original value, it means variables are missing
			if err != nil || resolvedValue == secretValue {
				missingVar := extractMissingVariable(secretValue)
				if missingVar != "" {
					missingSecretVars = append(missingSecretVars, missingVar)
					fmt.Printf("Missing secret variable in %s.%s: %s\n", groupName, secretKey, missingVar)
				}
			}
		}
	}

	// Return error if any missing secret variables found
	if len(missingSecretVars) > 0 {
		return fmt.Errorf("missing secret environment variables: %v", missingSecretVars)
	}

	return nil
}

// extractMissingVariable extracts the first missing variable from a secret value
func extractMissingVariable(secretValue string) string {
	re := regexp.MustCompile(`\${([^}:-]+)(?::-([^}]*))?}`)
	matches := re.FindAllStringSubmatch(secretValue, -1)

	for _, match := range matches {
		if len(match) < 2 {
			continue
		}

		envVar := match[1]
		defaultVal := ""
		if len(match) > 2 {
			defaultVal = match[2]
		}

		// If no default or secrets reference, return the env var
		if defaultVal == "" || !strings.HasPrefix(defaultVal, "secrets.") {
			return envVar
		}
	}

	return ""
}

// ValidateSecrets is a wrapper function for external use
func ValidateSecrets(envConfig *cfg.EnvConfig) error {
	transformer := NewSecretsTransformer(envConfig)
	return transformer.ValidateSecrets()
}
