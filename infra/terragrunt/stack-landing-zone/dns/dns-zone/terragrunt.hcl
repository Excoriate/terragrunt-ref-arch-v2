# Reference to the root Terragrunt configuration from the parent directory
# It allows sharing of common configuration across multiple Terragrunt modules
include "root" {
  path = find_in_parent_folders("terragrunt.hcl")
  expose = true
}

include "cfg" {
  path = find_in_parent_folders("config.hcl")
  expose = true
}

# Reference to the shared component configuration
# For modifications, please refer to the shared component configuration, in the `_shared/_components` folder
include "shared" {
  path = "${get_terragrunt_dir()}/../../../_shared/_components/cloudflare_dns_zone.hcl"
  expose = true 
  merge_strategy = "deep"
}

locals {
  # Use the base URL from git globals
  base_url = include.cfg.locals.git.base_url

  # 🔧 Module Version Override
  upstream_tf_module_version = "v0.1.8"

  # Local source path configuration for Terraform modules
  # If no local path is specified in the architecture configuration, 
  # leave the paths empty. Otherwise, construct a full path by joining 
  # the base local modules path with the specific cloudflare_dns_zone module path
  terraform_module_local_path = ""
  terraform_modules_local_path = include.cfg.locals.git.terraform_modules_local_path == "" ? "" : join("/", [include.cfg.locals.git.terraform_modules_local_path, "cloudflare_dns_zone"])

  # Component configuration
  component_cfg = read_terragrunt_config("${get_terragrunt_dir()}/component.hcl")
  layer_cfg = read_terragrunt_config("${find_in_parent_folders("layer.hcl")}")
  stack_cfg = read_terragrunt_config("${find_in_parent_folders("stack.hcl")}")

  # 🏗️ Resolved Inputs: Hierarchical Configuration Aggregation
  # Consolidating inputs across different infrastructure levels
  
  # 🌐 Stack-Level Inputs: Broad configuration for the entire stack
  stack_inputs = local.stack_cfg.locals.stack_inputs
  
  # 🏢 Layer-Level Inputs: Configuration specific to the infrastructure layer
  layer_inputs = local.layer_cfg.locals.layer_inputs
  
  # 🧩 Component-Level Inputs: Granular configuration for this specific component
  component_inputs = local.component_cfg.locals.component_inputs
}

# Terraform source configuration
terraform {
  source = local.terraform_module_local_path != "" ? local.terraform_module_local_path : (
    format("%s/Excoriate/terraform-cloudflare-modules.git//modules/cloudflare-zone?ref=%s", 
      local.base_url, 
      local.upstream_tf_module_version
    )
  )
}

inputs = merge(
  local.stack_inputs,
  local.layer_inputs,
  local.component_inputs,
  {}  # Placeholder for any additional inputs
)