package resource_remote_storage_location

import (
	"terraform-provider-nd/internal/registry"
)

func init() {
	registry.RegisterResource(ModuleKey, NewRemoteStorageLocationResource)
}
