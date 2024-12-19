# Terragrunt Reference Architecture V2

## Overview

This reference architecture demonstrates how to build and manage complex infrastructure using Terragrunt, focusing on maintainability, reusability, and scalability. It provides a practical approach to organizing infrastructure code across multiple environments and cloud providers.

## Core Concepts

### Configuration Management

The architecture uses a centralized, flat YAML-based configuration system. Here's a comprehensive example from the `local.yaml` configuration:

```yaml
# Root Configuration
config:
  version: "1.0.0"
  last_updated: "2024-01-15"
  description: "Centralized configuration for local environment"

# Git Configuration
git:
  base_url: "git::git@github.com:"

# Product Identification
product:
  name: ref-arch
  version: 0.0.1-local
  description: "Reference architecture for cloud infrastructure - demo environment"
  use_as_stack_tags: true

# Infrastructure as Code Configuration
iac:
  versions:
    terraform_version_default: "1.9.8"
    terragrunt_version_default: "0.62.1"
  remote_state:
    s3:
      bucket: ${TF_STATE_BUCKET}
      lock_table: ${TF_STATE_LOCK_TABLE}
      region: us-east-1

# Providers Configuration
providers:
  aws:
    config:
      access_key_id: ${AWS_ACCESS_KEY_ID:-secrets.aws.access_key_id}
      secret_access_key: ${AWS_SECRET_ACCESS_KEY:-secrets.aws.secret_access_key}
      region: us-east-1
    version_constraints:
      - name: aws
        source: "hashicorp/aws"
        required_version: "5.80.0"
        enabled: true

# Secrets Management
secrets:
  aws:
    access_key: ${AWS_ACCESS_KEY_ID}
    secret_key: ${AWS_SECRET_ACCESS_KEY}
```

Key Configuration Properties:

- **`config`**: Global version and metadata
- **`git`**: Repository configuration
- **`product`**: Product identification and tagging
- **`iac`**: Infrastructure as Code settings
- **`providers`**: Cloud provider configurations
- **`secrets`**: Secure secret management

### Stack Organization

Infrastructure is organized into a hierarchical structure:

1. **Stacks**: Logical groupings of infrastructure

   - Represent a complete, deployable unit
   - Composed of multiple layers
   - Example: `stack-platform`, `stack-datastore`

2. **Layers**: Functional groupings within stacks

   - Organize related infrastructure components
   - Provide modular organization
   - Example: `networking`, `compute`, `security`

3. **Components**: Individual infrastructure resources
   - Smallest deployable units
   - Inherit from shared configurations
   - Example: `vpc`, `eks-cluster`, `dynamodb-table`

Detailed Stack Structure:

```
stack-platform/                 # Stack
├── stack.hcl                  # Stack-wide configuration
├── networking/                # Layer
│   ├── layer.hcl             # Layer configuration
│   ├── vpc/                  # Component
│   │   ├── terragrunt.hcl    # Terragrunt configuration
│   │   └── component.hcl     # Component settings
│   └── subnets/              # Another component
└── compute/                   # Another layer
    └── eks/                  # EKS component
```

### Stack Configuration Example

From a given configuration file, E.g. `local.yaml`, here's a stack configuration:

```yaml
stacks:
  - name: stack-datastore
    tags:
      stack_purpose: demo-resource-generation
    layers:
      - name: db
        tags:
          layer_type: databases
        components:
          - name: id-generator
            providers:
              - "random"
          - name: aws-dynamodb-table
            providers:
              - "aws"
```

### Provider Integration

Multiple cloud providers are supported through a unified configuration:

```yaml
providers:
  aws:
    config:
      region: us-east-1
  random:
    config: {} # No specific configuration needed
    version_constraints:
      - name: random
        source: "hashicorp/random"
        required_version: "3.6.3"
        enabled: true
```
