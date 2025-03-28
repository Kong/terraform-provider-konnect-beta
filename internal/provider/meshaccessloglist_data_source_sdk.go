// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect-beta/internal/provider/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *MeshAccessLogListDataSourceModel) RefreshFromSharedMeshAccessLogList(resp *shared.MeshAccessLogList) {
	if resp != nil {
		r.Items = []tfTypes.MeshAccessLogItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items1 tfTypes.MeshAccessLogItem
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
			items1.Mesh = types.StringPointerValue(itemsItem.Mesh)
			if itemsItem.ModificationTime != nil {
				items1.ModificationTime = types.StringValue(itemsItem.ModificationTime.Format(time.RFC3339Nano))
			} else {
				items1.ModificationTime = types.StringNull()
			}
			items1.Name = types.StringValue(itemsItem.Name)
			items1.Spec.From = []tfTypes.From{}
			for fromCount, fromItem := range itemsItem.Spec.From {
				var from1 tfTypes.From
				from1.Default.Backends = []tfTypes.MeshAccessLogItemSpecFromBackends{}
				for backendsCount, backendsItem := range fromItem.Default.Backends {
					var backends1 tfTypes.MeshAccessLogItemSpecFromBackends
					if backendsItem.File == nil {
						backends1.File = nil
					} else {
						backends1.File = &tfTypes.File{}
						if backendsItem.File.Format == nil {
							backends1.File.Format = nil
						} else {
							backends1.File.Format = &tfTypes.Format{}
							backends1.File.Format.JSON = []tfTypes.JSON{}
							for jsonVarCount, jsonVarItem := range backendsItem.File.Format.JSON {
								var jsonVar1 tfTypes.JSON
								jsonVar1.Key = types.StringValue(jsonVarItem.Key)
								jsonVar1.Value = types.StringValue(jsonVarItem.Value)
								if jsonVarCount+1 > len(backends1.File.Format.JSON) {
									backends1.File.Format.JSON = append(backends1.File.Format.JSON, jsonVar1)
								} else {
									backends1.File.Format.JSON[jsonVarCount].Key = jsonVar1.Key
									backends1.File.Format.JSON[jsonVarCount].Value = jsonVar1.Value
								}
							}
							backends1.File.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem.File.Format.OmitEmptyValues)
							backends1.File.Format.Plain = types.StringPointerValue(backendsItem.File.Format.Plain)
							backends1.File.Format.Type = types.StringValue(string(backendsItem.File.Format.Type))
						}
						backends1.File.Path = types.StringValue(backendsItem.File.Path)
					}
					if backendsItem.OpenTelemetry == nil {
						backends1.OpenTelemetry = nil
					} else {
						backends1.OpenTelemetry = &tfTypes.MeshAccessLogItemSpecFromOpenTelemetry{}
						backends1.OpenTelemetry.Attributes = []tfTypes.JSON{}
						for attributesCount, attributesItem := range backendsItem.OpenTelemetry.Attributes {
							var attributes1 tfTypes.JSON
							attributes1.Key = types.StringValue(attributesItem.Key)
							attributes1.Value = types.StringValue(attributesItem.Value)
							if attributesCount+1 > len(backends1.OpenTelemetry.Attributes) {
								backends1.OpenTelemetry.Attributes = append(backends1.OpenTelemetry.Attributes, attributes1)
							} else {
								backends1.OpenTelemetry.Attributes[attributesCount].Key = attributes1.Key
								backends1.OpenTelemetry.Attributes[attributesCount].Value = attributes1.Value
							}
						}
						if backendsItem.OpenTelemetry.Body == nil {
							backends1.OpenTelemetry.Body = types.StringNull()
						} else {
							bodyResult, _ := json.Marshal(backendsItem.OpenTelemetry.Body)
							backends1.OpenTelemetry.Body = types.StringValue(string(bodyResult))
						}
						backends1.OpenTelemetry.Endpoint = types.StringValue(backendsItem.OpenTelemetry.Endpoint)
					}
					if backendsItem.TCP == nil {
						backends1.TCP = nil
					} else {
						backends1.TCP = &tfTypes.MeshAccessLogItemSpecFromTCP{}
						backends1.TCP.Address = types.StringValue(backendsItem.TCP.Address)
						if backendsItem.TCP.Format == nil {
							backends1.TCP.Format = nil
						} else {
							backends1.TCP.Format = &tfTypes.Format{}
							backends1.TCP.Format.JSON = []tfTypes.JSON{}
							for jsonVarCount1, jsonVarItem1 := range backendsItem.TCP.Format.JSON {
								var jsonVar2 tfTypes.JSON
								jsonVar2.Key = types.StringValue(jsonVarItem1.Key)
								jsonVar2.Value = types.StringValue(jsonVarItem1.Value)
								if jsonVarCount1+1 > len(backends1.TCP.Format.JSON) {
									backends1.TCP.Format.JSON = append(backends1.TCP.Format.JSON, jsonVar2)
								} else {
									backends1.TCP.Format.JSON[jsonVarCount1].Key = jsonVar2.Key
									backends1.TCP.Format.JSON[jsonVarCount1].Value = jsonVar2.Value
								}
							}
							backends1.TCP.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem.TCP.Format.OmitEmptyValues)
							backends1.TCP.Format.Plain = types.StringPointerValue(backendsItem.TCP.Format.Plain)
							backends1.TCP.Format.Type = types.StringValue(string(backendsItem.TCP.Format.Type))
						}
					}
					backends1.Type = types.StringValue(string(backendsItem.Type))
					if backendsCount+1 > len(from1.Default.Backends) {
						from1.Default.Backends = append(from1.Default.Backends, backends1)
					} else {
						from1.Default.Backends[backendsCount].File = backends1.File
						from1.Default.Backends[backendsCount].OpenTelemetry = backends1.OpenTelemetry
						from1.Default.Backends[backendsCount].TCP = backends1.TCP
						from1.Default.Backends[backendsCount].Type = backends1.Type
					}
				}
				from1.TargetRef.Kind = types.StringValue(string(fromItem.TargetRef.Kind))
				if len(fromItem.TargetRef.Labels) > 0 {
					from1.TargetRef.Labels = make(map[string]types.String, len(fromItem.TargetRef.Labels))
					for key4, value4 := range fromItem.TargetRef.Labels {
						from1.TargetRef.Labels[key4] = types.StringValue(value4)
					}
				}
				from1.TargetRef.Mesh = types.StringPointerValue(fromItem.TargetRef.Mesh)
				from1.TargetRef.Name = types.StringPointerValue(fromItem.TargetRef.Name)
				from1.TargetRef.Namespace = types.StringPointerValue(fromItem.TargetRef.Namespace)
				from1.TargetRef.ProxyTypes = make([]types.String, 0, len(fromItem.TargetRef.ProxyTypes))
				for _, v := range fromItem.TargetRef.ProxyTypes {
					from1.TargetRef.ProxyTypes = append(from1.TargetRef.ProxyTypes, types.StringValue(string(v)))
				}
				from1.TargetRef.SectionName = types.StringPointerValue(fromItem.TargetRef.SectionName)
				if len(fromItem.TargetRef.Tags) > 0 {
					from1.TargetRef.Tags = make(map[string]types.String, len(fromItem.TargetRef.Tags))
					for key5, value5 := range fromItem.TargetRef.Tags {
						from1.TargetRef.Tags[key5] = types.StringValue(value5)
					}
				}
				if fromCount+1 > len(items1.Spec.From) {
					items1.Spec.From = append(items1.Spec.From, from1)
				} else {
					items1.Spec.From[fromCount].Default = from1.Default
					items1.Spec.From[fromCount].TargetRef = from1.TargetRef
				}
			}
			items1.Spec.Rules = []tfTypes.Rules{}
			for rulesCount, rulesItem := range itemsItem.Spec.Rules {
				var rules1 tfTypes.Rules
				rules1.Default.Backends = []tfTypes.MeshAccessLogItemSpecFromBackends{}
				for backendsCount1, backendsItem1 := range rulesItem.Default.Backends {
					var backends3 tfTypes.MeshAccessLogItemSpecFromBackends
					if backendsItem1.File == nil {
						backends3.File = nil
					} else {
						backends3.File = &tfTypes.File{}
						if backendsItem1.File.Format == nil {
							backends3.File.Format = nil
						} else {
							backends3.File.Format = &tfTypes.Format{}
							backends3.File.Format.JSON = []tfTypes.JSON{}
							for jsonVarCount2, jsonVarItem2 := range backendsItem1.File.Format.JSON {
								var jsonVar3 tfTypes.JSON
								jsonVar3.Key = types.StringValue(jsonVarItem2.Key)
								jsonVar3.Value = types.StringValue(jsonVarItem2.Value)
								if jsonVarCount2+1 > len(backends3.File.Format.JSON) {
									backends3.File.Format.JSON = append(backends3.File.Format.JSON, jsonVar3)
								} else {
									backends3.File.Format.JSON[jsonVarCount2].Key = jsonVar3.Key
									backends3.File.Format.JSON[jsonVarCount2].Value = jsonVar3.Value
								}
							}
							backends3.File.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem1.File.Format.OmitEmptyValues)
							backends3.File.Format.Plain = types.StringPointerValue(backendsItem1.File.Format.Plain)
							backends3.File.Format.Type = types.StringValue(string(backendsItem1.File.Format.Type))
						}
						backends3.File.Path = types.StringValue(backendsItem1.File.Path)
					}
					if backendsItem1.OpenTelemetry == nil {
						backends3.OpenTelemetry = nil
					} else {
						backends3.OpenTelemetry = &tfTypes.MeshAccessLogItemSpecFromOpenTelemetry{}
						backends3.OpenTelemetry.Attributes = []tfTypes.JSON{}
						for attributesCount1, attributesItem1 := range backendsItem1.OpenTelemetry.Attributes {
							var attributes3 tfTypes.JSON
							attributes3.Key = types.StringValue(attributesItem1.Key)
							attributes3.Value = types.StringValue(attributesItem1.Value)
							if attributesCount1+1 > len(backends3.OpenTelemetry.Attributes) {
								backends3.OpenTelemetry.Attributes = append(backends3.OpenTelemetry.Attributes, attributes3)
							} else {
								backends3.OpenTelemetry.Attributes[attributesCount1].Key = attributes3.Key
								backends3.OpenTelemetry.Attributes[attributesCount1].Value = attributes3.Value
							}
						}
						if backendsItem1.OpenTelemetry.Body == nil {
							backends3.OpenTelemetry.Body = types.StringNull()
						} else {
							bodyResult1, _ := json.Marshal(backendsItem1.OpenTelemetry.Body)
							backends3.OpenTelemetry.Body = types.StringValue(string(bodyResult1))
						}
						backends3.OpenTelemetry.Endpoint = types.StringValue(backendsItem1.OpenTelemetry.Endpoint)
					}
					if backendsItem1.TCP == nil {
						backends3.TCP = nil
					} else {
						backends3.TCP = &tfTypes.MeshAccessLogItemSpecFromTCP{}
						backends3.TCP.Address = types.StringValue(backendsItem1.TCP.Address)
						if backendsItem1.TCP.Format == nil {
							backends3.TCP.Format = nil
						} else {
							backends3.TCP.Format = &tfTypes.Format{}
							backends3.TCP.Format.JSON = []tfTypes.JSON{}
							for jsonVarCount3, jsonVarItem3 := range backendsItem1.TCP.Format.JSON {
								var jsonVar4 tfTypes.JSON
								jsonVar4.Key = types.StringValue(jsonVarItem3.Key)
								jsonVar4.Value = types.StringValue(jsonVarItem3.Value)
								if jsonVarCount3+1 > len(backends3.TCP.Format.JSON) {
									backends3.TCP.Format.JSON = append(backends3.TCP.Format.JSON, jsonVar4)
								} else {
									backends3.TCP.Format.JSON[jsonVarCount3].Key = jsonVar4.Key
									backends3.TCP.Format.JSON[jsonVarCount3].Value = jsonVar4.Value
								}
							}
							backends3.TCP.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem1.TCP.Format.OmitEmptyValues)
							backends3.TCP.Format.Plain = types.StringPointerValue(backendsItem1.TCP.Format.Plain)
							backends3.TCP.Format.Type = types.StringValue(string(backendsItem1.TCP.Format.Type))
						}
					}
					backends3.Type = types.StringValue(string(backendsItem1.Type))
					if backendsCount1+1 > len(rules1.Default.Backends) {
						rules1.Default.Backends = append(rules1.Default.Backends, backends3)
					} else {
						rules1.Default.Backends[backendsCount1].File = backends3.File
						rules1.Default.Backends[backendsCount1].OpenTelemetry = backends3.OpenTelemetry
						rules1.Default.Backends[backendsCount1].TCP = backends3.TCP
						rules1.Default.Backends[backendsCount1].Type = backends3.Type
					}
				}
				if rulesCount+1 > len(items1.Spec.Rules) {
					items1.Spec.Rules = append(items1.Spec.Rules, rules1)
				} else {
					items1.Spec.Rules[rulesCount].Default = rules1.Default
				}
			}
			if itemsItem.Spec.TargetRef == nil {
				items1.Spec.TargetRef = nil
			} else {
				items1.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
				items1.Spec.TargetRef.Kind = types.StringValue(string(itemsItem.Spec.TargetRef.Kind))
				if len(itemsItem.Spec.TargetRef.Labels) > 0 {
					items1.Spec.TargetRef.Labels = make(map[string]types.String, len(itemsItem.Spec.TargetRef.Labels))
					for key9, value9 := range itemsItem.Spec.TargetRef.Labels {
						items1.Spec.TargetRef.Labels[key9] = types.StringValue(value9)
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
					items1.Spec.TargetRef.Tags = make(map[string]types.String, len(itemsItem.Spec.TargetRef.Tags))
					for key10, value10 := range itemsItem.Spec.TargetRef.Tags {
						items1.Spec.TargetRef.Tags[key10] = types.StringValue(value10)
					}
				}
			}
			items1.Spec.To = []tfTypes.From{}
			for toCount, toItem := range itemsItem.Spec.To {
				var to1 tfTypes.From
				to1.Default.Backends = []tfTypes.MeshAccessLogItemSpecFromBackends{}
				for backendsCount2, backendsItem2 := range toItem.Default.Backends {
					var backends5 tfTypes.MeshAccessLogItemSpecFromBackends
					if backendsItem2.File == nil {
						backends5.File = nil
					} else {
						backends5.File = &tfTypes.File{}
						if backendsItem2.File.Format == nil {
							backends5.File.Format = nil
						} else {
							backends5.File.Format = &tfTypes.Format{}
							backends5.File.Format.JSON = []tfTypes.JSON{}
							for jsonVarCount4, jsonVarItem4 := range backendsItem2.File.Format.JSON {
								var jsonVar5 tfTypes.JSON
								jsonVar5.Key = types.StringValue(jsonVarItem4.Key)
								jsonVar5.Value = types.StringValue(jsonVarItem4.Value)
								if jsonVarCount4+1 > len(backends5.File.Format.JSON) {
									backends5.File.Format.JSON = append(backends5.File.Format.JSON, jsonVar5)
								} else {
									backends5.File.Format.JSON[jsonVarCount4].Key = jsonVar5.Key
									backends5.File.Format.JSON[jsonVarCount4].Value = jsonVar5.Value
								}
							}
							backends5.File.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem2.File.Format.OmitEmptyValues)
							backends5.File.Format.Plain = types.StringPointerValue(backendsItem2.File.Format.Plain)
							backends5.File.Format.Type = types.StringValue(string(backendsItem2.File.Format.Type))
						}
						backends5.File.Path = types.StringValue(backendsItem2.File.Path)
					}
					if backendsItem2.OpenTelemetry == nil {
						backends5.OpenTelemetry = nil
					} else {
						backends5.OpenTelemetry = &tfTypes.MeshAccessLogItemSpecFromOpenTelemetry{}
						backends5.OpenTelemetry.Attributes = []tfTypes.JSON{}
						for attributesCount2, attributesItem2 := range backendsItem2.OpenTelemetry.Attributes {
							var attributes5 tfTypes.JSON
							attributes5.Key = types.StringValue(attributesItem2.Key)
							attributes5.Value = types.StringValue(attributesItem2.Value)
							if attributesCount2+1 > len(backends5.OpenTelemetry.Attributes) {
								backends5.OpenTelemetry.Attributes = append(backends5.OpenTelemetry.Attributes, attributes5)
							} else {
								backends5.OpenTelemetry.Attributes[attributesCount2].Key = attributes5.Key
								backends5.OpenTelemetry.Attributes[attributesCount2].Value = attributes5.Value
							}
						}
						if backendsItem2.OpenTelemetry.Body == nil {
							backends5.OpenTelemetry.Body = types.StringNull()
						} else {
							bodyResult2, _ := json.Marshal(backendsItem2.OpenTelemetry.Body)
							backends5.OpenTelemetry.Body = types.StringValue(string(bodyResult2))
						}
						backends5.OpenTelemetry.Endpoint = types.StringValue(backendsItem2.OpenTelemetry.Endpoint)
					}
					if backendsItem2.TCP == nil {
						backends5.TCP = nil
					} else {
						backends5.TCP = &tfTypes.MeshAccessLogItemSpecFromTCP{}
						backends5.TCP.Address = types.StringValue(backendsItem2.TCP.Address)
						if backendsItem2.TCP.Format == nil {
							backends5.TCP.Format = nil
						} else {
							backends5.TCP.Format = &tfTypes.Format{}
							backends5.TCP.Format.JSON = []tfTypes.JSON{}
							for jsonVarCount5, jsonVarItem5 := range backendsItem2.TCP.Format.JSON {
								var jsonVar6 tfTypes.JSON
								jsonVar6.Key = types.StringValue(jsonVarItem5.Key)
								jsonVar6.Value = types.StringValue(jsonVarItem5.Value)
								if jsonVarCount5+1 > len(backends5.TCP.Format.JSON) {
									backends5.TCP.Format.JSON = append(backends5.TCP.Format.JSON, jsonVar6)
								} else {
									backends5.TCP.Format.JSON[jsonVarCount5].Key = jsonVar6.Key
									backends5.TCP.Format.JSON[jsonVarCount5].Value = jsonVar6.Value
								}
							}
							backends5.TCP.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem2.TCP.Format.OmitEmptyValues)
							backends5.TCP.Format.Plain = types.StringPointerValue(backendsItem2.TCP.Format.Plain)
							backends5.TCP.Format.Type = types.StringValue(string(backendsItem2.TCP.Format.Type))
						}
					}
					backends5.Type = types.StringValue(string(backendsItem2.Type))
					if backendsCount2+1 > len(to1.Default.Backends) {
						to1.Default.Backends = append(to1.Default.Backends, backends5)
					} else {
						to1.Default.Backends[backendsCount2].File = backends5.File
						to1.Default.Backends[backendsCount2].OpenTelemetry = backends5.OpenTelemetry
						to1.Default.Backends[backendsCount2].TCP = backends5.TCP
						to1.Default.Backends[backendsCount2].Type = backends5.Type
					}
				}
				to1.TargetRef.Kind = types.StringValue(string(toItem.TargetRef.Kind))
				if len(toItem.TargetRef.Labels) > 0 {
					to1.TargetRef.Labels = make(map[string]types.String, len(toItem.TargetRef.Labels))
					for key14, value14 := range toItem.TargetRef.Labels {
						to1.TargetRef.Labels[key14] = types.StringValue(value14)
					}
				}
				to1.TargetRef.Mesh = types.StringPointerValue(toItem.TargetRef.Mesh)
				to1.TargetRef.Name = types.StringPointerValue(toItem.TargetRef.Name)
				to1.TargetRef.Namespace = types.StringPointerValue(toItem.TargetRef.Namespace)
				to1.TargetRef.ProxyTypes = make([]types.String, 0, len(toItem.TargetRef.ProxyTypes))
				for _, v := range toItem.TargetRef.ProxyTypes {
					to1.TargetRef.ProxyTypes = append(to1.TargetRef.ProxyTypes, types.StringValue(string(v)))
				}
				to1.TargetRef.SectionName = types.StringPointerValue(toItem.TargetRef.SectionName)
				if len(toItem.TargetRef.Tags) > 0 {
					to1.TargetRef.Tags = make(map[string]types.String, len(toItem.TargetRef.Tags))
					for key15, value15 := range toItem.TargetRef.Tags {
						to1.TargetRef.Tags[key15] = types.StringValue(value15)
					}
				}
				if toCount+1 > len(items1.Spec.To) {
					items1.Spec.To = append(items1.Spec.To, to1)
				} else {
					items1.Spec.To[toCount].Default = to1.Default
					items1.Spec.To[toCount].TargetRef = to1.TargetRef
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
