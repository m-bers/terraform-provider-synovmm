package synovmm

type Host struct {
	Data struct {
		Hosts []struct {
			FreeCPUCore  int    `json:"free_cpu_core"`
			FreeRAMSize  int    `json:"free_ram_size"`
			HostID       string `json:"host_id"`
			HostName     string `json:"host_name"`
			Status       string `json:"status"`
			TotalCPUCore int    `json:"total_cpu_core"`
			TotalRAMSize int    `json:"total_ram_size"`
		} `json:"hosts"`
	} `json:"data"`
	Success bool `json:"success"`
}
