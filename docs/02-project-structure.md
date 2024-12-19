# Project Structure

## Repository Organization

The repository follows a clear, hierarchical structure designed for scalability and maintainability:

```
.
├── infra/
│   ├── terraform/          # Terraform modules
│   │   └── module-name/    # Individual Terraform modules
│   └── terragrunt/         # Terragrunt configuration
│       ├── _ENVS/          # Environment configurations
│       ├── _shared/        # Shared components
│       ├── _templates/     # Templates for generated files
│       └── stack-*/        # Infrastructure stacks
└── tools/
    └── infractl/          # Custom CLI tool
```

## Key Directories

### `/infra/terraform/`

Contains reusable Terraform modules that define specific infrastructure components:

- Each module follows standard Terraform structure
- Includes `README.md`, `main.tf`, `variables.tf`, `outputs.tf`
- Modules are referenced by Terragrunt configurations

### `/infra/terragrunt/`

The heart of the infrastructure configuration:

#### `_ENVS/`

- Contains environment-specific YAML configurations
- `base.yaml`: Base configuration for all environments
- Environment-specific files (e.g., `local.yaml`, `production.yaml`)

#### `_shared/_components/`

- Reusable Terragrunt configurations
- Common patterns and configurations
- Referenced by specific components

#### `_templates/`

- Template files for generated configurations
- `providers.tf.tpl`: Provider configuration templates
- `versions.tf.tpl`: Version constraint templates

#### Stack Directories (`stack-*/`)

```
stack-example/
├── stack.hcl           # Stack-level configuration
└── layer-name/         # Logical infrastructure layer
    ├── layer.hcl       # Layer-level configuration
    └── component-name/ # Infrastructure component
        ├── component.hcl
        └── terragrunt.hcl
```

### `/tools/infractl/`

Custom CLI tool for managing infrastructure:

- Configuration validation
- Environment management
- Terragrunt command wrapper
- Built with Go for performance and reliability

## Configuration Hierarchy

1. **Stack Level** (`stack.hcl`)

   - Highest level of organization
   - Common configuration for all layers
   - Stack-wide tags and settings

2. **Layer Level** (`layer.hcl`)

   - Logical grouping of components
   - Layer-specific configuration
   - Common settings for components

3. **Component Level** (`component.hcl`, `terragrunt.hcl`)
   - Individual infrastructure resources
   - Component-specific configuration
   - References to Terraform modules

## File Naming Conventions

- Stack directories: `stack-<purpose>` (e.g., `stack-landing-zone`)
- Layer directories: Descriptive names (e.g., `networking`, `database`)
- Component files:
  - `component.hcl`: Component-specific configuration
  - `terragrunt.hcl`: Terragrunt execution configuration

## Best Practices

1. **Module Organization**

   - Keep modules focused and single-purpose
   - Include comprehensive documentation
   - Use consistent variable naming

2. **Stack Structure**

   - Group related components in layers
   - Use meaningful stack names
   - Maintain clear dependencies

3. **Configuration Management**
   - Keep environment configurations DRY
   - Use clear, descriptive names
   - Document all configuration options

## Next Steps

Continue to [Configuration System](03-configuration-system.md) to learn about the YAML-based configuration system.
