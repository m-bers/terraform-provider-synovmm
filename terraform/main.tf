terraform {
  required_providers {
    synovmm = {
      version = "0.2"
      source  = "terraform.local/m-bers/synovmm"
    }
  }
}

provider "synovmm" {
  host     = var.synology_host
  username = var.synology_username
  password = var.synology_password
}

data "synovmm_host" "mbers" {}

output "hosts" {
  value = data.synovmm_host.mbers.hosts
}

output "ds918_host_id" {
  value = data.synovmm_host.mbers.hosts[0].host_id
}
