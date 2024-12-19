locals {
  # Additional processing or formatting of the UUID
  formatted_uuid = var.uuid_format != "" ? format(var.uuid_format, random_uuid.resource_id.result) : random_uuid.resource_id.result

  # Optional prefix or suffix
  prefixed_uuid = var.uuid_prefix != "" ? "${var.uuid_prefix}-${local.formatted_uuid}" : local.formatted_uuid
  suffixed_uuid = var.uuid_suffix != "" ? "${local.prefixed_uuid}-${var.uuid_suffix}" : local.prefixed_uuid
}
