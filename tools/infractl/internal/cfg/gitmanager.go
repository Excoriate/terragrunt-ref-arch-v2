package cfg

import (
	"fmt"

	"github.com/Excoriate/terragrunt-ref-arch-v2/tools/infractl/pkg/utils"
)

// AddEntriesToGitignore ensures that the .infra-cache directory is added to the .gitignore file.
//
// This function:
// - Checks if .infra-cache is already in the .gitignore file
// - Adds the entry if it's not present
// - Is idempotent, meaning multiple calls will not duplicate the entry
//
// Returns:
//   - nil if the .infra-cache entry is successfully added or already exists
//   - An error if there are issues reading or writing the .gitignore file
func AddEntriesToGitignore() error {
	// Find the git repository root
	gitRepoRoot, err := GetGitRepoRoot()
	if err != nil {
		return fmt.Errorf("failed to find git repository root: %w", err)
	}

	// Check and add each ignore entry
	for _, entry := range InfraCtlGitIgnoreEntries {
		exists, err := utils.IsEntryInGitIgnore(gitRepoRoot, entry)
		if err != nil {
			return fmt.Errorf("failed to check .gitignore for %s: %w", entry, err)
		}

		if !exists {
			if err := utils.AddFolderToGitIgnoreIdempotent(gitRepoRoot, entry); err != nil {
				return fmt.Errorf("failed to add %s to .gitignore: %w", entry, err)
			}
		}
	}

	return nil
}

// IsInfraCacheDirInGitignore checks if any of the infrastructure cache directory entries
// are present in the .gitignore file.
//
// This function:
// - Finds the git repository root
// - Checks each predefined infrastructure cache directory ignore entry
// - Returns true if any entry is found in .gitignore
// - Returns false if no entries are found
//
// Returns:
//   - bool: true if any infrastructure cache directory is in .gitignore, false otherwise
//   - error: any error encountered while finding the git repo root or checking .gitignore
func IsInfraCacheDirInGitignore() (bool, error) {
	gitRepoRoot, err := utils.FindGitRepoRoot()
	if err != nil {
		return false, fmt.Errorf("failed to find git repository root: %w", err)
	}

	for _, entry := range InfraCtlGitIgnoreEntries {
		exists, err := utils.IsEntryInGitIgnore(gitRepoRoot, entry)
		if err != nil {
			return false, fmt.Errorf("failed to check .gitignore for %s: %w", entry, err)
		}
		if exists {
			return true, nil
		}
	}

	return false, nil
}
