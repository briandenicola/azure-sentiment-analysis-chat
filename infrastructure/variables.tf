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

variable "app_vnet_resource_group_name" {
  description = "The Resource Group name that contains the Vnet for Private Endpoints"
}

variable "app_vnet" {
  description = "The vnet name where the Private Endpoints will be created"
}

variable "dns_resource_group_name" {
  description = "The Resource Group name that contains Private DNS Zones"
}

variable "github_actions_identity_name" {
  description = "The name of the Github Task runner Managed Identity"
}

variable "github_actions_identity_resource_group" {
  description = "The Resource Group name that Github Taskrunner Identity"
}