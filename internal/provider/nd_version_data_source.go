package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &VersionDataSource{}

func NewVersionDataSource() datasource.DataSource {
	return &VersionDataSource{}
}

// VersionDataSource defines the data source implementation.
type VersionDataSource struct {
	client *Client
}

func (d *VersionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of datasource: nd_version")
	resp.TypeName = req.ProviderTypeName + "_version"
	tflog.Debug(ctx, "End metadata of datasource: nd_version")
}

type VersionResourceModel struct {
	Id          types.String  `tfsdk:"commit_id"`
	BuildTime   types.String  `tfsdk:"build_time"`
	BuildHost   types.String  `tfsdk:"build_host"`
	User        types.String  `tfsdk:"user"`
	ProductId   types.String  `tfsdk:"product_id"`
	ProductName types.String  `tfsdk:"product_name"`
	Release     types.Bool    `tfsdk:"release"`
	Major       types.Float64 `tfsdk:"major"`
	Minor       types.Float64 `tfsdk:"minor"`
	Maintenance types.Float64 `tfsdk:"maintenance"`
	Patch       types.String  `tfsdk:"patch"`
}

func (d *VersionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of datasource: nd_version")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The version datasource for the 'ND Platform Version' information",

		Attributes: map[string]schema.Attribute{
			"commit_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The commit id of the ND Platform Version.",
			},
			"build_time": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The build time of the ND Platform Version.",
			},
			"build_host": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The build host of the ND Platform Version.",
			},
			"user": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The build user name of the ND Platform Version.",
			},
			"product_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The product id of the ND Platform Version.",
			},
			"product_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The product name of the ND Platform Version.",
			},
			"release": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "The release status of the ND Platform Version.",
			},
			"major": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "The major version number of the ND Platform Version.",
			},
			"minor": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "The minor version number of the ND Platform Version.",
			},
			"maintenance": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "The maintenance version number of the ND Platform Version.",
			},
			"patch": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The patch version letter of the ND Platform Version.",
			},
		},
	}
	tflog.Debug(ctx, "End schema of datasource: nd_version")
}

func (d *VersionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of datasource: nd_version")
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
	tflog.Debug(ctx, "End configure of datasource: nd_version")
}

func (d *VersionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Start read of datasource: nd_version")
	var data *VersionResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setVersionId(ctx, data)

	tflog.Debug(ctx, fmt.Sprintf("Read of datasource nd_version with id '%s'", data.Id.ValueString()))

	getAndSetVersionAttributes(ctx, &resp.Diagnostics, d.client, data)

	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Failed to read nd_version data source", "")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End read of datasource nd_version with id '%s'", data.Id.ValueString()))
}

func setVersionId(ctx context.Context, data *VersionResourceModel) {
	data.Id = types.StringValue(data.Id.ValueString())
}

func getAndSetVersionAttributes(ctx context.Context, diags *diag.Diagnostics, client *Client, data *VersionResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, "version.json", "GET", nil)
	if diags.HasError() {
		return
	}
	if requestData.Data() != nil {
		classReadInfo := requestData.Data().(map[string]interface{})
		for attributeName, attributeValue := range classReadInfo {

			if attributeName == "commit_id" {
				data.Id = basetypes.NewStringValue(attributeValue.(string))
			}

			if attributeName == "build_time" {
				data.BuildTime = basetypes.NewStringValue(attributeValue.(string))
			}

			if attributeName == "build_host" {
				data.BuildHost = basetypes.NewStringValue(attributeValue.(string))
			}

			if attributeName == "user" {
				data.User = basetypes.NewStringValue(attributeValue.(string))
			}

			if attributeName == "product_id" {
				data.ProductId = basetypes.NewStringValue(attributeValue.(string))
			}

			if attributeName == "product_name" {
				data.ProductName = basetypes.NewStringValue(attributeValue.(string))
			}

			if attributeName == "release" {
				data.Release = basetypes.NewBoolValue(attributeValue.(bool))
			}

			if attributeName == "major" {
				data.Major = basetypes.NewFloat64Value(attributeValue.(float64))
			}

			if attributeName == "minor" {
				data.Minor = basetypes.NewFloat64Value(attributeValue.(float64))
			}

			if attributeName == "maintenance" {
				data.Maintenance = basetypes.NewFloat64Value(attributeValue.(float64))
			}

			if attributeName == "patch" {
				data.Patch = basetypes.NewStringValue(attributeValue.(string))
			}
		}
	} else {
		data.Id = basetypes.NewStringNull()
	}
}
