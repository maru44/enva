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
  cpu                      = 256
  memory                   = 512
  execution_role_arn       = aws_iam_role.task_execution.arn
  container_definitions = jsonencode([
    {
      name               = "${local.name}"
      image              = "${var.api_image}"
      tag                = "${var.image_tag}"
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

resource "aws_lb_target_group" "this" {
  name   = "enva"
  vpc_id = var.vpc_id

  port        = 8080
  protocol    = "HTTP"
  target_type = "ip"

  health_check {
    port = 8080
    path = "/health"
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

resource "aws_security_group_rule" "this" {
  security_group_id = aws_security_group.this.id
  type              = "ingress"

  from_port        = 80
  to_port          = 8080
  protocol         = "tcp"
  cidr_blocks      = ["10.0.0.0/16"]
  ipv6_cidr_blocks = []
  prefix_list_ids  = []
}

resource "aws_ecs_service" "this" {
  name = local.name
  # depends_on = [aws_lb_listener_rule.http, aws_lb_listener_rule.https]

  desired_count   = 1
  cluster         = var.cluster_name
  task_definition = aws_ecs_task_definition.this.arn
  launch_type     = "FARGATE"

  network_configuration {
    subnets         = var.subnet_ids
    security_groups = [aws_security_group.this.id]
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.this.arn
    container_name   = "enva-api"
    container_port   = 8080
  }
}

output "target_group_arn" {
  value = aws_lb_target_group.this.arn
}

output "security_group_id" {
  value = aws_security_group.this.id
}

output "task_execution_role_arn" {
  value = aws_iam_role.task_execution.arn
}
