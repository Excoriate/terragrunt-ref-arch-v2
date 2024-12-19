package envars

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"

	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/pkg/utils"
)

// findGitRepoRoot attempts to find the root directory of the current git repository.
// Returns the absolute path to the git repository root or an error
func findGitRepoRoot() (string, error) {
	return utils.FindGitRepoRoot()
}

// LoadDotenv attempts to load environment variables from multiple potential .env file locations.
// It searches for .env files in the git repository root and current directory.
//
// The method tries to load files in the following order:
//  1. .env file in git repository root
//  2. .env.local file in git repository root
//  3. .env file in current directory
//  4. .env.local file in current directory
//
// Returns an error if any critical loading issues occur
func LoadDotenv() error {
	// Try to find git repository root
	gitRepoRoot, err := findGitRepoRoot()
	if err != nil {
		// If git repo root can't be found, continue with current directory
		gitRepoRoot = "."
	}

	// Potential .env file locations
	envFiles := []string{
		filepath.Join(gitRepoRoot, EnvFile),      // Git repo root
		filepath.Join(gitRepoRoot, LocalEnvFile), // Local environment specific
		EnvFile,                                  // Current directory fallback
		LocalEnvFile,                             // Current directory fallback
	}

	var loadedFiles []string
	var notFoundFiles []string

	for _, file := range envFiles {
		if err := loadSingleDotenv(file); err != nil {
			// If the error is because the file doesn't exist, track it
			if os.IsNotExist(err) {
				notFoundFiles = append(notFoundFiles, file)
			}
		} else {
			loadedFiles = append(loadedFiles, file)
		}
	}

	// If no files were loaded, return a summary error
	if len(loadedFiles) == 0 {
		return fmt.Errorf("no .env files found in locations: %v", notFoundFiles)
	}

	return nil
}

// loadSingleDotenv attempts to load environment variables from a single .env file.
// Parameters:
//   - filename: Path to the .env file to be loaded
//
// Returns an error if the file cannot be loaded
func loadSingleDotenv(filename string) error {
	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return err
	}

	// Load environment variables from the file
	return godotenv.Load(filename)
}

// GetLoadedEnvVars retrieves all currently loaded environment variables.
// Returns a map of environment variable names to their values
func GetLoadedEnvVars() map[string]string {
	envMap := make(map[string]string)
	for _, env := range os.Environ() {
		key, value, _ := strings.Cut(env, "=")
		envMap[key] = value
	}
	return envMap
}
