# Terragrunt Reference Architecture V2

## Overview

This repository provides a highly DRY (Don't Repeat Yourself) Terragrunt reference architecture designed for scalable infrastructure deployments across multiple environments. It offers a centralized, flexible, and reproducible infrastructure-as-code (IaC) approach.

## Key Features

- **Centralized Configuration**: All configuration is managed through YAML files in the `_ENVS` folder, allowing for easy environment-specific overrides
- **DRY Terragrunt Architecture**: Minimizes code duplication while maintaining flexibility
- **Multi-Provider Support**: Built-in support for multiple cloud providers with clean configuration management
- **Environment-Specific Configurations**: Easy adaptation across different deployment environments
- **Custom CLI Tool**: Includes `infractl` for configuration validation and management
- **Hierarchical Stack Management**: Organized structure of stacks, layers, and components

## Why This Architecture?

### Solving Common IaC Challenges

1. **Configuration Management**:

   - Centralized YAML-based configuration
   - Environment-specific overrides
   - Clean separation of concerns

2. **Code Reusability**:

   - Shared components across stacks
   - DRY configuration principles
   - Modular design

3. **Scalability**:
   - Support for multiple environments
   - Easy addition of new components
   - Flexible provider configuration

### Benefits for Teams

- **Reduced Complexity**: Clear separation of configuration and implementation
- **Improved Maintainability**: Centralized configuration management
- **Enhanced Collaboration**: Standardized structure and conventions
- **Better Security**: Proper secrets management and environment separation

## Getting Started

To begin using this reference architecture:

1. Review the [Project Structure](02-project-structure.md)
2. Understand the [Configuration System](03-configuration-system.md)
3. Learn about the [CLI Tool](04-infractl-cli.md)
4. Explore [Stack Management](05-stack-management.md)

## Prerequisites

- Terraform (>= 1.9.8)
- Terragrunt (>= 0.62.1)
- Go (>= 1.20) for CLI tool
- AWS CLI (configured with appropriate credentials)
- Git

## Next Steps

Continue to [Project Structure](02-project-structure.md) to understand the repository organization and key components.
