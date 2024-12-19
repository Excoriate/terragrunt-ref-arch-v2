# AWS DynamoDB Generator

## Overview

A Terraform module that generates an AWS DynamoDB table with dynamically provided:

- Name
- Unique Identifier
- Quota

## Features

- Flexible input for table naming and tagging
- Supports dynamically generated or explicitly provided values
- Cost-effective configuration options

## Usage Example

```hcl
module "test_table" {
  source = "./aws-dynamodb-generator"

  generated_name = "my-dynamic-table"
  generated_id   = "unique-tracking-id"
  generated_quota = 10
  quota_type     = "read_capacity"

  # Optional customizations
  hash_key_name = "custom_id"
  additional_tags = {
    Environment = "test"
  }
}

output "table_name" {
  value = module.test_table.table_name
}
```

## Inputs for Dynamic Generation

| Name              | Description                       | Type          | Default         |
| ----------------- | --------------------------------- | ------------- | --------------- |
| `table_name`      | Explicit table name               | `string`      | `""`            |
| `generated_name`  | Generated name for naming/tagging | `string`      | `""`            |
| `generated_id`    | Generated unique identifier       | `string`      | `""`            |
| `generated_quota` | Generated quota value             | `number`      | `null`          |
| `quota_type`      | Type of quota                     | `string`      | `"unspecified"` |
| `additional_tags` | Extra tags                        | `map(string)` | `{}`            |

## Terragrunt Integration

This module is designed to receive dynamically generated values from Terragrunt:

- Name from Name Generator
- ID from ID Generator
- Quota from Quota Generator

## Requirements

- Terraform >= 1.9.8
- AWS Provider 5.81.0

## Cost Considerations

- Uses PAY_PER_REQUEST billing mode
- Configurable point-in-time recovery
- Optional Time-To-Live
- Minimal default configuration
