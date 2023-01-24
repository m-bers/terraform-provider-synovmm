package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"terraform-provider-synovmm/synovmm"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return synovmm.Provider()
		},
	})
}

// package main

// import (
// 	"github.com/m-bers/synovmm-client-go"
// )

// func main() {
// 	// Create a new client
// 	host := "http://192.168.50.216:5000"
// 	user := "josh"
// 	pass := "Neverlos7"
// 	c, err := synovmm.NewClient(&host, &user, &pass)
// 	if err != nil {
// 		panic(err)
// 	}

// 	println(c.SID)

// 	// Get all Hosts
// 	hosts, err := c.GetHosts(&c.SID)
// 	if err != nil {
// 		panic(err)
// 	}

// 	println(hosts.Data.Hosts[0].HostID)
// 	println(hosts.Data.Hosts[0].HostName)
// 	println(hosts.Data.Hosts[0].Status)
// 	println(hosts.Data.Hosts[0].TotalCPUCore)
// 	println(hosts.Data.Hosts[0].TotalRAMSize)
// 	println(hosts.Data.Hosts[0].FreeCPUCore)
// 	println(hosts.Data.Hosts[0].FreeRAMSize)

// 	// Sign out of the API
// 	err = c.SignOut(&c.SID)
// 	if err != nil {
// 		panic(err)
// 	}

// }
