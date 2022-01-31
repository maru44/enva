locals {
  name = "${var.name}-postgres"
}

resource "aws_security_group" "this" {
  name        = local.name
  description = local.name

  vpc_id = var.vpc_id

  tags = {
    Name = local.name
  }
}

resource "aws_security_group_rule" "postgres_ingress" {
  security_group_id = aws_security_group.this.id
  type              = "ingress"

  from_port        = 5432
  to_port          = 5432
  protocol         = "tcp"
  cidr_blocks      = var.vpc_main_cidr_blocks
  ipv6_cidr_blocks = []
  prefix_list_ids  = []
}

resource "aws_security_group_rule" "postgres_egress" {
  security_group_id = aws_security_group.this.id
  type              = "egress"

  from_port        = 0
  to_port          = 0
  protocol         = "-1"
  cidr_blocks      = ["0.0.0.0/0"]
  ipv6_cidr_blocks = ["::/0"]
  prefix_list_ids  = []
}

resource "aws_db_subnet_group" "this" {
  name        = local.name
  description = local.name
  subnet_ids  = var.subnet_ids
}

resource "aws_db_instance" "this" {
  name              = "enva0"
  allocated_storage = 20
  storage_type      = "gp2"

  auto_minor_version_upgrade = true
  apply_immediately          = false
  backup_retention_period    = 2
  backup_window              = "10:00-10:30" # UTC
  maintenance_window         = "Sun:11:00-Sun:11:30"

  vpc_security_group_ids = [aws_security_group.this.id]

  copy_tags_to_snapshot    = true
  db_subnet_group_name     = aws_db_subnet_group.this.name
  delete_automated_backups = true
  deletion_protection      = false # @TODO must true
  engine                   = "postgres"
  engine_version           = "13.4"

  enabled_cloudwatch_logs_exports = ["postgresql"]
  skip_final_snapshot             = true

  iam_database_authentication_enabled = true # managed by IAM
  identifier                          = "enva0"
  max_allocated_storage               = 0

  instance_class = "db.t4g.micro"
  #   monitoring_interval = 60 @TODO must monitor
  #   monitoring_role_arn = 
  multi_az = false

  username = var.master_user
  password = var.master_password

  port                = 5432
  publicly_accessible = false
  storage_encrypted   = false

  lifecycle {
    ignore_changes = [password]
  }
}
