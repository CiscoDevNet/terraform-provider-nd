package resource_multi_cluster_connectivity_nd

import (
	"context"
	"encoding/json"
	"fmt"
	"terraform-provider-nd/internal/infra/api"

	"log"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// setModelId sets the Id field on the model based on ClusterName.
// This is kept outside resource_codec_gen.go to avoid conflicts with the internal generator.
func setModelId(model *MultiClusterConnectivityNdModel) {
	if !model.ClusterName.IsNull() && !model.ClusterName.IsUnknown() {
		model.Id = types.StringValue(model.ClusterName.ValueString())
	} else {
		model.Id = types.StringNull()
	}
}

// RscCreateMultiClusterConnectivityNd creates a multi cluster connectivity nd resource
func (r *multiClusterConnectivityNdResource) rscCreateMultiClusterConnectivityNd(ctx context.Context, dg *diag.Diagnostics, input *MultiClusterConnectivityNdModel) {
	if input == nil {
		dg.AddError(
			"Invalid Input",
			"The input model is nil",
		)
		return
	}

	inData := input.GetModelData()

	// Create multi cluster connectivity nd API client
	clusterAPI := api.NewClusterAPI(nil, r.infraClient.ApiClient)

	// Convert model data to JSON
	clusterPayload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Creating Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not create multi cluster connectivity nd, Data Marshall error: %v", err),
		)
		return
	}

	// Call the API to create the multi cluster connectivity nd
	res, err := clusterAPI.Post(clusterPayload)
	if err != nil {
		dg.AddError(
			"Error Creating Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not create multi cluster connectivity nd, unexpected error: %v %v", err, res),
		)
		return
	}

	// Preserve sensitive fields that are not returned by the API
	preservedUsername := input.Username
	preservedPassword := input.Password
	preservedLoginDomain := input.LoginDomain
	preservedMultiClusterLoginDomain := input.MultiClusterLoginDomain

	r.rscGetMultiClusterConnectivityNd(ctx, dg, input)

	// Restore sensitive fields after read
	input.Username = preservedUsername
	input.Password = preservedPassword
	input.LoginDomain = preservedLoginDomain
	input.MultiClusterLoginDomain = preservedMultiClusterLoginDomain

	// Set Id from ClusterName (logic kept outside generated codec)
	setModelId(input)
}

// GetMultiClusterConnectivityNd retrieves multi cluster connectivity nd information by name
func (r *multiClusterConnectivityNdResource) rscGetMultiClusterConnectivityNd(ctx context.Context, dg *diag.Diagnostics, in *MultiClusterConnectivityNdModel) {

	clusterAPI := api.NewClusterAPI(nil, r.infraClient.ApiClient)
	clusterAPI.ClusterName = in.ClusterName.ValueString()
	respData, err := clusterAPI.Get()

	if err != nil {
		dg.AddError(
			"Error Reading Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not read multi cluster connectivity nd, unexpected error: %v %v", err, respData),
		)
		return
	}

	if clusterAPI.ClusterName == "" {
		var clustersResp map[string][]NDFCMultiClusterConnectivityNdModel
		err = json.Unmarshal(respData, &clustersResp)
		if err != nil {
			dg.AddError(
				"Error Reading Multi Cluster Connectivity ND",
				fmt.Sprintf("Could not unmarshal multi cluster connectivity nd response, unexpected error: %v", err),
			)
			return
		}

		hostname := in.Hostname.ValueString()
		clusterType := in.ClusterType.ValueString()
		for _, cluster := range clustersResp["clusters"] {
			if cluster.Spec.Hostname == hostname && cluster.Spec.ClusterType == clusterType {
				in.SetModelData(&cluster)
				setModelId(in)
				return
			}
		}

		dg.AddError(
			"Error Reading Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not find cluster with onboardUrl %q and clusterType %q in the response", hostname, clusterType),
		)
		return
	}

	var clusterResp NDFCMultiClusterConnectivityNdModel
	err = json.Unmarshal(respData, &clusterResp)
	if err != nil {
		dg.AddError(
			"Error Reading Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not unmarshal multi cluster connectivity nd response, unexpected error: %v", err),
		)
		return
	}

	in.SetModelData(&clusterResp)
	setModelId(in)
}

