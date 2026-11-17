package datasource_switch_config_diff

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	nd "github.com/netascode/go-nd"
)

func newSwitchConfigDiffTestSource(test *testing.T, handler http.HandlerFunc) *switchConfigDiffDataSource {
	test.Helper()
	server := httptest.NewServer(handler)
	test.Cleanup(server.Close)
	client, err := nd.NewClient(server.URL, "/api/v1", "", "", "", false, nd.MaxRetries(0))
	if err != nil {
		test.Fatalf("create mock client: %v", err)
	}
	client.Token = "test-token"
	client.AuthTimeStamp = time.Now()
	client.AuthTokenTimeout = time.Hour
	return &switchConfigDiffDataSource{manageClient: &manage.NexusDashboardManage{ApiClient: &client}}
}

func switchConfigDiffReadRequest(test *testing.T, fabricName, switchID types.String) (datasource.ReadRequest, datasource.ReadResponse) {
	test.Helper()
	ctx := context.Background()
	sourceSchema := SwitchConfigDiffDataSourceSchema(ctx)
	configState := tfsdk.State{Schema: sourceSchema}
	config := SwitchConfigDiffModel{
		FabricName:  fabricName,
		SwitchId:    switchID,
		DiffConfigs: types.ListNull(DiffConfigsValue{}.Type(ctx)),
	}
	if diagnostics := configState.Set(ctx, &config); diagnostics.HasError() {
		test.Fatalf("create configuration: %v", diagnostics)
	}
	return datasource.ReadRequest{Config: tfsdk.Config{Schema: sourceSchema, Raw: configState.Raw}},
		datasource.ReadResponse{State: tfsdk.State{Schema: sourceSchema}}
}

func TestSwitchConfigDiffRead(test *testing.T) {
	testCases := []struct {
		name        string
		payload     string
		statusCode  int
		wantError   string
		wantEntries int
	}{
		{
			name: "ordered differences",
			payload: `{"diffConfigs":[
				{"differenceType":"equal","expectedConfig":"interface Ethernet1/1","runningConfig":"interface Ethernet1/1"},
				{"differenceType":"insert","expectedConfig":"  description leaf uplink","runningConfig":""},
				{"differenceType":"delete","expectedConfig":"","runningConfig":"  shutdown"}
			]}`,
			wantEntries: 3,
		},
		{name: "empty differences", payload: `{"diffConfigs":[]}`},
		{name: "missing differences", payload: `{}`},
		{name: "null differences", payload: `{"diffConfigs":null}`},
		{name: "response identifiers do not replace inputs", payload: `{"fabricName":"other-fabric","switchId":"other-switch","diffConfigs":[]}`},
		{name: "API error", statusCode: http.StatusBadRequest, payload: `{"error":"switch not found"}`, wantError: "Error Reading Switch Configuration Diff"},
		{name: "invalid JSON", payload: `{`, wantError: "Error Decoding Switch Configuration Diff"},
		{name: "invalid differences type", payload: `{"diffConfigs":"invalid"}`, wantError: "Error Decoding Switch Configuration Diff"},
	}
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			var requestCount atomic.Int32
			source := newSwitchConfigDiffTestSource(test, func(writer http.ResponseWriter, request *http.Request) {
				requestCount.Add(1)
				if request.Method != http.MethodGet || request.URL.Path != "/api/v1/manage/fabrics/AK_VXLAN/switches/9QOQ1PKTKPU/diff" {
					test.Errorf("unexpected request: %s %s", request.Method, request.URL.Path)
				}
				if request.URL.RawQuery != "forceShowRun=false" {
					test.Errorf("unexpected query: %s", request.URL.RawQuery)
				}
				writer.Header().Set("Content-Type", "application/json")
				if testCase.statusCode != 0 {
					writer.WriteHeader(testCase.statusCode)
				}
				_, _ = writer.Write([]byte(testCase.payload))
			})
			request, response := switchConfigDiffReadRequest(test, types.StringValue("AK_VXLAN"), types.StringValue("9QOQ1PKTKPU"))
			source.Read(context.Background(), request, &response)
			if requestCount.Load() != 1 {
				test.Errorf("expected one API request, got %d", requestCount.Load())
			}
			if testCase.wantError != "" {
				if !response.Diagnostics.HasError() || response.Diagnostics[0].Summary() != testCase.wantError {
					test.Fatalf("expected %q, got %v", testCase.wantError, response.Diagnostics)
				}
				if !response.State.Raw.IsNull() {
					test.Error("failed read populated state")
				}
				return
			}
			if response.Diagnostics.HasError() {
				test.Fatalf("read diff: %v", response.Diagnostics)
			}
			var state SwitchConfigDiffModel
			if diagnostics := response.State.Get(context.Background(), &state); diagnostics.HasError() {
				test.Fatalf("decode state: %v", diagnostics)
			}
			if state.FabricName.ValueString() != "AK_VXLAN" || state.SwitchId.ValueString() != "9QOQ1PKTKPU" {
				test.Errorf("configured identifiers changed: %v", state)
			}
			if len(state.DiffConfigs.Elements()) != testCase.wantEntries {
				test.Fatalf("expected %d differences, got %d", testCase.wantEntries, len(state.DiffConfigs.Elements()))
			}
			if testCase.wantEntries == 0 {
				return
			}
			var entries []DiffConfigsValue
			if diagnostics := state.DiffConfigs.ElementsAs(context.Background(), &entries, false); diagnostics.HasError() {
				test.Fatalf("decode differences: %v", diagnostics)
			}
			for index, expectedType := range []string{"equal", "insert", "delete"} {
				if entries[index].DifferenceType.ValueString() != expectedType {
					test.Errorf("entry %d: expected %q, got %q", index, expectedType, entries[index].DifferenceType.ValueString())
				}
			}
			if entries[0].ExpectedConfig.ValueString() != "interface Ethernet1/1" || entries[0].RunningConfig.ValueString() != "interface Ethernet1/1" ||
				entries[1].ExpectedConfig.ValueString() != "  description leaf uplink" || !entries[1].RunningConfig.IsNull() ||
				!entries[2].ExpectedConfig.IsNull() || entries[2].RunningConfig.ValueString() != "  shutdown" {
				test.Errorf("configuration lines changed: %v", entries)
			}
		})
	}
}

