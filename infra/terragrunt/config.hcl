locals {
  # Dynamically load the configuration JSON file from the environment variable
  default_env_config_json_path = "${get_repo_root()}/infra/.infractl-cache/test.json"
  env_config_json_path = get_env("INFRACTL_CONFIG_JSON_PATH", local.default_env_config_json_path)
  
  # Ensure the file exists and is readable
  config_file = jsondecode(file(local.env_config_json_path))

# Top-level configurations matching YAML/JSON structure
  config = {
    version     = local.config_file.config.version
    last_updated = local.config_file.config.last_updated
    description = local.config_file.config.description
  }

  # Product Configuration
  product = {
    name        = local.config_file.product.name
    version     = local.config_file.product.version
    description = local.config_file.product.description
    use_as_stack_tags = local.config_file.product.use_as_stack_tags
  }

  # Git Configuration
  git = {
    base_url = local.config_file.git.base_url
    terraform_modules_local_path = local.config_file.git.terraform_modules_local_path
  }

  # Infrastructure as Code Configuration
  iac = {
    versions = {
      terraform_version_default  = local.config_file.iac.versions.terraform_version_default
      terragrunt_version_default = local.config_file.iac.versions.terragrunt_version_default
    }
    remote_state = {
      s3 = {
        bucket    = local.config_file.iac.remote_state.s3.bucket
        lock_table = local.config_file.iac.remote_state.s3.lock_table
        region    = local.config_file.iac.remote_state.s3.region
      }
    }
  }

  # Providers Configuration
  providers = {
    for name, provider in local.config_file.providers : name => {
      config = provider.config
      version_constraint = {
        source = provider.version_constraint.source
        required_version = provider.version_constraint.required_version
        enabled = provider.version_constraint.enabled
      }
    }
  }

  # Stacks Configuration
  stacks = local.config_file.stacks

  # Secrets Configuration (masked for security)
  secrets = {
    for name, secret in local.config_file.secrets : name => {
      # Only include keys, masking actual secret values
      # It's possible to get the secret values from the environment variables using the get_env function
      keys = keys(secret)
    }
  }
}