// updateSpecValue extends NDFCSpecValue with fields only needed during update.
type updateSpecValue struct {
	NDFCSpecValue
	ReRegister *bool `json:"reRegister,omitempty"`
}

// updatePayload wraps the update spec for JSON marshalling.
type updatePayload struct {
	Spec updateSpecValue `json:"spec,omitempty"`
}

// UpdateMultiClusterConnectivityNd updates a multi cluster connectivity nd with the provided payload
func (r *multiClusterConnectivityNdResource) rscUpdateMultiClusterConnectivityNd(ctx context.Context, dg *diag.Diagnostics, clusterModel *MultiClusterConnectivityNdModel) {
	inData := clusterModel.GetModelData()

	clusterAPI := api.NewClusterAPI(nil, r.infraClient.ApiClient)
	clusterAPI.ClusterName = clusterModel.ClusterName.ValueString()

	// This is only used for the update operation and not for create, as create will register the cluster for the first time.
	// For update, we want to ensure that the changes are applied by re-registering the cluster with the new details.
	reRegister := true
	payload := updatePayload{
		Spec: updateSpecValue{
			NDFCSpecValue: inData.Spec,
			ReRegister:    &reRegister,
		},
	}

	inDataBytes, err := json.Marshal(payload)
	if err != nil {
		dg.AddError(
			"Error Updating Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not update multi cluster connectivity nd, Data Marshall error: %v", err),
		)
		log.Printf("[ERROR] Error Updating Multi Cluster Connectivity ND: error=%s", err.Error())
		return
	}
	res, err := clusterAPI.Put(inDataBytes)

	if err != nil {
		dg.AddError(
			"Error Updating Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not update multi cluster connectivity nd, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Updating Multi Cluster Connectivity ND: error=%s", err.Error())
		return
	}
	// Preserve sensitive fields that are not returned by the API
	preservedUsername := clusterModel.Username
	preservedPassword := clusterModel.Password
	preservedLoginDomain := clusterModel.LoginDomain
	preservedMultiClusterLoginDomain := clusterModel.MultiClusterLoginDomain

	// Read the updated multi cluster connectivity nd
	r.rscGetMultiClusterConnectivityNd(ctx, dg, clusterModel)

	// Restore sensitive fields after read
	clusterModel.Username = preservedUsername
	clusterModel.Password = preservedPassword
	clusterModel.LoginDomain = preservedLoginDomain
	clusterModel.MultiClusterLoginDomain = preservedMultiClusterLoginDomain

	// Set Id from ClusterName (logic kept outside generated codec)
	setModelId(clusterModel)

}

// DeleteMultiClusterConnectivityNd deletes a multi cluster connectivity nd by name
func (r *multiClusterConnectivityNdResource) rscDeleteMultiClusterConnectivityNd(ctx context.Context, dg *diag.Diagnostics, state *MultiClusterConnectivityNdModel) {
	clusterAPI := api.NewClusterAPI(nil, r.infraClient.ApiClient)
	clusterAPI.ClusterName = state.ClusterName.ValueString()

	// Build the remove payload with credentials and force flag
	removePayload := api.ClusterRemovePayload{
		Credentials: api.ClusterRemoveCredentials{
			User:     state.Username.ValueString(),
			Password: state.Password.ValueString(),
		},
	}

	if !state.LoginDomain.IsNull() && !state.LoginDomain.IsUnknown() {
		removePayload.Credentials.LoginDomain = state.LoginDomain.ValueString()
	}

	payload, err := json.Marshal(removePayload)
	if err != nil {
		dg.AddError(
			"Error Deleting Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not delete multi cluster connectivity nd, Data Marshall error: %v", err),
		)
		return
	}

	res, err := clusterAPI.PostDelete(payload)
	if err != nil {
		dg.AddError(
			"Error Deleting Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not delete multi cluster connectivity nd, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Deleting Multi Cluster Connectivity ND: error=%s", err.Error())
		return
	}
}
