terraform {
  required_providers {
    portainer = {
      source = "grulicht/portainer"
    }
  }
}

provider "portainer" {
  endpoint = var.portainer_url
  api_key  = var.portainer_api_key
}
