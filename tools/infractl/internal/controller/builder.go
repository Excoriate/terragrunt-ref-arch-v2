package controller

import (
	"fmt"

	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/internal/cfg"
)

// buildBaseEnvConfig constructs the base environment configuration by retrieving
// the path to the base environment configuration file and loading its contents.
//
// This function utilizes the GetBaseEnvFilePath function to obtain the path
// of the base configuration file and then calls GetInfraEnvConfigFromFile
// to read and parse the configuration from that file. If any errors occur
// during this process, an error is returned with a detailed message indicating
// the failure reason, including the path that was attempted to be read.
//
// Returns:
//   - A pointer to the base EnvConfig structure, which contains the loaded
//     environment configuration.
//   - An error if the loading process fails, providing details about the
//     specific failure. If successful, the error will be nil.
func (c *Client) buildBaseEnvConfig() (*cfg.EnvConfig, error) {
	baseCfgPath, err := cfg.GetBaseEnvFilePathAbsolute()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve the absolute path for the base environment configuration file: %w", err)
	}

	baseCfg, err := cfg.GetInfraEnvConfigFromFile(baseCfgPath)

	if err != nil {
		return nil, fmt.Errorf("failed to build base environment configuration from path %s: %w", baseCfgPath, err)
	}

	return baseCfg, nil
}

// BuildTargetEnvConfig constructs the target environment configuration by resolving
// the configuration file path for the specified target environment and loading its
// contents from the file. This function is a part of the Client struct and is
// responsible for ensuring that the target environment configuration is correctly
// retrieved and parsed.
//
// Parameters:
//   - targetEnv: A string representing the name of the target environment for which
//     the configuration is to be built. This should correspond to a valid environment
//     name that has an associated configuration file.
//
// Returns:
//   - A pointer to the EnvConfig structure containing the loaded target environment
//     configuration.
//   - An error if the loading process fails, providing details about the specific
//     failure. If successful, the error will be nil.
//
// The function first attempts to resolve the file path of the target environment
// configuration using the ResolveEnvConfigFilepathByEnvName method. If this fails,
// it returns an error indicating the failure reason along with the attempted path.
// If the path is successfully resolved, it then attempts to load the configuration
// from the file using GetInfraEnvConfigFromFile. If this operation fails, it returns
// an error indicating the failure to load the configuration.
func (c *Client) BuildTargetEnvConfig(targetEnv string) (*cfg.EnvConfig, error) {
	targetEnvCfgPath, err := c.ResolveEnvConfigFilepathByEnvName(targetEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to build target environment configuration from path %s: %w", targetEnvCfgPath, err)
	}

	targetCfg, err := cfg.GetInfraEnvConfigFromFile(targetEnvCfgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load target environment configuration: %w", err)
	}

	return targetCfg, nil
}
