// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect-beta/internal/provider/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"time"
)

func (r *MeshRateLimitResourceModel) ToSharedMeshRateLimitItemInput() *shared.MeshRateLimitItemInput {
	typeVar := shared.MeshRateLimitItemType(r.Type.ValueString())
	mesh := new(string)
	if !r.Mesh.IsUnknown() && !r.Mesh.IsNull() {
		*mesh = r.Mesh.ValueString()
	} else {
		mesh = nil
	}
	var name string
	name = r.Name.ValueString()

	labels := make(map[string]string)
	for labelsKey, labelsValue := range r.Labels {
		var labelsInst string
		labelsInst = labelsValue.ValueString()

		labels[labelsKey] = labelsInst
	}
	var from []shared.MeshRateLimitItemFrom = []shared.MeshRateLimitItemFrom{}
	for _, fromItem := range r.Spec.From {
		var defaultVar *shared.MeshRateLimitItemDefault
		if fromItem.Default != nil {
			var local *shared.Local
			if fromItem.Default.Local != nil {
				var http *shared.MeshRateLimitItemHTTP
				if fromItem.Default.Local.HTTP != nil {
					disabled := new(bool)
					if !fromItem.Default.Local.HTTP.Disabled.IsUnknown() && !fromItem.Default.Local.HTTP.Disabled.IsNull() {
						*disabled = fromItem.Default.Local.HTTP.Disabled.ValueBool()
					} else {
						disabled = nil
					}
					var onRateLimit *shared.MeshRateLimitItemSpecFromOnRateLimit
					if fromItem.Default.Local.HTTP.OnRateLimit != nil {
						var headers *shared.MeshRateLimitItemHeaders
						if fromItem.Default.Local.HTTP.OnRateLimit.Headers != nil {
							var add []shared.MeshRateLimitItemAdd = []shared.MeshRateLimitItemAdd{}
							for _, addItem := range fromItem.Default.Local.HTTP.OnRateLimit.Headers.Add {
								var name1 string
								name1 = addItem.Name.ValueString()

								var value string
								value = addItem.Value.ValueString()

								add = append(add, shared.MeshRateLimitItemAdd{
									Name:  name1,
									Value: value,
								})
							}
							var set []shared.MeshRateLimitItemSet = []shared.MeshRateLimitItemSet{}
							for _, setItem := range fromItem.Default.Local.HTTP.OnRateLimit.Headers.Set {
								var name2 string
								name2 = setItem.Name.ValueString()

								var value1 string
								value1 = setItem.Value.ValueString()

								set = append(set, shared.MeshRateLimitItemSet{
									Name:  name2,
									Value: value1,
								})
							}
							headers = &shared.MeshRateLimitItemHeaders{
								Add: add,
								Set: set,
							}
						}
						status := new(int)
						if !fromItem.Default.Local.HTTP.OnRateLimit.Status.IsUnknown() && !fromItem.Default.Local.HTTP.OnRateLimit.Status.IsNull() {
							*status = int(fromItem.Default.Local.HTTP.OnRateLimit.Status.ValueInt32())
						} else {
							status = nil
						}
						onRateLimit = &shared.MeshRateLimitItemSpecFromOnRateLimit{
							Headers: headers,
							Status:  status,
						}
					}
					var requestRate *shared.MeshRateLimitItemSpecFromRequestRate
					if fromItem.Default.Local.HTTP.RequestRate != nil {
						var interval string
						interval = fromItem.Default.Local.HTTP.RequestRate.Interval.ValueString()

						var num int
						num = int(fromItem.Default.Local.HTTP.RequestRate.Num.ValueInt32())

						requestRate = &shared.MeshRateLimitItemSpecFromRequestRate{
							Interval: interval,
							Num:      num,
						}
					}
					http = &shared.MeshRateLimitItemHTTP{
						Disabled:    disabled,
						OnRateLimit: onRateLimit,
						RequestRate: requestRate,
					}
				}
				var tcp *shared.MeshRateLimitItemTCP
				if fromItem.Default.Local.TCP != nil {
					var connectionRate *shared.ConnectionRate
					if fromItem.Default.Local.TCP.ConnectionRate != nil {
						var interval1 string
						interval1 = fromItem.Default.Local.TCP.ConnectionRate.Interval.ValueString()

						var num1 int
						num1 = int(fromItem.Default.Local.TCP.ConnectionRate.Num.ValueInt32())

						connectionRate = &shared.ConnectionRate{
							Interval: interval1,
							Num:      num1,
						}
					}
					disabled1 := new(bool)
					if !fromItem.Default.Local.TCP.Disabled.IsUnknown() && !fromItem.Default.Local.TCP.Disabled.IsNull() {
						*disabled1 = fromItem.Default.Local.TCP.Disabled.ValueBool()
					} else {
						disabled1 = nil
					}
					tcp = &shared.MeshRateLimitItemTCP{
						ConnectionRate: connectionRate,
						Disabled:       disabled1,
					}
				}
				local = &shared.Local{
					HTTP: http,
					TCP:  tcp,
				}
			}
			defaultVar = &shared.MeshRateLimitItemDefault{
				Local: local,
			}
		}
		kind := shared.MeshRateLimitItemSpecKind(fromItem.TargetRef.Kind.ValueString())
		labels1 := make(map[string]string)
		for labelsKey1, labelsValue1 := range fromItem.TargetRef.Labels {
			var labelsInst1 string
			labelsInst1 = labelsValue1.ValueString()

			labels1[labelsKey1] = labelsInst1
		}
		mesh1 := new(string)
		if !fromItem.TargetRef.Mesh.IsUnknown() && !fromItem.TargetRef.Mesh.IsNull() {
			*mesh1 = fromItem.TargetRef.Mesh.ValueString()
		} else {
			mesh1 = nil
		}
		name3 := new(string)
		if !fromItem.TargetRef.Name.IsUnknown() && !fromItem.TargetRef.Name.IsNull() {
			*name3 = fromItem.TargetRef.Name.ValueString()
		} else {
			name3 = nil
		}
		namespace := new(string)
		if !fromItem.TargetRef.Namespace.IsUnknown() && !fromItem.TargetRef.Namespace.IsNull() {
			*namespace = fromItem.TargetRef.Namespace.ValueString()
		} else {
			namespace = nil
		}
		var proxyTypes []shared.MeshRateLimitItemSpecProxyTypes = []shared.MeshRateLimitItemSpecProxyTypes{}
		for _, proxyTypesItem := range fromItem.TargetRef.ProxyTypes {
			proxyTypes = append(proxyTypes, shared.MeshRateLimitItemSpecProxyTypes(proxyTypesItem.ValueString()))
		}
		sectionName := new(string)
		if !fromItem.TargetRef.SectionName.IsUnknown() && !fromItem.TargetRef.SectionName.IsNull() {
			*sectionName = fromItem.TargetRef.SectionName.ValueString()
		} else {
			sectionName = nil
		}
		tags := make(map[string]string)
		for tagsKey, tagsValue := range fromItem.TargetRef.Tags {
			var tagsInst string
			tagsInst = tagsValue.ValueString()

			tags[tagsKey] = tagsInst
		}
		targetRef := shared.MeshRateLimitItemSpecTargetRef{
			Kind:        kind,
			Labels:      labels1,
			Mesh:        mesh1,
			Name:        name3,
			Namespace:   namespace,
			ProxyTypes:  proxyTypes,
			SectionName: sectionName,
			Tags:        tags,
		}
		from = append(from, shared.MeshRateLimitItemFrom{
			Default:   defaultVar,
			TargetRef: targetRef,
		})
	}
	var rules []shared.MeshRateLimitItemRules = []shared.MeshRateLimitItemRules{}
	for _, rulesItem := range r.Spec.Rules {
		var default1 *shared.MeshRateLimitItemSpecDefault
		if rulesItem.Default != nil {
			var local1 *shared.MeshRateLimitItemLocal
			if rulesItem.Default.Local != nil {
				var http1 *shared.MeshRateLimitItemSpecHTTP
				if rulesItem.Default.Local.HTTP != nil {
					disabled2 := new(bool)
					if !rulesItem.Default.Local.HTTP.Disabled.IsUnknown() && !rulesItem.Default.Local.HTTP.Disabled.IsNull() {
						*disabled2 = rulesItem.Default.Local.HTTP.Disabled.ValueBool()
					} else {
						disabled2 = nil
					}
					var onRateLimit1 *shared.MeshRateLimitItemOnRateLimit
					if rulesItem.Default.Local.HTTP.OnRateLimit != nil {
						var headers1 *shared.MeshRateLimitItemSpecHeaders
						if rulesItem.Default.Local.HTTP.OnRateLimit.Headers != nil {
							var add1 []shared.MeshRateLimitItemSpecAdd = []shared.MeshRateLimitItemSpecAdd{}
							for _, addItem1 := range rulesItem.Default.Local.HTTP.OnRateLimit.Headers.Add {
								var name4 string
								name4 = addItem1.Name.ValueString()

								var value2 string
								value2 = addItem1.Value.ValueString()

								add1 = append(add1, shared.MeshRateLimitItemSpecAdd{
									Name:  name4,
									Value: value2,
								})
							}
							var set1 []shared.MeshRateLimitItemSpecSet = []shared.MeshRateLimitItemSpecSet{}
							for _, setItem1 := range rulesItem.Default.Local.HTTP.OnRateLimit.Headers.Set {
								var name5 string
								name5 = setItem1.Name.ValueString()

								var value3 string
								value3 = setItem1.Value.ValueString()

								set1 = append(set1, shared.MeshRateLimitItemSpecSet{
									Name:  name5,
									Value: value3,
								})
							}
							headers1 = &shared.MeshRateLimitItemSpecHeaders{
								Add: add1,
								Set: set1,
							}
						}
						status1 := new(int)
						if !rulesItem.Default.Local.HTTP.OnRateLimit.Status.IsUnknown() && !rulesItem.Default.Local.HTTP.OnRateLimit.Status.IsNull() {
							*status1 = int(rulesItem.Default.Local.HTTP.OnRateLimit.Status.ValueInt32())
						} else {
							status1 = nil
						}
						onRateLimit1 = &shared.MeshRateLimitItemOnRateLimit{
							Headers: headers1,
							Status:  status1,
						}
					}
					var requestRate1 *shared.MeshRateLimitItemRequestRate
					if rulesItem.Default.Local.HTTP.RequestRate != nil {
						var interval2 string
						interval2 = rulesItem.Default.Local.HTTP.RequestRate.Interval.ValueString()

						var num2 int
						num2 = int(rulesItem.Default.Local.HTTP.RequestRate.Num.ValueInt32())

						requestRate1 = &shared.MeshRateLimitItemRequestRate{
							Interval: interval2,
							Num:      num2,
						}
					}
					http1 = &shared.MeshRateLimitItemSpecHTTP{
						Disabled:    disabled2,
						OnRateLimit: onRateLimit1,
						RequestRate: requestRate1,
					}
				}
				var tcp1 *shared.MeshRateLimitItemSpecTCP
				if rulesItem.Default.Local.TCP != nil {
					var connectionRate1 *shared.MeshRateLimitItemConnectionRate
					if rulesItem.Default.Local.TCP.ConnectionRate != nil {
						var interval3 string
						interval3 = rulesItem.Default.Local.TCP.ConnectionRate.Interval.ValueString()

						var num3 int
						num3 = int(rulesItem.Default.Local.TCP.ConnectionRate.Num.ValueInt32())

						connectionRate1 = &shared.MeshRateLimitItemConnectionRate{
							Interval: interval3,
							Num:      num3,
						}
					}
					disabled3 := new(bool)
					if !rulesItem.Default.Local.TCP.Disabled.IsUnknown() && !rulesItem.Default.Local.TCP.Disabled.IsNull() {
						*disabled3 = rulesItem.Default.Local.TCP.Disabled.ValueBool()
					} else {
						disabled3 = nil
					}
					tcp1 = &shared.MeshRateLimitItemSpecTCP{
						ConnectionRate: connectionRate1,
						Disabled:       disabled3,
					}
				}
				local1 = &shared.MeshRateLimitItemLocal{
					HTTP: http1,
					TCP:  tcp1,
				}
			}
			default1 = &shared.MeshRateLimitItemSpecDefault{
				Local: local1,
			}
		}
		rules = append(rules, shared.MeshRateLimitItemRules{
			Default: default1,
		})
	}
	var targetRef1 *shared.MeshRateLimitItemTargetRef
	if r.Spec.TargetRef != nil {
		kind1 := shared.MeshRateLimitItemKind(r.Spec.TargetRef.Kind.ValueString())
		labels2 := make(map[string]string)
		for labelsKey2, labelsValue2 := range r.Spec.TargetRef.Labels {
			var labelsInst2 string
			labelsInst2 = labelsValue2.ValueString()

			labels2[labelsKey2] = labelsInst2
		}
		mesh2 := new(string)
		if !r.Spec.TargetRef.Mesh.IsUnknown() && !r.Spec.TargetRef.Mesh.IsNull() {
			*mesh2 = r.Spec.TargetRef.Mesh.ValueString()
		} else {
			mesh2 = nil
		}
		name6 := new(string)
		if !r.Spec.TargetRef.Name.IsUnknown() && !r.Spec.TargetRef.Name.IsNull() {
			*name6 = r.Spec.TargetRef.Name.ValueString()
		} else {
			name6 = nil
		}
		namespace1 := new(string)
		if !r.Spec.TargetRef.Namespace.IsUnknown() && !r.Spec.TargetRef.Namespace.IsNull() {
			*namespace1 = r.Spec.TargetRef.Namespace.ValueString()
		} else {
			namespace1 = nil
		}
		var proxyTypes1 []shared.MeshRateLimitItemProxyTypes = []shared.MeshRateLimitItemProxyTypes{}
		for _, proxyTypesItem1 := range r.Spec.TargetRef.ProxyTypes {
			proxyTypes1 = append(proxyTypes1, shared.MeshRateLimitItemProxyTypes(proxyTypesItem1.ValueString()))
		}
		sectionName1 := new(string)
		if !r.Spec.TargetRef.SectionName.IsUnknown() && !r.Spec.TargetRef.SectionName.IsNull() {
			*sectionName1 = r.Spec.TargetRef.SectionName.ValueString()
		} else {
			sectionName1 = nil
		}
		tags1 := make(map[string]string)
		for tagsKey1, tagsValue1 := range r.Spec.TargetRef.Tags {
			var tagsInst1 string
			tagsInst1 = tagsValue1.ValueString()

			tags1[tagsKey1] = tagsInst1
		}
		targetRef1 = &shared.MeshRateLimitItemTargetRef{
			Kind:        kind1,
			Labels:      labels2,
			Mesh:        mesh2,
			Name:        name6,
			Namespace:   namespace1,
			ProxyTypes:  proxyTypes1,
			SectionName: sectionName1,
			Tags:        tags1,
		}
	}
	var to []shared.MeshRateLimitItemTo = []shared.MeshRateLimitItemTo{}
	for _, toItem := range r.Spec.To {
		var default2 *shared.MeshRateLimitItemSpecToDefault
		if toItem.Default != nil {
			var local2 *shared.MeshRateLimitItemSpecLocal
			if toItem.Default.Local != nil {
				var http2 *shared.MeshRateLimitItemSpecToHTTP
				if toItem.Default.Local.HTTP != nil {
					disabled4 := new(bool)
					if !toItem.Default.Local.HTTP.Disabled.IsUnknown() && !toItem.Default.Local.HTTP.Disabled.IsNull() {
						*disabled4 = toItem.Default.Local.HTTP.Disabled.ValueBool()
					} else {
						disabled4 = nil
					}
					var onRateLimit2 *shared.MeshRateLimitItemSpecOnRateLimit
					if toItem.Default.Local.HTTP.OnRateLimit != nil {
						var headers2 *shared.MeshRateLimitItemSpecToHeaders
						if toItem.Default.Local.HTTP.OnRateLimit.Headers != nil {
							var add2 []shared.MeshRateLimitItemSpecToAdd = []shared.MeshRateLimitItemSpecToAdd{}
							for _, addItem2 := range toItem.Default.Local.HTTP.OnRateLimit.Headers.Add {
								var name7 string
								name7 = addItem2.Name.ValueString()

								var value4 string
								value4 = addItem2.Value.ValueString()

								add2 = append(add2, shared.MeshRateLimitItemSpecToAdd{
									Name:  name7,
									Value: value4,
								})
							}
							var set2 []shared.MeshRateLimitItemSpecToSet = []shared.MeshRateLimitItemSpecToSet{}
							for _, setItem2 := range toItem.Default.Local.HTTP.OnRateLimit.Headers.Set {
								var name8 string
								name8 = setItem2.Name.ValueString()

								var value5 string
								value5 = setItem2.Value.ValueString()

								set2 = append(set2, shared.MeshRateLimitItemSpecToSet{
									Name:  name8,
									Value: value5,
								})
							}
							headers2 = &shared.MeshRateLimitItemSpecToHeaders{
								Add: add2,
								Set: set2,
							}
						}
						status2 := new(int)
						if !toItem.Default.Local.HTTP.OnRateLimit.Status.IsUnknown() && !toItem.Default.Local.HTTP.OnRateLimit.Status.IsNull() {
							*status2 = int(toItem.Default.Local.HTTP.OnRateLimit.Status.ValueInt32())
						} else {
							status2 = nil
						}
						onRateLimit2 = &shared.MeshRateLimitItemSpecOnRateLimit{
							Headers: headers2,
							Status:  status2,
						}
					}
					var requestRate2 *shared.MeshRateLimitItemSpecRequestRate
					if toItem.Default.Local.HTTP.RequestRate != nil {
						var interval4 string
						interval4 = toItem.Default.Local.HTTP.RequestRate.Interval.ValueString()

						var num4 int
						num4 = int(toItem.Default.Local.HTTP.RequestRate.Num.ValueInt32())

						requestRate2 = &shared.MeshRateLimitItemSpecRequestRate{
							Interval: interval4,
							Num:      num4,
						}
					}
					http2 = &shared.MeshRateLimitItemSpecToHTTP{
						Disabled:    disabled4,
						OnRateLimit: onRateLimit2,
						RequestRate: requestRate2,
					}
				}
				var tcp2 *shared.MeshRateLimitItemSpecToTCP
				if toItem.Default.Local.TCP != nil {
					var connectionRate2 *shared.MeshRateLimitItemSpecConnectionRate
					if toItem.Default.Local.TCP.ConnectionRate != nil {
						var interval5 string
						interval5 = toItem.Default.Local.TCP.ConnectionRate.Interval.ValueString()

						var num5 int
						num5 = int(toItem.Default.Local.TCP.ConnectionRate.Num.ValueInt32())

						connectionRate2 = &shared.MeshRateLimitItemSpecConnectionRate{
							Interval: interval5,
							Num:      num5,
						}
					}
					disabled5 := new(bool)
					if !toItem.Default.Local.TCP.Disabled.IsUnknown() && !toItem.Default.Local.TCP.Disabled.IsNull() {
						*disabled5 = toItem.Default.Local.TCP.Disabled.ValueBool()
					} else {
						disabled5 = nil
					}
					tcp2 = &shared.MeshRateLimitItemSpecToTCP{
						ConnectionRate: connectionRate2,
						Disabled:       disabled5,
					}
				}
				local2 = &shared.MeshRateLimitItemSpecLocal{
					HTTP: http2,
					TCP:  tcp2,
				}
			}
			default2 = &shared.MeshRateLimitItemSpecToDefault{
				Local: local2,
			}
		}
		kind2 := shared.MeshRateLimitItemSpecToKind(toItem.TargetRef.Kind.ValueString())
		labels3 := make(map[string]string)
		for labelsKey3, labelsValue3 := range toItem.TargetRef.Labels {
			var labelsInst3 string
			labelsInst3 = labelsValue3.ValueString()

			labels3[labelsKey3] = labelsInst3
		}
		mesh3 := new(string)
		if !toItem.TargetRef.Mesh.IsUnknown() && !toItem.TargetRef.Mesh.IsNull() {
			*mesh3 = toItem.TargetRef.Mesh.ValueString()
		} else {
			mesh3 = nil
		}
		name9 := new(string)
		if !toItem.TargetRef.Name.IsUnknown() && !toItem.TargetRef.Name.IsNull() {
			*name9 = toItem.TargetRef.Name.ValueString()
		} else {
			name9 = nil
		}
		namespace2 := new(string)
		if !toItem.TargetRef.Namespace.IsUnknown() && !toItem.TargetRef.Namespace.IsNull() {
			*namespace2 = toItem.TargetRef.Namespace.ValueString()
		} else {
			namespace2 = nil
		}
		var proxyTypes2 []shared.MeshRateLimitItemSpecToProxyTypes = []shared.MeshRateLimitItemSpecToProxyTypes{}
		for _, proxyTypesItem2 := range toItem.TargetRef.ProxyTypes {
			proxyTypes2 = append(proxyTypes2, shared.MeshRateLimitItemSpecToProxyTypes(proxyTypesItem2.ValueString()))
		}
		sectionName2 := new(string)
		if !toItem.TargetRef.SectionName.IsUnknown() && !toItem.TargetRef.SectionName.IsNull() {
			*sectionName2 = toItem.TargetRef.SectionName.ValueString()
		} else {
			sectionName2 = nil
		}
		tags2 := make(map[string]string)
		for tagsKey2, tagsValue2 := range toItem.TargetRef.Tags {
			var tagsInst2 string
			tagsInst2 = tagsValue2.ValueString()

			tags2[tagsKey2] = tagsInst2
		}
		targetRef2 := shared.MeshRateLimitItemSpecToTargetRef{
			Kind:        kind2,
			Labels:      labels3,
			Mesh:        mesh3,
			Name:        name9,
			Namespace:   namespace2,
			ProxyTypes:  proxyTypes2,
			SectionName: sectionName2,
			Tags:        tags2,
		}
		to = append(to, shared.MeshRateLimitItemTo{
			Default:   default2,
			TargetRef: targetRef2,
		})
	}
	spec := shared.MeshRateLimitItemSpec{
		From:      from,
		Rules:     rules,
		TargetRef: targetRef1,
		To:        to,
	}
	out := shared.MeshRateLimitItemInput{
		Type:   typeVar,
		Mesh:   mesh,
		Name:   name,
		Labels: labels,
		Spec:   spec,
	}
	return &out
}

