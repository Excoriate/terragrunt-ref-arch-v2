# 🌐 ---------------------------------------------------------------------------------------------------------------------
# 🔧 COMPONENT CONFIGURATION
# 📦 Defines the core component parameters derived from the architectural configuration
# 🏷️ Centralizes product-specific metadata and tagging for consistent resource identification
# 🌐 ---------------------------------------------------------------------------------------------------------------------
locals {
  # 🏗️ Configuration Loaders
  # Reads configurations from different hierarchy levels to ensure consistent, centralized management
  cfg = read_terragrunt_config("${get_repo_root()}/infra/terragrunt/config.hcl")
  stack_cfg = read_terragrunt_config("${find_in_parent_folders("stack.hcl")}")
  layer_cfg = read_terragrunt_config("${find_in_parent_folders("layer.hcl")}")

  # 📥 Retrieve stack-level configuration inputs
  # Extracts predefined inputs from the stack configuration for consistent deployment
  stack_inputs = local.stack_cfg.locals.stack_inputs

  # 📥 Retrieve layer-level configuration inputs
  # Extracts predefined inputs from the layer configuration for granular control
  layer_inputs = local.layer_cfg.locals.layer_inputs

  # 🔧 Module Version Override
  # Allows for explicit module version specification, overriding default versions
  # Useful for testing, gradual changes, or precise version control
  upstream_tf_module_version_override = ""

  # 🏷️ Extract stack-level tags for consistent resource labeling
  # Retrieves predefined tags from the stack configuration
  stack_tags = local.stack_cfg.locals.tags

  # 🏷️ Extract layer-level tags for granular resource identification
  # Retrieves predefined tags from the layer configuration
  layer_tags = local.layer_cfg.locals.tags

  # 🏷️ Comprehensive Tagging Strategy
  # Merges tags from architectural, stack, layer, and component levels
  tags = merge(
    # Component-level tags
    {
      Component = "dns-zone"
      Enabled = local.is_enabled
    },
    # Stack-level tags
    local.stack_tags,
    # Layer-level tags
    local.layer_tags,
  )

  # 🌐 Component Inputs
  # Merges inputs from stack, layer, and component levels
  # Provides a flexible, hierarchical configuration mechanism
  component_inputs = merge(
    # Stack-level inputs (if any)
    local.stack_inputs,
    
    # Layer-level inputs
    local.layer_inputs,
    
    # Component-specific inputs
    {
      is_enabled = local.is_enabled
    }
  )

  is_enabled = true
}