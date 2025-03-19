// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	custom_listplanmodifier "github.com/kong/terraform-provider-konnect-beta/internal/planmodifiers/listplanmodifier"
	custom_stringplanmodifier "github.com/kong/terraform-provider-konnect-beta/internal/planmodifiers/stringplanmodifier"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-konnect-beta/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/kong/terraform-provider-konnect-beta/internal/provider/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect-beta/internal/validators"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-konnect-beta/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/kong/terraform-provider-konnect-beta/internal/validators/stringvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MeshTraceResource{}
var _ resource.ResourceWithImportState = &MeshTraceResource{}

func NewMeshTraceResource() resource.Resource {
	return &MeshTraceResource{}
}

// MeshTraceResource defines the resource implementation.
type MeshTraceResource struct {
	client *sdk.KonnectBeta
}

// MeshTraceResourceModel describes the resource data model.
type MeshTraceResourceModel struct {
	CpID             types.String              `tfsdk:"cp_id"`
	CreationTime     types.String              `tfsdk:"creation_time"`
	Labels           map[string]types.String   `tfsdk:"labels"`
	Mesh             types.String              `tfsdk:"mesh"`
	ModificationTime types.String              `tfsdk:"modification_time"`
	Name             types.String              `tfsdk:"name"`
	Spec             tfTypes.MeshTraceItemSpec `tfsdk:"spec"`
	Type             types.String              `tfsdk:"type"`
	Warnings         []types.String            `tfsdk:"warnings"`
}

func (r *MeshTraceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "konnect_mesh_trace"
}

