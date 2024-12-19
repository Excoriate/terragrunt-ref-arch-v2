include "root" {
  path = find_in_parent_folders("terragrunt.hcl")
  expose = true
}

include "parent" {
  path           = "${get_terragrunt_dir()}/../../../_shared/_components/namecheap_nameservers_updater.hcl"
  expose         = true
  merge_strategy = "deep"
}

locals {
  base_url = include.parent.locals.base_url
  upstream_tf_module_version = "v0.1.4"

  # This is the local path to the this module.
  # It is used to source the module from the local filesystem instead of the GitHub repository.
  source_local_path = "../../../../terraform/namecheap-domain-records"
}

terraform {
  source = local.source_local_path == "" ? format("%s/Excoriate/terraform-cloudflare-modules.git//modules/cloudflare-zone?ref=%s", local.base_url, local.upstream_tf_module_version) : local.source_local_path
}

inputs = {
}