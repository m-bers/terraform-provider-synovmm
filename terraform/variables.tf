variable "synology_host" {
  type        = string
  description = "Synology DSM VMM host with scheme and port (e.g.: http://10.10.10.10:5000)"
}

variable "synology_username" {
  type        = string
  description = "Synology DSM username"
}

variable "synology_password" {
  type        = string
  description = "Synology DSM password"
}
