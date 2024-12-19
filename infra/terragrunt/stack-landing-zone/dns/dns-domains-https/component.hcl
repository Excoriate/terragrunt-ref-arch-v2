locals {
  # ---------------------------------------------------------------------------------------------------------------------
  # COMPONENT CONFIGURATION
  # ---------------------------------------------------------------------------------------------------------------------
  is_enabled = true
  name       = "cloudflare-dns-domains-https"

  tags = {
    Name = "cloudflare-dns-domains-https"
    ArchitectureType = "component"
  }
}