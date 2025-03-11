// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/internal/utils"
	"time"
)

// MeshTrafficPermissionItemType - the type of the resource
type MeshTrafficPermissionItemType string

const (
	MeshTrafficPermissionItemTypeMeshTrafficPermission MeshTrafficPermissionItemType = "MeshTrafficPermission"
)

func (e MeshTrafficPermissionItemType) ToPointer() *MeshTrafficPermissionItemType {
	return &e
}
func (e *MeshTrafficPermissionItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MeshTrafficPermission":
		*e = MeshTrafficPermissionItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTrafficPermissionItemType: %v", v)
	}
}

// Action defines a behavior for the specified group of clients:
type Action string

const (
	ActionAllow               Action = "Allow"
	ActionDeny                Action = "Deny"
	ActionAllowWithShadowDeny Action = "AllowWithShadowDeny"
)

func (e Action) ToPointer() *Action {
	return &e
}
func (e *Action) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Allow":
		fallthrough
	case "Deny":
		fallthrough
	case "AllowWithShadowDeny":
		*e = Action(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Action: %v", v)
	}
}

// MeshTrafficPermissionItemDefault - Default is a configuration specific to the group of clients referenced in
// 'targetRef'
type MeshTrafficPermissionItemDefault struct {
	// Action defines a behavior for the specified group of clients:
	Action *Action `json:"action,omitempty"`
}

func (o *MeshTrafficPermissionItemDefault) GetAction() *Action {
	if o == nil {
		return nil
	}
	return o.Action
}

// MeshTrafficPermissionItemSpecKind - Kind of the referenced resource
type MeshTrafficPermissionItemSpecKind string

const (
	MeshTrafficPermissionItemSpecKindMesh                 MeshTrafficPermissionItemSpecKind = "Mesh"
	MeshTrafficPermissionItemSpecKindMeshSubset           MeshTrafficPermissionItemSpecKind = "MeshSubset"
	MeshTrafficPermissionItemSpecKindMeshGateway          MeshTrafficPermissionItemSpecKind = "MeshGateway"
	MeshTrafficPermissionItemSpecKindMeshService          MeshTrafficPermissionItemSpecKind = "MeshService"
	MeshTrafficPermissionItemSpecKindMeshExternalService  MeshTrafficPermissionItemSpecKind = "MeshExternalService"
	MeshTrafficPermissionItemSpecKindMeshMultiZoneService MeshTrafficPermissionItemSpecKind = "MeshMultiZoneService"
	MeshTrafficPermissionItemSpecKindMeshServiceSubset    MeshTrafficPermissionItemSpecKind = "MeshServiceSubset"
	MeshTrafficPermissionItemSpecKindMeshHTTPRoute        MeshTrafficPermissionItemSpecKind = "MeshHTTPRoute"
	MeshTrafficPermissionItemSpecKindDataplane            MeshTrafficPermissionItemSpecKind = "Dataplane"
)

func (e MeshTrafficPermissionItemSpecKind) ToPointer() *MeshTrafficPermissionItemSpecKind {
	return &e
}
func (e *MeshTrafficPermissionItemSpecKind) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Mesh":
		fallthrough
	case "MeshSubset":
		fallthrough
	case "MeshGateway":
		fallthrough
	case "MeshService":
		fallthrough
	case "MeshExternalService":
		fallthrough
	case "MeshMultiZoneService":
		fallthrough
	case "MeshServiceSubset":
		fallthrough
	case "MeshHTTPRoute":
		fallthrough
	case "Dataplane":
		*e = MeshTrafficPermissionItemSpecKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTrafficPermissionItemSpecKind: %v", v)
	}
}

type MeshTrafficPermissionItemSpecProxyTypes string

const (
	MeshTrafficPermissionItemSpecProxyTypesSidecar MeshTrafficPermissionItemSpecProxyTypes = "Sidecar"
	MeshTrafficPermissionItemSpecProxyTypesGateway MeshTrafficPermissionItemSpecProxyTypes = "Gateway"
)

func (e MeshTrafficPermissionItemSpecProxyTypes) ToPointer() *MeshTrafficPermissionItemSpecProxyTypes {
	return &e
}
func (e *MeshTrafficPermissionItemSpecProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshTrafficPermissionItemSpecProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTrafficPermissionItemSpecProxyTypes: %v", v)
	}
}

