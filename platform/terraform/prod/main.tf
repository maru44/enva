provider "aws" {
  region = "ap-northeast-1"
  access_key = var.aws_access_key_id
  secret_key = var.aws_secret_access_key
}

resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support = true

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

resource "aws_subnet" "private_1a" {
  vpc_id = aws_vpc.main.id
  availability_zone = "ap-northeast-1a"
  cidr_block = "10.0.10.0/24"

  tags = {
    Name = "enva-private-subnet_1a"
  }
}

resource "aws_subnet" "private_1c" {
  vpc_id = aws_vpc.main.id
  availability_zone = "ap-northeast-1c"
  cidr_block = "10.0.20.0/24"

  tags = {
    Name = "enva-private-subnet_1c"
  }
}

resource "aws_subnet" "public_1a" {
  vpc_id = aws_vpc.main.id
  availability_zone = "ap-northeast-1a"
  cidr_block = "10.0.1.0/24"

  tags = {
    Name = "enva-public-subnet_1a"
  }
}

resource "aws_subnet" "public_1c" {
  vpc_id = aws_vpc.main.id
  availability_zone = "ap-northeast-1c"
  cidr_block = "10.0.2.0/24"

  tags = {
    Name = "enva-public-subnet_1c"
  }
}

/********************************
**             ecr             **
********************************/

resource "aws_ecr_repository" "nginx" {
  name = "enva-nginx"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_repository" "api" {
  name = "enva0"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

/********************************
**              nat             **
********************************/

resource "aws_eip" "nat_1a" {
  vpc = true

  tags = {
    Name = "enva_natgw_1a"
  }
}

resource "aws_nat_gateway" "nat_1a" {
  subnet_id = "${aws_subnet.public_1a.id}"
  allocation_id = "${aws_eip.nat_1a.id}"

  tags = {
    Name = "enva_public_1a"
  }
}

resource "aws_eip" "nat_1c" {
  vpc = true

  tags = {
    Name = "enva_natgw_1c"
  }
}

resource "aws_nat_gateway" "nat_1c" {
  subnet_id = "${aws_subnet.public_1c.id}"
  allocation_id = "${aws_eip.nat_1c.id}"

  tags = {
    Name = "enva_public_1c"
  }
}

/********************************
**       route for internet            **
********************************/

resource "aws_route_table" "public" {
  vpc_id = "${aws_vpc.main.id}"

  tags = {
    Name = "enva_public"
  }
}

resource "aws_route" "public" {
  destination_cidr_block = "0.0.0.0/0"
  route_table_id = "${aws_route_table.public.id}"
  gateway_id = "${aws_internet_gateway.main.id}"
}

resource "aws_route_table_association" "public_1a" {
  subnet_id = aws_subnet.public_1a.id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "public_1c" {
  subnet_id = aws_subnet.public_1c.id
  route_table_id = aws_route_table.public.id
}

/********************************
**       route for internet            **
********************************/

resource "aws_route_table" "private_1a" {
  vpc_id = aws_vpc.main.id
  
  tags = {
    Name = "enva_private_1a"
  }
}

resource "aws_route_table" "private_1c" {
  vpc_id = aws_vpc.main.id
  
  tags = {
    Name = "enva_private_1c"
  }
}

resource "aws_route" "private_1a" {
  destination_cidr_block = "0.0.0.0/0"
  route_table_id = "${aws_route_table.private_1a.id}"
  nat_gateway_id = "${aws_nat_gateway.nat_1a.id}"
}

resource "aws_route" "private_1c" {
  destination_cidr_block = "0.0.0.0/0"
  route_table_id = "${aws_route_table.private_1c.id}"
  nat_gateway_id = "${aws_nat_gateway.nat_1c.id}"
}

resource "aws_route_table_association" "private_1a" {
  subnet_id      = "${aws_subnet.private_1a.id}"
  route_table_id = "${aws_route_table.private_1a.id}"
}

resource "aws_route_table_association" "private_1c" {
  subnet_id      = "${aws_subnet.private_1c.id}"
  route_table_id = "${aws_route_table.private_1c.id}"
}

/********************************
**         rds module          **
********************************/

module "rds" {
  source = "./rds"
  name = "${var.rds_name}"

  vpc_id = aws_vpc.main.id
  vpc_main_cidr_blocks = [ "10.0.0.0/16" ]
  subnet_ids = ["${aws_subnet.private_1a.id}","${aws_subnet.private_1c.id}"]

  database_name = var.database_name
  master_user = var.database_user
  master_password = var.database_password
}

/********************************
**       ecs module           **
********************************/

resource "aws_ecs_cluster" "main" {
  name = "enva"
}

resource "aws_acm_certificate_validation" "main" {
  certificate_arn = var.api_cert_arn
}

module "ecs_api" {
  source = "./ecs_api"

  name = "enva"
  vpc_id = aws_vpc.main.id
  subnet_ids = [ "${aws_subnet.public_1a.id}", "${aws_subnet.public_1c.id}" ]
  https_listener_arn = "${aws_lb_listener.main.arn}"
  cluster_name = "${aws_ecs_cluster.main.name}"

  db_host = "${module.rds.endpoint}"

  nginx_image = "${var.ecr_api_registory}/${var.ecr_api_repository}:latest"
  api_image = "${var.ecr_nginx_registory}/${var.ecr_nginx_repository}:latest"
}
