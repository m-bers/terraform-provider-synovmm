package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// Provider returns a terraform provider for the Synology Virtual Machine Manager API
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema {
			"username": &schema.Schema {
				Type: schema.TypeString,
				Optional: false,
				DefaultFunc: schema.EnvDefaultFunc("SYNOLOGY_USERNAME", nil),
			},
			"base_url": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				DefaultFunc: schema.EnvDefaultFunc("SYNOLOGY_URL", "http://ds918:5000/webapi/entry.cgi")
			}
		},
	}
}