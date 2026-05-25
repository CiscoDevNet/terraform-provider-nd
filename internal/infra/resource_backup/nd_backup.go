package resource_backup

import (
	"context"
	"encoding/json"
	"fmt"
	"terraform-provider-nd/internal/infra/api"

	"log"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// rscCreateBackup creates a nd backup resource
func (r *backupNdResource) rscCreateBackup(ctx context.Context, dg *diag.Diagnostics, input *BackupModel) {
	if input == nil {
		dg.AddError(
			"Invalid Input",
			"The input model is nil",
		)
		return
	}

	log.Printf("[INFO] Create nd_backup name=%s", input.Name.ValueString())
	inData := input.GetModelData()

	// Create nd backup API client
	backupAPI := api.NewBackupAPI(r.infraClient.ApiClient)

	// Convert model data to JSON
	backupPayload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Creating ND Backup",
			fmt.Sprintf("Could not create nd backup, Data Marshall error: %v", err),
		)
		return
	}

	// Call the API to create the nd backup
	res, err := backupAPI.Post(backupPayload, nil)
	if err != nil {
		dg.AddError(
			"Error Creating ND Backup",
			fmt.Sprintf("Could not create nd backup, unexpected error: %v %v", err, res),
		)
		return
	}
	r.rscGetBackup(ctx, dg, input)
}

// rscGetBackup retrieves nd backup information by name
func (r *backupNdResource) rscGetBackup(ctx context.Context, dg *diag.Diagnostics, in *BackupModel) {
	log.Printf("[INFO] Read nd_backup name=%s", in.Name.ValueString())
	// Preserve sensitive fields that are not returned by the API
	preservedEncryptionKey := in.EncryptionKey
	telemetryData := in.TelemetryData

	backupAPI := api.NewBackupAPI(r.infraClient.ApiClient)
	backupAPI.Name = in.Name.ValueString()
	respData, err := backupAPI.Get()

	if err != nil {
		dg.AddError(
			"Error Reading ND Backup",
			fmt.Sprintf("Could not read nd backup, unexpected error: %v %v", err, respData),
		)
		return
	}

	var backupResp NDFCBackupModel
	err = json.Unmarshal(respData, &backupResp)
	if err != nil {
		dg.AddError(
			"Error Reading ND Backup",
			fmt.Sprintf("Could not unmarshal nd backup response, unexpected error: %v", err),
		)
		return
	}

	in.SetModelData(&backupResp)

	// Destination = "" is valid and indicates the backup location is ND local storage, so it should be set in the state even if empty.
	in.Destination = types.StringValue(backupResp.Destination)

	// Restore sensitive fields after SetModelData (API does not return them)
	in.EncryptionKey = preservedEncryptionKey
	in.TelemetryData = telemetryData
}

// rscDeleteBackup deletes a nd backup by name
func (r *backupNdResource) rscDeleteBackup(ctx context.Context, dg *diag.Diagnostics, state *BackupModel) {
	log.Printf("[INFO] Delete nd_backup name=%s", state.Name.ValueString())
	backupAPI := api.NewBackupAPI(r.infraClient.ApiClient)
	backupAPI.Name = state.Name.ValueString()

	res, err := backupAPI.Delete(nil)
	if err != nil {
		dg.AddError(
			"Error Deleting ND Backup",
			fmt.Sprintf("Could not delete nd backup, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Deleting ND Backup: error=%s", err.Error())
		return
	}
}
