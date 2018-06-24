package main

import (
	"github.com/hashicorp/terraform/mockable"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mockable.Provider})
}
