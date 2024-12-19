  # 🏗️ ---------------------------------------------------------------------------------------------------------------------
  # 🔧 STACK CONFIGURATION
  # 📦 Defines the core stack parameters derived from the architectural configuration
  # 🏷️ Centralizes product-specific metadata and tagging for consistent resource identification
  # 🌐 ---------------------------------------------------------------------------------------------------------------------
locals {
  # 🏗️ Architecture Configuration Loader
  # Reads the centralized architectural configuration from the project's architecture definition file
  # This allows for consistent, centralized management of product-level configuration across the infrastructure
  cfg = read_terragrunt_config("${get_repo_root()}/infra/terragrunt/config.hcl")

  # Extracts the base name from the relative path of the current Terragrunt configuration
  # This name is typically used to identify the stack and can be useful for resource naming
  name = basename(path_relative_to_include())

  # 📦 Product Configuration
  # Extracts product-specific metadata from the architectural configuration
  product_name = local.cfg.locals.product.name
  product_version = local.cfg.locals.product.version
  product_description = local.cfg.locals.product.description
  use_as_stack_tags = local.cfg.locals.product.use_as_stack_tags

  # 🔍 Tags Configuration
  # Defines the tags for the stack, including product-specific metadata
  tags_products = {
    Name = local.product_name
    Product = local.product_name
    Description = local.product_description
    Version = local.product_version
  }

  tags_default = {
    ManagedBy = "terragrunt"
    Stack = local.name
  }

  tags = local.use_as_stack_tags ? merge(local.tags_products, local.tags_default) : local.tags_default

  # 🔧 Stack Inputs
  # Defines the inputs for the stack, including layer configurations
  stack_inputs = {
  }
}
