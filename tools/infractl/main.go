package main

import (
	"fmt"
	"os"

	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/internal/controller"
	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/internal/tui"
	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/pkg/logger"
	"github.com/alecthomas/kong"
)

var CLI struct {
	Plan     PlanCmd     `cmd:"" help:"Plan infrastructure changes"`
	Apply    ApplyCmd    `cmd:"" help:"Apply infrastructure changes"`
	Destroy  DestroyCmd  `cmd:"" help:"Destroy infrastructure"`
	Validate ValidateCmd `cmd:"" help:"Validate secrets and configurations - It does not compile, just pre-validate the configuration"`
}

type GenerateCmd struct {
	Stack            string `help:"Name of the stack to execute" required:"true"`
	Component        string `help:"Optional name of the component to execute" required:"true"`
	Layer            string `help:"Optional name of the layer to execute" required:"true"`
	Base             string `help:"Name of the base environment configuration. Defaults to 'base', which corresponds to _ENVS/base.yaml" default:"base" optional:"true"`
	TargetEnv        string `help:"Name of the target environment. E.g.: local, staging, production. If 'local' is passed, it means that there is a target configuration in _ENVS/local.yaml" required:""`
	OverrideJSONName string `help:"Optional name of the JSON file to override the default name of the JSON file. E.g.: 'my_custom_name.json'" optional:"true"`
}

type PlanCmd struct {
	Stack            string `help:"Name of the stack to execute" required:"true"`
	Component        string `help:"Optional name of the component to execute" required:"true"`
	Layer            string `help:"Optional name of the layer to execute" required:"true"`
	Base             string `help:"Name of the base environment configuration. Defaults to 'base', which corresponds to _ENVS/base.yaml" default:"base" optional:"true"`
	TargetEnv        string `help:"Name of the target environment. E.g.: local, staging, production. If 'local' is passed, it means that there is a target configuration in _ENVS/local.yaml" required:""`
	OverrideJSONName string `help:"Optional name of the JSON file to override the default name of the JSON file. E.g.: 'my_custom_name.json'" optional:"true"`
}

type ApplyCmd struct {
	Env              string `help:"Environment configuration file location (ENVS/<environment>.yaml)" required:"true"`
	Stack            string `help:"Name of the stack to execute" required:"true"`
	Layer            string `help:"Optional name of the layer to execute" required:"true"`
	Component        string `help:"Optional name of the component to execute" required:"true"`
	OverrideJSONName string `help:"Optional name of the JSON file to override the default name of the JSON file. E.g.: 'my_custom_name.json'" optional:"true"`
}

type DestroyCmd struct {
	Env              string `help:"Environment configuration file location (ENVS/<environment>.yaml)" required:"true"`
	Stack            string `help:"Name of the stack to execute" required:"true"`
	Layer            string `help:"Optional name of the layer to execute" required:"true"`
	Component        string `help:"Optional name of the component to execute" required:"true"`
	OverrideJSONName string `help:"Optional name of the JSON file to override the default name of the JSON file. E.g.: 'my_custom_name.json'" optional:"true"`
}

type ValidateCmd struct {
	Base      string `help:"Name of the base environment configuration. Defaults to 'base', which corresponds to _ENVS/base.yaml" default:"base" optional:"true"`
	TargetEnv string `help:"Name of the target environment. E.g.: local, staging, production. If 'local' is passed, it means that there is a target configuration in _ENVS/local.yaml" required:""`
}

func (v *ValidateCmd) Run() error {
	log := logger.DefaultLogger()

	targetEnvParam := v.TargetEnv
	baseEnvParam := v.Base

	// Log the actual values of Base and TargetEnv for clarity
	log.Info(fmt.Sprintf("🔍 Target Environment: %s", targetEnvParam))
	log.Info(fmt.Sprintf("🔍 Base Environment: %s", baseEnvParam))

	// Create a new infractl client
	log.Info("🔧 Initializing infractl client...")
	ic, clientErr := controller.NewClient(baseEnvParam, targetEnvParam)
	if clientErr != nil {
		return fmt.Errorf("❌ Error: Unable to create infractl client: %w", clientErr)
	}

	// Initialise the infractl client
	log.Info("🛠️ Initializing the infractl client...")
	if err := ic.Initialise(); err != nil {
		return fmt.Errorf("❌ Error: Failed to initialize infractl client: %w", err)
	}

	// Run sanity checks
	log.Info("🛠️ Running sanity check... 🧐")
	if err := ic.RunSanityCheck(targetEnvParam); err != nil {
		return fmt.Errorf("❌ Error: Sanity check failed: %w", err)
	}

	log.Info("✅ Sanity check passed successfully! 🎉")

	// Compile the target environment configuration
	log.Info("🔍 Compiling the target environment configuration...")
	if _, err := ic.Compile(v.TargetEnv); err != nil {
		return fmt.Errorf("❌ Error: Failed to compile target environment configuration: %w", err)
	}

	log.Info("✅ Target environment configuration compiled successfully!")
	log.Info("🎉 All checks completed successfully! 🎉")

	return nil
}

