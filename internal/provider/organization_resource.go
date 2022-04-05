package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/client/organizations"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type OrganizationResourceType struct{}

func (t OrganizationResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	// TODO - rework the schema based on learnings from this article(also check how hashmaps are implimented in the tf docs and scaffolding repos..):
	// https://medium.com/@steve_strutt/developing-a-public-terraform-provider-part-4-resource-schema-fb3cc93954e0
	// https://github.com/IBM-Cloud/terraform-provider-ibm/blob/master/ibm/service/cis/resource_ibm_cis_origin_pool.go
	return tfsdk.Schema{
		MarkdownDescription: "Organization resource",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				MarkdownDescription: "merakiOrganization id",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
			},
			"name": {
				MarkdownDescription: "merakiOrganization name",
				Required:            true,
				Type:                types.StringType,
			},
			"url": {
				MarkdownDescription: "merakiOrganization url",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
			},
			"cloud": {
				MarkdownDescription: "merakiOrganization region",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
			},
			"api": {
				Type:     types.MapType{ElemType: types.BoolType},
				Optional: true,
			},
			"licensing": {
				Type:     types.MapType{ElemType: types.StringType},
				Optional: true,
			},
		},
	}, nil
}

func (t OrganizationResourceType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return merakiOrganizationResource{
		provider: provider,
	}, diags
}

type merakiOrganizationData struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Url       types.String `tfsdk:"url"`
	Cloud     types.String `tfsdk:"cloud"`
	API       Api          `tfsdk:"api"`
	Licensing Licensing    `tfsdk:"licensing"`
}

type Api struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type Licensing struct {
	Model types.String `tfsdk:"model"`
}

type merakiOrganizationResource struct {
	provider provider
}

func valueConversion(data interface{}, results interface{}) {
	jsonString, _ := json.Marshal(data)

	// convert json to struct
	err := json.Unmarshal(jsonString, &results)
	if err != nil {
		return
	}
}

func (r merakiOrganizationResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var data merakiOrganizationData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve values from plan
	params := organizations.NewCreateOrganizationParams()
	params.CreateOrganization.Name = &data.Name.Value
	response, err := r.provider.client.Organizations.CreateOrganization(params, r.provider.transport.DefaultAuthentication)
	if err != nil {
		resp.Diagnostics.AddError(
			"Create Organization Error",
			fmt.Sprintf("Failed Create Organization Response: %s", err),
		)
		return
	}

	// Map response body to resource schema attribute
	valueConversion(response.GetPayload(), &data)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r merakiOrganizationResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var data merakiOrganizationData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := organizations.NewGetOrganizationParams()
	params.OrganizationID = data.ID.Value

	response, err := r.provider.client.Organizations.GetOrganization(params, r.provider.transport.DefaultAuthentication)
	if err != nil {
		resp.Diagnostics.AddError(
			"Read Organization Error",
			fmt.Sprintf("Failed Read Organization Response: %s", err),
		)
		return
	}

	// Map response body to resource schema attribute
	data.ID = types.String{Value: response.GetPayload().ID}
	data.Name = types.String{Value: response.GetPayload().Name}
	data.Url = types.String{Value: response.GetPayload().URL}
	data.Cloud = types.String{Value: response.GetPayload().Cloud.Region.Name}

	//TODO - breaking provider
	data.Licensing.Model = types.String{Value: response.GetPayload().Licensing.Model}
	data.API.Enabled = types.Bool{Value: response.GetPayload().API.Enabled}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r merakiOrganizationResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var data merakiOrganizationData

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update Organization
	params := organizations.NewUpdateOrganizationParams()
	params.OrganizationID = data.ID.Value
	params.UpdateOrganization.Name = data.Name.Value

	// TODO - breaking provider
	//params.UpdateOrganization.API.Enabled = data.Api.Enabled.Value

	response, err := r.provider.client.Organizations.UpdateOrganization(params, r.provider.transport.DefaultAuthentication)
	if err != nil {
		resp.Diagnostics.AddError(
			"Update Organization Error",
			fmt.Sprintf("Failed Update Organization Response: %s", err),
		)
		return
	}

	// Map response body to resource schema attribute
	valueConversion(response.GetPayload(), &data)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r merakiOrganizationResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var data merakiOrganizationData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete Organization
	params := organizations.NewDeleteOrganizationParams()
	params.OrganizationID = data.ID.Value
	_, err := r.provider.client.Organizations.DeleteOrganization(params, r.provider.transport.DefaultAuthentication)
	if err != nil {
		resp.Diagnostics.AddError(
			"Delete Organization Error",
			fmt.Sprintf("Failed Delete Organization Response: %s", err),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r merakiOrganizationResource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
