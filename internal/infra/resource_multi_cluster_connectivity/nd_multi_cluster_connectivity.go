package resource_multi_cluster_connectivity

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// setMultiClusterConnectivityID returns the identifier used by the cluster API.
// Nexus Dashboard identifies a connected cluster by its cluster name. In
// Terraform, cluster_name is Optional+Computed because users can omit it during
// create and let the backend assign the name. Once readback has populated the
// Terraform id, prefer id; before that, fall back to cluster_name when it is
// already known, and use an empty identifier for create-time list lookup.
func setMultiClusterConnectivityID(model *MultiClusterConnectivityModel) string {
	// Id is unknown during the create
	if !model.Id.IsNull() && !model.Id.IsUnknown() && model.Id.ValueString() != "" {
		return model.Id.ValueString()
	}
	// ClusterName is unknown during the create
	if !model.ClusterName.IsNull() && !model.ClusterName.IsUnknown() && model.ClusterName.ValueString() != "" {
		return model.ClusterName.ValueString()
	}
	// return empty string during the create
	return ""
}

// RscCreateMultiClusterConnectivity creates a multi cluster connectivity nd resource
func (r *multiClusterConnectivityNdResource) rscCreateMultiClusterConnectivity(dg *diag.Diagnostics, model *MultiClusterConnectivityModel) {
	id := setMultiClusterConnectivityID(model)
	log.Printf("[INFO] Create nd_multi_cluster_connectivity id=%s", id)

	inData := model.GetModelData()
	inData.Spec.ClusterType = "ND"

	// Create multi cluster connectivity nd API client
	clusterAPI := api.NewClusterAPI(r.infraClient.ApiClient)

	clusterPayload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Creating Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not create multi cluster connectivity nd, Data Marshall error: %s", err.Error()),
		)
		return
	}

	// Call the API to create the multi cluster connectivity nd
	// During create, the backend ignores any provided clusterName and assigns it asynchronously.
	res, err := clusterAPI.Post(clusterPayload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		dg.AddError(
			"Error Creating Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not create multi cluster connectivity nd, unexpected error: %s %s", err.Error(), res.String()),
		)
		return
	}

	if r.rscGetMultiClusterConnectivity(dg, model) {
		dg.AddError(
			"Error Reading Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not read multi cluster connectivity nd with id %q after create: resource not found", id),
		)
	}
}

