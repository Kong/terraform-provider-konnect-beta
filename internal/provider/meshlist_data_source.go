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
var _ datasource.DataSource = &MeshListDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshListDataSource{}

func NewMeshListDataSource() datasource.DataSource {
	return &MeshListDataSource{}
}

// MeshListDataSource is the data source implementation.
type MeshListDataSource struct {
	client *sdk.KonnectBeta
}

// MeshListDataSourceModel describes the data model.
type MeshListDataSourceModel struct {
	CpID   types.String       `tfsdk:"cp_id"`
	Items  []tfTypes.MeshItem `tfsdk:"items"`
	Key    types.String       `tfsdk:"key"`
	Next   types.String       `tfsdk:"next"`
	Offset types.Int64        `tfsdk:"offset"`
	Size   types.Int64        `tfsdk:"size"`
	Total  types.Number       `tfsdk:"total"`
	Value  types.String       `tfsdk:"value"`
}

// Metadata returns the data source type name.
func (r *MeshListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_list"
}

// Schema defines the schema for the data source.
func (r *MeshListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshList DataSource",

		Attributes: map[string]schema.Attribute{
			"cp_id": schema.StringAttribute{
				Required:    true,
				Description: `Id of the Konnect resource`,
			},
			"items": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"constraints": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"dataplane_proxy": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"requirements": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"tags": schema.MapAttribute{
														Computed:    true,
														ElementType: types.StringType,
														MarkdownDescription: `Tags defines set of required tags. You can specify '*' in value to` + "\n" +
															`require non empty value of tag`,
													},
												},
											},
											MarkdownDescription: `Requirements defines a set of requirements that data plane proxies must` + "\n" +
												`fulfill in order to join the mesh. A data plane proxy must fulfill at` + "\n" +
												`least one requirement in order to join the mesh. Empty list of allowed` + "\n" +
												`requirements means that any proxy that is not explicitly denied can join.`,
										},
										"restrictions": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"tags": schema.MapAttribute{
														Computed:    true,
														ElementType: types.StringType,
														MarkdownDescription: `Tags defines set of required tags. You can specify '*' in value to` + "\n" +
															`require non empty value of tag`,
													},
												},
											},
											MarkdownDescription: `Restrictions defines a set of restrictions that data plane proxies cannot` + "\n" +
												`fulfill in order to join the mesh. A data plane proxy cannot fulfill any` + "\n" +
												`requirement in order to join the mesh.` + "\n" +
												`Restrictions takes precedence over requirements.`,
										},
									},
									MarkdownDescription: `DataplaneProxyMembership defines a set of requirements for data plane` + "\n" +
										`proxies to be a member of the mesh.`,
								},
							},
							Description: `Constraints that applies to the mesh and its entities`,
						},
						"labels": schema.MapAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"logging": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"backends": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"conf": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"file_logging_backend_config": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"path": schema.StringAttribute{
																Computed:    true,
																Description: `Path to a file that logs will be written to`,
															},
														},
													},
													"tcp_logging_backend_config": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"address": schema.StringAttribute{
																Computed:    true,
																Description: `Address to TCP service that will receive logs`,
															},
														},
													},
												},
											},
											"format": schema.StringAttribute{
												Computed: true,
												MarkdownDescription: `Format of access logs. Placeholders available on` + "\n" +
													`https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log`,
											},
											"name": schema.StringAttribute{
												Computed: true,
												MarkdownDescription: `Name of the backend, can be then used in Mesh.logging.defaultBackend or in` + "\n" +
													`TrafficLogging`,
											},
											"type": schema.StringAttribute{
												Computed:    true,
												Description: `Type of the backend (Kuma ships with 'tcp' and 'file')`,
											},
										},
									},
									Description: `List of available logging backends`,
								},
								"default_backend": schema.StringAttribute{
									Computed:    true,
									Description: `Name of the default backend`,
								},
							},
							MarkdownDescription: `Logging settings.` + "\n" +
								`+optional`,
						},
						"mesh_services": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"mode": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"integer": schema.Int64Attribute{
											Computed: true,
										},
										"str": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
						},
						"metrics": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"backends": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"conf": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"prometheus_metrics_backend_config": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"aggregate": schema.ListNestedAttribute{
																Computed: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"address": schema.StringAttribute{
																			Computed:    true,
																			Description: `Address on which a service expose HTTP endpoint with Prometheus metrics.`,
																		},
																		"enabled": schema.BoolAttribute{
																			Computed: true,
																			MarkdownDescription: `If false then the application won't be scrapped. If nil, then it is treated` + "\n" +
																				`as true and kuma-dp scrapes metrics from the service.`,
																		},
																		"name": schema.StringAttribute{
																			Computed:    true,
																			Description: `Name which identify given configuration.`,
																		},
																		"path": schema.StringAttribute{
																			Computed:    true,
																			Description: `Path on which a service expose HTTP endpoint with Prometheus metrics.`,
																		},
																		"port": schema.Int64Attribute{
																			Computed:    true,
																			Description: `Port on which a service expose HTTP endpoint with Prometheus metrics.`,
																		},
																	},
																},
																MarkdownDescription: `Map with the configuration of applications which metrics are going to be` + "\n" +
																	`scrapped by kuma-dp.`,
															},
															"envoy": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"filter_regex": schema.StringAttribute{
																		Computed: true,
																		MarkdownDescription: `FilterRegex value that is going to be passed to Envoy for filtering` + "\n" +
																			`Envoy metrics.`,
																	},
																	"used_only": schema.BoolAttribute{
																		Computed: true,
																		MarkdownDescription: `If true then return metrics that Envoy has updated (counters incremented` + "\n" +
																			`at least once, gauges changed at least once, and histograms added to at` + "\n" +
																			`least once). If nil, then it is treated as false.`,
																	},
																},
																Description: `Configuration of Envoy's metrics.`,
															},
															"path": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `Path on which a dataplane should expose HTTP endpoint with Prometheus` + "\n" +
																	`metrics.`,
															},
															"port": schema.Int64Attribute{
																Computed: true,
																MarkdownDescription: `Port on which a dataplane should expose HTTP endpoint with Prometheus` + "\n" +
																	`metrics.`,
															},
															"skip_mtls": schema.BoolAttribute{
																Computed: true,
																MarkdownDescription: `If true then endpoints for scraping metrics won't require mTLS even if mTLS` + "\n" +
																	`is enabled in Mesh. If nil, then it is treated as false.`,
															},
															"tags": schema.MapAttribute{
																Computed:    true,
																ElementType: types.StringType,
																MarkdownDescription: `Tags associated with an application this dataplane is deployed next to,` + "\n" +
																	`e.g. service=web, version=1.0.` + "\n" +
																	`` + "`" + `service` + "`" + ` tag is mandatory.`,
															},
															"tls": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"mode": schema.SingleNestedAttribute{
																		Computed: true,
																		Attributes: map[string]schema.Attribute{
																			"integer": schema.Int64Attribute{
																				Computed: true,
																			},
																			"str": schema.StringAttribute{
																				Computed: true,
																			},
																		},
																		MarkdownDescription: `mode defines how configured is the TLS for Prometheus.` + "\n" +
																			`Supported values, delegated, disabled, activeMTLSBackend. Default to` + "\n" +
																			`` + "`" + `activeMTLSBackend` + "`" + `.`,
																	},
																},
																Description: `Configuration of TLS for prometheus listener.`,
															},
														},
													},
												},
											},
											"name": schema.StringAttribute{
												Computed:    true,
												Description: `Name of the backend, can be then used in Mesh.metrics.enabledBackend`,
											},
											"type": schema.StringAttribute{
												Computed:    true,
												Description: `Type of the backend (Kuma ships with 'prometheus')`,
											},
										},
									},
									Description: `List of available Metrics backends`,
								},
								"enabled_backend": schema.StringAttribute{
									Computed:    true,
									Description: `Name of the enabled backend`,
								},
							},
							MarkdownDescription: `Configuration for metrics collected and exposed by dataplanes.` + "\n" +
								`` + "\n" +
								`Settings defined here become defaults for every dataplane in a given Mesh.` + "\n" +
								`Additionally, it is also possible to further customize this configuration` + "\n" +
								`for each dataplane individually using Dataplane resource.` + "\n" +
								`+optional`,
						},
						"mtls": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"backends": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"conf": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"acm_certificate_authority_config": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"arn": schema.StringAttribute{
																Computed: true,
															},
															"auth": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"aws_credentials": schema.SingleNestedAttribute{
																		Computed: true,
																		Attributes: map[string]schema.Attribute{
																			"access_key": schema.SingleNestedAttribute{
																				Computed: true,
																				Attributes: map[string]schema.Attribute{
																					"type": schema.StringAttribute{
																						Computed:    true,
																						Description: `Parsed as JSON.`,
																					},
																				},
																			},
																			"access_key_secret": schema.SingleNestedAttribute{
																				Computed: true,
																				Attributes: map[string]schema.Attribute{
																					"type": schema.StringAttribute{
																						Computed:    true,
																						Description: `Parsed as JSON.`,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"ca_cert": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"type": schema.StringAttribute{
																		Computed:    true,
																		Description: `Parsed as JSON.`,
																	},
																},
															},
															"common_name": schema.StringAttribute{
																Computed: true,
															},
														},
													},
													"builtin_certificate_authority_config": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"ca_cert": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"expiration": schema.StringAttribute{
																		Computed: true,
																	},
																	"rs_abits": schema.Int64Attribute{
																		Computed: true,
																	},
																},
															},
														},
													},
													"cert_manager_certificate_authority_config": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"ca_cert": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"type": schema.StringAttribute{
																		Computed:    true,
																		Description: `Parsed as JSON.`,
																	},
																},
															},
															"common_name": schema.StringAttribute{
																Computed: true,
															},
															"dns_names": schema.ListAttribute{
																Computed:    true,
																ElementType: types.StringType,
															},
															"issuer_ref": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"group": schema.StringAttribute{
																		Computed: true,
																	},
																	"kind": schema.StringAttribute{
																		Computed: true,
																	},
																	"name": schema.StringAttribute{
																		Computed: true,
																	},
																},
															},
														},
													},
													"provided_certificate_authority_config": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"cert": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"type": schema.StringAttribute{
																		Computed:    true,
																		Description: `Parsed as JSON.`,
																	},
																},
															},
															"key": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"type": schema.StringAttribute{
																		Computed:    true,
																		Description: `Parsed as JSON.`,
																	},
																},
															},
														},
													},
													"vault_certificate_authority_config": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"mode": schema.StringAttribute{
																Computed:    true,
																Description: `Parsed as JSON.`,
															},
														},
													},
												},
											},
											"dp_cert": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"request_timeout": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"nanos": schema.Int64Attribute{
																Computed: true,
															},
															"seconds": schema.Int64Attribute{
																Computed: true,
															},
														},
														Description: `Timeout on request to CA for DP certificate generation and retrieval`,
													},
													"rotation": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"expiration": schema.StringAttribute{
																Computed:    true,
																Description: `Time after which generated certificate for Dataplane will expire`,
															},
														},
														Description: `Rotation settings`,
													},
												},
												Description: `Dataplane certificate settings`,
											},
											"mode": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"integer": schema.Int64Attribute{
														Computed: true,
													},
													"str": schema.StringAttribute{
														Computed: true,
													},
												},
												MarkdownDescription: `Mode defines the behaviour of inbound listeners with regard to traffic` + "\n" +
													`encryption`,
											},
											"name": schema.StringAttribute{
												Computed:    true,
												Description: `Name of the backend`,
											},
											"root_chain": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"request_timeout": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"nanos": schema.Int64Attribute{
																Computed: true,
															},
															"seconds": schema.Int64Attribute{
																Computed: true,
															},
														},
														MarkdownDescription: `Timeout on request for to CA for root certificate chain.` + "\n" +
															`If not specified, defaults to 10s.`,
													},
												},
											},
											"type": schema.StringAttribute{
												Computed: true,
												MarkdownDescription: `Type of the backend. Has to be one of the loaded plugins (Kuma ships with` + "\n" +
													`builtin and provided)`,
											},
										},
									},
									Description: `List of available Certificate Authority backends`,
								},
								"enabled_backend": schema.StringAttribute{
									Computed:    true,
									Description: `Name of the enabled backend`,
								},
								"skip_validation": schema.BoolAttribute{
									Computed:    true,
									Description: `If enabled, skips CA validation.`,
								},
							},
							MarkdownDescription: `mTLS settings.` + "\n" +
								`+optional`,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"networking": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"outbound": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"passthrough": schema.BoolAttribute{
											Computed:    true,
											Description: `Control the passthrough cluster`,
										},
									},
									Description: `Outbound settings`,
								},
							},
							Description: `Networking settings of the mesh`,
						},
						"routing": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"default_forbid_mesh_external_service_access": schema.BoolAttribute{
									Computed: true,
									MarkdownDescription: `If true, blocks traffic to MeshExternalServices.` + "\n" +
										`Default: false`,
								},
								"locality_aware_load_balancing": schema.BoolAttribute{
									Computed:    true,
									Description: `Enable the Locality Aware Load Balancing`,
								},
								"zone_egress": schema.BoolAttribute{
									Computed: true,
									MarkdownDescription: `Enable routing traffic to services in other zone or external services` + "\n" +
										`through ZoneEgress. Default: false`,
								},
							},
							Description: `Routing settings of the mesh`,
						},
						"skip_creating_initial_policies": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							MarkdownDescription: `List of policies to skip creating by default when the mesh is created.` + "\n" +
								`e.g. TrafficPermission, MeshRetry, etc. An '*' can be used to skip all` + "\n" +
								`policies.`,
						},
						"tracing": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"backends": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"conf": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"datadog_tracing_backend_config": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"address": schema.StringAttribute{
																Computed:    true,
																Description: `Address of datadog collector.`,
															},
															"port": schema.Int64Attribute{
																Computed:    true,
																Description: `Port of datadog collector`,
															},
															"split_service": schema.BoolAttribute{
																Computed: true,
																MarkdownDescription: `Determines if datadog service name should be split based on traffic` + "\n" +
																	`direction and destination. For example, with ` + "`" + `splitService: true` + "`" + ` and a` + "\n" +
																	`` + "`" + `backend` + "`" + ` service that communicates with a couple of databases, you would` + "\n" +
																	`get service names like ` + "`" + `backend_INBOUND` + "`" + `, ` + "`" + `backend_OUTBOUND_db1` + "`" + `, and` + "\n" +
																	`` + "`" + `backend_OUTBOUND_db2` + "`" + ` in Datadog. Default: false`,
															},
														},
													},
													"zipkin_tracing_backend_config": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"api_version": schema.StringAttribute{
																Computed: true,
																MarkdownDescription: `Version of the API. values: httpJson, httpJsonV1, httpProto. Default:` + "\n" +
																	`httpJson see` + "\n" +
																	`https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/trace/v3/trace.proto#envoy-v3-api-enum-config-trace-v3-zipkinconfig-collectorendpointversion`,
															},
															"shared_span_context": schema.BoolAttribute{
																Computed: true,
																MarkdownDescription: `Determines whether client and server spans will share the same span` + "\n" +
																	`context. Default: true.` + "\n" +
																	`https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/trace/v3/zipkin.proto#config-trace-v3-zipkinconfig`,
															},
															"trace_id128bit": schema.BoolAttribute{
																Computed:    true,
																Description: `Generate 128bit traces. Default: false`,
															},
															"url": schema.StringAttribute{
																Computed:    true,
																Description: `Address of Zipkin collector.`,
															},
														},
													},
												},
											},
											"name": schema.StringAttribute{
												Computed: true,
												MarkdownDescription: `Name of the backend, can be then used in Mesh.tracing.defaultBackend or in` + "\n" +
													`TrafficTrace`,
											},
											"sampling": schema.NumberAttribute{
												Computed: true,
												MarkdownDescription: `Percentage of traces that will be sent to the backend (range 0.0 - 100.0).` + "\n" +
													`Empty value defaults to 100.0%`,
											},
											"type": schema.StringAttribute{
												Computed:    true,
												Description: `Type of the backend (Kuma ships with 'zipkin')`,
											},
										},
									},
									Description: `List of available tracing backends`,
								},
								"default_backend": schema.StringAttribute{
									Computed:    true,
									Description: `Name of the default backend`,
								},
							},
							MarkdownDescription: `Tracing settings.` + "\n" +
								`+optional`,
						},
						"type": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"key": schema.StringAttribute{
				Optional: true,
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

func (r *MeshListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshListDataSourceModel
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
	var filter *operations.GetMeshListQueryParamFilter
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
	filter = &operations.GetMeshListQueryParamFilter{
		Key:   key,
		Value: value,
	}
	request := operations.GetMeshListRequest{
		CpID:   cpID,
		Offset: offset,
		Size:   size,
		Filter: filter,
	}
	res, err := r.client.Mesh.GetMeshList(ctx, request)
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
	if !(res.MeshList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshList(res.MeshList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
