package datasource_local_user

import (
	"terraform-provider-nd/internal/registry"
)

func init() {
	registry.RegisterDataSource(ModuleKey, NewLocalUserDataSource)
}
