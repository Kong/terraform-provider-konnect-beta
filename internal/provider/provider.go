// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
	"os"
)

var _ provider.Provider = &KonnectBetaProvider{}

type KonnectBetaProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// KonnectBetaProviderModel describes the provider data model.
type KonnectBetaProviderModel struct {
	KonnectAccessToken       types.String `tfsdk:"konnect_access_token"`
	PersonalAccessToken      types.String `tfsdk:"personal_access_token"`
	ServerURL                types.String `tfsdk:"server_url"`
	SystemAccountAccessToken types.String `tfsdk:"system_account_access_token"`
}

func (p *KonnectBetaProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "konnect-beta"
	resp.Version = p.version
}

func (p *KonnectBetaProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"konnect_access_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"personal_access_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"server_url": schema.StringAttribute{
				Description: `Server URL (defaults to https://global.api.konghq.com)`,
				Optional:    true,
			},
			"system_account_access_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
		MarkdownDescription: `Konnect API (BETA): This is a BETA specification. Endpoints in this specification may change with zero notice`,
	}
}

func (p *KonnectBetaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data KonnectBetaProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" && len(os.Getenv("KONNECT_SERVER_URL")) > 0 {
		ServerURL = os.Getenv("KONNECT_SERVER_URL")
	}
	if ServerURL == "" {
		ServerURL = "https://global.api.konghq.com"
	}

	personalAccessToken := new(string)
	if !data.PersonalAccessToken.IsUnknown() && !data.PersonalAccessToken.IsNull() {
		*personalAccessToken = data.PersonalAccessToken.ValueString()
	} else {
		if len(os.Getenv("KONNECT_TOKEN")) > 0 {
			*personalAccessToken = os.Getenv("KONNECT_TOKEN")
		} else {
			personalAccessToken = nil
		}
	}
	systemAccountAccessToken := new(string)
	if !data.SystemAccountAccessToken.IsUnknown() && !data.SystemAccountAccessToken.IsNull() {
		*systemAccountAccessToken = data.SystemAccountAccessToken.ValueString()
	} else {
		if len(os.Getenv("KONNECT_SPAT")) > 0 {
			*systemAccountAccessToken = os.Getenv("KONNECT_SPAT")
		} else {
			systemAccountAccessToken = nil
		}
	}
	konnectAccessToken := new(string)
	if !data.KonnectAccessToken.IsUnknown() && !data.KonnectAccessToken.IsNull() {
		*konnectAccessToken = data.KonnectAccessToken.ValueString()
	} else {
		konnectAccessToken = nil
	}
	security := shared.Security{
		PersonalAccessToken:      personalAccessToken,
		SystemAccountAccessToken: systemAccountAccessToken,
		KonnectAccessToken:       konnectAccessToken,
	}

	providerHTTPTransportOpts := ProviderHTTPTransportOpts{
		SetHeaders: make(map[string]string),
		Transport:  http.DefaultTransport,
	}

	httpClient := http.DefaultClient
	httpClient.Transport = NewProviderHTTPTransport(providerHTTPTransportOpts)

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
		sdk.WithClient(httpClient),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *KonnectBetaProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAPIResource,
		NewAPIDocumentResource,
		NewAPIImplementationResource,
		NewAPIPublicationResource,
		NewAPISpecificationResource,
		NewMeshResource,
		NewMeshAccessLogResource,
		NewMeshCircuitBreakerResource,
		NewMeshControlPlaneResource,
		NewMeshExternalServiceResource,
		NewMeshFaultInjectionResource,
		NewMeshGatewayResource,
		NewMeshGlobalRateLimitResource,
		NewMeshHealthCheckResource,
		NewMeshHostnameGeneratorResource,
		NewMeshHTTPRouteResource,
		NewMeshLoadBalancingStrategyResource,
		NewMeshMetricResource,
		NewMeshMultiZoneServiceResource,
		NewMeshOPAResource,
		NewMeshPassthroughResource,
		NewMeshProxyPatchResource,
		NewMeshRateLimitResource,
		NewMeshRetryResource,
		NewMeshServiceResource,
		NewMeshTCPRouteResource,
		NewMeshTimeoutResource,
		NewMeshTLSResource,
		NewMeshTraceResource,
		NewMeshTrafficPermissionResource,
		NewPortalResource,
		NewPortalAuthResource,
		NewPortalCustomDomainResource,
		NewPortalCustomizationResource,
		NewPortalPageResource,
		NewPortalSnippetResource,
		NewPortalTeamResource,
	}
}

func (p *KonnectBetaProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAPIDataSource,
		NewAPIDocumentDataSource,
		NewAPIImplementationDataSource,
		NewAPIPublicationDataSource,
		NewAPISpecificationDataSource,
		NewHostnameGeneratorListDataSource,
		NewMeshDataSource,
		NewMeshAccessLogDataSource,
		NewMeshAccessLogListDataSource,
		NewMeshCircuitBreakerDataSource,
		NewMeshCircuitBreakerListDataSource,
		NewMeshControlPlaneDataSource,
		NewMeshExternalServiceDataSource,
		NewMeshExternalServiceListDataSource,
		NewMeshFaultInjectionDataSource,
		NewMeshFaultInjectionListDataSource,
		NewMeshGatewayDataSource,
		NewMeshGatewayListDataSource,
		NewMeshGlobalRateLimitDataSource,
		NewMeshGlobalRateLimitListDataSource,
		NewMeshHealthCheckDataSource,
		NewMeshHealthCheckListDataSource,
		NewMeshHostnameGeneratorDataSource,
		NewMeshHTTPRouteDataSource,
		NewMeshHTTPRouteListDataSource,
		NewMeshListDataSource,
		NewMeshLoadBalancingStrategyDataSource,
		NewMeshLoadBalancingStrategyListDataSource,
		NewMeshMetricDataSource,
		NewMeshMetricListDataSource,
		NewMeshMultiZoneServiceDataSource,
		NewMeshMultiZoneServiceListDataSource,
		NewMeshOPADataSource,
		NewMeshOPAListDataSource,
		NewMeshPassthroughDataSource,
		NewMeshPassthroughListDataSource,
		NewMeshProxyPatchDataSource,
		NewMeshProxyPatchListDataSource,
		NewMeshRateLimitDataSource,
		NewMeshRateLimitListDataSource,
		NewMeshRetryDataSource,
		NewMeshRetryListDataSource,
		NewMeshServiceDataSource,
		NewMeshServiceListDataSource,
		NewMeshTCPRouteDataSource,
		NewMeshTCPRouteListDataSource,
		NewMeshTimeoutDataSource,
		NewMeshTimeoutListDataSource,
		NewMeshTLSDataSource,
		NewMeshTLSListDataSource,
		NewMeshTraceDataSource,
		NewMeshTraceListDataSource,
		NewMeshTrafficPermissionDataSource,
		NewMeshTrafficPermissionListDataSource,
		NewPortalDataSource,
		NewPortalAuthDataSource,
		NewPortalCustomDomainDataSource,
		NewPortalCustomizationDataSource,
		NewPortalPageDataSource,
		NewPortalSnippetDataSource,
		NewPortalTeamDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &KonnectBetaProvider{
			version: version,
		}
	}
}
