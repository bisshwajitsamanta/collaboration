terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
  backend "s3" {
    bucket = "redbus-infra-a6bb"
    key = "ec2instance.tfstate"
    region = "us-east-1"
  }
}
provider "aws" {
  region  = var.region
  profile = "cloud-guru"
}