// MeshTrafficPermissionItemSpecTargetRef - TargetRef is a reference to the resource that represents a group of
// clients.
type MeshTrafficPermissionItemSpecTargetRef struct {
	// Kind of the referenced resource
	Kind MeshTrafficPermissionItemSpecKind `json:"kind"`
	// Labels are used to select group of MeshServices that match labels. Either Labels or
	// Name and Namespace can be used.
	Labels map[string]string `json:"labels,omitempty"`
	// Mesh is reserved for future use to identify cross mesh resources.
	Mesh *string `json:"mesh,omitempty"`
	// Name of the referenced resource. Can only be used with kinds: `MeshService`,
	// `MeshServiceSubset` and `MeshGatewayRoute`
	Name *string `json:"name,omitempty"`
	// Namespace specifies the namespace of target resource. If empty only resources in policy namespace
	// will be targeted.
	Namespace *string `json:"namespace,omitempty"`
	// ProxyTypes specifies the data plane types that are subject to the policy. When not specified,
	// all data plane types are targeted by the policy.
	ProxyTypes []MeshTrafficPermissionItemSpecProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshTrafficPermissionItemSpecTargetRef) GetKind() MeshTrafficPermissionItemSpecKind {
	if o == nil {
		return MeshTrafficPermissionItemSpecKind("")
	}
	return o.Kind
}

func (o *MeshTrafficPermissionItemSpecTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTrafficPermissionItemSpecTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTrafficPermissionItemSpecTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshTrafficPermissionItemSpecTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshTrafficPermissionItemSpecTargetRef) GetProxyTypes() []MeshTrafficPermissionItemSpecProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshTrafficPermissionItemSpecTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshTrafficPermissionItemSpecTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type MeshTrafficPermissionItemFrom struct {
	// Default is a configuration specific to the group of clients referenced in
	// 'targetRef'
	Default *MeshTrafficPermissionItemDefault `json:"default,omitempty"`
	// TargetRef is a reference to the resource that represents a group of
	// clients.
	TargetRef MeshTrafficPermissionItemSpecTargetRef `json:"targetRef"`
}

func (o *MeshTrafficPermissionItemFrom) GetDefault() *MeshTrafficPermissionItemDefault {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *MeshTrafficPermissionItemFrom) GetTargetRef() MeshTrafficPermissionItemSpecTargetRef {
	if o == nil {
		return MeshTrafficPermissionItemSpecTargetRef{}
	}
	return o.TargetRef
}

// MeshTrafficPermissionItemKind - Kind of the referenced resource
type MeshTrafficPermissionItemKind string

const (
	MeshTrafficPermissionItemKindMesh                 MeshTrafficPermissionItemKind = "Mesh"
	MeshTrafficPermissionItemKindMeshSubset           MeshTrafficPermissionItemKind = "MeshSubset"
	MeshTrafficPermissionItemKindMeshGateway          MeshTrafficPermissionItemKind = "MeshGateway"
	MeshTrafficPermissionItemKindMeshService          MeshTrafficPermissionItemKind = "MeshService"
	MeshTrafficPermissionItemKindMeshExternalService  MeshTrafficPermissionItemKind = "MeshExternalService"
	MeshTrafficPermissionItemKindMeshMultiZoneService MeshTrafficPermissionItemKind = "MeshMultiZoneService"
	MeshTrafficPermissionItemKindMeshServiceSubset    MeshTrafficPermissionItemKind = "MeshServiceSubset"
	MeshTrafficPermissionItemKindMeshHTTPRoute        MeshTrafficPermissionItemKind = "MeshHTTPRoute"
	MeshTrafficPermissionItemKindDataplane            MeshTrafficPermissionItemKind = "Dataplane"
)

func (e MeshTrafficPermissionItemKind) ToPointer() *MeshTrafficPermissionItemKind {
	return &e
}
func (e *MeshTrafficPermissionItemKind) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Mesh":
		fallthrough
	case "MeshSubset":
		fallthrough
	case "MeshGateway":
		fallthrough
	case "MeshService":
		fallthrough
	case "MeshExternalService":
		fallthrough
	case "MeshMultiZoneService":
		fallthrough
	case "MeshServiceSubset":
		fallthrough
	case "MeshHTTPRoute":
		fallthrough
	case "Dataplane":
		*e = MeshTrafficPermissionItemKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTrafficPermissionItemKind: %v", v)
	}
}

type MeshTrafficPermissionItemProxyTypes string

const (
	MeshTrafficPermissionItemProxyTypesSidecar MeshTrafficPermissionItemProxyTypes = "Sidecar"
	MeshTrafficPermissionItemProxyTypesGateway MeshTrafficPermissionItemProxyTypes = "Gateway"
)

