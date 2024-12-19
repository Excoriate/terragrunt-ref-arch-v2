package transformers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/internal/cfg"
)

// StacksTransformer provides functionality for managing and validating stacks
type StacksTransformer struct {
	EnvConfig *cfg.EnvConfig
	BasePath  string // Base path for Terragrunt configurations
}

// NewStacksTransformer creates a new instance of StacksTransformer
func NewStacksTransformer(cfg *cfg.EnvConfig, basePath string) *StacksTransformer {
	return &StacksTransformer{
		EnvConfig: cfg,
		BasePath:  basePath,
	}
}

// ValidateStacks validates all stacks in the configuration
// ValidateStacks validates all stacks in the configuration. It checks if each stack's directory exists
// and validates the layers within each stack. If any validation fails, it returns an error.
// ValidateStacks iterates through all stacks defined in the EnvConfig and validates each one.
// It checks if each stack's directory exists on the filesystem and then validates the layers within each stack.
//
// Returns:
//   - nil if all stacks are valid.
//   - An error if no stacks are found in the configuration, if a stack's directory does not exist,
//     or if there is an error validating the layers within a stack. The error provides specific
//     details about the validation failure, such as the stack name and directory path, or the
//     nature of the layer validation error.
func (t *StacksTransformer) ValidateStacks() error {
	// Check if there are no stacks configured and return an error if true.
	if len(t.EnvConfig.Stacks) == 0 {
		return fmt.Errorf("no stacks found in the configuration")
	}

	// Iterate through each stack in the configuration to validate.
	for _, stack := range t.EnvConfig.Stacks {
		// Construct the path to the stack directory.
		stackPath := filepath.Join(t.BasePath, stack.Name)

		// Check if the stack directory exists.
		if _, err := os.Stat(stackPath); os.IsNotExist(err) {
			return fmt.Errorf("stack '%s' exists in configuration but directory not found at: %s",
				stack.Name, stackPath)
		}

		// Validate the layers within the stack.
		if err := t.validateStackLayers(stack); err != nil {
			return fmt.Errorf("error validating layers for stack '%s': %w", stack.Name, err)
		}
	}

	// If all validations pass, return nil indicating success.
	return nil
}

// StackExists checks if a stack exists in the configuration.
// It iterates through the configured stacks to find a match for the given stackName.
// If a match is found, it returns true; otherwise, it returns false.
func (t *StacksTransformer) StackExists(stackName string) bool {
	// Check if the environment configuration is nil or if there are no stacks configured.
	if t.EnvConfig == nil || len(t.EnvConfig.Stacks) == 0 {
		return false // Return false if the configuration is empty or nil.
	}

	// Iterate through each stack in the configuration.
	for _, stack := range t.EnvConfig.Stacks {
		// Check if the current stack's name matches the given stackName.
		if stack.Name == stackName {
			return true // Return true if a match is found.
		}
	}

	// If no match is found after iterating through all stacks, return false.
	return false
}

// GetStack retrieves a stack by its name from the configuration.
// It first checks if the stack exists in the configuration using StackExists.
// If the stack exists, it iterates through the configured stacks to find the one matching the given name.
// If found, it returns a copy of the stack configuration to avoid returning a pointer to the loop variable.
// If the stack does not exist or is not found, it returns an error indicating the stack does not exist in the configuration.
func (t *StacksTransformer) GetStack(stackName string) (*cfg.StackConfig, error) {
	if !t.StackExists(stackName) {
		return nil, fmt.Errorf("stack '%s' does not exist in the configuration", stackName)
	}

	for _, stack := range t.EnvConfig.Stacks {
		if stack.Name == stackName {
			stackCopy := stack // Create a copy to avoid returning a pointer to the loop variable
			return &stackCopy, nil
		}
	}

	// This should never happen due to the StackExists check, but it's included for code completeness.
	return nil, fmt.Errorf("stack '%s' was found to exist but not retrieved", stackName)
}

// validateStackExists checks if a stack exists in the filesystem
func (t *StacksTransformer) validateStackExists(stack cfg.StackConfig) error {
	stackPath := filepath.Join(t.BasePath, stack.Name)

	if _, err := os.Stat(stackPath); os.IsNotExist(err) {
		return fmt.Errorf("stack '%s' does not exist in %s", stack.Name, stackPath)
	}

	return nil
}

// GetLayer retrieves a layer from a stack by name
// It first retrieves the stack by name from the configuration, then iterates through the stack's layers
// to find the one matching the given layer name. If found, it returns a copy of the layer
// configuration. If not found, it returns an error indicating the layer does not exist.
func (t *StacksTransformer) GetLayer(stackName, layerName string) (*cfg.LayerConfig, error) {
	stack, err := t.GetStack(stackName)
	if err != nil {
		return nil, fmt.Errorf("failed to get stack '%s': %w", stackName, err)
	}

	for _, layer := range stack.Layers {
		if layer.Name == layerName {
			layerCopy := layer // Create a copy to avoid returning a pointer to the loop variable
			return &layerCopy, nil
		}
	}

	return nil, fmt.Errorf("layer '%s' does not exist in stack '%s'", layerName, stackName)
}

