package action_backup_restore

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra"
	"terraform-provider-nd/internal/infra/api"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/action"
	actionschema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	restoreTypeConfigOnly = "configOnly"
	restoreTypeFull       = "full"

	importStatusPollInterval  = 5 * time.Second
	restoreStatusPollInterval = 30 * time.Second
	restoreStatusTimeout      = 60 * time.Minute
)

const (
	// idle means no backup or restore operation is currently active.
	backupStatusIdle = "idle"
	// downloading means ND is downloading/importing the backup archive from the configured source.
	backupStatusDownloading = "downloading"
	// ready means the backup has been imported and validated, and restore can be triggered.
	backupStatusReady = "ready"
	// processing means ND is actively running the backup or restore operation.
	backupStatusProcessing = "processing"
	// completed means the restore operation finished successfully.
	backupStatusCompleted = "completed"
	// failed means the backup or restore operation failed after it started.
	backupStatusFailed = "failed"
	// failedToValidate means the imported backup failed validation before restore could run.
	backupStatusFailedToValidate = "failedToValidate"
)

const backupOperationRestore = "restore"

type restoreStatusState int

const (
	restoreStatusUnknown restoreStatusState = iota
	restoreStatusInProgress
	restoreStatusSucceeded
	restoreStatusFailed
)

var (
	_ action.Action                   = &backupRestoreAction{}
	_ action.ActionWithConfigure      = &backupRestoreAction{}
	_ action.ActionWithValidateConfig = &backupRestoreAction{}
)

type backupRestoreAction struct {
	infraClient *infra.NexusDashboardInfra
}

type backupRestoreModel struct {
	EncryptionKey                   types.String `tfsdk:"encryption_key"`
	Name                            types.String `tfsdk:"name"`
	Path                            types.String `tfsdk:"path"`
	Source                          types.String `tfsdk:"source"`
	IgnorePersistentIPs             types.Bool   `tfsdk:"ignore_persistent_ips"`
	Type                            types.String `tfsdk:"type"`
	IncludeTelemetryOperationalData types.Bool   `tfsdk:"include_telemetry_operational_data"`
}

func NewBackupRestoreAction() action.Action {
	return &backupRestoreAction{}
}

func (a *backupRestoreAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_backup_restore"
}

func (a *backupRestoreAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = actionschema.Schema{
		Description:         "Restores Nexus Dashboard from an existing backup.",
		MarkdownDescription: "Restores Nexus Dashboard from an existing backup.",
		Attributes: map[string]actionschema.Attribute{
			"encryption_key": actionschema.StringAttribute{
				Required:            true,
				Description:         "The encryption key used to decrypt the backup during import.",
				MarkdownDescription: "The encryption key used to decrypt the backup during import.",
			},
			"name": actionschema.StringAttribute{
				Optional:            true,
				Description:         "The imported backup name. Defaults to an empty string when omitted.",
				MarkdownDescription: "The imported backup name. Defaults to an empty string when omitted.",
			},
			"path": actionschema.StringAttribute{
				Required:            true,
				Description:         "The path to the backup archive on the import source.",
				MarkdownDescription: "The path to the backup archive on the import source.",
			},
			"source": actionschema.StringAttribute{
				Required:            true,
				Description:         "The remote storage location source for the backup import.",
				MarkdownDescription: "The remote storage location source for the backup import.",
			},
			"ignore_persistent_ips": actionschema.BoolAttribute{
				Optional:            true,
				Description:         "Whether to ignore persistent IPs during restore. Defaults to false when omitted.",
				MarkdownDescription: "Whether to ignore persistent IPs during restore. Defaults to `false` when omitted.",
			},
			"type": actionschema.StringAttribute{
				Optional:            true,
				Description:         "The restore type. Defaults to configOnly when omitted.",
				MarkdownDescription: "The restore type. Defaults to `configOnly` when omitted.",
				Validators: []validator.String{
					stringvalidator.OneOf(restoreTypeConfigOnly, restoreTypeFull),
				},
			},
			"include_telemetry_operational_data": actionschema.BoolAttribute{
				Optional:            true,
				Description:         "Whether to restore telemetry operational data. This is valid only when type is full. Defaults to false when omitted.",
				MarkdownDescription: "Whether to restore telemetry operational data. This is valid only when `type` is `full`. Defaults to `false` when omitted.",
			},
		},
	}
}

func (a *backupRestoreAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Action Configure Type",
			fmt.Sprintf("Expected registry.ClientProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	infraModule := client.GetModule(infra.ModuleKey)
	if infraModule == nil {
		resp.Diagnostics.AddError(
			"Infra Module Not Found",
			"The infra module was not registered with the provider.",
		)
		return
	}

	infraClient, ok := infraModule.(*infra.NexusDashboardInfra)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Infra Module Type",
			fmt.Sprintf("Expected *infra.NexusDashboardInfra, got: %T. Please report this issue to the provider developers.", infraModule),
		)
		return
	}

	a.infraClient = infraClient
}

