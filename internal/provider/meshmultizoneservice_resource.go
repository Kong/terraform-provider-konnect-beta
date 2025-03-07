// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	custom_listplanmodifier "github.com/kong/terraform-provider-konnect-beta/internal/planmodifiers/listplanmodifier"
	speakeasy_listplanmodifier "github.com/kong/terraform-provider-konnect-beta/internal/planmodifiers/listplanmodifier"
	speakeasy_objectplanmodifier "github.com/kong/terraform-provider-konnect-beta/internal/planmodifiers/objectplanmodifier"
	custom_stringplanmodifier "github.com/kong/terraform-provider-konnect-beta/internal/planmodifiers/stringplanmodifier"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-konnect-beta/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/kong/terraform-provider-konnect-beta/internal/provider/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect-beta/internal/validators"
	speakeasy_int64validators "github.com/kong/terraform-provider-konnect-beta/internal/validators/int64validators"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-konnect-beta/internal/validators/objectvalidators"
	"regexp"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MeshMultiZoneServiceResource{}
var _ resource.ResourceWithImportState = &MeshMultiZoneServiceResource{}

func NewMeshMultiZoneServiceResource() resource.Resource {
	return &MeshMultiZoneServiceResource{}
}

// MeshMultiZoneServiceResource defines the resource implementation.
type MeshMultiZoneServiceResource struct {
	client *sdk.KonnectBeta
}

// MeshMultiZoneServiceResourceModel describes the resource data model.
type MeshMultiZoneServiceResourceModel struct {
	CpID             types.String                            `tfsdk:"cp_id"`
	CreationTime     types.String                            `tfsdk:"creation_time"`
	Labels           map[string]types.String                 `tfsdk:"labels"`
	Mesh             types.String                            `tfsdk:"mesh"`
	ModificationTime types.String                            `tfsdk:"modification_time"`
	Name             types.String                            `tfsdk:"name"`
	Spec             tfTypes.MeshMultiZoneServiceItemSpec    `tfsdk:"spec"`
	Status           *tfTypes.MeshMultiZoneServiceItemStatus `tfsdk:"status"`
	Type             types.String                            `tfsdk:"type"`
	Warnings         []types.String                          `tfsdk:"warnings"`
}

func (r *MeshMultiZoneServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "konnect_mesh_multi_zone_service"
}

