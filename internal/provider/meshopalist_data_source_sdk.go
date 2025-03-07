// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect-beta/internal/provider/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *MeshOPAListDataSourceModel) RefreshFromSharedMeshOPAList(resp *shared.MeshOPAList) {
	if resp != nil {
		r.Items = []tfTypes.MeshOPAItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items1 tfTypes.MeshOPAItem
			if itemsItem.CreationTime != nil {
				items1.CreationTime = types.StringValue(itemsItem.CreationTime.Format(time.RFC3339Nano))
			} else {
				items1.CreationTime = types.StringNull()
			}
			if len(itemsItem.Labels) > 0 {
				items1.Labels = make(map[string]types.String)
				for key, value := range itemsItem.Labels {
					items1.Labels[key] = types.StringValue(value)
				}
			}
			items1.Mesh = types.StringPointerValue(itemsItem.Mesh)
			if itemsItem.ModificationTime != nil {
				items1.ModificationTime = types.StringValue(itemsItem.ModificationTime.Format(time.RFC3339Nano))
			} else {
				items1.ModificationTime = types.StringNull()
			}
			items1.Name = types.StringValue(itemsItem.Name)
			if itemsItem.Spec.Default == nil {
				items1.Spec.Default = nil
			} else {
				items1.Spec.Default = &tfTypes.MeshOPAItemDefault{}
				if itemsItem.Spec.Default.AgentConfig == nil {
					items1.Spec.Default.AgentConfig = nil
				} else {
					items1.Spec.Default.AgentConfig = &tfTypes.CaCert{}
					items1.Spec.Default.AgentConfig.Inline = types.StringPointerValue(itemsItem.Spec.Default.AgentConfig.Inline)
					items1.Spec.Default.AgentConfig.InlineString = types.StringPointerValue(itemsItem.Spec.Default.AgentConfig.InlineString)
					items1.Spec.Default.AgentConfig.Secret = types.StringPointerValue(itemsItem.Spec.Default.AgentConfig.Secret)
				}
				items1.Spec.Default.AppendPolicies = []tfTypes.AppendPolicies{}
				for appendPoliciesCount, appendPoliciesItem := range itemsItem.Spec.Default.AppendPolicies {
					var appendPolicies1 tfTypes.AppendPolicies
					appendPolicies1.IgnoreDecision = types.BoolPointerValue(appendPoliciesItem.IgnoreDecision)
					appendPolicies1.Rego.Inline = types.StringPointerValue(appendPoliciesItem.Rego.Inline)
					appendPolicies1.Rego.InlineString = types.StringPointerValue(appendPoliciesItem.Rego.InlineString)
					appendPolicies1.Rego.Secret = types.StringPointerValue(appendPoliciesItem.Rego.Secret)
					if appendPoliciesCount+1 > len(items1.Spec.Default.AppendPolicies) {
						items1.Spec.Default.AppendPolicies = append(items1.Spec.Default.AppendPolicies, appendPolicies1)
					} else {
						items1.Spec.Default.AppendPolicies[appendPoliciesCount].IgnoreDecision = appendPolicies1.IgnoreDecision
						items1.Spec.Default.AppendPolicies[appendPoliciesCount].Rego = appendPolicies1.Rego
					}
				}
				if itemsItem.Spec.Default.AuthConfig == nil {
					items1.Spec.Default.AuthConfig = nil
				} else {
					items1.Spec.Default.AuthConfig = &tfTypes.AuthConfig{}
					if itemsItem.Spec.Default.AuthConfig.OnAgentFailure != nil {
						items1.Spec.Default.AuthConfig.OnAgentFailure = types.StringValue(string(*itemsItem.Spec.Default.AuthConfig.OnAgentFailure))
					} else {
						items1.Spec.Default.AuthConfig.OnAgentFailure = types.StringNull()
					}
					if itemsItem.Spec.Default.AuthConfig.RequestBody == nil {
						items1.Spec.Default.AuthConfig.RequestBody = nil
					} else {
						items1.Spec.Default.AuthConfig.RequestBody = &tfTypes.RequestBody{}
						if itemsItem.Spec.Default.AuthConfig.RequestBody.MaxSize != nil {
							items1.Spec.Default.AuthConfig.RequestBody.MaxSize = types.Int64Value(int64(*itemsItem.Spec.Default.AuthConfig.RequestBody.MaxSize))
						} else {
							items1.Spec.Default.AuthConfig.RequestBody.MaxSize = types.Int64Null()
						}
						items1.Spec.Default.AuthConfig.RequestBody.SendRawBody = types.BoolPointerValue(itemsItem.Spec.Default.AuthConfig.RequestBody.SendRawBody)
					}
					if itemsItem.Spec.Default.AuthConfig.StatusOnError != nil {
						items1.Spec.Default.AuthConfig.StatusOnError = types.Int64Value(int64(*itemsItem.Spec.Default.AuthConfig.StatusOnError))
					} else {
						items1.Spec.Default.AuthConfig.StatusOnError = types.Int64Null()
					}
					items1.Spec.Default.AuthConfig.Timeout = types.StringPointerValue(itemsItem.Spec.Default.AuthConfig.Timeout)
				}
			}
			if itemsItem.Spec.TargetRef == nil {
				items1.Spec.TargetRef = nil
			} else {
				items1.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
				items1.Spec.TargetRef.Kind = types.StringValue(string(itemsItem.Spec.TargetRef.Kind))
				if len(itemsItem.Spec.TargetRef.Labels) > 0 {
					items1.Spec.TargetRef.Labels = make(map[string]types.String)
					for key1, value1 := range itemsItem.Spec.TargetRef.Labels {
						items1.Spec.TargetRef.Labels[key1] = types.StringValue(value1)
					}
				}
				items1.Spec.TargetRef.Mesh = types.StringPointerValue(itemsItem.Spec.TargetRef.Mesh)
				items1.Spec.TargetRef.Name = types.StringPointerValue(itemsItem.Spec.TargetRef.Name)
				items1.Spec.TargetRef.Namespace = types.StringPointerValue(itemsItem.Spec.TargetRef.Namespace)
				items1.Spec.TargetRef.ProxyTypes = make([]types.String, 0, len(itemsItem.Spec.TargetRef.ProxyTypes))
				for _, v := range itemsItem.Spec.TargetRef.ProxyTypes {
					items1.Spec.TargetRef.ProxyTypes = append(items1.Spec.TargetRef.ProxyTypes, types.StringValue(string(v)))
				}
				items1.Spec.TargetRef.SectionName = types.StringPointerValue(itemsItem.Spec.TargetRef.SectionName)
				if len(itemsItem.Spec.TargetRef.Tags) > 0 {
					items1.Spec.TargetRef.Tags = make(map[string]types.String)
					for key2, value2 := range itemsItem.Spec.TargetRef.Tags {
						items1.Spec.TargetRef.Tags[key2] = types.StringValue(value2)
					}
				}
			}
			items1.Type = types.StringValue(string(itemsItem.Type))
			if itemsCount+1 > len(r.Items) {
				r.Items = append(r.Items, items1)
			} else {
				r.Items[itemsCount].CreationTime = items1.CreationTime
				r.Items[itemsCount].Labels = items1.Labels
				r.Items[itemsCount].Mesh = items1.Mesh
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
