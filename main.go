package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform-providers/terraform-provider-reqres/reqres"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: reqres.Provider})
}
