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
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &APISpecificationDataSource{}
var _ datasource.DataSourceWithConfigure = &APISpecificationDataSource{}

func NewAPISpecificationDataSource() datasource.DataSource {
	return &APISpecificationDataSource{}
}

// APISpecificationDataSource is the data source implementation.
type APISpecificationDataSource struct {
	client *sdk.KonnectBeta
}

// APISpecificationDataSourceModel describes the data model.
type APISpecificationDataSourceModel struct {
	APIID     types.String `tfsdk:"api_id"`
	Content   types.String `tfsdk:"content"`
	CreatedAt types.String `tfsdk:"created_at"`
	ID        types.String `tfsdk:"id"`
	Type      types.String `tfsdk:"type"`
	UpdatedAt types.String `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *APISpecificationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_specification"
}

// Schema defines the schema for the data source.
func (r *APISpecificationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "APISpecification DataSource",

		Attributes: map[string]schema.Attribute{
			"api_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID API identifier`,
			},
			"content": schema.StringAttribute{
				Computed:    true,
				Description: `The raw content of your API specification.`,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity creation date.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The API specification identifier.`,
			},
			"type": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: `The type of specification being stored. This allows us to render the specification correctly.` + "\n" +
					`` + "\n" +
					`If this field is not set, it will be autodetected from ` + "`" + `content` + "`" + ``,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity update date.`,
			},
		},
	}
}

func (r *APISpecificationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *APISpecificationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *APISpecificationDataSourceModel
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

	var apiID string
	apiID = data.APIID.ValueString()

	var specID string
	specID = data.ID.ValueString()

	request := operations.FetchAPISpecRequest{
		APIID:  apiID,
		SpecID: specID,
	}
	res, err := r.client.APISpecification.FetchAPISpec(ctx, request)
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
	if !(res.APISpecResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAPISpecResponse(res.APISpecResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
