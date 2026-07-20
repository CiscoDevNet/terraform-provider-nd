package resource_multi_cluster_connectivity

import (
	"context"
	"fmt"

	"terraform-provider-nd/internal/infra"
	"terraform-provider-nd/internal/registry"

	"log"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ModuleKey is the key used to get the infra module from the provider.
const ModuleKey = "infra"

// Ensure the implementation satisfies the expected interfaces
var (
	_ resource.Resource                = &multiClusterConnectivityNdResource{}
	_ resource.ResourceWithConfigure   = &multiClusterConnectivityNdResource{}
	_ resource.ResourceWithImportState = &multiClusterConnectivityNdResource{}
)

// NewMultiClusterConnectivityResource is a helper function to simplify the provider implementation.
func NewMultiClusterConnectivityResource() resource.Resource {
	return &multiClusterConnectivityNdResource{}
}

// multiClusterConnectivityNdResource is the resource implementation.
type multiClusterConnectivityNdResource struct {
	infraClient *infra.NexusDashboardInfra
}

// Metadata returns the resource type name.
func (r *multiClusterConnectivityNdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_multi_cluster_connectivity"
}

// Schema defines the schema for the resource.
func (r *multiClusterConnectivityNdResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = MultiClusterConnectivityResourceSchema(ctx)
}

// Configure adds the provider configured client to the resource.
func (r *multiClusterConnectivityNdResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected registry.ClientProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	infraModule := client.GetModule(ModuleKey)
	if infraModule == nil {
		resp.Diagnostics.AddError(
			"Infra Module Not Found",
			"The infra module was not registered with the provider.",
		)
		return
	}

	r.infraClient = infraModule.(*infra.NexusDashboardInfra)
}

// Create creates the resource and sets the initial Terraform state.
func (r *multiClusterConnectivityNdResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	log.Printf("[DEBUG] Start create of resource: nd_multi_cluster_connectivity")

	var in MultiClusterConnectivityModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if !in.ClusterName.IsNull() && !in.ClusterName.IsUnknown() {
		in.Id = in.ClusterName
	}
	log.Printf("[DEBUG] Creating Multi Cluster Connectivity ND: id=%s", in.Id.ValueString())

	r.rscCreateMultiClusterConnectivity(&resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("[DEBUG] End create of resource nd_multi_cluster_connectivity with id '%s'", in.Id.ValueString())
}

// Read refreshes the Terraform state with the latest data.
func (r *multiClusterConnectivityNdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("[DEBUG] Start read of resource: nd_multi_cluster_connectivity")

	var state MultiClusterConnectivityModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := getMultiClusterConnectivityID(&state)
	log.Printf("[DEBUG] Reading Multi Cluster Connectivity ND: id=%s", id)

	if r.rscGetMultiClusterConnectivity(&resp.Diagnostics, &state) {
		log.Printf("[WARN] Multi Cluster Connectivity ND not found; removing resource from state: id=%s", id)
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("[DEBUG] End read of resource nd_multi_cluster_connectivity with id '%s'", state.Id.ValueString())
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *multiClusterConnectivityNdResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("[DEBUG] Start update of resource: nd_multi_cluster_connectivity")
	var plan MultiClusterConnectivityModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if !plan.ClusterName.IsNull() && !plan.ClusterName.IsUnknown() {
		plan.Id = plan.ClusterName
	}
	log.Printf("[DEBUG] Updating Multi Cluster Connectivity ND: id=%s", plan.Id.ValueString())

	r.rscUpdateMultiClusterConnectivity(&resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("[DEBUG] End update of resource nd_multi_cluster_connectivity with id '%s'", plan.Id.ValueString())
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *multiClusterConnectivityNdResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("[DEBUG] Start delete of resource: nd_multi_cluster_connectivity")
	var state MultiClusterConnectivityModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := getMultiClusterConnectivityID(&state)
	r.rscDeleteMultiClusterConnectivity(&resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("[DEBUG] End delete of resource nd_multi_cluster_connectivity with id '%s'", id)
}

// ImportState imports a multi cluster connectivity nd resource by cluster name.
func (r *multiClusterConnectivityNdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	log.Printf("[DEBUG] Start import state of resource: nd_multi_cluster_connectivity")

	var state MultiClusterConnectivityModel
	state.ClusterName = types.StringValue(req.ID)

	if r.rscGetMultiClusterConnectivity(&resp.Diagnostics, &state) {
		resp.Diagnostics.AddError(
			"Error Importing Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not import nd multi cluster connectivity with id %q: resource not found", req.ID),
		)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddWarning(
		"Sensitive Attributes Not Imported",
		"The values for `username`, `password`, `login_domain` and `multi_cluster_login_domain` attributes will not be imported when the nd_multi_cluster_connectivity resource imports an already registered cluster from Nexus Dashboard.",
	)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End import of state resource: nd_multi_cluster_connectivity")
}
