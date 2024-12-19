package cfg

import (
	"os"
)

const (
	// INFRACTL_CONFIG_FILE_PATH is the path to the file that will be used to store the compiled configuration.
	InfractlConfigFilePathEnvVar = "INFRACTL_CONFIG_FILE_PATH"
)

type TransmitterEnvVar struct {
	Key   string
	Value string
}

// GetTransmitterEnvVar converts a file path into an environment variable for the configuration file path.
//
// This function prepares the INFRACTL_CONFIG_FILE_PATH environment variable by taking a file path.
// It creates a FileTransmitterEnvVar struct that can be used to set the environment variable.
//
// Args:
//
//	value: The full path to the configuration file
//
// Returns:
//
//	A FileTransmitterEnvVar representing the configuration file path environment variable
func GetTransmitterEnvVar(value string) TransmitterEnvVar {
	return TransmitterEnvVar{
		Key:   InfractlConfigFilePathEnvVar,
		Value: value,
	}
}

// SetTransmitterTgEnvVar sets an environment variable for Terragrunt configuration.
//
// This function takes a FileTransmitterEnvVar pointer and sets the environment variable
// using the Key and Value fields. Note: The current implementation appears to have
// a duplicate os.Setenv call which may be unintentional.
//
// Args:
//
//	envVarsToSet: A pointer to a FileTransmitterEnvVar containing the key and value
//	              to be set as an environment variable
func SetTransmitterTgEnvVar(value string) {
	envVar := GetTransmitterEnvVar(value)
	os.Setenv(envVar.Key, envVar.Value)
}
