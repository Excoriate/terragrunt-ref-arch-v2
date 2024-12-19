include "root" {
  path = find_in_parent_folders()
  merge_strategy = "deep"
}

include "globals_git" {
  path           = "${get_terragrunt_dir()}/../../../_globals/git.hcl"
  expose         = true
  merge_strategy = "deep"
}

include "globals_module" {
  path           = "${get_terragrunt_dir()}/../../../_components/github_repo.hcl"
  expose         = true
  merge_strategy = "deep"
}

locals {
  base_url = include.globals_git.locals.github_base_url
  upstream_tf_module_version = "v0.1.4"

  # This is the local path to the terraform-cloudflare-modules repository.
  # It is used to source the module from the local filesystem instead of the GitHub repository.
  source_local_path = "../../../../terraform/github-repo"
}

terraform {
  source = local.source_local_path == "" ? format("%s/Excoriate/terraform-cloudflare-modules.git//modules/github-repo?ref=%s", local.base_url, local.upstream_tf_module_version) : local.source_local_path
}

inputs = {
}