package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-nd/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SiteDataSource{}

func NewSiteDataSource() datasource.DataSource {
	return &SiteDataSource{}
}

// SiteDataSource defines the data source implementation.
type SiteDataSource struct {
	client *client.Client
}

func (d *SiteDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of datasource: nd_site")
	resp.TypeName = req.ProviderTypeName + "_site"
	tflog.Debug(ctx, "End metadata of datasource: nd_site")
}

func (d *SiteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of datasource: nd_site")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Data source for Nexus Dashboard Sites",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The ID of the site.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the site.",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The URL of the site.",
			},
			"type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The type of the site.",
			},
			"username": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The username of the site.",
			},
			"password": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The password of the site.",
			},
			"login_domain": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The login domain of the site.",
			},
			"inband_epg": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The In-Band Endpoint Group (EPG) used to connect ND to the site.",
			},
			"latitude": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The latitude location of the site.",
			},
			"longitude": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The longitude location of the site.",
			},
			"use_proxy": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "The use proxy of the site.",
			},
		},
	}
	tflog.Debug(ctx, "End schema of datasource: nd_site")
}

func (d *SiteDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of datasource: nd_site")
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
	tflog.Debug(ctx, "End configure of datasource: nd_site")
}

func (d *SiteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Start read of datasource: nd_site")
	var data *SiteResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setSiteId(ctx, data)

	tflog.Debug(ctx, fmt.Sprintf("Read of datasource nd_site with id '%s'", data.Id.ValueString()))

	getAndSetSiteAttributes(ctx, &resp.Diagnostics, d.client, data)

	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Failed to read nd_site data source", "")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End read of datasource nd_site with id '%s'", data.Id.ValueString()))
}
