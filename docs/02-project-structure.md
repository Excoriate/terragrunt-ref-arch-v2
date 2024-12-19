# Project Structure

## Overview

The Terragrunt Reference Architecture is meticulously designed to provide a modular, scalable, and maintainable infrastructure-as-code (IaC) solution. This document provides an in-depth exploration of the repository's structure and its key components.

## Repository Layout

```
.
├── LICENSE
├── README.md
├── docs/                  # Documentation
├── infra/                 # Infrastructure core
│   ├── terraform/         # Terraform modules
│   │   ├── module-name/   # Individual module
│   │   │   ├── main.tf
│   │   │   ├── variables.tf
│   │   │   ├── outputs.tf
│   │   │   └── versions.tf
│   └── terragrunt/        # Terragrunt configuration
│       ├── _ENVS/         # Environment configurations
│       ├── _shared/       # Reusable components
│       ├── _templates/    # Generated file templates
│       ├── config.hcl     # Global configuration
│       └── stack-*/       # Infrastructure stacks
└── tools/                 # Infrastructure management tools
    └── infractl/          # Custom CLI tool
```

## Key Directories

### 1. Terraform Modules (`/infra/terraform/`)

Terraform modules are the building blocks of infrastructure, each representing a specific, reusable component.

#### Module Structure

```
infra/terraform/my-terraform-module/
├── README.md              # Module documentation
├── main.tf                # Core resource definitions
├── variables.tf           # Input variable definitions
├── outputs.tf             # Output value definitions
└── versions.tf            # Provider and version constraints
```

### 2. Terragrunt Configuration (`/infra/terragrunt/`)

The Terragrunt configuration provides a powerful, DRY (Don't Repeat Yourself) approach to infrastructure management.

#### Directory Breakdown

```
terragrunt/
├── _ENVS/                 # Environment-specific configurations
│   ├── base.yaml          # Base configuration
│   ├── local.yaml         # Local development settings
│   └── dev.yaml           # Development environment
├── _shared/               # Shared infrastructure components
│   └── _components/       # Reusable component configurations
│       ├── aws-vpc.hcl
│       └── eks-cluster.hcl
├── _templates/            # Configuration templates
│   ├── providers.tf.tpl
│   └── versions.tf.tpl
├── config.hcl             # Root configuration
├── stack-landing-zone/    # Infrastructure stack
│   ├── stack.hcl          # Stack-wide configuration
│   ├── dns/               # Layer
│   │   ├── layer.hcl
│   │   ├── dns-zone/      # Component
│   │   │   ├── component.hcl
│   │   │   └── terragrunt.hcl
│   │   └── dns-domains/   # Another component
│   │       ├── component.hcl
│   │       └── terragrunt.hcl
└── terragrunt.hcl         # Root Terragrunt configuration
```

##### Configuration Hierarchy

1. **`_ENVS/`**: Environment-specific configurations

   - `base.yaml`: Default settings
   - `local.yaml`, `dev.yaml`, etc.: Environment overrides

2. **`_shared/_components/`**: Reusable configuration snippets

   - Shared across different stacks and components
   - Promotes configuration reusability

3. **`stack-*/`**: Infrastructure stacks
   - Organized into layers and components
   - Supports modular infrastructure design

### 3. InfraCTL CLI Tool (`/tools/infractl/`)

A custom Go-based CLI for infrastructure management.

#### Tool Structure

```
infractl/
├── internal/              # Implementation details
│   ├── cfg/               # Configuration management
│   ├── controller/        # Business logic
│   ├── transformers/      # Data processing
│   └── tui/               # Text UI components
├── pkg/                   # Public packages
│   ├── envars/            # Environment variable handling
│   ├── logger/            # Logging utilities
│   ├── tg/                # Terragrunt integration
│   └── utils/             # Utility functions
└── main.go                # Entry point
```

## Development Tools

- `justfile`: Task automation
- `Makefile`: Build and deployment scripts
- `.editorconfig`: Consistent coding style
- GitHub Actions: CI/CD workflows
