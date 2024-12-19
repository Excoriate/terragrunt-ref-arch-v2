terraform {
  required_version = "~> 1.9.8"
  
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4.36.0"
    }
  }
}
