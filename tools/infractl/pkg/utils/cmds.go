package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// ExecuteCommand runs a generic shell command with flexible arguments
func ExecuteCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)

	// Capture both stdout and stderr
	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Set environment variables from current process
	cmd.Env = os.Environ()

	err := cmd.Run()

	output := outBuffer.String()
	errorOutput := errBuffer.String()

	if err != nil {
		fullErrorMessage := fmt.Sprintf("Command execution failed: %v\nStdout: %s\nStderr: %s",
			err, output, errorOutput)
		return output, fmt.Errorf(fullErrorMessage)
	}

	return output, nil
}

// ExecuteBinaryCommand runs specific binary commands with predefined configurations
func ExecuteBinaryCommand(binaryName string, workingDir string, args ...string) (string, error) {
	// Validate binary exists
	_, err := exec.LookPath(binaryName)
	if err != nil {
		return "", fmt.Errorf("binary %s not found in system path", binaryName)
	}

	cmd := exec.Command(binaryName, args...)

	// Set working directory if provided
	if workingDir != "" {
		cmd.Dir = workingDir
	}

	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Inherit environment variables
	cmd.Env = os.Environ()

	err = cmd.Run()

	output := outBuffer.String()
	errorOutput := errBuffer.String()

	if err != nil {
		fullErrorMessage := fmt.Sprintf("Binary command execution failed: %v\nStdout: %s\nStderr: %s",
			err, output, errorOutput)
		return output, fmt.Errorf(fullErrorMessage)
	}

	return output, nil
}