// GetComponent retrieves a component from a stack's layer by name
// It first retrieves the layer by name from the stack, then iterates through the layer's components
// to find the one matching the given component name. If found, it returns a copy of the component
// configuration. If not found, it returns an error indicating the component does not exist.
func (t *StacksTransformer) GetComponent(stackName, layerName, componentName string) (*cfg.ComponentConfig, error) {
	layer, err := t.GetLayer(stackName, layerName)
	if err != nil {
		return nil, fmt.Errorf("failed to get layer '%s' from stack '%s': %w", layerName, stackName, err)
	}

	for _, component := range layer.Components {
		if component.Name == componentName {
			componentCopy := component // Create a copy to avoid returning a pointer to the loop variable
			return &componentCopy, nil
		}
	}

	return nil, fmt.Errorf("component '%s' does not exist in layer '%s' of stack '%s'",
		componentName, layerName, stackName)
}

func (t *StacksTransformer) validateStackLayers(stack cfg.StackConfig) error {
	stackPath := filepath.Join(t.BasePath, stack.Name)

	for _, layer := range stack.Layers {
		layerPath := filepath.Join(stackPath, layer.Name)

		// Check if layer directory exists
		if _, err := os.Stat(layerPath); os.IsNotExist(err) {
			return fmt.Errorf("layer '%s' exists in configuration but directory not found at: %s",
				layer.Name, layerPath)
		}

		// Validate components within the layer
		if err := t.validateLayerComponents(layerPath, layer); err != nil {
			return err
		}
	}

	return nil
}

func (t *StacksTransformer) validateLayerComponents(layerPath string, layer cfg.LayerConfig) error {
	for _, component := range layer.Components {
		componentPath := filepath.Join(layerPath, component.Name)

		// Check if component directory exists
		if _, err := os.Stat(componentPath); os.IsNotExist(err) {
			return fmt.Errorf("component '%s' exists in configuration but directory not found at: %s",
				component.Name, componentPath)
		}

		// Check for required Terragrunt configuration file
		componentHCL := filepath.Join(componentPath, "component.hcl")
		if _, err := os.Stat(componentHCL); os.IsNotExist(err) {
			return fmt.Errorf("missing component.hcl for component '%s' in layer '%s' at: %s",
				component.Name, layer.Name, componentHCL)
		}
	}

	return nil
}

// ValidateStackExists checks if a stack exists in both configuration and filesystem.
// It first verifies the stack's presence in the configuration, then checks for its corresponding directory
// in the filesystem.
//
// Parameters:
//   - stackName: The name of the stack to validate.
//
// Returns:
//   - An error if the stack does not exist in the configuration or if its directory is not found in the filesystem.
//     Returns nil if the stack is valid.
func (t *StacksTransformer) ValidateStackExists(stackName string) error {
	// Check configuration first
	if !t.StackExists(stackName) {
		return fmt.Errorf("stack '%s' does not exist in the configuration", stackName)
	}

	// Then check filesystem
	stack, err := t.GetStack(stackName)
	if err != nil {
		return fmt.Errorf("error retrieving stack '%s' from configuration: %w", stackName, err)
	}

	return t.validateStackExists(*stack)
}

// ValidateLayerExists checks if a layer exists in both configuration and filesystem
// It first verifies the layer's presence in the configuration, then checks for its corresponding directory
// in the filesystem.
//
// Parameters:
//   - stackName: The name of the stack.
//   - layerName: The name of the layer within the stack.
//
// Returns:
//   - An error if the layer does not exist in the configuration or if its directory is not found in the filesystem.
//     Returns nil if the layer is valid.
func (t *StacksTransformer) ValidateLayerExists(stackName, layerName string) error {
	// Retrieve the layer configuration to validate its existence
	layer, err := t.GetLayer(stackName, layerName)
	if err != nil {
		return fmt.Errorf("error retrieving layer '%s' from configuration: %w", layerName, err)
	}

	// Construct the path to the layer directory in the filesystem
	stackPath := filepath.Join(t.BasePath, stackName)
	layerPath := filepath.Join(stackPath, layer.Name)

	// Check if the layer directory exists in the filesystem
	if _, err := os.Stat(layerPath); os.IsNotExist(err) {
		return fmt.Errorf("layer '%s' exists in configuration but not in filesystem at path: %s",
			layerName, layerPath)
	}

	return nil
}

