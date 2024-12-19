# 🌐 Root Terragrunt Configuration Inclusion
# Imports common configuration from the parent directory's terragrunt.hcl
# Enables consistent configuration sharing across multiple Terragrunt modules
include "root" {
  path = find_in_parent_folders("root.hcl")
  expose = true
}

# 🔧 Global Configuration Inclusion
# Imports global configuration settings from config.hcl
# Provides centralized access to repository-wide configuration parameters
include "cfg" {
  path = find_in_parent_folders("config.hcl")
  expose = true
}

# 🧩 Shared Component Configuration
# Imports standardized component configuration from the shared components directory
# IMPORTANT: Modifications should be made in the shared component configuration file
# Located at: `_shared/_components/quota-generator.hcl`
include "shared" {
  path = "${get_terragrunt_dir()}/../../../_shared/_components/quota-generator.hcl"
  expose = true
  merge_strategy = "deep"
}

locals {
  # 🌍 Base URL Configuration
  # Retrieves the base URL from global git configuration
  base_url = include.cfg.locals.git.base_url

  # 🏷️ Module Version Management
  # Explicitly define the upstream Terraform module version to ensure consistency
  # Update this version when upgrading or pinning to a specific module release
  upstream_tf_module_version = "v0.1.8"

  # 📦 Terraform Module Source Path Configuration
  # Dynamically constructs the source path for Terraform modules
  # Supports both local and remote module sources with flexible versioning
  terraform_module_name = local.component_cfg.locals.name
  terraform_modules_local_path = "../../../../terraform/quota-generator"
  terraform_module_remote_path = ""

  # 🔗 Source Path Resolution Strategy
  # - If no remote path: Uses local path + module name (e.g., "quota-generator")
  # - If remote path exists: Formats source using base URL, remote path, and version
  terraform_modules_source_path = local.terraform_module_remote_path == "" ? local.terraform_modules_local_path : format("%s/%s?ref=%s", local.base_url, local.terraform_module_remote_path, local.upstream_tf_module_version)

  # 📋 Configuration Aggregation
  # Reads configuration from different hierarchical levels to build a comprehensive input set
  component_cfg = read_terragrunt_config("${get_terragrunt_dir()}/component.hcl")
  layer_cfg = read_terragrunt_config("${find_in_parent_folders("layer.hcl")}")
  stack_cfg = read_terragrunt_config("${find_in_parent_folders("stack.hcl")}")

  # 🏗️ Hierarchical Input Resolution
  # Consolidates inputs from multiple infrastructure levels
  # Provides a flexible, layered configuration approach

  # 🌐 Stack-Level Inputs: Broad, overarching configuration for the entire stack
  stack_inputs = local.stack_cfg.locals.stack_inputs

  # 🏢 Layer-Level Inputs: Configuration specific to the infrastructure layer
  layer_inputs = local.layer_cfg.locals.layer_inputs

  # 🧩 Component-Level Inputs: Granular, specific configuration for this component
  component_inputs = local.component_cfg.locals.component_inputs
}

# 🚀 Terraform Module Source Configuration
# Dynamically sets the source path based on the resolved module path
terraform {
  source = local.terraform_modules_source_path
}

# 📥 Input Aggregation
# Merges inputs from stack, layer, and component levels
# Allows for flexible, hierarchical configuration management
inputs = merge(
  local.component_inputs,
  {}  # Placeholder for any additional, component-specific inputs
)
