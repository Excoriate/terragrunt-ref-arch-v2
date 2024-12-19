package controller

import (
	"fmt"
	"path/filepath"

	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/internal/cfg"
	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/pkg/envars"
	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/pkg/utils"
)

type RepoPaths struct {
	GitRepoRoot string
	Cache       string
	Terragrunt  string
	EnvsConfig  string
}

type Client struct {
	Paths RepoPaths
}

// NewClient creates and initializes a new InfraController client with the specified base and override environment configuration file paths.
//
// The function performs the following key steps:
// 1. Finds the root directory of the current Git repository
// 2. Resolves the absolute path for environment configuration files
// 3. Creates a new Client instance with the provided and resolved paths
//
// Parameters:
//   - baseFilePath: Path to the base environment configuration file
//   - overrideFilePath: Path to the environment override configuration file
//
// Returns:
//   - A pointer to the newly created Client instance
//   - An error if any path resolution fails (Git root or environment config paths)
func NewClient(baseFilePath, targetEnvFilePath string) (*Client, error) {
	gitRepoRoot, err := utils.FindGitRepoRoot()

	if err != nil {
		return nil, fmt.Errorf("failed to create a infractl client due to git repository root path resolution failure: %w", err)
	}

	// InfraCTL required paths
	envsFilesPath := cfg.GetEnvConfigFilesPathAbsoluteWithGitRepoRoot(gitRepoRoot)
	cachePath := cfg.GetInfraCacheDirPathAbsoluteWithGitRepoRoot(gitRepoRoot)
	terragruntDirPath := cfg.GetInfraTerragruntDirPathAbsoluteWithGitRepoRoot(gitRepoRoot)

	// Ensure the base and target environment files have the correct extension,
	// since it's expected to receive targetenv, and not necessarily target.yaml
	// baseEnvFileWithExtension := utils.ForceExtensionForFilepath(baseFilePath, ".yaml")
	// targetEnvFileWithExtension := utils.ForceExtensionForFilepath(targetEnvFilePath, ".yaml")

	// baseEnvFilePathResolved := filepath.Join(envsFilesPath, baseEnvFileWithExtension)
	// targetEnvFilePathResolved := filepath.Join(envsFilesPath, targetEnvFileWithExtension)

	client := &Client{
		Paths: RepoPaths{
			GitRepoRoot: gitRepoRoot,
			Cache:       cachePath,
			Terragrunt:  terragruntDirPath,
			EnvsConfig:  envsFilesPath,
		},
	}

	return client, nil
}

// ResolveEnvConfigFilepathByEnvName constructs the absolute file path for the environment configuration file
// based on the provided environment name. It ensures that the environment name is valid and appends the
// appropriate file extension before resolving the full path.
//
// Parameters:
//   - envName: A string representing the name of the environment for which the configuration file path is to be resolved.
//     This should be a valid environment name and cannot be empty.
//
// Returns:
//   - A string containing the absolute path to the environment configuration file with the ".yaml" extension.
//   - An error if the environment name is empty or if any other issues arise during path resolution.
//
// Example usage:
//
//	filepath, err := client.ResolveEnvConfigFilepathByEnvName("development")
//	if err != nil {
//	    // handle error
//	}
//	// use filepath as needed
func (c *Client) ResolveEnvConfigFilepathByEnvName(envName string) (string, error) {
	if envName == "" {
		return "", fmt.Errorf("environment name is not set, please provide a valid environment name")
	}

	// Force the environment file name to have a ".yaml" extension
	envFileWithExtension := utils.ForceExtensionForFilepath(envName, ".yaml")
	// Construct the full path to the environment configuration file
	envFilePathResolved := filepath.Join(c.Paths.EnvsConfig, envFileWithExtension)

	return envFilePathResolved, nil
}

// RunSanityCheck performs a series of validation checks to ensure that the environment
// is properly set up for operation. This method checks for the installation of Terragrunt,
// the existence of environment configuration files, and the validity of both the base
// and target environment configuration files.
//
// Parameters:
//   - targetEnv: A string representing the name of the target environment. This should
//     be a valid environment name that is used to validate the corresponding configuration file.
//
// Returns:
//   - An error if any of the checks fail, providing details about the specific failure.
//     If all checks pass, it returns nil, indicating that the environment is ready for use.
//
// Example usage:
//
//	err := client.RunSanityCheck("production")
//	if err != nil {
//	    // handle error
//	}
func (c *Client) RunSanityCheck(targetEnv string) error {
	// Check 1: Terragrunt is installed
	if err := isTerragruntInstalled(); err != nil {
		return fmt.Errorf("terragrunt installation check failed: %w", err)
	}

	// Check 2: Validate environment files exist
	if err := envConfigurationFilesExistInEnvsDir(c.Paths.EnvsConfig); err != nil {
		return fmt.Errorf("failed to validate if any environment configuration files exist in ENVS directory: %w", err)
	}

	// Check 3: Validate base environment file
	if err := IsBaseEnvConfigFileValid(c.Paths.EnvsConfig); err != nil {
		return fmt.Errorf("failed to validate base environment file: %w", err)
	}

	targetEnvCfgPath, err := c.ResolveEnvConfigFilepathByEnvName(targetEnv)

	if err != nil {
		return fmt.Errorf("failed to resolve target environment configuration file path: %w", err)
	}

	// Check 4: Validate target environment file
	if err := IsTargetEnvConfigFileValid(targetEnvCfgPath); err != nil {
		return fmt.Errorf("failed to validate target environment file: %w", err)
	}

	return nil
}