func (r *MeshTraceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshTrace Resource",
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
				Description: `name of the MeshTrace`,
			},
			"spec": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"default": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"backends": schema.ListNestedAttribute{
								Optional: true,
								PlanModifiers: []planmodifier.List{
									custom_listplanmodifier.SupressZeroNullModifier(),
								},
								NestedObject: schema.NestedAttributeObject{
									Validators: []validator.Object{
										speakeasy_objectvalidators.NotNull(),
									},
									Attributes: map[string]schema.Attribute{
										"datadog": schema.SingleNestedAttribute{
											Optional: true,
											Attributes: map[string]schema.Attribute{
												"split_service": schema.BoolAttribute{
													Computed: true,
													Optional: true,
													Default:  booldefault.StaticBool(false),
													MarkdownDescription: `Determines if datadog service name should be split based on traffic` + "\n" +
														`direction and destination. For example, with ` + "`" + `splitService: true` + "`" + ` and a` + "\n" +
														`` + "`" + `backend` + "`" + ` service that communicates with a couple of databases, you would` + "\n" +
														`get service names like ` + "`" + `backend_INBOUND` + "`" + `, ` + "`" + `backend_OUTBOUND_db1` + "`" + `, and` + "\n" +
														`` + "`" + `backend_OUTBOUND_db2` + "`" + ` in Datadog.` + "\n" +
														`Default: false`,
												},
												"url": schema.StringAttribute{
													Optional: true,
													MarkdownDescription: `Address of Datadog collector, only host and port are allowed (no paths,` + "\n" +
														`fragments etc.)` + "\n" +
														`Not Null`,
													Validators: []validator.String{
														speakeasy_stringvalidators.NotNull(),
													},
												},
											},
											Description: `Datadog backend configuration.`,
										},
										"open_telemetry": schema.SingleNestedAttribute{
											Optional: true,
											Attributes: map[string]schema.Attribute{
												"endpoint": schema.StringAttribute{
													Optional:    true,
													Description: `Address of OpenTelemetry collector. Not Null`,
													Validators: []validator.String{
														speakeasy_stringvalidators.NotNull(),
														stringvalidator.UTF8LengthAtLeast(1),
													},
												},
											},
											Description: `OpenTelemetry backend configuration.`,
										},
										"type": schema.StringAttribute{
											Optional:    true,
											Description: `Not Null; must be one of ["Zipkin", "Datadog", "OpenTelemetry"]`,
											Validators: []validator.String{
												speakeasy_stringvalidators.NotNull(),
												stringvalidator.OneOf(
													"Zipkin",
													"Datadog",
													"OpenTelemetry",
												),
											},
										},
										"zipkin": schema.SingleNestedAttribute{
											Optional: true,
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Computed: true,
													Optional: true,
													Default:  stringdefault.StaticString(`httpJson`),
													MarkdownDescription: `Version of the API.` + "\n" +
														`https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L66` + "\n" +
														`Default: "httpJson"; must be one of ["httpJson", "httpProto"]`,
													Validators: []validator.String{
														stringvalidator.OneOf(
															"httpJson",
															"httpProto",
														),
													},
												},
												"shared_span_context": schema.BoolAttribute{
													Computed: true,
													Optional: true,
													Default:  booldefault.StaticBool(true),
													MarkdownDescription: `Determines whether client and server spans will share the same span` + "\n" +
														`context.` + "\n" +
														`https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L63` + "\n" +
														`Default: true`,
												},
												"trace_id128bit": schema.BoolAttribute{
													Computed:    true,
													Optional:    true,
													Default:     booldefault.StaticBool(false),
													Description: `Generate 128bit traces. Default: false`,
												},
												"url": schema.StringAttribute{
													Optional:    true,
													Description: `Address of Zipkin collector. Not Null`,
													Validators: []validator.String{
														speakeasy_stringvalidators.NotNull(),
													},
												},
											},
											Description: `Zipkin backend configuration.`,
										},
									},
								},
								MarkdownDescription: `A one element array of backend definition.` + "\n" +
									`Envoy allows configuring only 1 backend, so the natural way of` + "\n" +
									`representing that would be just one object. Unfortunately due to the` + "\n" +
									`reasons explained in MADR 009-tracing-policy this has to be a one element` + "\n" +
									`array for now.`,
								Validators: []validator.List{
									listvalidator.SizeAtMost(1),
								},
							},
							"sampling": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"client": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"integer": schema.Int64Attribute{
												Optional: true,
												Validators: []validator.Int64{
													int64validator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("str"),
													}...),
												},
											},
											"str": schema.StringAttribute{
												Optional: true,
												Validators: []validator.String{
													stringvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("integer"),
													}...),
												},
											},
										},
										MarkdownDescription: `Target percentage of requests that will be force traced if the` + "\n" +
											`'x-client-trace-id' header is set. Mirror of client_sampling in Envoy` + "\n" +
											`https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L127-L133` + "\n" +
											`Either int or decimal represented as string.` + "\n" +
											`If not specified then the default value is 100.`,
									},
									"overall": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"integer": schema.Int64Attribute{
												Optional: true,
												Validators: []validator.Int64{
													int64validator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("str"),
													}...),
												},
											},
											"str": schema.StringAttribute{
												Optional: true,
												Validators: []validator.String{
													stringvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("integer"),
													}...),
												},
											},
										},
										MarkdownDescription: `Target percentage of requests will be traced` + "\n" +
											`after all other sampling checks have been applied (client, force tracing,` + "\n" +
											`random sampling). This field functions as an upper limit on the total` + "\n" +
											`configured sampling rate. For instance, setting client to 100` + "\n" +
											`but overall to 1 will result in only 1% of client requests with` + "\n" +
											`the appropriate headers to be force traced. Mirror of` + "\n" +
											`overall_sampling in Envoy` + "\n" +
											`https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L142-L150` + "\n" +
											`Either int or decimal represented as string.` + "\n" +
											`If not specified then the default value is 100.`,
									},
									"random": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"integer": schema.Int64Attribute{
												Optional: true,
												Validators: []validator.Int64{
													int64validator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("str"),
													}...),
												},
											},
											"str": schema.StringAttribute{
												Optional: true,
												Validators: []validator.String{
													stringvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("integer"),
													}...),
												},
											},
										},
										MarkdownDescription: `Target percentage of requests that will be randomly selected for trace` + "\n" +
											`generation, if not requested by the client or not forced.` + "\n" +
											`Mirror of random_sampling in Envoy` + "\n" +
											`https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L135-L140` + "\n" +
											`Either int or decimal represented as string.` + "\n" +
											`If not specified then the default value is 100.`,
									},
								},
								MarkdownDescription: `Sampling configuration.` + "\n" +
									`Sampling is the process by which a decision is made on whether to` + "\n" +
									`process/export a span or not.`,
							},
							"tags": schema.ListNestedAttribute{
								Optional: true,
								PlanModifiers: []planmodifier.List{
									custom_listplanmodifier.SupressZeroNullModifier(),
								},
								NestedObject: schema.NestedAttributeObject{
									Validators: []validator.Object{
										speakeasy_objectvalidators.NotNull(),
									},
									Attributes: map[string]schema.Attribute{
										"header": schema.SingleNestedAttribute{
											Optional: true,
											Attributes: map[string]schema.Attribute{
												"default": schema.StringAttribute{
													Optional: true,
													MarkdownDescription: `Default value to use if header is missing.` + "\n" +
														`If the default is missing and there is no value the tag will not be` + "\n" +
														`included.`,
												},
												"name": schema.StringAttribute{
													Optional:    true,
													Description: `Name of the header. Not Null`,
													Validators: []validator.String{
														speakeasy_stringvalidators.NotNull(),
													},
												},
											},
											Description: `Tag taken from a header.`,
										},
										"literal": schema.StringAttribute{
											Optional:    true,
											Description: `Tag taken from literal value.`,
										},
										"name": schema.StringAttribute{
											Optional:    true,
											Description: `Name of the tag. Not Null`,
											Validators: []validator.String{
												speakeasy_stringvalidators.NotNull(),
											},
										},
									},
								},
								MarkdownDescription: `Custom tags configuration. You can add custom tags to traces based on` + "\n" +
									`headers or literal values.`,
							},
						},
						Description: `MeshTrace configuration.`,
					},
					"target_ref": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Required:    true,
								Description: `Kind of the referenced resource. must be one of ["Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshMultiZoneService", "MeshServiceSubset", "MeshHTTPRoute", "Dataplane"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Mesh",
										"MeshSubset",
										"MeshGateway",
										"MeshService",
										"MeshExternalService",
										"MeshMultiZoneService",
										"MeshServiceSubset",
										"MeshHTTPRoute",
										"Dataplane",
									),
								},
							},
							"labels": schema.MapAttribute{
								Optional:    true,
								ElementType: types.StringType,
								MarkdownDescription: `Labels are used to select group of MeshServices that match labels. Either Labels or` + "\n" +
									`Name and Namespace can be used.`,
							},
							"mesh": schema.StringAttribute{
								Optional:    true,
								Description: `Mesh is reserved for future use to identify cross mesh resources.`,
							},
							"name": schema.StringAttribute{
								Optional: true,
								MarkdownDescription: `Name of the referenced resource. Can only be used with kinds: ` + "`" + `MeshService` + "`" + `,` + "\n" +
									`` + "`" + `MeshServiceSubset` + "`" + ` and ` + "`" + `MeshGatewayRoute` + "`" + ``,
							},
							"namespace": schema.StringAttribute{
								Optional: true,
								MarkdownDescription: `Namespace specifies the namespace of target resource. If empty only resources in policy namespace` + "\n" +
									`will be targeted.`,
							},
							"proxy_types": schema.ListAttribute{
								Optional: true,
								PlanModifiers: []planmodifier.List{
									custom_listplanmodifier.SupressZeroNullModifier(),
								},
								ElementType: types.StringType,
								MarkdownDescription: `ProxyTypes specifies the data plane types that are subject to the policy. When not specified,` + "\n" +
									`all data plane types are targeted by the policy.`,
							},
							"section_name": schema.StringAttribute{
								Optional: true,
								MarkdownDescription: `SectionName is used to target specific section of resource.` + "\n" +
									`For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.`,
							},
							"tags": schema.MapAttribute{
								Optional:    true,
								ElementType: types.StringType,
								MarkdownDescription: `Tags used to select a subset of proxies by tags. Can only be used with kinds` + "\n" +
									`` + "`" + `MeshSubset` + "`" + ` and ` + "`" + `MeshServiceSubset` + "`" + ``,
							},
						},
						MarkdownDescription: `TargetRef is a reference to the resource the policy takes an effect on.` + "\n" +
							`The resource could be either a real store object or virtual resource` + "\n" +
							`defined inplace.`,
					},
				},
				Description: `Spec is the specification of the Kuma MeshTrace resource.`,
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: `the type of the resource. must be "MeshTrace"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"MeshTrace",
					),
				},
			},
			"warnings": schema.ListAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					custom_listplanmodifier.SupressZeroNullModifier(),
				},
				ElementType: types.StringType,
				MarkdownDescription: `warnings is a list of warning messages to return to the requesting Kuma API clients.` + "\n" +
					`Warning messages describe a problem the client making the API request should correct or be aware of.`,
			},
		},
	}
}

