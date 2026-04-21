package resource_multi_cluster_connectivity_nd

import (
	"terraform-provider-nd/internal/registry"
)

func init() {
	registry.RegisterResource(ModuleKey, NewMultiClusterConnectivityNdResource)
}