func (p *PlanCmd) Run() error {
	log := logger.DefaultLogger()

	// Log the input parameters for traceability
	log.Info(fmt.Sprintf("🌍 Initiating infrastructure planning for environment: %s", p.TargetEnv))
	log.Info(fmt.Sprintf("🏗️ Targeting stack: %s", p.Stack))

	if p.Component != "" {
		log.Info(fmt.Sprintf("🧩 Focusing on specific component: %s", p.Component))
	}

	// Checking the stack hierarchy consistency
	if err := controller.IsStackHierarchyConsistent(p.Stack, p.Layer, p.Component); err != nil {
		return fmt.Errorf("❌ Error: Stack hierarchy is inconsistent: %w", err)
	}

	// Create and initialize the infractl client
	log.Info("🔧 Setting up the infrastructure client...")
	ic, clientErr := controller.NewClient(p.Base, p.TargetEnv)
	if clientErr != nil {
		return fmt.Errorf("❌ Error: Unable to create infractl client: %w", clientErr)
	}

	if err := ic.Initialise(); err != nil {
		return fmt.Errorf("❌ Error: Failed to initialize infractl client: %w", err)
	}

	// Run sanity checks to ensure system readiness
	log.Info("🕵️ Conducting initial system sanity check...")
	if err := ic.RunSanityCheck(p.TargetEnv); err != nil {
		return fmt.Errorf("❌ Error: Sanity check failed: %w", err)
	}

	log.Info("✅ Sanity check completed successfully!")

	// Compile the target environment configuration
	log.Info("🔍 Compiling the target environment configuration...")
	compiledConfig, compileErr := ic.Compile(p.TargetEnv)
	if compileErr != nil {
		return fmt.Errorf("❌ Error: Compilation of target environment configuration failed: %w", compileErr)
	}

	log.Info("✅ Target environment configuration compiled successfully!")

	// Validating the infrastructure hierarchy
	log.Info("🔍 Validating the infrastructure hierarchy...")
	if err := ic.ValidateInfrastructureHierarchy(compiledConfig, p.Stack, p.Layer, p.Component); err != nil {
		return fmt.Errorf("❌ Error: Infrastructure hierarchy validation failed: %w", err)
	}

	log.Info("✅ Infrastructure hierarchy validated successfully!")

	// Convert the compiled configuration to JSON format
	log.Info("🔄 Transforming compiled configuration into JSON format...")
	compiledEnvConfigInJSON, err := ic.EnvCfgCompiledToJSON(compiledConfig)
	if err != nil {
		return fmt.Errorf("❌ Error: Conversion to JSON format failed: %w", err)
	}

	log.Info("✅ Compiled environment configuration converted to JSON format successfully!")

	// Store the JSON file in the cache directory
	log.Info("💾 Storing compiled environment configuration in cache directory...")
	envConfigFilepathInCacheDir, err := ic.CreateCachedEnvCfgJSONFile(p.TargetEnv, compiledEnvConfigInJSON, p.OverrideJSONName)
	if err != nil {
		return fmt.Errorf("❌ Error: Unable to create cached environment configuration file: %w", err)
	}

	log.Info(fmt.Sprintf("💾 Compiled environment configuration saved in JSON format at: %s", envConfigFilepathInCacheDir))

	// Running Tg using the InfraRunner
	log.Info("🚀 Running Terragrunt plan command...")
	tgRunner, tgRunnerErr := controller.NewTgRunner(compiledConfig, envConfigFilepathInCacheDir)
	if tgRunnerErr != nil {
		return fmt.Errorf("❌ Error: Unable to create infractl client: %w", tgRunnerErr)
	}

	if err := tgRunner.Plan(controller.TgRunnerStackOptions{
		StackName:     p.Stack,
		LayerName:     p.Layer,
		ComponentName: p.Component,
	}); err != nil {
		return fmt.Errorf("❌ Error: Failed to run Terragrunt plan command: %w", err)
	}

	log.Info("✅ Terragrunt plan command executed successfully!")

	return nil
}

func main() {
	fmt.Println(tui.GetBanner())

	ctx := kong.Parse(&CLI,
		kong.Name("infra"),
		kong.Description("CLI tool to facilitate the IaaC configuration using Terragrunt"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
	)

	err := ctx.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
