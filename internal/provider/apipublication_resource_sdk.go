// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"time"
)

func (r *APIPublicationResourceModel) ToSharedAPIPublication() *shared.APIPublication {
	autoApproveRegistrations := new(bool)
	if !r.AutoApproveRegistrations.IsUnknown() && !r.AutoApproveRegistrations.IsNull() {
		*autoApproveRegistrations = r.AutoApproveRegistrations.ValueBool()
	} else {
		autoApproveRegistrations = nil
	}
	var authStrategyIds []string = []string{}
	for _, authStrategyIdsItem := range r.AuthStrategyIds {
		authStrategyIds = append(authStrategyIds, authStrategyIdsItem.ValueString())
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
	return &out
}

func (r *APIPublicationResourceModel) RefreshFromSharedAPIPublicationResponse(resp *shared.APIPublicationResponse) {
	if resp != nil {
		if resp.AuthStrategyIds != nil {
			r.AuthStrategyIds = make([]types.String, 0, len(resp.AuthStrategyIds))
			for _, v := range resp.AuthStrategyIds {
				r.AuthStrategyIds = append(r.AuthStrategyIds, types.StringValue(v))
			}
		}
		r.AutoApproveRegistrations = types.BoolPointerValue(resp.AutoApproveRegistrations)
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		if resp.Visibility != nil {
			r.Visibility = types.StringValue(string(*resp.Visibility))
		} else {
			r.Visibility = types.StringNull()
		}
	}
}
