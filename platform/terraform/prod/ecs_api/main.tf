data "aws_region" "current" {}

data "aws_caller_identity" "current" {}

locals {
  name       = "${var.name}-api"
  account_id = data.aws_caller_identity.current.account_id
  region     = data.aws_region.current.name
}

resource "aws_ecs_task_definition" "this" {
  family                   = "api"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = 512
  memory                   = 1024
  execution_role_arn       = aws_iam_role.task_execution.arn
  container_definitions = jsonencode([
    {
      name               = "nginx"
      image              = "${var.nginx_image}"
      esseitila          = true
      tag                = "latest"
      cpu                = 256
      memory             = 512
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
      name               = "${local.name}"
      image              = "${var.api_image}"
      tag                = "latest"
      region             = "${local.region}"
      cpu                = 256
      memory             = 512
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
    target_group_arn = var.target_group_arn
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
    target_group_arn = var.target_group_arn
  }

  condition {
    path_pattern {
      values = ["*"]
    }
  }
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

resource "aws_security_group_rule" "ecs_sec_http" {
  security_group_id = aws_security_group.this.id
  type              = "ingress"

  from_port        = 80
  to_port          = 80
  protocol         = "tcp"
  cidr_blocks      = ["10.0.0.0/16"]
  ipv6_cidr_blocks = []
  prefix_list_ids  = []
}

resource "aws_security_group_rule" "ecs_sec_https" {
  security_group_id = aws_security_group.this.id
  type              = "ingress"

  from_port        = 443
  to_port          = 443
  protocol         = "tcp"
  cidr_blocks      = ["10.0.0.0/16"]
  ipv6_cidr_blocks = []
  prefix_list_ids  = []
}

resource "aws_ecs_service" "this" {
  name       = local.name
  depends_on = [aws_lb_listener_rule.http, aws_lb_listener_rule.https]

  desired_count   = 1
  cluster         = var.cluster_name
  task_definition = aws_ecs_task_definition.this.arn
  launch_type     = "FARGATE"

  network_configuration {
    subnets         = var.subnet_ids
    security_groups = [aws_security_group.this.id]
  }

  load_balancer {
    target_group_arn = var.target_group_arn
    container_name   = "nginx"
    container_port   = "80"
  }
}
