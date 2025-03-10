// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect-beta/internal/provider/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"time"
)

func (r *MeshHealthCheckResourceModel) ToSharedMeshHealthCheckItemInput() *shared.MeshHealthCheckItemInput {
	typeVar := shared.MeshHealthCheckItemType(r.Type.ValueString())
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
	var targetRef *shared.MeshHealthCheckItemTargetRef
	if r.Spec.TargetRef != nil {
		kind := shared.MeshHealthCheckItemKind(r.Spec.TargetRef.Kind.ValueString())
		labels1 := make(map[string]string)
		for labelsKey1, labelsValue1 := range r.Spec.TargetRef.Labels {
			var labelsInst1 string
			labelsInst1 = labelsValue1.ValueString()

			labels1[labelsKey1] = labelsInst1
		}
		mesh1 := new(string)
		if !r.Spec.TargetRef.Mesh.IsUnknown() && !r.Spec.TargetRef.Mesh.IsNull() {
			*mesh1 = r.Spec.TargetRef.Mesh.ValueString()
		} else {
			mesh1 = nil
		}
		name1 := new(string)
		if !r.Spec.TargetRef.Name.IsUnknown() && !r.Spec.TargetRef.Name.IsNull() {
			*name1 = r.Spec.TargetRef.Name.ValueString()
		} else {
			name1 = nil
		}
		namespace := new(string)
		if !r.Spec.TargetRef.Namespace.IsUnknown() && !r.Spec.TargetRef.Namespace.IsNull() {
			*namespace = r.Spec.TargetRef.Namespace.ValueString()
		} else {
			namespace = nil
		}
		var proxyTypes []shared.MeshHealthCheckItemProxyTypes = []shared.MeshHealthCheckItemProxyTypes{}
		for _, proxyTypesItem := range r.Spec.TargetRef.ProxyTypes {
			proxyTypes = append(proxyTypes, shared.MeshHealthCheckItemProxyTypes(proxyTypesItem.ValueString()))
		}
		sectionName := new(string)
		if !r.Spec.TargetRef.SectionName.IsUnknown() && !r.Spec.TargetRef.SectionName.IsNull() {
			*sectionName = r.Spec.TargetRef.SectionName.ValueString()
		} else {
			sectionName = nil
		}
		tags := make(map[string]string)
		for tagsKey, tagsValue := range r.Spec.TargetRef.Tags {
			var tagsInst string
			tagsInst = tagsValue.ValueString()

			tags[tagsKey] = tagsInst
		}
		targetRef = &shared.MeshHealthCheckItemTargetRef{
			Kind:        kind,
			Labels:      labels1,
			Mesh:        mesh1,
			Name:        name1,
			Namespace:   namespace,
			ProxyTypes:  proxyTypes,
			SectionName: sectionName,
			Tags:        tags,
		}
	}
	var to []shared.MeshHealthCheckItemTo = []shared.MeshHealthCheckItemTo{}
	for _, toItem := range r.Spec.To {
		var defaultVar *shared.MeshHealthCheckItemDefault
		if toItem.Default != nil {
			alwaysLogHealthCheckFailures := new(bool)
			if !toItem.Default.AlwaysLogHealthCheckFailures.IsUnknown() && !toItem.Default.AlwaysLogHealthCheckFailures.IsNull() {
				*alwaysLogHealthCheckFailures = toItem.Default.AlwaysLogHealthCheckFailures.ValueBool()
			} else {
				alwaysLogHealthCheckFailures = nil
			}
			eventLogPath := new(string)
			if !toItem.Default.EventLogPath.IsUnknown() && !toItem.Default.EventLogPath.IsNull() {
				*eventLogPath = toItem.Default.EventLogPath.ValueString()
			} else {
				eventLogPath = nil
			}
			failTrafficOnPanic := new(bool)
			if !toItem.Default.FailTrafficOnPanic.IsUnknown() && !toItem.Default.FailTrafficOnPanic.IsNull() {
				*failTrafficOnPanic = toItem.Default.FailTrafficOnPanic.ValueBool()
			} else {
				failTrafficOnPanic = nil
			}
			var grpc *shared.Grpc
			if toItem.Default.Grpc != nil {
				authority := new(string)
				if !toItem.Default.Grpc.Authority.IsUnknown() && !toItem.Default.Grpc.Authority.IsNull() {
					*authority = toItem.Default.Grpc.Authority.ValueString()
				} else {
					authority = nil
				}
				disabled := new(bool)
				if !toItem.Default.Grpc.Disabled.IsUnknown() && !toItem.Default.Grpc.Disabled.IsNull() {
					*disabled = toItem.Default.Grpc.Disabled.ValueBool()
				} else {
					disabled = nil
				}
				serviceName := new(string)
				if !toItem.Default.Grpc.ServiceName.IsUnknown() && !toItem.Default.Grpc.ServiceName.IsNull() {
					*serviceName = toItem.Default.Grpc.ServiceName.ValueString()
				} else {
					serviceName = nil
				}
				grpc = &shared.Grpc{
					Authority:   authority,
					Disabled:    disabled,
					ServiceName: serviceName,
				}
			}
			var healthyPanicThreshold *shared.HealthyPanicThreshold
			if toItem.Default.HealthyPanicThreshold != nil {
				integer := new(int64)
				if !toItem.Default.HealthyPanicThreshold.Integer.IsUnknown() && !toItem.Default.HealthyPanicThreshold.Integer.IsNull() {
					*integer = toItem.Default.HealthyPanicThreshold.Integer.ValueInt64()
				} else {
					integer = nil
				}
				if integer != nil {
					healthyPanicThreshold = &shared.HealthyPanicThreshold{
						Integer: integer,
					}
				}
				str := new(string)
				if !toItem.Default.HealthyPanicThreshold.Str.IsUnknown() && !toItem.Default.HealthyPanicThreshold.Str.IsNull() {
					*str = toItem.Default.HealthyPanicThreshold.Str.ValueString()
				} else {
					str = nil
				}
				if str != nil {
					healthyPanicThreshold = &shared.HealthyPanicThreshold{
						Str: str,
					}
				}
			}
			healthyThreshold := new(int)
			if !toItem.Default.HealthyThreshold.IsUnknown() && !toItem.Default.HealthyThreshold.IsNull() {
				*healthyThreshold = int(toItem.Default.HealthyThreshold.ValueInt32())
			} else {
				healthyThreshold = nil
			}
			var http *shared.MeshHealthCheckItemHTTP
			if toItem.Default.HTTP != nil {
				disabled1 := new(bool)
				if !toItem.Default.HTTP.Disabled.IsUnknown() && !toItem.Default.HTTP.Disabled.IsNull() {
					*disabled1 = toItem.Default.HTTP.Disabled.ValueBool()
				} else {
					disabled1 = nil
				}
				var expectedStatuses []int64 = []int64{}
				for _, expectedStatusesItem := range toItem.Default.HTTP.ExpectedStatuses {
					expectedStatuses = append(expectedStatuses, expectedStatusesItem.ValueInt64())
				}
				path := new(string)
				if !toItem.Default.HTTP.Path.IsUnknown() && !toItem.Default.HTTP.Path.IsNull() {
					*path = toItem.Default.HTTP.Path.ValueString()
				} else {
					path = nil
				}
				var requestHeadersToAdd *shared.RequestHeadersToAdd
				if toItem.Default.HTTP.RequestHeadersToAdd != nil {
					var add []shared.Add = []shared.Add{}
					for _, addItem := range toItem.Default.HTTP.RequestHeadersToAdd.Add {
						var name2 string
						name2 = addItem.Name.ValueString()

						var value string
						value = addItem.Value.ValueString()

						add = append(add, shared.Add{
							Name:  name2,
							Value: value,
						})
					}
					var set []shared.Set = []shared.Set{}
					for _, setItem := range toItem.Default.HTTP.RequestHeadersToAdd.Set {
						var name3 string
						name3 = setItem.Name.ValueString()

						var value1 string
						value1 = setItem.Value.ValueString()

						set = append(set, shared.Set{
							Name:  name3,
							Value: value1,
						})
					}
					requestHeadersToAdd = &shared.RequestHeadersToAdd{
						Add: add,
						Set: set,
					}
				}
				http = &shared.MeshHealthCheckItemHTTP{
					Disabled:            disabled1,
					ExpectedStatuses:    expectedStatuses,
					Path:                path,
					RequestHeadersToAdd: requestHeadersToAdd,
				}
			}
			initialJitter := new(string)
			if !toItem.Default.InitialJitter.IsUnknown() && !toItem.Default.InitialJitter.IsNull() {
				*initialJitter = toItem.Default.InitialJitter.ValueString()
			} else {
				initialJitter = nil
			}
			interval := new(string)
			if !toItem.Default.Interval.IsUnknown() && !toItem.Default.Interval.IsNull() {
				*interval = toItem.Default.Interval.ValueString()
			} else {
				interval = nil
			}
			intervalJitter := new(string)
			if !toItem.Default.IntervalJitter.IsUnknown() && !toItem.Default.IntervalJitter.IsNull() {
				*intervalJitter = toItem.Default.IntervalJitter.ValueString()
			} else {
				intervalJitter = nil
			}
			intervalJitterPercent := new(int)
			if !toItem.Default.IntervalJitterPercent.IsUnknown() && !toItem.Default.IntervalJitterPercent.IsNull() {
				*intervalJitterPercent = int(toItem.Default.IntervalJitterPercent.ValueInt32())
			} else {
				intervalJitterPercent = nil
			}
			noTrafficInterval := new(string)
			if !toItem.Default.NoTrafficInterval.IsUnknown() && !toItem.Default.NoTrafficInterval.IsNull() {
				*noTrafficInterval = toItem.Default.NoTrafficInterval.ValueString()
			} else {
				noTrafficInterval = nil
			}
			reuseConnection := new(bool)
			if !toItem.Default.ReuseConnection.IsUnknown() && !toItem.Default.ReuseConnection.IsNull() {
				*reuseConnection = toItem.Default.ReuseConnection.ValueBool()
			} else {
				reuseConnection = nil
			}
			var tcp *shared.TCP
			if toItem.Default.TCP != nil {
				disabled2 := new(bool)
				if !toItem.Default.TCP.Disabled.IsUnknown() && !toItem.Default.TCP.Disabled.IsNull() {
					*disabled2 = toItem.Default.TCP.Disabled.ValueBool()
				} else {
					disabled2 = nil
				}
				var receive []string = []string{}
				for _, receiveItem := range toItem.Default.TCP.Receive {
					receive = append(receive, receiveItem.ValueString())
				}
				send := new(string)
				if !toItem.Default.TCP.Send.IsUnknown() && !toItem.Default.TCP.Send.IsNull() {
					*send = toItem.Default.TCP.Send.ValueString()
				} else {
					send = nil
				}
				tcp = &shared.TCP{
					Disabled: disabled2,
					Receive:  receive,
					Send:     send,
				}
			}
			timeout := new(string)
			if !toItem.Default.Timeout.IsUnknown() && !toItem.Default.Timeout.IsNull() {
				*timeout = toItem.Default.Timeout.ValueString()
			} else {
				timeout = nil
			}
			unhealthyThreshold := new(int)
			if !toItem.Default.UnhealthyThreshold.IsUnknown() && !toItem.Default.UnhealthyThreshold.IsNull() {
				*unhealthyThreshold = int(toItem.Default.UnhealthyThreshold.ValueInt32())
			} else {
				unhealthyThreshold = nil
			}
			defaultVar = &shared.MeshHealthCheckItemDefault{
				AlwaysLogHealthCheckFailures: alwaysLogHealthCheckFailures,
				EventLogPath:                 eventLogPath,
				FailTrafficOnPanic:           failTrafficOnPanic,
				Grpc:                         grpc,
				HealthyPanicThreshold:        healthyPanicThreshold,
				HealthyThreshold:             healthyThreshold,
				HTTP:                         http,
				InitialJitter:                initialJitter,
				Interval:                     interval,
				IntervalJitter:               intervalJitter,
				IntervalJitterPercent:        intervalJitterPercent,
				NoTrafficInterval:            noTrafficInterval,
				ReuseConnection:              reuseConnection,
				TCP:                          tcp,
				Timeout:                      timeout,
				UnhealthyThreshold:           unhealthyThreshold,
			}
		}
		kind1 := shared.MeshHealthCheckItemSpecKind(toItem.TargetRef.Kind.ValueString())
		labels2 := make(map[string]string)
		for labelsKey2, labelsValue2 := range toItem.TargetRef.Labels {
			var labelsInst2 string
			labelsInst2 = labelsValue2.ValueString()

			labels2[labelsKey2] = labelsInst2
		}
		mesh2 := new(string)
		if !toItem.TargetRef.Mesh.IsUnknown() && !toItem.TargetRef.Mesh.IsNull() {
			*mesh2 = toItem.TargetRef.Mesh.ValueString()
		} else {
			mesh2 = nil
		}
		name4 := new(string)
		if !toItem.TargetRef.Name.IsUnknown() && !toItem.TargetRef.Name.IsNull() {
			*name4 = toItem.TargetRef.Name.ValueString()
		} else {
			name4 = nil
		}
		namespace1 := new(string)
		if !toItem.TargetRef.Namespace.IsUnknown() && !toItem.TargetRef.Namespace.IsNull() {
			*namespace1 = toItem.TargetRef.Namespace.ValueString()
		} else {
			namespace1 = nil
		}
		var proxyTypes1 []shared.MeshHealthCheckItemSpecProxyTypes = []shared.MeshHealthCheckItemSpecProxyTypes{}
		for _, proxyTypesItem1 := range toItem.TargetRef.ProxyTypes {
			proxyTypes1 = append(proxyTypes1, shared.MeshHealthCheckItemSpecProxyTypes(proxyTypesItem1.ValueString()))
		}
		sectionName1 := new(string)
		if !toItem.TargetRef.SectionName.IsUnknown() && !toItem.TargetRef.SectionName.IsNull() {
			*sectionName1 = toItem.TargetRef.SectionName.ValueString()
		} else {
			sectionName1 = nil
		}
		tags1 := make(map[string]string)
		for tagsKey1, tagsValue1 := range toItem.TargetRef.Tags {
			var tagsInst1 string
			tagsInst1 = tagsValue1.ValueString()

			tags1[tagsKey1] = tagsInst1
		}
		targetRef1 := shared.MeshHealthCheckItemSpecTargetRef{
			Kind:        kind1,
			Labels:      labels2,
			Mesh:        mesh2,
			Name:        name4,
			Namespace:   namespace1,
			ProxyTypes:  proxyTypes1,
			SectionName: sectionName1,
			Tags:        tags1,
		}
		to = append(to, shared.MeshHealthCheckItemTo{
			Default:   defaultVar,
			TargetRef: targetRef1,
		})
	}
	spec := shared.MeshHealthCheckItemSpec{
		TargetRef: targetRef,
		To:        to,
	}
	out := shared.MeshHealthCheckItemInput{
		Type:   typeVar,
		Mesh:   mesh,
		Name:   name,
		Labels: labels,
		Spec:   spec,
	}
	return &out
}

