package infra

import (
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the infra module.
// Resources auto-register via init() in their packages.
func GetResources() []func() resource.Resource {
	return registry.GetResources(ModuleKey)
}

// GetDataSources returns all data sources for the infra module.
// Data sources auto-register via init() in their packages.
func GetDataSources() []func() datasource.DataSource {
	return registry.GetDataSources(ModuleKey)
}
