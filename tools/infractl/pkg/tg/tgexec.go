package tg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// TerragruntOptions represents comprehensive configuration options for Terragrunt commands
type TerragruntOptions struct {
	// Core command configuration
	WorkingDir     string
	ConfigPath     string
	Command        string
	NonInteractive bool
	AutoApprove    bool

	// Execution control
	Parallelism int
	IncludeDirs []string
	ExcludeDirs []string

	// Advanced options
	NoColor            bool
	Debug              bool
	StrictMode         bool
	IgnoreDependencies bool

	// Specific command options
	Target      string
	Replace     string
	Destroy     bool
	RefreshOnly bool

	// Additional arguments for maximum flexibility
	AdditionalArgs []string

	// Output and logging
	JsonOutputDir string
	OutputDir     string
}

// buildTerragruntCommand constructs a comprehensive Terragrunt command
func buildTerragruntCommand(opts TerragruntOptions) *exec.Cmd {
	args := []string{}

	// Add command
	if opts.Command != "" {
		if strings.HasPrefix(opts.Command, "run-all") {
			args = append(args, "run-all")
			args = append(args, strings.TrimPrefix(opts.Command, "run-all "))
		} else {
			args = append(args, opts.Command)
		}
	}

	// Global options
	if opts.ConfigPath != "" {
		args = append(args, "--terragrunt-config", opts.ConfigPath)
	}

	if opts.NonInteractive {
		args = append(args, "--terragrunt-non-interactive")
	}

	if opts.AutoApprove {
		args = append(args, "-auto-approve")
	}

	if opts.NoColor {
		args = append(args, "--terragrunt-no-color")
	}

	if opts.Debug {
		args = append(args, "--terragrunt-debug")
	}

	if opts.StrictMode {
		args = append(args, "--strict-mode")
	}

	if opts.IgnoreDependencies {
		args = append(args, "--terragrunt-ignore-dependency-order")
	}

	// Parallelism for run-all commands
	if opts.Parallelism > 0 && strings.Contains(opts.Command, "run-all") {
		args = append(args, fmt.Sprintf("--terragrunt-parallelism=%d", opts.Parallelism))
	}

	// Include/Exclude directories
	for _, dir := range opts.IncludeDirs {
		absDir, err := filepath.Abs(dir)
		if err == nil {
			args = append(args, fmt.Sprintf("--terragrunt-include-dir=%s", absDir))
		}
	}

	for _, dir := range opts.ExcludeDirs {
		absDir, err := filepath.Abs(dir)
		if err == nil {
			args = append(args, fmt.Sprintf("--terragrunt-exclude-dir=%s", absDir))
		}
	}

	// Specific command options
	if opts.Target != "" {
		args = append(args, "-target="+opts.Target)
	}

	if opts.Replace != "" {
		args = append(args, "-replace="+opts.Replace)
	}

	if opts.Destroy {
		args = append(args, "-destroy")
	}

	if opts.RefreshOnly {
		args = append(args, "-refresh-only")
	}

	// Output directories
	if opts.JsonOutputDir != "" {
		args = append(args, "--terragrunt-json-out-dir="+opts.JsonOutputDir)
	}

	if opts.OutputDir != "" {
		args = append(args, "--terragrunt-out-dir="+opts.OutputDir)
	}

	// Additional arguments for maximum flexibility
	args = append(args, opts.AdditionalArgs...)

	workingDir := opts.WorkingDir
	if workingDir == "" {
		workingDir = "."
	}

	cmd := exec.Command("terragrunt", args...)
	cmd.Dir = workingDir

	return cmd
}

// StreamOptions defines configuration for command streaming
type StreamOptions struct {
	// Inherit all TerragruntOptions
	TerragruntOptions

	// Optional custom output writers (defaults to os.Stdout and os.Stderr)
	OutWriter io.Writer
	ErrWriter io.Writer
}

// streamCommand is a generic method to stream command output
func streamCommand(opts StreamOptions) error {
	// Build the base Terragrunt command
	cmd := buildTerragruntCommand(opts.TerragruntOptions)
	if cmd == nil {
		return fmt.Errorf("failed to build terragrunt command: invalid configuration")
	}

	// Use default writers if not provided
	outWriter := opts.OutWriter
	if outWriter == nil {
		outWriter = os.Stdout
	}
	errWriter := opts.ErrWriter
	if errWriter == nil {
		errWriter = os.Stderr
	}

	// Create pipes for stdout and stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start terragrunt command: %w", err)
	}

	// Create scanners
	stdoutScanner := bufio.NewScanner(stdout)
	stderrScanner := bufio.NewScanner(stderr)

	// Stream output
	go func() {
		for stdoutScanner.Scan() {
			fmt.Fprintln(outWriter, stdoutScanner.Text())
		}
	}()

	go func() {
		for stderrScanner.Scan() {
			fmt.Fprintln(errWriter, stderrScanner.Text())
		}
	}()

	// Wait for command completion
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("%s command failed: %w", opts.Command, err)
	}

	return nil
}

// Plan runs terragrunt plan with streaming output
func Plan(opts TerragruntOptions) error {
	streamOpts := StreamOptions{
		TerragruntOptions: opts,
	}
	streamOpts.Command = "plan"
	return streamCommand(streamOpts)
}

// Apply runs terragrunt apply with streaming output
func Apply(opts TerragruntOptions) error {
	streamOpts := StreamOptions{
		TerragruntOptions: opts,
	}
	streamOpts.Command = "apply"
	return streamCommand(streamOpts)
}

// Destroy runs terragrunt destroy with streaming output
func Destroy(opts TerragruntOptions) error {
	streamOpts := StreamOptions{
		TerragruntOptions: opts,
	}
	streamOpts.Command = "destroy"
	return streamCommand(streamOpts)
}

// RunAll executes terragrunt run-all with specified command and streaming output
func RunAll(command string, opts TerragruntOptions) error {
	streamOpts := StreamOptions{
		TerragruntOptions: opts,
	}
	streamOpts.Command = "run-all " + command
	return streamCommand(streamOpts)
}

// RunAllPlan is a convenience method for run-all plan
func RunAllPlan(opts TerragruntOptions) error {
	return RunAll("plan", opts)
}

// RunAllApply is a convenience method for run-all apply
func RunAllApply(opts TerragruntOptions) error {
	return RunAll("apply", opts)
}

// RunAllDestroy is a convenience method for run-all destroy
func RunAllDestroy(opts TerragruntOptions) error {
	return RunAll("destroy", opts)
}
