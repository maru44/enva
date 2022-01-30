variable "name" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "subnet_ids" {
  type = list(string)
}

variable "cluster_name" {
  type = string
}

variable "api_image" {
  type = string
}

variable "image_tag" {
  type = string
}
