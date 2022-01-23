provider "aws" {
  region = "ap-northeast-1"
  access_key = var.aws_access_key_id
  secret_key = var.aws_secret_access_key
}

resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = "enva"
  }
}

resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id

  tags = {
      Name = "enva_gw"
  }
}

# resource "aws_subnet" "private" {
#   vpc_id = aws_vpc.main.id
#   availability_zone = "ap-northeast-1a"
#   cidr_block = "10.0.0.1/16"
# }
