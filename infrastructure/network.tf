data "azurerm_virtual_network" "vnet" {
  name                = var.app_vnet
  resource_group_name = var.app_vnet_resource_group_name
}

data "azurerm_subnet" "private_endpoint_subnet" {
  name                 = "private-endpoints"
  virtual_network_name = var.app_vnet
  resource_group_name  = var.app_vnet_resource_group_name
}

data "azurerm_private_dns_zone" "privatelink_vaultcore_azure_net" {
  name                      = "privatelink.vaultcore.azure.net"
  resource_group_name       = var.dns_resource_group_name
}
