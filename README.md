
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

# Portainer Terraform Provider
A [Terraform](https://www.terraform.io) provider to manage [Portainer](https://www.portainer.io/) resources via its REST API using Terraform.

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
| Resource                   | Documentation                                               | Example                                 | Status                                                                 | E2E Tests                                                            |
|----------------------------|-------------------------------------------------------------|-----------------------------------------|------------------------------------------------------------------------|----------------------------------------------------------------------|
| `portainer_user`           | [📘 user.md](docs/resources/user.md)                       | [📂 example](examples/user/)             | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_team`           | [📘 team.md](docs/resources/team.md)                       | [📂 example](examples/team/)             | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_team_membership`| [📘 team_membership.md](docs/resources/team_membership.md) | [📂 example](examples/team_membership/)  | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_environment`    | [📘 environment.md](docs/resources/environment.md)         | [📂 example](examples/environment/)      | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-false-grey)              |
| `portainer_tag`            | [📘 tag.md](docs/resources/tag.md)                         | [📂 example](examples/tag/)              | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_endpoint_group` | [📘 endpoint_group.md](docs/resources/endpoint_group.md)   | [📂 example](examples/endpoint_group/)   | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_registry`       | [📘 registry.md](docs/resources/registry.md)               | [📂 example](examples/registry/)         | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_backup`         | [📘 backup.md](docs/resources/backup.md)                   | [📂 example](examples/backup/)           | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_backup_s3`      | [📘 backup.md](docs/resources/backup_s3.md)                | [📂 example](examples/backup_s3/)        | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-false-grey)              |
| `portainer_auth`           | [📘 auth.md](docs/resources/auth.md)                       | [📂 example](examples/auth/)             | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_edge_group`     | [📘 edge_group.md](docs/resources/edge_group.md)           | [📂 example](examples/edge_group/)       | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-false-grey)              |
| `portainer_edge_stack`     | [📘 edge_stack.md](docs/resources/edge_stack.md)           | [📂 example](examples/edge_stack/)       | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-false-grey)              |
| `portainer_edge_job`       | [📘 edge_job.md](docs/resources/edge_job.md)               | [📂 example](examples/edge_job/)         | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-false-grey)              |
| `portainer_stack`          | [📘 stack.md](docs/resources/stack.md)                     | [📂 example](examples/stack/)            | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_custom_template`| [📘 custom_template.md](docs/resources/custom_template.md) | [📂 example](examples/custom_template/)  | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_container_exec` | [📘 container_exec.md](docs/resources/container_exec.md)   | [📂 example](examples/container_exec/)   | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_docker_network` | [📘 docker_network.md](docs/resources/docker_network.md)   | [📂 example](examples/docker_network/)   | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_docker_image`   | [📘 docker_image.md](docs/resources/docker_image.md)       | [📂 example](examples/docker_image/)     | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_docker_volume`  | [📘 docker_volume.md](docs/resources/docker_volume.md)     | [📂 example](examples/docker_volume/)    | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_open_amt`       | [📘 open_amt.md](docs/resources/open_amt.md)               | [📂 example](examples/open_amt/)         | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-false-grey)              |
| `portainer_settings`       | [📘 settings.md](docs/resources/settings.md)               | [📂 example](examples/settings/)         | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_ssl`            | [📘 ssl.md](docs/resources/ssl.md)                         | [📂 example](examples/ssl/)              | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_webhook`           | [📘 webhook.md](docs/resources/webhook.md)                         | [📂 example](examples/webhook/)                        | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-daily-blue)              |
| `portainer_webhook_execute`   | [📘 webhook_execute.md](docs/resources/webhook_execute.md)         | [📂 example](examples/webhook_execute/)                | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-false-grey)              |
| `portainer_licenses`          | [📘 licenses.md](docs/resources/licenses.md)                       | [📂 example](examples/licenses/)                       | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-false-grey)              |
| `portainer_cloud_credentials` | [📘 cloud_credentials.md](docs/resources/cloud_credentials.md)     | [📂 example](examples/cloud_credentials/)              | ![Done](https://img.shields.io/badge/status-done-brightgreen)         | ![Tested](https://img.shields.io/badge/running-false-grey)              |

---

### 💡 Missing a resource?
Is there a Portainer resource you'd like to see supported?

👉 [Open an issue](https://github.com/grulicht/terraform-provider-portainer/issues/new?template=feature_request.md) and we’ll consider it for implementation — or even better, submit a [Pull Request](https://github.com/grulicht/terraform-provider-portainer/pulls) to contribute directly!

📘 See [CONTRIBUTING.md](./.github/CONTRIBUTING.md) for guidelines.

## ✅ Daily End-to-End Testing
To ensure maximum reliability and functionality of this provider, **automated end-to-end tests are executed every day** via GitHub Actions.

These tests run against a real Portainer instance (started using docker compose) and validate the majority of supported resources using real Terraform plans and applies.

> 💡 This helps catch regressions early and ensures the provider remains fully operational and compatible with the Portainer API.

### 🔄 Workflows
The project uses GitHub Actions to automate validation and testing of the provider.

- Validate and lint documentation files (`README.md` and `docs/`)
- Initialize, test and check the Portainer provider with **Terraform** and **OpenTofu**
- Publish the new version of the Portainer Terraform provider to Terraform Registry
- Run daily **E2E Terraform tests** against a live Portainer instance spun up via Docker Compose (`make up`) at **07:00 UTC**

### 🧪 Localy Testing
To test the provider locally, start the Portainer Web UI using Docker Compose:
```sh
make up
```
Then open `http://localhost:9000` in your browser.

### 🔐 Predefined Test Credentials for Login (use also E2E tests)
Thanks to the `portainer_data` directory included in this repository, a test user and token are preloaded when you launch the local Portainer instance:

| **Field**    | **Value**                                                                  |
|--------------|----------------------------------------------------------------------------|
| Username     | `admin`                                                                    |
| Password     | `password123456789`                                                        |
| API Token    | `ptr_xrP7XWqfZEOoaCJRu5c8qKaWuDtVc2Zb07Q5g22YpS8=`                          |

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
- [Docker](https://www.docker.com/)
