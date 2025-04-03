<!-- BEGIN_TF_DOCS -->


## Providers

| Name | Version |
|------|---------|
| <a name="provider_portainer"></a> [portainer](#provider\_portainer) | n/a |

## Resources

| Name | Type |
|------|------|
| [portainer_kubernetes_namespace_services.test](https://registry.terraform.io/providers/grulicht/portainer/latest/docs/resources/kubernetes_namespace_services) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_environment_id"></a> [environment\_id](#input\_environment\_id) | Portainer environment (endpoint) ID | `number` | `4` | no |
| <a name="input_namespace"></a> [namespace](#input\_namespace) | Kubernetes namespace for the service | `string` | `"default"` | no |
| <a name="input_port"></a> [port](#input\_port) | Service port exposed to the outside | `number` | `80` | no |
| <a name="input_portainer_api_key"></a> [portainer\_api\_key](#input\_portainer\_api\_key) | Default Portainer Admin API Key | `string` | n/a | yes |
| <a name="input_portainer_url"></a> [portainer\_url](#input\_portainer\_url) | Default Portainer URL | `string` | n/a | yes |
| <a name="input_selector"></a> [selector](#input\_selector) | Selector labels to match pods | `map(string)` | <pre>{<br/>  "app": "nginx"<br/>}</pre> | no |
| <a name="input_service_name"></a> [service\_name](#input\_service\_name) | Name of the Kubernetes service | `string` | `"simple-service"` | no |
| <a name="input_service_type"></a> [service\_type](#input\_service\_type) | Type of Kubernetes service (ClusterIP, NodePort, LoadBalancer) | `string` | `"ClusterIP"` | no |
| <a name="input_target_port"></a> [target\_port](#input\_target\_port) | Pod port to forward traffic to | `number` | `8080` | no |
<!-- END_TF_DOCS -->