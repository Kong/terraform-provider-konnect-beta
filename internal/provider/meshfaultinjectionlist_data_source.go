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
var _ datasource.DataSource = &MeshFaultInjectionListDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshFaultInjectionListDataSource{}

func NewMeshFaultInjectionListDataSource() datasource.DataSource {
	return &MeshFaultInjectionListDataSource{}
}

// MeshFaultInjectionListDataSource is the data source implementation.
type MeshFaultInjectionListDataSource struct {
	client *sdk.KonnectBeta
}

// MeshFaultInjectionListDataSourceModel describes the data model.
type MeshFaultInjectionListDataSourceModel struct {
	CpID   types.String                     `tfsdk:"cp_id"`
	Items  []tfTypes.MeshFaultInjectionItem `tfsdk:"items"`
	Key    types.String                     `queryParam:"name=key" tfsdk:"key"`
	Mesh   types.String                     `tfsdk:"mesh"`
	Next   types.String                     `tfsdk:"next"`
	Offset types.Int64                      `queryParam:"style=form,explode=true,name=offset" tfsdk:"offset"`
	Size   types.Int64                      `queryParam:"style=form,explode=true,name=size" tfsdk:"size"`
	Total  types.Number                     `tfsdk:"total"`
	Value  types.String                     `queryParam:"name=value" tfsdk:"value"`
}

// Metadata returns the data source type name.
func (r *MeshFaultInjectionListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_fault_injection_list"
}

