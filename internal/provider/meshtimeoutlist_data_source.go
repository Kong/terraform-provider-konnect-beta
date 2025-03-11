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
var _ datasource.DataSource = &MeshTimeoutListDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshTimeoutListDataSource{}

func NewMeshTimeoutListDataSource() datasource.DataSource {
	return &MeshTimeoutListDataSource{}
}

// MeshTimeoutListDataSource is the data source implementation.
type MeshTimeoutListDataSource struct {
	client *sdk.KonnectBeta
}

// MeshTimeoutListDataSourceModel describes the data model.
type MeshTimeoutListDataSourceModel struct {
	CpID   types.String              `tfsdk:"cp_id"`
	Items  []tfTypes.MeshTimeoutItem `tfsdk:"items"`
	Key    types.String              `queryParam:"name=key" tfsdk:"key"`
	Mesh   types.String              `tfsdk:"mesh"`
	Next   types.String              `tfsdk:"next"`
	Offset types.Int64               `queryParam:"style=form,explode=true,name=offset" tfsdk:"offset"`
	Size   types.Int64               `queryParam:"style=form,explode=true,name=size" tfsdk:"size"`
	Total  types.Number              `tfsdk:"total"`
	Value  types.String              `queryParam:"name=value" tfsdk:"value"`
}

// Metadata returns the data source type name.
func (r *MeshTimeoutListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_timeout_list"
}

