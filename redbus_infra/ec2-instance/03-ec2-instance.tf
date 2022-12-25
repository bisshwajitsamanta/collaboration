locals {
  availability_zone = "${local.region}a"
  name              = "redbus-jumphost"
  region            = "us-east-1"
  tags = {
    Owner       = "Bisshwajit"
    Environment = "Dev"
  }
}

module "ec2_public" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "4.2.1"

  name = "${local.name}-instance"

  ami                    = data.aws_ami.amazonLinux2.id
  instance_type          = var.instance_type
  key_name               = var.instance_keypair
  monitoring             = true
  vpc_security_group_ids = [module.ec2_public_sg.security_group_id]
  subnet_id              = "subnet-eddcdzz4"

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}