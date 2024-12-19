package controller

import (
	"encoding/json"
	"fmt"

	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/internal/cfg"
	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/internal/transformers"
)

// Compile constructs the environment configuration for a specified target environment.
// It first builds the base environment configuration and then the target environment configuration.
// After that, it merges both configurations and applies transformations to generate the final compiled configuration.
//
// Parameters:
//   - targetEnv: A string representing the name of the target environment for which the configuration is to be compiled.
//
// Returns:
//   - A pointer to the compiled environment configuration (*cfg.EnvConfig) if successful.
//   - An error if any step in the process fails, providing context about the failure.
func (c *Client) Compile(targetEnv string) (*cfg.EnvConfig, error) {
	// Build the base environment configuration.
	baseEnvCfgBuilt, err := c.buildBaseEnvConfig()
	if err != nil {
		return nil, fmt.Errorf("error building base environment configuration: %w; ensure that the base configuration is correctly defined and accessible", err)
	}

	// Build the target environment configuration.
	targetEnvCfgBuilt, err := c.BuildTargetEnvConfig(targetEnv)
	if err != nil {
		return nil, fmt.Errorf("error building target environment configuration for '%s': %w; please check the target environment name and its associated settings", targetEnv, err)
	}

	// Merge the base and target environment configurations.
	mergedCfg, err := cfg.MergeConfigs(baseEnvCfgBuilt, targetEnvCfgBuilt)
	if err != nil {
		return nil, fmt.Errorf("failed to compile target environment %s: error occurred during the merging of base and target environment configurations", targetEnv)
	}

	// Create a new transformer for environment variables based on the merged configuration.
	envVarsTransformer := transformers.NewEnvVarsTransformer(mergedCfg)

	// Get the updated configuration after applying environment variable transformations.
	compiledConfig, err := envVarsTransformer.GetUpdatedConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to compile configuration: %w", err)
	}

	// Create a new stacks transformer for validating the stacks in the compiled configuration.
	stacksTransformer := transformers.NewStacksTransformer(compiledConfig, c.Paths.Terragrunt)

	// Validate the stacks to ensure they are correctly configured.
	if err := stacksTransformer.ValidateStacks(); err != nil {
		return nil, fmt.Errorf("failed to compile stacks in %s: %w", c.Paths.Terragrunt, err)
	}

	// Return the successfully compiled configuration.
	return compiledConfig, nil
}

// EnvCfgCompiledToJSON converts the provided compiled environment configuration into a JSON string format.
// It utilizes indentation for better readability of the JSON output.
//
// Parameters:
//   - compiledCfg: A pointer to the compiled environment configuration (*cfg.EnvConfig) that needs to be converted to JSON.
//
// Returns:
//   - A string containing the JSON representation of the compiled environment configuration if successful.
//   - An error if the marshaling process fails, providing context about the failure.
//
// This function is useful for logging, debugging, or exporting the configuration in a human-readable format.
func (c *Client) EnvCfgCompiledToJSON(compiledCfg *cfg.EnvConfig) (string, error) {
	// Marshal the compiled configuration to JSON with indentation for better readability.
	jsonConfig, err := json.MarshalIndent(compiledCfg, "", "  ")

	if err != nil {
		return "", fmt.Errorf("error marshaling the compiled environment configuration to JSON: %w; ensure the configuration is valid and properly structured", err)
	}

	// Return the JSON string representation of the compiled configuration.
	return string(jsonConfig), nil
}

// ValidateInfrastructureHierarchy checks the validity and existence of infrastructure components
// across configuration and filesystem, ensuring the requested stack, layer, and component
// are correctly defined and accessible.
//
// Parameters:
//   - compiledCfg: A pointer to the compiled environment configuration (*cfg.EnvConfig) that needs to be validated.
//   - stackName: The name of the stack to be validated.
//   - layerName: The name of the layer to be validated.
//   - componentName: The name of the component to be validated.
//
// Returns:
//   - An error if the validation process fails, providing context about the failure.
//
// This function is useful for ensuring the correctness of the infrastructure hierarchy.
func (c *Client) ValidateInfrastructureHierarchy(compiledCfg *cfg.EnvConfig, stackName string, layerName string, componentName string) error {
	// Initialize stacks transformer
	stacksTransformer := transformers.NewStacksTransformer(compiledCfg, c.Paths.Terragrunt)

	// Validate stack existence (required for all operations)
	if stackName == "" {
		return fmt.Errorf("stack name is required")
	}

	if err := stacksTransformer.ValidateRequestedStack(stackName); err != nil {
		return fmt.Errorf("invalid stack: %w", err)
	}

	// Validate layer hierarchy and existence
	if layerName != "" {
		// If layer is specified, validate it exists in the stack
		if err := stacksTransformer.ValidateRequestedLayer(stackName, layerName); err != nil {
			return fmt.Errorf("invalid layer: %w", err)
		}
	}

	// Validate component hierarchy and existence
	if componentName != "" {
		// Component requires both stack and layer to be specified
		if layerName == "" {
			return fmt.Errorf("cannot specify component '%s' without a layer", componentName)
		}

		// If component is specified, validate it exists in the layer
		if err := stacksTransformer.ValidateRequestedComponent(stackName, layerName, componentName); err != nil {
			return fmt.Errorf("invalid component: %w", err)
		}
	}

	return nil
}
