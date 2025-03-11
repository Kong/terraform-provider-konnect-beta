// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect-beta/internal/provider/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MeshTraceDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshTraceDataSource{}

func NewMeshTraceDataSource() datasource.DataSource {
	return &MeshTraceDataSource{}
}

// MeshTraceDataSource is the data source implementation.
type MeshTraceDataSource struct {
	client *sdk.KonnectBeta
}

// MeshTraceDataSourceModel describes the data model.
type MeshTraceDataSourceModel struct {
	CpID             types.String              `tfsdk:"cp_id"`
	CreationTime     types.String              `tfsdk:"creation_time"`
	Labels           map[string]types.String   `tfsdk:"labels"`
	Mesh             types.String              `tfsdk:"mesh"`
	ModificationTime types.String              `tfsdk:"modification_time"`
	Name             types.String              `tfsdk:"name"`
	Spec             tfTypes.MeshTraceItemSpec `tfsdk:"spec"`
	Type             types.String              `tfsdk:"type"`
}

// Metadata returns the data source type name.
func (r *MeshTraceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_trace"
}

// Schema defines the schema for the data source.
func (r *MeshTraceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshTrace DataSource",

		Attributes: map[string]schema.Attribute{
			"cp_id": schema.StringAttribute{
				Required:    true,
				Description: `Id of the Konnect resource`,
			},
			"creation_time": schema.StringAttribute{
				Computed:    true,
				Description: `Time at which the resource was created`,
			},
			"labels": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `The labels to help identity resources`,
			},
			"mesh": schema.StringAttribute{
				Required:    true,
				Description: `name of the mesh`,
			},
			"modification_time": schema.StringAttribute{
				Computed:    true,
				Description: `Time at which the resource was updated`,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `name of the MeshTrace`,
			},
			"spec": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"default": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"backends": schema.ListNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"datadog": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"split_service": schema.BoolAttribute{
													Computed: true,
													MarkdownDescription: `Determines if datadog service name should be split based on traffic` + "\n" +
														`direction and destination. For example, with ` + "`" + `splitService: true` + "`" + ` and a` + "\n" +
														`` + "`" + `backend` + "`" + ` service that communicates with a couple of databases, you would` + "\n" +
														`get service names like ` + "`" + `backend_INBOUND` + "`" + `, ` + "`" + `backend_OUTBOUND_db1` + "`" + `, and` + "\n" +
														`` + "`" + `backend_OUTBOUND_db2` + "`" + ` in Datadog.`,
												},
												"url": schema.StringAttribute{
													Computed: true,
													MarkdownDescription: `Address of Datadog collector, only host and port are allowed (no paths,` + "\n" +
														`fragments etc.)`,
												},
											},
											Description: `Datadog backend configuration.`,
										},
										"open_telemetry": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"endpoint": schema.StringAttribute{
													Computed:    true,
													Description: `Address of OpenTelemetry collector.`,
												},
											},
											Description: `OpenTelemetry backend configuration.`,
										},
										"type": schema.StringAttribute{
											Computed: true,
										},
										"zipkin": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Computed: true,
													MarkdownDescription: `Version of the API.` + "\n" +
														`https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L66`,
												},
												"shared_span_context": schema.BoolAttribute{
													Computed: true,
													MarkdownDescription: `Determines whether client and server spans will share the same span` + "\n" +
														`context.` + "\n" +
														`https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L63`,
												},
												"trace_id128bit": schema.BoolAttribute{
													Computed:    true,
													Description: `Generate 128bit traces.`,
												},
												"url": schema.StringAttribute{
													Computed:    true,
													Description: `Address of Zipkin collector.`,
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
							},
							"sampling": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"client": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"integer": schema.Int64Attribute{
												Computed: true,
											},
											"str": schema.StringAttribute{
												Computed: true,
											},
										},
										MarkdownDescription: `Target percentage of requests that will be force traced if the` + "\n" +
											`'x-client-trace-id' header is set. Mirror of client_sampling in Envoy` + "\n" +
											`https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L127-L133` + "\n" +
											`Either int or decimal represented as string.` + "\n" +
											`If not specified then the default value is 100.`,
									},
									"overall": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"integer": schema.Int64Attribute{
												Computed: true,
											},
											"str": schema.StringAttribute{
												Computed: true,
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
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"integer": schema.Int64Attribute{
												Computed: true,
											},
											"str": schema.StringAttribute{
												Computed: true,
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
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"header": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"default": schema.StringAttribute{
													Computed: true,
													MarkdownDescription: `Default value to use if header is missing.` + "\n" +
														`If the default is missing and there is no value the tag will not be` + "\n" +
														`included.`,
												},
												"name": schema.StringAttribute{
													Computed:    true,
													Description: `Name of the header.`,
												},
											},
											Description: `Tag taken from a header.`,
										},
										"literal": schema.StringAttribute{
											Computed:    true,
											Description: `Tag taken from literal value.`,
										},
										"name": schema.StringAttribute{
											Computed:    true,
											Description: `Name of the tag.`,
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
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Computed:    true,
								Description: `Kind of the referenced resource`,
							},
							"labels": schema.MapAttribute{
								Computed:    true,
								ElementType: types.StringType,
								MarkdownDescription: `Labels are used to select group of MeshServices that match labels. Either Labels or` + "\n" +
									`Name and Namespace can be used.`,
							},
							"mesh": schema.StringAttribute{
								Computed:    true,
								Description: `Mesh is reserved for future use to identify cross mesh resources.`,
							},
							"name": schema.StringAttribute{
								Computed: true,
								MarkdownDescription: `Name of the referenced resource. Can only be used with kinds: ` + "`" + `MeshService` + "`" + `,` + "\n" +
									`` + "`" + `MeshServiceSubset` + "`" + ` and ` + "`" + `MeshGatewayRoute` + "`" + ``,
							},
							"namespace": schema.StringAttribute{
								Computed: true,
								MarkdownDescription: `Namespace specifies the namespace of target resource. If empty only resources in policy namespace` + "\n" +
									`will be targeted.`,
							},
							"proxy_types": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
								MarkdownDescription: `ProxyTypes specifies the data plane types that are subject to the policy. When not specified,` + "\n" +
									`all data plane types are targeted by the policy.`,
							},
							"section_name": schema.StringAttribute{
								Computed: true,
								MarkdownDescription: `SectionName is used to target specific section of resource.` + "\n" +
									`For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.`,
							},
							"tags": schema.MapAttribute{
								Computed:    true,
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
				Computed:    true,
				Description: `the type of the resource`,
			},
		},
	}
}

func (r *MeshTraceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshTraceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshTraceDataSourceModel
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
