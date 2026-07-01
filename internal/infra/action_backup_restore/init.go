package action_backup_restore

import (
	"terraform-provider-nd/internal/infra/resource_backup"
	"terraform-provider-nd/internal/registry"
)

func init() {
	registry.RegisterAction(resource_backup.ModuleKey, NewBackupRestoreAction)
}
