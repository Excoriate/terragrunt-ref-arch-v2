locals {
  # Calculate percentage of quota range
  quota_percentage = (random_integer.resource_quota.result - var.min_quota) / (var.max_quota - var.min_quota) * 100

  # Optional scaling or adjustment of the quota
  scaled_quota = var.scale_factor > 0 ? random_integer.resource_quota.result * var.scale_factor : random_integer.resource_quota.result

  # Optional quota adjustment based on predefined rules
  adjusted_quota = var.adjustment_type == "floor" ? floor(local.scaled_quota) : var.adjustment_type == "ceil" ? ceil(local.scaled_quota) : local.scaled_quota
}
