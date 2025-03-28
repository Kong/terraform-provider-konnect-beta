// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect-beta/internal/provider/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *HostnameGeneratorListDataSourceModel) RefreshFromSharedHostnameGeneratorList(resp *shared.HostnameGeneratorList) {
	if resp != nil {
		r.Items = []tfTypes.HostnameGeneratorItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items1 tfTypes.HostnameGeneratorItem
			if itemsItem.CreationTime != nil {
				items1.CreationTime = types.StringValue(itemsItem.CreationTime.Format(time.RFC3339Nano))
			} else {
				items1.CreationTime = types.StringNull()
			}
			if len(itemsItem.Labels) > 0 {
				items1.Labels = make(map[string]types.String, len(itemsItem.Labels))
				for key, value := range itemsItem.Labels {
					items1.Labels[key] = types.StringValue(value)
				}
			}
			if itemsItem.ModificationTime != nil {
				items1.ModificationTime = types.StringValue(itemsItem.ModificationTime.Format(time.RFC3339Nano))
			} else {
				items1.ModificationTime = types.StringNull()
			}
			items1.Name = types.StringValue(itemsItem.Name)
			if itemsItem.Spec.Selector == nil {
				items1.Spec.Selector = nil
			} else {
				items1.Spec.Selector = &tfTypes.Selector{}
				if itemsItem.Spec.Selector.MeshExternalService == nil {
					items1.Spec.Selector.MeshExternalService = nil
				} else {
					items1.Spec.Selector.MeshExternalService = &tfTypes.MeshExternalService{}
					if len(itemsItem.Spec.Selector.MeshExternalService.MatchLabels) > 0 {
						items1.Spec.Selector.MeshExternalService.MatchLabels = make(map[string]types.String, len(itemsItem.Spec.Selector.MeshExternalService.MatchLabels))
						for key1, value1 := range itemsItem.Spec.Selector.MeshExternalService.MatchLabels {
							items1.Spec.Selector.MeshExternalService.MatchLabels[key1] = types.StringValue(value1)
						}
					}
				}
				if itemsItem.Spec.Selector.MeshMultiZoneService == nil {
					items1.Spec.Selector.MeshMultiZoneService = nil
				} else {
					items1.Spec.Selector.MeshMultiZoneService = &tfTypes.MeshExternalService{}
					if len(itemsItem.Spec.Selector.MeshMultiZoneService.MatchLabels) > 0 {
						items1.Spec.Selector.MeshMultiZoneService.MatchLabels = make(map[string]types.String, len(itemsItem.Spec.Selector.MeshMultiZoneService.MatchLabels))
						for key2, value2 := range itemsItem.Spec.Selector.MeshMultiZoneService.MatchLabels {
							items1.Spec.Selector.MeshMultiZoneService.MatchLabels[key2] = types.StringValue(value2)
						}
					}
				}
				if itemsItem.Spec.Selector.MeshService == nil {
					items1.Spec.Selector.MeshService = nil
				} else {
					items1.Spec.Selector.MeshService = &tfTypes.MeshExternalService{}
					if len(itemsItem.Spec.Selector.MeshService.MatchLabels) > 0 {
						items1.Spec.Selector.MeshService.MatchLabels = make(map[string]types.String, len(itemsItem.Spec.Selector.MeshService.MatchLabels))
						for key3, value3 := range itemsItem.Spec.Selector.MeshService.MatchLabels {
							items1.Spec.Selector.MeshService.MatchLabels[key3] = types.StringValue(value3)
						}
					}
				}
			}
			items1.Spec.Template = types.StringPointerValue(itemsItem.Spec.Template)
			items1.Type = types.StringValue(string(itemsItem.Type))
			if itemsCount+1 > len(r.Items) {
				r.Items = append(r.Items, items1)
			} else {
				r.Items[itemsCount].CreationTime = items1.CreationTime
				r.Items[itemsCount].Labels = items1.Labels
				r.Items[itemsCount].ModificationTime = items1.ModificationTime
				r.Items[itemsCount].Name = items1.Name
				r.Items[itemsCount].Spec = items1.Spec
				r.Items[itemsCount].Type = items1.Type
			}
		}
		r.Next = types.StringPointerValue(resp.Next)
		if resp.Total != nil {
			r.Total = types.NumberValue(big.NewFloat(float64(*resp.Total)))
		} else {
			r.Total = types.NumberNull()
		}
	}
}
