package resource_fabric_aci

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type aciClusterListResponse struct {
	Clusters []aciClusterResponse `json:"clusters,omitempty"`
}

type aciClusterResponse struct {
	Spec aciClusterResponseSpec `json:"spec,omitempty"`
}

type aciClusterResponseSpec struct {
	NDFCSpecValue
	FabricName string `json:"name,omitempty"`
}

type aciClusterRemovePayload struct {
	Credentials NDFCSpecCredentialsValue `json:"credentials"`
	Force       bool                     `json:"force,omitempty"`
}

func (r aciClusterResponse) modelData() NDFCFabricAciModel {
	data := NDFCFabricAciModel{
		Spec: r.Spec.NDFCSpecValue,
	}
	if data.Spec.Aci.FabricName == "" {
		data.Spec.Aci.FabricName = r.Spec.FabricName
	}
	return data
}

// rscCreateFabricAci creates a fabric ACI resource.
func (r *fabricAciResource) rscCreateFabricAci(ctx context.Context, dg *diag.Diagnostics, input *FabricAciModel) {
	if input == nil {
		dg.AddError(
			"Invalid Input",
			"The input model is nil",
		)
		return
	}

	inData := input.GetModelData()
	inData.Spec.ClusterType = "APIC"

	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient)

	clusterPayload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Creating Fabric ACI",
			fmt.Sprintf("Could not create fabric ACI, Data Marshall error: %v", err),
		)
		return
	}

	res, err := clusterAPI.Post(clusterPayload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		dg.AddError(
			"Error Creating Fabric ACI",
			fmt.Sprintf("Could not create fabric ACI, unexpected error: %v %v", err, res),
		)
		return
	}

	if r.rscGetFabricAci(ctx, dg, input) {
		dg.AddError(
			"Error Reading Fabric ACI",
			"Could not read fabric ACI after create because it was not found.",
		)
	}
}

// rscGetFabricAci retrieves fabric ACI information by name.
func (r *fabricAciResource) rscGetFabricAci(ctx context.Context, dg *diag.Diagnostics, in *FabricAciModel) bool {
	preservedUsername := in.Username
	preservedPassword := in.Password
	preservedLoginDomain := in.LoginDomain
	if preservedLoginDomain.IsUnknown() {
		preservedLoginDomain = types.StringNull()
	}

	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient)
	clusterAPI.ClusterName = in.FabricName.ValueString()
	respData, err := clusterAPI.Get()

	if err != nil {
		if isNotFoundError(err) {
			return true
		}
		dg.AddError(
			"Error Reading Fabric ACI",
			fmt.Sprintf("Could not read fabric ACI, unexpected error: %v %v", err, respData),
		)
		return false
	}

	if clusterAPI.ClusterName == "" {
		var clustersResp aciClusterListResponse
		err = json.Unmarshal(respData, &clustersResp)
		if err != nil {
			dg.AddError(
				"Error Reading Fabric ACI",
				fmt.Sprintf("Could not unmarshal fabric ACI response, unexpected error: %v", err),
			)
			return false
		}

		hostname := in.Hostname.ValueString()
		for _, cluster := range clustersResp.Clusters {
			modelData := cluster.modelData()
			if modelData.Spec.Hostname == hostname && modelData.Spec.ClusterType == "APIC" {
				in.SetModelData(&modelData)
				in.NormalizeTelemetryNetworkState()
				in.Username = preservedUsername
				in.Password = preservedPassword
				in.LoginDomain = preservedLoginDomain
				return false
			}
		}

		dg.AddError(
			"Error Reading Fabric ACI",
			fmt.Sprintf("Could not find cluster with onboardUrl %q and clusterType APIC in the response", hostname),
		)
		return false
	}

	var clusterResp aciClusterResponse
	err = json.Unmarshal(respData, &clusterResp)
	if err != nil {
		dg.AddError(
			"Error Reading Fabric ACI",
			fmt.Sprintf("Could not unmarshal fabric ACI response, unexpected error: %v", err),
		)
		return false
	}

	modelData := clusterResp.modelData()
	in.SetModelData(&modelData)
	in.NormalizeTelemetryNetworkState()

	in.Username = preservedUsername
	in.Password = preservedPassword
	in.LoginDomain = preservedLoginDomain
	return false
}

// rscUpdateFabricAci updates a fabric ACI resource.
func (r *fabricAciResource) rscUpdateFabricAci(ctx context.Context, dg *diag.Diagnostics, clusterModel *FabricAciModel) {
	if clusterModel == nil {
		dg.AddError(
			"Invalid Input",
			"The input model is nil",
		)
		return
	}

	if clusterModel.FabricName.IsNull() || clusterModel.FabricName.IsUnknown() || clusterModel.FabricName.ValueString() == "" {
		dg.AddError(
			"Invalid Input",
			"The fabric_name value is required to update fabric ACI.",
		)
		return
	}

	inData := clusterModel.GetManageModelData()

	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient)
	clusterAPI.ClusterName = clusterModel.FabricName.ValueString()

	inDataBytes, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Updating Fabric ACI",
			fmt.Sprintf("Could not update fabric ACI, Data Marshall error: %v", err),
		)
		log.Printf("[ERROR] Error Updating Fabric ACI: error=%s", err.Error())
		return
	}

	res, err := clusterAPI.Put(inDataBytes, nil)
	if err != nil {
		dg.AddError(
			"Error Updating Fabric ACI",
			fmt.Sprintf("Could not update fabric ACI, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Updating Fabric ACI: error=%s", err.Error())
		return
	}

	if r.rscGetFabricAci(ctx, dg, clusterModel) {
		dg.AddError(
			"Error Reading Fabric ACI",
			"Could not read fabric ACI after update because it was not found.",
		)
	}
}

// rscDeleteFabricAci deletes a fabric ACI resource by name.
func (r *fabricAciResource) rscDeleteFabricAci(_ context.Context, dg *diag.Diagnostics, state *FabricAciModel) {

	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient)
	clusterAPI.ClusterName = state.FabricName.ValueString()
	clusterAPI.Delete = true

	removePayload := aciClusterRemovePayload{
		Credentials: NDFCSpecCredentialsValue{
			Username: state.Username.ValueString(),
			Password: state.Password.ValueString(),
		},
	}

	if !state.LoginDomain.IsNull() && !state.LoginDomain.IsUnknown() {
		removePayload.Credentials.LoginDomain = state.LoginDomain.ValueString()
	}

	payload, err := json.Marshal(removePayload)
	if err != nil {
		dg.AddError(
			"Error Deleting Fabric ACI",
			fmt.Sprintf("Could not delete fabric ACI, Data Marshall error: %v", err),
		)
		return
	}

	res, err := clusterAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		dg.AddError(
			"Error Deleting Fabric ACI",
			fmt.Sprintf("Could not delete fabric ACI, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Deleting Fabric ACI: error=%s", err.Error())
		return
	}
}

func isNotFoundError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "StatusCode 404")
}
