package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-nd/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ClusterDataSource{}

func NewClusterDataSource() datasource.DataSource {
	return &ClusterDataSource{}
}

// ClusterDataSource defines the data source implementation.
type ClusterDataSource struct {
	client *client.Client
}

// ClusterDataModel describes the resource data model.
type ClusterDataModel struct {
	Id                         types.String  `tfsdk:"id"`
	ClusterType                types.String  `tfsdk:"cluster_type"`
	ClusterHostname            types.String  `tfsdk:"cluster_hostname"`
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

func (d *ClusterDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of datasource: nd_multi_cluster_connectivity")
	resp.TypeName = req.ProviderTypeName + "_multi_cluster_connectivity"
	tflog.Debug(ctx, "End metadata of datasource: nd_multi_cluster_connectivity")
}

func (d *ClusterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of datasource: nd_multi_cluster_connectivity")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Data source for Multi-cluster connectivity for Nexus Dashboard",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The ID of the cluster.",
			},
			"cluster_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The type of the cluster.",
			},
			"cluster_hostname": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The URL or Hostname of the cluster.",
			},
			"fabric_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the cluster.",
			},
			"license_tier": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The license tier of the cluster.",
			},
			"features": schema.SetAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "The features of the cluster.",
			},
			"inband_epg": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The Inband EPG name of the cluster.",
			},
			"security_domain": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The security domain of the cluster.",
			},
			"validate_peer_certificate": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "The validate peer certificate flag of the cluster.",
			},
			"latitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "The latitude coordinate of the cluster.",
			},
			"longitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "The longitude coordinate of the cluster.",
			},
			"telemetry_streaming_protocol": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The telemetry streaming protocol of the cluster.",
			},
			"telemetry_network": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The telemetry network type of the cluster.",
			},
		},
	}
	tflog.Debug(ctx, "End schema of datasource: nd_multi_cluster_connectivity")
}

func (d *ClusterDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of datasource: nd_multi_cluster_connectivity")
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
	tflog.Debug(ctx, "End configure of datasource: nd_multi_cluster_connectivity")
}

func (d *ClusterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Start read of datasource: nd_multi_cluster_connectivity")
	var data *ClusterDataModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = types.StringValue(data.FabricName.ValueString())
	tflog.Debug(ctx, fmt.Sprintf("Read of datasource nd_multi_cluster_connectivity with id '%s'", data.Id.ValueString()))

	getAndSetDataClusterAttributes(ctx, &resp.Diagnostics, d.client, data)
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Failed to read nd_multi_cluster_connectivity data source", "")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End read of datasource nd_multi_cluster_connectivity with id '%s'", data.Id.ValueString()))
}

func getAndSetDataClusterAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *ClusterDataModel) {
	responseData := client.DoRestRequest(ctx, diags, fmt.Sprintf("%s/%s", clusterPath, data.Id.ValueString()), "GET", nil)

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
