package controller

import (
	"fmt"
	"path/filepath"

	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/internal/cfg"
	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/pkg/tg"
	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/pkg/utils"
)

type Tg struct {
	cfgCompiled         *cfg.EnvConfig
	cfgCompiledJSONPath string
}

type TgRunnerStackOptions struct {
	StackName     string
	LayerName     string
	ComponentName string
}

type TgRunner interface {
	Plan(stackOpts TgRunnerStackOptions, tgArgs ...string) error
}

func NewTgRunner(cfgCompiled *cfg.EnvConfig, cfgCompiledJSONPath string) (*Tg, error) {
	if cfgCompiled == nil {
		return nil, fmt.Errorf("configuration for the target environment (cfgCompiled) must not be nil; ensure that the environment is properly initialized")
	}

	if cfgCompiledJSONPath == "" {
		return nil, fmt.Errorf("the path to the compiled JSON configuration (cfgCompiledJSONPath) cannot be empty; please provide a valid file path")
	}

	return &Tg{
		cfgCompiled:         cfgCompiled,
		cfgCompiledJSONPath: cfgCompiledJSONPath,
	}, nil
}

func (t *Tg) setTgEnvVar() error {
	cfg.SetTransmitterTgEnvVar(t.cfgCompiledJSONPath)

	return nil
}

// getWorkdir constructs the working directory path for the specified stack, layer, and component.
// It retrieves the absolute path to the infrastructure Terragrunt directory and builds the path
// based on the provided stack options. It also validates the existence of the directory and
// checks for the presence of the 'terragrunt.hcl' file if a component is specified.
//
// Parameters:
//
//	stackOpts: TgRunnerStackOptions containing the stack name, layer name, and component name.
//
// Returns:
//
//	A string representing the absolute path to the working directory if successful.
//	An error if any issues occur during path resolution or validation, providing context
//	about the failure.
func (t *Tg) getWorkdir(stackOpts TgRunnerStackOptions) (string, error) {
	terragruntDirPath, err := cfg.GetInfraTerragruntDirPathAbsolute()
	if err != nil {
		return "", fmt.Errorf("failed to determine the absolute path to the infrastructure Terragrunt directory: %w", err)
	}

	// Construct the base path for Terragrunt configurations, starting with the Terragrunt directory and stack name.
	basePath := filepath.Join(terragruntDirPath, stackOpts.StackName)

	// Append the layer name to the base path if it's provided.
	if stackOpts.LayerName != "" {
		basePath = filepath.Join(basePath, stackOpts.LayerName)
	}

	// Append the component name to the base path if it's provided.
	if stackOpts.ComponentName != "" {
		basePath = filepath.Join(basePath, stackOpts.ComponentName)
	}

	// Resolve the constructed path to an absolute file path.
	workdirPath, err := filepath.Abs(basePath)
	if err != nil {
		return "", fmt.Errorf("failed to resolve the absolute path for '%s': %w", basePath, err)
	}

	// Verify that the resolved path corresponds to an existing directory.
	if err := utils.DirExists(workdirPath); err != nil {
		return "", fmt.Errorf("the resolved path '%s' does not correspond to an existing directory: %w", workdirPath, err)
	}

	// If a component name is specified, check for the presence of 'terragrunt.hcl' within the component directory.
	if stackOpts.ComponentName != "" {
		terragruntHclPath := filepath.Join(workdirPath, "terragrunt.hcl")
		if err := utils.FileExists(terragruntHclPath); err != nil {
			return "", fmt.Errorf("the 'terragrunt.hcl' configuration file is missing in the component directory '%s': %w", workdirPath, err)
		}
	}

	return workdirPath, nil
}

// Plan wraps the Terragrunt plan command with hierarchical validation
func (t *Tg) Plan(stackOpts TgRunnerStackOptions, tgArgs ...string) error {
	// Set Terragrunt environment variable before execution
	if err := t.setTgEnvVar(); err != nil {
		return fmt.Errorf("failed to set Terragrunt environment variable: %w", err)
	}

	// Get the workdir for the stack, layer, or component
	workdir, workdirErr := t.getWorkdir(stackOpts)
	if workdirErr != nil {
		return fmt.Errorf("failed to get workdir: %w", workdirErr)
	}

	fmt.Println("Running Terragrunt plan command in workdir:", workdir)

	// Prepare Terragrunt options
	planOpts := tg.TerragruntOptions{
		WorkingDir:     workdir,
		Command:        "plan",
		NonInteractive: true,
		AdditionalArgs: tgArgs,
	}

	// Execute Terragrunt plan with streaming output
	return tg.Plan(planOpts)
}
