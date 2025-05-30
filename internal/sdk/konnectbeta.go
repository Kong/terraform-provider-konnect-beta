// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/internal/hooks"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/retry"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://global.api.konghq.com",
	"https://us.api.konghq.com",
	"https://eu.api.konghq.com",
	"https://au.api.konghq.com",
	"https://me.api.konghq.com",
	"https://in.api.konghq.com",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// KonnectBeta - Konnect API (BETA): This is a BETA specification. Endpoints in this specification may change with zero notice
type KonnectBeta struct {
	// APIs related to configuration of Konnect Developer Portals.
	Portals *Portals
	// APIs related to configuration of Konnect Developer Portals custom domains.
	PortalCustomDomains *PortalCustomDomains
	// APIs related to customization of Konnect Developer Portals.
	PortalCustomization *PortalCustomization
	// APIs related to Konnect Developer Portal Custom Pages.
	Pages *Pages
	// APIs related to Konnect Developer Portal Custom Snippets.
	Snippets *Snippets
	// APIs related to configuration of Konnect Developer Portal auth settings.
	PortalAuthSettings *PortalAuthSettings
	// APIs related to configuration of Konnect Developer Portal developer teams.
	PortalTeams               *PortalTeams
	API                       *API
	APIDocumentation          *APIDocumentation
	APISpecification          *APISpecification
	APIPublication            *APIPublication
	APIImplementation         *APIImplementation
	MeshAccessLog             *MeshAccessLog
	MeshCircuitBreaker        *MeshCircuitBreaker
	MeshFaultInjection        *MeshFaultInjection
	MeshHealthCheck           *MeshHealthCheck
	MeshHTTPRoute             *MeshHTTPRoute
	MeshLoadBalancingStrategy *MeshLoadBalancingStrategy
	MeshMetric                *MeshMetric
	MeshPassthrough           *MeshPassthrough
	MeshProxyPatch            *MeshProxyPatch
	MeshRateLimit             *MeshRateLimit
	MeshRetry                 *MeshRetry
	MeshTCPRoute              *MeshTCPRoute
	MeshTimeout               *MeshTimeout
	MeshTLS                   *MeshTLS
	MeshTrace                 *MeshTrace
	MeshTrafficPermission     *MeshTrafficPermission
	Mesh                      *Mesh
	MeshGateway               *MeshGateway
	HostnameGenerator         *HostnameGenerator
	MeshExternalService       *MeshExternalService
	MeshMultiZoneService      *MeshMultiZoneService
	MeshService               *MeshService
	MeshGlobalRateLimit       *MeshGlobalRateLimit
	MeshOPA                   *MeshOPA

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*KonnectBeta)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *KonnectBeta) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *KonnectBeta) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *KonnectBeta) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *KonnectBeta) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *KonnectBeta) {
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *KonnectBeta) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *KonnectBeta) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *KonnectBeta) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *KonnectBeta {
	sdk := &KonnectBeta{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "2.0.0",
			SDKVersion:        "0.5.0",
			GenVersion:        "2.610.0",
			UserAgent:         "speakeasy-sdk/terraform 0.5.0 2.610.0 2.0.0 github.com/kong/terraform-provider-konnect-beta/internal/sdk",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Portals = newPortals(sdk.sdkConfiguration)

	sdk.PortalCustomDomains = newPortalCustomDomains(sdk.sdkConfiguration)

	sdk.PortalCustomization = newPortalCustomization(sdk.sdkConfiguration)

	sdk.Pages = newPages(sdk.sdkConfiguration)

	sdk.Snippets = newSnippets(sdk.sdkConfiguration)

	sdk.PortalAuthSettings = newPortalAuthSettings(sdk.sdkConfiguration)

	sdk.PortalTeams = newPortalTeams(sdk.sdkConfiguration)

	sdk.API = newAPI(sdk.sdkConfiguration)

	sdk.APIDocumentation = newAPIDocumentation(sdk.sdkConfiguration)

	sdk.APISpecification = newAPISpecification(sdk.sdkConfiguration)

	sdk.APIPublication = newAPIPublication(sdk.sdkConfiguration)

	sdk.APIImplementation = newAPIImplementation(sdk.sdkConfiguration)

	sdk.MeshAccessLog = newMeshAccessLog(sdk.sdkConfiguration)

	sdk.MeshCircuitBreaker = newMeshCircuitBreaker(sdk.sdkConfiguration)

	sdk.MeshFaultInjection = newMeshFaultInjection(sdk.sdkConfiguration)

	sdk.MeshHealthCheck = newMeshHealthCheck(sdk.sdkConfiguration)

	sdk.MeshHTTPRoute = newMeshHTTPRoute(sdk.sdkConfiguration)

	sdk.MeshLoadBalancingStrategy = newMeshLoadBalancingStrategy(sdk.sdkConfiguration)

	sdk.MeshMetric = newMeshMetric(sdk.sdkConfiguration)

	sdk.MeshPassthrough = newMeshPassthrough(sdk.sdkConfiguration)

	sdk.MeshProxyPatch = newMeshProxyPatch(sdk.sdkConfiguration)

	sdk.MeshRateLimit = newMeshRateLimit(sdk.sdkConfiguration)

	sdk.MeshRetry = newMeshRetry(sdk.sdkConfiguration)

	sdk.MeshTCPRoute = newMeshTCPRoute(sdk.sdkConfiguration)

	sdk.MeshTimeout = newMeshTimeout(sdk.sdkConfiguration)

	sdk.MeshTLS = newMeshTLS(sdk.sdkConfiguration)

	sdk.MeshTrace = newMeshTrace(sdk.sdkConfiguration)

	sdk.MeshTrafficPermission = newMeshTrafficPermission(sdk.sdkConfiguration)

	sdk.Mesh = newMesh(sdk.sdkConfiguration)

	sdk.MeshGateway = newMeshGateway(sdk.sdkConfiguration)

	sdk.HostnameGenerator = newHostnameGenerator(sdk.sdkConfiguration)

	sdk.MeshExternalService = newMeshExternalService(sdk.sdkConfiguration)

	sdk.MeshMultiZoneService = newMeshMultiZoneService(sdk.sdkConfiguration)

	sdk.MeshService = newMeshService(sdk.sdkConfiguration)

	sdk.MeshGlobalRateLimit = newMeshGlobalRateLimit(sdk.sdkConfiguration)

	sdk.MeshOPA = newMeshOPA(sdk.sdkConfiguration)

	return sdk
}
