output "raw_quota" {
  description = "The raw generated quota value"
  value       = random_integer.resource_quota.result
}

output "scaled_quota" {
  description = "The quota after optional scaling"
  value       = local.scaled_quota
}

output "adjusted_quota" {
  description = "The final adjusted quota value"
  value       = local.adjusted_quota
}

output "quota_percentage" {
  description = "Percentage of the quota within the specified range"
  value       = local.quota_percentage
}

output "quota_type" {
  description = "The type of quota generated"
  value       = var.quota_type
}