// Schema defines the schema for the data source.
func (r *MeshTimeoutListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshTimeoutList DataSource",

		Attributes: map[string]schema.Attribute{
			"cp_id": schema.StringAttribute{
				Required:    true,
				Description: `Id of the Konnect resource`,
			},
			"items": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
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
							Computed:    true,
							Description: `Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.`,
						},
						"modification_time": schema.StringAttribute{
							Computed:    true,
							Description: `Time at which the resource was updated`,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: `Name of the Kuma resource`,
						},
						"spec": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"from": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"default": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"connection_timeout": schema.StringAttribute{
														Computed: true,
														MarkdownDescription: `ConnectionTimeout specifies the amount of time proxy will wait for an TCP connection to be established.` + "\n" +
															`Default value is 5 seconds. Cannot be set to 0.`,
													},
													"http": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"max_connection_duration": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `MaxConnectionDuration is the time after which a connection will be drained and/or closed,` + "\n" +
																	`starting from when it was first established. Setting this timeout to 0 will disable it.` + "\n" +
																	`Disabled by default.`,
															},
															"max_stream_duration": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `MaxStreamDuration is the maximum time that a stream’s lifetime will span.` + "\n" +
																	`Setting this timeout to 0 will disable it. Disabled by default.`,
															},
															"request_headers_timeout": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `RequestHeadersTimeout The amount of time that proxy will wait for the request headers to be received. The timer is` + "\n" +
																	`activated when the first byte of the headers is received, and is disarmed when the last byte of` + "\n" +
																	`the headers has been received. If not specified or set to 0, this timeout is disabled.` + "\n" +
																	`Disabled by default.`,
															},
															"request_timeout": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `RequestTimeout The amount of time that proxy will wait for the entire request to be received.` + "\n" +
																	`The timer is activated when the request is initiated, and is disarmed when the last byte of the request is sent,` + "\n" +
																	`OR when the response is initiated. Setting this timeout to 0 will disable it.` + "\n" +
																	`Default is 15s.`,
															},
															"stream_idle_timeout": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `StreamIdleTimeout is the amount of time that proxy will allow a stream to exist with no activity.` + "\n" +
																	`Setting this timeout to 0 will disable it. Default is 30m`,
															},
														},
														Description: `Http provides configuration for HTTP specific timeouts`,
													},
													"idle_timeout": schema.StringAttribute{
														Computed: true,
														MarkdownDescription: `IdleTimeout is defined as the period in which there are no bytes sent or received on connection` + "\n" +
															`Setting this timeout to 0 will disable it. Be cautious when disabling it because` + "\n" +
															`it can lead to connection leaking. Default value is 1h.`,
													},
												},
												MarkdownDescription: `Default is a configuration specific to the group of clients referenced in` + "\n" +
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
													`clients.`,
											},
										},
									},
									Description: `From list makes a match between clients and corresponding configurations`,
								},
								"rules": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"default": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"connection_timeout": schema.StringAttribute{
														Computed: true,
														MarkdownDescription: `ConnectionTimeout specifies the amount of time proxy will wait for an TCP connection to be established.` + "\n" +
															`Default value is 5 seconds. Cannot be set to 0.`,
													},
													"http": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"max_connection_duration": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `MaxConnectionDuration is the time after which a connection will be drained and/or closed,` + "\n" +
																	`starting from when it was first established. Setting this timeout to 0 will disable it.` + "\n" +
																	`Disabled by default.`,
															},
															"max_stream_duration": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `MaxStreamDuration is the maximum time that a stream’s lifetime will span.` + "\n" +
																	`Setting this timeout to 0 will disable it. Disabled by default.`,
															},
															"request_headers_timeout": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `RequestHeadersTimeout The amount of time that proxy will wait for the request headers to be received. The timer is` + "\n" +
																	`activated when the first byte of the headers is received, and is disarmed when the last byte of` + "\n" +
																	`the headers has been received. If not specified or set to 0, this timeout is disabled.` + "\n" +
																	`Disabled by default.`,
															},
															"request_timeout": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `RequestTimeout The amount of time that proxy will wait for the entire request to be received.` + "\n" +
																	`The timer is activated when the request is initiated, and is disarmed when the last byte of the request is sent,` + "\n" +
																	`OR when the response is initiated. Setting this timeout to 0 will disable it.` + "\n" +
																	`Default is 15s.`,
															},
															"stream_idle_timeout": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `StreamIdleTimeout is the amount of time that proxy will allow a stream to exist with no activity.` + "\n" +
																	`Setting this timeout to 0 will disable it. Default is 30m`,
															},
														},
														Description: `Http provides configuration for HTTP specific timeouts`,
													},
													"idle_timeout": schema.StringAttribute{
														Computed: true,
														MarkdownDescription: `IdleTimeout is defined as the period in which there are no bytes sent or received on connection` + "\n" +
															`Setting this timeout to 0 will disable it. Be cautious when disabling it because` + "\n" +
															`it can lead to connection leaking. Default value is 1h.`,
													},
												},
												Description: `Default contains configuration of the inbound timeouts`,
											},
										},
									},
									MarkdownDescription: `Rules defines inbound timeout configurations. Currently limited to exactly one rule containing` + "\n" +
										`default timeouts that apply to all inbound traffic, as L7 matching is not yet implemented.`,
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
								"to": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"default": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"connection_timeout": schema.StringAttribute{
														Computed: true,
														MarkdownDescription: `ConnectionTimeout specifies the amount of time proxy will wait for an TCP connection to be established.` + "\n" +
															`Default value is 5 seconds. Cannot be set to 0.`,
													},
													"http": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"max_connection_duration": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `MaxConnectionDuration is the time after which a connection will be drained and/or closed,` + "\n" +
																	`starting from when it was first established. Setting this timeout to 0 will disable it.` + "\n" +
																	`Disabled by default.`,
															},
															"max_stream_duration": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `MaxStreamDuration is the maximum time that a stream’s lifetime will span.` + "\n" +
																	`Setting this timeout to 0 will disable it. Disabled by default.`,
															},
															"request_headers_timeout": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `RequestHeadersTimeout The amount of time that proxy will wait for the request headers to be received. The timer is` + "\n" +
																	`activated when the first byte of the headers is received, and is disarmed when the last byte of` + "\n" +
																	`the headers has been received. If not specified or set to 0, this timeout is disabled.` + "\n" +
																	`Disabled by default.`,
															},
															"request_timeout": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `RequestTimeout The amount of time that proxy will wait for the entire request to be received.` + "\n" +
																	`The timer is activated when the request is initiated, and is disarmed when the last byte of the request is sent,` + "\n" +
																	`OR when the response is initiated. Setting this timeout to 0 will disable it.` + "\n" +
																	`Default is 15s.`,
															},
															"stream_idle_timeout": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `StreamIdleTimeout is the amount of time that proxy will allow a stream to exist with no activity.` + "\n" +
																	`Setting this timeout to 0 will disable it. Default is 30m`,
															},
														},
														Description: `Http provides configuration for HTTP specific timeouts`,
													},
													"idle_timeout": schema.StringAttribute{
														Computed: true,
														MarkdownDescription: `IdleTimeout is defined as the period in which there are no bytes sent or received on connection` + "\n" +
															`Setting this timeout to 0 will disable it. Be cautious when disabling it because` + "\n" +
															`it can lead to connection leaking. Default value is 1h.`,
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
							Description: `Spec is the specification of the Kuma MeshTimeout resource.`,
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: `the type of the resource`,
						},
					},
				},
			},
			"key": schema.StringAttribute{
				Optional: true,
			},
			"mesh": schema.StringAttribute{
				Required:    true,
				Description: `name of the mesh`,
			},
			"next": schema.StringAttribute{
				Computed:    true,
				Description: `URL to the next page`,
			},
			"offset": schema.Int64Attribute{
				Optional:    true,
				Description: `offset in the list of entities`,
			},
			"size": schema.Int64Attribute{
				Optional:    true,
				Description: `the number of items per page`,
			},
			"total": schema.NumberAttribute{
				Computed:    true,
				Description: `The total number of entities`,
			},
			"value": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *MeshTimeoutListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshTimeoutListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshTimeoutListDataSourceModel
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

	offset := new(int64)
	if !data.Offset.IsUnknown() && !data.Offset.IsNull() {
		*offset = data.Offset.ValueInt64()
	} else {
		offset = nil
	}
	size := new(int64)
	if !data.Size.IsUnknown() && !data.Size.IsNull() {
		*size = data.Size.ValueInt64()
	} else {
		size = nil
	}
	var filter *operations.GetMeshTimeoutListQueryParamFilter
	key := new(string)
	if !data.Key.IsUnknown() && !data.Key.IsNull() {
		*key = data.Key.ValueString()
	} else {
		key = nil
	}
	value := new(string)
	if !data.Value.IsUnknown() && !data.Value.IsNull() {
		*value = data.Value.ValueString()
	} else {
		value = nil
	}
	filter = &operations.GetMeshTimeoutListQueryParamFilter{
		Key:   key,
		Value: value,
	}
	var mesh string
	mesh = data.Mesh.ValueString()

	request := operations.GetMeshTimeoutListRequest{
		CpID:   cpID,
		Offset: offset,
		Size:   size,
		Filter: filter,
		Mesh:   mesh,
	}
	res, err := r.client.MeshTimeout.GetMeshTimeoutList(ctx, request)
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
	if !(res.MeshTimeoutList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshTimeoutList(res.MeshTimeoutList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
