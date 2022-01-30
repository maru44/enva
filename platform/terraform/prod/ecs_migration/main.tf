data "aws_region" "current" {}

data "aws_caller_identity" "current" {}

locals {
  name       = "${var.name}-migration"
  account_id = data.aws_caller_identity.current.account_id
  region     = data.aws_region.current.name
}

resource "aws_ecs_task_definition" "migration" {
  family                   = "migration"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = 256
  memory                   = 512
  execution_role_arn       = var.task_execution_role_arn
  container_definitions = jsonencode([
    {
      name               = local.name
      image              = var.migration_image
      tag                = var.image_tag
      region             = local.region
      cpu                = 256
      memory             = 512
      execution_role_arn = var.task_execution_role_arn
    }
  ])
}

resource "aws_security_group" "this" {
  name        = local.name
  description = local.name

  vpc_id = var.vpc_id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = local.name
  }
}

resource "aws_cloudwatch_event_rule" "migration_schedule" {
  name                = "migration_schedule"
  schedule_expression = "cron(0 0 * * ? *)"
  is_enabled          = true
}

resource "aws_cloudwatch_event_target" "migration_task" {
  rule     = aws_cloudwatch_event_rule.migration_schedule.name
  arn      = var.cluster_arn
  role_arn = var.task_execution_role_arn

  ecs_target {
    task_definition_arn = aws_ecs_task_definition.migration.arn
    task_count          = 1
    launch_type         = "FARGATE"

    network_configuration {
      subnets         = var.subnet_ids
      security_groups = [aws_security_group.this.id]
    }
  }
}
