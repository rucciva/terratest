output "resource_group_name" {
  value = azurerm_resource_group.ml_rg.name
}

output "workspace_name" {
  value = azurerm_machine_learning_workspace.ml_workspace.name
}


