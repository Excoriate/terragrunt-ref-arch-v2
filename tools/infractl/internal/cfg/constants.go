package cfg

import "path/filepath"

const (
	// Directories
	EnvsDirectory = "_ENVS"
	InfraDir      = "infra"
	TerragruntDir = "terragrunt"
	CacheDir      = ".infractl-cache"
	// Defaults
	EnvCfgBaseFilenameDefault = "base.yaml"
	// Cache and Temporal
	TempDirPrefix = ".temp-"
)

var (
	// Cache path, expected: infra/.infractl-cache
	CacheDirPathRelative = filepath.Join(InfraDir, CacheDir)

	// Env configuration path, expected: infra/terragrunt/_ENVS
	EnvCfgDirPathRelative = filepath.Join(InfraDir, TerragruntDir, EnvsDirectory)

	// gitignore entries for the infrastructure cache directory
	InfraCtlGitIgnoreEntries = []string{
		".infractl-cache/",               // Ignore the cache directory
		".infractl-cache/*",              // Ignore contents of cache directory
		"infra/.infractl-cache/",         // Fully qualified path for infra cache
		"infra/.infractl-cache/*",        // Contents of infra cache
		"tools/infractl/target/infractl", // Specific binary in target directory
		"tools/infractl/infractl",        // Binary in infractl directory
	}
)
