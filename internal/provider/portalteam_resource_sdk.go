// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"time"
)

func (r *PortalTeamResourceModel) ToSharedPortalCreateTeamRequest() *shared.PortalCreateTeamRequest {
	var name string
	name = r.Name.ValueString()

	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	out := shared.PortalCreateTeamRequest{
		Name:        name,
		Description: description,
	}
	return &out
}

func (r *PortalTeamResourceModel) RefreshFromSharedPortalTeamResponse(resp *shared.PortalTeamResponse) {
	if resp != nil {
		if resp.CreatedAt != nil {
			r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		} else {
			r.CreatedAt = types.StringNull()
		}
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringPointerValue(resp.Name)
		if resp.UpdatedAt != nil {
			r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		} else {
			r.UpdatedAt = types.StringNull()
		}
	}
}

func (r *PortalTeamResourceModel) ToSharedPortalUpdateTeamRequest() *shared.PortalUpdateTeamRequest {
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	out := shared.PortalUpdateTeamRequest{
		Name:        name,
		Description: description,
	}
	return &out
}
