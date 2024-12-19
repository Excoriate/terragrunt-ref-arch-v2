terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4.0"
    }
    namecheap = {
      source  = "namecheap/namecheap"
      version = "~> 2.0"
    }
  }
  required_version = ">= 1.0"
}