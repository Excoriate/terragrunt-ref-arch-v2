package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// GenerateHashFileWithSHA256 calculates the SHA-256 hash of a file's contents.
//
// Parameters:
//   - filePath: The full path to the file to be hashed
//
// Returns:
//   - A string representing the SHA-256 hash of the file contents
//   - An error if the file cannot be read or hashed
func GenerateHashFileWithSHA256(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s: %w", filePath, err)
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	hash := sha256.Sum256(content)

	return hex.EncodeToString(hash[:]), nil
}