func TestSwitchConfigDiffReadValidation(test *testing.T) {
	testCases := []struct {
		name       string
		fabricName types.String
		switchID   types.String
		noClient   bool
		wantError  string
	}{
		{name: "null fabric", fabricName: types.StringNull(), switchID: types.StringValue("switch1"), wantError: "Fabric Name Required"},
		{name: "unknown fabric", fabricName: types.StringUnknown(), switchID: types.StringValue("switch1"), wantError: "Fabric Name Required"},
		{name: "blank fabric", fabricName: types.StringValue(" \t"), switchID: types.StringValue("switch1"), wantError: "Fabric Name Required"},
		{name: "null switch", fabricName: types.StringValue("fabric1"), switchID: types.StringNull(), wantError: "Switch ID Required"},
		{name: "unknown switch", fabricName: types.StringValue("fabric1"), switchID: types.StringUnknown(), wantError: "Switch ID Required"},
		{name: "blank switch", fabricName: types.StringValue("fabric1"), switchID: types.StringValue(" \t"), wantError: "Switch ID Required"},
		{name: "missing client", fabricName: types.StringValue("fabric1"), switchID: types.StringValue("switch1"), noClient: true, wantError: "Manage Client Not Configured"},
	}
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			source := newSwitchConfigDiffTestSource(test, func(writer http.ResponseWriter, request *http.Request) {
				test.Error("invalid configuration made an API request")
			})
			if testCase.noClient {
				source.manageClient = nil
			}
			request, response := switchConfigDiffReadRequest(test, testCase.fabricName, testCase.switchID)
			source.Read(context.Background(), request, &response)
			if !response.Diagnostics.HasError() || response.Diagnostics[0].Summary() != testCase.wantError {
				test.Fatalf("expected %q, got %v", testCase.wantError, response.Diagnostics)
			}
		})
	}
}

