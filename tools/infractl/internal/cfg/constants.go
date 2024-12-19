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
		".infractl-cache/",  // Standard directory ignore
		".infractl-cache",   // Exact match for file/directory
		".infractl-cache/*", // Contents of immediate directory
		"infractl",          // infractl binary
	}
)
