resource "aws_ecs_task_definition" "main" {
  family = "enva"

  requires_compatibilities = [ "FARGATE" ]

  cpu = 256
  memory = 512

  network_mode = "awsvpc"

  container_definitions = <<EOL
[
    {
        "name": "nginx",
        "image": "nginx:1.21.5-alpine",
        "portMappings": [
            {
                "containerPort": 80,
                "hostPort": 80
            }
        ]
    }
]
EOL
}

resource "aws_ecs_cluster" "main" {
  name = "enva"
}

resource "aws_lb_target_group" "main" {
  name = "enva"

  vpc_id = aws_vpc.main.id

  port = 80
  protocol = "HTTP"
  target_type = "ip"

  health_check {
    port = 80
    path = "/"
  }
}

resource "aws_lb_listener_rule" "main" {
  listener_arn = aws_lb_listener.main.arn

  action {
    type = "forward"
    target_group_arn = aws_lb_target_group.main.id
  }

  condition {
    path_pattern {
      values = ["*"]
    }
  }
}

resource "aws_ecs_service" "main" {
  name = "enva"

  depends_on = [aws_lb_listener_rule.main]
  cluster = aws_ecs_cluster.main.name
  launch_type = "FARGATE"
  desired_count = 1
  network_configuration {
    subnets = [ "${aws_subnet.private_1a.id}", "${aws_subnet.private_1c.id}" ]
    security_groups = [ "${aws_security_group.internal.id}" ]
  }

  load_balancer {
    target_group_arn = "${aws_lb_target_group.main.arn}"
    container_name = "nginx"
    container_port = 80
  }

  task_definition = "${aws_ecs_task_definition.main.arn}"
}

# resource "aws_acm_certificate" "main" {
#   domain_name = var.api_domain
# #   subject_alternative_names = [ "value" ]
#   validation_method = "DNS"

#   lifecycle {
#     create_before_destroy = true
#   }

#   tags = {
#     Name = "enva_acm"
#   }
# }

resource "aws_acm_certificate_validation" "main" {
  certificate_arn = var.api_cert_arn
}
