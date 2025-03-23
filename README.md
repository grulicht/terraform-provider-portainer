
<p align="center">
  <a href="https://github.com/grulicht/terraform-provider-portainer">
    <img src="https://www.portainer.io/hubfs/portainer-logo-black.svg" alt="portainer-provider-terraform" width="200">
  </a>
  <h3 align="center" style="font-weight: bold">Terraform Provider for Portainer</h3>
  <p align="center">
    <a href="https://github.com/grulicht/terraform-provider-portainer/graphs/contributors">
      <img alt="Contributors" src="https://img.shields.io/github/contributors/grulicht/terraform-provider-portainer">
    </a>
    <a href="https://golang.org/doc/devel/release.html">
      <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/grulicht/terraform-provider-portainer">
    </a>
    <a href="https://github.com/grulicht/terraform-provider-portainer/actions?query=workflow%3A%22Terraform+Provider+CI%22">
      <img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/grulicht/terraform-provider-portainer/go.yml?branch=main">
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

A [Terraform](https://www.terraform.io) provider to manage[Portainer](https://www.portainer.io/) resources via its REST API using Terraform. It supports provisioning and configuration of Portainer users and will be extended to support other objects such as teams, stacks, endpoints, and access control.

## Requirements

- Terraform v0.13+
- Portainer 2.x with admin API key support enabled
- Go 1.21+ (if building from source)

## Building and Installing


## Example Provider Configuration

```hcl
provider "portainer" {
  endpoint = "https://portainer.example.com"
  api_key  = "your-api-key"
}
```
> 🔐 **Authentication:** This provider supports only **API keys** via the `X-API-Key` header. JWT tokens are not supported.

## Authentication
- Static API key

Static credentials can be provided by adding the `api_key` variables in-line in the Portainer provider block:

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
| `api_key`  | string | ✅ yes   | API key used to authenticate requests. Only `X-API-Key` is supported.       |


## Resources

| Resource                             | Description                      |
|--------------------------------------|----------------------------------|
| [`portainer_user`](docs/resources/README-user.md)                | Manages Portainer users         |
| [`portainer_team`](docs/resources/README-team.md)                | Manages Portainer teams         |
| [`portainer_environment`](docs/resources/README-environment.md)  | Manages Portainer environments  |
| [`portainer_tag`](docs/resources/README-tag.md)                  | Manages Portainer tags          |
| [`portainer_endpoint_group`](docs/resources/README-endpoint-group.md) | Manages Portainer endpoint groups |

## Testing

For testing locally, run the docker compose to spin up a portainer web ui:

```sh
docker compose up
```

Access `http://localhost:9000` on your browser, apply your terraform templates and watch them going live.

## License

All versions of this provider starting from v2.0.0 are distributed under the AGPL-3.0 License. See [LICENSE](./LICENSE) for more information.
