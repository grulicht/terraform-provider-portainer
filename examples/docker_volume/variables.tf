variable "portainer_url" {
  description = "Default Portainer URL"
  type        = string
  # default     = "http://localhost:9000"
}

variable "portainer_api_key" {
  description = "Default Portainer Admin API Key"
  type        = string
  sensitive   = true
  # default     = "your-api-key-from-portainer"
}

variable "endpoint_id" {
  type        = number
  description = "ID of the Portainer endpoint/environment"
}

variable "volume_name" {
  type        = string
  description = "Name of the Docker volume"
}

variable "volume_driver" {
  type        = string
  default     = "local"
  description = "Docker volume driver to use"
}

variable "volume_driver_opts" {
  type = map(string)
  default = {
    device = "tmpfs"
    o      = "size=100m,uid=1000"
    type   = "tmpfs"
  }
  description = "Driver-specific options"
}

variable "volume_labels" {
  type = map(string)
  default = {
    env     = "test"
    managed = "terraform"
  }
  description = "Labels to apply to the volume"
}
