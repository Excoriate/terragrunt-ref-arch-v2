variable "is_enabled" {
  type        = bool
  description = "Whether this module should create resources or not"
  default     = true
}

variable "zone_id" {
  type        = string
  description = "The Cloudflare Zone ID"
}

variable "domain" {
  type        = string
  description = "The domain name to configure HTTPS for"
}

variable "min_tls_version" {
  type        = string
  description = "Minimum TLS version to allow"
  default     = "1.2"

  validation {
    condition     = contains(["1.0", "1.1", "1.2", "1.3"], var.min_tls_version)
    error_message = "min_tls_version must be one of: 1.0, 1.1, 1.2, 1.3"
  }
}

variable "enable_waf" {
  description = "Enable WAF (requires Enterprise plan)"
  type        = bool
  default     = false
}

variable "enable_polish" {
  description = "Enable Polish image optimization (requires Pro plan or higher)"
  type        = bool
  default     = false
}

variable "enable_webp" {
  description = "Enable WebP optimization (requires Pro plan or higher)"
  type        = bool
  default     = false
}

variable "enable_privacy_pass" {
  description = "Enable Privacy Pass (requires Enterprise plan)"
  type        = bool
  default     = false
}

variable "enable_true_client_ip" {
  description = "Enable True Client IP (requires Enterprise plan)"
  type        = bool
  default     = false
}

variable "enable_rate_limiting" {
  description = "Enable rate limiting protection"
  type        = bool
  default     = true
}

variable "rate_limit_api_requests" {
  description = "Maximum API requests per minute per IP/datacenter"
  type        = number
  default     = 100
}

variable "rate_limit_api_timeout" {
  description = "Duration in seconds to challenge after exceeding API rate limit"
  type        = number
  default     = 600
}

variable "rate_limit_post_requests" {
  description = "Maximum POST requests per 10 seconds per IP"
  type        = number
  default     = 50
}

variable "rate_limit_post_timeout" {
  description = "Duration in seconds to challenge after exceeding POST rate limit"
  type        = number
  default     = 300
}

variable "rate_limit_global_requests" {
  description = "Maximum global requests per 30 seconds per IP/path"
  type        = number
  default     = 200
}

variable "rate_limit_global_timeout" {
  description = "Duration in seconds to challenge after exceeding global rate limit"
  type        = number
  default     = 600
}