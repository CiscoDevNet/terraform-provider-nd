package action_backup_restore

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"terraform-provider-nd/internal/infra/api"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type fakeClientProvider struct {
	modules map[string]interface{}
}

func (f fakeClientProvider) GetModule(name string) interface{} {
	return f.modules[name]
}

func TestBackupRestoreActionMetadata(t *testing.T) {
	t.Parallel()

	a := NewBackupRestoreAction()
	resp := &action.MetadataResponse{}

	a.Metadata(context.Background(), action.MetadataRequest{ProviderTypeName: "nd"}, resp)

	if resp.TypeName != "nd_backup_restore" {
		t.Fatalf("expected action type nd_backup_restore, got %q", resp.TypeName)
	}
}

func TestBackupRestoreModelRestorePayloadDefaults(t *testing.T) {
	t.Parallel()

	payload := backupRestoreModel{}.restorePayload()

	if payload.IgnorePersistentIPs {
		t.Fatalf("expected ignorePersistentIPs default false")
	}

	if payload.Type != restoreTypeConfigOnly {
		t.Fatalf("expected type default %q, got %q", restoreTypeConfigOnly, payload.Type)
	}

	if payload.IncludeTelemetryOperationalData {
		t.Fatalf("expected includeTelemetryOperationalData default false")
	}
}

func TestBackupRestoreModelRestorePayloadConfigured(t *testing.T) {
	t.Parallel()

	payload := backupRestoreModel{
		IgnorePersistentIPs:             types.BoolValue(true),
		Type:                            types.StringValue(restoreTypeFull),
		IncludeTelemetryOperationalData: types.BoolValue(true),
	}.restorePayload()

	if !payload.IgnorePersistentIPs {
		t.Fatalf("expected ignorePersistentIPs true")
	}

	if payload.Type != restoreTypeFull {
		t.Fatalf("expected type %q, got %q", restoreTypeFull, payload.Type)
	}

	if !payload.IncludeTelemetryOperationalData {
		t.Fatalf("expected includeTelemetryOperationalData true")
	}
}

func TestBackupRestoreModelImportPayloadDefaults(t *testing.T) {
	t.Parallel()

	payload := backupRestoreModel{}.importPayload()

	if payload.EncryptionKey != "" {
		t.Fatalf("expected encryptionKey default empty, got %q", payload.EncryptionKey)
	}

	if payload.Name != "" {
		t.Fatalf("expected name default empty, got %q", payload.Name)
	}

	if payload.Path != "" {
		t.Fatalf("expected path default empty, got %q", payload.Path)
	}

	if payload.Source != "" {
		t.Fatalf("expected source default empty, got %q", payload.Source)
	}
}

func TestBackupRestoreModelImportPayloadConfigured(t *testing.T) {
	t.Parallel()

	payload := backupRestoreModel{
		EncryptionKey: types.StringValue("mykey-12345"),
		Name:          types.StringValue(""),
		Path:          types.StringValue("mybackups/mybackup.tar.gz"),
		Source:        types.StringValue("sftp-server"),
	}.importPayload()

	if payload.EncryptionKey != "mykey-12345" {
		t.Fatalf("expected encryptionKey mykey-12345, got %q", payload.EncryptionKey)
	}

	if payload.Name != "" {
		t.Fatalf("expected name empty, got %q", payload.Name)
	}

	if payload.Path != "mybackups/mybackup.tar.gz" {
		t.Fatalf("expected path mybackups/mybackup.tar.gz, got %q", payload.Path)
	}

	if payload.Source != "sftp-server" {
		t.Fatalf("expected source sftp-server, got %q", payload.Source)
	}
}

func TestBackupRestoreModelImportPayloadJSON(t *testing.T) {
	t.Parallel()

	payload := backupRestoreModel{
		EncryptionKey: types.StringValue("mykey-12345"),
		Name:          types.StringValue(""),
		Path:          types.StringValue("mybackups/mybackup.tar.gz"),
		Source:        types.StringValue("sftp-server"),
	}.importPayload()

	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("unexpected marshal error: %v", err)
	}

	expected := `{"encryptionKey":"mykey-12345","name":"","path":"mybackups/mybackup.tar.gz","source":"sftp-server"}`
	if string(data) != expected {
		t.Fatalf("expected JSON %s, got %s", expected, string(data))
	}
}

func TestBackupRestoreModelRestorePayloadJSON(t *testing.T) {
	t.Parallel()

	payload := backupRestoreModel{
		IgnorePersistentIPs:             types.BoolValue(true),
		Type:                            types.StringValue(restoreTypeFull),
		IncludeTelemetryOperationalData: types.BoolValue(false),
	}.restorePayload()

	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("unexpected marshal error: %v", err)
	}

	expected := `{"ignorePersistentIPs":true,"type":"full","includeTelemetryOperationalData":false}`
	if string(data) != expected {
		t.Fatalf("expected JSON %s, got %s", expected, string(data))
	}
}

