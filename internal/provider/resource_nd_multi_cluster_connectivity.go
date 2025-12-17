package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/CiscoDevNet/terraform-provider-nd/internal/client"
	"github.com/Jeffail/gabs/v2"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ClusterResource{}
var _ resource.ResourceWithImportState = &ClusterResource{}

var clusterPath = "/api/v1/infra/clusters"

func NewClusterResource() resource.Resource {
	return &ClusterResource{}
}

// ClusterResource defines the resource implementation.
type ClusterResource struct {
	client *client.Client
}

// ClusterResourceModel describes the resource data model.
type ClusterResourceModel struct {
	Id                         types.String  `tfsdk:"id"`
	ClusterType                types.String  `tfsdk:"cluster_type"`
	ClusterHostname            types.String  `tfsdk:"cluster_hostname"`
	ClusterUsername            types.String  `tfsdk:"cluster_username"`
	ClusterPassword            types.String  `tfsdk:"cluster_password"`
	ClusterLoginDomain         types.String  `tfsdk:"cluster_login_domain"`
	MultiClusterLoginDomain    types.String  `tfsdk:"multi_cluster_login_domain"`
	FabricName                 types.String  `tfsdk:"fabric_name"`
	LicenseTier                types.String  `tfsdk:"license_tier"`
	Features                   types.Set     `tfsdk:"features"`
	InbandEpg                  types.String  `tfsdk:"inband_epg"`
	SecurityDomain             types.String  `tfsdk:"security_domain"`
	ValidatePeerCertificate    types.Bool    `tfsdk:"validate_peer_certificate"`
	Latitude                   types.Float64 `tfsdk:"latitude"`
	Longitude                  types.Float64 `tfsdk:"longitude"`
	TelemetryStreamingProtocol types.String  `tfsdk:"telemetry_streaming_protocol"`
	TelemetryNetwork           types.String  `tfsdk:"telemetry_network"`
}

func getBaseClusterResourceModel(username, password, clusterLoginDomain, multiClusterLoginDomain basetypes.StringValue) *ClusterResourceModel {
	return &ClusterResourceModel{
		Id:                         basetypes.NewStringNull(),
		ClusterType:                basetypes.NewStringNull(),
		ClusterHostname:            basetypes.NewStringNull(),
		ClusterUsername:            basetypes.NewStringValue(username.ValueString()),
		ClusterPassword:            basetypes.NewStringValue(password.ValueString()),
		ClusterLoginDomain:         basetypes.NewStringValue(clusterLoginDomain.ValueString()),
		MultiClusterLoginDomain:    basetypes.NewStringValue(multiClusterLoginDomain.ValueString()),
		FabricName:                 basetypes.NewStringNull(),
		LicenseTier:                basetypes.NewStringNull(),
		Features:                   basetypes.NewSetNull(types.StringType),
		InbandEpg:                  basetypes.NewStringNull(),
		SecurityDomain:             basetypes.NewStringNull(),
		ValidatePeerCertificate:    basetypes.NewBoolNull(),
		Latitude:                   basetypes.NewFloat64Null(),
		Longitude:                  basetypes.NewFloat64Null(),
		TelemetryStreamingProtocol: basetypes.NewStringNull(),
		TelemetryNetwork:           basetypes.NewStringNull(),
	}
}

