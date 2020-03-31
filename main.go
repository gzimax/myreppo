package main

import (
	"github.com/gzimax/terraform-provider-pingaccess/pingaccess"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pingaccess.Provider})
}
