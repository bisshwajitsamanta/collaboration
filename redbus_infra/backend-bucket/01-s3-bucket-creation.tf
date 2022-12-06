resource "random_id" "randomGenerator" {
  byte_length = 2
}

resource "aws_s3_bucket" "redbus_infra_bucket" {
  bucket = "${var.bucketName}-${random_id.randomGenerator.hex}"

  versioning {
    enabled = true
  }
}

