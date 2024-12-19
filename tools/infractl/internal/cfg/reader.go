package cfg

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/pkg/utils"
	"gopkg.in/yaml.v3"
)

// ConfigurationError represents detailed configuration parsing errors
type ConfigurationError struct {
	Stage        string
	FieldPath    string
	ActualValue  interface{}
	ExpectedType string
	Reason       string
}

func (e *ConfigurationError) Error() string {
	return fmt.Sprintf(
		"Configuration Error in %s: Field '%s' conversion failed. "+
			"Expected %s, got %T with value %v. %s",
		e.Stage, e.FieldPath, e.ExpectedType,
		e.ActualValue, e.ActualValue, e.Reason,
	)
}

// RawEnvConfig represents the intermediate YAML structure
type RawEnvConfig struct {
	Config    map[string]interface{}   `yaml:"config"`
	Git       map[string]interface{}   `yaml:"git"`
	Product   map[string]interface{}   `yaml:"product"`
	IAC       map[string]interface{}   `yaml:"iac"`
	Providers map[string]interface{}   `yaml:"providers"`
	Secrets   map[string]interface{}   `yaml:"secrets"`
	Stacks    []map[string]interface{} `yaml:"stacks"`
}

// transformRootConfig converts raw config to RootConfig
func transformRootConfig(rawConfig map[string]interface{}) (RootConfig, *ConfigurationError) {
	rootConfig := RootConfig{}

	rootConfig.Version = utils.SafeStringConvert(rawConfig["version"], "")
	rootConfig.LastUpdated = utils.SafeStringConvert(rawConfig["last_updated"], time.Now().Format(time.RFC3339))
	rootConfig.Description = utils.SafeStringConvert(rawConfig["description"], "")

	if rootConfig.Version == "" {
		return rootConfig, &ConfigurationError{
			Stage:        "RootConfig",
			FieldPath:    "version",
			ActualValue:  rawConfig["version"],
			ExpectedType: "non-empty string",
			Reason:       "Version is required",
		}
	}

	return rootConfig, nil
}

// GetInfraEnvConfigFromFile reads and parses a YAML configuration file
func GetInfraEnvConfigFromFile(path string) (*EnvConfig, error) {
	// Validate that the file is a YAML file
	if err := utils.IsYAMLFile(path); err != nil {
		return nil, fmt.Errorf("environment configuration file is not a YAML file: %w", err)
	}

	// Check that the file is not empty
	if err := utils.FileHasContent(path); err != nil {
		return nil, fmt.Errorf("environment configuration file is empty: %w", err)
	}

	// Read the entire file contents
	cfgFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading environment configuration file: %w", err)
	}

	// Unmarshal into RawEnvConfig first
	var rawCfg RawEnvConfig
	if err := yaml.Unmarshal(cfgFile, &rawCfg); err != nil {
		return nil, fmt.Errorf("unmarshalling raw environment configuration: %w", err)
	}

	// Create final EnvConfig
	envConfig := &EnvConfig{}

	// Transform and populate each section
	if rootConfig, cfgErr := transformRootConfig(rawCfg.Config); cfgErr != nil {
		return nil, fmt.Errorf("transforming root configuration: %w", cfgErr)
	} else {
		envConfig.Config = rootConfig
	}

	// Populate Git configuration
	if err := utils.MapToStruct(rawCfg.Git, &envConfig.Git); err != nil {
		return nil, fmt.Errorf("populating git configuration: %w", err)
	}

	// Populate Product configuration
	if err := utils.MapToStruct(rawCfg.Product, &envConfig.Product); err != nil {
		return nil, fmt.Errorf("populating product configuration: %w", err)
	}

	// Populate IAC configuration
	if err := utils.MapToStruct(rawCfg.IAC, &envConfig.IAC); err != nil {
		return nil, fmt.Errorf("populating iac configuration: %w", err)
	}

	// Populate Providers configuration
	if err := utils.MapToStruct(rawCfg.Providers, &envConfig.Providers); err != nil {
		return nil, fmt.Errorf("populating providers configuration: %w", err)
	}

	// Populate Secrets configuration
	if err := utils.MapToStruct(rawCfg.Secrets, &envConfig.Secrets); err != nil {
		return nil, fmt.Errorf("populating secrets configuration: %w", err)
	}

	// Populate Stacks configuration
	envConfig.Stacks = make([]StackConfig, len(rawCfg.Stacks))
	for i, stackMap := range rawCfg.Stacks {
		if err := utils.MapToStruct(stackMap, &envConfig.Stacks[i]); err != nil {
			return nil, fmt.Errorf("populating stack configuration at index %d: %w", i, err)
		}
	}

	return envConfig, nil
}

// MergeConfigs merges two EnvConfig structs with target overriding base
// MergeConfigs merges two EnvConfig structs, where the target configuration
// overrides the base configuration. This function is useful for combining
// configurations from different sources, ensuring that the most specific
// settings take precedence over the more general ones.
//
// Parameters:
//   - base: A pointer to the base EnvConfig struct. If this is nil, the
//     function will return the target configuration.
//   - target: A pointer to the target EnvConfig struct. If this is nil,
//     the function will return the base configuration.
//
// Returns:
//   - A pointer to a new EnvConfig struct that contains the merged
//     configurations. If both base and target are nil, the function will
//     return nil.
//   - An error if any issues occur during the merging process.
//
// Best Practices:
//   - Ensure that both base and target configurations are properly
//     initialized before calling this function to avoid unexpected behavior.
func MergeConfigs(base, target *EnvConfig) (*EnvConfig, error) {
	if base == nil {
		return target, nil
	}
	if target == nil {
		return base, nil
	}

	// Create a deep copy of the base configuration to avoid mutating the original
	merged := &EnvConfig{}
	*merged = *base

	// Merge Config (RootConfig) if the target's Config is not empty
	if !reflect.DeepEqual(target.Config, RootConfig{}) {
		merged.Config = target.Config
	}

	// Merge Git configuration if the target's Git is not empty
	if !reflect.DeepEqual(target.Git, Git{}) {
		merged.Git = target.Git
	}

	// Merge Product configuration if the target's Product is not empty
	if !reflect.DeepEqual(target.Product, Product{}) {
		merged.Product = target.Product
	}

	// Merge IAC configuration if the target's IAC is not empty
	if !reflect.DeepEqual(target.IAC, IaC{}) {
		merged.IAC = target.IAC
	}

	// Merge Providers if the target's Providers slice is not empty
	if len(target.Providers) > 0 {
		merged.Providers = target.Providers
	}

	// Merge Secrets if the target's Secrets slice is not empty
	if len(target.Secrets) > 0 {
		merged.Secrets = target.Secrets
	}

	// Merge Stacks if the target's Stacks slice is not empty
	if len(target.Stacks) > 0 {
		merged.Stacks = target.Stacks
	}

	return merged, nil
}
