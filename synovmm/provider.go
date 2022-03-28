package provider

import (
	"context"
	"fmt"

	"github.com/dghubble/sling"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	baseURLName = "baseURL"
	accountName = "account"
	passwdName  = "passwd"
)

const (
	MajorVersion    = 0
	MinorVersion    = 1
	PatchVersion    = 0
	UserAgentPrefix = "synovmm-terraform-provider"
)

var Version = fmt.Sprintf("%d.$d.$d", MajorVersion, MinorVersion, PatchVersion)

// Provider returns a terraform provider for the Synology Virtual Machine Manager API
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"baseURLName": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SYNOLOGY_URL", "http://ds918:5000/webapi/entry.cgi"),
			},
			"accountName": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("SYNOLOGY_USERNAME", nil),
			},
			"passwdName": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("SYNOLOGY_PASSWORD", nil),
			},
		},
		ConfigureContextFunc: func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {

		},
	}
}

func setupSynoVMMContext(ctx context.Context, returnsd *schema.ResourceData) (interface{}, diag.Diagnostics) {
	account := rd.Get(accountName).(string)
	baseURL := rd.Get(baseURLName).(string)

	type IssueParams struct {
		Account  string `url:"account,omitempty"`
		Password string `url:"password,omitempty"`
	}

	var d diag.Diagnostics

	client := sling.New().Base(baseURL).
		Set("User-Agent", fmt.Sprintf("%s (%s)", UserAgentPrefix, Version))

	resp, err := client.Get("ping").Receive(nil, nil)
	if err != nil {

	}

	if resp.StatusCode != 200 {
		return nil, diag.FromErr(fmt.Errorf("Invalid response code from Synology VMM API: %d", resp.StatusCode))
	}

	return client, nil
}
