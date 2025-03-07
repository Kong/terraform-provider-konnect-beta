// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/internal/utils"
	"time"
)

// MeshTLSItemType - the type of the resource
type MeshTLSItemType string

const (
	MeshTLSItemTypeMeshTLS MeshTLSItemType = "MeshTLS"
)

func (e MeshTLSItemType) ToPointer() *MeshTLSItemType {
	return &e
}
func (e *MeshTLSItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MeshTLS":
		*e = MeshTLSItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTLSItemType: %v", v)
	}
}

// MeshTLSItemSpecMode - Mode defines the behavior of inbound listeners with regard to traffic encryption.
type MeshTLSItemSpecMode string

const (
	MeshTLSItemSpecModePermissive MeshTLSItemSpecMode = "Permissive"
	MeshTLSItemSpecModeStrict     MeshTLSItemSpecMode = "Strict"
)

func (e MeshTLSItemSpecMode) ToPointer() *MeshTLSItemSpecMode {
	return &e
}
func (e *MeshTLSItemSpecMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Permissive":
		fallthrough
	case "Strict":
		*e = MeshTLSItemSpecMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTLSItemSpecMode: %v", v)
	}
}

type TLSCiphers string

const (
	TLSCiphersEcdheEcdsaAes128GcmSha256  TLSCiphers = "ECDHE-ECDSA-AES128-GCM-SHA256"
	TLSCiphersEcdheEcdsaAes256GcmSha384  TLSCiphers = "ECDHE-ECDSA-AES256-GCM-SHA384"
	TLSCiphersEcdheEcdsaChacha20Poly1305 TLSCiphers = "ECDHE-ECDSA-CHACHA20-POLY1305"
	TLSCiphersEcdheRsaAes128GcmSha256    TLSCiphers = "ECDHE-RSA-AES128-GCM-SHA256"
	TLSCiphersEcdheRsaAes256GcmSha384    TLSCiphers = "ECDHE-RSA-AES256-GCM-SHA384"
	TLSCiphersEcdheRsaChacha20Poly1305   TLSCiphers = "ECDHE-RSA-CHACHA20-POLY1305"
)

func (e TLSCiphers) ToPointer() *TLSCiphers {
	return &e
}
func (e *TLSCiphers) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ECDHE-ECDSA-AES128-GCM-SHA256":
		fallthrough
	case "ECDHE-ECDSA-AES256-GCM-SHA384":
		fallthrough
	case "ECDHE-ECDSA-CHACHA20-POLY1305":
		fallthrough
	case "ECDHE-RSA-AES128-GCM-SHA256":
		fallthrough
	case "ECDHE-RSA-AES256-GCM-SHA384":
		fallthrough
	case "ECDHE-RSA-CHACHA20-POLY1305":
		*e = TLSCiphers(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TLSCiphers: %v", v)
	}
}

// MeshTLSItemSpecMax - Max defines maximum supported version. One of `TLSAuto`, `TLS10`, `TLS11`, `TLS12`, `TLS13`.
type MeshTLSItemSpecMax string

const (
	MeshTLSItemSpecMaxTLSAuto MeshTLSItemSpecMax = "TLSAuto"
	MeshTLSItemSpecMaxTls10   MeshTLSItemSpecMax = "TLS10"
	MeshTLSItemSpecMaxTls11   MeshTLSItemSpecMax = "TLS11"
	MeshTLSItemSpecMaxTls12   MeshTLSItemSpecMax = "TLS12"
	MeshTLSItemSpecMaxTls13   MeshTLSItemSpecMax = "TLS13"
)

func (e MeshTLSItemSpecMax) ToPointer() *MeshTLSItemSpecMax {
	return &e
}
func (e *MeshTLSItemSpecMax) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSAuto":
		fallthrough
	case "TLS10":
		fallthrough
	case "TLS11":
		fallthrough
	case "TLS12":
		fallthrough
	case "TLS13":
		*e = MeshTLSItemSpecMax(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTLSItemSpecMax: %v", v)
	}
}

// MeshTLSItemSpecMin - Min defines minimum supported version. One of `TLSAuto`, `TLS10`, `TLS11`, `TLS12`, `TLS13`.
type MeshTLSItemSpecMin string

const (
	MeshTLSItemSpecMinTLSAuto MeshTLSItemSpecMin = "TLSAuto"
	MeshTLSItemSpecMinTls10   MeshTLSItemSpecMin = "TLS10"
	MeshTLSItemSpecMinTls11   MeshTLSItemSpecMin = "TLS11"
	MeshTLSItemSpecMinTls12   MeshTLSItemSpecMin = "TLS12"
	MeshTLSItemSpecMinTls13   MeshTLSItemSpecMin = "TLS13"
)