func (r *MeshHealthCheckResourceModel) RefreshFromSharedMeshHealthCheckCreateOrUpdateSuccessResponse(resp *shared.MeshHealthCheckCreateOrUpdateSuccessResponse) {
	if resp != nil {
		r.Warnings = make([]types.String, 0, len(resp.Warnings))
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}
}

func (r *MeshHealthCheckResourceModel) RefreshFromSharedMeshHealthCheckItem(resp *shared.MeshHealthCheckItem) {
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
		if resp.Spec.TargetRef == nil {
			r.Spec.TargetRef = nil
		} else {
			r.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
			r.Spec.TargetRef.Kind = types.StringValue(string(resp.Spec.TargetRef.Kind))
			if len(resp.Spec.TargetRef.Labels) > 0 {
				r.Spec.TargetRef.Labels = make(map[string]types.String, len(resp.Spec.TargetRef.Labels))
				for key1, value1 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key1] = types.StringValue(value1)
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
				for key2, value2 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key2] = types.StringValue(value2)
				}
			}
		}
		r.Spec.To = []tfTypes.MeshHealthCheckItemTo{}
		if len(r.Spec.To) > len(resp.Spec.To) {
			r.Spec.To = r.Spec.To[:len(resp.Spec.To)]
		}
		for toCount, toItem := range resp.Spec.To {
			var to1 tfTypes.MeshHealthCheckItemTo
			if toItem.Default == nil {
				to1.Default = nil
			} else {
				to1.Default = &tfTypes.MeshHealthCheckItemDefault{}
				to1.Default.AlwaysLogHealthCheckFailures = types.BoolPointerValue(toItem.Default.AlwaysLogHealthCheckFailures)
				to1.Default.EventLogPath = types.StringPointerValue(toItem.Default.EventLogPath)
				to1.Default.FailTrafficOnPanic = types.BoolPointerValue(toItem.Default.FailTrafficOnPanic)
				if toItem.Default.Grpc == nil {
					to1.Default.Grpc = nil
				} else {
					to1.Default.Grpc = &tfTypes.Grpc{}
					to1.Default.Grpc.Authority = types.StringPointerValue(toItem.Default.Grpc.Authority)
					to1.Default.Grpc.Disabled = types.BoolPointerValue(toItem.Default.Grpc.Disabled)
					to1.Default.Grpc.ServiceName = types.StringPointerValue(toItem.Default.Grpc.ServiceName)
				}
				if toItem.Default.HealthyPanicThreshold == nil {
					to1.Default.HealthyPanicThreshold = nil
				} else {
					to1.Default.HealthyPanicThreshold = &tfTypes.MeshItemMode{}
					if toItem.Default.HealthyPanicThreshold.Integer != nil {
						to1.Default.HealthyPanicThreshold.Integer = types.Int64PointerValue(toItem.Default.HealthyPanicThreshold.Integer)
					}
					if toItem.Default.HealthyPanicThreshold.Str != nil {
						to1.Default.HealthyPanicThreshold.Str = types.StringPointerValue(toItem.Default.HealthyPanicThreshold.Str)
					}
				}
				if toItem.Default.HealthyThreshold != nil {
					to1.Default.HealthyThreshold = types.Int32Value(int32(*toItem.Default.HealthyThreshold))
				} else {
					to1.Default.HealthyThreshold = types.Int32Null()
				}
				if toItem.Default.HTTP == nil {
					to1.Default.HTTP = nil
				} else {
					to1.Default.HTTP = &tfTypes.MeshHealthCheckItemHTTP{}
					to1.Default.HTTP.Disabled = types.BoolPointerValue(toItem.Default.HTTP.Disabled)
					to1.Default.HTTP.ExpectedStatuses = make([]types.Int64, 0, len(toItem.Default.HTTP.ExpectedStatuses))
					for _, v := range toItem.Default.HTTP.ExpectedStatuses {
						to1.Default.HTTP.ExpectedStatuses = append(to1.Default.HTTP.ExpectedStatuses, types.Int64Value(v))
					}
					to1.Default.HTTP.Path = types.StringPointerValue(toItem.Default.HTTP.Path)
					if toItem.Default.HTTP.RequestHeadersToAdd == nil {
						to1.Default.HTTP.RequestHeadersToAdd = nil
					} else {
						to1.Default.HTTP.RequestHeadersToAdd = &tfTypes.MeshGlobalRateLimitItemHeaders{}
						to1.Default.HTTP.RequestHeadersToAdd.Add = []tfTypes.MeshGlobalRateLimitItemAdd{}
						for addCount, addItem := range toItem.Default.HTTP.RequestHeadersToAdd.Add {
							var add1 tfTypes.MeshGlobalRateLimitItemAdd
							add1.Name = types.StringValue(addItem.Name)
							add1.Value = types.StringValue(addItem.Value)
							if addCount+1 > len(to1.Default.HTTP.RequestHeadersToAdd.Add) {
								to1.Default.HTTP.RequestHeadersToAdd.Add = append(to1.Default.HTTP.RequestHeadersToAdd.Add, add1)
							} else {
								to1.Default.HTTP.RequestHeadersToAdd.Add[addCount].Name = add1.Name
								to1.Default.HTTP.RequestHeadersToAdd.Add[addCount].Value = add1.Value
							}
						}
						to1.Default.HTTP.RequestHeadersToAdd.Set = []tfTypes.MeshGlobalRateLimitItemAdd{}
						for setCount, setItem := range toItem.Default.HTTP.RequestHeadersToAdd.Set {
							var set1 tfTypes.MeshGlobalRateLimitItemAdd
							set1.Name = types.StringValue(setItem.Name)
							set1.Value = types.StringValue(setItem.Value)
							if setCount+1 > len(to1.Default.HTTP.RequestHeadersToAdd.Set) {
								to1.Default.HTTP.RequestHeadersToAdd.Set = append(to1.Default.HTTP.RequestHeadersToAdd.Set, set1)
							} else {
								to1.Default.HTTP.RequestHeadersToAdd.Set[setCount].Name = set1.Name
								to1.Default.HTTP.RequestHeadersToAdd.Set[setCount].Value = set1.Value
							}
						}
					}
				}
				to1.Default.InitialJitter = types.StringPointerValue(toItem.Default.InitialJitter)
				to1.Default.Interval = types.StringPointerValue(toItem.Default.Interval)
				to1.Default.IntervalJitter = types.StringPointerValue(toItem.Default.IntervalJitter)
				if toItem.Default.IntervalJitterPercent != nil {
					to1.Default.IntervalJitterPercent = types.Int32Value(int32(*toItem.Default.IntervalJitterPercent))
				} else {
					to1.Default.IntervalJitterPercent = types.Int32Null()
				}
				to1.Default.NoTrafficInterval = types.StringPointerValue(toItem.Default.NoTrafficInterval)
				to1.Default.ReuseConnection = types.BoolPointerValue(toItem.Default.ReuseConnection)
				if toItem.Default.TCP == nil {
					to1.Default.TCP = nil
				} else {
					to1.Default.TCP = &tfTypes.TCP{}
					to1.Default.TCP.Disabled = types.BoolPointerValue(toItem.Default.TCP.Disabled)
					to1.Default.TCP.Receive = make([]types.String, 0, len(toItem.Default.TCP.Receive))
					for _, v := range toItem.Default.TCP.Receive {
						to1.Default.TCP.Receive = append(to1.Default.TCP.Receive, types.StringValue(v))
					}
					to1.Default.TCP.Send = types.StringPointerValue(toItem.Default.TCP.Send)
				}
				to1.Default.Timeout = types.StringPointerValue(toItem.Default.Timeout)
				if toItem.Default.UnhealthyThreshold != nil {
					to1.Default.UnhealthyThreshold = types.Int32Value(int32(*toItem.Default.UnhealthyThreshold))
				} else {
					to1.Default.UnhealthyThreshold = types.Int32Null()
				}
			}
			to1.TargetRef.Kind = types.StringValue(string(toItem.TargetRef.Kind))
			if len(toItem.TargetRef.Labels) > 0 {
				to1.TargetRef.Labels = make(map[string]types.String, len(toItem.TargetRef.Labels))
				for key3, value5 := range toItem.TargetRef.Labels {
					to1.TargetRef.Labels[key3] = types.StringValue(value5)
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
				for key4, value6 := range toItem.TargetRef.Tags {
					to1.TargetRef.Tags[key4] = types.StringValue(value6)
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
