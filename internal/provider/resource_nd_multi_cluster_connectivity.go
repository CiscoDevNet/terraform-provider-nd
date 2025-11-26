package provider

import (
	"context"
	"encoding/json"
	"fmt"
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
}

func getBaseClusterResourceModel(username, password, clusterLoginDomain, multiClusterLoginDomain string) *ClusterResourceModel {
	return &ClusterResourceModel{
		Id:                         basetypes.NewStringNull(),
		ClusterType:                basetypes.NewStringNull(),
		ClusterHostname:            basetypes.NewStringNull(),
		ClusterUsername:            basetypes.NewStringValue(username),
		ClusterPassword:            basetypes.NewStringValue(password),
		ClusterLoginDomain:         basetypes.NewStringValue(clusterLoginDomain),
		MultiClusterLoginDomain:    basetypes.NewStringValue(multiClusterLoginDomain),
		FabricName:                 basetypes.NewStringNull(),
		LicenseTier:                basetypes.NewStringNull(),
		Features:                   basetypes.NewSetNull(types.StringType),
		InbandEpg:                  basetypes.NewStringNull(),
		SecurityDomain:             basetypes.NewStringNull(),
		ValidatePeerCertificate:    basetypes.NewBoolNull(),
		Latitude:                   basetypes.NewFloat64Null(),
		Longitude:                  basetypes.NewFloat64Null(),
		TelemetryStreamingProtocol: basetypes.NewStringNull(),
	}
}

func (r *ClusterResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if !req.Plan.Raw.IsNull() {
		var planData, stateData, configData *ClusterResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)
		resp.Diagnostics.Append(req.Config.Get(ctx, &configData)...)

		if stateData != nil {
			if resp.Diagnostics.HasError() {
				return
			}

			if !configData.ClusterUsername.IsNull() && stateData.ClusterUsername.ValueString() == "" {
				planData.ClusterUsername = basetypes.NewStringValue("")
			}

			if !configData.ClusterPassword.IsNull() && stateData.ClusterPassword.ValueString() == "" {
				planData.ClusterPassword = basetypes.NewStringValue("")
			}

			if !configData.ClusterLoginDomain.IsNull() && stateData.ClusterLoginDomain.ValueString() == "" {
				planData.ClusterLoginDomain = basetypes.NewStringValue("")
			}

			if !configData.MultiClusterLoginDomain.IsNull() && stateData.MultiClusterLoginDomain.ValueString() == "" {
				planData.MultiClusterLoginDomain = basetypes.NewStringValue("")
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
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cluster_type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The type of the cluster. Allowed values are 'nd', or 'apic'.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("nd", "apic"),
				},
			},
			"cluster_hostname": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The URL or Hostname of the cluster.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cluster_username": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The username of the cluster.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cluster_password": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The password of the cluster.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cluster_login_domain": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The login domain of the cluster.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"multi_cluster_login_domain": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The multi cluster login domain of the cluster.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"fabric_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the cluster.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"license_tier": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The clusters license tier. Only one value can be specified at a time. Allowed values are 'advantage', or 'essentials', or 'premier'.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("advantage", "essentials", "premier"),
				},
			},
			"features": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "The features of the cluster. Allowed values are 'telemetry', 'orchestration'.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf("telemetry", "orchestration"),
					),
				},
			},
			"inband_epg": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The Inband EPG name of the cluster.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"security_domain": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The security domain of the cluster.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"validate_peer_certificate": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The validate peer certificate flag of the cluster.",
				PlanModifiers:       []planmodifier.Bool{},
			},
			"latitude": schema.Float64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The latitude location of the cluster.",
			},
			"longitude": schema.Float64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The longitude location of the cluster.",
			},
			"telemetry_streaming_protocol": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The telemetry streaming protocol of the cluster. Only one value can be specified at a time. Allowed values are 'ipv4', or 'ipv6'.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("ipv4", "ipv6"),
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

	var data *ClusterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	jsonPayload := getClusterCreateJsonPayload(ctx, &resp.Diagnostics, data, "POST")

	if resp.Diagnostics.HasError() {
		return
	}

	r.client.DoRestRequest(ctx, &resp.Diagnostics, clusterPath, "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	setClusterId(ctx, data, data.FabricName.ValueString())
	getAndSetClusterAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource nd_multi_cluster_connectivity with id '%s'", data.Id.ValueString()))
}

func (r *ClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: nd_multi_cluster_connectivity")
	var data *ClusterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource nd_multi_cluster_connectivity with id '%s'", data.Id.ValueString()))

	getAndSetClusterAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *ClusterResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource nd_multi_cluster_connectivity with id '%s'", data.Id.ValueString()))
}

func (r *ClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: nd_multi_cluster_connectivity")

	var stateData *ClusterResourceModel
	var data *ClusterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource nd_multi_cluster_connectivity with id '%s'", data.Id.ValueString()))

	jsonPayload := getClusterCreateJsonPayload(ctx, &resp.Diagnostics, data, "PUT")

	if resp.Diagnostics.HasError() {
		return
	}

	r.client.DoRestRequest(ctx, &resp.Diagnostics, fmt.Sprintf("%s/%s", clusterPath, data.Id.ValueString()), "PUT", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	setClusterId(ctx, data, data.Id.ValueString())
	getAndSetClusterAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, "End update of resource nd_multi_cluster_connectivity")
}

