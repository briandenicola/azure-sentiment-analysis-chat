resource "azurerm_cognitive_account" "app" {
  name                = "${var.app_name}-cogs01"
  resource_group_name = azurerm_resource_group.app.name
  location            = azurerm_resource_group.app.location
  kind                = "TextAnalytics"

  sku_name            = "S0"
}

resource "azurerm_key_vault_secret" "azurerm_cognitive_account_key" {
  name         = "${var.app_name}-cogs-key"
  value        = azurerm_cognitive_account.app.primary_access_key
  key_vault_id = azurerm_key_vault.app.id
}
