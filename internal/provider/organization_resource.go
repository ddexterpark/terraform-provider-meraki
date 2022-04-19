package provider

import (
	"context"
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/client/organizations"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type OrganizationResourceType struct{}
type OrganizationData struct {

func (t OrganizationResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"name": {
				Type:     types.StringType,
				Computed: true,
				Optional: true,
			},
			"url": {
				Type:     types.StringType,
				Computed: true,
			},
			"cloud": {
				Type:     types.StringType,
				Computed: true,
			},
			"api": {
				Computed: true,
				Optional: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"enabled": {
						Type:     types.BoolType,
						Required: true,
					},
				}),
			},
			"licensing": {
				Computed: true,
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

type OrganizationData struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Url       types.String `tfsdk:"url"`
	Cloud     types.String `tfsdk:"cloud"`
	Api       *Api         `tfsdk:"api"`
	Licensing *Licensing   `tfsdk:"licensing"`
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

func (r merakiOrganizationResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {

	if !r.provider.configured {
	if !r.provider.Configured {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}
	var data OrganizationData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Create Organization Error",
			fmt.Sprintf("data: %v", data),
		)
		return
	}

	// Retrieve values from plan
	params := organizations.NewCreateOrganizationParams()
	params.CreateOrganization.Name = &data.Name.Value
	response, err := r.provider.client.Organizations.CreateOrganization(params, r.provider.transport.DefaultAuthentication)
	if err != nil {
		resp.Diagnostics.AddError(
			"Create Organization Error",
			fmt.Sprintf("Config from .tf file: %v", data),
		)
		resp.Diagnostics.AddError(
			"Create Organization Error",
			fmt.Sprintf("Meraki API Response: %s", err),
		)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("api response: %s", response))

	// Map response body to resource schema attribute
	data.ID = types.String{Value: response.GetPayload().ID}
	data.Name = types.String{Value: response.GetPayload().Name}
	data.Url = types.String{Value: response.GetPayload().URL}
	data.Cloud = types.String{Value: response.GetPayload().Cloud.Region.Name}

	if licensing := response.GetPayload().Licensing; licensing != nil {
		data.Licensing = &Licensing{
			Model: types.String{Value: licensing.Model},
		}
	} else {
		data.Licensing = &Licensing{Model: types.String{Null: true}}
	}

	if api := response.GetPayload().API; api != nil {
		data.Api = &Api{
			Enabled: types.Bool{Value: api.Enabled},
		}
	} else {
		data.Api = &Api{Enabled: types.Bool{Null: true}}
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r merakiOrganizationResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var data OrganizationData

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
			fmt.Sprintf("Failure: %s", err),
		)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("api response: %s", response))

	// Map response body to resource schema attribute
	data.ID = types.String{Value: response.GetPayload().ID}
	data.Name = types.String{Value: response.GetPayload().Name}
	data.Url = types.String{Value: response.GetPayload().URL}
	data.Cloud = types.String{Value: response.GetPayload().Cloud.Region.Name}

	if licensing := response.GetPayload().Licensing; licensing != nil {
		data.Licensing = &Licensing{
			Model: types.String{Value: licensing.Model},
		}
	} else {
		data.Licensing = &Licensing{Model: types.String{Null: true}}
	}

	if api := response.GetPayload().API; api != nil {
		data.Api = &Api{
			Enabled: types.Bool{Value: api.Enabled},
		}
	} else {
		data.Api = &Api{Enabled: types.Bool{Null: true}}
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r merakiOrganizationResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var data OrganizationData

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Update Organization Error",
			fmt.Sprintf("data: %v", data),
		)
		return
	}

	// Update Organization
	params := organizations.NewUpdateOrganizationParams()
	params.OrganizationID = data.ID.Value
	params.UpdateOrganization.Name = data.Name.Value

	var enabled = organizations.UpdateOrganizationParamsBodyAPI{
		Enabled: data.Api.Enabled.Value,
	}
	params.UpdateOrganization.API = &enabled

	if err != nil {
		resp.Diagnostics.AddError(
			"Update Organization Error",
			fmt.Sprintf("Failed Update Organization Response: %s", err),
		)
		return
	}

	// Map response body to resource schema attribute
	data.ID = types.String{Value: response.GetPayload().ID}
	data.Name = types.String{Value: response.GetPayload().Name}
	data.Url = types.String{Value: response.GetPayload().URL}
	data.Cloud = types.String{Value: response.GetPayload().Cloud.Region.Name}

	if licensing := response.GetPayload().Licensing; licensing != nil {
		data.Licensing = &Licensing{
			Model: types.String{Value: licensing.Model},
		}
	} else {
		data.Licensing = &Licensing{Model: types.String{Null: true}}
	}

	if api := response.GetPayload().API; api != nil {
		data.Api = &Api{
			Enabled: types.Bool{Value: api.Enabled},
		}
	} else {
		data.Api = &Api{Enabled: types.Bool{Null: true}}
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r merakiOrganizationResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var data OrganizationData

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
