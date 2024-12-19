# YAML Merger CLI Tool

## Overview

A Go-based CLI tool for merging YAML configurations with deep merging capabilities.

## Features

- Deep merging of nested configurations
- JSON-encoded input
- Optional output file writing
- Robust error handling and logging

## Usage

### Basic Parsing

```bash
./yaml-merger parse -input '{"base":{"key":"value"},"override":{"newkey":"newvalue"}}'
```

### Output to File

```bash
./yaml-merger parse -input '...' -output /path/to/merged/config.json
```

## Build Process

```bash
# Navigate to the script directory
cd infra/terragrunt/_scripts/yaml-merger

# Build the binary
./build.sh
```

## Integration with Terragrunt

The tool is automatically built and validated by `arch.hcl` before use.

## Error Handling

- Logs are printed to stdout
- Non-zero exit codes for failures
- Validates input and output configurations

## Dependencies

- Go 1.21+
- gopkg.in/yaml.v3 library