func (e MeshTLSItemSpecMin) ToPointer() *MeshTLSItemSpecMin {
	return &e
}
func (e *MeshTLSItemSpecMin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSAuto":
		fallthrough
	case "TLS10":
		fallthrough
	case "TLS11":
		fallthrough
	case "TLS12":
		fallthrough
	case "TLS13":
		*e = MeshTLSItemSpecMin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTLSItemSpecMin: %v", v)
	}
}

// TLSVersion - Version section for providing version specification.
type TLSVersion struct {
	// Max defines maximum supported version. One of `TLSAuto`, `TLS10`, `TLS11`, `TLS12`, `TLS13`.
	Max *MeshTLSItemSpecMax `default:"TLSAuto" json:"max"`
	// Min defines minimum supported version. One of `TLSAuto`, `TLS10`, `TLS11`, `TLS12`, `TLS13`.
	Min *MeshTLSItemSpecMin `default:"TLSAuto" json:"min"`
}

func (t TLSVersion) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSVersion) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSVersion) GetMax() *MeshTLSItemSpecMax {
	if o == nil {
		return nil
	}
	return o.Max
}

func (o *TLSVersion) GetMin() *MeshTLSItemSpecMin {
	if o == nil {
		return nil
	}
	return o.Min
}

// MeshTLSItemDefault - Default is a configuration specific to the group of clients referenced in
// 'targetRef'
type MeshTLSItemDefault struct {
	// Mode defines the behavior of inbound listeners with regard to traffic encryption.
	Mode *MeshTLSItemSpecMode `json:"mode,omitempty"`
	// TlsCiphers section for providing ciphers specification.
	TLSCiphers []TLSCiphers `json:"tlsCiphers,omitempty"`
	// Version section for providing version specification.
	TLSVersion *TLSVersion `json:"tlsVersion,omitempty"`
}

func (o *MeshTLSItemDefault) GetMode() *MeshTLSItemSpecMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *MeshTLSItemDefault) GetTLSCiphers() []TLSCiphers {
	if o == nil {
		return nil
	}
	return o.TLSCiphers
}

func (o *MeshTLSItemDefault) GetTLSVersion() *TLSVersion {
	if o == nil {
		return nil
	}
	return o.TLSVersion
}

// MeshTLSItemSpecKind - Kind of the referenced resource
type MeshTLSItemSpecKind string

const (
	MeshTLSItemSpecKindMesh                 MeshTLSItemSpecKind = "Mesh"
	MeshTLSItemSpecKindMeshSubset           MeshTLSItemSpecKind = "MeshSubset"
	MeshTLSItemSpecKindMeshGateway          MeshTLSItemSpecKind = "MeshGateway"
	MeshTLSItemSpecKindMeshService          MeshTLSItemSpecKind = "MeshService"
	MeshTLSItemSpecKindMeshExternalService  MeshTLSItemSpecKind = "MeshExternalService"
	MeshTLSItemSpecKindMeshMultiZoneService MeshTLSItemSpecKind = "MeshMultiZoneService"
	MeshTLSItemSpecKindMeshServiceSubset    MeshTLSItemSpecKind = "MeshServiceSubset"
	MeshTLSItemSpecKindMeshHTTPRoute        MeshTLSItemSpecKind = "MeshHTTPRoute"
	MeshTLSItemSpecKindDataplane            MeshTLSItemSpecKind = "Dataplane"
)

func (e MeshTLSItemSpecKind) ToPointer() *MeshTLSItemSpecKind {
	return &e
}
func (e *MeshTLSItemSpecKind) UnmarshalJSON(data []byte) error {
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
		*e = MeshTLSItemSpecKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTLSItemSpecKind: %v", v)
	}
}

type MeshTLSItemSpecProxyTypes string

const (
	MeshTLSItemSpecProxyTypesSidecar MeshTLSItemSpecProxyTypes = "Sidecar"
	MeshTLSItemSpecProxyTypesGateway MeshTLSItemSpecProxyTypes = "Gateway"
)

func (e MeshTLSItemSpecProxyTypes) ToPointer() *MeshTLSItemSpecProxyTypes {
	return &e
}
func (e *MeshTLSItemSpecProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshTLSItemSpecProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTLSItemSpecProxyTypes: %v", v)
	}
}

