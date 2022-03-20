package provider

import (
	"context"
	"fmt"
	apiclient "github.com/ddexterpark/dashboard-api-golang/client"
	"github.com/go-openapi/runtime"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"os"
)

var stderr = os.Stderr

func New(version string) func() tfsdk.Provider {
	return func() tfsdk.Provider {
		return &provider{
			version: version,
		}
	}
}

type provider struct {
	configured bool

	// meraki-golang-api client
	client    *apiclient.MerakiAPIGolang
	transport *httptransport.Runtime

	// Header Based Authentication is an input passed to client API calls
	apiKeyHeaderAuth runtime.ClientAuthInfoWriter
	version          string
}

// GetSchema
func (p *provider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
				Optional:  true,
				Computed:  true,
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

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	// Retrieve provider data from configuration
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a user to the provider
	var host string
	if config.Host.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as host",
		)
		return
	}

	if config.Host.Null {
		host = os.Getenv("MERAKI_DASHBOARD_API_URL")
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
	if config.Path.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as path",
		)
		return
	}

	if config.Path.Null {
		path = os.Getenv("MERAKI_DASHBOARD_API_URL")
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
	p.transport = httptransport.New(host, path, []string{"https"})
	p.transport.DefaultAuthentication = httptransport.APIKeyAuth("X-Cisco-Meraki-API-Key", "header", apikey)

	p.client = MerakiApiClient(p.transport)
	p.configured = true
}

func MerakiApiClient(transport *httptransport.Runtime) *apiclient.MerakiAPIGolang {
	merakiApiClient, _ := MerakiApiClientWithErrors(transport)
	return merakiApiClient
}

func MerakiApiClientWithErrors(transport *httptransport.Runtime) (*apiclient.MerakiAPIGolang, error) {

	// perhaps an API call as a health check could be error pass condition?
	client := apiclient.New(transport, nil)

	return client, nil
}

// GetResources - Defines provider resources
func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"meraki_organization": OrganizationResourceType{},
	}, nil
}

// GetDataSources - Defines provider data sources
func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"meraki_organizations": OrganizationsDataSourceType{},
	}, nil
}

func convertProviderType(in tfsdk.Provider) (provider, diag.Diagnostics) {
	var diags diag.Diagnostics

	p, ok := in.(*provider)

	if !ok {
		diags.AddError(
			"Unexpected Provider Instance Type",
			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
		)
		return provider{}, diags
	}

	if p == nil {
		diags.AddError(
			"Unexpected Provider Instance Type",
			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
		)
		return provider{}, diags
	}

	return *p, diags

}
