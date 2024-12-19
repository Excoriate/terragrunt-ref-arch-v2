# Configuration System

## Overview

The configuration system is built around YAML files in the `_ENVS` directory, providing a centralized and flexible way to manage infrastructure configurations across different environments.

## Configuration Structure

### Root Configuration

```yaml
config:
  version: "1.0.0"
  last_updated: "2024-01-15"
  description: "Environment configuration"
```

### Git Configuration

```yaml
git:
  base_url: "git::git@github.com:"
  terraform_modules_local_path: "terraform"
```

### Product Configuration

```yaml
product:
  name: "my-product"
  version: "0.0.1-local"
  description: "Product description"
  use_as_stack_tags: true
```

### Infrastructure as Code Configuration

```yaml
iac:
  versions:
    terraform_version_default: "1.9.8"
    terragrunt_version_default: "0.62.1"
  remote_state:
    s3:
      bucket: "terraform-state-bucket"
      lock_table: "terraform-state-lock"
      region: "us-east-1"
```

### Stacks Configuration

```yaml
stacks:
  - name: stack-example
    tags:
      stack_tag: "value"
    layers:
      - name: layer-name
        tags:
          layer_tag: "value"
        components:
          - name: component-name
            providers:
              - "aws"
            tags:
              component_tag: "value"
            inputs:
              variable_name: "value"
```

### Providers Configuration

```yaml
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
```

### Secrets Configuration

```yaml
secrets:
  aws:
    access_key: ${AWS_ACCESS_KEY_ID}
    secret_key: ${AWS_SECRET_ACCESS_KEY}
```

## Environment Variable Resolution

The configuration system supports sophisticated environment variable resolution:

1. **Direct References**: `${ENV_VAR}`
2. **Default Values**: `${ENV_VAR:-default}`
3. **Secret References**: `${ENV_VAR:-secrets.provider.key}`

## Configuration Inheritance

1. **Base Configuration** (`base.yaml`)

   - Contains default settings
   - Shared across environments
   - Foundation for all deployments

2. **Environment Overrides** (`<env>.yaml`)
   - Environment-specific settings
   - Overrides base configuration
   - Maintains environment isolation

## Configuration Processing

1. **Loading**

   - Read base configuration
   - Load environment-specific configuration
   - Merge configurations

2. **Validation**

   - Schema validation
   - Environment variable resolution
   - Provider configuration validation
   - Stack structure validation

3. **Compilation**
   - Generate final JSON configuration
   - Cache compiled configuration
   - Make available to Terragrunt

## Best Practices

### Environment Variables

1. **Sensitive Data**

   - Use environment variables for secrets
   - Never commit sensitive values
   - Utilize secret references

2. **Default Values**
   - Provide sensible defaults
   - Document required variables
   - Use clear naming conventions

### Configuration Management

1. **Version Control**

   - Track configuration changes
   - Review configuration updates
   - Maintain change history

2. **Documentation**

   - Document all configuration options
   - Explain environment requirements
   - Provide example configurations

3. **Testing**
   - Validate configurations before deployment
   - Test environment variable resolution
   - Verify provider configurations

## Common Patterns

### Stack Configuration

```yaml
stacks:
  - name: stack-networking
    layers:
      - name: vpc
        components:
          - name: main-vpc
            providers:
              - "aws"
            inputs:
              cidr_block: "10.0.0.0/16"
```

### Provider Configuration

```yaml
providers:
  aws:
    config:
      region: ${AWS_REGION:-us-east-1}
    version_constraints:
      - name: aws
        source: "hashicorp/aws"
        required_version: "5.80.0"
```

## Next Steps

Continue to [InfraCTL CLI](04-infractl-cli.md) to learn about the custom CLI tool for managing infrastructure.