// MeshTLSItemSpecTargetRef - TargetRef is a reference to the resource that represents a group of
// clients.
type MeshTLSItemSpecTargetRef struct {
	// Kind of the referenced resource
	Kind MeshTLSItemSpecKind `json:"kind"`
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
	ProxyTypes []MeshTLSItemSpecProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshTLSItemSpecTargetRef) GetKind() MeshTLSItemSpecKind {
	if o == nil {
		return MeshTLSItemSpecKind("")
	}
	return o.Kind
}

func (o *MeshTLSItemSpecTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTLSItemSpecTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTLSItemSpecTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshTLSItemSpecTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshTLSItemSpecTargetRef) GetProxyTypes() []MeshTLSItemSpecProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshTLSItemSpecTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshTLSItemSpecTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type MeshTLSItemFrom struct {
	// Default is a configuration specific to the group of clients referenced in
	// 'targetRef'
	Default *MeshTLSItemDefault `json:"default,omitempty"`
	// TargetRef is a reference to the resource that represents a group of
	// clients.
	TargetRef MeshTLSItemSpecTargetRef `json:"targetRef"`
}

func (o *MeshTLSItemFrom) GetDefault() *MeshTLSItemDefault {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *MeshTLSItemFrom) GetTargetRef() MeshTLSItemSpecTargetRef {
	if o == nil {
		return MeshTLSItemSpecTargetRef{}
	}
	return o.TargetRef
}

// MeshTLSItemMode - Mode defines the behavior of inbound listeners with regard to traffic encryption.
type MeshTLSItemMode string

const (
	MeshTLSItemModePermissive MeshTLSItemMode = "Permissive"
	MeshTLSItemModeStrict     MeshTLSItemMode = "Strict"
)

func (e MeshTLSItemMode) ToPointer() *MeshTLSItemMode {
	return &e
}
func (e *MeshTLSItemMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Permissive":
		fallthrough
	case "Strict":
		*e = MeshTLSItemMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTLSItemMode: %v", v)
	}
}

type MeshTLSItemTLSCiphers string

const (
	MeshTLSItemTLSCiphersEcdheEcdsaAes128GcmSha256  MeshTLSItemTLSCiphers = "ECDHE-ECDSA-AES128-GCM-SHA256"
	MeshTLSItemTLSCiphersEcdheEcdsaAes256GcmSha384  MeshTLSItemTLSCiphers = "ECDHE-ECDSA-AES256-GCM-SHA384"
	MeshTLSItemTLSCiphersEcdheEcdsaChacha20Poly1305 MeshTLSItemTLSCiphers = "ECDHE-ECDSA-CHACHA20-POLY1305"
	MeshTLSItemTLSCiphersEcdheRsaAes128GcmSha256    MeshTLSItemTLSCiphers = "ECDHE-RSA-AES128-GCM-SHA256"
	MeshTLSItemTLSCiphersEcdheRsaAes256GcmSha384    MeshTLSItemTLSCiphers = "ECDHE-RSA-AES256-GCM-SHA384"
	MeshTLSItemTLSCiphersEcdheRsaChacha20Poly1305   MeshTLSItemTLSCiphers = "ECDHE-RSA-CHACHA20-POLY1305"
)

func (e MeshTLSItemTLSCiphers) ToPointer() *MeshTLSItemTLSCiphers {
	return &e
}
func (e *MeshTLSItemTLSCiphers) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ECDHE-ECDSA-AES128-GCM-SHA256":
		fallthrough
	case "ECDHE-ECDSA-AES256-GCM-SHA384":
		fallthrough
	case "ECDHE-ECDSA-CHACHA20-POLY1305":
		fallthrough
	case "ECDHE-RSA-AES128-GCM-SHA256":
		fallthrough
	case "ECDHE-RSA-AES256-GCM-SHA384":
		fallthrough
	case "ECDHE-RSA-CHACHA20-POLY1305":
		*e = MeshTLSItemTLSCiphers(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTLSItemTLSCiphers: %v", v)
	}
}

// MeshTLSItemMax - Max defines maximum supported version. One of `TLSAuto`, `TLS10`, `TLS11`, `TLS12`, `TLS13`.
type MeshTLSItemMax string

const (
	MeshTLSItemMaxTLSAuto MeshTLSItemMax = "TLSAuto"
	MeshTLSItemMaxTls10   MeshTLSItemMax = "TLS10"
	MeshTLSItemMaxTls11   MeshTLSItemMax = "TLS11"
	MeshTLSItemMaxTls12   MeshTLSItemMax = "TLS12"
	MeshTLSItemMaxTls13   MeshTLSItemMax = "TLS13"
)

