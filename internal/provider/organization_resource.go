package provider

import (
	"context"
	"github.com/ddexterpark/dashboard-api-golang/client/organizations"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type merakiOrganizationResourceType struct{}

func (t merakiOrganizationResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"items": {
				Required: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						MarkdownDescription: "merakiOrganization configurable attribute",
						Optional:            true,
						Type:                types.StringType,
					},
					"name": {
						Computed:            true,
						MarkdownDescription: "merakiOrganization identifier",
						PlanModifiers: tfsdk.AttributePlanModifiers{
							tfsdk.UseStateForUnknown(),
						},
						Type: types.StringType,
					},
					"url": {
						MarkdownDescription: "merakiOrganization configurable attribute",
						Optional:            true,
						Type:                types.StringType,
					},
					"api": {
						Required: true,
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"enabled": {
								Type:     types.BoolType,
								Required: true,
							},
						}),
					},
					"licensing": {
						Required: true,
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"model": {
								Type:     types.StringType,
								Required: true,
							},
						}),
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (t merakiOrganizationResourceType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return merakiOrganizationResource{
		provider: provider,
	}, diags
}

type merakiOrganizationResourceData struct {
	Id        types.String  `tfsdk:"id"`
	Name      types.String  `tfsdk:"name"`
	Url       types.String  `tfsdk:"url"`
	Api       types.MapType `tfsdk:"api"`
	Licensing types.MapType `tfsdk:"licensing"`
}

type merakiOrganizationResource struct {
	provider provider
}

func (r merakiOrganizationResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {

	if !r.provider.configured {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	var data []merakiOrganizationResourceData
	diags := req.Plan.Get(ctx, &data)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Trace(ctx, "created a resource")

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r merakiOrganizationResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state merakiOrganizationResourceData

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	iD := state.Id.Value
	params := organizations.NewGetOrganizationParams()

	org, err := r.provider.client.Organizations.GetOrganization(params, r.provider.apiKeyHeaderAuth)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading order",
			"Could not complete read Organization request for: "+iD+": "+err.Error(),
		)
		return
	}

	diags = resp.State.Set(ctx, org.Payload)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r merakiOrganizationResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var data merakiOrganizationResourceData

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r merakiOrganizationResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var data merakiOrganizationResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r merakiOrganizationResource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
