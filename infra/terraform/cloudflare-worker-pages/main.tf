terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4.36.0"
    }
  }
}

provider "cloudflare" {
  # Credentials pulled from environment variables:
  # CLOUDFLARE_API_TOKEN or CLOUDFLARE_EMAIL and CLOUDFLARE_API_KEY
}

resource "cloudflare_pages_project" "astro_site" {
  account_id        = var.cloudflare_account_id
  name             = var.project_name
  production_branch = var.production_branch

  build_config {
    build_command       = "npm run build"
    destination_dir     = "dist"
    root_dir           = var.root_dir
  }

  source {
    type = "github"
    config {
      owner                         = var.github_owner
      repo_name                     = var.github_repo_name
      production_branch            = var.production_branch
      pr_comments_enabled          = true
      deployments_enabled          = true
      production_deployment_enabled = true
      preview_deployment_setting   = "all"
      preview_branch_includes      = ["dev", "staging"]
    }
  }

  deployment_configs {
    preview {
      environment_variables = {
        NODE_VERSION = "20"
      }
      compatibility_date = "2024-01-01"
    }
    production {
      environment_variables = {
        NODE_VERSION = "20"
      }
      compatibility_date = "2024-01-01"
    }
  }
}
