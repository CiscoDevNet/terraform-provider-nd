package resource_multi_cluster_connectivity

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
func setModelId(model *MultiClusterConnectivityModel) {
	if !model.ClusterName.IsNull() && !model.ClusterName.IsUnknown() {
		model.Id = types.StringValue(model.ClusterName.ValueString())
	} else {
		model.Id = types.StringNull()
	}
}

// RscCreateMultiClusterConnectivity creates a multi cluster connectivity nd resource
func (r *multiClusterConnectivityNdResource) rscCreateMultiClusterConnectivity(ctx context.Context, dg *diag.Diagnostics, input *MultiClusterConnectivityModel) {
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

	r.rscGetMultiClusterConnectivity(ctx, dg, input)

	// Set Id from ClusterName (logic kept outside generated codec)
	setModelId(input)
}

// GetMultiClusterConnectivity retrieves multi cluster connectivity nd information by name
func (r *multiClusterConnectivityNdResource) rscGetMultiClusterConnectivity(ctx context.Context, dg *diag.Diagnostics, in *MultiClusterConnectivityModel) {

	// Preserve sensitive fields that are not returned by the API
	preservedUsername := in.Username
	preservedPassword := in.Password
	preservedLoginDomain := in.LoginDomain
	preservedMultiClusterLoginDomain := in.MultiClusterLoginDomain

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

	// ClusterName is optional when onboard a ND cluster.
	// If ClusterName is empty, API return list of clusters instead of single cluster
	if clusterAPI.ClusterName == "" {
		var clustersResp map[string][]NDFCMultiClusterConnectivityModel
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
				in.Username = preservedUsername
				in.Password = preservedPassword
				in.LoginDomain = preservedLoginDomain
				in.MultiClusterLoginDomain = preservedMultiClusterLoginDomain
				setModelId(in)
				return
			}
		}

		dg.AddError(
			"Error Reading Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not find cluster with onboardUrl %q and clusterType ND in the response", hostname),
		)
		return
	} else {
		var clusterResp NDFCMultiClusterConnectivityModel
		err = json.Unmarshal(respData, &clusterResp)
		if err != nil {
			dg.AddError(
				"Error Reading Multi Cluster Connectivity ND",
				fmt.Sprintf("Could not unmarshal multi cluster connectivity nd response, unexpected error: %v", err),
			)
			return
		}

		in.SetModelData(&clusterResp)

		// Restore sensitive fields after SetModelData (API does not return them)
		in.Username = preservedUsername
		in.Password = preservedPassword
		in.LoginDomain = preservedLoginDomain
		in.MultiClusterLoginDomain = preservedMultiClusterLoginDomain

		setModelId(in)
	}
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

// UpdateMultiClusterConnectivity updates a multi cluster connectivity nd with the provided payload
func (r *multiClusterConnectivityNdResource) rscUpdateMultiClusterConnectivity(ctx context.Context, dg *diag.Diagnostics, clusterModel *MultiClusterConnectivityModel) {
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
	// Read the updated multi cluster connectivity nd
	r.rscGetMultiClusterConnectivity(ctx, dg, clusterModel)

	// Set Id from ClusterName (logic kept outside generated codec)
	setModelId(clusterModel)

}

// DeleteMultiClusterConnectivity deletes a multi cluster connectivity nd by name
func (r *multiClusterConnectivityNdResource) rscDeleteMultiClusterConnectivity(ctx context.Context, dg *diag.Diagnostics, state *MultiClusterConnectivityModel) {
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