func (r *MeshMultiZoneServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshMultiZoneService Resource",
		Attributes: map[string]schema.Attribute{
			"cp_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					custom_stringplanmodifier.RequiresReplaceModifier(),
				},
				Description: `Id of the Konnect resource`,
			},
			"creation_time": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `Time at which the resource was created`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"labels": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: `The labels to help identity resources`,
			},
			"mesh": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					custom_stringplanmodifier.RequiresReplaceModifier(),
				},
				Description: `name of the mesh`,
			},
			"modification_time": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `Time at which the resource was updated`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					custom_stringplanmodifier.RequiresReplaceModifier(),
				},
				Description: `name of the MeshMultiZoneService`,
			},
			"spec": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"ports": schema.ListNestedAttribute{
						Required: true,
						PlanModifiers: []planmodifier.List{
							custom_listplanmodifier.SupressZeroNullModifier(),
						},
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							Attributes: map[string]schema.Attribute{
								"app_protocol": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Default:     stringdefault.StaticString("tcp"),
									Description: `Protocol identifies a protocol supported by a service. Default: "tcp"`,
								},
								"name": schema.StringAttribute{
									Optional: true,
								},
								"port": schema.Int64Attribute{
									Optional:    true,
									Description: `Not Null`,
									Validators: []validator.Int64{
										speakeasy_int64validators.NotNull(),
									},
								},
							},
						},
						Description: `Ports is a list of ports from selected MeshServices`,
						Validators: []validator.List{
							listvalidator.SizeAtLeast(1),
						},
					},
					"selector": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"mesh_service": schema.SingleNestedAttribute{
								Required: true,
								Attributes: map[string]schema.Attribute{
									"match_labels": schema.MapAttribute{
										Required:    true,
										ElementType: types.StringType,
										Description: `MatchLabels matches multiple MeshServices by labels`,
									},
								},
								Description: `MeshService selects MeshServices`,
							},
						},
						Description: `Selector is a way to select multiple MeshServices`,
					},
				},
				Description: `Spec is the specification of the Kuma MeshMultiZoneService resource.`,
			},
			"status": schema.SingleNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Attributes: map[string]schema.Attribute{
					"addresses": schema.ListNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.List{
							custom_listplanmodifier.SupressZeroNullModifier(),
							speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
						},
						NestedObject: schema.NestedAttributeObject{
							PlanModifiers: []planmodifier.Object{
								speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
							},
							Attributes: map[string]schema.Attribute{
								"hostname": schema.StringAttribute{
									Computed: true,
								},
								"hostname_generator_ref": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"core_name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"origin": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						Description: `Addresses is a list of addresses generated by HostnameGenerator`,
					},
					"hostname_generators": schema.ListNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.List{
							custom_listplanmodifier.SupressZeroNullModifier(),
							speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
						},
						NestedObject: schema.NestedAttributeObject{
							PlanModifiers: []planmodifier.Object{
								speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
							},
							Attributes: map[string]schema.Attribute{
								"conditions": schema.ListNestedAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.List{
										custom_listplanmodifier.SupressZeroNullModifier(),
										speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
									},
									NestedObject: schema.NestedAttributeObject{
										PlanModifiers: []planmodifier.Object{
											speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
										},
										Attributes: map[string]schema.Attribute{
											"message": schema.StringAttribute{
												Computed: true,
												MarkdownDescription: `message is a human readable message indicating details about the transition.` + "\n" +
													`This may be an empty string.`,
												Validators: []validator.String{
													stringvalidator.UTF8LengthAtMost(32768),
												},
											},
											"reason": schema.StringAttribute{
												Computed: true,
												MarkdownDescription: `reason contains a programmatic identifier indicating the reason for the condition's last transition.` + "\n" +
													`Producers of specific condition types may define expected values and meanings for this field,` + "\n" +
													`and whether the values are considered a guaranteed API.` + "\n" +
													`The value should be a CamelCase string.` + "\n" +
													`This field may not be empty.`,
												Validators: []validator.String{
													stringvalidator.UTF8LengthBetween(1, 1024),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$`), "must match pattern "+regexp.MustCompile(`^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$`).String()),
												},
											},
											"status": schema.StringAttribute{
												Computed: true,
												PlanModifiers: []planmodifier.String{
													speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
												},
												Description: `status of the condition, one of True, False, Unknown. must be one of ["True", "False", "Unknown"]`,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"True",
														"False",
														"Unknown",
													),
												},
											},
											"type": schema.StringAttribute{
												Computed:    true,
												Description: `type of condition in CamelCase or in foo.example.com/CamelCase.`,
												Validators: []validator.String{
													stringvalidator.UTF8LengthAtMost(316),
													stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`), "must match pattern "+regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`).String()),
												},
											},
										},
									},
									Description: `Conditions is an array of hostname generator conditions.`,
								},
								"hostname_generator_ref": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"core_name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
						},
						Description: `Status of hostnames generator applied on this resource`,
					},
					"mesh_services": schema.ListNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.List{
							custom_listplanmodifier.SupressZeroNullModifier(),
							speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
						},
						NestedObject: schema.NestedAttributeObject{
							PlanModifiers: []planmodifier.Object{
								speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
							},
							Attributes: map[string]schema.Attribute{
								"mesh": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed:    true,
									Description: `Name is a core name of MeshService`,
								},
								"namespace": schema.StringAttribute{
									Computed: true,
								},
								"zone": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						Description: `MeshServices is a list of matched MeshServices`,
					},
					"vips": schema.ListNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.List{
							custom_listplanmodifier.SupressZeroNullModifier(),
							speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
						},
						NestedObject: schema.NestedAttributeObject{
							PlanModifiers: []planmodifier.Object{
								speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
							},
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						Description: `VIPs is a list of assigned Kuma VIPs.`,
					},
				},
				Description: `Status is the current status of the Kuma MeshMultiZoneService resource.`,
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: `the type of the resource. must be "MeshMultiZoneService"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"MeshMultiZoneService",
					),
				},
			},
			"warnings": schema.ListAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					custom_listplanmodifier.SupressZeroNullModifier(),
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				ElementType: types.StringType,
				MarkdownDescription: `warnings is a list of warning messages to return to the requesting Kuma API clients.` + "\n" +
					`Warning messages describe a problem the client making the API request should correct or be aware of.`,
			},
		},
	}
}

