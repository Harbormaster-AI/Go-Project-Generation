#!/bin/bash

cd /code/gitRoot/terraform

terraform version
terraform init --input=false 
terraform plan --input=false --out=terraform.plan -var aws-access-key= -var aws-secret-key= 
terraform apply --input=false "terraform.plan"