// ValidateComponentExists checks if a component exists in both configuration and filesystem
// It first verifies the component's presence in the configuration, then checks for its corresponding directory
// in the filesystem.
//
// Parameters:
//   - stackName: The name of the stack.
//   - layerName: The name of the layer within the stack.
//   - componentName: The name of the component within the layer.
//
// Returns:
//   - An error if the component does not exist in the configuration or if its directory is not found in the filesystem.
//     Returns nil if the component is valid.
func (t *StacksTransformer) ValidateComponentExists(stackName, layerName, componentName string) error {
	// Check configuration first
	component, err := t.GetComponent(stackName, layerName, componentName)
	if err != nil {
		return fmt.Errorf("error retrieving component '%s' from configuration: %w", componentName, err)
	}

	// Then check filesystem
	stackPath := filepath.Join(t.BasePath, stackName)
	componentPath := filepath.Join(stackPath, layerName, component.Name)
	if _, err := os.Stat(componentPath); os.IsNotExist(err) {
		return fmt.Errorf("component '%s' exists in configuration but not in filesystem at path: %s",
			componentName, componentPath)
	}

	return nil
}

// ValidateRequestedStack validates if a specific stack exists in both configuration and filesystem
// It first verifies the stack's presence in the configuration, then checks for its corresponding directory
// in the filesystem.
//
// Parameters:
//   - stackName: The name of the stack to validate.
//
// Returns:
//   - An error if the stack does not exist in the configuration or if its directory is not found in the filesystem.
//     Returns nil if the stack is valid.
func (t *StacksTransformer) ValidateRequestedStack(stackName string) error {
	// First check if the stack exists in configuration
	if !t.StackExists(stackName) {
		return fmt.Errorf("stack '%s' does not exist in the configuration", stackName)
	}

	// Then validate the stack in filesystem
	stackPath := filepath.Join(t.BasePath, stackName)
	if _, err := os.Stat(stackPath); os.IsNotExist(err) {
		return fmt.Errorf("stack '%s' directory not found at: %s", stackName, stackPath)
	}

	return nil
}

// ValidateRequestedLayer validates if a specific layer exists in both configuration and filesystem
// It first verifies the layer's presence in the configuration, then checks for its corresponding directory
// in the filesystem.
//
// Parameters:
//   - stackName: The name of the stack.
//   - layerName: The name of the layer within the stack.
//
// Returns:
//   - An error if the layer does not exist in the configuration or if its directory is not found in the filesystem.
//     Returns nil if the layer is valid.
func (t *StacksTransformer) ValidateRequestedLayer(stackName, layerName string) error {
	// Check configuration first
	layer, err := t.GetLayer(stackName, layerName)
	if err != nil {
		return fmt.Errorf("layer '%s' does not exist in stack '%s' configuration", layerName, stackName)
	}

	// Then check filesystem
	stackPath := filepath.Join(t.BasePath, stackName)
	layerPath := filepath.Join(stackPath, layer.Name)
	if _, err := os.Stat(layerPath); os.IsNotExist(err) {
		return fmt.Errorf("layer '%s' directory not found at: %s", layerName, layerPath)
	}

	return nil
}

// ValidateRequestedComponent validates if a specific component exists in both configuration and filesystem
// It first verifies the component's presence in the configuration, then checks for its corresponding directory
// in the filesystem. Additionally, it ensures the existence of the required 'component.hcl' Terragrunt configuration
// file within the component's directory.
//
// Parameters:
//   - stackName: The name of the stack.
//   - layerName: The name of the layer within the stack.
//   - componentName: The name of the component within the layer.
//
// Returns:
//   - An error if the component does not exist in the configuration, if its directory is not found in the filesystem,
//     or if the 'component.hcl' file is missing. Returns nil if the component is valid.
func (t *StacksTransformer) ValidateRequestedComponent(stackName, layerName, componentName string) error {
	// Check configuration first
	component, err := t.GetComponent(stackName, layerName, componentName)
	if err != nil {
		return fmt.Errorf("component '%s' does not exist in layer '%s' of stack '%s' configuration",
			componentName, layerName, stackName)
	}

	// Then check filesystem
	stackPath := filepath.Join(t.BasePath, stackName)
	componentPath := filepath.Join(stackPath, layerName, component.Name)
	if _, err := os.Stat(componentPath); os.IsNotExist(err) {
		return fmt.Errorf("component '%s' directory not found at: %s", componentName, componentPath)
	}

	// Check for required Terragrunt configuration file
	componentHCL := filepath.Join(componentPath, "component.hcl")
	if _, err := os.Stat(componentHCL); os.IsNotExist(err) {
		return fmt.Errorf("missing component.hcl for component '%s' at: %s", componentName, componentHCL)
	}

	return nil
}
