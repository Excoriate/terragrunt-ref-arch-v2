resource "random_uuid" "resource_id" {
  # Keepers allow regeneration based on specific attribute changes
  keepers = {
    id_type        = var.id_type
    generation_key = var.generation_key
  }
}
