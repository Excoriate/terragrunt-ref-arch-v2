locals {
#   # ---------------------------------------------------------------------------------------------------------------------
#   # PROVIDER CONFIGURATIONS
#   # This section centralizes the configuration of Terraform providers, such as Cloudflare and AWS, using Terraform's
#   # heredoc syntax for inline definition. This approach allows for dynamic, environment-specific configuration of
#   # providers through environment variables, enhancing the flexibility and security of provider setups. Direct use of
#   # heredoc syntax within the Terragrunt configuration eliminates the need for external template files, streamlining
#   # the codebase and simplifying the management of provider configurations.
#   #
#   # Each provider configuration includes:
#   # - `content`: The Terraform configuration for the provider, including authentication details and any other
#   #              provider-specific settings. Sensitive information, such as API keys, is securely sourced from
#   #              environment variables.
#   #
#   # This modular and dynamic approach to configuring providers supports best practices in security and infrastructure
#   # code management, enabling selective provider use and environment-specific configurations without altering the
#   # core codebase.
#   # ---------------------------------------------------------------------------------------------------------------------
#   providers_globals_cfg = read_terragrunt_config("${get_terragrunt_dir()}/../../../_providers/namecheap.hcl")

#   providers = {
#     cloudflare = {
#       content = <<EOF
# provider "namecheap" {
#   api_key     = "${local.providers_globals_cfg.locals.namecheap.api_key}"
#   user_name    = "${local.providers_globals_cfg.locals.namecheap.username}"
#   api_user    = "${local.providers_globals_cfg.locals.namecheap.api_user}"
#   use_sandbox = "${local.providers_globals_cfg.locals.namecheap.use_sandbox}"
# }
# EOF
#     },
#     # Additional providers can be added here following the same pattern.
#   }

#   # ---------------------------------------------------------------------------------------------------------------------
#   # PROVIDERS CONTENT
#   # Generate the providers' configuration content only for enabled providers.
#   # ---------------------------------------------------------------------------------------------------------------------
#   providers_content = [
#     for provider, details in local.providers : details.content]
}