// GetMultiClusterConnectivity retrieves multi cluster connectivity nd information by name
func (r *multiClusterConnectivityNdResource) rscGetMultiClusterConnectivity(dg *diag.Diagnostics, model *MultiClusterConnectivityModel) bool {
	id := setMultiClusterConnectivityID(model)
	log.Printf("[INFO] Read nd_multi_cluster_connectivity id=%s", id)

	// Preserve sensitive fields that are not returned by the API.
	// Coerce unknown (planned but unset Optional+Computed) to null so the
	// provider does not return unknowns after apply.
	preservedUsername := model.Username
	preservedPassword := model.Password
	preservedLoginDomain := model.LoginDomain
	if preservedLoginDomain.IsUnknown() {
		preservedLoginDomain = types.StringNull()
	}
	preservedMultiClusterLoginDomain := model.MultiClusterLoginDomain
	if preservedMultiClusterLoginDomain.IsUnknown() {
		preservedMultiClusterLoginDomain = types.StringNull()
	}

	clusterAPI := api.NewClusterAPI(r.infraClient.ApiClient)
	clusterAPI.ClusterName = id
	respData, err := clusterAPI.Get()

	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return true
		}
		dg.AddError(
			"Error Reading Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not read multi cluster connectivity nd, unexpected error: %s %s", err.Error(), string(respData)),
		)
		return false
	}

	// On create, the backend does not use the clusterName from the request.
	// The cluster_name is still empty during the first read, so the GET call is
	// made without a cluster name. In that case, /api/v1/infra/clusters returns
	// all clusters, and we find the created ND cluster from that list.
	if clusterAPI.ClusterName == "" {
		var clustersResp map[string][]NDFCMultiClusterConnectivityModel
		err = json.Unmarshal(respData, &clustersResp)
		if err != nil {
			dg.AddError(
				"Error Reading Multi Cluster Connectivity ND",
				fmt.Sprintf("Could not unmarshal multi cluster connectivity nd response, unexpected error: %s", err.Error()),
			)
			return false
		}

		hostname := model.Hostname.ValueString()
		for _, cluster := range clustersResp["clusters"] {
			if cluster.Spec.Hostname == hostname && cluster.Spec.ClusterType == "ND" {
				dg.Append(model.SetModelData(&cluster)...)
				if dg.HasError() {
					return false
				}
				model.Id = model.ClusterName
				model.Username = preservedUsername
				model.Password = preservedPassword
				model.LoginDomain = preservedLoginDomain
				model.MultiClusterLoginDomain = preservedMultiClusterLoginDomain
				return false
			}
		}

		return true
	} else {
		var clusterResp NDFCMultiClusterConnectivityModel
		err = json.Unmarshal(respData, &clusterResp)
		if err != nil {
			dg.AddError(
				"Error Reading Multi Cluster Connectivity ND",
				fmt.Sprintf("Could not unmarshal multi cluster connectivity nd response, unexpected error: %s", err.Error()),
			)
			return false
		}

		dg.Append(model.SetModelData(&clusterResp)...)
		if dg.HasError() {
			return false
		}
		model.Id = model.ClusterName

		// Restore sensitive fields after SetModelData (API does not return them)
		model.Username = preservedUsername
		model.Password = preservedPassword
		model.LoginDomain = preservedLoginDomain
		model.MultiClusterLoginDomain = preservedMultiClusterLoginDomain
	}
	return false
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
func (r *multiClusterConnectivityNdResource) rscUpdateMultiClusterConnectivity(dg *diag.Diagnostics, model *MultiClusterConnectivityModel) {
	id := setMultiClusterConnectivityID(model)
	log.Printf("[INFO] Update nd_multi_cluster_connectivity id=%s", id)
	inData := model.GetModelData()

	clusterAPI := api.NewClusterAPI(r.infraClient.ApiClient)
	clusterAPI.ClusterName = id

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
			fmt.Sprintf("Could not update multi cluster connectivity nd, Data Marshall error: %s", err.Error()),
		)
		log.Printf("[ERROR] Error Updating Multi Cluster Connectivity ND: error=%s", err.Error())
		return
	}
	res, err := clusterAPI.Put(inDataBytes, &ndapi.APIOptions{DisablePayloadLog: true})

	if err != nil {
		dg.AddError(
			"Error Updating Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not update multi cluster connectivity nd, unexpected error: %s %s", err.Error(), res.String()),
		)
		log.Printf("[ERROR] Error Updating Multi Cluster Connectivity ND: error=%s", err.Error())
		return
	}
	// Read the updated multi cluster connectivity nd
	if r.rscGetMultiClusterConnectivity(dg, model) {
		dg.AddError(
			"Error Reading Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not read multi cluster connectivity nd with id %q after update: resource not found", id),
		)
	}
}

// DeleteMultiClusterConnectivity deletes a multi cluster connectivity nd by name
func (r *multiClusterConnectivityNdResource) rscDeleteMultiClusterConnectivity(dg *diag.Diagnostics, model *MultiClusterConnectivityModel) {
	id := setMultiClusterConnectivityID(model)
	log.Printf("[INFO] Delete nd_multi_cluster_connectivity id=%s", id)
	clusterAPI := api.NewClusterAPI(r.infraClient.ApiClient)
	clusterAPI.ClusterName = id
	clusterAPI.Delete = true

	payload := []byte("{}")

	res, err := clusterAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			log.Printf("[DEBUG] Multi Cluster Connectivity ND already absent during delete: id=%s", id)
			return
		}
		dg.AddError(
			"Error Deleting Multi Cluster Connectivity ND",
			fmt.Sprintf("Could not delete multi cluster connectivity nd, unexpected error: %s %s", err.Error(), res.String()),
		)
		log.Printf("[ERROR] Error Deleting Multi Cluster Connectivity ND: error=%s", err.Error())
		return
	}
}
