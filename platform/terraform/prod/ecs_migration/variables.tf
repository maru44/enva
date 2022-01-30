variable "migration_image" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "subnet_ids" {
  type = list(string)
}

variable "cluster_arn" {
  type = string
}

variable "name" {
  type = string
}

variable "task_execution_role_arn" {
  type = string
}
