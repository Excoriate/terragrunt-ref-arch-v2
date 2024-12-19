variable "cloudflare_account_id" {
  description = "Cloudflare Account ID"
  type        = string
}

variable "project_name" {
  description = "Name of the Pages project"
  type        = string
}

variable "production_branch" {
  description = "Git branch to use for production deployments"
  type        = string
  default     = "main"
}

variable "root_dir" {
  description = "Directory where the build command should be run"
  type        = string
  default     = "/"
}

variable "github_owner" {
  description = "GitHub repository owner"
  type        = string
}

variable "github_repo_name" {
  description = "GitHub repository name"
  type        = string
}
