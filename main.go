package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"terraform-provider-portainer/internal"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: internal.Provider,
		ProviderAddr: "localdomain/local/portainer",
	})
}
