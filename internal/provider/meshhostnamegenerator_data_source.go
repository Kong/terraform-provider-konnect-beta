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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MeshHostnameGeneratorDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshHostnameGeneratorDataSource{}

func NewMeshHostnameGeneratorDataSource() datasource.DataSource {
	return &MeshHostnameGeneratorDataSource{}
}

// MeshHostnameGeneratorDataSource is the data source implementation.
type MeshHostnameGeneratorDataSource struct {
	client *sdk.KonnectBeta
}

// MeshHostnameGeneratorDataSourceModel describes the data model.
type MeshHostnameGeneratorDataSourceModel struct {
	CpID             types.String                      `tfsdk:"cp_id"`
	CreationTime     types.String                      `tfsdk:"creation_time"`
	Labels           map[string]types.String           `tfsdk:"labels"`
	ModificationTime types.String                      `tfsdk:"modification_time"`
	Name             types.String                      `tfsdk:"name"`
	Spec             tfTypes.HostnameGeneratorItemSpec `tfsdk:"spec"`
	Type             types.String                      `tfsdk:"type"`
}

// Metadata returns the data source type name.
func (r *MeshHostnameGeneratorDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_hostname_generator"
}

// Schema defines the schema for the data source.
func (r *MeshHostnameGeneratorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshHostnameGenerator DataSource",

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
			"modification_time": schema.StringAttribute{
				Computed:    true,
				Description: `Time at which the resource was updated`,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `name of the HostnameGenerator`,
			},
			"spec": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"extension": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"config": schema.StringAttribute{
								Computed:    true,
								Description: `Config freeform configuration for the extension. Parsed as JSON.`,
							},
							"type": schema.StringAttribute{
								Computed:    true,
								Description: `Type of the extension.`,
							},
						},
						Description: `Extension struct for a plugin configuration`,
					},
					"selector": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"mesh_external_service": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"match_labels": schema.MapAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
								},
							},
							"mesh_multi_zone_service": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"match_labels": schema.MapAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
								},
							},
							"mesh_service": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"match_labels": schema.MapAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
								},
							},
						},
					},
					"template": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `Spec is the specification of the Kuma HostnameGenerator resource.`,
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: `the type of the resource`,
			},
		},
	}
}

func (r *MeshHostnameGeneratorDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshHostnameGeneratorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshHostnameGeneratorDataSourceModel
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

	request, requestDiags := data.ToOperationsGetHostnameGeneratorRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.HostnameGenerator.GetHostnameGenerator(ctx, *request)
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
	if !(res.HostnameGeneratorItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedHostnameGeneratorItem(ctx, res.HostnameGeneratorItem)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
