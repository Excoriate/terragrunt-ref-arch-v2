# Terragrunt Deployment Blueprint
DEFAULT_ENV := "local"

# Path configurations
TERRAGRUNT_DIR := "./infra/terragrunt"
ENVS_DIR := "./infra/terragrunt/_ENVS"

# 🌍 Load environment variables from .env file
set dotenv-load

# 🐚 Set the default shell to bash with error handling
set shell := ["bash", "-uce"]

# 📋 List all available recipes
default:
    @just --list

# 📦 Install all dependencies using Turborepo
install:
    bunx turbo run install


# 🗑️ Remove all .DS_Store files
clean-ds:
    find . -name '.DS_Store' -type f -delete

build-cli:
    @echo "Building InfraCTL CLI 👨🏻‍💻"
    @cd tools/infractl && go build -o target/infractl main.go

run-cli *ARGS: build-cli
    @echo "Running InfraCTL CLI 👨🏻‍💻 with args: {{ARGS}}"
    @./tools/infractl/target/infractl {{ARGS}}

tg-clean:
    @echo "Cleaning Terragrunt cache for all environments"
    @cd infra/terragrunt && find . -type d -name ".terragrunt-cache" -exec rm -rf {} +

tg-plan stack='stack-datastore' layer='db' component='quota-generator':
    @echo "Planning Terragrunt blueprint for environment"
    @cd infra/terragrunt/{{stack}}/{{layer}}/{{component}} && terragrunt plan

tg-datastore-db cmd='plan':
    @echo "Planning Terragrunt blueprint for environment"
    @cd infra/terragrunt/stack-datastore/db && terragrunt run-all {{cmd}} --terragrunt-non-interactive

tg-plan-demo component='quota-generator':
    @echo "Planning Terragrunt blueprint for environment"
    @cd infra/terragrunt/stack-datastore/db/{{component}} && terragrunt plan
