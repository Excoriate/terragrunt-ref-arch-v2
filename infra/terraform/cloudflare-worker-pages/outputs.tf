output "project_name" {
  description = "Name of the created Pages project"
  value       = cloudflare_pages_project.astro_site.name
}

output "project_id" {
  description = "ID of the created Pages project"
  value       = cloudflare_pages_project.astro_site.id
}

output "project_subdomain" {
  description = "The *.pages.dev subdomain for the project"
  value       = "${cloudflare_pages_project.astro_site.name}.pages.dev"
}
