package resource_fabric_vxlan

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"terraform-provider-nd/internal/manage/api"

	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RscCreateFabric creates a fabric resource using the VXLAN fabric model
func (r *fabricVxlanResource) rscCreateFabric(ctx context.Context, dg *diag.Diagnostics, input *FabricVxlanModel) {
	if input == nil {
		dg.AddError(
			"Invalid Input",
			"The input model is nil",
		)
		return
	}

	inData := input.GetModelData()

	log.Printf("Creating fabric %s with category %s", inData.FabricName, inData.Category)

	// Create fabric API client
	fabricAPI := api.NewFabricAPI(nil, r.manageClient.ApiClient)

	// Convert model data to JSON
	fabricPayload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Creating Fabric",
			fmt.Sprintf("Could not create fabric, Data Marshall error: %v", err),
		)
		return
	}

	// Call the API to create the fabric
	res, err := fabricAPI.Post(fabricPayload)
	if err != nil {
		dg.AddError(
			"Error Creating Fabric",
			fmt.Sprintf("Could not create fabric, unexpected error: %v %v", err, res),
		)
		return
	}
	// Read the created fabric
	// ND BUG - license_tier is not set in the response. Try delaying the read
	time.Sleep(2 * time.Second)
	r.rscGetFabric(ctx, dg, input)

}

// GetFabric retrieves fabric information by name
func (r *fabricVxlanResource) rscGetFabric(ctx context.Context, dg *diag.Diagnostics, in *FabricVxlanModel) {

	fabricAPI := api.NewFabricAPI(nil, r.manageClient.ApiClient)
	fabricAPI.FabricName = in.FabricName.ValueString()
	respData, err := fabricAPI.Get()
	if err != nil {
		dg.AddError(
			"Error Creating Fabric",
			fmt.Sprintf("Could not read fabric, unexpected error: %v %v", err, respData),
		)
		return
	}
	var outData NDFCFabricVxlanModel

	err = json.Unmarshal(respData, &outData)
	if err != nil {
		dg.AddError(
			"Error Creating Fabric",
			fmt.Sprintf("Could not read fabric, unexpected error: %v %v", err, respData),
		)
		return
	}
	log.Printf("Location = %v %v", *outData.Location.Latitude, *outData.Location.Longitude)
	log.Printf("Netflow = %v", *outData.Management.NetflowSettings.NetflowEnable)
	in.SetModelData(&outData)
	log.Printf("Location from model=%v,%v", in.Location.Latitude.ValueFloat64(), in.Location.Longitude.ValueFloat64())
}

// UpdateFabricEVPN updates a fabric with the provided payload
func (r *fabricVxlanResource) rscUpdateFabric(ctx context.Context, dg *diag.Diagnostics, fabricModel *FabricVxlanModel) {
	inData := fabricModel.GetModelData()
	log.Printf("Creating fabric %s with category %s", inData.FabricName, inData.Category)

	fabricAPI := api.NewFabricAPI(nil, r.manageClient.ApiClient)
	fabricAPI.FabricName = inData.FabricName

	inDataBytes, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Updating Fabric",
			fmt.Sprintf("Could not update fabric, Data Marshall error: %v", err),
		)
		tflog.Error(ctx, "Error Updating Fabric", map[string]interface{}{"error": err.Error()})
		return
	}
	res, err := fabricAPI.Put(inDataBytes)
	if err != nil {
		dg.AddError(
			"Error Updating Fabric",
			fmt.Sprintf("Could not update fabric, unexpected error: %v %v", err, res),
		)
		tflog.Error(ctx, "Error Updating Fabric", map[string]interface{}{"error": err.Error()})
		return
	}
	// Read the updated fabric
	r.rscGetFabric(ctx, dg, fabricModel)
	log.Printf("Updated fabric %s with category %s", inData.FabricName, inData.Category)

}

// DeleteFabricEVPN deletes a fabric by name
func (r *fabricVxlanResource) rscDeleteFabric(ctx context.Context, dg *diag.Diagnostics, fabricName string) {
	fabricAPI := api.NewFabricAPI(nil, r.manageClient.ApiClient)
	fabricAPI.FabricName = fabricName
	res, err := fabricAPI.Delete()
	if err != nil {
		dg.AddError(
			"Error Deleting Fabric",
			fmt.Sprintf("Could not delete fabric, unexpected error: %v %v", err, res),
		)
		tflog.Error(ctx, "Error Deleting Fabric", map[string]interface{}{"error": err.Error()})
		return
	}
}