func (e MeshTrafficPermissionItemProxyTypes) ToPointer() *MeshTrafficPermissionItemProxyTypes {
	return &e
}
func (e *MeshTrafficPermissionItemProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshTrafficPermissionItemProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTrafficPermissionItemProxyTypes: %v", v)
	}
}

// MeshTrafficPermissionItemTargetRef - TargetRef is a reference to the resource the policy takes an effect on.
// The resource could be either a real store object or virtual resource
// defined inplace.
type MeshTrafficPermissionItemTargetRef struct {
	// Kind of the referenced resource
	Kind MeshTrafficPermissionItemKind `json:"kind"`
	// Labels are used to select group of MeshServices that match labels. Either Labels or
	// Name and Namespace can be used.
	Labels map[string]string `json:"labels,omitempty"`
	// Mesh is reserved for future use to identify cross mesh resources.
	Mesh *string `json:"mesh,omitempty"`
	// Name of the referenced resource. Can only be used with kinds: `MeshService`,
	// `MeshServiceSubset` and `MeshGatewayRoute`
	Name *string `json:"name,omitempty"`
	// Namespace specifies the namespace of target resource. If empty only resources in policy namespace
	// will be targeted.
	Namespace *string `json:"namespace,omitempty"`
	// ProxyTypes specifies the data plane types that are subject to the policy. When not specified,
	// all data plane types are targeted by the policy.
	ProxyTypes []MeshTrafficPermissionItemProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshTrafficPermissionItemTargetRef) GetKind() MeshTrafficPermissionItemKind {
	if o == nil {
		return MeshTrafficPermissionItemKind("")
	}
	return o.Kind
}

func (o *MeshTrafficPermissionItemTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTrafficPermissionItemTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTrafficPermissionItemTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshTrafficPermissionItemTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshTrafficPermissionItemTargetRef) GetProxyTypes() []MeshTrafficPermissionItemProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshTrafficPermissionItemTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshTrafficPermissionItemTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// MeshTrafficPermissionItemSpec - Spec is the specification of the Kuma MeshTrafficPermission resource.
type MeshTrafficPermissionItemSpec struct {
	// From list makes a match between clients and corresponding configurations
	From []MeshTrafficPermissionItemFrom `json:"from,omitempty"`
	// TargetRef is a reference to the resource the policy takes an effect on.
	// The resource could be either a real store object or virtual resource
	// defined inplace.
	TargetRef *MeshTrafficPermissionItemTargetRef `json:"targetRef,omitempty"`
}

func (o *MeshTrafficPermissionItemSpec) GetFrom() []MeshTrafficPermissionItemFrom {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *MeshTrafficPermissionItemSpec) GetTargetRef() *MeshTrafficPermissionItemTargetRef {
	if o == nil {
		return nil
	}
	return o.TargetRef
}

// MeshTrafficPermissionItem - Successful response
type MeshTrafficPermissionItem struct {
	// the type of the resource
	Type MeshTrafficPermissionItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshTrafficPermission resource.
	Spec MeshTrafficPermissionItemSpec `json:"spec"`
	// Time at which the resource was created
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// Time at which the resource was updated
	ModificationTime *time.Time `json:"modificationTime,omitempty"`
}

func (m MeshTrafficPermissionItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshTrafficPermissionItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshTrafficPermissionItem) GetType() MeshTrafficPermissionItemType {
	if o == nil {
		return MeshTrafficPermissionItemType("")
	}
	return o.Type
}

func (o *MeshTrafficPermissionItem) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTrafficPermissionItem) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshTrafficPermissionItem) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTrafficPermissionItem) GetSpec() MeshTrafficPermissionItemSpec {
	if o == nil {
		return MeshTrafficPermissionItemSpec{}
	}
	return o.Spec
}

func (o *MeshTrafficPermissionItem) GetCreationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *MeshTrafficPermissionItem) GetModificationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModificationTime
}

type MeshTrafficPermissionItemInput struct {
	// the type of the resource
	Type MeshTrafficPermissionItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshTrafficPermission resource.
	Spec MeshTrafficPermissionItemSpec `json:"spec"`
}

func (m MeshTrafficPermissionItemInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshTrafficPermissionItemInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshTrafficPermissionItemInput) GetType() MeshTrafficPermissionItemType {
	if o == nil {
		return MeshTrafficPermissionItemType("")
	}
	return o.Type
}

func (o *MeshTrafficPermissionItemInput) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTrafficPermissionItemInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshTrafficPermissionItemInput) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTrafficPermissionItemInput) GetSpec() MeshTrafficPermissionItemSpec {
	if o == nil {
		return MeshTrafficPermissionItemSpec{}
	}
	return o.Spec
}
