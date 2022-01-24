# variable "name" {
#   type = string
# }

# variable "vpc_id" {
#   type = string
# }

# variable "subnet_ids" {
#   type = "list"
# }

# variable "database_name" {
#   type = string
# }

# variable "master_user" {
#   type = string
# }

# variable "master_password" {
#   type = string
# }

# variable "vpc_main_cidr_blocks" {
#   type = "list"
# }

# locals {
#   name = "${var.name}-postgres"
# }

# resource "aws_security_group" "this" {
#   name = local.name
#   description = local.name

#   vpc_id = var.vpc_id

#   tags = {
#     Name = local.name
#   }
# }

# resource "aws_security_group_rule" "postgres_ingress" {
#   security_group_id = aws_security_group.this.id
#   type = "ingress"

#   from_port = 5432
#   to_port = 5432
#   protocol = "tcp"
#   cidr_blocks = var.vpc_main_cidr_blocks
#   ipv6_cidr_blocks = [  ]
#   prefix_list_ids = []
# }

# resource "aws_security_group_rule" "postgres_egress" {
#   security_group_id = aws_security_group.this.id
#   type = "egress"

#   from_port = 0
#   to_port = 0
#   protocol = "-1"
#   cidr_blocks = [ "0.0.0.0/16" ]
#   ipv6_cidr_blocks = [ "::/0" ]
#   prefix_list_ids = []
# }

# resource "aws_db_subnet_group" "this" {
#   name = local.name
#   description = local.name
#   subnet_ids = var.subnet_ids
# }

# resource "aws_rds_cluster" "this" {
#   cluster_identifier = local.name

#   db_subnet_group_name = aws_db_subnet_group.this.name
#   vpc_security_group_ids = ["${aws_db_subnet_group.this.id}"]

#   engine = "postgres"
#   port = "5432"

#   database_name = var.database_name
#   master_username = var.master_user
#   master_password = var.master_password
# }

# resource "aws_rds_cluster_instance" "this" {
#   identifier = local.name
#   cluster_identifier = aws_rds_cluster.this.id

#   engine = "postgres"
#   instance_class = "db.t4.small"
# }

# output "endpoint" {
#   value = "${aws_rds_cluster.this.endpoint}"
# }
