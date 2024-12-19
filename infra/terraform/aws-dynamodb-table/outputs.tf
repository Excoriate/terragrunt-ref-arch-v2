output "table_name" {
  description = "The full generated DynamoDB table name"
  value       = aws_dynamodb_table.generated_table.name
}

output "table_arn" {
  description = "The ARN of the generated DynamoDB table"
  value       = aws_dynamodb_table.generated_table.arn
}

output "generated_name" {
  description = "The generated name from the name generator"
  value       = module.table_name.full_resource_name
}

output "generated_id" {
  description = "The generated UUID from the ID generator"
  value       = module.table_id.full_uuid
}

output "generated_quota" {
  description = "The generated quota from the quota generator"
  value       = module.table_quota.adjusted_quota
}