// Initialise sets up the infractl client by performing necessary initialization tasks.
// This includes creating a cache directory and adding entries to the .gitignore file
// to ensure that generated files are not tracked by version control.
//
// Returns:
//   - An error if any of the initialization steps fail, providing details about the
//     specific failure. If all steps are successful, it returns nil, indicating that
//     the client has been successfully initialized.
//
// Example usage:
//
//	err := client.Initialise()
//	if err != nil {
//	    // handle error
//	}
func (c *Client) Initialise() error {
	// Attempt to create the cache directory. If it fails, return an error with context.
	if _, err := cfg.CreateCacheDir(); err != nil {
		return fmt.Errorf("failed to initialise infractl client: %w", err)
	}

	// Attempt to add entries to the .gitignore file. If it fails, return an error with context.
	if err := cfg.AddEntriesToGitignore(); err != nil {
		return fmt.Errorf("failed to initialise infractl client: %w", err)
	}

	// Load from dotenv
	if err := envars.LoadDotenv(); err != nil {
		return fmt.Errorf("failed to initialise infractl client: %w", err)
	}

	// Clean the transmitter environment variables set previously
	c.CleanTransmitterEnvVars()

	// If both operations succeed, return nil indicating successful initialization.
	return nil
}

// CleanTransmitterEnvVars removes environment variables associated with the transmitter
// by utilizing the keys defined in the InfraCtlGitIgnoreEntries configuration. This
// function is particularly useful for ensuring that any sensitive or unnecessary
// environment variables are cleared from the environment, thereby maintaining a clean
// state for subsequent operations.
//
// Returns:
//   - An error if the cleaning process fails, providing details about the specific
//     failure. If the operation is successful, it returns nil, indicating that the
//     environment variables have been successfully cleaned.
//
// Example usage:
//
//	client.CleanTransmitterEnvVars()
func (c *Client) CleanTransmitterEnvVars() {
	// Attempt to clean environment variables by keys defined in the InfraCtlGitIgnoreEntries.
	// If an error occurs during this process, we can ignore it as it means the env vars weren't set.
	_ = envars.CleanEnvVarsByKeys(cfg.InfraCtlGitIgnoreEntries)
}

// CreateCachedEnvCfgJSONFile generates a unique filename for the specified target environment
// and creates a JSON configuration file in the infra cache directory.
//
// This method performs the following steps:
//  1. Generates a unique filename based on the provided target environment name.
//  2. Creates a new file in the infra cache directory using the generated filename and the
//     provided JSON configuration content.
//
// If any of these steps fail, an error is returned with a descriptive message indicating
// the failure point. If successful, the method returns the full file path of the created
// JSON configuration file.
//
// Parameters:
//   - targetEnv: A string representing the name of the target environment for which the
//     configuration file is being created. This should be a valid environment name.
//   - jsonCfg: A string containing the JSON configuration content to be written to the file.
//   - overrideJSONName: A string representing the name of the JSON file to override the default name of the JSON file.
//
// Returns:
//   - A string containing the full file path of the created JSON configuration file.
//   - An error if any step in the process fails, providing details about the specific failure.
func (c *Client) CreateCachedEnvCfgJSONFile(targetEnv, jsonCfg, overrideJSONName string) (string, error) {
	var jsonFilename string
	if overrideJSONName != "" {
		jsonFilename = overrideJSONName
		// for extension if it's not provided
		jsonFilename = utils.ForceExtensionForFilepath(jsonFilename, ".json")
	} else {
		conventionalFilename, err := cfg.GenerateUniqueEnvConfigFilename(targetEnv)

		if err != nil {
			return "", fmt.Errorf("failed to create cached environment configuration JSON file: %w", err)
		}

		jsonFilename = conventionalFilename
	}

	envCfgJSONFilepath, creationErr := cfg.CreateFileInInfraCacheDir(jsonFilename, jsonCfg)

	if creationErr != nil {
		return "", fmt.Errorf("failed to create cached environment configuration JSON file: %w", creationErr)
	}

	return envCfgJSONFilepath, nil
}
