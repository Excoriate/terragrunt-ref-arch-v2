output "domain_records" {
  value       = try(namecheap_domain_records.domain[0], null)
  description = "The created domain records"
}

output "is_enabled" {
  value       = local.is_enabled
  description = "Whether the module is enabled or not"
}