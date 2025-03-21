// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect-beta/internal/provider/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"time"
)

func (r *MeshMultiZoneServiceResourceModel) ToSharedMeshMultiZoneServiceItemInput() *shared.MeshMultiZoneServiceItemInput {
	typeVar := shared.MeshMultiZoneServiceItemType(r.Type.ValueString())
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
	var ports []shared.Ports = []shared.Ports{}
	for _, portsItem := range r.Spec.Ports {
		appProtocol := new(string)
		if !portsItem.AppProtocol.IsUnknown() && !portsItem.AppProtocol.IsNull() {
			*appProtocol = portsItem.AppProtocol.ValueString()
		} else {
			appProtocol = nil
		}
		name1 := new(string)
		if !portsItem.Name.IsUnknown() && !portsItem.Name.IsNull() {
			*name1 = portsItem.Name.ValueString()
		} else {
			name1 = nil
		}
		var port int
		port = int(portsItem.Port.ValueInt32())

		ports = append(ports, shared.Ports{
			AppProtocol: appProtocol,
			Name:        name1,
			Port:        port,
		})
	}
	matchLabels := make(map[string]string)
	for matchLabelsKey, matchLabelsValue := range r.Spec.Selector.MeshService.MatchLabels {
		var matchLabelsInst string
		matchLabelsInst = matchLabelsValue.ValueString()

		matchLabels[matchLabelsKey] = matchLabelsInst
	}
	meshService := shared.MeshMultiZoneServiceItemMeshService{
		MatchLabels: matchLabels,
	}
	selector := shared.MeshMultiZoneServiceItemSelector{
		MeshService: meshService,
	}
	spec := shared.MeshMultiZoneServiceItemSpec{
		Ports:    ports,
		Selector: selector,
	}
	out := shared.MeshMultiZoneServiceItemInput{
		Type:   typeVar,
		Mesh:   mesh,
		Name:   name,
		Labels: labels,
		Spec:   spec,
	}
	return &out
}

func (r *MeshMultiZoneServiceResourceModel) RefreshFromSharedMeshMultiZoneServiceCreateOrUpdateSuccessResponse(resp *shared.MeshMultiZoneServiceCreateOrUpdateSuccessResponse) {
	if resp != nil {
		r.Warnings = make([]types.String, 0, len(resp.Warnings))
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}
}

