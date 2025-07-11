// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-konnect-beta/internal/planmodifiers/stringplanmodifier"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk"
	"github.com/kong/terraform-provider-konnect-beta/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PortalResource{}
var _ resource.ResourceWithImportState = &PortalResource{}

func NewPortalResource() resource.Resource {
	return &PortalResource{}
}

// PortalResource defines the resource implementation.
type PortalResource struct {
	client *sdk.KonnectBeta
}

// PortalResourceModel describes the resource data model.
type PortalResourceModel struct {
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
	ForceDestroy                     types.String            `queryParam:"style=form,explode=true,name=force" tfsdk:"force_destroy"`
	ID                               types.String            `tfsdk:"id"`
	Labels                           map[string]types.String `tfsdk:"labels"`
	Name                             types.String            `tfsdk:"name"`
	RbacEnabled                      types.Bool              `tfsdk:"rbac_enabled"`
	UpdatedAt                        types.String            `tfsdk:"updated_at"`
}

func (r *PortalResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "konnect_portal"
}

func (r *PortalResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Portal Resource",
		Attributes: map[string]schema.Attribute{
			"authentication_enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the portal supports developer authentication. If disabled, developers cannot register for accounts or create applications.`,
			},
			"auto_approve_applications": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether requests from applications to register for APIs will be automatically approved, or if they will be set to pending until approved by an admin.`,
			},
			"auto_approve_developers": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether developer account registrations will be automatically approved, or if they will be set to pending until approved by an admin.`,
			},
			"canonical_domain": schema.StringAttribute{
				Computed:    true,
				Description: `The canonical domain of the developer portal`,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `An ISO-8601 timestamp representation of entity creation date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"default_api_visibility": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The default visibility of APIs in the portal. If set to ` + "`" + `public` + "`" + `, newly published APIs are visible to unauthenticated developers. If set to ` + "`" + `private` + "`" + `, newly published APIs are hidden from unauthenticated developers. must be one of ["public", "private"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"public",
						"private",
					),
				},
			},
			"default_application_auth_strategy_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The default authentication strategy for APIs published to the portal. Newly published APIs will use this authentication strategy unless overridden during publication. If set to ` + "`" + `null` + "`" + `, API publications will not use an authentication strategy unless set during publication.`,
			},
			"default_domain": schema.StringAttribute{
				Computed:    true,
				Description: `The domain assigned to the portal by Konnect. This is the default place to access the portal and its API if not using a ` + "`" + `custom_domain` + "``" + `.`,
			},
			"default_page_visibility": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The default visibility of pages in the portal. If set to ` + "`" + `public` + "`" + `, newly created pages are visible to unauthenticated developers. If set to ` + "`" + `private` + "`" + `, newly created pages are hidden from unauthenticated developers. must be one of ["public", "private"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"public",
						"private",
					),
				},
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `A description of the portal.`,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(512),
				},
			},
			"display_name": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The display name of the portal. This value will be the portal's ` + "`" + `name` + "`" + ` in Portal API.`,
				Validators: []validator.String{
					stringvalidator.UTF8LengthBetween(1, 255),
				},
			},
			"force_destroy": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Default:  stringdefault.StaticString(`false`),
				MarkdownDescription: `If set to "true", the portal and all child entities will be deleted when running ` + "`" + `terraform destroy` + "`" + `.` + "\n" +
					`If set to "false", the portal will not be deleted until all child entities are manually removed.` + "\n" +
					`This will IRREVERSIBLY DELETE ALL REGISTERED DEVELOPERS AND THEIR CREDENTIALS. Only set to "true" if you want this behavior.` + "\n" +
					`Default: "false"; must be one of ["true", "false"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"true",
						"false",
					),
				},
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `Contains a unique identifier used for this resource.`,
			},
			"labels": schema.MapAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				MarkdownDescription: `Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types. ` + "\n" +
					`` + "\n" +
					`Labels are intended to store **INTERNAL** metadata.` + "\n" +
					`` + "\n" +
					`Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".`,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `The name of the portal, used to distinguish it from other portals. Name must be unique.`,
				Validators: []validator.String{
					stringvalidator.UTF8LengthBetween(1, 255),
				},
			},
			"rbac_enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the portal resources are protected by Role Based Access Control (RBAC). If enabled, developers view or register for APIs until unless assigned to teams with access to view and consume specific APIs. Authentication must be enabled to use RBAC.`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `An ISO-8601 timestamp representation of entity update date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
		},
	}
}

func (r *PortalResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.KonnectBeta)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.KonnectBeta, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *PortalResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PortalResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request, requestDiags := data.ToSharedCreatePortal(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Portals.CreatePortal(ctx, *request)
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
	if res.StatusCode != 201 {
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

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PortalResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PortalResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
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

func (r *PortalResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PortalResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	request, requestDiags := data.ToOperationsUpdatePortalRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Portals.UpdatePortal(ctx, *request)
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

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PortalResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PortalResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	request, requestDiags := data.ToOperationsDeletePortalRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Portals.DeletePortal(ctx, *request)
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
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *PortalResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
