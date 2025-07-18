// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
	"os"
)

var _ provider.Provider = (*KonnectBetaProvider)(nil)
var _ provider.ProviderWithEphemeralResources = (*KonnectBetaProvider)(nil)

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

	security := shared.Security{}

	if !data.PersonalAccessToken.IsUnknown() {
		security.PersonalAccessToken = data.PersonalAccessToken.ValueStringPointer()
	}

	if personalAccessTokenEnvVar := os.Getenv("KONNECT_TOKEN"); security.PersonalAccessToken == nil && personalAccessTokenEnvVar != "" {
		security.PersonalAccessToken = &personalAccessTokenEnvVar
	}

	if !data.SystemAccountAccessToken.IsUnknown() {
		security.SystemAccountAccessToken = data.SystemAccountAccessToken.ValueStringPointer()
	}

	if systemAccountAccessTokenEnvVar := os.Getenv("KONNECT_SPAT"); security.SystemAccountAccessToken == nil && systemAccountAccessTokenEnvVar != "" {
		security.SystemAccountAccessToken = &systemAccountAccessTokenEnvVar
	}

	if !data.KonnectAccessToken.IsUnknown() {
		security.KonnectAccessToken = data.KonnectAccessToken.ValueStringPointer()
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
	resp.EphemeralResourceData = client
	resp.ResourceData = client
}

func (p *KonnectBetaProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAPIResource,
		NewAPIDocumentResource,
		NewAPIImplementationResource,
		NewAPIPublicationResource,
		NewAPISpecificationResource,
		NewAPIVersionResource,
		NewAuthServerResource,
		NewAuthServerClaimsResource,
		NewAuthServerClientsResource,
		NewAuthServerScopesResource,
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
		NewAPIVersionDataSource,
		NewAuthServerDataSource,
		NewAuthServerClaimsDataSource,
		NewAuthServerClientsDataSource,
		NewAuthServerScopesDataSource,
		NewHostnameGeneratorListDataSource,
		NewMeshDataSource,
		NewMeshAccessLogDataSource,
		NewMeshAccessLogListDataSource,
		NewMeshCircuitBreakerDataSource,
		NewMeshCircuitBreakerListDataSource,
		NewMeshControlPlaneDataSource,
		NewMeshControlPlanesDataSource,
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

func (p *KonnectBetaProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &KonnectBetaProvider{
			version: version,
		}
	}
}
