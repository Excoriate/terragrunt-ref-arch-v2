package controller

import (
	"fmt"
	"strings"

	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/internal/cfg"
	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/pkg/utils"
)

// isTerragruntInstalled verifies Terragrunt is installed and accessible
func isTerragruntInstalled() error {
	// Simple check using 'which' command to verify Terragrunt installation
	output, err := utils.ExecuteCommand("which", "terragrunt")
	if err != nil {
		return fmt.Errorf("terragrunt is not installed or not in PATH")
	}

	// Trim any whitespace and verify the output is not empty
	terragruntPath := strings.TrimSpace(output)
	if terragruntPath == "" {
		return fmt.Errorf("terragrunt executable path is empty")
	}

	return nil
}

func envConfigurationFilesExistInEnvsDir(envsFilesPath string) error {
	extensions := []string{".yaml", ".yml"}
	var foundExtensions []string

	for _, ext := range extensions {
		foundFiles, err := utils.FoundFilesWithExtensionInPath(envsFilesPath, ext)
		if err != nil {
			return fmt.Errorf("cannot scan configuration files in %s. Check directory permissions", envsFilesPath)
		}

		if len(foundFiles) > 0 {
			return nil
		}
		foundExtensions = append(foundExtensions, ext)
	}

	return fmt.Errorf("no configuration files found in %s with extensions %v. "+
		"Add base.yaml or environment-specific configurations",
		envsFilesPath, foundExtensions)
}

// IsBaseEnvConfigFileValid checks the validity of the base environment configuration file.
// It retrieves the path to the base environment file using the configuration package and
// validates its contents. If the file does not exist or contains invalid configuration,
// an error is returned with a descriptive message.
//
// Returns:
//
//	error: An error indicating the validation failure, or nil if the validation is successful.
func IsBaseEnvConfigFileValid(envsFilesPath string) error {
	baseYamlFilePath := cfg.GetBaseEnvFilePath()

	// Validate the environment file at the specified path
	if err := validateEnvironmentFile(baseYamlFilePath); err != nil {
		return fmt.Errorf("base configuration validation failed for %s: %w. "+
			"Ensure the file exists and contains valid infrastructure configuration", baseYamlFilePath, err)
	}

	return nil
}

// IsTargetEnvConfigFileValid checks the validity of the target environment configuration file.
// It constructs the path to the target environment YAML file based on the provided environment files path
// and the target environment name. The function then validates the contents of the file. If the file does not
// exist or contains invalid configuration, an error is returned with a descriptive message.
//
// Parameters:
//
//	envsFilesPath string: The path to the directory containing environment configuration files.
//	targetEnv string: The name of the target environment for which the configuration file is being validated.
//
// Returns:
//
//	error: An error indicating the validation failure, or nil if the validation is successful.
func IsTargetEnvConfigFileValid(targetEnv string) error {
	targetEnvName := utils.ForceExtensionForFilepath(targetEnv, ".yaml")

	// Validate the environment file at the constructed path
	if err := validateEnvironmentFile(targetEnvName); err != nil {
		return fmt.Errorf("target environment configuration invalid for %s at %s: %w",
			targetEnv, targetEnvName, err)
	}

	return nil
}

// validateEnvironmentFile performs detailed checks on a single environment file
func validateEnvironmentFile(filename string) error {
	if filename == "" {
		return fmt.Errorf("no configuration file path provided")
	}

	// Check if file is a valid YAML
	if err := utils.IsYAMLFile(filename); err != nil {
		return fmt.Errorf("invalid YAML file format: %s. Ensure the file has a .yaml or .yml extension", filename)
	}

	// Check if file is empty
	if err := utils.FileIsEmpty(filename); err == nil {
		return fmt.Errorf("empty configuration file: %s. Add required configuration parameters", filename)
	}

	return nil
}

// IsStackHierarchyConsistent validates the consistency of the infrastructure hierarchy
// Validation rules:
// 1. Stack can be specified alone
// 2. Layer requires a stack
// 3. Component always requires both stack and layer
func IsStackHierarchyConsistent(stack, layer, component string) error {
	// Stack validation
	if stack == "" {
		// If no stack is provided, reject any layer or component
		if layer != "" {
			return fmt.Errorf("layer '%s' requires a stack to be specified", layer)
		}
		if component != "" {
			return fmt.Errorf("component '%s' requires a stack to be specified", component)
		}
		return fmt.Errorf("stack name is required for infrastructure hierarchy")
	}

	// Layer validation - requires stack
	if layer != "" && stack == "" {
		return fmt.Errorf("layer '%s' requires a stack to be specified", layer)
	}

	// Component validation - requires both stack and layer
	if component != "" {
		if stack == "" {
			return fmt.Errorf("component '%s' requires a stack to be specified", component)
		}
		if layer == "" {
			return fmt.Errorf("component '%s' requires a layer to be specified", component)
		}
	}

	return nil
}
