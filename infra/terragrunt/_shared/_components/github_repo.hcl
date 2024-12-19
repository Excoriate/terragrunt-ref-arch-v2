locals {
  # ---------------------------------------------------------------------------------------------------------------------
  # STACK CONFIGURATION
  # This section is used to configure the module with attributes that are specific to the stack.
  # ---------------------------------------------------------------------------------------------------------------------
  stack_cfg  = read_terragrunt_config(find_in_parent_folders("stack.hcl"))
  layer_cfg  = read_terragrunt_config(find_in_parent_folders("layer.hcl"))
  component_cfg = read_terragrunt_config("${get_terragrunt_dir()}/component.hcl")
  architecture_cfg = read_terragrunt_config("${get_repo_root()}/infra/terragrunt/arch.hcl")

  # ---------------------------------------------------------------------------------------------------------------------
  # PROVIDERS GLOBAL CONFIGURATION
  # ---------------------------------------------------------------------------------------------------------------------
  providers_globals_cfg = read_terragrunt_config("${get_terragrunt_dir()}/../../../_shared/_providers/github.hcl")

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

  # Merge tags with a clear precedence order
  all_tags = merge(
    local.global_tags,     # Base tags from architecture
    local.stack_tags,      # Stack-level tags
    local.layer_tags,      # Layer-specific tags
    local.component_tags   # Most specific component tags
  )

  # Git configuration
  git = local.architecture_cfg.locals.git
  base_url = local.git.base_url
}

inputs = {
  is_enabled            = local.component_cfg.locals.is_enabled
  github_token          = local.providers_globals_cfg.locals.github.token
  github_owner          = local.providers_globals_cfg.locals.github.owner
}