locals {
  # ---------------------------------------------------------------------------------------------------------------------
  # COMPONENT CONFIGURATION
  # ---------------------------------------------------------------------------------------------------------------------
  is_enabled = true
  name       = "namecheap-ns-updater"

  tags = {
    Name = "namecheap-nameservers-updater"
    ArchitectureType = "component"
  }
}