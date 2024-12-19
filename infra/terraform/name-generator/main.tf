resource "random_string" "resource_prefix" {
  length  = var.prefix_length
  special = var.include_special_chars
  upper   = var.include_uppercase
  numeric = var.include_numeric

  keepers = {
    # Ensure regeneration only happens when specific attributes change
    prefix_type = var.prefix_type
  }
}
