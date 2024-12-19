variable "table_name" {
  description = "Optional explicit table name. If not provided, will be generated from name and ID"
  type        = string
  default     = ""
}

variable "generated_name" {
  description = "Generated name to be used in table naming or tagging"
  type        = string
  default     = ""
}

variable "generated_id" {
  description = "Generated unique identifier to be used in table naming or tagging"
  type        = string
  default     = ""
}

variable "generated_quota" {
  description = "Generated quota value to be used in table tagging"
  type        = number
  default     = null
}

variable "quota_type" {
  description = "Type of quota being applied"
  type        = string
  default     = "unspecified"
}

variable "additional_tags" {
  description = "Additional tags to be applied to the DynamoDB table"
  type        = map(string)
  default     = {}
}

variable "hash_key_name" {
  description = "Name of the hash key attribute"
  type        = string
  default     = "id"
  validation {
    condition     = length(var.hash_key_name) > 0 && length(var.hash_key_name) <= 255
    error_message = "Hash key name must be between 1 and 255 characters."
  }
}

variable "hash_key_type" {
  description = "Type of the hash key attribute"
  type        = string
  default     = "S"
  validation {
    condition     = contains(["S", "N", "B"], var.hash_key_type)
    error_message = "Hash key type must be 'S' (string), 'N' (number), or 'B' (binary)."
  }
}

variable "enable_point_in_time_recovery" {
  description = "Enable point-in-time recovery for the DynamoDB table"
  type        = bool
  default     = false
}

variable "enable_ttl" {
  description = "Enable Time-To-Live for the DynamoDB table"
  type        = bool
  default     = false
}

variable "ttl_attribute" {
  description = "Attribute name for Time-To-Live"
  type        = string
  default     = "TimeToExist"
  validation {
    condition     = length(var.ttl_attribute) > 0 && length(var.ttl_attribute) <= 255
    error_message = "TTL attribute name must be between 1 and 255 characters."
  }
}
