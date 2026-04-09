package resource_fabric_vxlan

import (
	"terraform-provider-nd/internal/registry"
)

func init() {
	registry.RegisterResource(ModuleKey, NewFabricVxlanResource)
}
