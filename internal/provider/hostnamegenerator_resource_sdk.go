// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect-beta/internal/provider/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"time"
)

func (r *HostnameGeneratorResourceModel) ToSharedHostnameGeneratorItemInput() *shared.HostnameGeneratorItemInput {
	typeVar := shared.HostnameGeneratorItemType(r.Type.ValueString())
	var name string
	name = r.Name.ValueString()

	labels := make(map[string]string)
	for labelsKey, labelsValue := range r.Labels {
		var labelsInst string
		labelsInst = labelsValue.ValueString()

		labels[labelsKey] = labelsInst
	}
	var selector *shared.Selector
	if r.Spec.Selector != nil {
		var meshExternalService *shared.MeshExternalService
		if r.Spec.Selector.MeshExternalService != nil {
			matchLabels := make(map[string]string)
			for matchLabelsKey, matchLabelsValue := range r.Spec.Selector.MeshExternalService.MatchLabels {
				var matchLabelsInst string
				matchLabelsInst = matchLabelsValue.ValueString()

				matchLabels[matchLabelsKey] = matchLabelsInst
			}
			meshExternalService = &shared.MeshExternalService{
				MatchLabels: matchLabels,
			}
		}
		var meshMultiZoneService *shared.MeshMultiZoneService
		if r.Spec.Selector.MeshMultiZoneService != nil {
			matchLabels1 := make(map[string]string)
			for matchLabelsKey1, matchLabelsValue1 := range r.Spec.Selector.MeshMultiZoneService.MatchLabels {
				var matchLabelsInst1 string
				matchLabelsInst1 = matchLabelsValue1.ValueString()

				matchLabels1[matchLabelsKey1] = matchLabelsInst1
			}
			meshMultiZoneService = &shared.MeshMultiZoneService{
				MatchLabels: matchLabels1,
			}
		}
		var meshService *shared.MeshService
		if r.Spec.Selector.MeshService != nil {
			matchLabels2 := make(map[string]string)
			for matchLabelsKey2, matchLabelsValue2 := range r.Spec.Selector.MeshService.MatchLabels {
				var matchLabelsInst2 string
				matchLabelsInst2 = matchLabelsValue2.ValueString()

				matchLabels2[matchLabelsKey2] = matchLabelsInst2
			}
			meshService = &shared.MeshService{
				MatchLabels: matchLabels2,
			}
		}
		selector = &shared.Selector{
			MeshExternalService:  meshExternalService,
			MeshMultiZoneService: meshMultiZoneService,
			MeshService:          meshService,
		}
	}
	template := new(string)
	if !r.Spec.Template.IsUnknown() && !r.Spec.Template.IsNull() {
		*template = r.Spec.Template.ValueString()
	} else {
		template = nil
	}
	spec := shared.HostnameGeneratorItemSpec{
		Selector: selector,
		Template: template,
	}
	out := shared.HostnameGeneratorItemInput{
		Type:   typeVar,
		Name:   name,
		Labels: labels,
		Spec:   spec,
	}
	return &out
}

func (r *HostnameGeneratorResourceModel) RefreshFromSharedHostnameGeneratorCreateOrUpdateSuccessResponse(resp *shared.HostnameGeneratorCreateOrUpdateSuccessResponse) {
	if resp != nil {
		r.Warnings = make([]types.String, 0, len(resp.Warnings))
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}
}

func (r *HostnameGeneratorResourceModel) RefreshFromSharedHostnameGeneratorItem(resp *shared.HostnameGeneratorItem) {
	if resp != nil {
		if resp.CreationTime != nil {
			r.CreationTime = types.StringValue(resp.CreationTime.Format(time.RFC3339Nano))
		} else {
			r.CreationTime = types.StringNull()
		}
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String, len(resp.Labels))
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringValue(value)
			}
		}
		if resp.ModificationTime != nil {
			r.ModificationTime = types.StringValue(resp.ModificationTime.Format(time.RFC3339Nano))
		} else {
			r.ModificationTime = types.StringNull()
		}
		r.Name = types.StringValue(resp.Name)
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
}
