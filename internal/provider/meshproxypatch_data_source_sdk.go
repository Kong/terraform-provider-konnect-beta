// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect-beta/internal/provider/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"time"
)

func (r *MeshProxyPatchDataSourceModel) RefreshFromSharedMeshProxyPatchItem(resp *shared.MeshProxyPatchItem) {
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
		r.Mesh = types.StringPointerValue(resp.Mesh)
		if resp.ModificationTime != nil {
			r.ModificationTime = types.StringValue(resp.ModificationTime.Format(time.RFC3339Nano))
		} else {
			r.ModificationTime = types.StringNull()
		}
		r.Name = types.StringValue(resp.Name)
		r.Spec.Default.AppendModifications = []tfTypes.AppendModifications{}
		if len(r.Spec.Default.AppendModifications) > len(resp.Spec.Default.AppendModifications) {
			r.Spec.Default.AppendModifications = r.Spec.Default.AppendModifications[:len(resp.Spec.Default.AppendModifications)]
		}
		for appendModificationsCount, appendModificationsItem := range resp.Spec.Default.AppendModifications {
			var appendModifications1 tfTypes.AppendModifications
			if appendModificationsItem.Cluster == nil {
				appendModifications1.Cluster = nil
			} else {
				appendModifications1.Cluster = &tfTypes.Cluster{}
				appendModifications1.Cluster.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount, jsonPatchesItem := range appendModificationsItem.Cluster.JSONPatches {
					var jsonPatches1 tfTypes.JSONPatches
					jsonPatches1.From = types.StringPointerValue(jsonPatchesItem.From)
					jsonPatches1.Op = types.StringValue(string(jsonPatchesItem.Op))
					jsonPatches1.Path = types.StringValue(jsonPatchesItem.Path)
					if jsonPatchesItem.Value == nil {
						jsonPatches1.Value = types.StringNull()
					} else {
						valueResult, _ := json.Marshal(jsonPatchesItem.Value)
						jsonPatches1.Value = types.StringValue(string(valueResult))
					}
					if jsonPatchesCount+1 > len(appendModifications1.Cluster.JSONPatches) {
						appendModifications1.Cluster.JSONPatches = append(appendModifications1.Cluster.JSONPatches, jsonPatches1)
					} else {
						appendModifications1.Cluster.JSONPatches[jsonPatchesCount].From = jsonPatches1.From
						appendModifications1.Cluster.JSONPatches[jsonPatchesCount].Op = jsonPatches1.Op
						appendModifications1.Cluster.JSONPatches[jsonPatchesCount].Path = jsonPatches1.Path
						appendModifications1.Cluster.JSONPatches[jsonPatchesCount].Value = jsonPatches1.Value
					}
				}
				if appendModificationsItem.Cluster.Match == nil {
					appendModifications1.Cluster.Match = nil
				} else {
					appendModifications1.Cluster.Match = &tfTypes.MeshProxyPatchItemSpecDefaultAppendModificationsClusterMatch{}
					appendModifications1.Cluster.Match.Name = types.StringPointerValue(appendModificationsItem.Cluster.Match.Name)
					appendModifications1.Cluster.Match.Origin = types.StringPointerValue(appendModificationsItem.Cluster.Match.Origin)
				}
				appendModifications1.Cluster.Operation = types.StringValue(string(appendModificationsItem.Cluster.Operation))
				appendModifications1.Cluster.Value = types.StringPointerValue(appendModificationsItem.Cluster.Value)
			}
			if appendModificationsItem.HTTPFilter == nil {
				appendModifications1.HTTPFilter = nil
			} else {
				appendModifications1.HTTPFilter = &tfTypes.HTTPFilter{}
				appendModifications1.HTTPFilter.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount1, jsonPatchesItem1 := range appendModificationsItem.HTTPFilter.JSONPatches {
					var jsonPatches3 tfTypes.JSONPatches
					jsonPatches3.From = types.StringPointerValue(jsonPatchesItem1.From)
					jsonPatches3.Op = types.StringValue(string(jsonPatchesItem1.Op))
					jsonPatches3.Path = types.StringValue(jsonPatchesItem1.Path)
					if jsonPatchesItem1.Value == nil {
						jsonPatches3.Value = types.StringNull()
					} else {
						valueResult1, _ := json.Marshal(jsonPatchesItem1.Value)
						jsonPatches3.Value = types.StringValue(string(valueResult1))
					}
					if jsonPatchesCount1+1 > len(appendModifications1.HTTPFilter.JSONPatches) {
						appendModifications1.HTTPFilter.JSONPatches = append(appendModifications1.HTTPFilter.JSONPatches, jsonPatches3)
					} else {
						appendModifications1.HTTPFilter.JSONPatches[jsonPatchesCount1].From = jsonPatches3.From
						appendModifications1.HTTPFilter.JSONPatches[jsonPatchesCount1].Op = jsonPatches3.Op
						appendModifications1.HTTPFilter.JSONPatches[jsonPatchesCount1].Path = jsonPatches3.Path
						appendModifications1.HTTPFilter.JSONPatches[jsonPatchesCount1].Value = jsonPatches3.Value
					}
				}
				if appendModificationsItem.HTTPFilter.Match == nil {
					appendModifications1.HTTPFilter.Match = nil
				} else {
					appendModifications1.HTTPFilter.Match = &tfTypes.MeshProxyPatchItemMatch{}
					appendModifications1.HTTPFilter.Match.ListenerName = types.StringPointerValue(appendModificationsItem.HTTPFilter.Match.ListenerName)
					if len(appendModificationsItem.HTTPFilter.Match.ListenerTags) > 0 {
						appendModifications1.HTTPFilter.Match.ListenerTags = make(map[string]types.String, len(appendModificationsItem.HTTPFilter.Match.ListenerTags))
						for key1, value4 := range appendModificationsItem.HTTPFilter.Match.ListenerTags {
							appendModifications1.HTTPFilter.Match.ListenerTags[key1] = types.StringValue(value4)
						}
					}
					appendModifications1.HTTPFilter.Match.Name = types.StringPointerValue(appendModificationsItem.HTTPFilter.Match.Name)
					appendModifications1.HTTPFilter.Match.Origin = types.StringPointerValue(appendModificationsItem.HTTPFilter.Match.Origin)
				}
				appendModifications1.HTTPFilter.Operation = types.StringValue(string(appendModificationsItem.HTTPFilter.Operation))
				appendModifications1.HTTPFilter.Value = types.StringPointerValue(appendModificationsItem.HTTPFilter.Value)
			}
			if appendModificationsItem.Listener == nil {
				appendModifications1.Listener = nil
			} else {
				appendModifications1.Listener = &tfTypes.Listener{}
				appendModifications1.Listener.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount2, jsonPatchesItem2 := range appendModificationsItem.Listener.JSONPatches {
					var jsonPatches5 tfTypes.JSONPatches
					jsonPatches5.From = types.StringPointerValue(jsonPatchesItem2.From)
					jsonPatches5.Op = types.StringValue(string(jsonPatchesItem2.Op))
					jsonPatches5.Path = types.StringValue(jsonPatchesItem2.Path)
					if jsonPatchesItem2.Value == nil {
						jsonPatches5.Value = types.StringNull()
					} else {
						valueResult2, _ := json.Marshal(jsonPatchesItem2.Value)
						jsonPatches5.Value = types.StringValue(string(valueResult2))
					}
					if jsonPatchesCount2+1 > len(appendModifications1.Listener.JSONPatches) {
						appendModifications1.Listener.JSONPatches = append(appendModifications1.Listener.JSONPatches, jsonPatches5)
					} else {
						appendModifications1.Listener.JSONPatches[jsonPatchesCount2].From = jsonPatches5.From
						appendModifications1.Listener.JSONPatches[jsonPatchesCount2].Op = jsonPatches5.Op
						appendModifications1.Listener.JSONPatches[jsonPatchesCount2].Path = jsonPatches5.Path
						appendModifications1.Listener.JSONPatches[jsonPatchesCount2].Value = jsonPatches5.Value
					}
				}
				if appendModificationsItem.Listener.Match == nil {
					appendModifications1.Listener.Match = nil
				} else {
					appendModifications1.Listener.Match = &tfTypes.MeshProxyPatchItemSpecMatch{}
					appendModifications1.Listener.Match.Name = types.StringPointerValue(appendModificationsItem.Listener.Match.Name)
					appendModifications1.Listener.Match.Origin = types.StringPointerValue(appendModificationsItem.Listener.Match.Origin)
					if len(appendModificationsItem.Listener.Match.Tags) > 0 {
						appendModifications1.Listener.Match.Tags = make(map[string]types.String, len(appendModificationsItem.Listener.Match.Tags))
						for key2, value7 := range appendModificationsItem.Listener.Match.Tags {
							appendModifications1.Listener.Match.Tags[key2] = types.StringValue(value7)
						}
					}
				}
				appendModifications1.Listener.Operation = types.StringValue(string(appendModificationsItem.Listener.Operation))
				appendModifications1.Listener.Value = types.StringPointerValue(appendModificationsItem.Listener.Value)
			}
			if appendModificationsItem.NetworkFilter == nil {
				appendModifications1.NetworkFilter = nil
			} else {
				appendModifications1.NetworkFilter = &tfTypes.HTTPFilter{}
				appendModifications1.NetworkFilter.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount3, jsonPatchesItem3 := range appendModificationsItem.NetworkFilter.JSONPatches {
					var jsonPatches7 tfTypes.JSONPatches
					jsonPatches7.From = types.StringPointerValue(jsonPatchesItem3.From)
					jsonPatches7.Op = types.StringValue(string(jsonPatchesItem3.Op))
					jsonPatches7.Path = types.StringValue(jsonPatchesItem3.Path)
					if jsonPatchesItem3.Value == nil {
						jsonPatches7.Value = types.StringNull()
					} else {
						valueResult3, _ := json.Marshal(jsonPatchesItem3.Value)
						jsonPatches7.Value = types.StringValue(string(valueResult3))
					}
					if jsonPatchesCount3+1 > len(appendModifications1.NetworkFilter.JSONPatches) {
						appendModifications1.NetworkFilter.JSONPatches = append(appendModifications1.NetworkFilter.JSONPatches, jsonPatches7)
					} else {
						appendModifications1.NetworkFilter.JSONPatches[jsonPatchesCount3].From = jsonPatches7.From
						appendModifications1.NetworkFilter.JSONPatches[jsonPatchesCount3].Op = jsonPatches7.Op
						appendModifications1.NetworkFilter.JSONPatches[jsonPatchesCount3].Path = jsonPatches7.Path
						appendModifications1.NetworkFilter.JSONPatches[jsonPatchesCount3].Value = jsonPatches7.Value
					}
				}
				if appendModificationsItem.NetworkFilter.Match == nil {
					appendModifications1.NetworkFilter.Match = nil
				} else {
					appendModifications1.NetworkFilter.Match = &tfTypes.MeshProxyPatchItemMatch{}
					appendModifications1.NetworkFilter.Match.ListenerName = types.StringPointerValue(appendModificationsItem.NetworkFilter.Match.ListenerName)
					if len(appendModificationsItem.NetworkFilter.Match.ListenerTags) > 0 {
						appendModifications1.NetworkFilter.Match.ListenerTags = make(map[string]types.String, len(appendModificationsItem.NetworkFilter.Match.ListenerTags))
						for key3, value10 := range appendModificationsItem.NetworkFilter.Match.ListenerTags {
							appendModifications1.NetworkFilter.Match.ListenerTags[key3] = types.StringValue(value10)
						}
					}
					appendModifications1.NetworkFilter.Match.Name = types.StringPointerValue(appendModificationsItem.NetworkFilter.Match.Name)
					appendModifications1.NetworkFilter.Match.Origin = types.StringPointerValue(appendModificationsItem.NetworkFilter.Match.Origin)
				}
				appendModifications1.NetworkFilter.Operation = types.StringValue(string(appendModificationsItem.NetworkFilter.Operation))
				appendModifications1.NetworkFilter.Value = types.StringPointerValue(appendModificationsItem.NetworkFilter.Value)
			}
			if appendModificationsItem.VirtualHost == nil {
				appendModifications1.VirtualHost = nil
			} else {
				appendModifications1.VirtualHost = &tfTypes.VirtualHost{}
				appendModifications1.VirtualHost.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount4, jsonPatchesItem4 := range appendModificationsItem.VirtualHost.JSONPatches {
					var jsonPatches9 tfTypes.JSONPatches
					jsonPatches9.From = types.StringPointerValue(jsonPatchesItem4.From)
					jsonPatches9.Op = types.StringValue(string(jsonPatchesItem4.Op))
					jsonPatches9.Path = types.StringValue(jsonPatchesItem4.Path)
					if jsonPatchesItem4.Value == nil {
						jsonPatches9.Value = types.StringNull()
					} else {
						valueResult4, _ := json.Marshal(jsonPatchesItem4.Value)
						jsonPatches9.Value = types.StringValue(string(valueResult4))
					}
					if jsonPatchesCount4+1 > len(appendModifications1.VirtualHost.JSONPatches) {
						appendModifications1.VirtualHost.JSONPatches = append(appendModifications1.VirtualHost.JSONPatches, jsonPatches9)
					} else {
						appendModifications1.VirtualHost.JSONPatches[jsonPatchesCount4].From = jsonPatches9.From
						appendModifications1.VirtualHost.JSONPatches[jsonPatchesCount4].Op = jsonPatches9.Op
						appendModifications1.VirtualHost.JSONPatches[jsonPatchesCount4].Path = jsonPatches9.Path
						appendModifications1.VirtualHost.JSONPatches[jsonPatchesCount4].Value = jsonPatches9.Value
					}
				}
				appendModifications1.VirtualHost.Match.Name = types.StringPointerValue(appendModificationsItem.VirtualHost.Match.Name)
				appendModifications1.VirtualHost.Match.Origin = types.StringPointerValue(appendModificationsItem.VirtualHost.Match.Origin)
				appendModifications1.VirtualHost.Match.RouteConfigurationName = types.StringPointerValue(appendModificationsItem.VirtualHost.Match.RouteConfigurationName)
				appendModifications1.VirtualHost.Operation = types.StringValue(string(appendModificationsItem.VirtualHost.Operation))
				appendModifications1.VirtualHost.Value = types.StringPointerValue(appendModificationsItem.VirtualHost.Value)
			}
			if appendModificationsCount+1 > len(r.Spec.Default.AppendModifications) {
				r.Spec.Default.AppendModifications = append(r.Spec.Default.AppendModifications, appendModifications1)
			} else {
				r.Spec.Default.AppendModifications[appendModificationsCount].Cluster = appendModifications1.Cluster
				r.Spec.Default.AppendModifications[appendModificationsCount].HTTPFilter = appendModifications1.HTTPFilter
				r.Spec.Default.AppendModifications[appendModificationsCount].Listener = appendModifications1.Listener
				r.Spec.Default.AppendModifications[appendModificationsCount].NetworkFilter = appendModifications1.NetworkFilter
				r.Spec.Default.AppendModifications[appendModificationsCount].VirtualHost = appendModifications1.VirtualHost
			}
		}
		if resp.Spec.TargetRef == nil {
			r.Spec.TargetRef = nil
		} else {
			r.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
			r.Spec.TargetRef.Kind = types.StringValue(string(resp.Spec.TargetRef.Kind))
			if len(resp.Spec.TargetRef.Labels) > 0 {
				r.Spec.TargetRef.Labels = make(map[string]types.String, len(resp.Spec.TargetRef.Labels))
				for key4, value14 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key4] = types.StringValue(value14)
				}
			}
			r.Spec.TargetRef.Mesh = types.StringPointerValue(resp.Spec.TargetRef.Mesh)
			r.Spec.TargetRef.Name = types.StringPointerValue(resp.Spec.TargetRef.Name)
			r.Spec.TargetRef.Namespace = types.StringPointerValue(resp.Spec.TargetRef.Namespace)
			r.Spec.TargetRef.ProxyTypes = make([]types.String, 0, len(resp.Spec.TargetRef.ProxyTypes))
			for _, v := range resp.Spec.TargetRef.ProxyTypes {
				r.Spec.TargetRef.ProxyTypes = append(r.Spec.TargetRef.ProxyTypes, types.StringValue(string(v)))
			}
			r.Spec.TargetRef.SectionName = types.StringPointerValue(resp.Spec.TargetRef.SectionName)
			if len(resp.Spec.TargetRef.Tags) > 0 {
				r.Spec.TargetRef.Tags = make(map[string]types.String, len(resp.Spec.TargetRef.Tags))
				for key5, value15 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key5] = types.StringValue(value15)
				}
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}
}
