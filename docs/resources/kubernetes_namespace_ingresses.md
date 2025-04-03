# 🌐 **Resource Documentation: `portainer_kubernetes_namespace_ingresses`**

# portainer_kubernetes_namespace_ingresses
The `portainer_kubernetes_namespace_ingresses` resource allows you to manage Kubernetes Ingress resources inside a specific namespace in a Portainer-managed Kubernetes environment.

## Example Usage
```hcl
resource "portainer_kubernetes_namespace_ingresses" "example" {
  environment_id = 4
  namespace      = "default"
  name           = "testest1-ingress-1"
  class_name     = "nginx"

  annotations = {
    "kubernetes.io/ingress.class" = "nginx"
  }

  labels = {
    "app" = "nginx"
  }

  hosts = ["example.com"]

  tls {
    hosts       = ["example.com"]
    secret_name = "example-tls"
  }

  paths {
    host         = "example.com"
    path         = "/"
    path_type    = "Prefix"
    port         = 80
    service_name = "nginx-service"
  }
}
```

## Lifecycle & Behavior
- Terraform updates the Ingress if any of the fields change (`annotations, labels, tls, paths` etc.).
- The Ingress is created or updated via the Portainer Kubernetes API using the appropriate HTTP method (`POST` or `PUT`).
- 🚫 Ingress deletion is not currently supported, as the Portainer API does not provide a DELETE endpoint for Ingresses.

### Arguments Reference
| Name           | Type           | Required | Description                                                                 |
|----------------|----------------|----------|-----------------------------------------------------------------------------|
| `environment_id` | number       | ✅ yes   | ID of the Portainer environment (Kubernetes endpoint).                     |
| `namespace`    | string         | ✅ yes   | Namespace in which to create the Ingress.                                  |
| `name`         | string         | ✅ yes   | Name of the Ingress. Must be unique within the namespace.                  |
| `class_name`   | string         | ✅ yes   | Ingress controller class name (e.g. `nginx`).                               |
| `annotations`  | map(string)    | 🚫 optional | Key-value pairs to annotate the Ingress object.                         |
| `labels`       | map(string)    | 🚫 optional | Labels for organizing and selecting the Ingress.                        |
| `hosts`        | list(string)   | 🚫 optional | List of hostnames that the Ingress will match.                           |
| `tls`          | list(object)   | 🚫 optional | TLS block for securing connections. Requires `hosts` and `secret_name`.  |
| `paths`        | list(object)   | 🚫 optional | Routing rules for the Ingress. See below for structure.                   |

#### `tls` Block
| Name          | Type         | Required | Description                             |
|---------------|--------------|----------|-----------------------------------------|
| `hosts`       | list(string) | ✅ yes   | List of TLS hostnames.                  |
| `secret_name` | string       | ✅ yes   | Name of the TLS secret to use.          |

#### `paths` Block
| Name          | Type         | Required | Description                                                                  |
|---------------|--------------|----------|------------------------------------------------------------------------------|
| `host`        | string       | ✅ yes   | Host for the rule (e.g. `example.com`).                                     |
| `path`        | string       | ✅ yes   | Path to match for the incoming request.                                     |
| `path_type`   | string       | ✅ yes   | Path matching strategy (`Prefix`, `Exact`, etc.).                           |
| `port`        | number       | ✅ yes   | Target port of the backend service.                                         |
| `service_name`| string       | ✅ yes   | Name of the backend Kubernetes service handling the matched path.          |

---

### Attributes Reference
| Name  | Description                                               |
|-------|-----------------------------------------------------------|
| `id`  | Composite ID in the format `environmentID:namespace:name` |
