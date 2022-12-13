variable "region"{
    description = "Ec2 Region "
    default = "us-east-1"
}
variable "instance_type" {
    description = "EC2 instance Type"
    type = string
    default = "t3.micro"
}
variable "instance_keypair" {
    description = "AWS Ec2 Key pair that need to be associated"
    type = string
    default = "redbus-terraform-key"
}