func (a *backupRestoreAction) ValidateConfig(ctx context.Context, req action.ValidateConfigRequest, resp *action.ValidateConfigResponse) {
	var config backupRestoreModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	payload := config.restorePayload()
	if payload.IncludeTelemetryOperationalData && payload.Type != restoreTypeFull {
		resp.Diagnostics.AddAttributeError(
			path.Root("include_telemetry_operational_data"),
			"Invalid Backup Restore Configuration",
			"include_telemetry_operational_data can be true only when type is full.",
		)
	}
}

func (a *backupRestoreAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	log.Printf("[DEBUG] Start action: nd_backup_restore")

	if a.infraClient == nil {
		resp.Diagnostics.AddError(
			"Nexus Dashboard Infra Client Not Configured",
			"The provider did not configure the Nexus Dashboard infra client before invoking nd_backup_restore.",
		)
		return
	}

	var config backupRestoreModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	payloadData := config.restorePayload()
	if payloadData.IncludeTelemetryOperationalData && payloadData.Type != restoreTypeFull {
		resp.Diagnostics.AddAttributeError(
			path.Root("include_telemetry_operational_data"),
			"Invalid Backup Restore Configuration",
			"include_telemetry_operational_data can be true only when type is full.",
		)
		return
	}

	importPayload, err := json.Marshal(config.importPayload())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Importing ND Backup",
			fmt.Sprintf("Could not marshal backup import payload: %v", err),
		)
		return
	}

	restorePayload, err := json.Marshal(payloadData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Restoring ND Backup",
			fmt.Sprintf("Could not marshal backup restore payload: %v", err),
		)
		return
	}

	backupAPI := api.NewBackupAPI(a.infraClient.ApiClient)
	guard := ndapi.Acquire(backupAPI.FabricScope(), backupAPI.RscName(), ndapi.LockGlobal)
	defer guard.Release()

	res, err := backupAPI.ClearImportedBackup()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Clearing Imported ND Backup",
			fmt.Sprintf("Could not execute DELETE %s, unexpected error: %v %v", api.UrlBackupImport, err, res),
		)
		return
	}

	res, err = backupAPI.ImportBackup(importPayload)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Importing ND Backup",
			fmt.Sprintf("Could not execute POST %s, unexpected error: %v %v", api.UrlBackupImport, err, res),
		)
		return
	}

	a.sendProgress(resp, "Nexus Dashboard backup import request accepted. Waiting for imported backup validation.")
	if !a.waitForImportedBackupReady(ctx, backupAPI, resp, importStatusPollInterval) {
		return
	}

	res, err = backupAPI.Restore(restorePayload)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Restoring ND Backup",
			fmt.Sprintf("Could not execute POST %s, unexpected error: %v %v", api.UrlBackupRestore, err, res),
		)
		return
	}

	a.sendProgress(resp, "Nexus Dashboard backup restore request accepted. Waiting for restore status.")

	if !a.waitForRestoreCompletion(ctx, backupAPI, resp, restoreStatusPollInterval) {
		return
	}

	log.Printf("[DEBUG] End action: nd_backup_restore")
}

func (a *backupRestoreAction) waitForImportedBackupReady(ctx context.Context, backupAPI *api.BackupAPI, resp *action.InvokeResponse, pollInterval time.Duration) bool {
	deadline := time.NewTimer(restoreStatusTimeout)
	defer deadline.Stop()

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		status, err := backupAPI.Status()
		if err != nil {
			a.sendProgress(resp, fmt.Sprintf("Waiting for Nexus Dashboard backup import status. Status check failed: %v", err))
		} else {
			state, detail := evaluateRestoreStatus(status)
			a.sendProgress(resp, fmt.Sprintf("Nexus Dashboard backup import status: %s", detail))

			if strings.EqualFold(status.Operation, backupOperationRestore) &&
				strings.EqualFold(status.State, backupStatusReady) {
				return true
			}

			if state == restoreStatusFailed {
				resp.Diagnostics.AddError(
					"Error Importing ND Backup",
					fmt.Sprintf("Nexus Dashboard reported backup import validation failure: %s", detail),
				)
				return false
			}
		}

		select {
		case <-ctx.Done():
			resp.Diagnostics.AddError(
				"Error Importing ND Backup",
				fmt.Sprintf("Context canceled while waiting for backup import validation: %v", ctx.Err()),
			)
			return false
		case <-deadline.C:
			resp.Diagnostics.AddError(
				"Error Importing ND Backup",
				fmt.Sprintf("Timed out after %s waiting for backup import validation.", restoreStatusTimeout),
			)
			return false
		case <-ticker.C:
		}
	}
}