// Schema defines the schema for the data source.
func (r *MeshFaultInjectionListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshFaultInjectionList DataSource",

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
													"http": schema.ListNestedAttribute{
														Computed: true,
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"abort": schema.SingleNestedAttribute{
																	Computed: true,
																	Attributes: map[string]schema.Attribute{
																		"http_status": schema.Int32Attribute{
																			Computed:    true,
																			Description: `HTTP status code which will be returned to source side`,
																		},
																		"percentage": schema.SingleNestedAttribute{
																			Computed: true,
																			Attributes: map[string]schema.Attribute{
																				"integer": schema.Int64Attribute{
																					Computed: true,
																				},
																				"str": schema.StringAttribute{
																					Computed: true,
																				},
																			},
																			MarkdownDescription: `Percentage of requests on which abort will be injected, has to be` + "\n" +
																				`either int or decimal represented as string.`,
																		},
																	},
																	MarkdownDescription: `Abort defines a configuration of not delivering requests to destination` + "\n" +
																		`service and replacing the responses from destination dataplane by` + "\n" +
																		`predefined status code`,
																},
																"delay": schema.SingleNestedAttribute{
																	Computed: true,
																	Attributes: map[string]schema.Attribute{
																		"percentage": schema.SingleNestedAttribute{
																			Computed: true,
																			Attributes: map[string]schema.Attribute{
																				"integer": schema.Int64Attribute{
																					Computed: true,
																				},
																				"str": schema.StringAttribute{
																					Computed: true,
																				},
																			},
																			MarkdownDescription: `Percentage of requests on which delay will be injected, has to be` + "\n" +
																				`either int or decimal represented as string.`,
																		},
																		"value": schema.StringAttribute{
																			Computed:    true,
																			Description: `The duration during which the response will be delayed`,
																		},
																	},
																	Description: `Delay defines configuration of delaying a response from a destination`,
																},
																"response_bandwidth": schema.SingleNestedAttribute{
																	Computed: true,
																	Attributes: map[string]schema.Attribute{
																		"limit": schema.StringAttribute{
																			Computed: true,
																			MarkdownDescription: `Limit is represented by value measure in Gbps, Mbps, kbps, e.g.` + "\n" +
																				`10kbps`,
																		},
																		"percentage": schema.SingleNestedAttribute{
																			Computed: true,
																			Attributes: map[string]schema.Attribute{
																				"integer": schema.Int64Attribute{
																					Computed: true,
																				},
																				"str": schema.StringAttribute{
																					Computed: true,
																				},
																			},
																			MarkdownDescription: `Percentage of requests on which response bandwidth limit will be` + "\n" +
																				`either int or decimal represented as string.`,
																		},
																	},
																	MarkdownDescription: `ResponseBandwidth defines a configuration to limit the speed of` + "\n" +
																		`responding to the requests`,
																},
															},
														},
														Description: `Http allows to define list of Http faults between dataplanes.`,
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
									Description: `From list makes a match between clients and corresponding configurations`,
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
													"http": schema.ListNestedAttribute{
														Computed: true,
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"abort": schema.SingleNestedAttribute{
																	Computed: true,
																	Attributes: map[string]schema.Attribute{
																		"http_status": schema.Int32Attribute{
																			Computed:    true,
																			Description: `HTTP status code which will be returned to source side`,
																		},
																		"percentage": schema.SingleNestedAttribute{
																			Computed: true,
																			Attributes: map[string]schema.Attribute{
																				"integer": schema.Int64Attribute{
																					Computed: true,
																				},
																				"str": schema.StringAttribute{
																					Computed: true,
																				},
																			},
																			MarkdownDescription: `Percentage of requests on which abort will be injected, has to be` + "\n" +
																				`either int or decimal represented as string.`,
																		},
																	},
																	MarkdownDescription: `Abort defines a configuration of not delivering requests to destination` + "\n" +
																		`service and replacing the responses from destination dataplane by` + "\n" +
																		`predefined status code`,
																},
																"delay": schema.SingleNestedAttribute{
																	Computed: true,
																	Attributes: map[string]schema.Attribute{
																		"percentage": schema.SingleNestedAttribute{
																			Computed: true,
																			Attributes: map[string]schema.Attribute{
																				"integer": schema.Int64Attribute{
																					Computed: true,
																				},
																				"str": schema.StringAttribute{
																					Computed: true,
																				},
																			},
																			MarkdownDescription: `Percentage of requests on which delay will be injected, has to be` + "\n" +
																				`either int or decimal represented as string.`,
																		},
																		"value": schema.StringAttribute{
																			Computed:    true,
																			Description: `The duration during which the response will be delayed`,
																		},
																	},
																	Description: `Delay defines configuration of delaying a response from a destination`,
																},
																"response_bandwidth": schema.SingleNestedAttribute{
																	Computed: true,
																	Attributes: map[string]schema.Attribute{
																		"limit": schema.StringAttribute{
																			Computed: true,
																			MarkdownDescription: `Limit is represented by value measure in Gbps, Mbps, kbps, e.g.` + "\n" +
																				`10kbps`,
																		},
																		"percentage": schema.SingleNestedAttribute{
																			Computed: true,
																			Attributes: map[string]schema.Attribute{
																				"integer": schema.Int64Attribute{
																					Computed: true,
																				},
																				"str": schema.StringAttribute{
																					Computed: true,
																				},
																			},
																			MarkdownDescription: `Percentage of requests on which response bandwidth limit will be` + "\n" +
																				`either int or decimal represented as string.`,
																		},
																	},
																	MarkdownDescription: `ResponseBandwidth defines a configuration to limit the speed of` + "\n" +
																		`responding to the requests`,
																},
															},
														},
														Description: `Http allows to define list of Http faults between dataplanes.`,
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
									Description: `To list makes a match between clients and corresponding configurations`,
								},
							},
							Description: `Spec is the specification of the Kuma MeshFaultInjection resource.`,
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

func (r *MeshFaultInjectionListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshFaultInjectionListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshFaultInjectionListDataSourceModel
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
	var filter *operations.GetMeshFaultInjectionListQueryParamFilter
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
	filter = &operations.GetMeshFaultInjectionListQueryParamFilter{
		Key:   key,
		Value: value,
	}
	var mesh string
	mesh = data.Mesh.ValueString()

	request := operations.GetMeshFaultInjectionListRequest{
		CpID:   cpID,
		Offset: offset,
		Size:   size,
		Filter: filter,
		Mesh:   mesh,
	}
	res, err := r.client.MeshFaultInjection.GetMeshFaultInjectionList(ctx, request)
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
	if !(res.MeshFaultInjectionList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshFaultInjectionList(res.MeshFaultInjectionList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
