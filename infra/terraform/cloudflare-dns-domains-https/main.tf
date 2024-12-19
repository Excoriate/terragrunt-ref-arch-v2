locals {
  is_enabled = var.is_enabled
}

resource "cloudflare_zone_settings_override" "https_settings" {
  count   = local.is_enabled ? 1 : 0
  zone_id = var.zone_id

  settings {
    ssl                      = "strict"
    always_use_https        = "on"
    min_tls_version        = var.min_tls_version
    tls_1_3                = "on"
    automatic_https_rewrites = "on"
    http3                   = "on"
    browser_check          = "on"
    challenge_ttl          = 2700
    security_level         = "medium"
    brotli                 = "on"
    early_hints            = "on"
    opportunistic_encryption = "on"
    universal_ssl          = "on"
    tls_client_auth       = "off"
  }
}

resource "cloudflare_page_rule" "force_https" {
  count    = local.is_enabled ? 1 : 0
  zone_id  = var.zone_id
  target   = "*.${var.domain}/*"
  priority = 1
  status   = "active"

  actions {
    ssl = "strict"
    cache_level = "aggressive"
    browser_cache_ttl = "14400"
    edge_cache_ttl = 7200
  }
}

resource "cloudflare_ruleset" "transform_rule" {
  count       = local.is_enabled ? 1 : 0
  zone_id     = var.zone_id
  name        = "HTTPS Transform Rules"
  description = "Force HTTPS and security headers"
  kind        = "zone"
  phase       = "http_response_headers_transform"

  rules {
    action = "rewrite"
    action_parameters {
      headers {
        name      = "Strict-Transport-Security"
        value     = "max-age=31536000; includeSubDomains; preload"
        operation = "set"
      }
      headers {
        name      = "X-Content-Type-Options"
        value     = "nosniff"
        operation = "set"
      }
      headers {
        name      = "X-Frame-Options"
        value     = "DENY"
        operation = "set"
      }
      headers {
        name      = "X-XSS-Protection"
        value     = "1; mode=block"
        operation = "set"
      }
    }
    expression  = "true"
    description = "Add security headers"
    enabled     = true
  }
}

resource "cloudflare_zone_dnssec" "dnssec" {
  count   = local.is_enabled ? 1 : 0
  zone_id = var.zone_id
}