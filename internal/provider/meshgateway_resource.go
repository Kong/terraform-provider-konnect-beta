// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	custom_boolplanmodifier "github.com/kong/terraform-provider-konnect-beta/internal/planmodifiers/boolplanmodifier"
	custom_listplanmodifier "github.com/kong/terraform-provider-konnect-beta/internal/planmodifiers/listplanmodifier"
	custom_stringplanmodifier "github.com/kong/terraform-provider-konnect-beta/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/kong/terraform-provider-konnect-beta/internal/provider/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect-beta/internal/validators"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-konnect-beta/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/kong/terraform-provider-konnect-beta/internal/validators/stringvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MeshGatewayResource{}
var _ resource.ResourceWithImportState = &MeshGatewayResource{}

func NewMeshGatewayResource() resource.Resource {
	return &MeshGatewayResource{}
}

// MeshGatewayResource defines the resource implementation.
type MeshGatewayResource struct {
	client *sdk.KonnectBeta
}

// MeshGatewayResourceModel describes the resource data model.
type MeshGatewayResourceModel struct {
	Conf      *tfTypes.Conf           `tfsdk:"conf"`
	CpID      types.String            `tfsdk:"cp_id"`
	Labels    map[string]types.String `tfsdk:"labels"`
	Mesh      types.String            `tfsdk:"mesh"`
	Name      types.String            `tfsdk:"name"`
	Selectors []tfTypes.Selectors     `tfsdk:"selectors"`
	Tags      map[string]types.String `tfsdk:"tags"`
	Type      types.String            `tfsdk:"type"`
	Warnings  []types.String          `tfsdk:"warnings"`
}

func (r *MeshGatewayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "konnect_mesh_gateway"
}

