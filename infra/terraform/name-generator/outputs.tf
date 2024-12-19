output "random_prefix" {
  description = "The generated random string prefix"
  value       = random_string.resource_prefix.result
}

output "full_resource_name" {
  description = "The full generated resource name using the template"
  value       = local.full_resource_name
}

output "prefix_length" {
  description = "The length of the generated random prefix"
  value       = length(random_string.resource_prefix.result)
}
