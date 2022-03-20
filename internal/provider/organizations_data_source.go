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

		//TODO Figure out how to make this an array if needed
		Attributes: map[string]tfsdk.Attribute{
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
						Required: true,
					},
				}),
			},
			"licensing": {
				Optional: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"model": {
						Type:     types.StringType,
						Required: true,
					},
				}),
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

type OrganizationsDataSource struct {
	provider provider
}

func (d OrganizationsDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var data Organizations

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := organizations.NewGetOrganizationsParams()

	orgs, err := d.provider.client.Organizations.GetOrganizations(params, d.provider.apiKeyHeaderAuth)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading meraki organization",
			"Could not complete read Organization request for org: "+err.Error(),
		)
		return
	}

	var result []Organization
	for _, org := range orgs.Payload {
		result = append(result, Organization{
			ID:        types.String{Value: org.ID},
			Name:      types.String{Value: org.Name},
			Url:       types.String{Value: org.URL},
			Cloud:     types.String{Value: org.Cloud.Region.Name},
			Api:       Api{Enabled: types.Bool{Value: org.API.Enabled}},
			Licensing: Licensing{Model: types.String{Value: org.Licensing.Model}},
		})
	}

	diags = resp.State.Set(ctx, &result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Error setting state",
			"Could not set state, unexpected error: "+err.Error(),
		)
		return
	}
}
