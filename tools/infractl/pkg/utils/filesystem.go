package utils

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// IsYAMLFile determines whether the provided filename has a YAML file extension.
// It checks for both ".yaml" and ".yml" extensions, which are commonly used for YAML files.
// The function returns true if the file extension matches either of these, otherwise false.
func IsYAMLFile(filename string) error {
	ext := filepath.Ext(filename)

	if ext != ".yaml" && ext != ".yml" {
		return fmt.Errorf("file %s is not a YAML file", filename)
	}

	return nil
}

// FileHasContent checks if a file has content and returns nil if the file contains content.
// It performs the following steps:
// 1. Retrieves file information using os.Stat
// 2. Returns any file access errors encountered
// 3. Returns nil if the file has content (size > 0)
// 4. Returns an error if the file is empty
func FileHasContent(filename string) error {
	info, err := os.Stat(filename)
	if err != nil {
		return fmt.Errorf("error accessing file %s: %w", filename, err)
	}

	if info.Size() == 0 {
		return fmt.Errorf("file %s is empty", filename)
	}

	return nil
}

// FileIsEmpty checks if a file is empty and returns nil if the file is empty.
// It performs the following steps:
// 1. Retrieves file information using os.Stat
// 2. Returns any file access errors encountered
// 3. Returns nil if the file is empty (size == 0)
// 4. Returns an error if the file has content
func FileIsEmpty(filename string) error {
	info, err := os.Stat(filename)
	if err != nil {
		return fmt.Errorf("error accessing file %s: %w", filename, err)
	}

	if info.Size() > 0 {
		return fmt.Errorf("file %s is not empty", filename)
	}

	return nil
}

// DirectoryExists checks if a directory exists
func DirectoryExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

// CreateDirectory creates a directory with specified permissions
func CreateDirectory(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// ListDirectoryContents returns a list of entries in a directory
func ListDirectoryContents(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var contents []string
	for _, entry := range entries {
		contents = append(contents, entry.Name())
	}

	return contents, nil
}

// MergeYAMLFiles merges two YAML files with a deep merge strategy
// It reads both input files, performs a deep merge, and returns the merged result
// If a key exists in both files, the value from the override file takes precedence
func MergeYAMLFiles(baseFile, overrideFile string) ([]byte, error) {
	// Validate input files
	if err := IsYAMLFile(baseFile); err != nil {
		return nil, fmt.Errorf("base file is not a valid YAML file: %w", err)
	}
	if err := IsYAMLFile(overrideFile); err != nil {
		return nil, fmt.Errorf("override file is not a valid YAML file: %w", err)
	}

	// Read base file
	baseData, err := os.ReadFile(baseFile)
	if err != nil {
		return nil, fmt.Errorf("reading base file: %w", err)
	}

	// Read override file
	overrideData, err := os.ReadFile(overrideFile)
	if err != nil {
		return nil, fmt.Errorf("reading override file: %w", err)
	}

	// Unmarshal base configuration
	var baseMap map[string]interface{}
	if err := yaml.Unmarshal(baseData, &baseMap); err != nil {
		return nil, fmt.Errorf("unmarshaling base file: %w", err)
	}

	// Unmarshal override configuration
	var overrideMap map[string]interface{}
	if err := yaml.Unmarshal(overrideData, &overrideMap); err != nil {
		return nil, fmt.Errorf("unmarshaling override file: %w", err)
	}

	// Perform deep merge
	mergedMap := deepMerge(baseMap, overrideMap)

	// Marshal merged configuration
	mergedYAML, err := yaml.Marshal(mergedMap)
	if err != nil {
		return nil, fmt.Errorf("marshaling merged configuration: %w", err)
	}

	return mergedYAML, nil
}

// deepMerge recursively merges two maps with override strategy
func deepMerge(base, override map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	// Copy base first
	for k, v := range base {
		result[k] = v
	}

	// Override with values from override map
	for k, v := range override {
		switch ov := v.(type) {
		case map[string]interface{}:
			// Recursively merge nested maps
			if bv, exists := result[k]; exists {
				if baseMap, ok := bv.(map[string]interface{}); ok {
					result[k] = deepMerge(baseMap, ov)
				} else {
					result[k] = ov
				}
			} else {
				result[k] = ov
			}
		default:
			result[k] = v
		}
	}

	return result
}

// FoundFilesWithExtensionInPath searches a given directory path for files with a specific file extension.
// It returns a slice of full file paths that match the specified extension.
//
// Parameters:
//   - path: The directory path to search
//   - extension: The file extension to match (including the dot, e.g. ".go")
//
// Returns:
//   - A slice of file paths with matching extensions
//   - An error if directory contents cannot be listed
func FoundFilesWithExtensionInPath(path string, extension string) ([]string, error) {
	entries, err := ListDirectoryContents(path)
	if err != nil {
		return nil, fmt.Errorf("failed to list directory contents: %w", err)
	}

	var foundFiles []string
	for _, entry := range entries {
		if filepath.Ext(entry) == extension {
			foundFiles = append(foundFiles, entry)
		}
	}

	return foundFiles, nil
}

// CreateDirIdempotent creates a directory at the specified path.
// If the directory already exists, it does nothing and returns nil.
// If the directory cannot be created, it returns an error.
//
// Parameters:
//   - path: The full path of the directory to create
//
// Returns:
//   - An error if directory creation fails, nil otherwise
func CreateDirIdempotent(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}
	return nil
}