func (r *ClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: nd_multi_cluster_connectivity")
	var data *ClusterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource nd_multi_cluster_connectivity with id '%s'", data.Id.ValueString()))
	if resp.Diagnostics.HasError() {
		return
	}

	deletePayload := map[string]interface{}{}
	if data.ClusterType.ValueString() == "apic" {
		deletePayload["force"] = true
		deletePayload["credentials"] = map[string]string{
			"loginDomain": data.ClusterLoginDomain.ValueString(),
			"password":    data.ClusterPassword.ValueString(),
			"user":        data.ClusterUsername.ValueString(),
		}

	}
	payload, _ := json.Marshal(deletePayload)
	jsonPayload, _ := gabs.ParseJSON(payload)
	r.client.DoRestRequest(ctx, &resp.Diagnostics, fmt.Sprintf("%s/%s/remove", clusterPath, data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	// Detaching the cluster from the controller may take a few seconds.
	time.Sleep(5 * time.Second)
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource nd_multi_cluster_connectivity with id '%s'", data.Id.ValueString()))
}

func (r *ClusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: nd_multi_cluster_connectivity")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *ClusterResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)

	tflog.Debug(ctx, fmt.Sprintf("Import state of resource nd_multi_cluster_connectivity with id '%s'", stateData.Id.ValueString()))
	tflog.Debug(ctx, "End import of state resource: nd_multi_cluster_connectivity")
}

func getClusterCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, data *ClusterResourceModel, method string) *gabs.Container {
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
		clusterType = "APIC"
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

		orchestrationStatus := "disabled"
		telemetryStatus := "disabled"
		if !data.Features.IsNull() && !data.Features.IsUnknown() {
			featureStrings := make([]string, 0)
			data.Features.ElementsAs(ctx, &featureStrings, false)
			for _, feature := range featureStrings {
				switch feature {
				case "telemetry":
					telemetryStatus = "enabled"
				case "orchestration":
					orchestrationStatus = "enabled"
				}
			}
		}

		aciMap["orchestration"] = map[string]interface{}{
			"status": orchestrationStatus,
		}

		epgDn := ""
		telemetryNetworkType := "outband"
		if !data.InbandEpg.IsNull() && !data.InbandEpg.IsUnknown() && data.InbandEpg.ValueString() != "" {
			epgDn = fmt.Sprintf("uni/tn-mgmt/mgmtp-default/inb-%s", data.InbandEpg.ValueString())
			telemetryNetworkType = "inband"
		}

		telemetryStreamingProtocol := "ipv4"
		if !data.TelemetryStreamingProtocol.IsNull() && !data.TelemetryStreamingProtocol.IsUnknown() {
			telemetryStreamingProtocol = data.TelemetryStreamingProtocol.ValueString()
		}

		aciMap["telemetry"] = map[string]interface{}{
			"status":            telemetryStatus,
			"network":           telemetryNetworkType,
			"streamingProtocol": telemetryStreamingProtocol,
			"epg":               epgDn,
		}
		payloadMap["aci"] = aciMap
	} else if clusterType == "nd" {
		clusterType = "ND"
		if !data.MultiClusterLoginDomain.IsNull() && !data.MultiClusterLoginDomain.IsUnknown() {
			payloadMap["nd"] = map[string]interface{}{"multiClusterLoginDomainName": data.MultiClusterLoginDomain.ValueString()}
		}
	}

	payloadMap["clusterType"] = clusterType

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
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	jsonPayload, err := gabs.ParseJSON(payload)

	if err != nil {
		diags.AddError(
			"Construction of json payload failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}
	return jsonPayload
}

func setClusterId(ctx context.Context, data *ClusterResourceModel, fabricName string) {
	data.Id = types.StringValue(fabricName)
}

func getAndSetClusterAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *ClusterResourceModel) {
	responseData := client.DoRestRequest(ctx, diags, fmt.Sprintf("%s/%s", clusterPath, data.Id.ValueString()), "GET", nil)
	// The API does not return the username, password, and login_domain attributes.
	// Therefore, these attributes will be assigned based on the user's configuration settings.
	*data = *getBaseClusterResourceModel(data.ClusterUsername.ValueString(), data.ClusterPassword.ValueString(), data.ClusterLoginDomain.ValueString(), data.MultiClusterLoginDomain.ValueString())
	if diags.HasError() {
		return
	}

	if responseData.Data() != nil {
		responseReadInfo := responseData.Data().(map[string]interface{})
		specReadInfo := responseReadInfo["spec"].(map[string]interface{})

		for attributeName, attributeValue := range specReadInfo {
			if attributeName == "name" {
				data.FabricName = basetypes.NewStringValue(attributeValue.(string))
				setClusterId(ctx, data, attributeValue.(string))
			}

			if attributeName == "clusterType" {
				if attributeValue.(string) == "APIC" {
					data.ClusterType = basetypes.NewStringValue("apic")
				} else {
					data.ClusterType = basetypes.NewStringValue("nd")
				}
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
