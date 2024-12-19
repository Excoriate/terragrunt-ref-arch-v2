locals {
  # ---------------------------------------------------------------------------------------------------------------------
  # COMPONENT CONFIGURATION
  # ---------------------------------------------------------------------------------------------------------------------

  # 🔧 Flag to enable or disable the component.
  # Setting this to true means the component will be active and provisioned.
  is_enabled = true

  # 🏷️ Extracts the base name from the relative path of the current Terragrunt configuration.
  # This name is typically used to identify the stack and can be useful for resource naming,
  # ensuring that resources are easily identifiable and organized.
  name = basename(path_relative_to_include())

  # 🔧 Layer Configuration Loader
  # This section reads the layer configuration from the parent folder.
  # The layer configuration contains inputs and settings that are shared across multiple components,
  # allowing for a more modular and maintainable infrastructure setup.
  layer_cfg = read_terragrunt_config("${find_in_parent_folders("layer.hcl")}")

  # 📥 Extracts the layer inputs from the loaded layer configuration.
  # These inputs can include various settings and parameters that the component can utilize.
  layer_inputs = local.layer_cfg.locals.layer_inputs

  # 🔧 Component Inputs: Configurable Infrastructure Hub
  # ---------------------------------------------
  # This section defines the inputs specific to this component.
  # It merges the layer inputs with any additional component-specific configurations.
  # Key Features of this approach:
  # ✅ Centralized configuration management: Allows for easier updates and consistency across components.
  # ✅ Hierarchical configuration inheritance: Enables components to inherit settings from parent layers.
  # ✅ Consistent infrastructure defaults: Ensures that all components start with a standard set of configurations.
  component_inputs = merge(local.layer_inputs, {
    # 📝 Additional component-specific inputs can be added here as key-value pairs.
  })

  # 🏷️ Tags are used for resource organization and identification in cloud environments.
  # Here, we are tagging the component with its name for easier tracking and management.
  tags = {
    Component = local.name
  }
}
