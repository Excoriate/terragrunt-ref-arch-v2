# Name Generator Terraform Module

## Overview

This Terraform module generates random strings for resource naming, providing flexible configuration for prefix generation.

## Features

- Generate random string prefixes
- Configurable length and character composition
- Optional name templating
- Predictable regeneration based on specific attributes

## Usage Example

```hcl
module "resource_name" {
  source = "./name-generator"

  prefix_length         = 10
  include_special_chars = false
  prefix_type           = "dev"
  name_template         = "myapp-%s-instance"
}

output "resource_name" {
  value = module.resource_name.full_resource_name
}
```

## Variables

| Name                    | Description                | Type     | Default         | Validation      |
| ----------------------- | -------------------------- | -------- | --------------- | --------------- |
| `prefix_length`         | Length of random prefix    | `number` | `8`             | 4-32 characters |
| `include_special_chars` | Include special characters | `bool`   | `false`         | -               |
| `include_uppercase`     | Include uppercase letters  | `bool`   | `true`          | -               |
| `include_numeric`       | Include numeric characters | `bool`   | `true`          | -               |
| `prefix_type`           | Prefix generation type     | `string` | `"generic"`     | -               |
| `name_template`         | Resource name template     | `string` | `"%s-resource"` | -               |

## Outputs

- `random_prefix`: Generated random string
- `full_resource_name`: Complete resource name
- `prefix_length`: Length of generated prefix

## Requirements

- Terraform >= 1.5.0
- Random Provider ~> 3.5.0
