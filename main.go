package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/reqres"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: reqres.Provider})
}
