/********************************
**         public web          **
********************************/

resource "aws_security_group" "web" {
  name = "enva_alb_security_group"
  description = "security group"
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "enva_alb_security_group"
  }
}

resource "aws_security_group_rule" "enva_web_https" {
  security_group_id = aws_security_group.web.id
  type = "ingress"
  
  cidr_blocks = [ "0.0.0.0/0" ]
  ipv6_cidr_blocks = [ "::/0" ]
  description = "https"
  from_port = 443
  to_port = 443
  protocol = "tcp"
}

resource "aws_security_group_rule" "enva_web_http" {
  security_group_id = aws_security_group.web.id
  type = "ingress"
  
  cidr_blocks = [ "0.0.0.0/0" ]
  ipv6_cidr_blocks = [ "::/0" ]
  description = "http"
  from_port = 80
  to_port = 80
  protocol = "tcp"
}

resource "aws_security_group_rule" "egress_web" {
  security_group_id = aws_security_group.web.id
  type = "egress"

  cidr_blocks = [ "0.0.0.0/0" ]
  ipv6_cidr_blocks = [ "::/0" ]
  description = "output web"
  prefix_list_ids = []
  from_port = 0
  to_port = 0
  protocol = "-1"
}

/********************************
**         private             **
********************************/

resource "aws_security_group" "internal" {
  name = "enva_alb_security_group_internal"
  description = "security group"
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "enva_alb_security_group_internal"
  }
}

resource "aws_security_group_rule" "enva_internal_https" {
  security_group_id = aws_security_group.internal.id
  type = "ingress"
  
  cidr_blocks = [ "${aws_vpc.main.cidr_block}" ]
  ipv6_cidr_blocks = []
  prefix_list_ids = []
  description = "https"
  from_port = 443
  to_port = 443
  protocol = "tcp"
}

resource "aws_security_group_rule" "enva_internal_http" {
  security_group_id = aws_security_group.internal.id
  type = "ingress"
  
  cidr_blocks = [ "${aws_vpc.main.cidr_block}" ]
  ipv6_cidr_blocks = []
  prefix_list_ids = []
  description = "http"
  from_port = 80
  to_port = 80
  protocol = "tcp"
}

resource "aws_security_group_rule" "egress_internal" {
  security_group_id = aws_security_group.internal.id
  type = "egress"

  cidr_blocks = [ "0.0.0.0/0" ]
  ipv6_cidr_blocks = [ "::/0" ]
  description = "output internal"
  prefix_list_ids = []
  from_port = 0
  to_port = 0
  protocol = "-1"
}

/********************************
**              lb             **
********************************/

resource "aws_lb" "main" {
  load_balancer_type = "application"
  idle_timeout = 30

  subnets = [ "${aws_subnet.public_1a.id}", "${aws_subnet.public_1c.id}" ]
  security_groups = [ "${aws_security_group.web.id}" ]
  tags = {
    Name = "enva_alb"
  }
}

resource "aws_lb_target_group" "main" {
  name = "enva"

  vpc_id = aws_vpc.main.id

  port = 80
  protocol = "HTTP"

  health_check {
    port = 80
    path = "/"
  }
}

resource "aws_lb_listener" "http" {
  port = 80
  protocol = "HTTP"

  load_balancer_arn = aws_lb.main.arn
  default_action {
    type = "redirect"
    redirect {
      port = 443
      protocol = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}

resource "aws_lb_listener" "https" {
  port = 443
  protocol = "HTTPS"

  load_balancer_arn = aws_lb.main.arn
  certificate_arn = var.api_cert_arn
  default_action {
    # type = "forward"
    # target_group_arn = aws_lb_target_group.main.id

    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      status_code  = "200"
      message_body = "ok"
    }
  }
}

