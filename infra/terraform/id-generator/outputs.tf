output "uuid" {
  description = "The generated UUID"
  value       = random_uuid.resource_id.result
}

output "formatted_uuid" {
  description = "The UUID after optional formatting"
  value       = local.formatted_uuid
}

output "full_uuid" {
  description = "The complete UUID with optional prefix and suffix"
  value       = local.suffixed_uuid
}

output "id_type" {
  description = "The type of identifier generated"
  value       = var.id_type
}
