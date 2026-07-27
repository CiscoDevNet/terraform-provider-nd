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
	"terraform-provider-nd/internal/common/utils"
	"terraform-provider-nd/internal/infra/api"

	frameworktimeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

	createTimeout, pollInterval := backupCreateTimeouts(ctx, dg, input.Timeouts)
	if dg.HasError() {
		return
	}

	inData := input.GetModelData()

	// Create nd backup API client
	backupAPI := api.NewBackupAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)

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

	if !r.waitForBackupCreateCompletion(ctx, dg, backupAPI, input, createTimeout, pollInterval) {
		return
	}

	if r.rscGetBackup(ctx, dg, input) {
		dg.AddError(
			"Error Creating ND Backup",
			fmt.Sprintf("Could not read nd backup %q after create: resource not found", input.Id),
		)
	}
}

// waitForBackupCreateCompletion polls the backend until creation completes or
// the configured create deadline is reached.
func (r *backupNdResource) waitForBackupCreateCompletion(ctx context.Context, dg *diag.Diagnostics, backupAPI *api.BackupAPI, input *BackupModel, createTimeout time.Duration, pollInterval time.Duration) bool {
	id := input.Id.ValueString()
	backupAPI.Name = id

	err := utils.PollUntil(ctx, pollInterval, createTimeout, func(pollCtx context.Context) (bool, error) {
		return backupCreateStatusCheck(pollCtx, backupAPI, input)
	})

	if err == nil {
		return true
	}

	if errors.Is(err, utils.ErrPollTimeout) {
		dg.AddError(
			"Error Creating ND Backup",
			fmt.Sprintf("Timed out after %s waiting for backup %q creation to complete.", createTimeout, id),
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

// backupCreateTimeouts resolves the create deadline and status polling interval
// from the Terraform timeout object.
func backupCreateTimeouts(ctx context.Context, dg *diag.Diagnostics, configured TimeoutsValue) (time.Duration, time.Duration) {
	// The generated schema keeps timeouts as a custom TimeoutsValue, so adapt it
	// before using the HashiCorp timeout helper methods. For nd_backup,
	// timeouts.create is the overall create deadline and timeouts.read is the
	// status polling interval during create.
	if !validateBackupTimeoutValuesKnown(dg, configured) {
		return 0, 0
	}

	timeoutValue, timeoutStateDiags := newBackupFrameworkTimeouts(ctx, configured)
	dg.Append(timeoutStateDiags...)
	if dg.HasError() {
		return 0, 0
	}

	// Schema defaults make both values known. Zero is only the timeout module's
	// required fallback argument and is not used for a valid resource plan.
	createTimeout, createDiags := timeoutValue.Create(ctx, 0)
	dg.Append(createDiags...)

	pollInterval, readDiags := timeoutValue.Read(ctx, 0)
	dg.Append(readDiags...)
	if dg.HasError() {
		return 0, 0
	}

	if createTimeout <= 0 {
		dg.AddError(
			"Invalid Backup Create Timeout",
			"The nd_backup timeouts.create value must be greater than zero.",
		)
	}
	if pollInterval <= 0 {
		dg.AddError(
			"Invalid Backup Poll Interval",
			"The nd_backup timeouts.read value must be greater than zero.",
		)
	}

	return createTimeout, pollInterval
}

// validateBackupTimeoutValuesKnown prevents unknown or null timeout values from
// reaching the timeout module's zero fallback path.
func validateBackupTimeoutValuesKnown(dg *diag.Diagnostics, configured TimeoutsValue) bool {
	if configured.IsNull() || configured.IsUnknown() ||
		configured.Create.IsNull() || configured.Create.IsUnknown() ||
		configured.Read.IsNull() || configured.Read.IsUnknown() {
		dg.AddError(
			"Invalid Backup Timeouts",
			"The nd_backup timeouts.create and timeouts.read values must be known and non-null after schema defaults are applied.",
		)
		return false
	}

	return true
}

// newBackupFrameworkTimeouts adapts the generated timeout object to the value
// type expected by terraform-plugin-framework-timeouts.
func newBackupFrameworkTimeouts(ctx context.Context, configured TimeoutsValue) (frameworktimeouts.Value, diag.Diagnostics) {
	objectValue, diags := configured.ToObjectValue(ctx)
	return frameworktimeouts.Value{Object: objectValue}, diags
}

// backupCreateStatusCheck reads and classifies the current asynchronous backup
// creation status.
func backupCreateStatusCheck(ctx context.Context, backupAPI *api.BackupAPI, input *BackupModel) (bool, error) {
	id := input.Id.ValueString()
	respData, err := backupAPI.Get()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return false, fmt.Errorf("Nexus Dashboard reported backup creation failure for %q: GET /infra/backups/%s returned 404. Response: %s", id, id, string(respData))
		}

		if backupCreateStatusErrorRetryable(err) {
			log.Printf("[INFO] Waiting for nd_backup id=%s, transient read failure: %v", id, err)
			return false, nil
		}

		return false, fmt.Errorf("could not read nd backup %q creation status: %w. Response: %s", id, err, string(respData))
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

	switch strings.ToLower(backupResp.Status) {
	case "completed":
		return true, nil
	case "pending", "inprogress":
		return false, nil
	case "failed":
		return false, fmt.Errorf("Nexus Dashboard reported backup creation failure for %q: %s. Response: %s", id, detail, string(backupResp.Raw))
	default:
		return false, fmt.Errorf("Nexus Dashboard returned unexpected creation status %q for backup %q. Response: %s", backupResp.Status, id, string(backupResp.Raw))
	}
}

// backupCreateStatusErrorRetryable reports whether a status read failed with a
// transient HTTP status that can be retried until the create deadline.
func backupCreateStatusErrorRetryable(err error) bool {
	for _, statusCode := range []int{408, 429, 500, 502, 503, 504} {
		if strings.Contains(err.Error(), fmt.Sprintf("StatusCode %d", statusCode)) {
			return true
		}
	}

	return false
}

// backupCreateStatusMismatches verifies that the status response belongs to
// the backup requested by Terraform.
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

// backupCreateStatusDetail formats the useful status fields for logs and
// diagnostics while retaining the raw response as a fallback.
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
	// Preserve Terraform-only/API-omitted fields across response hydration.
	preservedDestination := in.Destination
	preservedEncryptionKey := in.EncryptionKey
	preservedTelemetryData := in.TelemetryData
	preservedTimeouts := in.Timeouts

	backupAPI := api.NewBackupAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
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

	// Restore Terraform-only/API-omitted fields after response hydration.
	in.EncryptionKey = preservedEncryptionKey
	in.TelemetryData = preservedTelemetryData
	in.Timeouts = preservedTimeouts
	in.Id = in.Name
	return false
}

// rscDeleteBackup deletes a nd backup by id
func (r *backupNdResource) rscDeleteBackup(ctx context.Context, dg *diag.Diagnostics, state *BackupModel) {
	id := state.Id.ValueString()
	log.Printf("[INFO] Delete nd_backup id=%s", id)
	backupAPI := api.NewBackupAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
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
