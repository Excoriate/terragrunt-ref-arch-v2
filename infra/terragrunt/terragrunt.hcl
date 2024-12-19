locals {
  # Import deployment configuration from architecture.hcl
  cfg = read_terragrunt_config("${get_repo_root()}/infra/terragrunt/config.hcl")
  
  # Configuration extraction
  product_name     = local.cfg.locals.product.name
  product_version  = local.cfg.locals.product.version
  product_description = local.cfg.locals.product.description

  # Remote state configuration from iac block
  remote_state_config = local.cfg.locals.iac.remote_state.s3

  bucket_name   = local.remote_state_config.bucket
  bucket_region = local.remote_state_config.region
  lock_table    = local.remote_state_config.lock_table
  
  # Tags
  tags = {
    "git-sha"       = run_cmd("--terragrunt-global-cache", "--terragrunt-quiet", "git", "rev-parse", "--short", "HEAD")
    "managed_by"    = "terragrunt"
    "product_name"  = local.product_name
    "version"       = local.product_version
    "description"   = local.product_description
  }

  # Provider configuration content
  providers_content = join("\n\n", [
    for provider_name, provider_config in local.cfg.locals.providers : 
    format(
      <<-EOT
      provider "%s" {
        %s
      }
      EOT
      ,
      provider_name,
      join("\n  ", [
        for key, value in provider_config.config :
        format("%s = %q", key, value)
      ])
    )
  ])

  # Version constraints content
  versions_content = length([
    for provider_config in local.cfg.locals.providers :
    provider_config if provider_config.version_constraint.enabled
  ]) > 0 ? join("\n", [
    "terraform {",
    "  required_providers {",
    join("\n", [
      for provider_name, provider_config in local.cfg.locals.providers :
      provider_config.version_constraint.enabled ? format(
        <<-EOT
        %s = {
          source  = "%s"
          version = "%s"
        }
        EOT
        ,
        provider_name,
        provider_config.version_constraint.source,
        provider_config.version_constraint.required_version
      ) : ""
    ]),
    "  }",
    "}"
  ]) : ""
}

# Generate provider configuration using template
generate "providers" {
  path      = "providers.tf"
  if_exists = "overwrite"
  contents  = templatefile(
    "${get_repo_root()}/infra/terragrunt/_templates/providers.tf.tpl",
    {
      providers_content = local.providers_content
    }
  )
}

# Generate version constraints using template
generate "versions" {
  path      = "versions.tf"
  if_exists = "overwrite"
  contents  = templatefile(
    "${get_repo_root()}/infra/terragrunt/_templates/versions.tf.tpl",
    {
      versions_content = local.versions_content
    }
  )
}

# Remote State Configuration
remote_state {
  backend = "s3"
  generate = {
    path      = "_backend.tf"
    if_exists = "overwrite"
  }

  config = {
    region         = local.bucket_region
    bucket         = local.bucket_name
    dynamodb_table = local.lock_table
    encrypt        = true
    key            = format("%s/%s.tfstate", 
      local.product_name,
      replace(path_relative_to_include(), "/", "-")
    )

    s3_bucket_tags      = local.tags
    dynamodb_table_tags = local.tags
  }
}