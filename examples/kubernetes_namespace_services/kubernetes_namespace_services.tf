resource "portainer_kubernetes_namespace_services" "test" {
  environment_id = var.environment_id
  namespace      = var.namespace
  name           = var.service_name
  type           = var.service_type

  selector = var.selector

  ports {
    port        = var.port
    target_port = var.target_port
  }
}
