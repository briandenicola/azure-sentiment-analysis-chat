variable "location" {
  description = "The Azure Region to deploy AKS"
  default     = "centralus"
}

variable "app_name" {
  description = "The name of this application"
}

variable "resource_group_name" {
  description = "The Azure Resource Group to deploy this application"
}

variable "environment" {
  description = "The environment this cluster is"
}
