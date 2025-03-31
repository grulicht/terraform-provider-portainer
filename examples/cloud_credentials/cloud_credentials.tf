resource "portainer_cloud_credentials" "example" {
  name        = var.cloud_credentials_name
  provider    = var.cloud_credentials_provider
  credentials = var.cloud_credentials_data
}
