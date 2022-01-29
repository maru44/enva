provider "aws" {
  region     = "ap-northeast-1"
  access_key = var.aws_access_key_id
  secret_key = var.aws_secret_access_key
}

/********************************
**           network           **
********************************/

module "network" {
  source = "./network"
}

/********************************
**              lb             **
********************************/

resource "aws_acm_certificate_validation" "main" {
  certificate_arn = var.api_cert_arn
}

module "alb" {
  source = "./alb"

  vpc_id            = module.network.vpc_id
  vpc_cidr_block    = module.network.vpc_cidr_block
  public_subnet_ids = module.network.public_subnet_ids
  certificate_arn   = var.api_cert_arn
}

/********************************
**             ecr             **
********************************/

resource "aws_ecr_repository" "nginx" {
  name                 = var.ecr_nginx_repository
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_repository" "api" {
  name                 = var.ecr_api_repository
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_vpc_endpoint" "ecr_dkr" {
  vpc_id              = module.network.vpc_id
  service_name        = "com.amazonaws.ap-northeast-1.ecr.dkr"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = module.network.private_subnet_ids
  security_group_ids  = [module.alb.aws_security_group_internal_id]
  private_dns_enabled = true
}

resource "aws_vpc_endpoint" "ecr_api" {
  vpc_id              = module.network.vpc_id
  service_name        = "com.amazonaws.ap-northeast-1.ecr.api"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = module.network.private_subnet_ids
  security_group_ids  = [module.alb.aws_security_group_internal_id]
  private_dns_enabled = true
}

/********************************
**         rds module          **
********************************/

module "rds" {
  source = "./rds"
  name   = var.rds_name

  vpc_id               = module.network.vpc_id
  vpc_main_cidr_blocks = [module.network.vpc_cidr_block]
  subnet_ids           = module.network.public_subnet_ids

  database_name   = var.database_name
  master_user     = var.database_user
  master_password = var.database_password
}

/********************************
**       ecs module           **
********************************/

resource "aws_ecs_cluster" "main" {
  name = "enva"
}

module "ecs_api" {
  source = "./ecs_api"

  name               = "enva"
  vpc_id             = module.network.vpc_id
  subnet_ids         = module.network.private_subnet_ids
  http_listener_arn  = module.alb.aws_lb_listener_http_arn
  https_listener_arn = module.alb.aws_lb_listener_https_arn
  cluster_name       = aws_ecs_cluster.main.name

  nginx_image = "${var.ecr_nginx_registory}/${var.ecr_nginx_repository}:latest"
  api_image   = "${var.ecr_api_registory}/${var.ecr_api_repository}:latest"
}
