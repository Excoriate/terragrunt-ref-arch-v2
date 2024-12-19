package cfg

import (
	"fmt"
	"path/filepath"

	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/pkg/utils"
)

// GetGitRepoRoot attempts to find and return the root directory of the current Git repository.
//
// This function uses the utils.FindGitRepoRoot() method to locate the root directory
// of the Git repository containing the current working directory.
//
// Returns:
//   - A string containing the absolute path to the Git repository root
//   - An error if the repository root cannot be found or there's an issue resolving the path
func GetGitRepoRoot() (string, error) {
	gitRepoRootPath, err := utils.FindGitRepoRoot()

	if err != nil {
		return "", fmt.Errorf("failed to resolve the Git repository root path: %w", err)
	}

	return gitRepoRootPath, nil
}

// GetEnvConfigFilesPathAbsolute resolves and returns the absolute filesystem path to the environment configuration files.
//
// This function performs the following steps:
// 1. Attempts to find the root directory of the current Git repository
// 2. Validates that a valid Git repository root was found
// 3. Constructs the full absolute path to the environment configuration files
//
// The function returns:
//   - A string containing the full absolute path to the environment configuration files
//   - An error if the Git repository root cannot be found or there's an issue resolving the path
//
// The returned path is constructed by joining the Git repository root with the relative
// path to the environment configuration files as defined by GetEnvConfigFilesPath().
func GetEnvConfigFilesPathAbsolute() (string, error) {
	var gitRepoRootPath string
	gitRepoRootPath, err := GetGitRepoRoot()

	if err != nil {
		return "", fmt.Errorf("failed to get the environment configuration files path: %w", err)
	}

	envConfigFilesPath := filepath.Join(gitRepoRootPath, EnvCfgDirPathRelative)

	return envConfigFilesPath, nil
}

// GetEnvConfigFilesPathAbsoluteWithGitRepoRoot constructs the absolute filesystem path to the environment
// configuration files based on the provided Git repository root path.
//
// This function takes a string representing the absolute path to the root of a Git repository and
// appends the relative path to the environment configuration files, as defined by the constant
// EnvCfgDirPathRelative. It is important to ensure that the provided gitRepoRootPath is a valid
// directory path that corresponds to the root of a Git repository.
//
// Parameters:
//   - gitRepoRootPath: A string representing the absolute path to the Git repository root.
//
// Returns:
//   - A string containing the full absolute path to the environment configuration files.
//   - An error if there is an issue constructing the path (though this function does not perform
//     validation on the gitRepoRootPath itself).
//
// Example usage:
//
//	path, err := GetEnvConfigFilesPathAbsoluteWithGitRepoRoot("/path/to/git/repo")

func GetEnvConfigFilesPathAbsoluteWithGitRepoRoot(gitRepoRootPath string) string {
	return filepath.Join(gitRepoRootPath, EnvCfgDirPathRelative)
}

// GetInfraCacheDirPathAbsolute resolves and returns the absolute filesystem path to the infrastructure cache directory.
//
// This function performs the following steps:
// 1. Attempts to find the root directory of the current Git repository
// 2. Validates that a valid Git repository root was found
// 3. Constructs the full absolute path to the infrastructure cache directory
//
// The function returns:
//   - A string containing the full absolute path to the infrastructure cache directory
//   - An error if the Git repository root cannot be found or there's an issue resolving the path
//
// The returned path is constructed by joining the Git repository root with the relative
// path to the infrastructure cache directory as defined by GetInfraCacheDirPath().
func GetInfraCacheDirPathAbsolute() (string, error) {
	gitRepoRootPath, err := GetGitRepoRoot()

	if err != nil {
		return "", fmt.Errorf("failed to get the infrastructure cache directory path: %w", err)
	}

	return filepath.Join(gitRepoRootPath, CacheDirPathRelative), nil
}

// GetInfraCacheDirPathAbsoluteWithGitRepoRoot constructs and returns the absolute filesystem path
// to the infrastructure cache directory based on the provided Git repository root path.
//
// This function takes a string representing the absolute path to the root of a Git repository
// and appends the relative path to the infrastructure cache directory, as defined by the constant
// CacheDirPathRelative. It is crucial to ensure that the provided gitRepoRootPath is a valid
// directory path that corresponds to the root of a Git repository.
//
// Parameters:
//   - gitRepoRootPath: A string representing the absolute path to the Git repository root.
//
// Returns:
//   - A string containing the full absolute path to the infrastructure cache directory.
//
// Example usage:
//
//	path, err := GetInfraCacheDirPathAbsoluteWithGitRepoRoot("/path/to/git/repo")
//	if err != nil {
//	    // handle error
func GetInfraCacheDirPathAbsoluteWithGitRepoRoot(gitRepoRootPath string) string {
	return filepath.Join(gitRepoRootPath, CacheDirPathRelative)
}

