locals {
  # Generate a full resource name with the random prefix
  full_resource_name = var.name_template != "" ? format(var.name_template, random_string.resource_prefix.result) : random_string.resource_prefix.result
}
