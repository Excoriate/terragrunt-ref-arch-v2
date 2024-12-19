# InfraCTL CLI Tool

## Overview

The InfraCTL CLI tool is a custom-built utility designed to manage and validate infrastructure configurations, wrapping Terragrunt commands with additional functionality for configuration management and validation.

## Installation

The CLI tool is located in the `tools/infractl` directory and can be built using Go:

```bash
cd tools/infractl
go build -o infractl
```

## Command Structure

### Basic Command Format

```bash
infractl <command> --stack=<stack_name> --layer=<layer_name> --component=<component_name> --target-env=<environment>
```

### Available Commands

1. **Plan**

   ```bash
   infractl plan --stack=landingzone --layer=dns --component=dns-zone --target-env=local
   ```

2. **Apply**

   ```bash
   infractl apply --stack=webapp --layer=web --component=web-app --target-env=staging
   ```

3. **Destroy**

   ```bash
   infractl destroy --stack=database --layer=storage --component=dynamodb --target-env=dev
   ```

4. **Validate**
   ```bash
   infractl validate --target-env=production
   ```

## Configuration Processing

### 1. Configuration Loading

- Reads `base.yaml` from `_ENVS` directory
- Loads environment-specific YAML file
- Merges configurations

### 2. Validation Steps

- Schema validation
- Environment variable resolution
- Provider configuration validation
- Stack structure validation
- Terragrunt binary verification

### 3. Configuration Compilation

- Generates JSON configuration
- Caches in `.infractl-cache`
- Sets environment variables

## Cache Management

### Cache Directory

- Location: `infra/.infractl-cache`
- Permissions: 0700
- Git ignored by default

### Cache Files

- Format: `config-compiled-{env}-{ID}.json`
- Permissions: 0600
- Environment Variables:
  - `INFRACTL_CONFIG_FILE`
  - `INFRACTL_CONFIG_FILE_PATH`

## Validation Features

### 1. Sanity Checks

- YAML file existence
- YAML syntax validation
- Schema validation
- Terragrunt version verification

### 2. Structural Validation

Validates mandatory sections:

- config
- git
- product
- environment
- stacks
- providers
- iac
- secrets (optional)

### 3. Provider Validation

- Configuration presence
- Version constraints
- Required attributes
- Secret references

## Environment Variable Handling

### 1. Variable Resolution

- Direct environment variables
- Default values
- Secret references

### 2. Secret Management

- Environment variable injection
- Secret reference resolution
- Provider-specific secrets

## Best Practices

### 1. Command Usage

- Always specify target environment
- Use meaningful stack/layer/component names
- Validate before applying

### 2. Configuration Management

- Keep configurations DRY
- Use environment variables for secrets
- Maintain clear documentation

### 3. Error Handling

- Check validation output
- Review compiled configuration
- Monitor cache directory

## Common Workflows

### 1. New Component Deployment

```bash
# 1. Validate configuration
infractl validate --target-env=dev

# 2. Plan changes
infractl plan --stack=mystack --layer=mylayer --component=mycomponent --target-env=dev

# 3. Apply changes
infractl apply --stack=mystack --layer=mylayer --component=mycomponent --target-env=dev
```

### 2. Stack-wide Operations

```bash
# Plan entire stack
infractl plan --stack=mystack --target-env=dev

# Apply entire stack
infractl apply --stack=mystack --target-env=dev
```

## Troubleshooting

### Common Issues

1. **Configuration Errors**

   - Check YAML syntax
   - Verify environment variables
   - Validate provider configuration

2. **Cache Issues**

   - Clear `.infractl-cache`
   - Check file permissions
   - Verify environment variables

3. **Terragrunt Integration**
   - Check Terragrunt version
   - Verify binary location
   - Review command output

## Next Steps

Continue to [Stack Management](05-stack-management.md) to learn about managing infrastructure stacks.
