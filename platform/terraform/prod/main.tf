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
  target_group_arn  = module.ecs_api.target_group_arn
  certificate_arn   = var.api_cert_arn
  domain            = var.api_domain
}

/********************************
**             ecr             **
********************************/

resource "aws_ecr_repository" "api" {
  name                 = var.ecr_api_repository
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_repository" "migration" {
  name                 = var.ecr_migration_repository
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_vpc_endpoint" "ecr_dkr" {
  vpc_id              = module.network.vpc_id
  service_name        = "com.amazonaws.ap-northeast-1.ecr.dkr"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = module.network.private_subnet_ids
  security_group_ids  = [module.ecs_api.security_group_id]
  private_dns_enabled = true
}

resource "aws_vpc_endpoint" "ecr_api" {
  vpc_id              = module.network.vpc_id
  service_name        = "com.amazonaws.ap-northeast-1.ecr.api"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = module.network.private_subnet_ids
  security_group_ids  = [module.ecs_api.security_group_id]
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

  name         = "enva"
  vpc_id       = module.network.vpc_id
  subnet_ids   = module.network.private_subnet_ids
  cluster_name = aws_ecs_cluster.main.name

  api_image = "${var.ecr_api_registory}/${var.ecr_api_repository}:latest"
}

module "ecs_migration" {
  source = "./ecs_migration"

  name                    = "enva"
  vpc_id                  = module.network.vpc_id
  subnet_ids              = module.network.private_subnet_ids
  cluster_arn             = aws_ecs_cluster.main.arn
  task_execution_role_arn = module.ecs_api.task_execution_role_arn
  migration_image         = "${var.ecr_api_registory}/${var.ecr_migration_repository}:latest"
}
