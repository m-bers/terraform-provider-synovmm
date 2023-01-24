#!/bin/bash
go build -o terraform-provider-synovmm
mv terraform-provider-synovmm ~/.terraform.d/plugins/terraform.local/m-bers/synovmm/0.2/linux_arm64/
cd terraform
rm -rf .terraform terraform.tfstate .terraform.lock.hcl
# terraform init
# terraform apply -auto-approve
cd ..