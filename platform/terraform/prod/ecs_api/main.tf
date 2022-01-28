data "aws_region" "current" {}

data "aws_caller_identity" "current" {}

locals {
  name       = "${var.name}-api"
  account_id = data.aws_caller_identity.current.account_id
  region     = data.aws_region.current.name
}

resource "aws_ecs_task_definition" "this" {
  family = "api"
  container_definitions = jsonencode([
    {
      name                     = "nginx"
      image                    = "${var.nginx_image}"
      esseitila                = true
      tag                      = "latest"
      cpu                      = 256
      memory                   = 512
      network_mode             = "awsvpc"
      requires_compatibilities = ["FARGATE"]

      task_role_arn      = "${aws_iam_role.task_execution.arn}"
      execution_role_arn = "${aws_iam_role.task_execution.arn}"

      portMappings = [
        {
          containerPort = 80
          hostPort      = 80
        }
      ]
    },
    {
      name                     = "${local.name}"
      image                    = "${var.api_image}"
      tag                      = "latest"
      region                   = "${local.region}"
      cpu                      = 256
      memory                   = 512
      network_mode             = "awsvpc"
      requires_compatibilities = ["FARGATE"]

      task_role_arn      = "${aws_iam_role.task_execution.arn}"
      execution_role_arn = "${aws_iam_role.task_execution.arn}"

      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]
    }
  ])
}

resource "aws_cloudwatch_log_group" "this" {
  name              = "/${var.name}/ecs"
  retention_in_days = "7"
}

resource "aws_iam_role" "task_execution" {
  name = "${var.name}-TaskExecution"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ecs-tasks.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "task_execution" {
  role = aws_iam_role.task_execution.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents",
        "logs:DescribeLogGroups",
        "logs:DescribeLogStreams"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "task_execution" {
  role       = aws_iam_role.task_execution.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_lb_listener_rule" "http" {
  listener_arn = var.http_listener_arn

  action {
    type             = "forward"
    target_group_arn = var.aws_lb_target_group_id
  }

  condition {
    path_pattern {
      values = ["*"]
    }
  }
}

resource "aws_lb_listener_rule" "https" {
  listener_arn = var.https_listener_arn

  action {
    type             = "forward"
    target_group_arn = var.aws_lb_target_group_id
  }

  condition {
    path_pattern {
      values = ["*"]
    }
  }
}

resource "aws_ecs_service" "this" {
  name       = local.name
  depends_on = [aws_lb_listener_rule.http, aws_lb_listener_rule.https]

  desired_count   = 1
  cluster         = var.cluster_name
  task_definition = aws_ecs_task_definition.this.arn

  # network_configuration {
  #   security_groups  = [var.security_group_id]
  #   subnets          = var.subnet_ids
  #   assign_public_ip = true
  # }

  load_balancer {
    target_group_arn = var.aws_lb_target_group_arn
    container_name   = "nginx"
    container_port   = "80"
  }
}
