module "ec2_public_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "4.16.2"
  name = "${local.name}-public-sg"
  description = "Security Group with SSH port"
#  vpc_id =
  #Ingress Rules
  ingress_rules = ["ssh-tcp"]
  ingress_cidr_blocks = ["0.0.0.0/0"]
  #Egress Rules
  egress_rules = ["all-all"]
  tags = local.tags
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 3.0"

  name = local.name
  cidr = "10.99.0.0/18"

  azs              = ["${local.region}a", "${local.region}b", "${local.region}c"]
  public_subnets   = ["10.99.0.0/24", "10.99.1.0/24", "10.99.2.0/24"]
  private_subnets  = ["10.99.3.0/24", "10.99.4.0/24", "10.99.5.0/24"]

  enable_dns_hostnames = true
  enable_dns_support = true

}