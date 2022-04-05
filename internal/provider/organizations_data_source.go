package provider

import (
	"context"
	"github.com/ddexterpark/dashboard-api-golang/client/organizations"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OrganizationsDataSourceType struct{}

func (t OrganizationsDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Organizations data source",
		Attributes: map[string]tfsdk.Attribute{
			"organizations": {
				MarkdownDescription: "List of organizations",
				Optional:            true,
				Attributes: tfsdk.SetNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						MarkdownDescription: "merakiOrganization id",
						Optional:            true,
						Type:                types.StringType,
					},
					"name": {
						MarkdownDescription: "merakiOrganization name",
						Optional:            true,
						Type:                types.StringType,
					},
					"url": {
						MarkdownDescription: "merakiOrganization url",
						Optional:            true,
						Type:                types.StringType,
					},
					"cloud": {
						MarkdownDescription: "merakiOrganization region",
						Optional:            true,
						Type:                types.StringType,
					},
					"api": {
						Optional: true,
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"enabled": {
								Type:     types.BoolType,
								Optional: true,
							},
						}),
					},
					"licensing": {
						Optional: true,
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"model": {
								Type:     types.StringType,
								Optional: true,
							},
						}),
					},
				}, tfsdk.SetNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (t OrganizationsDataSourceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return OrganizationsDataSource{
		provider: provider,
	}, diags
}

// OrganizationDataSourceData -
type OrganizationDataSourceData struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Url       types.String `tfsdk:"url"`
	Cloud     types.String `tfsdk:"cloud"`
	Api       Api          `tfsdk:"api"`
	Licensing Licensing    `tfsdk:"licensing"`
}

// OrganizationsDataSourceData -
type OrganizationsDataSourceData struct {
	Organizations []OrganizationDataSourceData `tfsdk:"organizations"`
}

type OrganizationsDataSource struct {
	provider provider
}

func (d OrganizationsDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var data OrganizationsDataSourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := organizations.NewGetOrganizationsParams()
	response, err := d.provider.client.Organizations.GetOrganizations(params, d.provider.transport.DefaultAuthentication)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading meraki organization",
			"Could not complete read Organization request: "+err.Error(),
		)
		return
	}

	// Map response body to resource schema attribute
	for _, organization := range response.Payload {
		data.Organizations = append(data.Organizations, OrganizationDataSourceData{
			ID:        types.String{Value: organization.ID},
			Name:      types.String{Value: organization.Name},
			Url:       types.String{Value: organization.URL},
			Cloud:     types.String{Value: organization.Cloud.Region.Name},
			Api:       Api{Enabled: types.Bool{Value: organization.API.Enabled}},
			Licensing: Licensing{Model: types.String{Value: organization.Licensing.Model}},
		})
	}

	// Map response body to resource schema attribute
	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}
