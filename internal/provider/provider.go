package provider

import (
	"context"
	"fmt"
	apiclient "github.com/ddexterpark/dashboard-api-golang/client"
	httptransport "github.com/go-openapi/runtime/client"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"os"
)

func New(version string) func() tfsdk.Provider {
	return func() tfsdk.Provider {
		return &Provider{
			Version: version,
		}
	}
}

// Provider satisfies the tfsdk.Provider interface and usually is included
// with all Resource and DataSource implementations.
type Provider struct {
	// client can contain the upstream provider SDK or HTTP client used to
	// communicate with the upstream service. Resource and DataSource
	// implementations can then make calls using this client.
	// meraki-golang-api client
	Client    *apiclient.MerakiAPIGolang
	Transport *httptransport.Runtime

	// Configured is set to true at the end of the Configure method.
	// This can be used in Resource and DataSource implementations to verify
	// that the provider was previously configured.
	Configured bool

	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	Version string
}

func (p *Provider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"path": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
			"host": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
			"apikey": {
				Type:      types.StringType,
				Required:  true,
				Sensitive: true,
			},
		},
	}, nil
}

// Provider schema struct
type providerData struct {
	Host   types.String `tfsdk:"host"`
	Path   types.String `tfsdk:"path"`
	ApiKey types.String `tfsdk:"apikey"`
}

func (p *Provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	// Retrieve provider data from configuration
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a user to the provider
	var host string
	if config.Host.Null {
		host = "api.meraki.com"
	} else {
		host = config.Host.Value
	}

	if host == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find host",
			"host cannot be an empty string",
		)
		return
	}

	// User must provide a user to the provider
	var path string
	if config.Path.Null {
		path = "/api/v1"
	} else {
		path = config.Path.Value
	}

	if path == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find path",
			"path cannot be an empty string",
		)
		return
	}

	// User must provide a apikey to the provider
	var apikey string
	if config.ApiKey.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Cannot use unknown value as apikey",
		)
		return
	}

	if config.ApiKey.Null {
		apikey = os.Getenv("MERAKI_DASHBOARD_API_KEY")
	} else {
		apikey = config.ApiKey.Value
	}

	if apikey == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find apikey",
			"apikey cannot be an empty string",
		)
		return
	}

	// Create a new Meraki Dashboard API client and set it to the provider client
	p.Transport = httptransport.New(host, path, []string{"https"})
	p.Transport.DefaultAuthentication = httptransport.APIKeyAuth("X-Cisco-Meraki-API-Key", "header", apikey)
	p.Client = apiclient.New(p.Transport, nil)
	p.Configured = true
}

// GetResources - Defines provider resources
func (p *Provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"meraki_organization": OrganizationResourceType{},
	}, nil
}

// GetDataSources - Defines provider data sources
func (p *Provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"meraki_organizations": OrganizationsDataSourceType{},
	}, nil
}

// convertProviderType is a helper function for NewResource and NewDataSource
// implementations to associate the concrete provider type. Alternatively,
// this helper can be skipped and the provider type can be directly type
// asserted (e.g. provider: in.(*provider)), however using this can prevent
// potential panics.
func ConvertProviderType(in tfsdk.Provider) (Provider, diag.Diagnostics) {
	var diags diag.Diagnostics

	p, ok := in.(*Provider)

	if !ok {
		diags.AddError(
			"Unexpected Provider Instance Type",
			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
		)
		return Provider{}, diags
	}

	if p == nil {
		diags.AddError(
			"Unexpected Provider Instance Type",
			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
		)
		return Provider{}, diags
	}

	return *p, diags

}
