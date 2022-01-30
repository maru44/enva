variable "name" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "subnet_ids" {
  type = list(string)
}

variable "master_user" {
  type = string
}

variable "master_password" {
  type = string
}

variable "vpc_main_cidr_blocks" {
  type = list(string)
}
