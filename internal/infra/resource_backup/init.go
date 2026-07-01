package resource_backup

import (
	"terraform-provider-nd/internal/registry"
)

func init() {
	registry.RegisterResource(ModuleKey, NewBackupResource)
}
