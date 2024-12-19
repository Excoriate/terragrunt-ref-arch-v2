# 🌐 ---------------------------------------------------------------------------------------------------------------------
# 🔧 COMPONENT CONFIGURATION
# 📦 Defines the core component parameters derived from the architectural configuration
# 🏷️ Centralizes product-specific metadata and tagging for consistent resource identification
# 🌐 ---------------------------------------------------------------------------------------------------------------------
locals {
  # ---------------------------------------------------------------------------------------------------------------------
  # COMPONENT CONFIGURATION
  # ---------------------------------------------------------------------------------------------------------------------
  is_enabled = true
  name       = "name-generator"

  tags = {
    Component = "name-generator"
    ComponentTag = "component-tag"
  }
}
