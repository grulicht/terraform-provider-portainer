
<p align="center">
  <a href="https://registry.terraform.io/providers/grulicht/portainer/latest/docs">
    <img src="https://www.terraform.io/_next/static/media/terraform-community_on-light.cda79e7c.svg" alt="Terraform Logo" width="200">
  </a>
  &nbsp;&nbsp;&nbsp;
  <a href="https://github.com/grulicht/terraform-provider-portainer">
    <img src="https://www.portainer.io/hubfs/portainer-logo-black.svg" alt="portainer-provider-terraform" width="200">
  </a>
  &nbsp;&nbsp;&nbsp;
  <a href="https://search.opentofu.org/provider/grulicht/portainer/latest">
    <img src="https://raw.githubusercontent.com/opentofu/brand-artifacts/main/full/transparent/SVG/on-dark.svg#gh-dark-mode-only" alt="portainer-provider-opentofu" width="200">
  </a>
  <h3 align="center" style="font-weight: bold">Terraform Provider for Portainer</h3>
  <p align="center">
    <a href="https://github.com/grulicht/terraform-provider-portainer/graphs/contributors">
      <img alt="Contributors" src="https://img.shields.io/github/contributors/grulicht/terraform-provider-portainer">
    </a>
    <a href="https://golang.org/doc/devel/release.html">
      <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/grulicht/terraform-provider-portainer">
    </a>
    <a href="https://github.com/grulicht/terraform-provider-portainer/actions?query=workflow%3Arelease">
      <img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/grulicht/terraform-provider-portainer/release.yml?tag=latest&label=release">
    </a>
    <a href="https://github.com/grulicht/terraform-provider-portainer/releases">
      <img alt="GitHub release (latest by date including pre-releases)" src="https://img.shields.io/github/v/release/grulicht/terraform-provider-portainer?include_prereleases">
    </a>
  </p>
  <p align="center">
    <a href="https://github.com/grulicht/terraform-provider-portainer/tree/main/docs"><strong>Explore the docs »</strong></a>
  </p>
</p>

# Portainer CE Terraform Provider

