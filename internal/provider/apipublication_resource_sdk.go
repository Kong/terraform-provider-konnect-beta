// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/provider/typeconvert"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
)

func (r *APIPublicationResourceModel) ToSharedAPIPublication(ctx context.Context) (*shared.APIPublication, diag.Diagnostics) {
	var diags diag.Diagnostics

	autoApproveRegistrations := new(bool)
	if !r.AutoApproveRegistrations.IsUnknown() && !r.AutoApproveRegistrations.IsNull() {
		*autoApproveRegistrations = r.AutoApproveRegistrations.ValueBool()
	} else {
		autoApproveRegistrations = nil
	}
	var authStrategyIds []string
	if r.AuthStrategyIds != nil {
		authStrategyIds = make([]string, 0, len(r.AuthStrategyIds))
		for _, authStrategyIdsItem := range r.AuthStrategyIds {
			authStrategyIds = append(authStrategyIds, authStrategyIdsItem.ValueString())
		}
	}
	visibility := new(shared.APIPublicationVisibility)
	if !r.Visibility.IsUnknown() && !r.Visibility.IsNull() {
		*visibility = shared.APIPublicationVisibility(r.Visibility.ValueString())
	} else {
		visibility = nil
	}
	out := shared.APIPublication{
		AutoApproveRegistrations: autoApproveRegistrations,
		AuthStrategyIds:          authStrategyIds,
		Visibility:               visibility,
	}

	return &out, diags
}

func (r *APIPublicationResourceModel) ToOperationsPublishAPIToPortalRequest(ctx context.Context) (*operations.PublishAPIToPortalRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var apiID string
	apiID = r.APIID.ValueString()

	var portalID string
	portalID = r.PortalID.ValueString()

	apiPublication, apiPublicationDiags := r.ToSharedAPIPublication(ctx)
	diags.Append(apiPublicationDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.PublishAPIToPortalRequest{
		APIID:          apiID,
		PortalID:       portalID,
		APIPublication: *apiPublication,
	}

	return &out, diags
}

func (r *APIPublicationResourceModel) ToOperationsFetchPublicationRequest(ctx context.Context) (*operations.FetchPublicationRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var apiID string
	apiID = r.APIID.ValueString()

	var portalID string
	portalID = r.PortalID.ValueString()

	out := operations.FetchPublicationRequest{
		APIID:    apiID,
		PortalID: portalID,
	}

	return &out, diags
}

func (r *APIPublicationResourceModel) ToOperationsDeletePublicationRequest(ctx context.Context) (*operations.DeletePublicationRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var apiID string
	apiID = r.APIID.ValueString()

	var portalID string
	portalID = r.PortalID.ValueString()

	out := operations.DeletePublicationRequest{
		APIID:    apiID,
		PortalID: portalID,
	}

	return &out, diags
}

func (r *APIPublicationResourceModel) RefreshFromSharedAPIPublicationResponse(ctx context.Context, resp *shared.APIPublicationResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.AuthStrategyIds != nil {
			r.AuthStrategyIds = make([]types.String, 0, len(resp.AuthStrategyIds))
			for _, v := range resp.AuthStrategyIds {
				r.AuthStrategyIds = append(r.AuthStrategyIds, types.StringValue(v))
			}
		}
		r.AutoApproveRegistrations = types.BoolPointerValue(resp.AutoApproveRegistrations)
		r.CreatedAt = types.StringValue(typeconvert.TimeToString(resp.CreatedAt))
		r.UpdatedAt = types.StringValue(typeconvert.TimeToString(resp.UpdatedAt))
		if resp.Visibility != nil {
			r.Visibility = types.StringValue(string(*resp.Visibility))
		} else {
			r.Visibility = types.StringNull()
		}
	}

	return diags
}
