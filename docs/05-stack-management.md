# Stack Management

## Overview

Stacks are the highest-level organizational unit in the infrastructure, representing complete, deployable sets of infrastructure components. This guide explains how to work with stacks effectively.

## Stack Structure

### Hierarchy

```
stack-example/
├── stack.hcl           # Stack-level configuration
└── layer-name/         # Logical grouping of components
    ├── layer.hcl       # Layer-level configuration
    └── component-name/ # Individual infrastructure component
        ├── component.hcl
        └── terragrunt.hcl
```

### Configuration Files

1. **stack.hcl**

   - Stack-wide settings
   - Common tags
   - Shared variables

2. **layer.hcl**

   - Layer-specific configuration
   - Component grouping
   - Layer-level dependencies

3. **component.hcl**
   - Component-specific settings
   - Resource configuration
   - Component dependencies

## Stack Configuration

### YAML Configuration

```yaml
stacks:
  - name: stack-datastore
    tags:
      stack_purpose: "data-storage"
    layers:
      - name: db
        tags:
          layer_type: "database"
        components:
          - name: dynamodb-table
            providers:
              - "aws"
            tags:
              component_type: "nosql"
            inputs:
              table_name: "my-table"
```

### Stack-Level Settings

1. **Tags**

   - Applied to all resources
   - Environment identification
   - Cost tracking

2. **Variables**
   - Shared across layers
   - Environment-specific values
   - Common configuration

## Layer Management

### Layer Organization

- Group related components
- Maintain clear boundaries
- Define dependencies

### Layer Configuration

```hcl
# layer.hcl
locals {
  layer_name = "database"
  layer_tags = {
    layer = local.layer_name
    environment = "production"
  }
}
```

## Component Management

### Component Structure

1. **Configuration**

   ```hcl
   # component.hcl
   locals {
     component_name = "dynamodb-table"
     component_tags = {
       component = local.component_name
     }
   }
   ```

2. **Terragrunt Configuration**

   ```hcl
   # terragrunt.hcl
   include "root" {
     path = find_in_parent_folders()
   }

   include "component" {
     path = "${get_repo_root()}/infra/terragrunt/_shared/_components/aws-dynamodb-table.hcl"
   }
   ```

## Dependencies

### Between Components

```hcl
dependency "vpc" {
  config_path = "../../networking/vpc"
}

inputs = {
  vpc_id = dependency.vpc.outputs.vpc_id
}
```

### Between Layers

- Use explicit dependencies
- Maintain clear documentation
- Consider deployment order

## Best Practices

### 1. Stack Organization

- Use meaningful names
- Group related components
- Maintain clear documentation

### 2. Layer Management

- Logical grouping
- Clear boundaries
- Explicit dependencies

### 3. Component Design

- Single responsibility
- Clear interfaces
- Proper documentation

## Common Patterns

### 1. Landing Zone Stack

```yaml
stacks:
  - name: stack-landing-zone
    layers:
      - name: networking
        components:
          - name: vpc
          - name: subnets
      - name: security
        components:
          - name: iam-roles
          - name: security-groups
```

### 2. Application Stack

```yaml
stacks:
  - name: stack-application
    layers:
      - name: database
        components:
          - name: rds-instance
          - name: redis-cluster
      - name: compute
        components:
          - name: eks-cluster
          - name: node-groups
```

## Deployment Strategies

### 1. Full Stack Deployment

```bash
infractl plan --stack=mystack --target-env=prod
infractl apply --stack=mystack --target-env=prod
```

### 2. Layer Deployment

```bash
infractl plan --stack=mystack --layer=networking --target-env=prod
infractl apply --stack=mystack --layer=networking --target-env=prod
```

### 3. Component Deployment

```bash
infractl plan --stack=mystack --layer=database --component=rds --target-env=prod
infractl apply --stack=mystack --layer=database --component=rds --target-env=prod
```

## Troubleshooting

### Common Issues

1. **Dependency Problems**

   - Check dependency paths
   - Verify output variables
   - Review dependency order

2. **Configuration Issues**

   - Validate YAML syntax
   - Check variable references
   - Verify provider configuration

3. **Deployment Failures**
   - Review error messages
   - Check state files
   - Verify permissions

## Next Steps

This concludes the documentation series. You now have a comprehensive understanding of the Terragrunt Reference Architecture v2. For additional support:

1. Review the example stacks in the repository
2. Check the InfraCTL CLI documentation
3. Consult the Terragrunt official documentation
