# 🌐 ---------------------------------------------------------------------------------------------------------------------
# 🔧 LAYER CONFIGURATION
# 📦 Defines the core layer parameters derived from the architectural configuration
# 🏷️ Centralizes product-specific metadata and tagging for consistent resource identification
# 🌐 ---------------------------------------------------------------------------------------------------------------------
locals {
  # 🏗️ Architecture Configuration Loader
  # Reads the centralized architectural configuration from the project's architecture definition file
  # This allows for consistent, centralized management of product-level configuration across the infrastructure
  cfg = read_terragrunt_config("${get_repo_root()}/infra/terragrunt/config.hcl")
  stack_cfg = read_terragrunt_config("${find_in_parent_folders("stack.hcl")}")
  stack_inputs = local.stack_cfg.locals.stack_inputs

  # 🏷️ Comprehensive Tagging Strategy
  # Defines tags for the layer, combining architectural, product, and layer-specific metadata
  tags_default = {
    Layer = "dns"
  }

  tags = merge(local.tags_default, local.stack_cfg.locals.tags)

  # 🔧 Layer Inputs: Configurable Infrastructure Hub
  # ---------------------------------------------
  # Key Features:
  # ✅ Centralized configuration management
  # ✅ Hierarchical configuration inheritance
  # ✅ Consistent infrastructure defaults
  layer_inputs = merge(local.stack_inputs, {
  })
}