func TestSwitchConfigDiffFabricLock(test *testing.T) {
	source := newSwitchConfigDiffTestSource(test, func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{"diffConfigs":[]}`))
	})
	lockedFabric := "switch-config-diff-locked-fabric"
	guard := ndapi.Acquire(lockedFabric, "config", ndapi.LockDeploy)
	test.Cleanup(func() {
		if guard != nil {
			guard.Release()
		}
	})
	request, response := switchConfigDiffReadRequest(test, types.StringValue(lockedFabric), types.StringValue("switch1"))
	started := make(chan struct{})
	finished := make(chan struct{})
	go func() {
		close(started)
		source.Read(context.Background(), request, &response)
		close(finished)
	}()
	<-started
	select {
	case <-finished:
		test.Fatal("same-fabric read bypassed the deployment lock")
	case <-time.After(100 * time.Millisecond):
	}
	otherRequest, otherResponse := switchConfigDiffReadRequest(test, types.StringValue("switch-config-diff-other-fabric"), types.StringValue("switch2"))
	otherFinished := make(chan struct{})
	go func() {
		source.Read(context.Background(), otherRequest, &otherResponse)
		close(otherFinished)
	}()
	select {
	case <-otherFinished:
		if otherResponse.Diagnostics.HasError() {
			test.Fatalf("read other fabric: %v", otherResponse.Diagnostics)
		}
	case <-time.After(5 * time.Second):
		test.Fatal("deployment lock blocked a different fabric")
	}
	guard.Release()
	guard = nil
	select {
	case <-finished:
		if response.Diagnostics.HasError() {
			test.Fatalf("read after deployment unlock: %v", response.Diagnostics)
		}
	case <-time.After(5 * time.Second):
		test.Fatal("same-fabric read did not resume after deployment unlock")
	}
}

type switchConfigDiffTestProvider struct {
	module interface{}
}

func (provider switchConfigDiffTestProvider) GetModule(name string) interface{} {
	if name != ModuleKey {
		return nil
	}
	return provider.module
}

func TestSwitchConfigDiffConfigure(test *testing.T) {
	manageClient := &manage.NexusDashboardManage{}
	testCases := []struct {
		name         string
		providerData interface{}
		wantError    string
		wantClient   *manage.NexusDashboardManage
	}{
		{name: "nil provider"},
		{name: "wrong provider type", providerData: "invalid", wantError: "Unexpected DataSource Configure Type"},
		{name: "missing module", providerData: switchConfigDiffTestProvider{}, wantError: "Manage Module Not Found"},
		{name: "wrong module type", providerData: switchConfigDiffTestProvider{module: "invalid"}, wantError: "Unexpected Manage Module Type"},
		{name: "configured", providerData: switchConfigDiffTestProvider{module: manageClient}, wantClient: manageClient},
	}
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			source := NewSwitchConfigDiffDataSource().(*switchConfigDiffDataSource)
			var response datasource.ConfigureResponse
			source.Configure(context.Background(), datasource.ConfigureRequest{ProviderData: testCase.providerData}, &response)
			if testCase.wantError != "" {
				if !response.Diagnostics.HasError() || response.Diagnostics[0].Summary() != testCase.wantError {
					test.Fatalf("expected %q, got %v", testCase.wantError, response.Diagnostics)
				}
				return
			}
			if response.Diagnostics.HasError() || source.manageClient != testCase.wantClient {
				test.Fatalf("unexpected configuration result: client=%p, diagnostics=%v", source.manageClient, response.Diagnostics)
			}
		})
	}
}

func TestSwitchConfigDiffMetadataAndSchema(test *testing.T) {
	ctx := context.Background()
	source := NewSwitchConfigDiffDataSource()
	var metadata datasource.MetadataResponse
	source.Metadata(ctx, datasource.MetadataRequest{ProviderTypeName: "nd"}, &metadata)
	if metadata.TypeName != "nd_switch_config_diff" {
		test.Errorf("unexpected data source name: %s", metadata.TypeName)
	}
	var response datasource.SchemaResponse
	source.Schema(ctx, datasource.SchemaRequest{}, &response)
	if diagnostics := response.Schema.ValidateImplementation(ctx); diagnostics.HasError() {
		test.Fatalf("invalid schema: %v", diagnostics)
	}
	if !response.Schema.Attributes["fabric_name"].IsRequired() || !response.Schema.Attributes["switch_id"].IsRequired() || !response.Schema.Attributes["diff_configs"].IsComputed() {
		test.Error("unexpected data source attribute modes")
	}
}