func (e MeshTLSItemMax) ToPointer() *MeshTLSItemMax {
	return &e
}
func (e *MeshTLSItemMax) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSAuto":
		fallthrough
	case "TLS10":
		fallthrough
	case "TLS11":
		fallthrough
	case "TLS12":
		fallthrough
	case "TLS13":
		*e = MeshTLSItemMax(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTLSItemMax: %v", v)
	}
}

// MeshTLSItemMin - Min defines minimum supported version. One of `TLSAuto`, `TLS10`, `TLS11`, `TLS12`, `TLS13`.
type MeshTLSItemMin string

const (
	MeshTLSItemMinTLSAuto MeshTLSItemMin = "TLSAuto"
	MeshTLSItemMinTls10   MeshTLSItemMin = "TLS10"
	MeshTLSItemMinTls11   MeshTLSItemMin = "TLS11"
	MeshTLSItemMinTls12   MeshTLSItemMin = "TLS12"
	MeshTLSItemMinTls13   MeshTLSItemMin = "TLS13"
)

func (e MeshTLSItemMin) ToPointer() *MeshTLSItemMin {
	return &e
}
func (e *MeshTLSItemMin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSAuto":
		fallthrough
	case "TLS10":
		fallthrough
	case "TLS11":
		fallthrough
	case "TLS12":
		fallthrough
	case "TLS13":
		*e = MeshTLSItemMin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTLSItemMin: %v", v)
	}
}

// MeshTLSItemTLSVersion - Version section for providing version specification.
type MeshTLSItemTLSVersion struct {
	// Max defines maximum supported version. One of `TLSAuto`, `TLS10`, `TLS11`, `TLS12`, `TLS13`.
	Max *MeshTLSItemMax `default:"TLSAuto" json:"max"`
	// Min defines minimum supported version. One of `TLSAuto`, `TLS10`, `TLS11`, `TLS12`, `TLS13`.
	Min *MeshTLSItemMin `default:"TLSAuto" json:"min"`
}

func (m MeshTLSItemTLSVersion) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshTLSItemTLSVersion) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshTLSItemTLSVersion) GetMax() *MeshTLSItemMax {
	if o == nil {
		return nil
	}
	return o.Max
}

func (o *MeshTLSItemTLSVersion) GetMin() *MeshTLSItemMin {
	if o == nil {
		return nil
	}
	return o.Min
}

// MeshTLSItemSpecDefault - Default contains configuration of the inbound tls
type MeshTLSItemSpecDefault struct {
	// Mode defines the behavior of inbound listeners with regard to traffic encryption.
	Mode *MeshTLSItemMode `json:"mode,omitempty"`
	// TlsCiphers section for providing ciphers specification.
	TLSCiphers []MeshTLSItemTLSCiphers `json:"tlsCiphers,omitempty"`
	// Version section for providing version specification.
	TLSVersion *MeshTLSItemTLSVersion `json:"tlsVersion,omitempty"`
}

func (o *MeshTLSItemSpecDefault) GetMode() *MeshTLSItemMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *MeshTLSItemSpecDefault) GetTLSCiphers() []MeshTLSItemTLSCiphers {
	if o == nil {
		return nil
	}
	return o.TLSCiphers
}

func (o *MeshTLSItemSpecDefault) GetTLSVersion() *MeshTLSItemTLSVersion {
	if o == nil {
		return nil
	}
	return o.TLSVersion
}

type MeshTLSItemRules struct {
	// Default contains configuration of the inbound tls
	Default *MeshTLSItemSpecDefault `json:"default,omitempty"`
}

func (o *MeshTLSItemRules) GetDefault() *MeshTLSItemSpecDefault {
	if o == nil {
		return nil
	}
	return o.Default
}

// MeshTLSItemKind - Kind of the referenced resource
type MeshTLSItemKind string

const (
	MeshTLSItemKindMesh                 MeshTLSItemKind = "Mesh"
	MeshTLSItemKindMeshSubset           MeshTLSItemKind = "MeshSubset"
	MeshTLSItemKindMeshGateway          MeshTLSItemKind = "MeshGateway"
	MeshTLSItemKindMeshService          MeshTLSItemKind = "MeshService"
	MeshTLSItemKindMeshExternalService  MeshTLSItemKind = "MeshExternalService"
	MeshTLSItemKindMeshMultiZoneService MeshTLSItemKind = "MeshMultiZoneService"
	MeshTLSItemKindMeshServiceSubset    MeshTLSItemKind = "MeshServiceSubset"
	MeshTLSItemKindMeshHTTPRoute        MeshTLSItemKind = "MeshHTTPRoute"
	MeshTLSItemKindDataplane            MeshTLSItemKind = "Dataplane"
)

