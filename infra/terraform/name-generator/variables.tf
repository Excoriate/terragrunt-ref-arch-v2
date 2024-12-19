variable "prefix_length" {
  description = "Length of the random prefix to generate"
  type        = number
  default     = 8
  validation {
    condition     = var.prefix_length >= 4 && var.prefix_length <= 32
    error_message = "Prefix length must be between 4 and 32 characters."
  }
}

variable "include_special_chars" {
  description = "Whether to include special characters in the random string"
  type        = bool
  default     = false
}

variable "include_uppercase" {
  description = "Whether to include uppercase letters in the random string"
  type        = bool
  default     = true
}

variable "include_numeric" {
  description = "Whether to include numeric characters in the random string"
  type        = bool
  default     = true
}

variable "prefix_type" {
  description = "Type of prefix to generate (e.g., 'dev', 'prod', 'test')"
  type        = string
  default     = "generic"
}

variable "name_template" {
  description = "Optional template for generating full resource names. Use '%s' as placeholder for random string"
  type        = string
  default     = "%s-resource"
}
