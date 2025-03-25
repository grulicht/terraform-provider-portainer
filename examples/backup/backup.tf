resource "portainer_backup" "snapshot" {
  password     = var.portainer_backup_password
  output_path  = var.portainer_backup_outputh_path
}
