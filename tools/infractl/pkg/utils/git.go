package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FindGitRepoRoot attempts to find the root of the git repository
func FindGitRepoRoot() (string, error) {
	// Start from the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Walk up the directory tree looking for .git directory
	for {
		gitDir := filepath.Join(currentDir, ".git")
		if _, err := os.Stat(gitDir); err == nil {
			return currentDir, nil
		}

		// Move up one directory
		parentDir := filepath.Dir(currentDir)

		// Stop if we've reached the filesystem root
		if parentDir == currentDir {
			return "", fmt.Errorf("git repository root not found")
		}

		currentDir = parentDir
	}
}

func AddFolderToGitIgnoreIdempotent(gitRepoRootPath, folderPath string) error {
	// Validate input paths
	if gitRepoRootPath == "" {
		return fmt.Errorf("git repository root path cannot be empty")
	}
	if folderPath == "" {
		return fmt.Errorf("folder path cannot be empty")
	}

	gitignorePath := filepath.Join(gitRepoRootPath, ".gitignore")

	// Normalize the folder path to be relative to the repo root
	relativeFolderPath := filepath.Clean(folderPath)

	// Ensure the .gitignore file exists, create if not
	file, err := os.OpenFile(gitignorePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open or create .gitignore: %w", err)
	}
	defer file.Close()

	// Read existing content
	scanner := bufio.NewScanner(file)
	var lines []string
	alreadyExists := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == relativeFolderPath {
			alreadyExists = true
			break
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading .gitignore: %w", err)
	}

	// If already exists, no need to modify
	if alreadyExists {
		return nil
	}

	// Prepare to write, ensuring a newline at the end if needed
	lines = append(lines, relativeFolderPath)

	// Truncate the file and write updated content
	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate .gitignore: %w", err)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to reset file pointer: %w", err)
	}

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("failed to write line to .gitignore: %w", err)
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush .gitignore: %w", err)
	}

	return nil
}

// IsEntryInGitIgnore checks whether a specific entry exists in the .gitignore file.
//
// This function performs a comprehensive check to determine if a given path or pattern
// is already present in the .gitignore file. It handles various scenarios such as:
// - Empty .gitignore files
// - Case-sensitive matching
// - Trimming whitespace
// - Handling absolute and relative paths
//
// Parameters:
//   - gitRepoRootPath: The root path of the git repository
//   - entry: The path or pattern to check in .gitignore
//
// Returns:
//   - bool: True if the entry exists, false otherwise
//   - error: Any error encountered during file reading or processing
func IsEntryInGitIgnore(gitRepoRootPath, entry string) (bool, error) {
	// Validate input parameters
	if gitRepoRootPath == "" {
		return false, fmt.Errorf("git repository root path cannot be empty")
	}
	if entry == "" {
		return false, fmt.Errorf("gitignore entry cannot be empty")
	}

	// Construct the full path to .gitignore
	gitignorePath := filepath.Join(gitRepoRootPath, ".gitignore")

	// Read the entire .gitignore file
	content, err := os.ReadFile(gitignorePath)
	if os.IsNotExist(err) {
		// .gitignore doesn't exist, so entry is not present
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("failed to read .gitignore: %w", err)
	}

	// Normalize the entry and convert file content to lines
	normalizedEntry := filepath.Clean(strings.TrimSpace(entry))
	lines := strings.Split(string(content), "\n")

	// Check each line for a match
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Skip empty lines and comments
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			continue
		}

		if trimmedLine == normalizedEntry {
			return true, nil
		}
	}

	return false, nil
}
