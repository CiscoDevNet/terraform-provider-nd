package resource_multi_cluster_connectivity

import (
	"terraform-provider-nd/internal/registry"
)

func init() {
	registry.RegisterResource(ModuleKey, NewMultiClusterConnectivityResource)
}
