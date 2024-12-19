# Astro Site Cloudflare Pages Configuration

This Terragrunt configuration manages the Cloudflare Pages deployment for the Astro site.

## Prerequisites

- Terragrunt installed
- Terraform >= 1.0.0
- Environment variables set:
  - CLOUDFLARE_API_TOKEN or (CLOUDFLARE_EMAIL and CLOUDFLARE_API_KEY)
  - CLOUDFLARE_ACCOUNT_ID
  - CLOUDFLARE_PROJECT_NAME

## Usage