func TestBackupRestoreActionConfigureNilProviderData(t *testing.T) {
	t.Parallel()

	a := &backupRestoreAction{}
	resp := &action.ConfigureResponse{}

	a.Configure(context.Background(), action.ConfigureRequest{}, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("expected no diagnostics for nil ProviderData, got %v", resp.Diagnostics)
	}
}

func TestBackupRestoreActionConfigureWrongProviderData(t *testing.T) {
	t.Parallel()

	a := &backupRestoreAction{}
	resp := &action.ConfigureResponse{}

	a.Configure(context.Background(), action.ConfigureRequest{ProviderData: "wrong"}, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics for wrong ProviderData")
	}
}

func TestBackupRestoreActionConfigureMissingInfraModule(t *testing.T) {
	t.Parallel()

	a := &backupRestoreAction{}
	resp := &action.ConfigureResponse{}

	a.Configure(context.Background(), action.ConfigureRequest{
		ProviderData: fakeClientProvider{modules: map[string]interface{}{}},
	}, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics for missing infra module")
	}
}

func TestEvaluateRestoreStatusSuccess(t *testing.T) {
	t.Parallel()

	state, detail := evaluateRestoreStatus(mustBackupStatus(t, `{"id":"restore-20260629-143142","operation":"restore","state":"completed","details":{"progress":100}}`))

	if state != restoreStatusSucceeded {
		t.Fatalf("expected success state, got %v detail %q", state, detail)
	}
}

func TestEvaluateRestoreStatusFailure(t *testing.T) {
	t.Parallel()

	state, detail := evaluateRestoreStatus(mustBackupStatus(t, `{"operation":"restore","state":"failed","details":{"progress":40}}`))

	if state != restoreStatusFailed {
		t.Fatalf("expected failure state, got %v detail %q", state, detail)
	}
}

func TestEvaluateRestoreStatusFailedToValidate(t *testing.T) {
	t.Parallel()

	state, detail := evaluateRestoreStatus(mustBackupStatus(t, `{"operation":"restore","state":"failedToValidate","details":{"progress":0}}`))

	if state != restoreStatusFailed {
		t.Fatalf("expected failure state, got %v detail %q", state, detail)
	}
}

func TestEvaluateRestoreStatusInProgress(t *testing.T) {
	t.Parallel()

	state, detail := evaluateRestoreStatus(mustBackupStatus(t, `{"id":"restore-20260629-143142","operation":"restore","state":"processing","details":{"progress":85}}`))

	if state != restoreStatusInProgress {
		t.Fatalf("expected in-progress state, got %v detail %q", state, detail)
	}
}

func TestEvaluateRestoreStatusSpecInProgressStates(t *testing.T) {
	t.Parallel()

	for _, statusState := range []string{
		backupStatusIdle,
		backupStatusDownloading,
		backupStatusReady,
		backupStatusProcessing,
	} {
		state, detail := evaluateRestoreStatus(mustBackupStatus(t, fmt.Sprintf(`{"operation":"restore","state":%q}`, statusState)))
		if state != restoreStatusInProgress {
			t.Fatalf("expected %q to be in-progress, got %v detail %q", statusState, state, detail)
		}
	}
}

func TestEvaluateRestoreStatusNonRestoreOperation(t *testing.T) {
	t.Parallel()

	state, detail := evaluateRestoreStatus(mustBackupStatus(t, `{"operation":"backup","state":"processing","details":{"progress":85}}`))

	if state != restoreStatusUnknown {
		t.Fatalf("expected unknown state for non-restore operation, got %v detail %q", state, detail)
	}
}

func TestEvaluateRestoreStatusUnknown(t *testing.T) {
	t.Parallel()

	state, detail := evaluateRestoreStatus(mustBackupStatus(t, `{"operation":"restore","state":"somethingElse"}`))

	if state != restoreStatusUnknown {
		t.Fatalf("expected unknown state, got %v detail %q", state, detail)
	}
}

func TestRestoreStatusDetailIncludesError(t *testing.T) {
	t.Parallel()

	detail := restoreStatusDetail(mustBackupStatus(t, `{"operation":"restore","state":"failed","error":"restore service failed"}`))

	if !strings.Contains(detail, "error=restore service failed") {
		t.Fatalf("expected detail to include API error, got %q", detail)
	}
}

func mustBackupStatus(t *testing.T, raw string) api.BackupStatus {
	t.Helper()

	status := api.BackupStatus{Raw: json.RawMessage(raw)}
	if err := json.Unmarshal([]byte(raw), &status); err != nil {
		t.Fatalf("failed to unmarshal backup status: %v", err)
	}
	return status
}
