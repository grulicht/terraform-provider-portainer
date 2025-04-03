# 📘 **Resource Documentation: `portainer_kubernetes_namespace_services`**

# portainer_kubernetes_namespace_services
The `portainer_kubernetes_namespace_services` resource allows you to manage Kubernetes Services in a specific namespace within a Portainer-managed environment.

## Example Usage
```hcl
resource "portainer_kubernetes_namespace_services" "test" {
  environment_id = 4
  namespace      = "default"
  name           = "simple-service"
  type           = "ClusterIP"

  selector = {
    app = "nginx"
  }

  ports {
    port         = 80
    target_port  = 8080
  }
}
```

## Lifecycle & Behavior
- Terraform updates the Seriveces if any of the fields change (`annotations, labels, type, ports` etc.).
- The Seriveces is created or updated via the Portainer Kubernetes API using the appropriate HTTP method (`POST` or `PUT`).

### Arguments Reference
| Name                              | Type             | Required | Default      | Description                                                               |
|-----------------------------------|------------------|----------|--------------|---------------------------------------------------------------------------|
| `environment_id`                  | number           | ✅ Yes   | –            | ID of the Portainer environment (Kubernetes endpoint).                    |
| `namespace`                       | string           | ✅ Yes   | –            | Kubernetes namespace where the service will be created.                  |
| `name`                            | string           | ✅ Yes   | –            | Name of the Kubernetes Service.                                           |
| `type`                            | string           | 🚫 No    | `ClusterIP`  | Type of service (e.g., `ClusterIP`, `NodePort`, `LoadBalancer`).         |
| `allocate_load_balancer_node_ports` | bool           | 🚫 No    | `false`      | Only applicable when service type is `LoadBalancer`.                     |
| `selector`                        | map(string)      | ✅ Yes   | –            | Selector labels to match pods.                                            |
| `annotations`                     | map(string)      | 🚫 No    | `{}`         | Annotations for the service.                                              |
| `labels`                          | map(string)      | 🚫 No    | `{}`         | Labels for the service.                                                   |
| `external_ips`                    | list(string)     | 🚫 No    | `[]`         | External IPs for the service.                                             |
| `external_name`                   | string           | 🚫 No    | `""`         | External DNS name (used for `ExternalName` service type).                |
| `load_balancer_ip`                | string           | 🚫 No    | `""`         | Static IP to assign to `LoadBalancer`.                                   |
| `session_affinity`               | string           | 🚫 No    | `None`       | Session affinity settings (`None`, `ClientIP`).                           |
| `ports`                           | list(object)     | ✅ Yes   | –            | List of ports to expose. See [Nested `ports` Block](#️nested-ports-block) |

### Nested `ports` Block
| Name         | Type           | Required | Default | Description                                             |
|--------------|----------------|----------|---------|---------------------------------------------------------|
| `name`       | string         | 🚫 No    | –       | Optional name for the port.                             |
| `port`       | number         | ✅ Yes   | –       | Port to expose from the service.                        |
| `target_port`| string/number  | ✅ Yes   | –       | Port the traffic should be directed to.                 |

---

### Attributes Reference
| Name  | Description                                              |
|-------|----------------------------------------------------------|
| `id`  | Composite ID in format `environment_id:namespace:name`   |
