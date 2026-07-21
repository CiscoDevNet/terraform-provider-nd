package resource_backup

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	// backupCreateStatusPollInterval controls how often the provider checks
	// ND backup creation status while waiting for status=completed.
	backupCreateStatusPollInterval = 30 * time.Second

	// backupCreateDefaultTimeout is the default maximum duration Terraform waits
	// for ND backup creation to complete.
	backupCreateDefaultTimeout = 90 * time.Minute
)

type backupCreateStatusResponse struct {
	Raw         json.RawMessage `json:"-"`
	Destination string          `json:"destination"`
	Name        string          `json:"name"`
	Status      string          `json:"status"`
	Type        string          `json:"type"`
}

// rscCreateBackup creates a nd backup resource
func (r *backupNdResource) rscCreateBackup(ctx context.Context, dg *diag.Diagnostics, input *BackupModel) {
	input.Id = input.Name
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
	res, err := backupAPI.Post(backupPayload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		dg.AddError(
			"Error Creating ND Backup",
			fmt.Sprintf("Could not create nd backup, unexpected error: %v %v", err, res),
		)
		return
	}

	if !r.waitForBackupCreateCompletion(ctx, dg, backupAPI, input) {
		return
	}

	if r.rscGetBackup(ctx, dg, input) {
		dg.AddError(
			"Error Creating ND Backup",
			fmt.Sprintf("Could not read nd backup %q after create: resource not found", input.Id),
		)
	}
}

func (r *backupNdResource) waitForBackupCreateCompletion(ctx context.Context, dg *diag.Diagnostics, backupAPI *api.BackupAPI, input *BackupModel) bool {
	id := input.Id.ValueString()
	backupAPI.Name = id

	err := ndapi.PollUntil(ctx, backupCreateStatusPollInterval, backupCreateDefaultTimeout, func() (bool, error) {
		return backupCreateStatusCheck(backupAPI, input)
	})

	if err == nil {
		return true
	}

	if errors.Is(err, ndapi.ErrPollTimeout) {
		dg.AddError(
			"Error Creating ND Backup",
			fmt.Sprintf("Timed out after %s waiting for backup %q creation to complete.", backupCreateDefaultTimeout, id),
		)
		return false
	}

	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		dg.AddError(
			"Error Creating ND Backup",
			fmt.Sprintf("Context canceled while waiting for backup %q creation to complete: %v", id, err),
		)
		return false
	}

	dg.AddError("Error Creating ND Backup", err.Error())
	return false
}

func backupCreateStatusCheck(backupAPI *api.BackupAPI, input *BackupModel) (bool, error) {
	id := input.Id.ValueString()
	respData, err := backupAPI.Get()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return false, fmt.Errorf("Nexus Dashboard reported backup creation failure for %q: GET /infra/backups/%s returned 404. Response: %s", id, id, string(respData))
		}

		log.Printf("[INFO] Waiting for nd_backup id=%s, read failed: %v", id, err)
		return false, nil
	}

	var backupResp backupCreateStatusResponse
	backupResp.Raw = append(json.RawMessage(nil), respData...)
	if err := json.Unmarshal(respData, &backupResp); err != nil {
		return false, fmt.Errorf("Could not unmarshal nd backup status response, unexpected error: %v", err)
	}

	detail := backupCreateStatusDetail(backupResp)
	log.Printf("[INFO] nd_backup id=%s create status: %s", id, detail)

	if mismatches := backupCreateStatusMismatches(input, backupResp); len(mismatches) > 0 {
		return false, fmt.Errorf("Nexus Dashboard returned backup %q, but the response did not match the requested configuration: %s. Response: %s", id, strings.Join(mismatches, ", "), detail)
	}

	return strings.EqualFold(backupResp.Status, "completed"), nil
}

func backupCreateStatusMismatches(input *BackupModel, status backupCreateStatusResponse) []string {
	var mismatches []string

	expectedName := input.Name.ValueString()
	if status.Name != expectedName {
		mismatches = append(mismatches, fmt.Sprintf("name expected %q got %q", expectedName, status.Name))
	}

	expectedType := input.Type.ValueString()
	if status.Type != expectedType {
		mismatches = append(mismatches, fmt.Sprintf("type expected %q got %q", expectedType, status.Type))
	}

	expectedDestination := ""
	if !input.Destination.IsNull() && !input.Destination.IsUnknown() {
		expectedDestination = input.Destination.ValueString()
	}
	if status.Destination != expectedDestination {
		mismatches = append(mismatches, fmt.Sprintf("destination expected %q got %q", expectedDestination, status.Destination))
	}

	return mismatches
}

func backupCreateStatusDetail(status backupCreateStatusResponse) string {
	var parts []string
	if status.Name != "" {
		parts = append(parts, fmt.Sprintf("name=%s", status.Name))
	}
	if status.Status != "" {
		parts = append(parts, fmt.Sprintf("status=%s", status.Status))
	}
	if status.Type != "" {
		parts = append(parts, fmt.Sprintf("type=%s", status.Type))
	}
	if status.Destination != "" {
		parts = append(parts, fmt.Sprintf("destination=%q", status.Destination))
	}

	if len(parts) == 0 {
		return string(status.Raw)
	}

	return strings.Join(parts, ", ")
}

// rscGetBackup retrieves nd backup information by id
func (r *backupNdResource) rscGetBackup(ctx context.Context, dg *diag.Diagnostics, in *BackupModel) bool {
	id := in.Id.ValueString()
	log.Printf("[INFO] Read nd_backup id=%s", id)
	// Preserve sensitive fields that are not returned by the API
	preservedDestination := in.Destination
	preservedEncryptionKey := in.EncryptionKey
	preservedTelemetryData := in.TelemetryData

	backupAPI := api.NewBackupAPI(r.infraClient.ApiClient)
	backupAPI.Name = id
	respData, err := backupAPI.Get()

	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return true
		}
		dg.AddError(
			"Error Reading ND Backup",
			fmt.Sprintf("Could not read nd backup, unexpected error: %v %s", err, string(respData)),
		)
		return false
	}

	var backupResp NDFCBackupModel
	err = json.Unmarshal(respData, &backupResp)
	if err != nil {
		dg.AddError(
			"Error Reading ND Backup",
			fmt.Sprintf("Could not unmarshal nd backup response, unexpected error: %v", err),
		)
		return false
	}

	in.SetModelData(&backupResp)

	if backupResp.Destination == "" {
		in.Destination = preservedDestination
	} else {
		in.Destination = types.StringValue(backupResp.Destination)
	}

	// Restore sensitive fields after SetModelData (API does not return them)
	in.EncryptionKey = preservedEncryptionKey
	in.TelemetryData = preservedTelemetryData
	in.Id = in.Name
	return false
}

// rscDeleteBackup deletes a nd backup by id
func (r *backupNdResource) rscDeleteBackup(ctx context.Context, dg *diag.Diagnostics, state *BackupModel) {
	id := state.Id.ValueString()
	log.Printf("[INFO] Delete nd_backup id=%s", id)
	backupAPI := api.NewBackupAPI(r.infraClient.ApiClient)
	backupAPI.Name = id

	res, err := backupAPI.Delete(nil)
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return
		}
		dg.AddError(
			"Error Deleting ND Backup",
			fmt.Sprintf("Could not delete nd backup, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Deleting ND Backup: error=%s", err.Error())
		return
	}
}
