variable "name" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "subnet_ids" {
  type = list(string)
}

variable "http_listener_arn" {
  type = string
}

variable "https_listener_arn" {
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

# variable "aws_lb_target_group_id" {
#   type = string
# }

# variable "aws_lb_target_group_arn" {
#   type = string
# }
