locals {
  # ---------------------------------------------------------------------------------------------------------------------
  # COMPONENT CONFIGURATION
  # ---------------------------------------------------------------------------------------------------------------------
  is_enabled = true
  # Extracts the base name from the relative path of the current Terragrunt configuration
  # This name is typically used to identify the stack and can be useful for resource naming
  name = basename(path_relative_to_include())

  # 🔧 Layer Configuration Loader
  # Reads the layer configuration from the parent folder
  layer_cfg = read_terragrunt_config("${find_in_parent_folders("layer.hcl")}")
  layer_inputs = local.layer_cfg.locals.layer_inputs

  # 🔧 Component Inputs: Configurable Infrastructure Hub
  # ---------------------------------------------
  # Key Features:
  # ✅ Centralized configuration management
  # ✅ Hierarchical configuration inheritance
  # ✅ Consistent infrastructure defaults
  component_inputs = merge(local.layer_inputs, {
  })

  tags = {
    Component = local.name
  }
}
