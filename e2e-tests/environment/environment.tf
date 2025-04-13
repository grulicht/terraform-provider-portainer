resource "portainer_environment" "your-host" {
  name                = var.portainer_environment_name
  environment_address = var.portainer_environment_address
  type                = var.portainer_environment_type

  tls_enabled            = var.tls_enabled
  tls_skip_verify        = var.tls_skip_verify
  tls_skip_client_verify = var.tls_skip_client_verify
}
