# 🌐 Root Terragrunt Configuration Inclusion
# This block imports the common configuration from the parent directory's terragrunt.hcl file.
# It enables consistent configuration sharing across multiple Terragrunt modules, ensuring that
# all modules can access shared settings and parameters defined at the root level.
include "root" {
  path = find_in_parent_folders("root.hcl")  # Path to the root configuration file
  expose = true  # Exposes the included configuration to child modules
}

# 🔧 Global Configuration Inclusion
# This block imports global configuration settings from the config.hcl file located in the parent directory.
# It provides centralized access to repository-wide configuration parameters, allowing for
# consistent settings across different modules and environments.
include "cfg" {
  path = find_in_parent_folders("config.hcl")  # Path to the global configuration file
  expose = true  # Exposes the included configuration to child modules
}

# 🧩 Shared Component Configuration
# This block imports standardized component configuration from the shared components directory.
# It is important to note that any modifications should be made in the shared component configuration file
# located at: `_shared/_components/quota-generator.hcl`. This ensures that changes are reflected
# across all modules that utilize this shared configuration.
include "shared" {
  path = "${get_terragrunt_dir()}/../../../_shared/_components/quota-generator.hcl"  # Path to the shared component configuration
  expose = true  # Exposes the included configuration to child modules
  merge_strategy = "deep"  # Merges the shared configuration deeply with local configurations
}

locals {
  # 🌍 Base URL Configuration
  # This local variable retrieves the base URL from the global git configuration.
  # It is used to construct the source path for remote Terraform modules.
  base_url = include.cfg.locals.git.base_url

  # 🏷️ Module Version Management
  # This local variable explicitly defines the upstream Terraform module version to ensure consistency.
  # It is crucial to update this version when upgrading or pinning to a specific module release.
  upstream_tf_module_version = "v0.1.8"

  # 📦 Terraform Module Source Path Configuration
  # These local variables dynamically construct the source path for Terraform modules.
  # They support both local and remote module sources with flexible versioning.
  terraform_module_name = local.component_cfg.locals.name  # Name of the component
  terraform_modules_local_path = "../../../../terraform/quota-generator"  # Local path to the Terraform module
  terraform_module_remote_path = ""  # Placeholder for remote module path

  # 🔗 Source Path Resolution Strategy
  # This local variable determines the source path for the Terraform module.
  # - If no remote path is provided, it uses the local path combined with the module name (e.g., "quota-generator").
  # - If a remote path exists, it formats the source using the base URL, remote path, and version.
  terraform_modules_source_path = local.terraform_module_remote_path == "" ? local.terraform_modules_local_path : format("%s/%s?ref=%s", local.base_url, local.terraform_module_remote_path, local.upstream_tf_module_version)

  # 📋 Configuration Aggregation
  # This local variable reads configuration from different hierarchical levels to build a comprehensive input set.
  # It allows for the aggregation of settings from various sources, ensuring that all necessary configurations are included.
  component_cfg = read_terragrunt_config("${get_terragrunt_dir()}/component.hcl")  # Reads the component configuration file

  # 🧩 Component-Level Inputs
  # This local variable holds granular, specific configuration for this component.
  # It allows for the customization of inputs specific to the component's requirements.
  component_inputs = local.component_cfg.locals.component_inputs
}

# 🚀 Terraform Module Source Configuration
# This block dynamically sets the source path for the Terraform module based on the resolved module path.
# It ensures that the correct module is referenced during the Terraform execution.
terraform {
  source = local.terraform_modules_source_path  # Source path for the Terraform module
}

# 📥 Input Aggregation
# This block merges inputs from stack, layer, and component levels.
# It allows for flexible, hierarchical configuration management, ensuring that all relevant inputs are included.
inputs = merge(
  local.component_inputs,  # Component-specific inputs
  {}  # Placeholder for any additional, component-specific inputs
)