// CreateFileWithContent creates a new file at the specified path with the given content.
// It returns an error if the file already exists or cannot be created.
// On successful creation, it returns the full path to the created file.
//
// Parameters:
//   - path: The full path where the file should be created
//   - content: The string content to write to the file
//
// Returns:
//   - The full path of the created file
//   - An error if the file exists or cannot be created
func CreateFileWithContentIdempotent(filepath string, content string) (string, error) {
	// Check if file exists and has same content
	existingContent, err := os.ReadFile(filepath)
	if err == nil {
		// File exists, check if content matches
		if string(existingContent) == content {
			return filepath, nil // File exists with same content, return success
		}
		return "", fmt.Errorf("file exists with different content at path: %s", filepath)
	} else if !os.IsNotExist(err) {
		return "", fmt.Errorf("error checking file existence: %w", err)
	}

	// Create the directory structure if it doesn't exist
	dir := path.Dir(filepath)
	if err := CreateDirIdempotent(dir); err != nil {
		return "", fmt.Errorf("failed to create directory structure: %w", err)
	}

	// Create the file
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write content to file
	if _, err := file.WriteString(content); err != nil {
		return "", fmt.Errorf("failed to write content to file: %w", err)
	}

	return filepath, nil
}

// ForceExtensionForFilepath modifies the given file path to ensure it has the specified extension.
// If the file path already has the desired extension (case insensitive), it returns the original path.
// If the file path has a different extension, it replaces the existing extension with the new one.
// If the file path does not have an extension (including cases where it is a directory),
// it appends the specified extension to the path.
//
// Parameters:
//   - configFilePath: The original file path that may or may not have an extension.
//   - extensionToForce: The desired extension to enforce on the file path.
//     This should include the leading dot (e.g., ".yaml").
//
// Returns:
//   - A string representing the modified file path with the enforced extension.
//
// Example:
//   - If the input is "example.txt" and the extension to force is ".yaml",
//     the output will be "example.yaml".
//   - If the input is "example" and the extension to force is ".yaml",
//     the output will be "example.yaml".
func ForceExtensionForFilepath(configFilePath string, extensionToForce string) string {
	if strings.EqualFold(filepath.Ext(configFilePath), extensionToForce) {
		return configFilePath
	}

	ext := filepath.Ext(configFilePath)
	if ext != "" {
		return configFilePath[:len(configFilePath)-len(ext)] + extensionToForce
	}

	// If it's a directory or path without extension, append the specified extension
	return configFilePath + extensionToForce
}

// DirExists checks if a directory exists at the given path.
//
// It uses os.Stat to get file information for the given path and checks if it's a directory.
//
// Parameters:
//   - path: The path to the directory to check.
//
// Returns:
//   - nil if the directory exists.
//   - An error if:
//   - There is an error accessing the path (e.g., path does not exist, permissions issues).
//     The error returned will wrap the underlying error from os.Stat and provide context about the operation.
//   - The path exists but is not a directory. The error will indicate that the path is not a directory.
func DirExists(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("error accessing directory %s: %w", path, err)
	}

	if !info.IsDir() {
		return fmt.Errorf("path %s is not a directory", path)
	}

	return nil
}

// FileExists checks if a file exists at the given path.
//
// It uses os.Stat to get file information for the given path and checks if it's a regular file.
//
// Parameters:
//   - path: The path to the file to check.
//
// Returns:
//   - nil if the file exists.
//   - An error if:
//   - There is an error accessing the path (e.g., path does not exist, permissions issues).
//     The error returned will wrap the underlying error from os.Stat and provide context about the operation.
//   - The path exists but is a directory. The error will indicate that the path is a directory.
func FileExists(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("error accessing file %s: %w", path, err)
	}

	if info.IsDir() {
		return fmt.Errorf("path %s is a directory", path)
	}

	return nil
}
