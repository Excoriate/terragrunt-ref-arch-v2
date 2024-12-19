locals {
  is_enabled = var.is_enabled
}

resource "namecheap_domain_records" "domain" {
  count  = local.is_enabled ? 1 : 0
  domain = var.domain
  mode   = var.mode

  nameservers = var.nameservers
}