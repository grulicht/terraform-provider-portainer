# Portainer Provider

A [Terraform](https://www.terraform.io) provider to manage[Portainer](https://www.portainer.io/) resources via its REST API using Terraform.

It supports provisioning and configuration of Portainer users and will be extended to support other objects such as teams, compose/stacks, endpoints, and access control.

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
| `portainer_environment`    | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_tag`            | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_endpoint_group` | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_registry`       | ![Planned](https://img.shields.io/badge/status-planned-blue)          |
| `portainer_backup`         | ![Planned](https://img.shields.io/badge/status-planned-blue)          |
| `portainer_stack`          | ![Planned](https://img.shields.io/badge/status-planned-blue)          |
| `portainer_auth`           | ![Planned](https://img.shields.io/badge/status-planned-blue)          |
| `portainer_edge_group`     | ![Planned](https://img.shields.io/badge/status-planned-blue)          |
| `portainer_edge_stack`     | ![Planned](https://img.shields.io/badge/status-planned-blue)          |
| `portainer_edge_job`       | ![Planned](https://img.shields.io/badge/status-planned-blue)          |
| `portainer_custom_template`| ![Planned](https://img.shields.io/badge/status-planned-blue)          |
| `portainer_ldap_check`     | ![Planned](https://img.shields.io/badge/status-planned-blue)          |

---

### 💡 Missing a resource?

Is there a Portainer resource you'd like to see supported?

👉 [Open an issue](https://github.com/grulicht/terraform-provider-portainer/issues/new?template=feature_request.md) and we’ll consider it for implementation — or even better, submit a [Pull Request](https://github.com/grulicht/terraform-provider-portainer/pulls) to contribute directly!

📘 See [CONTRIBUTING.md](https://github.com/grulicht/terraform-provider-portainer/blob/main/.github/CONTRIBUTING.md) for guidelines.
