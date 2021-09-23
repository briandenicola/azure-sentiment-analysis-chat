terraform {
  backend "azurerm" {
    storage_account_name = "bjdterraform001"
    container_name       = "plans"
  }
}

data "azurerm_client_config" "current" {}

resource "azurerm_resource_group" "app" {
  name                  = var.resource_group_name
  location              = var.location

  tags = {
    "Envrionment" = var.environment
    "Application"     = "Sentiment Analysis Demo Chat App"
    "Deployer"        = data.azurerm_client_config.current.client_id
    "DeploymentDate"  = "${timestamp()}"
  }
}

