package synovmm

import (
	"context"

	"github.com/m-bers/synovmm-client-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SYNOVMM_USERNAME", nil),
				Description: "The username for API operations.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SYNOVMM_PASSWORD", nil),
				Description: "The password for API operations.",
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SYNOVMM_HOST", nil),
				Description: "The SynoVMM host to connect to.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"synovmm_host": dataSourceHost(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	host := d.Get("host").(string)

	var diags diag.Diagnostics

	if (username != "") && (password != "") && (host != "") {
		c, err := synovmm.NewClient(&host, &username, &password)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}

	c, err := synovmm.NewClient(&host, &username, &password)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
