terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
  backend "s3" {
    bucket = "redbus-infra-eb1a"
    key = "rds.tfstate"
    region = "us-east-1"
  }
}
provider "aws" {
  region  = var.region
  profile = "cloud-guru"
}