func (r *ClusterResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if !req.Plan.Raw.IsNull() {
		var planData, stateData, configData *ClusterResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)
		resp.Diagnostics.Append(req.Config.Get(ctx, &configData)...)

		if configData.ClusterType.ValueString() == "nd" {
			if configData.LicenseTier.ValueString() != "" {
				resp.Diagnostics.AddError("The license_tier is invalid attribute for cluster_type: nd", "The license_tier attribute is only applicable when cluster_type is set to apic.")
			}
			if !configData.Features.IsNull() && !configData.Features.IsUnknown() {
				resp.Diagnostics.AddError("The features is invalid attribute for cluster_type: nd", "The features attribute is only applicable when cluster_type is set to apic.")
			}
			if configData.InbandEpg.ValueString() != "" {
				resp.Diagnostics.AddError("The inband_epg is invalid attribute for cluster_type: nd", "The inband_epg attribute is only applicable when cluster_type is set to apic.")
			}
			if configData.SecurityDomain.ValueString() != "" {
				resp.Diagnostics.AddError("The security_domain is invalid attribute for cluster_type: nd", "The security_domain attribute is only applicable when cluster_type is set to apic.")
			}
			if !configData.ValidatePeerCertificate.IsNull() && !configData.ValidatePeerCertificate.IsUnknown() {
				resp.Diagnostics.AddError("The validate_peer_certificate is invalid attribute for cluster_type: nd", "The validate_peer_certificate attribute is only applicable when cluster_type is set to apic.")
			}
			if configData.TelemetryStreamingProtocol.ValueString() != "" {
				resp.Diagnostics.AddError("The telemetry_streaming_protocol is invalid attribute for cluster_type: nd", "The telemetry_streaming_protocol attribute is only applicable when cluster_type is set to apic.")
			}
			if configData.TelemetryNetwork.ValueString() != "" {
				resp.Diagnostics.AddError("The telemetry_network is invalid attribute for cluster_type: nd", "The telemetry_network attribute is only applicable when cluster_type is set to apic.")
			}
		} else if planData.ClusterType.ValueString() == "apic" {
			if configData.ClusterLoginDomain.ValueString() != "" {
				resp.Diagnostics.AddError("The cluster_login_domain is invalid attribute for cluster_type: apic", "The cluster_login_domain attribute is only applicable when cluster_type is set to nd.")
			}
			if configData.MultiClusterLoginDomain.ValueString() != "" {
				resp.Diagnostics.AddError("The multi_cluster_login_domain is invalid attribute for cluster_type: apic", "The multi_cluster_login_domain attribute is only applicable when cluster_type is set to nd.")
			}
		}

		if resp.Diagnostics.HasError() {
			return
		}

		if stateData != nil {
			if resp.Diagnostics.HasError() {
				return
			}

			// The cluster_username attribute is required, but it is not included in the ND API response data, so its state will be set to an empty string.
			// When this condition is met, an empty string will be assigned to the plan's cluster_username attribute to ignore any plantime changes.
			if configData.ClusterUsername.ValueString() != "" && stateData.ClusterUsername.ValueString() == "" {
				planData.ClusterUsername = basetypes.NewStringValue("")
			}

			// The cluster_password attribute is required, but it is not included in the ND API response data, so its state will be set to an empty string.
			// When this condition is met, an empty string will be assigned to the plan's cluster_password attribute to ignore any plantime changes.
			if configData.ClusterPassword.ValueString() != "" && stateData.ClusterPassword.ValueString() == "" {
				planData.ClusterPassword = basetypes.NewStringValue("")
			}

		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &planData)...)
	}
}

func (r *ClusterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: nd_multi_cluster_connectivity")
	resp.TypeName = req.ProviderTypeName + "_multi_cluster_connectivity"
	tflog.Debug(ctx, "End metadata of resource: nd_multi_cluster_connectivity")
}

func (r *ClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: nd_multi_cluster_connectivity")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Manages Multi-cluster connectivity for Nexus Dashboard",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The ID of the cluster.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cluster_type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The type of the cluster. Allowed values are 'nd', or 'apic'.",
				Validators: []validator.String{
					stringvalidator.OneOf("nd", "apic"),
				},
			},
			"cluster_hostname": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The URL or Hostname of the cluster.",
			},
			"cluster_username": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The username of the cluster.",
			},
			"cluster_password": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The password of the cluster.",
				Sensitive:           true,
			},
			"cluster_login_domain": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The login domain of the cluster. This attribute is only applicable when cluster_type is set to nd.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"multi_cluster_login_domain": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The multi cluster login domain of the cluster. This attribute is only applicable when cluster_type is set to nd.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"fabric_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the cluster.",
			},
			"license_tier": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The clusters license tier. Only one value can be specified at a time. Allowed values are 'advantage', or 'essentials', or 'premier'. This attribute is only applicable when cluster_type is set to apic.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("advantage", "essentials", "premier"),
				},
			},
			"features": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "The features of the cluster. Allowed values are 'telemetry', 'orchestration'. This attribute is only applicable when cluster_type is set to apic.",
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf("telemetry", "orchestration"),
					),
				},
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
					setplanmodifier.RequiresReplace(),
				},
			},
			"inband_epg": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The Inband EPG name of the cluster. This attribute is only applicable when cluster_type is set to apic.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"security_domain": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The security domain of the cluster. This attribute is only applicable when cluster_type is set to apic.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"validate_peer_certificate": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The validate peer certificate flag of the cluster. This attribute is only applicable when cluster_type is set to apic.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"latitude": schema.Float64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The latitude coordinate of the cluster.",
			},
			"longitude": schema.Float64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The longitude coordinate of the cluster.",
			},
			"telemetry_streaming_protocol": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The telemetry streaming protocol of the cluster. Allowed values are 'ipv4', or 'ipv6'. This attribute is only applicable when cluster_type is set to apic.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("ipv4", "ipv6"),
				},
			},
			"telemetry_network": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The telemetry network type of the cluster. Allowed values are 'inband', or 'outband'. This attribute is only applicable when cluster_type is set to apic.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("inband", "outband"),
				},
			},
		},
	}
	tflog.Debug(ctx, "End schema of resource: nd_multi_cluster_connectivity")
}

