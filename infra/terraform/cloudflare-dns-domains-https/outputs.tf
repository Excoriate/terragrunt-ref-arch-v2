output "is_enabled" {
  description = "Whether the module is enabled or not"
  value       = local.is_enabled
}

output "zone_settings" {
  description = "The Cloudflare zone settings configuration"
  value       = try(cloudflare_zone_settings_override.https_settings[0], null)
  sensitive   = true
}

output "security_headers" {
  description = "The configured security headers ruleset"
  value       = try(cloudflare_ruleset.transform_rule[0].id, null)
}