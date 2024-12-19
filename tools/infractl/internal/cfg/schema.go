package cfg

// RootConfig represents the root level config field
type RootConfig struct {
	Version     string `yaml:"version" json:"version"`
	LastUpdated string `yaml:"last_updated" json:"last_updated"`
	Description string `yaml:"description" json:"description"`
}

// Git represents the Git configuration
type Git struct {
	BaseURL string `yaml:"base_url" json:"base_url"`
}

// Product represents the global product identification
type Product struct {
	Name           string `yaml:"name" json:"name"`
	Version        string `yaml:"version" json:"version"`
	Description    string `yaml:"description" json:"description"`
	UseAsStackTags bool   `yaml:"use_as_stack_tags" json:"use_as_stack_tags"`
}

// IaCVersions represents the version configurations for infrastructure tools
type IaCVersions struct {
	TerraformVersionDefault  string `yaml:"terraform_version_default" json:"terraform_version_default"`
	TerragruntVersionDefault string `yaml:"terragrunt_version_default" json:"terragrunt_version_default"`
}

// RemoteStateS3 represents the S3 configuration for remote state storage
type RemoteStateS3 struct {
	Bucket    string `yaml:"bucket" json:"bucket"`
	LockTable string `yaml:"lock_table" json:"lock_table"`
	Region    string `yaml:"region" json:"region"`
}

// RemoteState represents the remote state configuration
type RemoteState struct {
	S3 RemoteStateS3 `yaml:"s3" json:"s3"`
}

// IaC represents Infrastructure as Code configuration
type IaC struct {
	Versions    IaCVersions `yaml:"versions" json:"versions"`
	RemoteState RemoteState `yaml:"remote_state" json:"remote_state"`
}

// ProviderConfig represents the configuration for a single provider
type ProviderConfig struct {
	Config            map[string]interface{} `yaml:"config" json:"config"`
	VersionConstraint VersionConstraint      `yaml:"version_constraint" json:"version_constraint"`
}

// Providers represents a dynamic map of provider configurations
type Providers map[string]ProviderConfig

// Secrets represents a dynamic, flexible secrets configuration
type Secrets map[string]map[string]string

// ComponentConfig represents a single component configuration
type ComponentConfig struct {
	Name      string                 `yaml:"name" json:"name"`
	Providers []string               `yaml:"providers" json:"providers"`
	Tags      map[string]string      `yaml:"tags" json:"tags"`
	Inputs    map[string]interface{} `yaml:"inputs,omitempty" json:"inputs,omitempty"`
}

// LayerConfig represents a layer configuration within a stack
type LayerConfig struct {
	Name       string                 `yaml:"name" json:"name"`
	Tags       map[string]string      `yaml:"tags" json:"tags"`
	Components []ComponentConfig      `yaml:"components" json:"components"`
	Inputs     map[string]interface{} `yaml:"inputs,omitempty" json:"inputs,omitempty"`
}

// StackConfig represents a stack configuration
type StackConfig struct {
	Name   string                 `yaml:"name" json:"name"`
	Tags   map[string]string      `yaml:"tags" json:"tags"`
	Layers []LayerConfig          `yaml:"layers" json:"layers"`
	Inputs map[string]interface{} `yaml:"inputs,omitempty" json:"inputs,omitempty"`
}

// VersionConstraint represents the version constraint configuration for a provider
type VersionConstraint struct {
	Source          string `yaml:"source" json:"source"`
	RequiredVersion string `yaml:"required_version" json:"required_version"`
	Enabled         bool   `yaml:"enabled" json:"enabled"`
}

// EnvConfig represents the complete environment configuration
type EnvConfig struct {
	Config    RootConfig    `yaml:"config" json:"config"`
	Git       Git           `yaml:"git" json:"git"`
	Product   Product       `yaml:"product" json:"product"`
	IAC       IaC           `yaml:"iac" json:"iac"`
	Providers Providers     `yaml:"providers" json:"providers"`
	Secrets   Secrets       `yaml:"secrets" json:"secrets"`
	Stacks    []StackConfig `yaml:"stacks" json:"stacks"`
}

// ToMap method can be removed or simplified if no longer needed
func (e *EnvConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"config":    e.Config,
		"git":       e.Git,
		"product":   e.Product,
		"iac":       e.IAC,
		"providers": e.Providers,
		"secrets":   e.Secrets,
		"stacks":    e.Stacks,
	}
}
