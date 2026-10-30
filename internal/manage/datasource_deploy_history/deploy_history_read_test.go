package datasource_deploy_history

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"terraform-provider-nd/internal/manage"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	nd "github.com/netascode/go-nd"
)

func TestDeployHistoryReadRecordLimit(test *testing.T) {
	testCases := []struct {
		name              string
		recordLimit       types.Int64
		returnedCount     int
		wantCount         int
		wantMax           int64
		timestamps        []string
		wantRecordIndexes []int
	}{
		{name: "oversized response", recordLimit: types.Int64Value(1), returnedCount: 2, wantCount: 1, wantMax: 1},
		{name: "exact response", recordLimit: types.Int64Value(2), returnedCount: 2, wantCount: 2, wantMax: 2},
		{name: "short response", recordLimit: types.Int64Value(3), returnedCount: 2, wantCount: 2, wantMax: 3},
		{name: "empty response", recordLimit: types.Int64Value(1), returnedCount: 0, wantCount: 0, wantMax: 1},
		{name: "default limit", recordLimit: types.Int64Null(), returnedCount: 12, wantCount: 10, wantMax: 10},
		{name: "short default response", recordLimit: types.Int64Null(), returnedCount: 2, wantCount: 2, wantMax: 10},
		{
			name: "sort before trimming", recordLimit: types.Int64Value(2), returnedCount: 3, wantCount: 2, wantMax: 2,
			timestamps:        []string{"2025-09-18T03:14:00.477Z", "2025-09-18T03:14:00.491Z", "2025-09-18T03:14:00.480Z"},
			wantRecordIndexes: []int{1, 2},
		},
		{
			name: "sort without trimming", recordLimit: types.Int64Value(4), returnedCount: 3, wantCount: 3, wantMax: 4,
			timestamps:        []string{"2025-09-18T03:14:00.477Z", "2025-09-18T03:14:00.491Z", "2025-09-18T03:14:00.480Z"},
			wantRecordIndexes: []int{1, 2, 0},
		},
		{
			name: "timezone offsets", recordLimit: types.Int64Value(1), returnedCount: 2, wantCount: 1, wantMax: 1,
			timestamps:        []string{"2025-09-18T04:14:00.491+01:00", "2025-09-18T03:14:00.500Z"},
			wantRecordIndexes: []int{1},
		},
		{
			name: "nanosecond precision", recordLimit: types.Int64Value(1), returnedCount: 2, wantCount: 1, wantMax: 1,
			timestamps:        []string{"2025-09-18T03:14:00Z", "2025-09-18T03:14:00.000000001Z"},
			wantRecordIndexes: []int{1},
		},
		{
			name: "equal timestamps preserve response order", recordLimit: types.Int64Value(2), returnedCount: 3, wantCount: 2, wantMax: 2,
			timestamps:        []string{"2025-09-18T03:14:00.491Z", "2025-09-18T03:14:00.491000000Z", "2025-09-18T03:14:00.491+00:00"},
			wantRecordIndexes: []int{0, 1},
		},
		{
			name: "missing and invalid timestamps sort last", recordLimit: types.Int64Value(4), returnedCount: 4, wantCount: 4, wantMax: 4,
			timestamps:        []string{"", "not-a-timestamp", "2025-09-18T03:14:00.491Z", "2025-09-18T03:14:00.477Z"},
			wantRecordIndexes: []int{2, 3, 0, 1},
		},
		{
			name: "trim missing and invalid timestamps first", recordLimit: types.Int64Value(2), returnedCount: 4, wantCount: 2, wantMax: 2,
			timestamps:        []string{"", "not-a-timestamp", "2025-09-18T03:14:00.491Z", "2025-09-18T03:14:00.477Z"},
			wantRecordIndexes: []int{2, 3},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			ctx := context.Background()
			records := make([]map[string]any, testCase.returnedCount)
			baseTime := time.Date(2025, time.September, 18, 3, 14, 0, 0, time.UTC)
			for recordIndex := range records {
				completionTimestamp := baseTime.Add(-time.Duration(recordIndex) * time.Millisecond).Format(time.RFC3339Nano)
				if testCase.timestamps != nil {
					completionTimestamp = testCase.timestamps[recordIndex]
				}
				records[recordIndex] = map[string]any{
					"hostname":          fmt.Sprintf("v1-bgw%d", recordIndex+2),
					"completeTimestamp": completionTimestamp,
					"startTimestamp":    baseTime.Add(time.Duration(recordIndex) * time.Millisecond).Format(time.RFC3339Nano),
					"status":            "success",
					"ticketId":          nil,
					"configCommandResponses": []map[string]any{
						{"command": "interface ethernet1/2", "status": "success", "commandExecutionStatus": "redundant", "cliResponse": ""},
						{"command": "configure terminal", "status": "success", "cliResponse": ""},
					},
				}
			}
			payload := map[string]any{
				"deploymentRecords": records,
				"meta": map[string]any{
					"counts": map[string]int{"remaining": 2, "total": 10},
					"links": map[string]string{
						"next":     "https://api.example.com/deploymentHistory?offset=8&max=2",
						"previous": "https://api.example.com/deploymentHistory?offset=4&max=2",
					},
				},
			}

			var requestCount atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				requestCount.Add(1)
				if request.Method != http.MethodGet || request.URL.Path != "/api/v1/manage/fabrics/fabric1/deploymentHistory" {
					test.Errorf("unexpected request: %s %s", request.Method, request.URL.Path)
				}
				query := request.URL.Query()
				if query.Get("max") != strconv.FormatInt(testCase.wantMax, 10) || query.Get("offset") != "0" || query.Get("sort") != "completeTimestamp:desc" {
					test.Errorf("unexpected query parameters: %v", query)
				}
				writer.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(writer).Encode(payload); err != nil {
					test.Errorf("encode response: %v", err)
				}
			}))
			defer server.Close()

			client, err := nd.NewClient(server.URL, "/api/v1", "", "", "", false, nd.MaxRetries(0))
			if err != nil {
				test.Fatalf("create mock client: %v", err)
			}
			client.Token = "test-token"
			client.AuthTimeStamp = time.Now()
			client.AuthTokenTimeout = time.Hour
			source := &deployHistoryDataSource{
				manageClient: &manage.NexusDashboardManage{ApiClient: &client},
			}

			schema := DeployHistoryDataSourceSchema(ctx)
			configState := tfsdk.State{Schema: schema}
			config := DeployHistoryModel{
				FabricName:        types.StringValue("fabric1"),
				NoOfRecords:       testCase.recordLimit,
				DeploymentRecords: types.ListNull(DeploymentRecordsValue{}.Type(ctx)),
			}
			if diagnostics := configState.Set(ctx, &config); diagnostics.HasError() {
				test.Fatalf("create configuration: %v", diagnostics)
			}
			request := datasource.ReadRequest{
				Config: tfsdk.Config{Schema: schema, Raw: configState.Raw},
			}
			response := datasource.ReadResponse{State: tfsdk.State{Schema: schema}}
			source.Read(ctx, request, &response)
			if response.Diagnostics.HasError() {
				test.Fatalf("read deployment history: %v", response.Diagnostics)
			}
			if requestCount.Load() != 1 {
				test.Errorf("expected one API request, got %d", requestCount.Load())
			}

			var state DeployHistoryModel
			if diagnostics := response.State.Get(ctx, &state); diagnostics.HasError() {
				test.Fatalf("decode state: %v", diagnostics)
			}
			if len(state.DeploymentRecords.Elements()) != testCase.wantCount {
				test.Fatalf("expected %d records in state, got %d", testCase.wantCount, len(state.DeploymentRecords.Elements()))
			}
			if !state.FabricName.Equal(config.FabricName) || !state.NoOfRecords.Equal(config.NoOfRecords) {
				test.Error("configured fabric name or record limit changed in state")
			}
			var savedRecords []DeploymentRecordsValue
			if diagnostics := state.DeploymentRecords.ElementsAs(ctx, &savedRecords, false); diagnostics.HasError() {
				test.Fatalf("decode deployment records: %v", diagnostics)
			}
			for recordIndex, record := range savedRecords {
				wantRecordIndex := recordIndex
				if testCase.wantRecordIndexes != nil {
					wantRecordIndex = testCase.wantRecordIndexes[recordIndex]
				}
				wantHostname := fmt.Sprintf("v1-bgw%d", wantRecordIndex+2)
				if record.Hostname.ValueString() != wantHostname {
					test.Errorf("record %d: expected hostname %q, got %q", recordIndex, wantHostname, record.Hostname.ValueString())
				}
				if !record.TicketId.IsNull() {
					test.Errorf("record %d: null ticketId was not preserved", recordIndex)
				}
				var commands []ConfigCommandResponsesValue
				if diagnostics := record.ConfigCommandResponses.ElementsAs(ctx, &commands, false); diagnostics.HasError() {
					test.Fatalf("decode command responses: %v", diagnostics)
				}
				if len(commands) != 2 || commands[0].Command.ValueString() != "interface ethernet1/2" || commands[1].Command.ValueString() != "configure terminal" {
					test.Errorf("record %d: nested command responses changed", recordIndex)
				}
				for _, command := range commands {
					if command.Status.ValueString() != "success" {
						test.Errorf("record %d: command status was not preserved", recordIndex)
					}
				}
			}
		})
	}
}

func TestDeployHistoryCommandResponseSchema(test *testing.T) {
	ctx := context.Background()
	dataSourceSchema := DeployHistoryDataSourceSchema(ctx)
	records := dataSourceSchema.Attributes["deployment_records"].(schema.ListNestedAttribute)
	commands := records.NestedObject.Attributes["config_command_responses"].(schema.ListNestedAttribute)
	if _, exists := commands.NestedObject.Attributes["command_execution_status"]; exists {
		test.Error("redundant command_execution_status remains in the schema")
	}
	if _, exists := commands.NestedObject.Attributes["status"]; !exists {
		test.Error("command status is missing from the schema")
	}
	if _, exists := (ConfigCommandResponsesValue{}).AttributeTypes(ctx)["command_execution_status"]; exists {
		test.Error("redundant command_execution_status remains in the object type")
	}
}
