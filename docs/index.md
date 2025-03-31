# Portainer Provider

A [Terraform](https://www.terraform.io) provider to manage [Portainer](https://www.portainer.io/) resources via its REST API using Terraform.

It supports provisioning and configuration of Portainer users and will be extended to support other objects such as teams, compose/stacks, endpoints, and access control.

## Provider Support
| Provider       | Provider Support Status              |
|----------------|--------------------------------------|
| [Terraform](https://registry.terraform.io/providers/grulicht/portainer/latest)      | ![Done](https://img.shields.io/badge/status-done-brightgreen)           |
| [OpenTofu](https://search.opentofu.org/provider/grulicht/portainer/latest)       | ![Done](https://img.shields.io/badge/status-done-brightgreen) |

## Example Provider Configuration
```hcl
provider "portainer" {
  endpoint = "..."
  api_key  = "..."
}
```

## Authentication
- Static API key

Static credentials can be provided by adding the `api_key` variables in-line in the Portainer provider block:

> 🔐 **Authentication:** This provider supports only **API keys** via the `X-API-Key` header. JWT tokens curentlly are not supported in this provider.

Usage:

```hcl
provider "portainer" {
  api_key  = "..."
}
```
### Environment variables
You can provide your configuration via the environment variables representing your minio credentials:

```hcl
$ export PORTAINER_ENDPOINT="http://portainer.example.com"
$ export PORTAINER_API_KEY="your-api-key"
```

## Arguments Reference
| Name       | Type   | Required | Description                                                                 |
|------------|--------|----------|-----------------------------------------------------------------------------|
| `endpoint` | string | ✅ yes   | The URL of the Portainer instance. `/api` will be appended automatically if missing. |
| `api_key`  | string | ✅ yes   | API key used to authenticate requests.                                      |


## 🧩 Supported Resources
| Resource                   | Status                                                                 |
|----------------------------|------------------------------------------------------------------------|
| `portainer_user`           | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_team`           | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_team_membrship` | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_environment`    | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_tag`            | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_endpoint_group` | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_registry`       | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_backup`         | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_backup_s3`      | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_auth`           | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_edge_group`     | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_edge_stack`     | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_edge_job`       | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_stack`          | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_custom_template`| ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_container_exec` | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_docker_network` | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_docker_image`   | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_docker_volume`  | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_open_amt`       | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_settings`       | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_webhook`        | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_webhook_execute`| ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_licenses`       | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_cloud_credentials`| ![Done](https://img.shields.io/badge/status-done-brightgreen)       |

---

### 💡 Missing a resource?
Is there a Portainer resource you'd like to see supported?

👉 [Open an issue](https://github.com/grulicht/terraform-provider-portainer/issues/new?template=feature_request.md) and we’ll consider it for implementation — or even better, submit a [Pull Request](https://github.com/grulicht/terraform-provider-portainer/pulls) to contribute directly!

📘 See [CONTRIBUTING.md](https://github.com/grulicht/terraform-provider-portainer/blob/main/.github/CONTRIBUTING.md) for guidelines.

## ✅ Daily End-to-End Testing
To ensure maximum reliability and functionality of this provider, **automated end-to-end tests are executed every day** via GitHub Actions.

These tests run against a real Portainer instance (started using docker compose) and validate the majority of supported resources using real Terraform plans and applies.

> 💡 This helps catch regressions early and ensures the provider remains fully operational and compatible with the Portainer API.

## License
This module is 100% Open Source and all versions of this provider starting from v2.0.0 are distributed under the AGPL-3.0 License. See [LICENSE](https://github.com/grulicht/terraform-provider-portainer/blob/main/LICENSE) for more information.

## Authors
Created by [Tomáš Grulich](https://github.com/grulicht) - to.grulich@gmail.com

## Acknowledgements
- [HashiCorp Terraform](https://www.hashicorp.com/products/terraform)
- [Portainer](https://portainer.io)
- [OpenTofu](https://opentofu.org/)
- [Docker](https://www.docker.com/)