func (a *backupRestoreAction) waitForRestoreCompletion(ctx context.Context, backupAPI *api.BackupAPI, resp *action.InvokeResponse, pollInterval time.Duration) bool {
	deadline := time.NewTimer(restoreStatusTimeout)
	defer deadline.Stop()

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		status, err := backupAPI.Status()
		if err != nil {
			a.sendProgress(resp, fmt.Sprintf("Waiting for Nexus Dashboard backup restore status. Status check failed: %v", err))
		} else {
			state, detail := evaluateRestoreStatus(status)
			a.sendProgress(resp, fmt.Sprintf("Nexus Dashboard backup restore status: %s", detail))

			switch state {
			case restoreStatusSucceeded:
				return true
			case restoreStatusFailed:
				resp.Diagnostics.AddError(
					"Error Restoring ND Backup",
					fmt.Sprintf("Nexus Dashboard reported backup restore failure: %s", detail),
				)
				return false
			}
		}

		select {
		case <-ctx.Done():
			resp.Diagnostics.AddError(
				"Error Restoring ND Backup",
				fmt.Sprintf("Context canceled while waiting for backup restore completion: %v", ctx.Err()),
			)
			return false
		case <-deadline.C:
			resp.Diagnostics.AddError(
				"Error Restoring ND Backup",
				fmt.Sprintf("Timed out after %s waiting for backup restore completion.", restoreStatusTimeout),
			)
			return false
		case <-ticker.C:
		}
	}
}

func (a *backupRestoreAction) sendProgress(resp *action.InvokeResponse, message string) {
	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: message})
	}
}

func evaluateRestoreStatus(status api.BackupStatus) (restoreStatusState, string) {
	if len(status.Raw) == 0 {
		return restoreStatusUnknown, "status response is empty"
	}

	operation := strings.ToLower(status.Operation)
	state := strings.ToLower(status.State)
	detail := restoreStatusDetail(status)

	if operation != "" && operation != backupOperationRestore {
		return restoreStatusUnknown, detail
	}

	switch state {
	case backupStatusCompleted:
		return restoreStatusSucceeded, detail
	case backupStatusFailed, strings.ToLower(backupStatusFailedToValidate):
		return restoreStatusFailed, detail
	case backupStatusIdle, backupStatusDownloading, backupStatusReady, backupStatusProcessing:
		return restoreStatusInProgress, detail
	default:
		return restoreStatusUnknown, detail
	}
}

func restoreStatusDetail(status api.BackupStatus) string {
	var parts []string
	if status.Id != "" {
		parts = append(parts, fmt.Sprintf("id=%s", status.Id))
	}
	if status.Operation != "" {
		parts = append(parts, fmt.Sprintf("operation=%s", status.Operation))
	}
	if status.State != "" {
		parts = append(parts, fmt.Sprintf("state=%s", status.State))
	}
	if status.Error != "" {
		parts = append(parts, fmt.Sprintf("error=%s", status.Error))
	}
	if status.Details.Progress != nil {
		parts = append(parts, fmt.Sprintf("progress=%d%%", *status.Details.Progress))
	}

	if len(parts) == 0 {
		return string(status.Raw)
	}

	return strings.Join(parts, ", ")
}

func (m backupRestoreModel) importPayload() api.BackupImportPayload {
	return api.BackupImportPayload{
		EncryptionKey: stringValueOrEmpty(m.EncryptionKey),
		Name:          stringValueOrEmpty(m.Name),
		Path:          stringValueOrEmpty(m.Path),
		Source:        stringValueOrEmpty(m.Source),
	}
}

func (m backupRestoreModel) restorePayload() api.BackupRestorePayload {
	restoreType := restoreTypeConfigOnly
	if !m.Type.IsNull() && !m.Type.IsUnknown() && m.Type.ValueString() != "" {
		restoreType = m.Type.ValueString()
	}

	return api.BackupRestorePayload{
		IgnorePersistentIPs:             boolValueOrFalse(m.IgnorePersistentIPs),
		Type:                            restoreType,
		IncludeTelemetryOperationalData: boolValueOrFalse(m.IncludeTelemetryOperationalData),
	}
}

func boolValueOrFalse(v types.Bool) bool {
	if v.IsNull() || v.IsUnknown() {
		return false
	}

	return v.ValueBool()
}

func stringValueOrEmpty(v types.String) string {
	if v.IsNull() || v.IsUnknown() {
		return ""
	}

	return v.ValueString()
}
