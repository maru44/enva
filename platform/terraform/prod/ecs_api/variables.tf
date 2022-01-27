variable "name" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "subnet_ids" {
  type = list(string)
}

variable "https_listener_arn" {
  type = string
}

variable "db_host" {
  type = string
}

variable "cluster_name" {
  type = string
}

variable "nginx_image" {
  type = string
}

variable "api_image" {
  type = string
}