// GetInfraTerragruntDirPath returns the relative path to the Terragrunt directory within the infrastructure directory.
// Normally, it's infra/terragrunt/
func GetInfraTerragruntDirPath() string {
	return filepath.Join(InfraDir, TerragruntDir)
}

// GetInfraTerragruntDirPathAbsolute resolves and returns the absolute filesystem path to the Terragrunt directory.
//
// This function performs the following steps:
// 1. Attempts to find the root directory of the current Git repository
// 2. Returns an empty string if the Git repository root cannot be found
// 3. Constructs the full absolute path to the Terragrunt directory
// The returned path is constructed by joining the Git repository root with the relative
// path to the Terragrunt directory as defined by GetInfraTerragruntDirPath().
func GetInfraTerragruntDirPathAbsolute() (string, error) {
	gitRepoRootPath, err := GetGitRepoRoot()

	if err != nil {
		return "", fmt.Errorf("failed to get the infrastructure terragrunt directory path: %w", err)
	}

	return filepath.Join(gitRepoRootPath, GetInfraTerragruntDirPath()), nil
}

// GetInfraTerragruntDirPathAbsoluteWithGitRepoRoot constructs and returns the absolute filesystem path
// to the Terragrunt directory based on the provided Git repository root path.
//
// This function takes a string representing the absolute path to the root of a Git repository
// and appends the relative path to the Terragrunt directory, as defined by the constant
// GetInfraTerragruntDirPath(). It is crucial to ensure that the provided gitRepoRootPath is a valid
// directory path that corresponds to the root of a Git repository.
//
// Parameters:
//   - gitRepoRootPath: A string representing the absolute path to the Git repository root.
//
// Returns:
//   - A string containing the full absolute path to the Terragrunt directory.
func GetInfraTerragruntDirPathAbsoluteWithGitRepoRoot(gitRepoRootPath string) string {
	return filepath.Join(gitRepoRootPath, GetInfraTerragruntDirPath())
}

// GetInfraTargetEnvFilePath returns the filesystem path to the target environment configuration file
// located in the Terragrunt directory. This typically represents a specific or target environment configuration.
func GetInfraTargetEnvFilePath() string {
	return filepath.Join(GetInfraTerragruntDirPath(), "target.yaml")
}

// GetBaseEnvFilePath constructs and returns the absolute filesystem path to the base environment
// configuration file located within the Terragrunt directory. This file typically contains the
// default environment configuration settings used by Terragrunt.
//
// The function performs the following steps:
//  1. Calls GetInfraTerragruntDirPath() to retrieve the absolute path to the Terragrunt directory.
//  2. Joins the Terragrunt directory path with the default environment configuration filename,
//     defined by EnvCfgBaseFilenameDefault, to create the full path to the base environment file.
//
// Returns:
// - A string representing the absolute path to the base environment configuration file.
//
// Note: The returned path may not be valid if the Terragrunt directory does not exist or if
// the filename is incorrect. It is the caller's responsibility to handle any potential errors
// related to file existence or accessibility.
func GetBaseEnvFilePath() string {
	return filepath.Join(GetInfraTerragruntDirPath(), EnvsDirectory, EnvCfgBaseFilenameDefault)
}

// GetBaseEnvFilePathAbsolute constructs and returns the absolute filesystem path to the base
// environment configuration file located within the Terragrunt directory. This function first
// retrieves the root path of the Git repository and then combines it with the relative path to
// the base environment configuration file, as defined by the GetBaseEnvFilePath function.
//
// This function is essential for obtaining the full path to the base environment file, which
// is typically used for default environment configuration settings in Terragrunt.
//
// Returns:
//   - A string containing the absolute path to the base environment configuration file.
//   - An error if the Git repository root path cannot be determined or if any other issue occurs
//     during the retrieval process.
//
// Note: The returned path may not be valid if the Terragrunt directory does not exist or if
// the filename is incorrect. It is the caller's responsibility to handle any potential errors
// related to file existence or accessibility.
func GetBaseEnvFilePathAbsolute() (string, error) {
	gitRepoRootPath, err := GetGitRepoRoot()

	if err != nil {
		return "", fmt.Errorf("failed to get the base environment file path: %w", err)
	}

	return filepath.Join(gitRepoRootPath, GetBaseEnvFilePath()), nil
}

