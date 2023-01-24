package synovmm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sv "github.com/m-bers/synovmm-client-go"
)

func dataSourceHost() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHostRead,
		Schema: map[string]*schema.Schema{
			"hosts": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"free_cpu_core": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"free_ram_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"host_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_cpu_core": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"total_ram_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func dataSourceHostRead(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {

	c := m.(*sv.Client)
	var diags diag.Diagnostics
	hosts, err := c.GetHosts(&c.SID)

	if err != nil {
		return diag.FromErr(err)
	}

	hostsList := make([]map[string]interface{}, 0, len(hosts.Data.Hosts))

	for _, host := range hosts.Data.Hosts {
		hostMap := make(map[string]interface{})
		hostMap["free_cpu_core"] = host.FreeCPUCore
		hostMap["free_ram_size"] = host.FreeRAMSize
		hostMap["host_id"] = host.HostID
		hostMap["host_name"] = host.HostName
		hostMap["status"] = host.Status
		hostMap["total_cpu_core"] = host.TotalCPUCore
		hostMap["total_ram_size"] = host.TotalRAMSize
		hostsList = append(hostsList, hostMap)
	}

	if err := d.Set("hosts", hostsList); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("hosts")
	return diags // We return the diagnostics.
}