func (r *MeshMultiZoneServiceResourceModel) RefreshFromSharedMeshMultiZoneServiceItem(resp *shared.MeshMultiZoneServiceItem) {
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
		r.Spec.Ports = []tfTypes.Ports{}
		if len(r.Spec.Ports) > len(resp.Spec.Ports) {
			r.Spec.Ports = r.Spec.Ports[:len(resp.Spec.Ports)]
		}
		for portsCount, portsItem := range resp.Spec.Ports {
			var ports1 tfTypes.Ports
			ports1.AppProtocol = types.StringPointerValue(portsItem.AppProtocol)
			ports1.Name = types.StringPointerValue(portsItem.Name)
			ports1.Port = types.Int32Value(int32(portsItem.Port))
			if portsCount+1 > len(r.Spec.Ports) {
				r.Spec.Ports = append(r.Spec.Ports, ports1)
			} else {
				r.Spec.Ports[portsCount].AppProtocol = ports1.AppProtocol
				r.Spec.Ports[portsCount].Name = ports1.Name
				r.Spec.Ports[portsCount].Port = ports1.Port
			}
		}
		if len(resp.Spec.Selector.MeshService.MatchLabels) > 0 {
			r.Spec.Selector.MeshService.MatchLabels = make(map[string]types.String, len(resp.Spec.Selector.MeshService.MatchLabels))
			for key1, value1 := range resp.Spec.Selector.MeshService.MatchLabels {
				r.Spec.Selector.MeshService.MatchLabels[key1] = types.StringValue(value1)
			}
		}
		if resp.Status == nil {
			r.Status = nil
		} else {
			r.Status = &tfTypes.MeshMultiZoneServiceItemStatus{}
			r.Status.Addresses = []tfTypes.Addresses{}
			if len(r.Status.Addresses) > len(resp.Status.Addresses) {
				r.Status.Addresses = r.Status.Addresses[:len(resp.Status.Addresses)]
			}
			for addressesCount, addressesItem := range resp.Status.Addresses {
				var addresses1 tfTypes.Addresses
				addresses1.Hostname = types.StringPointerValue(addressesItem.Hostname)
				if addressesItem.HostnameGeneratorRef == nil {
					addresses1.HostnameGeneratorRef = nil
				} else {
					addresses1.HostnameGeneratorRef = &tfTypes.HostnameGeneratorRef{}
					addresses1.HostnameGeneratorRef.CoreName = types.StringValue(addressesItem.HostnameGeneratorRef.CoreName)
				}
				addresses1.Origin = types.StringPointerValue(addressesItem.Origin)
				if addressesCount+1 > len(r.Status.Addresses) {
					r.Status.Addresses = append(r.Status.Addresses, addresses1)
				} else {
					r.Status.Addresses[addressesCount].Hostname = addresses1.Hostname
					r.Status.Addresses[addressesCount].HostnameGeneratorRef = addresses1.HostnameGeneratorRef
					r.Status.Addresses[addressesCount].Origin = addresses1.Origin
				}
			}
			r.Status.HostnameGenerators = []tfTypes.HostnameGenerators{}
			if len(r.Status.HostnameGenerators) > len(resp.Status.HostnameGenerators) {
				r.Status.HostnameGenerators = r.Status.HostnameGenerators[:len(resp.Status.HostnameGenerators)]
			}
			for hostnameGeneratorsCount, hostnameGeneratorsItem := range resp.Status.HostnameGenerators {
				var hostnameGenerators1 tfTypes.HostnameGenerators
				hostnameGenerators1.Conditions = []tfTypes.Conditions{}
				for conditionsCount, conditionsItem := range hostnameGeneratorsItem.Conditions {
					var conditions1 tfTypes.Conditions
					conditions1.Message = types.StringValue(conditionsItem.Message)
					conditions1.Reason = types.StringValue(conditionsItem.Reason)
					conditions1.Status = types.StringValue(string(conditionsItem.Status))
					conditions1.Type = types.StringValue(conditionsItem.Type)
					if conditionsCount+1 > len(hostnameGenerators1.Conditions) {
						hostnameGenerators1.Conditions = append(hostnameGenerators1.Conditions, conditions1)
					} else {
						hostnameGenerators1.Conditions[conditionsCount].Message = conditions1.Message
						hostnameGenerators1.Conditions[conditionsCount].Reason = conditions1.Reason
						hostnameGenerators1.Conditions[conditionsCount].Status = conditions1.Status
						hostnameGenerators1.Conditions[conditionsCount].Type = conditions1.Type
					}
				}
				hostnameGenerators1.HostnameGeneratorRef.CoreName = types.StringValue(hostnameGeneratorsItem.HostnameGeneratorRef.CoreName)
				if hostnameGeneratorsCount+1 > len(r.Status.HostnameGenerators) {
					r.Status.HostnameGenerators = append(r.Status.HostnameGenerators, hostnameGenerators1)
				} else {
					r.Status.HostnameGenerators[hostnameGeneratorsCount].Conditions = hostnameGenerators1.Conditions
					r.Status.HostnameGenerators[hostnameGeneratorsCount].HostnameGeneratorRef = hostnameGenerators1.HostnameGeneratorRef
				}
			}
			r.Status.MeshServices = []tfTypes.MeshMultiZoneServiceItemMeshServices{}
			if len(r.Status.MeshServices) > len(resp.Status.MeshServices) {
				r.Status.MeshServices = r.Status.MeshServices[:len(resp.Status.MeshServices)]
			}
			for meshServicesCount, meshServicesItem := range resp.Status.MeshServices {
				var meshServices1 tfTypes.MeshMultiZoneServiceItemMeshServices
				meshServices1.Mesh = types.StringValue(meshServicesItem.Mesh)
				meshServices1.Name = types.StringValue(meshServicesItem.Name)
				meshServices1.Namespace = types.StringValue(meshServicesItem.Namespace)
				meshServices1.Zone = types.StringValue(meshServicesItem.Zone)
				if meshServicesCount+1 > len(r.Status.MeshServices) {
					r.Status.MeshServices = append(r.Status.MeshServices, meshServices1)
				} else {
					r.Status.MeshServices[meshServicesCount].Mesh = meshServices1.Mesh
					r.Status.MeshServices[meshServicesCount].Name = meshServices1.Name
					r.Status.MeshServices[meshServicesCount].Namespace = meshServices1.Namespace
					r.Status.MeshServices[meshServicesCount].Zone = meshServices1.Zone
				}
			}
			r.Status.Vips = []tfTypes.Vip{}
			if len(r.Status.Vips) > len(resp.Status.Vips) {
				r.Status.Vips = r.Status.Vips[:len(resp.Status.Vips)]
			}
			for vipsCount, vipsItem := range resp.Status.Vips {
				var vips1 tfTypes.Vip
				vips1.IP = types.StringPointerValue(vipsItem.IP)
				if vipsCount+1 > len(r.Status.Vips) {
					r.Status.Vips = append(r.Status.Vips, vips1)
				} else {
					r.Status.Vips[vipsCount].IP = vips1.IP
				}
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}
}