// GetConfigPathForStack returns the absolute filesystem path for a given stack
//
// This function constructs the path to a specific stack's configuration directory
// by joining the Git repository root with the stack configuration from the provided EnvConfig.
//
// Parameters:
//   - config: The environment configuration containing stack information
//   - stackName: The name of the stack (e.g., "landing-zone")
//
// Returns:
//   - A string containing the absolute path to the stack's configuration directory
//   - An error if the Git repository root cannot be determined or stack is not found
//
// Example:
//
//	path, err := GetConfigPathForStack(envConfig, "landing-zone")
//	// path might be: /path/to/repo/infra/terragrunt/stack-landing-zone
func GetConfigPathForStack(config *EnvConfig, stackName string) (string, error) {
	// Validate input
	if config == nil {
		return "", fmt.Errorf("configuration is nil")
	}

	// Find the stack in the configuration
	var targetStack *StackConfig
	for _, stack := range config.Stacks {
		if stack.Name == stackName {
			targetStack = &stack
			break
		}
	}

	if targetStack == nil {
		return "", fmt.Errorf("stack '%s' not found in configuration", stackName)
	}

	// Get Git repository root
	gitRepoRoot, err := GetGitRepoRoot()
	if err != nil {
		return "", fmt.Errorf("failed to get config path for stack '%s': %w", stackName, err)
	}

	// Construct stack path using Git repo root and Terragrunt directory
	return filepath.Join(gitRepoRoot, GetInfraTerragruntDirPath(), fmt.Sprintf("stack-%s", stackName)), nil
}

// GetConfigPathForLayer returns the absolute filesystem path for a given layer within a stack
//
// This function constructs the path to a specific layer's configuration directory
// by joining the stack's configuration path with the layer name from the provided EnvConfig.
//
// Parameters:
//   - config: The environment configuration containing stack and layer information
//   - stackName: The name of the stack (e.g., "landing-zone")
//   - layerName: The name of the layer (e.g., "dns")
//
// Returns:
//   - A string containing the absolute path to the layer's configuration directory
//   - An error if the stack or layer cannot be found
//
// Example:
//
//	path, err := GetConfigPathForLayer(envConfig, "landing-zone", "dns")
//	// path might be: /path/to/repo/infra/terragrunt/stack-landing-zone/dns
func GetConfigPathForLayer(config *EnvConfig, stackName, layerName string) (string, error) {
	// Get stack path first
	stackPath, err := GetConfigPathForStack(config, stackName)
	if err != nil {
		return "", fmt.Errorf("failed to get config path for layer '%s' in stack '%s': %w", layerName, stackName, err)
	}

	// Find the layer in the stack configuration
	var targetLayer *LayerConfig
	for _, stack := range config.Stacks {
		if stack.Name == stackName {
			for _, layer := range stack.Layers {
				if layer.Name == layerName {
					targetLayer = &layer
					break
				}
			}
			break
		}
	}

	if targetLayer == nil {
		return "", fmt.Errorf("layer '%s' not found in stack '%s'", layerName, stackName)
	}

	return filepath.Join(stackPath, layerName), nil
}

// GetConfigPathForComponent returns the absolute filesystem path for a given component within a layer
//
// This function constructs the path to a specific component's configuration directory
// by joining the layer's configuration path with the component name from the provided EnvConfig.
//
// Parameters:
//   - config: The environment configuration containing stack, layer, and component information
//   - stackName: The name of the stack (e.g., "landing-zone")
//   - layerName: The name of the layer (e.g., "dns")
//   - componentName: The name of the component (e.g., "dns-zone")
//
// Returns:
//   - A string containing the absolute path to the component's configuration directory
//   - An error if the stack, layer, or component cannot be found
//
// Example:
//
//	path, err := GetConfigPathForComponent(envConfig, "landing-zone", "dns", "dns-zone")
//	// path might be: /path/to/repo/infra/terragrunt/stack-landing-zone/dns/dns-zone
func GetConfigPathForComponent(config *EnvConfig, stackName, layerName, componentName string) (string, error) {
	// Get layer path first
	layerPath, err := GetConfigPathForLayer(config, stackName, layerName)
	if err != nil {
		return "", fmt.Errorf("failed to get config path for component '%s' in layer '%s' of stack '%s': %w",
			componentName, layerName, stackName, err)
	}

	// Find the component in the layer configuration
	var targetComponent *ComponentConfig
	for _, stack := range config.Stacks {
		if stack.Name == stackName {
			for _, layer := range stack.Layers {
				if layer.Name == layerName {
					for _, component := range layer.Components {
						if component.Name == componentName {
							targetComponent = &component
							break
						}
					}
					break
				}
			}
			break
		}
	}

	if targetComponent == nil {
		return "", fmt.Errorf("component '%s' not found in layer '%s' of stack '%s'",
			componentName, layerName, stackName)
	}

	return filepath.Join(layerPath, componentName), nil
}
