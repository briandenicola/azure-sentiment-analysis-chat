resource "azurerm_user_assigned_identity" "app_identity" {
  name                = "${var.app_name}-identity"
  resource_group_name = azurerm_resource_group.app.name
  location            = azurerm_resource_group.app.location
}