func (r *MeshTraceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MeshTraceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MeshTraceResourceModel
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

	meshTraceItem := *data.ToSharedMeshTraceItemInput()
	request := operations.CreateMeshTraceRequest{
		CpID:          cpID,
		Mesh:          mesh,
		Name:          name,
		MeshTraceItem: meshTraceItem,
	}
	res, err := r.client.MeshTrace.CreateMeshTrace(ctx, request)
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
	if !(res.MeshTraceCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshTraceCreateOrUpdateSuccessResponse(res.MeshTraceCreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var cpId1 string
	cpId1 = data.CpID.ValueString()

	var mesh1 string
	mesh1 = data.Mesh.ValueString()

	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetMeshTraceRequest{
		CpID: cpId1,
		Mesh: mesh1,
		Name: name1,
	}
	res1, err := r.client.MeshTrace.GetMeshTrace(ctx, request1)
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
	if !(res1.MeshTraceItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedMeshTraceItem(res1.MeshTraceItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshTraceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MeshTraceResourceModel
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

	request := operations.GetMeshTraceRequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshTrace.GetMeshTrace(ctx, request)
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
	if !(res.MeshTraceItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshTraceItem(res.MeshTraceItem)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshTraceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *MeshTraceResourceModel
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

	meshTraceItem := *data.ToSharedMeshTraceItemInput()
	request := operations.UpdateMeshTraceRequest{
		CpID:          cpID,
		Mesh:          mesh,
		Name:          name,
		MeshTraceItem: meshTraceItem,
	}
	res, err := r.client.MeshTrace.UpdateMeshTrace(ctx, request)
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
	if !(res.MeshTraceCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshTraceCreateOrUpdateSuccessResponse(res.MeshTraceCreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var cpId1 string
	cpId1 = data.CpID.ValueString()

	var mesh1 string
	mesh1 = data.Mesh.ValueString()

	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetMeshTraceRequest{
		CpID: cpId1,
		Mesh: mesh1,
		Name: name1,
	}
	res1, err := r.client.MeshTrace.GetMeshTrace(ctx, request1)
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
	if !(res1.MeshTraceItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedMeshTraceItem(res1.MeshTraceItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshTraceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MeshTraceResourceModel
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

	request := operations.DeleteMeshTraceRequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshTrace.DeleteMeshTrace(ctx, request)
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

func (r *MeshTraceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
