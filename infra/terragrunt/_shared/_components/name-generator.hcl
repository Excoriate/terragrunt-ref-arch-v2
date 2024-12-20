locals {
  # ---------------------------------------------------------------------------------------------------------------------
  # 🏗️ ARCHITECTURE CONFIGURATION
  # This section dynamically loads configuration files to set up the infrastructure stack.
  # It reads various configuration layers to provide a flexible and modular infrastructure setup.
  # ---------------------------------------------------------------------------------------------------------------------
  cfg = read_terragrunt_config("${get_repo_root()}/infra/terragrunt/config.hcl")
}

inputs = {
  prefix_length = 8
  prefix_type   = "generic"
  name_template = "%s-resource"
}
