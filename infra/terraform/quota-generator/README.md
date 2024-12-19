# Quota Generator Terraform Module

## Overview

This Terraform module generates random integer quotas with advanced configuration options for scaling, adjustment, and type specification.

## Features

- Generate random integer quotas within a specified range
- Optional quota scaling
- Numerical adjustment (floor/ceil)
- Quota type specification
- Percentage calculation within quota range

## Usage Example

```hcl
module "user_quota" {
  source = "./quota-generator"

  min_quota        = 50
  max_quota        = 200
  quota_type       = "concurrent_users"
  scale_factor     = 1.5
  adjustment_type  = "floor"
}

output "max_concurrent_users" {
  value = module.user_quota.adjusted_quota
}
```

## Variables

| Name              | Description                    | Type     | Default     | Validation              |
| ----------------- | ------------------------------ | -------- | ----------- | ----------------------- |
| `min_quota`       | Minimum quota value            | `number` | `10`        | ≥ 0                     |
| `max_quota`       | Maximum quota value            | `number` | `100`       | > min_quota             |
| `quota_type`      | Type of quota                  | `string` | `"generic"` | -                       |
| `generation_seed` | Seed for consistent generation | `string` | `null`      | -                       |
| `scale_factor`    | Quota scaling factor           | `number` | `1.0`       | > 0                     |
| `adjustment_type` | Numerical adjustment type      | `string` | `"none"`    | 'none', 'floor', 'ceil' |

## Outputs

- `raw_quota`: Original generated quota
- `scaled_quota`: Scaled quota value
- `adjusted_quota`: Final adjusted quota
- `quota_percentage`: Quota percentage within range
- `quota_type`: Specified quota type

## Requirements

- Terraform >= 1.5.0
- Random Provider ~> 3.5.0
