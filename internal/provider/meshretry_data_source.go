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
var _ datasource.DataSource = &MeshRetryDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshRetryDataSource{}

func NewMeshRetryDataSource() datasource.DataSource {
	return &MeshRetryDataSource{}
}

// MeshRetryDataSource is the data source implementation.
type MeshRetryDataSource struct {
	client *sdk.KonnectBeta
}

// MeshRetryDataSourceModel describes the data model.
type MeshRetryDataSourceModel struct {
	CpID             types.String              `tfsdk:"cp_id"`
	CreationTime     types.String              `tfsdk:"creation_time"`
	Labels           map[string]types.String   `tfsdk:"labels"`
	Mesh             types.String              `tfsdk:"mesh"`
	ModificationTime types.String              `tfsdk:"modification_time"`
	Name             types.String              `tfsdk:"name"`
	Spec             tfTypes.MeshRetryItemSpec `tfsdk:"spec"`
	Type             types.String              `tfsdk:"type"`
}

// Metadata returns the data source type name.
func (r *MeshRetryDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_retry"
}

// Schema defines the schema for the data source.
func (r *MeshRetryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshRetry DataSource",

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
				Description: `name of the MeshRetry`,
			},
			"spec": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
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
					"to": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"grpc": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"back_off": schema.SingleNestedAttribute{
													Computed: true,
													Attributes: map[string]schema.Attribute{
														"base_interval": schema.StringAttribute{
															Computed: true,
															MarkdownDescription: `BaseInterval is an amount of time which should be taken between retries.` + "\n" +
																`Must be greater than zero. Values less than 1 ms are rounded up to 1 ms.` + "\n" +
																`If not specified then the default value is "25ms".`,
														},
														"max_interval": schema.StringAttribute{
															Computed: true,
															MarkdownDescription: `MaxInterval is a maximal amount of time which will be taken between retries.` + "\n" +
																`Default is 10 times the "BaseInterval".`,
														},
													},
													MarkdownDescription: `BackOff is a configuration of durations which will be used in an exponential` + "\n" +
														`backoff strategy between retries.`,
												},
												"num_retries": schema.Int64Attribute{
													Computed: true,
													MarkdownDescription: `NumRetries is the number of attempts that will be made on failed (and` + "\n" +
														`retriable) requests. If not set, the default value is 1.`,
												},
												"per_try_timeout": schema.StringAttribute{
													Computed: true,
													MarkdownDescription: `PerTryTimeout is the maximum amount of time each retry attempt can take` + "\n" +
														`before it times out. If not set, the global request timeout for the route` + "\n" +
														`will be used. Setting this value to 0 will disable the per-try timeout.`,
												},
												"rate_limited_back_off": schema.SingleNestedAttribute{
													Computed: true,
													Attributes: map[string]schema.Attribute{
														"max_interval": schema.StringAttribute{
															Computed: true,
															MarkdownDescription: `MaxInterval is a maximal amount of time which will be taken between retries.` + "\n" +
																`If not specified then the default value is "300s".`,
														},
														"reset_headers": schema.ListNestedAttribute{
															Computed: true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"format": schema.StringAttribute{
																		Computed:    true,
																		Description: `The format of the reset header.`,
																	},
																	"name": schema.StringAttribute{
																		Computed:    true,
																		Description: `The Name of the reset header.`,
																	},
																},
															},
															MarkdownDescription: `ResetHeaders specifies the list of headers (like Retry-After or X-RateLimit-Reset)` + "\n" +
																`to match against the response. Headers are tried in order, and matched` + "\n" +
																`case-insensitive. The first header to be parsed successfully is used.` + "\n" +
																`If no headers match the default exponential BackOff is used instead.`,
														},
													},
													MarkdownDescription: `RateLimitedBackOff is a configuration of backoff which will be used when` + "\n" +
														`the upstream returns one of the headers configured.`,
												},
												"retry_on": schema.ListAttribute{
													Computed:    true,
													ElementType: types.StringType,
													Description: `RetryOn is a list of conditions which will cause a retry.`,
												},
											},
											Description: `GRPC defines a configuration of retries for GRPC traffic`,
										},
										"http": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"back_off": schema.SingleNestedAttribute{
													Computed: true,
													Attributes: map[string]schema.Attribute{
														"base_interval": schema.StringAttribute{
															Computed: true,
															MarkdownDescription: `BaseInterval is an amount of time which should be taken between retries.` + "\n" +
																`Must be greater than zero. Values less than 1 ms are rounded up to 1 ms.` + "\n" +
																`If not specified then the default value is "25ms".`,
														},
														"max_interval": schema.StringAttribute{
															Computed: true,
															MarkdownDescription: `MaxInterval is a maximal amount of time which will be taken between retries.` + "\n" +
																`Default is 10 times the "BaseInterval".`,
														},
													},
													MarkdownDescription: `BackOff is a configuration of durations which will be used in exponential` + "\n" +
														`backoff strategy between retries.`,
												},
												"host_selection": schema.ListNestedAttribute{
													Computed: true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"predicate": schema.StringAttribute{
																Computed:    true,
																Description: `Type is requested predicate mode.`,
															},
															"tags": schema.MapAttribute{
																Computed:    true,
																ElementType: types.StringType,
																MarkdownDescription: `Tags is a map of metadata to match against for selecting the omitted hosts. Required if Type is` + "\n" +
																	`OmitHostsWithTags`,
															},
															"update_frequency": schema.Int64Attribute{
																Computed: true,
																MarkdownDescription: `UpdateFrequency is how often the priority load should be updated based on previously attempted priorities.` + "\n" +
																	`Used for OmitPreviousPriorities.`,
															},
														},
													},
													MarkdownDescription: `HostSelection is a list of predicates that dictate how hosts should be selected` + "\n" +
														`when requests are retried.`,
												},
												"host_selection_max_attempts": schema.Int64Attribute{
													Computed: true,
													MarkdownDescription: `HostSelectionMaxAttempts is the maximum number of times host selection will be` + "\n" +
														`reattempted before giving up, at which point the host that was last selected will` + "\n" +
														`be routed to. If unspecified, this will default to retrying once.`,
												},
												"num_retries": schema.Int64Attribute{
													Computed: true,
													MarkdownDescription: `NumRetries is the number of attempts that will be made on failed (and` + "\n" +
														`retriable) requests.  If not set, the default value is 1.`,
												},
												"per_try_timeout": schema.StringAttribute{
													Computed: true,
													MarkdownDescription: `PerTryTimeout is the amount of time after which retry attempt should time out.` + "\n" +
														`If left unspecified, the global route timeout for the request will be used.` + "\n" +
														`Consequently, when using a 5xx based retry policy, a request that times out` + "\n" +
														`will not be retried as the total timeout budget would have been exhausted.` + "\n" +
														`Setting this timeout to 0 will disable it.`,
												},
												"rate_limited_back_off": schema.SingleNestedAttribute{
													Computed: true,
													Attributes: map[string]schema.Attribute{
														"max_interval": schema.StringAttribute{
															Computed: true,
															MarkdownDescription: `MaxInterval is a maximal amount of time which will be taken between retries.` + "\n" +
																`If not specified then the default value is "300s".`,
														},
														"reset_headers": schema.ListNestedAttribute{
															Computed: true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"format": schema.StringAttribute{
																		Computed:    true,
																		Description: `The format of the reset header.`,
																	},
																	"name": schema.StringAttribute{
																		Computed:    true,
																		Description: `The Name of the reset header.`,
																	},
																},
															},
															MarkdownDescription: `ResetHeaders specifies the list of headers (like Retry-After or X-RateLimit-Reset)` + "\n" +
																`to match against the response. Headers are tried in order, and matched` + "\n" +
																`case-insensitive. The first header to be parsed successfully is used.` + "\n" +
																`If no headers match the default exponential BackOff is used instead.`,
														},
													},
													MarkdownDescription: `RateLimitedBackOff is a configuration of backoff which will be used` + "\n" +
														`when the upstream returns one of the headers configured.`,
												},
												"retriable_request_headers": schema.ListNestedAttribute{
													Computed: true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `Name is the name of the HTTP Header to be matched. Name MUST be lower case` + "\n" +
																	`as they will be handled with case insensitivity (See https://tools.ietf.org/html/rfc7230#section-3.2).`,
															},
															"type": schema.StringAttribute{
																Computed:    true,
																Description: `Type specifies how to match against the value of the header.`,
															},
															"value": schema.StringAttribute{
																Computed:    true,
																Description: `Value is the value of HTTP Header to be matched.`,
															},
														},
													},
													MarkdownDescription: `RetriableRequestHeaders is an HTTP headers which must be present in the request` + "\n" +
														`for retries to be attempted.`,
												},
												"retriable_response_headers": schema.ListNestedAttribute{
													Computed: true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `Name is the name of the HTTP Header to be matched. Name MUST be lower case` + "\n" +
																	`as they will be handled with case insensitivity (See https://tools.ietf.org/html/rfc7230#section-3.2).`,
															},
															"type": schema.StringAttribute{
																Computed:    true,
																Description: `Type specifies how to match against the value of the header.`,
															},
															"value": schema.StringAttribute{
																Computed:    true,
																Description: `Value is the value of HTTP Header to be matched.`,
															},
														},
													},
													MarkdownDescription: `RetriableResponseHeaders is an HTTP response headers that trigger a retry` + "\n" +
														`if present in the response. A retry will be triggered if any of the header` + "\n" +
														`matches the upstream response headers.`,
												},
												"retry_on": schema.ListAttribute{
													Computed:    true,
													ElementType: types.StringType,
													MarkdownDescription: `RetryOn is a list of conditions which will cause a retry. Available values are:` + "\n" +
														`[5XX, GatewayError, Reset, Retriable4xx, ConnectFailure, EnvoyRatelimited,` + "\n" +
														`RefusedStream, Http3PostConnectFailure, HttpMethodConnect, HttpMethodDelete,` + "\n" +
														`HttpMethodGet, HttpMethodHead, HttpMethodOptions, HttpMethodPatch,` + "\n" +
														`HttpMethodPost, HttpMethodPut, HttpMethodTrace].` + "\n" +
														`Also, any HTTP status code (500, 503, etc.).`,
												},
											},
											Description: `HTTP defines a configuration of retries for HTTP traffic`,
										},
										"tcp": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"max_connect_attempt": schema.Int64Attribute{
													Computed: true,
													MarkdownDescription: `MaxConnectAttempt is a maximal amount of TCP connection attempts` + "\n" +
														`which will be made before giving up`,
												},
											},
											Description: `TCP defines a configuration of retries for TCP traffic`,
										},
									},
									MarkdownDescription: `Default is a configuration specific to the group of destinations referenced in` + "\n" +
										`'targetRef'`,
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
									MarkdownDescription: `TargetRef is a reference to the resource that represents a group of` + "\n" +
										`destinations.`,
								},
							},
						},
						Description: `To list makes a match between the consumed services and corresponding configurations`,
					},
				},
				Description: `Spec is the specification of the Kuma MeshRetry resource.`,
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: `the type of the resource`,
			},
		},
	}
}

func (r *MeshRetryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshRetryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshRetryDataSourceModel
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

	request := operations.GetMeshRetryRequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshRetry.GetMeshRetry(ctx, request)
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
	if !(res.MeshRetryItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshRetryItem(res.MeshRetryItem)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
