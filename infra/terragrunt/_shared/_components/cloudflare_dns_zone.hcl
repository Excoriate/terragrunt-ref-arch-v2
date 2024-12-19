locals {
  # ---------------------------------------------------------------------------------------------------------------------
  # 🏗️ STACK CONFIGURATION
  # This section dynamically loads configuration files to set up the infrastructure stack.
  # It reads various configuration layers to provide a flexible and modular infrastructure setup.
  # 
  # Key configurations loaded:
  # - Stack-level settings from stack.hcl
  # - Layer-specific configurations from layer.hcl
  # - Component-specific details from component.hcl
  # - Overarching architecture settings from arch.hcl
  # ---------------------------------------------------------------------------------------------------------------------
  stack_cfg  = read_terragrunt_config(find_in_parent_folders("stack.hcl"))
  layer_cfg = read_terragrunt_config(find_in_parent_folders("layer.hcl"))
  component_cfg = read_terragrunt_config("${get_terragrunt_dir()}/component.hcl")

  # ---------------------------------------------------------------------------------------------------------------------
  # 🏗️ ARCHITECTURE CONFIGURATION
  # This section dynamically loads configuration files to set up the infrastructure stack.
  # It reads various configuration layers to provide a flexible and modular infrastructure setup.
  # ---------------------------------------------------------------------------------------------------------------------
  cfg = read_terragrunt_config("${get_repo_root()}/infra/terragrunt/config.hcl")

  # ---------------------------------------------------------------------------------------------------------------------
  # 🏷️ TAG ORCHESTRATION
  # A sophisticated tag management system that layers and merges tags from multiple configuration levels
  # 
  # Tag Hierarchy (from broadest to most specific):
  # 1. Global Architecture Tags 🌐
  # 2. Stack-Level Tags 🏗️
  # 3. Layer-Specific Tags 📦
  # 4. Component-Level Tags 🧩
  #
  # Source: '_globals/ownership.hcl' provides ownership-related tags
  # Allows flexible, hierarchical tag application across infrastructure resources
  # ---------------------------------------------------------------------------------------------------------------------
  stack_tags = local.stack_cfg.locals.tags
  layer_tags = local.layer_cfg.locals.tags
  component_tags = local.component_cfg.locals.tags

  all_tags = merge(
    local.stack_tags,      # Stack-level tags
    local.layer_tags,      # Layer-specific tags
    local.component_tags   # Most specific component tags
  )
}

inputs = {
  cloudflare_account_id = "601ae11910f6e4fe96f8757d6e0c9a3c" 
  is_enabled            = true
  domains               = [
    {
      name   = local.cfg.locals.product.name
      domain = "seko.io"
    }
  ]
  
  # Apply the merged tags to the inputs
  tags = local.all_tags
}