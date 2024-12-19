locals {
  # ---------------------------------------------------------------------------------------------------------------------
  # 🏗️ ARCHITECTURE CONFIGURATION
  # This section dynamically loads configuration files to set up the infrastructure stack.
  # It reads various configuration layers to provide a flexible and modular infrastructure setup.
  # ---------------------------------------------------------------------------------------------------------------------
  cfg = read_terragrunt_config("${get_repo_root()}/infra/terragrunt/config.hcl")

  # Dynamic naming convention
  table_name_prefix = "dynamodb"
  environment       = "dev"
}

dependency "name_generator" {
  config_path = "${get_repo_root()}/infra/terragrunt/stack-datastore/db/name-generator"
  mock_outputs = {
    full_resource_name = "mock-dynamodb-name"
    random_prefix      = "mockpfx"
    prefix_length      = 8
  }
}

dependency "id_generator" {
  config_path = "${get_repo_root()}/infra/terragrunt/stack-datastore/db/id-generator"
  mock_outputs = {
    full_uuid     = "00000000-0000-0000-0000-000000000000"
    formatted_uuid = "mock-uuid"
    uuid           = "00000000-0000-0000-0000-000000000000"
    id_type        = "generic"
  }
}

dependency "quota_generator" {
  config_path = find_in_parent_folders("db/quota-generator")
  mock_outputs = {
    adjusted_quota   = 100
    quota_percentage = 50
    raw_quota        = 100
    scaled_quota     = 100
    quota_type       = "concurrent_users"
  }
}

inputs = {
    # Generated name configuration
    generated_name = dependency.name_generator.outputs.full_resource_name
    generated_id   = dependency.id_generator.outputs.full_uuid
    generated_quota = dependency.quota_generator.outputs.adjusted_quota
    quota_type     = dependency.quota_generator.outputs.quota_type

    # Explicit table name (optional)
    table_name = "${local.table_name_prefix}-${dependency.name_generator.outputs.full_resource_name}"

    # Hash key configuration
    hash_key_name = "resource_id"
    hash_key_type = "S"

    # Additional configuration
    enable_point_in_time_recovery = true
    enable_ttl                    = true
    ttl_attribute                 = "TimeToExist"

    # Dynamic tagging strategy
    additional_tags = {
      Name           = "${local.table_name_prefix}-${dependency.name_generator.outputs.full_resource_name}"
      Environment    = local.environment
      GeneratedID    = dependency.id_generator.outputs.full_uuid
    ResourcePrefix = local.table_name_prefix
    QuotaType      = dependency.quota_generator.outputs.quota_type
    Provisioner    = "Terragrunt"
  }
}