func (r *MeshGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshGateway Resource",
		Attributes: map[string]schema.Attribute{
			"conf": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"listeners": schema.ListNestedAttribute{
						Optional: true,
						PlanModifiers: []planmodifier.List{
							custom_listplanmodifier.SupressZeroNullModifier(),
						},
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							Attributes: map[string]schema.Attribute{
								"cross_mesh": schema.BoolAttribute{
									Optional: true,
									PlanModifiers: []planmodifier.Bool{
										custom_boolplanmodifier.SupressZeroNullModifier(),
									},
									MarkdownDescription: `CrossMesh enables traffic to flow to this listener only from other` + "\n" +
										`meshes.`,
								},
								"hostname": schema.StringAttribute{
									Optional: true,
									MarkdownDescription: `Hostname specifies the virtual hostname to match for protocol types that` + "\n" +
										`define this concept. When unspecified, "", or ` + "`" + `*` + "`" + `, all hostnames are` + "\n" +
										`matched. This field can be omitted for protocols that don't require` + "\n" +
										`hostname based matching.`,
								},
								"port": schema.Int64Attribute{
									Optional: true,
									MarkdownDescription: `Port is the network port. Multiple listeners may use the` + "\n" +
										`same port, subject to the Listener compatibility rules.`,
								},
								"protocol": schema.SingleNestedAttribute{
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
									Description: `Protocol specifies the network protocol this listener expects to receive.`,
								},
								"resources": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"connection_limit": schema.Int64Attribute{
											Optional: true,
										},
									},
									Description: `Resources is used to specify listener-specific resource settings.`,
								},
								"tags": schema.MapAttribute{
									Optional:    true,
									ElementType: types.StringType,
									MarkdownDescription: `Tags specifies a unique combination of tags that routes can use` + "\n" +
										`to match themselves to this listener.` + "\n" +
										`` + "\n" +
										`When matching routes to listeners, the control plane constructs a` + "\n" +
										`set of matching tags for each listener by forming the union of the` + "\n" +
										`gateway tags and the listener tags. A route will be attached to the` + "\n" +
										`listener if all of the route's tags are preset in the matching tags`,
								},
								"tls": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"certificates": schema.ListNestedAttribute{
											Optional: true,
											PlanModifiers: []planmodifier.List{
												custom_listplanmodifier.SupressZeroNullModifier(),
											},
											NestedObject: schema.NestedAttributeObject{
												Validators: []validator.Object{
													speakeasy_objectvalidators.NotNull(),
												},
												Attributes: map[string]schema.Attribute{
													"type": schema.StringAttribute{
														Optional: true,
														MarkdownDescription: `Types that are assignable to Type:` + "\n" +
															`` + "\n" +
															`	*DataSource_Secret` + "\n" +
															`	*DataSource_File` + "\n" +
															`	*DataSource_Inline` + "\n" +
															`	*DataSource_InlineString` + "\n" +
															`Not Null; Parsed as JSON.`,
														Validators: []validator.String{
															speakeasy_stringvalidators.NotNull(),
															validators.IsValidJSON(),
														},
													},
												},
											},
											MarkdownDescription: `Certificates is an array of datasources that contain TLS` + "\n" +
												`certificates and private keys.  Each datasource must contain a` + "\n" +
												`sequence of PEM-encoded objects. The server certificate and private` + "\n" +
												`key are required, but additional certificates are allowed and will` + "\n" +
												`be added to the certificate chain.  The server certificate must` + "\n" +
												`be the first certificate in the datasource.` + "\n" +
												`` + "\n" +
												`When multiple certificate datasources are configured, they must have` + "\n" +
												`different key types. In practice, this means that one datasource` + "\n" +
												`should contain an RSA key and certificate, and the other an` + "\n" +
												`ECDSA key and certificate.`,
										},
										"mode": schema.SingleNestedAttribute{
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
											MarkdownDescription: `Mode defines the TLS behavior for the TLS session initiated` + "\n" +
												`by the client.`,
										},
										"options": schema.SingleNestedAttribute{
											Optional: true,
											MarkdownDescription: `Options should eventually configure how TLS is configured. This` + "\n" +
												`is where cipher suite and version configuration can be specified,` + "\n" +
												`client certificates enforced, and so on.`,
										},
									},
									MarkdownDescription: `TLS is the TLS configuration for the Listener. This field` + "\n" +
										`is required if the Protocol field is "HTTPS" or "TLS" and` + "\n" +
										`ignored otherwise.`,
								},
							},
						},
						MarkdownDescription: `Listeners define logical endpoints that are bound on this MeshGateway's` + "\n" +
							`address(es).`,
					},
				},
				Description: `The desired configuration of the MeshGateway.`,
			},
			"cp_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					custom_stringplanmodifier.RequiresReplaceModifier(),
				},
				Description: `Id of the Konnect resource`,
			},
			"labels": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"mesh": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					custom_stringplanmodifier.RequiresReplaceModifier(),
				},
				Description: `name of the mesh`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					custom_stringplanmodifier.RequiresReplaceModifier(),
				},
				Description: `name of the MeshGateway`,
			},
			"selectors": schema.ListNestedAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.List{
					custom_listplanmodifier.SupressZeroNullModifier(),
				},
				NestedObject: schema.NestedAttributeObject{
					Validators: []validator.Object{
						speakeasy_objectvalidators.NotNull(),
					},
					Attributes: map[string]schema.Attribute{
						"match": schema.MapAttribute{
							Optional:    true,
							ElementType: types.StringType,
							Description: `Tags to match, can be used for both source and destinations`,
						},
					},
				},
				MarkdownDescription: `Selectors is a list of selectors that are used to match builtin` + "\n" +
					`gateway dataplanes that will receive this MeshGateway configuration.`,
			},
			"tags": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				MarkdownDescription: `Tags is the set of tags common to all of the gateway's listeners.` + "\n" +
					`` + "\n" +
					`This field must not include a ` + "`" + `kuma.io/service` + "`" + ` tag (the service is always` + "\n" +
					`defined on the dataplanes).`,
			},
			"type": schema.StringAttribute{
				Required: true,
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

func (r *MeshGatewayResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MeshGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MeshGatewayResourceModel
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

	meshGatewayItem := *data.ToSharedMeshGatewayItem()
	request := operations.CreateMeshGatewayRequest{
		CpID:            cpID,
		Mesh:            mesh,
		Name:            name,
		MeshGatewayItem: meshGatewayItem,
	}
	res, err := r.client.MeshGateway.CreateMeshGateway(ctx, request)
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
	if !(res.MeshGatewayCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshGatewayCreateOrUpdateSuccessResponse(res.MeshGatewayCreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var cpId1 string
	cpId1 = data.CpID.ValueString()

	var mesh1 string
	mesh1 = data.Mesh.ValueString()

	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetMeshGatewayRequest{
		CpID: cpId1,
		Mesh: mesh1,
		Name: name1,
	}
	res1, err := r.client.MeshGateway.GetMeshGateway(ctx, request1)
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
	if !(res1.MeshGatewayItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedMeshGatewayItem(res1.MeshGatewayItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MeshGatewayResourceModel
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

	request := operations.GetMeshGatewayRequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshGateway.GetMeshGateway(ctx, request)
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
	if !(res.MeshGatewayItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshGatewayItem(res.MeshGatewayItem)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *MeshGatewayResourceModel
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

	meshGatewayItem := *data.ToSharedMeshGatewayItem()
	request := operations.UpdateMeshGatewayRequest{
		CpID:            cpID,
		Mesh:            mesh,
		Name:            name,
		MeshGatewayItem: meshGatewayItem,
	}
	res, err := r.client.MeshGateway.UpdateMeshGateway(ctx, request)
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
	if !(res.MeshGatewayCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshGatewayCreateOrUpdateSuccessResponse(res.MeshGatewayCreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var cpId1 string
	cpId1 = data.CpID.ValueString()

	var mesh1 string
	mesh1 = data.Mesh.ValueString()

	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetMeshGatewayRequest{
		CpID: cpId1,
		Mesh: mesh1,
		Name: name1,
	}
	res1, err := r.client.MeshGateway.GetMeshGateway(ctx, request1)
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
	if !(res1.MeshGatewayItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedMeshGatewayItem(res1.MeshGatewayItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MeshGatewayResourceModel
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

	request := operations.DeleteMeshGatewayRequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshGateway.DeleteMeshGateway(ctx, request)
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

func (r *MeshGatewayResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
