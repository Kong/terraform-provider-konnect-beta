// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/provider/typeconvert"
	tfTypes "github.com/kong/terraform-provider-konnect-beta/internal/provider/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
)

func (r *MeshHostnameGeneratorDataSourceModel) ToOperationsGetHostnameGeneratorRequest(ctx context.Context) (*operations.GetHostnameGeneratorRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var cpID string
	cpID = r.CpID.ValueString()

	var name string
	name = r.Name.ValueString()

	out := operations.GetHostnameGeneratorRequest{
		CpID: cpID,
		Name: name,
	}

	return &out, diags
}

func (r *MeshHostnameGeneratorDataSourceModel) RefreshFromSharedHostnameGeneratorItem(ctx context.Context, resp *shared.HostnameGeneratorItem) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CreationTime = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.CreationTime))
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String, len(resp.Labels))
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringValue(value)
			}
		}
		r.ModificationTime = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.ModificationTime))
		r.Name = types.StringValue(resp.Name)
		if resp.Spec.Extension == nil {
			r.Spec.Extension = nil
		} else {
			r.Spec.Extension = &tfTypes.MeshExternalServiceItemExtension{}
			if resp.Spec.Extension.Config == nil {
				r.Spec.Extension.Config = types.StringNull()
			} else {
				configResult, _ := json.Marshal(resp.Spec.Extension.Config)
				r.Spec.Extension.Config = types.StringValue(string(configResult))
			}
			r.Spec.Extension.Type = types.StringValue(resp.Spec.Extension.Type)
		}
		if resp.Spec.Selector == nil {
			r.Spec.Selector = nil
		} else {
			r.Spec.Selector = &tfTypes.Selector{}
			if resp.Spec.Selector.MeshExternalService == nil {
				r.Spec.Selector.MeshExternalService = nil
			} else {
				r.Spec.Selector.MeshExternalService = &tfTypes.MeshExternalService{}
				if len(resp.Spec.Selector.MeshExternalService.MatchLabels) > 0 {
					r.Spec.Selector.MeshExternalService.MatchLabels = make(map[string]types.String, len(resp.Spec.Selector.MeshExternalService.MatchLabels))
					for key1, value1 := range resp.Spec.Selector.MeshExternalService.MatchLabels {
						r.Spec.Selector.MeshExternalService.MatchLabels[key1] = types.StringValue(value1)
					}
				}
			}
			if resp.Spec.Selector.MeshMultiZoneService == nil {
				r.Spec.Selector.MeshMultiZoneService = nil
			} else {
				r.Spec.Selector.MeshMultiZoneService = &tfTypes.MeshExternalService{}
				if len(resp.Spec.Selector.MeshMultiZoneService.MatchLabels) > 0 {
					r.Spec.Selector.MeshMultiZoneService.MatchLabels = make(map[string]types.String, len(resp.Spec.Selector.MeshMultiZoneService.MatchLabels))
					for key2, value2 := range resp.Spec.Selector.MeshMultiZoneService.MatchLabels {
						r.Spec.Selector.MeshMultiZoneService.MatchLabels[key2] = types.StringValue(value2)
					}
				}
			}
			if resp.Spec.Selector.MeshService == nil {
				r.Spec.Selector.MeshService = nil
			} else {
				r.Spec.Selector.MeshService = &tfTypes.MeshExternalService{}
				if len(resp.Spec.Selector.MeshService.MatchLabels) > 0 {
					r.Spec.Selector.MeshService.MatchLabels = make(map[string]types.String, len(resp.Spec.Selector.MeshService.MatchLabels))
					for key3, value3 := range resp.Spec.Selector.MeshService.MatchLabels {
						r.Spec.Selector.MeshService.MatchLabels[key3] = types.StringValue(value3)
					}
				}
			}
		}
		r.Spec.Template = types.StringPointerValue(resp.Spec.Template)
		r.Type = types.StringValue(string(resp.Type))
	}

	return diags
}
