variable "is_enabled" {
  type        = bool
  description = "Whether this module should create resources or not"
  default     = false
}

# variable "namecheap_api_user" {
#   type        = string
#   description = "Namecheap API username"
# }

# variable "namecheap_api_key" {
#   type        = string
#   description = "Namecheap API key"
# }

# variable "namecheap_username" {
#   type        = string
#   description = "Namecheap username"
# }

# variable "namecheap_client_ip" {
#   type        = string
#   description = "IP address allowed to access the Namecheap API"
#   default     = null
# }

variable "domain" {
  type        = string
  description = "The domain name to update NS records for"
}

variable "nameservers" {
  type        = list(string)
  description = "List of nameservers to set for the domain"
}

variable "mode" {
  type        = string
  description = "The mode to use for the domain records"
  default     = "MERGE"
}