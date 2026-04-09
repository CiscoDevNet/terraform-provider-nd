package manage

import (
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the manage module.
// Resources auto-register via init() in their packages.
func GetResources() []func() resource.Resource {
	return registry.GetResources(ModuleKey)
}

// GetDataSources returns all data sources for the manage module.
// Data sources auto-register via init() in their packages.
func GetDataSources() []func() datasource.DataSource {
	return registry.GetDataSources(ModuleKey)
}
