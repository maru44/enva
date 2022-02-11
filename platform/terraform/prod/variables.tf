variable "aws_access_key_id" {
  type = string
}

variable "aws_secret_access_key" {
  type = string
}

variable "api_domain" {
  type = string
}

variable "api_cert_arn" {
  type = string
}

variable "database_name" {
  type = string
}

variable "database_user" {
  type = string
}

variable "database_password" {
  type = string
}

variable "ecr_api_registory" {
  type = string
}

variable "ecr_api_repository" {
  type = string
}

variable "ecr_migration_repository" {
  type = string
}

variable "rds_id" {
  type = string
}

variable "rds_name" {
  type = string
}
