package provider

import (
	"context"
	"github.com/ddexterpark/dashboard-api-golang/client/organizations"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type OrganizationResourceType struct{}

func (t OrganizationResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
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

func (t OrganizationResourceType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return merakiOrganizationResource{
		provider: provider,
	}, diags
}

type merakiOrganizationResource struct {
	provider provider
}

func (r merakiOrganizationResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.provider.configured {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, "+
				"likely because it depends on an unknown value from another resource.")
		return
	}

	// Retrieve values from plan
	var plan Organization
	err := req.Plan.Get(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading plan",
			"An unexpected error was encountered while reading the plan")
		return
	}

	params := organizations.CreateOrganizationParams{}
	Name := plan.Name.Value
	params.CreateOrganization.Name = &Name

	data, errr := r.provider.client.Organizations.CreateOrganization(&params, r.provider.apiKeyHeaderAuth)
	if errr != nil {
		resp.Diagnostics.AddError(
			"Error reading meraki organization",
			"Could not complete read Organization request for org: "+Name+": "+errr.Error(),
		)
		return
	}

	var result = organizations.CreateOrganizationCreated{
		Payload: data,
	}

	diags := req.Plan.Get(ctx, &result)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r merakiOrganizationResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state Organization

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := organizations.NewGetOrganizationParams()

	iD := state.ID.Value
	params.OrganizationID = iD

	data, err := r.provider.client.Organizations.GetOrganization(params, r.provider.apiKeyHeaderAuth)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading meraki organization",
			"Could not complete read Organization request for org: "+iD+": "+err.Error(),
		)
		return
	}

	var result []Organization
	result = append(result, Organization{
		ID:        types.String{Value: data.GetPayload().ID},
		Name:      types.String{Value: data.GetPayload().Name},
		Url:       types.String{Value: data.GetPayload().URL},
		Cloud:     types.String{Value: data.GetPayload().Cloud.Region.Name},
		Api:       Api{Enabled: types.Bool{Value: data.GetPayload().API.Enabled}},
		Licensing: Licensing{Model: types.String{Value: data.GetPayload().Licensing.Model}},
	})

	diags = resp.State.Set(ctx, result)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Error setting state",
			"Could not set state, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r merakiOrganizationResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var state Organization

	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	params := organizations.UpdateOrganizationParams{}
	iD := state.ID.Value
	params.OrganizationID = iD
	params.UpdateOrganization.Name = state.Name.Value
	params.UpdateOrganization.API.Enabled = state.Api.Enabled.Value

	data, err := r.provider.client.Organizations.UpdateOrganization(&params, r.provider.apiKeyHeaderAuth)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading meraki organization",
			"Could not complete read Organization request for org: "+iD+": "+err.Error(),
		)
		return
	}

	var result = organizations.UpdateOrganizationOK{
		Payload: data,
	}

	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &result)
	resp.Diagnostics.Append(diags...)
}

func (r merakiOrganizationResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state Organization

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	params := organizations.DeleteOrganizationParams{}
	iD := state.ID.Value
	params.OrganizationID = iD

	_, err := r.provider.client.Organizations.DeleteOrganization(&params, r.provider.apiKeyHeaderAuth)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading meraki organization",
			"Could not complete read Organization request for org: "+iD+": "+err.Error(),
		)
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r merakiOrganizationResource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
