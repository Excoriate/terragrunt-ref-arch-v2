package tg

import (
	"fmt"
	"os/exec"
	"strings"
)

// ValidateTerragruntInstallation checks if Terragrunt is installed and meets version requirements
func ValidateTerragruntInstallation(minVersion string) error {
	// Check if Terragrunt binary exists
	_, err := exec.LookPath("terragrunt")
	if err != nil {
		return fmt.Errorf("terragrunt not found in system path: ensure Terragrunt is installed")
	}

	// Get Terragrunt version
	cmd := exec.Command("terragrunt", "version")
	versionOutput, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to retrieve Terragrunt version: %v", err)
	}

	// Extract version from output
	version := extractVersion(string(versionOutput))

	// Compare version with minimum required version
	if compareVersions(version, minVersion) < 0 {
		return fmt.Errorf("terragrunt version %s is lower than required minimum %s", version, minVersion)
	}

	return nil
}

// extractVersion extracts the version string from the version command output
func extractVersion(versionOutput string) string {
	lines := strings.Split(versionOutput, "\n")
	if len(lines) > 0 {
		parts := strings.Fields(lines[0])
		for _, part := range parts {
			if strings.Count(part, ".") >= 1 {
				return part
			}
		}
	}
	return "0.0.0"
}

// compareVersions compares two version strings
// Returns:
// -1 if v1 < v2
//
//	0 if v1 == v2
//	1 if v1 > v2
func compareVersions(v1, v2 string) int {
	v1Parts := strings.Split(v1, ".")
	v2Parts := strings.Split(v2, ".")

	for i := 0; i < len(v1Parts) && i < len(v2Parts); i++ {
		v1Part := strings.TrimSpace(v1Parts[i])
		v2Part := strings.TrimSpace(v2Parts[i])

		v1Num := parseVersionPart(v1Part)
		v2Num := parseVersionPart(v2Part)

		if v1Num < v2Num {
			return -1
		}
		if v1Num > v2Num {
			return 1
		}
	}

	// If we've gotten this far, the common parts are equal
	if len(v1Parts) < len(v2Parts) {
		return -1
	}
	if len(v1Parts) > len(v2Parts) {
		return 1
	}

	return 0
}

// parseVersionPart converts a version part to an integer, handling potential non-numeric parts
func parseVersionPart(part string) int {
	var num int
	fmt.Sscanf(part, "%d", &num)
	return num
}
