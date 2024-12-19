# Terraform Random Resource Generators

## Overview

This directory contains a set of Terraform modules demonstrating the usage of HashiCorp's `random` provider for generating dynamic, configurable random resources.

## Modules

### 1. Name Generator

- **Resource**: `random_string`
- **Purpose**: Generate unique, configurable string prefixes for resource naming
- **Key Features**:
  - Customizable string length
  - Character composition control
  - Optional name templating

### 2. ID Generator

- **Resource**: `random_uuid`
- **Purpose**: Create unique identifiers with flexible formatting
- **Key Features**:
  - UUID generation with optional prefix/suffix
  - Formatting capabilities
  - Regeneration control

### 3. Quota Generator

- **Resource**: `random_integer`
- **Purpose**: Generate dynamic integer quotas for resource allocation
- **Key Features**:
  - Configurable quota range
  - Scaling and adjustment options
  - Percentage calculation

## Demo Scenario: Dynamic Resource Provisioning

The modules simulate a cloud resource management scenario where:

- Resource names are dynamically generated
- Unique tracking IDs are created
- Resource quotas are dynamically allocated

### Example Use Case

```hcl
module "web_service" {
  source = "./name-generator"
  prefix_length = 6
  name_template = "webservice-%s"
}

module "deployment_tracking" {
  source = "./id-generator"
  id_type = "deployment"
}

module "user_quota" {
  source = "./quota-generator"
  min_quota = 50
  max_quota = 200
  quota_type = "concurrent_users"
}

output "service_name" {
  value = module.web_service.full_resource_name
}
```

## Requirements

- Terraform >= 1.9.8
- Random Provider ~> 3.5.0
