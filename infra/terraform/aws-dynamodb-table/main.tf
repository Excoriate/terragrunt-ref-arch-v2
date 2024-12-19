locals {
  # Concatenate name and ID for unique table name
  table_name = var.table_name != "" ? var.table_name : "${var.generated_name}-${var.generated_id}"
}

resource "aws_dynamodb_table" "generated_table" {
  name           = local.table_name
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = var.hash_key_name

  attribute {
    name = var.hash_key_name
    type = var.hash_key_type
  }

  tags = merge(
    {
      Name            = local.table_name
      GeneratedName   = var.generated_name
      GeneratedID     = var.generated_id
      Quota           = var.generated_quota
      QuotaType       = var.quota_type
      ManagedBy       = "terraform-random-generator"
    },
    var.additional_tags
  )

  # Minimal, cost-effective configuration
  point_in_time_recovery {
    enabled = var.enable_point_in_time_recovery
  }

  # Minimal TTL to keep costs down
  ttl {
    attribute_name = var.ttl_attribute
    enabled        = var.enable_ttl
  }
}
