# ID Generator Terraform Module

## Overview

This Terraform module generates unique identifiers (UUIDs) with flexible configuration options for prefix, suffix, and formatting.

## Features

- Generate unique UUIDs
- Optional UUID formatting
- Configurable prefix and suffix
- Regeneration control via keepers

## Usage Example

```hcl
module "deployment_id" {
  source = "./id-generator"

  id_type        = "deployment"
  generation_key = timestamp()
  uuid_prefix    = "myapp"
  uuid_format    = "deploy-%s"
}

output "deployment_id" {
  value = module.deployment_id.full_uuid
}
```

## Variables

| Name             | Description                   | Type     | Default     | Validation             |
| ---------------- | ----------------------------- | -------- | ----------- | ---------------------- |
| `id_type`        | Type of identifier            | `string` | `"generic"` | -                      |
| `generation_key` | Trigger for UUID regeneration | `string` | `null`      | -                      |
| `uuid_format`    | Optional UUID format string   | `string` | `""`        | Valid format specifier |
| `uuid_prefix`    | Prefix for UUID               | `string` | `""`        | ≤ 32 characters        |
| `uuid_suffix`    | Suffix for UUID               | `string` | `""`        | ≤ 32 characters        |

## Outputs

- `uuid`: Raw generated UUID
- `formatted_uuid`: Formatted UUID
- `full_uuid`: UUID with prefix/suffix
- `id_type`: Identifier type

## Requirements

- Terraform >= 1.5.0
- Random Provider ~> 3.5.0