func (r *MeshRateLimitResourceModel) RefreshFromSharedMeshRateLimitCreateOrUpdateSuccessResponse(resp *shared.MeshRateLimitCreateOrUpdateSuccessResponse) {
	if resp != nil {
		r.Warnings = make([]types.String, 0, len(resp.Warnings))
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}
}

func (r *MeshRateLimitResourceModel) RefreshFromSharedMeshRateLimitItem(resp *shared.MeshRateLimitItem) {
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
		r.Spec.From = []tfTypes.MeshRateLimitItemFrom{}
		if len(r.Spec.From) > len(resp.Spec.From) {
			r.Spec.From = r.Spec.From[:len(resp.Spec.From)]
		}
		for fromCount, fromItem := range resp.Spec.From {
			var from1 tfTypes.MeshRateLimitItemFrom
			if fromItem.Default == nil {
				from1.Default = nil
			} else {
				from1.Default = &tfTypes.MeshRateLimitItemDefault{}
				if fromItem.Default.Local == nil {
					from1.Default.Local = nil
				} else {
					from1.Default.Local = &tfTypes.Local{}
					if fromItem.Default.Local.HTTP == nil {
						from1.Default.Local.HTTP = nil
					} else {
						from1.Default.Local.HTTP = &tfTypes.MeshRateLimitItemHTTP{}
						from1.Default.Local.HTTP.Disabled = types.BoolPointerValue(fromItem.Default.Local.HTTP.Disabled)
						if fromItem.Default.Local.HTTP.OnRateLimit == nil {
							from1.Default.Local.HTTP.OnRateLimit = nil
						} else {
							from1.Default.Local.HTTP.OnRateLimit = &tfTypes.OnRateLimit{}
							if fromItem.Default.Local.HTTP.OnRateLimit.Headers == nil {
								from1.Default.Local.HTTP.OnRateLimit.Headers = nil
							} else {
								from1.Default.Local.HTTP.OnRateLimit.Headers = &tfTypes.MeshGlobalRateLimitItemHeaders{}
								from1.Default.Local.HTTP.OnRateLimit.Headers.Add = []tfTypes.MeshGlobalRateLimitItemAdd{}
								for addCount, addItem := range fromItem.Default.Local.HTTP.OnRateLimit.Headers.Add {
									var add1 tfTypes.MeshGlobalRateLimitItemAdd
									add1.Name = types.StringValue(addItem.Name)
									add1.Value = types.StringValue(addItem.Value)
									if addCount+1 > len(from1.Default.Local.HTTP.OnRateLimit.Headers.Add) {
										from1.Default.Local.HTTP.OnRateLimit.Headers.Add = append(from1.Default.Local.HTTP.OnRateLimit.Headers.Add, add1)
									} else {
										from1.Default.Local.HTTP.OnRateLimit.Headers.Add[addCount].Name = add1.Name
										from1.Default.Local.HTTP.OnRateLimit.Headers.Add[addCount].Value = add1.Value
									}
								}
								from1.Default.Local.HTTP.OnRateLimit.Headers.Set = []tfTypes.MeshGlobalRateLimitItemAdd{}
								for setCount, setItem := range fromItem.Default.Local.HTTP.OnRateLimit.Headers.Set {
									var set1 tfTypes.MeshGlobalRateLimitItemAdd
									set1.Name = types.StringValue(setItem.Name)
									set1.Value = types.StringValue(setItem.Value)
									if setCount+1 > len(from1.Default.Local.HTTP.OnRateLimit.Headers.Set) {
										from1.Default.Local.HTTP.OnRateLimit.Headers.Set = append(from1.Default.Local.HTTP.OnRateLimit.Headers.Set, set1)
									} else {
										from1.Default.Local.HTTP.OnRateLimit.Headers.Set[setCount].Name = set1.Name
										from1.Default.Local.HTTP.OnRateLimit.Headers.Set[setCount].Value = set1.Value
									}
								}
							}
							if fromItem.Default.Local.HTTP.OnRateLimit.Status != nil {
								from1.Default.Local.HTTP.OnRateLimit.Status = types.Int32Value(int32(*fromItem.Default.Local.HTTP.OnRateLimit.Status))
							} else {
								from1.Default.Local.HTTP.OnRateLimit.Status = types.Int32Null()
							}
						}
						if fromItem.Default.Local.HTTP.RequestRate == nil {
							from1.Default.Local.HTTP.RequestRate = nil
						} else {
							from1.Default.Local.HTTP.RequestRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
							from1.Default.Local.HTTP.RequestRate.Interval = types.StringValue(fromItem.Default.Local.HTTP.RequestRate.Interval)
							from1.Default.Local.HTTP.RequestRate.Num = types.Int32Value(int32(fromItem.Default.Local.HTTP.RequestRate.Num))
						}
					}
					if fromItem.Default.Local.TCP == nil {
						from1.Default.Local.TCP = nil
					} else {
						from1.Default.Local.TCP = &tfTypes.MeshRateLimitItemTCP{}
						if fromItem.Default.Local.TCP.ConnectionRate == nil {
							from1.Default.Local.TCP.ConnectionRate = nil
						} else {
							from1.Default.Local.TCP.ConnectionRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
							from1.Default.Local.TCP.ConnectionRate.Interval = types.StringValue(fromItem.Default.Local.TCP.ConnectionRate.Interval)
							from1.Default.Local.TCP.ConnectionRate.Num = types.Int32Value(int32(fromItem.Default.Local.TCP.ConnectionRate.Num))
						}
						from1.Default.Local.TCP.Disabled = types.BoolPointerValue(fromItem.Default.Local.TCP.Disabled)
					}
				}
			}
			from1.TargetRef.Kind = types.StringValue(string(fromItem.TargetRef.Kind))
			if len(fromItem.TargetRef.Labels) > 0 {
				from1.TargetRef.Labels = make(map[string]types.String, len(fromItem.TargetRef.Labels))
				for key1, value3 := range fromItem.TargetRef.Labels {
					from1.TargetRef.Labels[key1] = types.StringValue(value3)
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
				for key2, value4 := range fromItem.TargetRef.Tags {
					from1.TargetRef.Tags[key2] = types.StringValue(value4)
				}
			}
			if fromCount+1 > len(r.Spec.From) {
				r.Spec.From = append(r.Spec.From, from1)
			} else {
				r.Spec.From[fromCount].Default = from1.Default
				r.Spec.From[fromCount].TargetRef = from1.TargetRef
			}
		}
		r.Spec.Rules = []tfTypes.MeshRateLimitItemRules{}
		if len(r.Spec.Rules) > len(resp.Spec.Rules) {
			r.Spec.Rules = r.Spec.Rules[:len(resp.Spec.Rules)]
		}
		for rulesCount, rulesItem := range resp.Spec.Rules {
			var rules1 tfTypes.MeshRateLimitItemRules
			if rulesItem.Default == nil {
				rules1.Default = nil
			} else {
				rules1.Default = &tfTypes.MeshRateLimitItemDefault{}
				if rulesItem.Default.Local == nil {
					rules1.Default.Local = nil
				} else {
					rules1.Default.Local = &tfTypes.Local{}
					if rulesItem.Default.Local.HTTP == nil {
						rules1.Default.Local.HTTP = nil
					} else {
						rules1.Default.Local.HTTP = &tfTypes.MeshRateLimitItemHTTP{}
						rules1.Default.Local.HTTP.Disabled = types.BoolPointerValue(rulesItem.Default.Local.HTTP.Disabled)
						if rulesItem.Default.Local.HTTP.OnRateLimit == nil {
							rules1.Default.Local.HTTP.OnRateLimit = nil
						} else {
							rules1.Default.Local.HTTP.OnRateLimit = &tfTypes.OnRateLimit{}
							if rulesItem.Default.Local.HTTP.OnRateLimit.Headers == nil {
								rules1.Default.Local.HTTP.OnRateLimit.Headers = nil
							} else {
								rules1.Default.Local.HTTP.OnRateLimit.Headers = &tfTypes.MeshGlobalRateLimitItemHeaders{}
								rules1.Default.Local.HTTP.OnRateLimit.Headers.Add = []tfTypes.MeshGlobalRateLimitItemAdd{}
								for addCount1, addItem1 := range rulesItem.Default.Local.HTTP.OnRateLimit.Headers.Add {
									var add3 tfTypes.MeshGlobalRateLimitItemAdd
									add3.Name = types.StringValue(addItem1.Name)
									add3.Value = types.StringValue(addItem1.Value)
									if addCount1+1 > len(rules1.Default.Local.HTTP.OnRateLimit.Headers.Add) {
										rules1.Default.Local.HTTP.OnRateLimit.Headers.Add = append(rules1.Default.Local.HTTP.OnRateLimit.Headers.Add, add3)
									} else {
										rules1.Default.Local.HTTP.OnRateLimit.Headers.Add[addCount1].Name = add3.Name
										rules1.Default.Local.HTTP.OnRateLimit.Headers.Add[addCount1].Value = add3.Value
									}
								}
								rules1.Default.Local.HTTP.OnRateLimit.Headers.Set = []tfTypes.MeshGlobalRateLimitItemAdd{}
								for setCount1, setItem1 := range rulesItem.Default.Local.HTTP.OnRateLimit.Headers.Set {
									var set3 tfTypes.MeshGlobalRateLimitItemAdd
									set3.Name = types.StringValue(setItem1.Name)
									set3.Value = types.StringValue(setItem1.Value)
									if setCount1+1 > len(rules1.Default.Local.HTTP.OnRateLimit.Headers.Set) {
										rules1.Default.Local.HTTP.OnRateLimit.Headers.Set = append(rules1.Default.Local.HTTP.OnRateLimit.Headers.Set, set3)
									} else {
										rules1.Default.Local.HTTP.OnRateLimit.Headers.Set[setCount1].Name = set3.Name
										rules1.Default.Local.HTTP.OnRateLimit.Headers.Set[setCount1].Value = set3.Value
									}
								}
							}
							if rulesItem.Default.Local.HTTP.OnRateLimit.Status != nil {
								rules1.Default.Local.HTTP.OnRateLimit.Status = types.Int32Value(int32(*rulesItem.Default.Local.HTTP.OnRateLimit.Status))
							} else {
								rules1.Default.Local.HTTP.OnRateLimit.Status = types.Int32Null()
							}
						}
						if rulesItem.Default.Local.HTTP.RequestRate == nil {
							rules1.Default.Local.HTTP.RequestRate = nil
						} else {
							rules1.Default.Local.HTTP.RequestRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
							rules1.Default.Local.HTTP.RequestRate.Interval = types.StringValue(rulesItem.Default.Local.HTTP.RequestRate.Interval)
							rules1.Default.Local.HTTP.RequestRate.Num = types.Int32Value(int32(rulesItem.Default.Local.HTTP.RequestRate.Num))
						}
					}
					if rulesItem.Default.Local.TCP == nil {
						rules1.Default.Local.TCP = nil
					} else {
						rules1.Default.Local.TCP = &tfTypes.MeshRateLimitItemTCP{}
						if rulesItem.Default.Local.TCP.ConnectionRate == nil {
							rules1.Default.Local.TCP.ConnectionRate = nil
						} else {
							rules1.Default.Local.TCP.ConnectionRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
							rules1.Default.Local.TCP.ConnectionRate.Interval = types.StringValue(rulesItem.Default.Local.TCP.ConnectionRate.Interval)
							rules1.Default.Local.TCP.ConnectionRate.Num = types.Int32Value(int32(rulesItem.Default.Local.TCP.ConnectionRate.Num))
						}
						rules1.Default.Local.TCP.Disabled = types.BoolPointerValue(rulesItem.Default.Local.TCP.Disabled)
					}
				}
			}
			if rulesCount+1 > len(r.Spec.Rules) {
				r.Spec.Rules = append(r.Spec.Rules, rules1)
			} else {
				r.Spec.Rules[rulesCount].Default = rules1.Default
			}
		}
		if resp.Spec.TargetRef == nil {
			r.Spec.TargetRef = nil
		} else {
			r.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
			r.Spec.TargetRef.Kind = types.StringValue(string(resp.Spec.TargetRef.Kind))
			if len(resp.Spec.TargetRef.Labels) > 0 {
				r.Spec.TargetRef.Labels = make(map[string]types.String, len(resp.Spec.TargetRef.Labels))
				for key3, value7 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key3] = types.StringValue(value7)
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
				for key4, value8 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key4] = types.StringValue(value8)
				}
			}
		}
		r.Spec.To = []tfTypes.MeshRateLimitItemFrom{}
		if len(r.Spec.To) > len(resp.Spec.To) {
			r.Spec.To = r.Spec.To[:len(resp.Spec.To)]
		}
		for toCount, toItem := range resp.Spec.To {
			var to1 tfTypes.MeshRateLimitItemFrom
			if toItem.Default == nil {
				to1.Default = nil
			} else {
				to1.Default = &tfTypes.MeshRateLimitItemDefault{}
				if toItem.Default.Local == nil {
					to1.Default.Local = nil
				} else {
					to1.Default.Local = &tfTypes.Local{}
					if toItem.Default.Local.HTTP == nil {
						to1.Default.Local.HTTP = nil
					} else {
						to1.Default.Local.HTTP = &tfTypes.MeshRateLimitItemHTTP{}
						to1.Default.Local.HTTP.Disabled = types.BoolPointerValue(toItem.Default.Local.HTTP.Disabled)
						if toItem.Default.Local.HTTP.OnRateLimit == nil {
							to1.Default.Local.HTTP.OnRateLimit = nil
						} else {
							to1.Default.Local.HTTP.OnRateLimit = &tfTypes.OnRateLimit{}
							if toItem.Default.Local.HTTP.OnRateLimit.Headers == nil {
								to1.Default.Local.HTTP.OnRateLimit.Headers = nil
							} else {
								to1.Default.Local.HTTP.OnRateLimit.Headers = &tfTypes.MeshGlobalRateLimitItemHeaders{}
								to1.Default.Local.HTTP.OnRateLimit.Headers.Add = []tfTypes.MeshGlobalRateLimitItemAdd{}
								for addCount2, addItem2 := range toItem.Default.Local.HTTP.OnRateLimit.Headers.Add {
									var add5 tfTypes.MeshGlobalRateLimitItemAdd
									add5.Name = types.StringValue(addItem2.Name)
									add5.Value = types.StringValue(addItem2.Value)
									if addCount2+1 > len(to1.Default.Local.HTTP.OnRateLimit.Headers.Add) {
										to1.Default.Local.HTTP.OnRateLimit.Headers.Add = append(to1.Default.Local.HTTP.OnRateLimit.Headers.Add, add5)
									} else {
										to1.Default.Local.HTTP.OnRateLimit.Headers.Add[addCount2].Name = add5.Name
										to1.Default.Local.HTTP.OnRateLimit.Headers.Add[addCount2].Value = add5.Value
									}
								}
								to1.Default.Local.HTTP.OnRateLimit.Headers.Set = []tfTypes.MeshGlobalRateLimitItemAdd{}
								for setCount2, setItem2 := range toItem.Default.Local.HTTP.OnRateLimit.Headers.Set {
									var set5 tfTypes.MeshGlobalRateLimitItemAdd
									set5.Name = types.StringValue(setItem2.Name)
									set5.Value = types.StringValue(setItem2.Value)
									if setCount2+1 > len(to1.Default.Local.HTTP.OnRateLimit.Headers.Set) {
										to1.Default.Local.HTTP.OnRateLimit.Headers.Set = append(to1.Default.Local.HTTP.OnRateLimit.Headers.Set, set5)
									} else {
										to1.Default.Local.HTTP.OnRateLimit.Headers.Set[setCount2].Name = set5.Name
										to1.Default.Local.HTTP.OnRateLimit.Headers.Set[setCount2].Value = set5.Value
									}
								}
							}
							if toItem.Default.Local.HTTP.OnRateLimit.Status != nil {
								to1.Default.Local.HTTP.OnRateLimit.Status = types.Int32Value(int32(*toItem.Default.Local.HTTP.OnRateLimit.Status))
							} else {
								to1.Default.Local.HTTP.OnRateLimit.Status = types.Int32Null()
							}
						}
						if toItem.Default.Local.HTTP.RequestRate == nil {
							to1.Default.Local.HTTP.RequestRate = nil
						} else {
							to1.Default.Local.HTTP.RequestRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
							to1.Default.Local.HTTP.RequestRate.Interval = types.StringValue(toItem.Default.Local.HTTP.RequestRate.Interval)
							to1.Default.Local.HTTP.RequestRate.Num = types.Int32Value(int32(toItem.Default.Local.HTTP.RequestRate.Num))
						}
					}
					if toItem.Default.Local.TCP == nil {
						to1.Default.Local.TCP = nil
					} else {
						to1.Default.Local.TCP = &tfTypes.MeshRateLimitItemTCP{}
						if toItem.Default.Local.TCP.ConnectionRate == nil {
							to1.Default.Local.TCP.ConnectionRate = nil
						} else {
							to1.Default.Local.TCP.ConnectionRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
							to1.Default.Local.TCP.ConnectionRate.Interval = types.StringValue(toItem.Default.Local.TCP.ConnectionRate.Interval)
							to1.Default.Local.TCP.ConnectionRate.Num = types.Int32Value(int32(toItem.Default.Local.TCP.ConnectionRate.Num))
						}
						to1.Default.Local.TCP.Disabled = types.BoolPointerValue(toItem.Default.Local.TCP.Disabled)
					}
				}
			}
			to1.TargetRef.Kind = types.StringValue(string(toItem.TargetRef.Kind))
			if len(toItem.TargetRef.Labels) > 0 {
				to1.TargetRef.Labels = make(map[string]types.String, len(toItem.TargetRef.Labels))
				for key5, value11 := range toItem.TargetRef.Labels {
					to1.TargetRef.Labels[key5] = types.StringValue(value11)
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
				for key6, value12 := range toItem.TargetRef.Tags {
					to1.TargetRef.Tags[key6] = types.StringValue(value12)
				}
			}
			if toCount+1 > len(r.Spec.To) {
				r.Spec.To = append(r.Spec.To, to1)
			} else {
				r.Spec.To[toCount].Default = to1.Default
				r.Spec.To[toCount].TargetRef = to1.TargetRef
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}
}
