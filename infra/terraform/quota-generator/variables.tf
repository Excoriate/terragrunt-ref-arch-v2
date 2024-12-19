variable "min_quota" {
  description = "Minimum value for the random quota generation"
  type        = number
  default     = 10
  validation {
    condition     = var.min_quota >= 0
    error_message = "Minimum quota must be a non-negative number."
  }
}

variable "max_quota" {
  description = "Maximum value for the random quota generation"
  type        = number
  default     = 100
  validation {
    condition     = var.max_quota > var.min_quota
    error_message = "Maximum quota must be greater than the minimum quota."
  }
}

variable "quota_type" {
  description = "Type of quota being generated (e.g., 'concurrent_users', 'rate_limit', 'storage')"
  type        = string
  default     = "generic"
}

variable "generation_seed" {
  description = "Optional seed to control quota generation consistency"
  type        = string
  default     = null
}

variable "scale_factor" {
  description = "Optional factor to scale the generated quota"
  type        = number
  default     = 1.0
  validation {
    condition     = var.scale_factor > 0
    error_message = "Scale factor must be a positive number."
  }
}

variable "adjustment_type" {
  description = "Type of numerical adjustment to apply to the quota"
  type        = string
  default     = "none"
  validation {
    condition     = contains(["none", "floor", "ceil"], var.adjustment_type)
    error_message = "Adjustment type must be 'none', 'floor', or 'ceil'."
  }
}