func (r *ClusterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: nd_multi_cluster_connectivity")
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
	tflog.Debug(ctx, "End configure of resource: nd_multi_cluster_connectivity")
}

func (r *ClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: nd_multi_cluster_connectivity")

	var stateData *ClusterResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)

	var planData *ClusterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	jsonPayload := getClusterJsonPayload(ctx, &resp.Diagnostics, planData, "POST")

	if resp.Diagnostics.HasError() {
		return
	}

	r.client.DoRestRequest(ctx, &resp.Diagnostics, clusterPath, "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	planData.Id = types.StringValue(planData.FabricName.ValueString())
	getAndSetResourceClusterAttributes(ctx, &resp.Diagnostics, r.client, planData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource nd_multi_cluster_connectivity with id '%s'", planData.Id.ValueString()))
}

func (r *ClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: nd_multi_cluster_connectivity")
	var stateData *ClusterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource nd_multi_cluster_connectivity with id '%s'", stateData.Id.ValueString()))

	getAndSetResourceClusterAttributes(ctx, &resp.Diagnostics, r.client, stateData)

	// Save updated data into Terraform state
	if stateData.Id.IsNull() {
		var emptyData *ClusterResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &stateData)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource nd_multi_cluster_connectivity with id '%s'", stateData.Id.ValueString()))
}

func (r *ClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: nd_multi_cluster_connectivity")

	var stateData *ClusterResourceModel
	var planData *ClusterResourceModel

	// Read Terraform plan data and state data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource nd_multi_cluster_connectivity with id '%s'", planData.Id.ValueString()))

	jsonPayload := getClusterJsonPayload(ctx, &resp.Diagnostics, planData, "PUT")

	if resp.Diagnostics.HasError() {
		return
	}

	r.client.DoRestRequest(ctx, &resp.Diagnostics, fmt.Sprintf("%s/%s", clusterPath, planData.Id.ValueString()), "PUT", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetResourceClusterAttributes(ctx, &resp.Diagnostics, r.client, planData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
	tflog.Debug(ctx, "End update of resource nd_multi_cluster_connectivity")
}

func (r *ClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: nd_multi_cluster_connectivity")
	var stateData *ClusterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource nd_multi_cluster_connectivity with id '%s'", stateData.Id.ValueString()))
	if resp.Diagnostics.HasError() {
		return
	}

	deletePayload := map[string]interface{}{}
	if stateData.ClusterType.ValueString() == "apic" {
		deletePayload["force"] = true
		deletePayload["credentials"] = map[string]string{
			"loginDomain": stateData.ClusterLoginDomain.ValueString(),
			"password":    stateData.ClusterPassword.ValueString(),
			"user":        stateData.ClusterUsername.ValueString(),
		}

	}
	payload, _ := json.Marshal(deletePayload)
	jsonPayload, _ := gabs.ParseJSON(payload)
	r.client.DoRestRequest(ctx, &resp.Diagnostics, fmt.Sprintf("%s/%s/remove", clusterPath, stateData.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	// Detaching the cluster from the controller may take a few seconds.
	time.Sleep(5 * time.Second)
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource nd_multi_cluster_connectivity with id '%s'", stateData.Id.ValueString()))
}

func (r *ClusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: nd_multi_cluster_connectivity")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *ClusterResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)

	// The username and password is required to update and delete the APIC cluster
	var clusterUsername, clusterPassword string

	// Read username and password from config file
	configPath := os.Getenv("CLUSTER_CREDENTIALS_FILE_LOCATION")
	if configPath != "" {
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			tflog.Error(ctx, fmt.Sprintf("Error: Config file does not exist at path: %s", configPath))
			return
		}

		data, err := os.ReadFile(configPath)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error: Unable to read config file: %v", err))
			return
		}

		var config map[string]map[string]string
		if err := json.Unmarshal(data, &config); err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error: Invalid JSON in config file: %v", err))
			return
		}

		importId := stateData.Id.ValueString()
		if config[importId] != nil {
			clusterUsername = config[importId]["username"]
			clusterPassword = config[importId]["password"]
		}
	} else {
		// Read username and password from environment variables
		clusterUsername = getStringAttributeValue(ctx, stateData.ClusterUsername, "CLUSTER_USERNAME")
		clusterPassword = getStringAttributeValue(ctx, stateData.ClusterPassword, "CLUSTER_PASSWORD")
	}

	resp.State.SetAttribute(ctx, path.Root("cluster_username"), basetypes.NewStringValue(clusterUsername))
	resp.State.SetAttribute(ctx, path.Root("cluster_password"), basetypes.NewStringValue(clusterPassword))

	tflog.Debug(ctx, fmt.Sprintf("Import state of resource nd_multi_cluster_connectivity with id '%s'", stateData.Id.ValueString()))
	tflog.Debug(ctx, "End import of state resource: nd_multi_cluster_connectivity")
}

func getClusterJsonPayload(ctx context.Context, diags *diag.Diagnostics, data *ClusterResourceModel, method string) *gabs.Container {
	payloadMap := map[string]interface{}{}

	clusterType := ""
	if !data.ClusterType.IsNull() && !data.ClusterType.IsUnknown() {
		clusterType = data.ClusterType.ValueString()
	}

	fabricName := ""
	if !data.FabricName.IsNull() && !data.FabricName.IsUnknown() {
		fabricName = data.FabricName.ValueString()
	}

	if clusterType == "apic" {
		aciMap := map[string]interface{}{}
		if !data.LicenseTier.IsNull() && !data.LicenseTier.IsUnknown() {
			aciMap["licenseTier"] = data.LicenseTier.ValueString()
		}
		aciMap["name"] = fabricName

		if !data.SecurityDomain.IsNull() && !data.SecurityDomain.IsUnknown() {
			aciMap["securityDomain"] = data.SecurityDomain.ValueString()
		}
		if !data.ValidatePeerCertificate.IsNull() && !data.ValidatePeerCertificate.IsUnknown() {
			aciMap["verifyCA"] = data.ValidatePeerCertificate.ValueBool()
		}

		if !data.Features.IsNull() && !data.Features.IsUnknown() {
			featureStrings := make([]string, 0)
			data.Features.ElementsAs(ctx, &featureStrings, false)
			for _, feature := range featureStrings {
				switch feature {
				case "telemetry":
					telemetryMap := map[string]string{"status": "enabled"}
					if !data.TelemetryNetwork.IsNull() && !data.TelemetryNetwork.IsUnknown() {
						telemetryMap["network"] = data.TelemetryNetwork.ValueString()
					}
					if !data.InbandEpg.IsNull() && !data.InbandEpg.IsUnknown() && data.InbandEpg.ValueString() != "" {
						telemetryMap["epg"] = fmt.Sprintf("uni/tn-mgmt/mgmtp-default/inb-%s", data.InbandEpg.ValueString())
					}
					if !data.TelemetryStreamingProtocol.IsNull() && !data.TelemetryStreamingProtocol.IsUnknown() {
						telemetryMap["streamingProtocol"] = data.TelemetryStreamingProtocol.ValueString()
					}
					aciMap["telemetry"] = telemetryMap
				case "orchestration":
					aciMap["orchestration"] = map[string]interface{}{
						"status": "enabled",
					}

				}
			}
		}

		payloadMap["aci"] = aciMap
	} else if clusterType == "nd" {
		if !data.MultiClusterLoginDomain.IsNull() && !data.MultiClusterLoginDomain.IsUnknown() {
			payloadMap["nd"] = map[string]interface{}{"multiClusterLoginDomainName": data.MultiClusterLoginDomain.ValueString()}
		}
	}

	payloadMap["clusterType"] = strings.ToUpper(clusterType)

	if method == "PUT" {
		payloadMap["name"] = fabricName
	}
	credentialsMap := map[string]interface{}{}
	if !data.ClusterUsername.IsNull() && !data.ClusterUsername.IsUnknown() {
		credentialsMap["user"] = data.ClusterUsername.ValueString()
	}

	if !data.ClusterPassword.IsNull() && !data.ClusterPassword.IsUnknown() {
		credentialsMap["password"] = data.ClusterPassword.ValueString()
	}

	if !data.ClusterLoginDomain.IsNull() && !data.ClusterLoginDomain.IsUnknown() {
		credentialsMap["loginDomain"] = data.ClusterLoginDomain.ValueString()
	}
	payloadMap["credentials"] = credentialsMap

	if !data.ClusterHostname.IsNull() && !data.ClusterHostname.IsUnknown() {
		payloadMap["onboardUrl"] = data.ClusterHostname.ValueString()
	}

	var latitude, longitude float64
	if !data.Latitude.IsNull() && !data.Latitude.IsUnknown() {
		latitude = data.Latitude.ValueFloat64()
	}
	if !data.Longitude.IsNull() && !data.Longitude.IsUnknown() {
		longitude = data.Longitude.ValueFloat64()
	}

	payloadMap["location"] = map[string]interface{}{
		"latitude":  latitude,
		"longitude": longitude,
	}

	payload, err := json.Marshal(map[string]interface{}{"spec": payloadMap})
	if err != nil {
		diags.AddError(
			"Marshalling of json payload failed",
			fmt.Sprintf("Error: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	jsonPayload, err := gabs.ParseJSON(payload)

	if err != nil {
		diags.AddError(
			"Construction of json payload failed",
			fmt.Sprintf("Error: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}
	return jsonPayload
}

func getAndSetResourceClusterAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *ClusterResourceModel) {
	responseData := client.DoRestRequest(ctx, diags, fmt.Sprintf("%s/%s", clusterPath, data.Id.ValueString()), "GET", nil)
	// When creating or updating the object the username, password, clusterLoginDomain, and multiClusterLoginDomain will be stored in the state file.
	// When importing the object the username, password, clusterLoginDomain, and multiClusterLoginDomain will be set to empty strings in the state file.
	// The API does not return the username, password, and login_domain attributes.
	// Therefore, these attributes will be assigned based on the user's configuration settings.
	*data = *getBaseClusterResourceModel(data.ClusterUsername, data.ClusterPassword, data.ClusterLoginDomain, data.MultiClusterLoginDomain)
	if diags.HasError() {
		return
	}

	if responseData.Data() != nil {
		specReadInfo := responseData.Data().(map[string]interface{})["spec"].(map[string]interface{})

		for attributeName, attributeValue := range specReadInfo {
			if attributeName == "name" {
				data.FabricName = basetypes.NewStringValue(attributeValue.(string))
				data.Id = types.StringValue(data.FabricName.ValueString())
			}

			if attributeName == "clusterType" {
				data.ClusterType = basetypes.NewStringValue(strings.ToLower(attributeValue.(string)))
			}
			if attributeName == "onboardUrl" {
				data.ClusterHostname = basetypes.NewStringValue(attributeValue.(string))
			}

			if attributeName == "location" {
				locationMap := attributeValue.(map[string]interface{})
				if locationMap["latitude"] != nil {
					data.Latitude = basetypes.NewFloat64Value(locationMap["latitude"].(float64))
				}
				if locationMap["longitude"] != nil {
					data.Longitude = basetypes.NewFloat64Value(locationMap["longitude"].(float64))
				}
			}

			if attributeName == "aci" && specReadInfo["clusterType"] == "APIC" {
				aciValueMap := attributeValue.(map[string]interface{})
				data.LicenseTier = basetypes.NewStringValue(aciValueMap["licenseTier"].(string))
				data.SecurityDomain = basetypes.NewStringValue(aciValueMap["securityDomain"].(string))

				if aciValueMap["verifyCA"] != nil {
					data.ValidatePeerCertificate = basetypes.NewBoolValue(aciValueMap["verifyCA"].(bool))
				}

				telemetryValueMap := aciValueMap["telemetry"].(map[string]interface{})
				if telemetryValueMap["network"] != nil {
					data.TelemetryNetwork = basetypes.NewStringValue(telemetryValueMap["network"].(string))
				}
				epgDn := telemetryValueMap["epg"].(string)
				epgName := ""
				if epgDn != "" {
					epgSeparator := "/inb-"
					lastIndex := strings.LastIndex(epgDn, epgSeparator)
					if lastIndex != -1 {
						epgNameStartIndex := lastIndex + len(epgSeparator)
						epgName = epgDn[epgNameStartIndex:]
					}
				}
				data.InbandEpg = basetypes.NewStringValue(epgName)

				featuresList := []string{}
				if telemetryValueMap["status"].(string) != "" && telemetryValueMap["status"].(string) == "enabled" {
					featuresList = append(featuresList, "telemetry")
				}
				orchestrationValueMap := aciValueMap["orchestration"].(map[string]interface{})
				if orchestrationValueMap["status"].(string) != "" && orchestrationValueMap["status"].(string) == "enabled" {
					featuresList = append(featuresList, "orchestration")
				}
				featuresSet, _ := types.SetValueFrom(ctx, basetypes.StringType{}, featuresList)
				data.Features = featuresSet
			}
		}
	} else {
		data.Id = basetypes.NewStringNull()
	}
}
