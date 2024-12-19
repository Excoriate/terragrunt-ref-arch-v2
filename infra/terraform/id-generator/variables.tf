variable "id_type" {
  description = "Type of identifier being generated (e.g., 'deployment', 'trace', 'session')"
  type        = string
  default     = "generic"
}

variable "generation_key" {
  description = "Optional key to trigger UUID regeneration when specific attributes change"
  type        = string
  default     = null
}

variable "uuid_format" {
  description = "Optional format string for the UUID (e.g., 'custom-%s')"
  type        = string
  default     = ""
  validation {
    condition     = can(format(var.uuid_format, "test"))
    error_message = "Invalid UUID format string. Must be a valid format specifier."
  }
}

variable "uuid_prefix" {
  description = "Optional prefix to prepend to the generated UUID"
  type        = string
  default     = ""
  validation {
    condition     = length(var.uuid_prefix) <= 32
    error_message = "UUID prefix must be 32 characters or less."
  }
}

variable "uuid_suffix" {
  description = "Optional suffix to append to the generated UUID"
  type        = string
  default     = ""
  validation {
    condition     = length(var.uuid_suffix) <= 32
    error_message = "UUID suffix must be 32 characters or less."
  }
}
