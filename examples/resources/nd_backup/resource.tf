
resource "nd_backup" "test_resource_backup_1" {
  name           = "tf-backup-123"
  type           = "configOnly"
  destination    = "sftp-server"
  encryption_key = "backupKey123"
  telemetry_data = false
  timeouts = {
    create = "90m"
    read   = "30s"
  }
}