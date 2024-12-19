locals {
  # ---------------------------------------------------------------------------------------------------------------------
  # HIERARCHY CONFIGURATION
  # ---------------------------------------------------------------------------------------------------------------------
  # This section dynamically loads configuration files from different levels of the project hierarchy
  # to provide a flexible and modular configuration management approach.
  #
  # The configurations are loaded in order of increasing specificity:
  # 1. Architecture-level configuration (most generic)
  # 2. Stack-level configuration
  # 3. Layer-level configuration
  # 4. Component-level configuration (most specific)
  #
  # This approach allows for:
  # - Centralized configuration management
  # - Overriding of settings at different project levels
  # - Maintaining a clear configuration inheritance structure
  # ---------------------------------------------------------------------------------------------------------------------
  stack_cfg  = read_terragrunt_config(find_in_parent_folders("stack.hcl"))
  layer_cfg = read_terragrunt_config(find_in_parent_folders("layer.hcl"))
  component_cfg = read_terragrunt_config("${get_terragrunt_dir()}/component.hcl")
  architecture_cfg = read_terragrunt_config("${get_repo_root()}/infra/terragrunt/arch.hcl")



  # ---------------------------------------------------------------------------------------------------------------------
  # PROVIDERS GLOBAL CONFIGURATION
  # ---------------------------------------------------------------------------------------------------------------------
  providers_globals_cfg = read_terragrunt_config("${get_terragrunt_dir()}/../../../_shared/_providers/cloudflare.hcl")

  # ---------------------------------------------------------------------------------------------------------------------
  # COMPONENT GLOBAL TAGS
  # ---------------------------------------------------------------------------------------------------------------------
  # This local variable defines global tags that can be applied to all resources within this Cloudflare DNS domains HTTPS component.
  # 
  # Purpose:
  # - Provides a centralized location to define tags that should be universally applied
  # - Allows for easy addition of consistent metadata across all resources
  # - Can be merged with more specific component-level tags
  #
  # By default, this is an empty map, which means no global tags are automatically applied.
  # To add global tags, simply add key-value pairs to this map.
  # 
  # Example usage:
  # component_global_tags = {
  #   "managed-by" = "terragrunt"
  #   "component"  = "cloudflare-dns-domains-https"
  # }
  component_global_tags = {}

  # ---------------------------------------------------------------------------------------------------------------------
  # TAGS
  # These tags are a mix of global tags and stack specific tags, including 'ownership' tags
  # That are defined in the '_globals/ownership.hcl' file.
  # These tags are referenced in the child terragrunt configurations, and from there they can be merged
  # with specific tags for each resource.
  # ---------------------------------------------------------------------------------------------------------------------
  global_tags = local.architecture_cfg.locals.global_tags
  stack_tags = local.stack_cfg.locals.tags
  layer_tags = local.layer_cfg.locals.tags
  component_tags = local.component_cfg.locals.tags
  component_tags_merged = merge(local.component_tags, local.component_global_tags)

  # Merge tags with a clear precedence order
  all_tags = merge(
    local.global_tags,     # Base tags from architecture
    local.stack_tags,      # Stack-level tags
    local.layer_tags,      # Layer-specific tags
    local.component_tags_merged,  # Most specific component tags
  )

  # Git configuration
  git = local.architecture_cfg.locals.git
  base_url = local.git.base_url

  # Layer Inputs.
  layer_inputs = try(
    local.layer_cfg.locals.layer_inputs,
    {}
  )

  # Stack Inputs
  stack_inputs = try(
    local.stack_cfg.locals.stack_inputs,
    {}
  )
}

dependency "cloudflare_dns_zone" {
  config_path = find_in_parent_folders("dns-zone")
  mock_outputs = {
    cloudflare_zone_ids = {
      "fake-zone-id" = "fake-zone-id"
    }
  }
}

inputs = merge(
  local.layer_inputs,
  local.stack_inputs,
  {
    zone_id = one(values(dependency.cloudflare_dns_zone.outputs.cloudflare_zone_ids))
    cloudflare_account_id = local.providers_globals_cfg.locals.cloudflare.account_id
    is_enabled = local.component_cfg.locals.is_enabled
    # domain = local.layer_cfg.locals.layer_inputs.domain
  }
)
