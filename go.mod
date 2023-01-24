module terraform-provider-synovmm

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.0.0-rc.2
	github.com/m-bers/synovmm-client-go v0.0.0-00010101000000-000000000000
)

replace github.com/m-bers/synovmm-client-go => ./synovmm-client-go
