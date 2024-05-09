package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/container"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SiteResource{}
var _ resource.ResourceWithImportState = &SiteResource{}

var sitePath = "nexus/api/sitemanagement/v4/sites"
var siteTypeMap = map[string]string{
	"ACI":         "aci",
	"DCNM":        "dcnm",
	"ThirdParty":  "third_party",
	"CloudACI":    "cloud_aci",
	"DCNMNG":      "dcnm_ng",
	"NDFC":        "ndfc",
	"aci":         "ACI",
	"dcnm":        "DCNM",
	"third_party": "ThirdParty",
	"cloud_aci":   "CloudACI",
	"dcnm_ng":     "DCNMNG",
	"ndfc":        "NDFC",
}

func NewSiteResource() resource.Resource {
	return &SiteResource{}
}

// SiteResource defines the resource implementation.
type SiteResource struct {
	client *client.Client
}

// SiteResourceModel describes the resource data model.
type SiteResourceModel struct {
	Id           types.String `tfsdk:"site_name"`
	SitePassword types.String `tfsdk:"site_password"`
	SiteUsername types.String `tfsdk:"site_username"`
	LoginDomain  types.String `tfsdk:"login_domain"`
	InbandEpg    types.String `tfsdk:"inband_epg"`
	Url          types.String `tfsdk:"url"`
	SiteType     types.String `tfsdk:"site_type"`
	Latitude     types.String `tfsdk:"latitude"`
	Longitude    types.String `tfsdk:"longitude"`
}

func (r *SiteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: nd_site")
	resp.TypeName = req.ProviderTypeName + "_site"
	tflog.Debug(ctx, "End metadata of resource: nd_site")
}

func (r *SiteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: nd_site")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The site resource for the Nexus Dashboard",

		Attributes: map[string]schema.Attribute{
			"site_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the site.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"url": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The URL to reference the APICs.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"site_type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The site type of the APICs.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("aci", "dcnm", "third_party", "cloud_aci", "dcnm_ng", "ndfc"),
				},
			},
			"site_username": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The username for the APIC.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
			"site_password": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The password for the APIC.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"login_domain": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The AAA login domain for the username of the APIC.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"inband_epg": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The In-Band Endpoint Group (EPG) used to connect Nexus Dashboard to the fabric.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"latitude": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The latitude of the location of the site.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"longitude": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The longitude of the location of the site.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
	tflog.Debug(ctx, "End schema of resource: nd_site")
}

func (r *SiteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: nd_site")
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
	tflog.Debug(ctx, "End configure of resource: nd_site")
}

func (r *SiteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: nd_site")

	var stateData *SiteResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)

	var data *SiteResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	jsonPayload := getSiteCreateJsonPayload(ctx, &resp.Diagnostics, data)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, sitePath, "POST", jsonPayload)

	setSiteId(ctx, data)

	if resp.Diagnostics.HasError() {
		return
	}
	getAndSetSiteAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource nd_site with id '%s'", data.Id.ValueString()))
}

func (r *SiteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: nd_site")
	var data *SiteResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource nd_site with id '%s'", data.Id.ValueString()))

	getAndSetSiteAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *SiteResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource nd_site with id '%s'", data.Id.ValueString()))
}

func (r *SiteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: nd_site")
	// Update function is not supported
	tflog.Debug(ctx, "End update of resource nd_site")
}

func (r *SiteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: nd_site")
	var data *SiteResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource nd_site with id '%s'", data.Id.ValueString()))
	if resp.Diagnostics.HasError() {
		return
	}
	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("%s/%s", sitePath, data.Id.ValueString()), "DELETE", nil)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource nd_site with id '%s'", data.Id.ValueString()))
}

func (r *SiteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: nd_site")
	resource.ImportStatePassthroughID(ctx, path.Root("site_name"), req, resp)

	var stateData *SiteResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)

	username := os.Getenv("ND_SITE_USERNAME")
	if username == "" {
		resp.Diagnostics.AddError("Missing input", "A username must be provided during import, please set the ND_SITE_USERNAME environment variable")
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_username"), username)...)

	password := os.Getenv("ND_SITE_PASSWORD")
	if password == "" {
		resp.Diagnostics.AddError("Missing input", "A password must be provided during import, please set the ND_SITE_PASSWORD environment variable")
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_password"), password)...)

	loginDomain := os.Getenv("ND_LOGIN_DOMAIN")
	if loginDomain == "" {
		resp.Diagnostics.AddError("Missing input", "A login_domain must be provided during import, please set the ND_LOGIN_DOMAIN environment variable")
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("login_domain"), loginDomain)...)

	tflog.Debug(ctx, fmt.Sprintf("Import state of resource nd_site with id '%s'", stateData.Id.ValueString()))
	tflog.Debug(ctx, "End import of state resource: nd_site")
}

func getSiteCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, data *SiteResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	siteType := ""

	if !data.SitePassword.IsNull() && !data.SitePassword.IsUnknown() {
		payloadMap["password"] = data.SitePassword.ValueString()
	}

	if !data.SiteUsername.IsNull() && !data.SiteUsername.IsUnknown() {
		payloadMap["userName"] = data.SiteUsername.ValueString()
	}

	if !data.LoginDomain.IsNull() && !data.LoginDomain.IsUnknown() {
		payloadMap["loginDomain"] = data.LoginDomain.ValueString()
	}

	if !data.InbandEpg.IsNull() && !data.InbandEpg.IsUnknown() {
		payloadMap["inband_epg"] = data.InbandEpg.ValueString()
	}

	if !data.Id.IsNull() && !data.Id.IsUnknown() {
		payloadMap["name"] = data.Id.ValueString()
	}

	if !data.Url.IsNull() && !data.Url.IsUnknown() {
		payloadMap["host"] = data.Url.ValueString()
	}

	if !data.SiteType.IsNull() && !data.SiteType.IsUnknown() {
		payloadMap["siteType"] = data.SiteType.ValueString()
		siteType = data.SiteType.ValueString()
	}

	if !data.Latitude.IsNull() && !data.Latitude.IsUnknown() {
		payloadMap["latitude"] = data.Latitude.ValueString()
	}

	if !data.Longitude.IsNull() && !data.Longitude.IsUnknown() {
		payloadMap["longitude"] = data.Longitude.ValueString()
	}

	siteConfiguration := map[string]interface{}{}
	if siteType == "aci" || siteType == "cloud_aci" {
		siteTypeParam := siteType
		if siteType == "cloud_aci" {
			siteTypeParam = siteTypeMap[siteType]
		}

		inbandEpg := ""
		if payloadMap["inband_epg"] != nil {
			inbandEpg = payloadMap["inband_epg"].(string)
		}
		siteConfiguration[siteTypeParam] = map[string]interface{}{
			"InbandEPGDN": inbandEpg,
		}
	} else if siteType == "ndfc" || siteType == "dcnm" {
		siteConfiguration[siteType] = map[string]string{
			"fabricName":       payloadMap["name"].(string),
			"fabricTechnology": "External",
			"fabricType":       "External",
		}
	}

	payloadMap["siteConfig"] = siteConfiguration
	payloadMap["siteType"] = siteTypeMap[siteType]

	payload, err := json.Marshal(map[string]interface{}{"spec": payloadMap})
	if err != nil {
		diags.AddError(
			"Marshalling of json payload failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	jsonPayload, err := container.ParseJSON(payload)

	if err != nil {
		diags.AddError(
			"Construction of json payload failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}
	return jsonPayload
}

func setSiteId(ctx context.Context, data *SiteResourceModel) {
	data.Id = types.StringValue(data.Id.ValueString())
}

func getAndSetSiteAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *SiteResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, fmt.Sprintf("%s/%s", sitePath, data.Id.ValueString()), "GET", nil)
	if diags.HasError() {
		return
	}
	if requestData.Data() != nil {
		metaReadInfo := requestData.Data().(map[string]interface{})
		specReadInfo := metaReadInfo["spec"].(map[string]interface{})
		for attributeName, attributeValue := range specReadInfo {
			if attributeName == "name" {
				data.Id = basetypes.NewStringValue(attributeValue.(string))
			}

			if attributeName == "siteConfig" {
				data.InbandEpg = basetypes.NewStringValue(attributeValue.(map[string]interface{})[siteTypeMap[specReadInfo["siteType"].(string)]].(map[string]interface{})["InbandEPGDN"].(string))
			}

			if attributeName == "host" {
				data.Url = basetypes.NewStringValue(attributeValue.(string))
			}

			if attributeName == "siteType" {
				data.SiteType = basetypes.NewStringValue(siteTypeMap[attributeValue.(string)])
			}

			if attributeName == "latitude" {
				data.Latitude = basetypes.NewStringValue(attributeValue.(string))
			}

			if attributeName == "longitude" {
				data.Longitude = basetypes.NewStringValue(attributeValue.(string))
			}
		}
	} else {
		data.Id = basetypes.NewStringNull()
	}
}
