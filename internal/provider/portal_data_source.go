// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &PortalDataSource{}
var _ datasource.DataSourceWithConfigure = &PortalDataSource{}

func NewPortalDataSource() datasource.DataSource {
	return &PortalDataSource{}
}

// PortalDataSource is the data source implementation.
type PortalDataSource struct {
	client *sdk.KonnectBeta
}

// PortalDataSourceModel describes the data model.
type PortalDataSourceModel struct {
	AuthenticationEnabled            types.Bool              `tfsdk:"authentication_enabled"`
	AutoApproveApplications          types.Bool              `tfsdk:"auto_approve_applications"`
	AutoApproveDevelopers            types.Bool              `tfsdk:"auto_approve_developers"`
	CanonicalDomain                  types.String            `tfsdk:"canonical_domain"`
	CreatedAt                        types.String            `tfsdk:"created_at"`
	DefaultAPIVisibility             types.String            `tfsdk:"default_api_visibility"`
	DefaultApplicationAuthStrategyID types.String            `tfsdk:"default_application_auth_strategy_id"`
	DefaultDomain                    types.String            `tfsdk:"default_domain"`
	DefaultPageVisibility            types.String            `tfsdk:"default_page_visibility"`
	Description                      types.String            `tfsdk:"description"`
	DisplayName                      types.String            `tfsdk:"display_name"`
	ID                               types.String            `tfsdk:"id"`
	Labels                           map[string]types.String `tfsdk:"labels"`
	Name                             types.String            `tfsdk:"name"`
	RbacEnabled                      types.Bool              `tfsdk:"rbac_enabled"`
	UpdatedAt                        types.String            `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *PortalDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_portal"
}

// Schema defines the schema for the data source.
func (r *PortalDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Portal DataSource",

		Attributes: map[string]schema.Attribute{
			"authentication_enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the portal supports developer authentication. If disabled, developers cannot register for accounts or create applications.`,
			},
			"auto_approve_applications": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether requests from applications to register for APIs will be automatically approved, or if they will be set to pending until approved by an admin.`,
			},
			"auto_approve_developers": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether developer account registrations will be automatically approved, or if they will be set to pending until approved by an admin.`,
			},
			"canonical_domain": schema.StringAttribute{
				Computed:    true,
				Description: `The canonical domain of the developer portal`,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity creation date.`,
			},
			"default_api_visibility": schema.StringAttribute{
				Computed:    true,
				Description: `The default visibility of APIs in the portal. If set to ` + "`" + `public` + "`" + `, newly published APIs are visible to unauthenticated developers. If set to ` + "`" + `private` + "`" + `, newly published APIs are hidden from unauthenticated developers.`,
			},
			"default_application_auth_strategy_id": schema.StringAttribute{
				Computed:    true,
				Description: `The default authentication strategy for APIs published to the portal. Newly published APIs will use this authentication strategy unless overridden during publication. If set to ` + "`" + `null` + "`" + `, API publications will not use an authentication strategy unless set during publication.`,
			},
			"default_domain": schema.StringAttribute{
				Computed:    true,
				Description: `The domain assigned to the portal by Konnect. This is the default place to access the portal and its API if not using a ` + "`" + `custom_domain` + "``" + `.`,
			},
			"default_page_visibility": schema.StringAttribute{
				Computed:    true,
				Description: `The default visibility of pages in the portal. If set to ` + "`" + `public` + "`" + `, newly created pages are visible to unauthenticated developers. If set to ` + "`" + `private` + "`" + `, newly created pages are hidden from unauthenticated developers.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `A description of the portal.`,
			},
			"display_name": schema.StringAttribute{
				Computed:    true,
				Description: `The display name of the portal. This value will be the portal's ` + "`" + `name` + "`" + ` in Portal API.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `Contains a unique identifier used for this resource.`,
			},
			"labels": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
				MarkdownDescription: `Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types. ` + "\n" +
					`` + "\n" +
					`Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the portal, used to distinguish it from other portals. Name must be unique.`,
			},
			"rbac_enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the portal resources are protected by Role Based Access Control (RBAC). If enabled, developers view or register for APIs until unless assigned to teams with access to view and consume specific APIs. Authentication must be enabled to use RBAC.`,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity update date.`,
			},
		},
	}
}

func (r *PortalDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.KonnectBeta)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.KonnectBeta, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *PortalDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PortalDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request, requestDiags := data.ToOperationsGetPortalRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Portals.GetPortal(ctx, *request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.PortalResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedPortalResponse(ctx, res.PortalResponse)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