func (e MeshTLSItemKind) ToPointer() *MeshTLSItemKind {
	return &e
}
func (e *MeshTLSItemKind) UnmarshalJSON(data []byte) error {
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
		*e = MeshTLSItemKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTLSItemKind: %v", v)
	}
}

type MeshTLSItemProxyTypes string

const (
	MeshTLSItemProxyTypesSidecar MeshTLSItemProxyTypes = "Sidecar"
	MeshTLSItemProxyTypesGateway MeshTLSItemProxyTypes = "Gateway"
)

func (e MeshTLSItemProxyTypes) ToPointer() *MeshTLSItemProxyTypes {
	return &e
}
func (e *MeshTLSItemProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshTLSItemProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTLSItemProxyTypes: %v", v)
	}
}

// MeshTLSItemTargetRef - TargetRef is a reference to the resource the policy takes an effect on.
// The resource could be either a real store object or virtual resource
// defined in-place.
type MeshTLSItemTargetRef struct {
	// Kind of the referenced resource
	Kind MeshTLSItemKind `json:"kind"`
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
	ProxyTypes []MeshTLSItemProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshTLSItemTargetRef) GetKind() MeshTLSItemKind {
	if o == nil {
		return MeshTLSItemKind("")
	}
	return o.Kind
}

func (o *MeshTLSItemTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTLSItemTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTLSItemTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshTLSItemTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshTLSItemTargetRef) GetProxyTypes() []MeshTLSItemProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshTLSItemTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshTLSItemTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// MeshTLSItemSpec - Spec is the specification of the Kuma MeshTLS resource.
type MeshTLSItemSpec struct {
	// From list makes a match between clients and corresponding configurations
	From []MeshTLSItemFrom `json:"from,omitempty"`
	// Rules defines inbound tls configurations. Currently limited to
	// selecting all inbound traffic, as L7 matching is not yet implemented.
	Rules []MeshTLSItemRules `json:"rules,omitempty"`
	// TargetRef is a reference to the resource the policy takes an effect on.
	// The resource could be either a real store object or virtual resource
	// defined in-place.
	TargetRef *MeshTLSItemTargetRef `json:"targetRef,omitempty"`
}

func (o *MeshTLSItemSpec) GetFrom() []MeshTLSItemFrom {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *MeshTLSItemSpec) GetRules() []MeshTLSItemRules {
	if o == nil {
		return nil
	}
	return o.Rules
}

func (o *MeshTLSItemSpec) GetTargetRef() *MeshTLSItemTargetRef {
	if o == nil {
		return nil
	}
	return o.TargetRef
}

// MeshTLSItem - Successful response
type MeshTLSItem struct {
	// the type of the resource
	Type MeshTLSItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshTLS resource.
	Spec MeshTLSItemSpec `json:"spec"`
	// Time at which the resource was created
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// Time at which the resource was updated
	ModificationTime *time.Time `json:"modificationTime,omitempty"`
}

func (m MeshTLSItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshTLSItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshTLSItem) GetType() MeshTLSItemType {
	if o == nil {
		return MeshTLSItemType("")
	}
	return o.Type
}

func (o *MeshTLSItem) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTLSItem) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshTLSItem) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTLSItem) GetSpec() MeshTLSItemSpec {
	if o == nil {
		return MeshTLSItemSpec{}
	}
	return o.Spec
}

func (o *MeshTLSItem) GetCreationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *MeshTLSItem) GetModificationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModificationTime
}

type MeshTLSItemInput struct {
	// the type of the resource
	Type MeshTLSItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshTLS resource.
	Spec MeshTLSItemSpec `json:"spec"`
}

func (m MeshTLSItemInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshTLSItemInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshTLSItemInput) GetType() MeshTLSItemType {
	if o == nil {
		return MeshTLSItemType("")
	}
	return o.Type
}

func (o *MeshTLSItemInput) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTLSItemInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshTLSItemInput) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTLSItemInput) GetSpec() MeshTLSItemSpec {
	if o == nil {
		return MeshTLSItemSpec{}
	}
	return o.Spec
}
