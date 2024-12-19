resource "random_integer" "resource_quota" {
  min = var.min_quota
  max = var.max_quota

  keepers = {
    quota_type        = var.quota_type
    generation_seed   = var.generation_seed
  }
}