A [Terraform](https://www.terraform.io) provider to manage[Portainer](https://www.portainer.io/) resources via its REST API using Terraform.

It supports provisioning and configuration of Portainer users and will be extended to support other objects such as teams, stacks, endpoints, and access control.

## Requirements

- Terraform v0.13+
- Portainer 2.x with admin API key support enabled
- Go 1.21+ (if building from source)

## Building and Installing
```hcl
make build
```

## Provider Support
| Provider       | Provider Support Status              |
|----------------|--------------------------------------|
| [Terraform](https://registry.terraform.io/providers/grulicht/portainer/latest)      | ![Done](https://img.shields.io/badge/status-done-brightgreen)           |
| [OpenTofu](https://search.opentofu.org/provider/grulicht/portainer/latest)       | ![Done](https://img.shields.io/badge/status-done-brightgreen) |


## Example Provider Configuration

```hcl
provider "portainer" {
  endpoint = "https://portainer.example.com"
  api_key  = "your-api-key"
}
```

## Authentication
- Static API key

Static credentials can be provided by adding the `api_key` variables in-line in the Portainer provider block:
> 🔐 **Authentication:** This provider supports only **API keys** via the `X-API-Key` header. JWT tokens curentlly are not supported in this provider.

Usage:

```hcl
provider "portainer" {
  api_key  = "your-api-key"
}
```
### Environment variables

You can provide your configuration via the environment variables representing your minio credentials:

```hcl
$ export PORTAINER_ENDPOINT="https://portainer.example.com"
$ export PORTAINER_API_KEY="your-api-key"
```

## Arguments Reference

| Name       | Type   | Required | Description                                                                 |
|------------|--------|----------|-----------------------------------------------------------------------------|
| `endpoint` | string | ✅ yes   | The URL of the Portainer instance. `/api` will be appended automatically if missing. |
| `api_key`  | string | ✅ yes   | API key used to authenticate requests.                                      |

## Usage

See our [examples](./docs/resources/) per resources in docs.

## 🧩 Supported Resources

| Resource                   | Documentation                                               | Example                                 | Status                                                                 |
|----------------------------|-------------------------------------------------------------|-----------------------------------------|------------------------------------------------------------------------|
| `portainer_user`           | [📘 user.md](docs/resources/user.md)                       | [📂 example](examples/user/)             | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_team`           | [📘 team.md](docs/resources/team.md)                       | [📂 example](examples/team/)             | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_environment`    | [📘 environment.md](docs/resources/environment.md)         | [📂 example](examples/environment/)      | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_tag`            | [📘 tag.md](docs/resources/tag.md)                         | [📂 example](examples/tag/)              | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_endpoint_group` | [📘 endpoint_group.md](docs/resources/endpoint_group.md)   | [📂 example](examples/endpoint_group/)   | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_registry`       | [📘 registry.md](docs/resources/registry.md)               | [📂 example](examples/registry/)         | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_backup`         | [📘 backup.md](docs/resources/backup.md)                   | [📂 example](examples/backup/)           | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_backup_s3`      | [📘 backup.md](docs/resources/backup_s3.md)                | [📂 example](examples/backup_s3/)        | ![Done](https://img.shields.io/badge/status-done-brightgreen)         |
| `portainer_stack`          | [📘 stack.md](docs/resources/stack.md)                     | [📂 example](examples/stack/)            | ![Planned](https://img.shields.io/badge/status-planned-blue)          |
| `portainer_auth`           | [📘 auth.md](docs/resources/auth.md)                       | [📂 example](examples/auth/)             | ![Planned](https://img.shields.io/badge/status-planned-blue)          |
| `portainer_edge_group`     | [📘 edge_group.md](docs/resources/edge_group.md)           | [📂 example](examples/edge_group/)       | ![Planned](https://img.shields.io/badge/status-planned-blue)          |
| `portainer_edge_stack`     | [📘 edge_stack.md](docs/resources/edge_stack.md)           | [📂 example](examples/edge_stack/)       | ![Planned](https://img.shields.io/badge/status-planned-blue)          |
| `portainer_edge_job`       | [📘 edge_job.md](docs/resources/edge_job.md)               | [📂 example](examples/edge_job/)         | ![Planned](https://img.shields.io/badge/status-planned-blue)          |
| `portainer_custom_template`| [📘 custom_template.md](docs/resources/custom_template.md) | [📂 example](examples/custom_template/)  | ![Planned](https://img.shields.io/badge/status-planned-blue)          |
| `portainer_ldap_check`     | [📘 ldap_check.md](docs/resources/ldap_check.md)           | [📂 example](examples/ldap_check/)       | ![Planned](https://img.shields.io/badge/status-planned-blue)          |

---

### 💡 Missing a resource?

Is there a Portainer resource you'd like to see supported?

👉 [Open an issue](https://github.com/grulicht/terraform-provider-portainer/issues/new?template=feature_request.md) and we’ll consider it for implementation — or even better, submit a [Pull Request](https://github.com/grulicht/terraform-provider-portainer/pulls) to contribute directly!

📘 See [CONTRIBUTING.md](./.github/CONTRIBUTING.md) for guidelines.

## Testing
To test the provider locally, start the Portainer Web UI using Docker Compose:
```sh
make up
```
Then open `http://localhost:9000` in your browser.
You can now apply your Terraform templates and observe changes live in the UI.

### Testing a new version of the Portainer provider
After making changes to the provider source code, follow these steps:
Build the provider binary:
```sh
make build
```
Install the binary into the local Terraform plugin directory:
```sh
make install-plugin
```
Update your main.tf to use the local provider source
Add the following to your Terraform configuration:
```sh
terraform {
  required_providers {
    portainer = {
      source  = "localdomain/local/portainer"
    }
  }
}
```
Now you're ready to test your provider against the local Portainer instance.

## Roadmap

See the [open issues](https://github.com/grulicht/terraform-provider-portainer/issues) for a list of proposed features (and known issues). See [CONTRIBUTING](./.github/CONTRIBUTING.md) for more information.

## License

This module is 100% Open Source and all versions of this provider starting from v2.0.0 are distributed under the AGPL-3.0 License. See [LICENSE](https://github.com/grulicht/terraform-provider-portainer/blob/main/LICENSE) for more information.

## Authors
Created by [Tomáš Grulich](https://github.com/grulicht) - to.grulich@gmail.com

## Acknowledgements
- [HashiCorp Terraform](https://www.hashicorp.com/products/terraform)
- [Portainer](https://portainer.io)
- [OpenTofu](https://opentofu.org/)
