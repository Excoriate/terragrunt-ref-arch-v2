locals {
  # ---------------------------------------------------------------------------------------------------------------------
  # 🏗️ ARCHITECTURE CONFIGURATION
  # This section dynamically loads configuration files to set up the infrastructure stack.
  # It reads various configuration layers to provide a flexible and modular infrastructure setup.
  # ---------------------------------------------------------------------------------------------------------------------
  cfg = read_terragrunt_config("${get_repo_root()}/infra/terragrunt/config.hcl")
}

inputs = {
   min_quota         = 50
   max_quota         = 200
   quota_type        = "concurrent_users"
   generation_seed   = null
   scale_factor      = 1.5
   adjustment_type   = "floor"
}
