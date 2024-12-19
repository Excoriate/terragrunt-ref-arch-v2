# Reference to the root Terragrunt configuration from the parent directory
# It allows sharing of common configuration across multiple Terragrunt modules
include "root" {
  path = find_in_parent_folders("terragrunt.hcl")
  expose = true
}

include "cfg" {
  path = find_in_parent_folders("arch.hcl")
  expose = true
}

# Reference to the shared component configuration
# For modifications, please refer to the shared component configuration, in the `_shared/_components` folder
include "shared" {
  path = "${get_terragrunt_dir()}/../../../_shared/_components/cloudflare_dns_zone.hcl"
  expose = true 
  merge_strategy = "deep"
}

locals {
  base_url = include.parent.locals.base_url
  upstream_tf_module_version = "v0.1.4"

  # This is the local path to the terraform-cloudflare-modules repository.
  # It is used to source the module from the local filesystem instead of the GitHub repository.
  source_local_path = "../../../../terraform/cloudflare-dns-domains-https"
}

terraform {
  source = local.source_local_path == "" ? format("%s/Excoriate/terraform-cloudflare-modules.git//modules/cloudflare-zone?ref=%s", local.base_url, local.upstream_tf_module_version) : local.source_local_path
}

inputs = merge(
  local.stack_inputs,
  local.layer_inputs,
  local.component_inputs,
  {}  # Placeholder for any additional inputs
)