func (r *MeshMultiZoneServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MeshMultiZoneServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MeshMultiZoneServiceResourceModel
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

	var cpID string
	cpID = data.CpID.ValueString()

	var mesh string
	mesh = data.Mesh.ValueString()

	var name string
	name = data.Name.ValueString()

	meshMultiZoneServiceItem := *data.ToSharedMeshMultiZoneServiceItemInput()
	request := operations.CreateMeshMultiZoneServiceRequest{
		CpID:                     cpID,
		Mesh:                     mesh,
		Name:                     name,
		MeshMultiZoneServiceItem: meshMultiZoneServiceItem,
	}
	res, err := r.client.MeshMultiZoneService.CreateMeshMultiZoneService(ctx, request)
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
	if !(res.MeshMultiZoneServiceCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshMultiZoneServiceCreateOrUpdateSuccessResponse(res.MeshMultiZoneServiceCreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var cpId1 string
	cpId1 = data.CpID.ValueString()

	var mesh1 string
	mesh1 = data.Mesh.ValueString()

	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetMeshMultiZoneServiceRequest{
		CpID: cpId1,
		Mesh: mesh1,
		Name: name1,
	}
	res1, err := r.client.MeshMultiZoneService.GetMeshMultiZoneService(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.MeshMultiZoneServiceItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedMeshMultiZoneServiceItem(res1.MeshMultiZoneServiceItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshMultiZoneServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MeshMultiZoneServiceResourceModel
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

	var cpID string
	cpID = data.CpID.ValueString()

	var mesh string
	mesh = data.Mesh.ValueString()

	var name string
	name = data.Name.ValueString()

	request := operations.GetMeshMultiZoneServiceRequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshMultiZoneService.GetMeshMultiZoneService(ctx, request)
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
	if !(res.MeshMultiZoneServiceItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshMultiZoneServiceItem(res.MeshMultiZoneServiceItem)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshMultiZoneServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *MeshMultiZoneServiceResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var cpID string
	cpID = data.CpID.ValueString()

	var mesh string
	mesh = data.Mesh.ValueString()

	var name string
	name = data.Name.ValueString()

	meshMultiZoneServiceItem := *data.ToSharedMeshMultiZoneServiceItemInput()
	request := operations.UpdateMeshMultiZoneServiceRequest{
		CpID:                     cpID,
		Mesh:                     mesh,
		Name:                     name,
		MeshMultiZoneServiceItem: meshMultiZoneServiceItem,
	}
	res, err := r.client.MeshMultiZoneService.UpdateMeshMultiZoneService(ctx, request)
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
	if !(res.MeshMultiZoneServiceCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshMultiZoneServiceCreateOrUpdateSuccessResponse(res.MeshMultiZoneServiceCreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var cpId1 string
	cpId1 = data.CpID.ValueString()

	var mesh1 string
	mesh1 = data.Mesh.ValueString()

	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetMeshMultiZoneServiceRequest{
		CpID: cpId1,
		Mesh: mesh1,
		Name: name1,
	}
	res1, err := r.client.MeshMultiZoneService.GetMeshMultiZoneService(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.MeshMultiZoneServiceItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedMeshMultiZoneServiceItem(res1.MeshMultiZoneServiceItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshMultiZoneServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MeshMultiZoneServiceResourceModel
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

	var cpID string
	cpID = data.CpID.ValueString()

	var mesh string
	mesh = data.Mesh.ValueString()

	var name string
	name = data.Name.ValueString()

	request := operations.DeleteMeshMultiZoneServiceRequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshMultiZoneService.DeleteMeshMultiZoneService(ctx, request)
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

}

func (r *MeshMultiZoneServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		CpID string `json:"cp_id"`
		Mesh string `json:"mesh"`
		Name string `json:"name"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "cp_id": "bf138ba2-c9b1-4229-b268-04d9d8a6410b",  "mesh": "",  "name": ""}': `+err.Error())
		return
	}

	if len(data.CpID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field cp_id is required but was not found in the json encoded ID. It's expected to be a value alike '"bf138ba2-c9b1-4229-b268-04d9d8a6410b"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cp_id"), data.CpID)...)
	if len(data.Mesh) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field mesh is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("mesh"), data.Mesh)...)
	if len(data.Name) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field name is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), data.Name)...)

}
