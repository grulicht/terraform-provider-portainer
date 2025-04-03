variable "portainer_url" {
  description = "Default Portainer URL"
  type        = string
  # default     = "http://localhost:9000"
}

variable "portainer_api_key" {
  description = "Default Portainer Admin API Key"
  type        = string
  sensitive   = true
  # default     = "your-portainer-api-key"
}

variable "environment_id" {
  type        = number
  description = "Portainer environment (endpoint) ID"
  default     = 4
}

variable "namespace" {
  type        = string
  description = "Kubernetes namespace for the service"
  default     = "default"
}

variable "service_name" {
  type        = string
  description = "Name of the Kubernetes service"
  default     = "simple-service"
}

variable "service_type" {
  type        = string
  description = "Type of Kubernetes service (ClusterIP, NodePort, LoadBalancer)"
  default     = "ClusterIP"
}

variable "selector" {
  type        = map(string)
  description = "Selector labels to match pods"
  default = {
    app = "nginx"
  }
}

variable "port" {
  type        = number
  description = "Service port exposed to the outside"
  default     = 80
}

variable "target_port" {
  type        = number
  description = "Pod port to forward traffic to"
  default     = 8080
}
