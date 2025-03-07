// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/internal/utils"
	"time"
)

// MeshExternalServiceItemType - the type of the resource
type MeshExternalServiceItemType string

const (
	MeshExternalServiceItemTypeMeshExternalService MeshExternalServiceItemType = "MeshExternalService"
)

func (e MeshExternalServiceItemType) ToPointer() *MeshExternalServiceItemType {
	return &e
}
func (e *MeshExternalServiceItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MeshExternalService":
		*e = MeshExternalServiceItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshExternalServiceItemType: %v", v)
	}
}

type Endpoints struct {
	// Address defines an address to which a user want to send a request. Is possible to provide `domain`, `ip`.
	Address string `json:"address"`
	// Port of the endpoint
	Port int64 `json:"port"`
}

func (o *Endpoints) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *Endpoints) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

// Extension struct for a plugin configuration, in the presence of an extension `endpoints` and `tls` are not required anymore - it's up to the extension to validate them independently.
type Extension struct {
	// Config freeform configuration for the extension.
	Config any `json:"config,omitempty"`
	// Type of the extension.
	Type string `json:"type"`
}

func (o *Extension) GetConfig() any {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *Extension) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

// MeshExternalServiceItemProtocol - Protocol defines a protocol of the communication. Possible values: `tcp`, `grpc`, `http`, `http2`.
type MeshExternalServiceItemProtocol string

const (
	MeshExternalServiceItemProtocolTCP   MeshExternalServiceItemProtocol = "tcp"
	MeshExternalServiceItemProtocolGrpc  MeshExternalServiceItemProtocol = "grpc"
	MeshExternalServiceItemProtocolHTTP  MeshExternalServiceItemProtocol = "http"
	MeshExternalServiceItemProtocolHttp2 MeshExternalServiceItemProtocol = "http2"
)

func (e MeshExternalServiceItemProtocol) ToPointer() *MeshExternalServiceItemProtocol {
	return &e
}
func (e *MeshExternalServiceItemProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tcp":
		fallthrough
	case "grpc":
		fallthrough
	case "http":
		fallthrough
	case "http2":
		*e = MeshExternalServiceItemProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshExternalServiceItemProtocol: %v", v)
	}
}

// MeshExternalServiceItemSpecType - Type of the match, only `HostnameGenerator` is available at the moment.
type MeshExternalServiceItemSpecType string

const (
	MeshExternalServiceItemSpecTypeHostnameGenerator MeshExternalServiceItemSpecType = "HostnameGenerator"
)

func (e MeshExternalServiceItemSpecType) ToPointer() *MeshExternalServiceItemSpecType {
	return &e
}
func (e *MeshExternalServiceItemSpecType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HostnameGenerator":
		*e = MeshExternalServiceItemSpecType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshExternalServiceItemSpecType: %v", v)
	}
}

// Match defines traffic that should be routed through the sidecar.
type Match struct {
	// Port defines a port to which a user does request.
	Port int64 `json:"port"`
	// Protocol defines a protocol of the communication. Possible values: `tcp`, `grpc`, `http`, `http2`.
	Protocol *MeshExternalServiceItemProtocol `default:"tcp" json:"protocol"`
	// Type of the match, only `HostnameGenerator` is available at the moment.
	Type *MeshExternalServiceItemSpecType `default:"HostnameGenerator" json:"type"`
}

func (m Match) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *Match) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Match) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *Match) GetProtocol() *MeshExternalServiceItemProtocol {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *Match) GetType() *MeshExternalServiceItemSpecType {
	if o == nil {
		return nil
	}
	return o.Type
}

// CaCert defines a certificate of CA.
type CaCert struct {
	// Data source is inline bytes.
	Inline *string `json:"inline,omitempty"`
	// Data source is inline string`
	InlineString *string `json:"inlineString,omitempty"`
	// Data source is a secret with given Secret key.
	Secret *string `json:"secret,omitempty"`
}

func (o *CaCert) GetInline() *string {
	if o == nil {
		return nil
	}
	return o.Inline
}

func (o *CaCert) GetInlineString() *string {
	if o == nil {
		return nil
	}
	return o.InlineString
}

func (o *CaCert) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

// ClientCert defines a certificate of a client.
type ClientCert struct {
	// Data source is inline bytes.
	Inline *string `json:"inline,omitempty"`
	// Data source is inline string`
	InlineString *string `json:"inlineString,omitempty"`
	// Data source is a secret with given Secret key.
	Secret *string `json:"secret,omitempty"`
}

func (o *ClientCert) GetInline() *string {
	if o == nil {
		return nil
	}
	return o.Inline
}

func (o *ClientCert) GetInlineString() *string {
	if o == nil {
		return nil
	}
	return o.InlineString
}

func (o *ClientCert) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

// ClientKey defines a client private key.
type ClientKey struct {
	// Data source is inline bytes.
	Inline *string `json:"inline,omitempty"`
	// Data source is inline string`
	InlineString *string `json:"inlineString,omitempty"`
	// Data source is a secret with given Secret key.
	Secret *string `json:"secret,omitempty"`
}

func (o *ClientKey) GetInline() *string {
	if o == nil {
		return nil
	}
	return o.Inline
}

func (o *ClientKey) GetInlineString() *string {
	if o == nil {
		return nil
	}
	return o.InlineString
}

func (o *ClientKey) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

// MeshExternalServiceItemMode - Mode defines if proxy should skip verification, one of `SkipSAN`, `SkipCA`, `Secured`, `SkipAll`. Default `Secured`.
type MeshExternalServiceItemMode string

const (
	MeshExternalServiceItemModeSkipSan MeshExternalServiceItemMode = "SkipSAN"
	MeshExternalServiceItemModeSkipCa  MeshExternalServiceItemMode = "SkipCA"
	MeshExternalServiceItemModeSecured MeshExternalServiceItemMode = "Secured"
	MeshExternalServiceItemModeSkipAll MeshExternalServiceItemMode = "SkipAll"
)

func (e MeshExternalServiceItemMode) ToPointer() *MeshExternalServiceItemMode {
	return &e
}
func (e *MeshExternalServiceItemMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SkipSAN":
		fallthrough
	case "SkipCA":
		fallthrough
	case "Secured":
		fallthrough
	case "SkipAll":
		*e = MeshExternalServiceItemMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshExternalServiceItemMode: %v", v)
	}
}

// MeshExternalServiceItemSpecTLSType - Type specifies matching type, one of `Exact`, `Prefix`. Default: `Exact`
type MeshExternalServiceItemSpecTLSType string

const (
	MeshExternalServiceItemSpecTLSTypeExact  MeshExternalServiceItemSpecTLSType = "Exact"
	MeshExternalServiceItemSpecTLSTypePrefix MeshExternalServiceItemSpecTLSType = "Prefix"
)

func (e MeshExternalServiceItemSpecTLSType) ToPointer() *MeshExternalServiceItemSpecTLSType {
	return &e
}
func (e *MeshExternalServiceItemSpecTLSType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Exact":
		fallthrough
	case "Prefix":
		*e = MeshExternalServiceItemSpecTLSType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshExternalServiceItemSpecTLSType: %v", v)
	}
}

type SubjectAltNames struct {
	// Type specifies matching type, one of `Exact`, `Prefix`. Default: `Exact`
	Type *MeshExternalServiceItemSpecTLSType `default:"Exact" json:"type"`
	// Value to match.
	Value string `json:"value"`
}

func (s SubjectAltNames) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SubjectAltNames) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SubjectAltNames) GetType() *MeshExternalServiceItemSpecTLSType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SubjectAltNames) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// Verification section for providing TLS verification details.
type Verification struct {
	// CaCert defines a certificate of CA.
	CaCert *CaCert `json:"caCert,omitempty"`
	// ClientCert defines a certificate of a client.
	ClientCert *ClientCert `json:"clientCert,omitempty"`
	// ClientKey defines a client private key.
	ClientKey *ClientKey `json:"clientKey,omitempty"`
	// Mode defines if proxy should skip verification, one of `SkipSAN`, `SkipCA`, `Secured`, `SkipAll`. Default `Secured`.
	Mode *MeshExternalServiceItemMode `default:"Secured" json:"mode"`
	// ServerName overrides the default Server Name Indicator set by Kuma.
	ServerName *string `json:"serverName,omitempty"`
	// SubjectAltNames list of names to verify in the certificate.
	SubjectAltNames []SubjectAltNames `json:"subjectAltNames,omitempty"`
}

func (v Verification) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *Verification) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Verification) GetCaCert() *CaCert {
	if o == nil {
		return nil
	}
	return o.CaCert
}

func (o *Verification) GetClientCert() *ClientCert {
	if o == nil {
		return nil
	}
	return o.ClientCert
}

func (o *Verification) GetClientKey() *ClientKey {
	if o == nil {
		return nil
	}
	return o.ClientKey
}

func (o *Verification) GetMode() *MeshExternalServiceItemMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *Verification) GetServerName() *string {
	if o == nil {
		return nil
	}
	return o.ServerName
}

func (o *Verification) GetSubjectAltNames() []SubjectAltNames {
	if o == nil {
		return nil
	}
	return o.SubjectAltNames
}

// Max defines maximum supported version. One of `TLSAuto`, `TLS10`, `TLS11`, `TLS12`, `TLS13`.
type Max string

const (
	MaxTLSAuto Max = "TLSAuto"
	MaxTls10   Max = "TLS10"
	MaxTls11   Max = "TLS11"
	MaxTls12   Max = "TLS12"
	MaxTls13   Max = "TLS13"
)

func (e Max) ToPointer() *Max {
	return &e
}
func (e *Max) UnmarshalJSON(data []byte) error {
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
		*e = Max(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Max: %v", v)
	}
}

// Min defines minimum supported version. One of `TLSAuto`, `TLS10`, `TLS11`, `TLS12`, `TLS13`.
type Min string

const (
	MinTLSAuto Min = "TLSAuto"
	MinTls10   Min = "TLS10"
	MinTls11   Min = "TLS11"
	MinTls12   Min = "TLS12"
	MinTls13   Min = "TLS13"
)

func (e Min) ToPointer() *Min {
	return &e
}
func (e *Min) UnmarshalJSON(data []byte) error {
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
		*e = Min(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Min: %v", v)
	}
}

// Version section for providing version specification.
type Version struct {
	// Max defines maximum supported version. One of `TLSAuto`, `TLS10`, `TLS11`, `TLS12`, `TLS13`.
	Max *Max `default:"TLSAuto" json:"max"`
	// Min defines minimum supported version. One of `TLSAuto`, `TLS10`, `TLS11`, `TLS12`, `TLS13`.
	Min *Min `default:"TLSAuto" json:"min"`
}

func (v Version) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *Version) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Version) GetMax() *Max {
	if o == nil {
		return nil
	}
	return o.Max
}

func (o *Version) GetMin() *Min {
	if o == nil {
		return nil
	}
	return o.Min
}

// TLS - Tls provides a TLS configuration when proxy is resposible for a TLS origination
type TLS struct {
	// AllowRenegotiation defines if TLS sessions will allow renegotiation.
	// Setting this to true is not recommended for security reasons.
	AllowRenegotiation *bool `default:"false" json:"allowRenegotiation"`
	// Enabled defines if proxy should originate TLS.
	Enabled *bool `default:"false" json:"enabled"`
	// Verification section for providing TLS verification details.
	Verification *Verification `json:"verification,omitempty"`
	// Version section for providing version specification.
	Version *Version `json:"version,omitempty"`
}

func (t TLS) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLS) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLS) GetAllowRenegotiation() *bool {
	if o == nil {
		return nil
	}
	return o.AllowRenegotiation
}

func (o *TLS) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *TLS) GetVerification() *Verification {
	if o == nil {
		return nil
	}
	return o.Verification
}

func (o *TLS) GetVersion() *Version {
	if o == nil {
		return nil
	}
	return o.Version
}

// MeshExternalServiceItemSpec - Spec is the specification of the Kuma MeshExternalService resource.
type MeshExternalServiceItemSpec struct {
	// Endpoints defines a list of destinations to send traffic to.
	Endpoints []Endpoints `json:"endpoints,omitempty"`
	// Extension struct for a plugin configuration, in the presence of an extension `endpoints` and `tls` are not required anymore - it's up to the extension to validate them independently.
	Extension *Extension `json:"extension,omitempty"`
	// Match defines traffic that should be routed through the sidecar.
	Match Match `json:"match"`
	// Tls provides a TLS configuration when proxy is resposible for a TLS origination
	TLS *TLS `json:"tls,omitempty"`
}

func (o *MeshExternalServiceItemSpec) GetEndpoints() []Endpoints {
	if o == nil {
		return nil
	}
	return o.Endpoints
}

func (o *MeshExternalServiceItemSpec) GetExtension() *Extension {
	if o == nil {
		return nil
	}
	return o.Extension
}

func (o *MeshExternalServiceItemSpec) GetMatch() Match {
	if o == nil {
		return Match{}
	}
	return o.Match
}

func (o *MeshExternalServiceItemSpec) GetTLS() *TLS {
	if o == nil {
		return nil
	}
	return o.TLS
}

type HostnameGeneratorRef struct {
	CoreName string `json:"coreName"`
}

func (o *HostnameGeneratorRef) GetCoreName() string {
	if o == nil {
		return ""
	}
	return o.CoreName
}

type Addresses struct {
	Hostname             *string               `json:"hostname,omitempty"`
	HostnameGeneratorRef *HostnameGeneratorRef `json:"hostnameGeneratorRef,omitempty"`
	Origin               *string               `json:"origin,omitempty"`
}

func (o *Addresses) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *Addresses) GetHostnameGeneratorRef() *HostnameGeneratorRef {
	if o == nil {
		return nil
	}
	return o.HostnameGeneratorRef
}

func (o *Addresses) GetOrigin() *string {
	if o == nil {
		return nil
	}
	return o.Origin
}

// MeshExternalServiceItemStatus - status of the condition, one of True, False, Unknown.
type MeshExternalServiceItemStatus string

const (
	MeshExternalServiceItemStatusTrue    MeshExternalServiceItemStatus = "True"
	MeshExternalServiceItemStatusFalse   MeshExternalServiceItemStatus = "False"
	MeshExternalServiceItemStatusUnknown MeshExternalServiceItemStatus = "Unknown"
)

func (e MeshExternalServiceItemStatus) ToPointer() *MeshExternalServiceItemStatus {
	return &e
}
func (e *MeshExternalServiceItemStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "True":
		fallthrough
	case "False":
		fallthrough
	case "Unknown":
		*e = MeshExternalServiceItemStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshExternalServiceItemStatus: %v", v)
	}
}

type Conditions struct {
	// message is a human readable message indicating details about the transition.
	// This may be an empty string.
	Message string `json:"message"`
	// reason contains a programmatic identifier indicating the reason for the condition's last transition.
	// Producers of specific condition types may define expected values and meanings for this field,
	// and whether the values are considered a guaranteed API.
	// The value should be a CamelCase string.
	// This field may not be empty.
	Reason string `json:"reason"`
	// status of the condition, one of True, False, Unknown.
	Status MeshExternalServiceItemStatus `json:"status"`
	// type of condition in CamelCase or in foo.example.com/CamelCase.
	Type string `json:"type"`
}

func (o *Conditions) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *Conditions) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}

func (o *Conditions) GetStatus() MeshExternalServiceItemStatus {
	if o == nil {
		return MeshExternalServiceItemStatus("")
	}
	return o.Status
}

func (o *Conditions) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

type MeshExternalServiceItemHostnameGeneratorRef struct {
	CoreName string `json:"coreName"`
}

func (o *MeshExternalServiceItemHostnameGeneratorRef) GetCoreName() string {
	if o == nil {
		return ""
	}
	return o.CoreName
}

type HostnameGenerators struct {
	// Conditions is an array of hostname generator conditions.
	Conditions           []Conditions                                `json:"conditions,omitempty"`
	HostnameGeneratorRef MeshExternalServiceItemHostnameGeneratorRef `json:"hostnameGeneratorRef"`
}

func (o *HostnameGenerators) GetConditions() []Conditions {
	if o == nil {
		return nil
	}
	return o.Conditions
}

func (o *HostnameGenerators) GetHostnameGeneratorRef() MeshExternalServiceItemHostnameGeneratorRef {
	if o == nil {
		return MeshExternalServiceItemHostnameGeneratorRef{}
	}
	return o.HostnameGeneratorRef
}

// Vip section for allocated IP
type Vip struct {
	// Value allocated IP for a provided domain with `HostnameGenerator` type in a match section.
	IP *string `json:"ip,omitempty"`
}

func (o *Vip) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}

// Status is the current status of the Kuma MeshExternalService resource.
type Status struct {
	// Addresses section for generated domains
	Addresses          []Addresses          `json:"addresses,omitempty"`
	HostnameGenerators []HostnameGenerators `json:"hostnameGenerators,omitempty"`
	// Vip section for allocated IP
	Vip *Vip `json:"vip,omitempty"`
}

func (o *Status) GetAddresses() []Addresses {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *Status) GetHostnameGenerators() []HostnameGenerators {
	if o == nil {
		return nil
	}
	return o.HostnameGenerators
}

func (o *Status) GetVip() *Vip {
	if o == nil {
		return nil
	}
	return o.Vip
}

// MeshExternalServiceItem - Successful response
type MeshExternalServiceItem struct {
	// the type of the resource
	Type MeshExternalServiceItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshExternalService resource.
	Spec MeshExternalServiceItemSpec `json:"spec"`
	// Time at which the resource was created
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// Time at which the resource was updated
	ModificationTime *time.Time `json:"modificationTime,omitempty"`
	// Status is the current status of the Kuma MeshExternalService resource.
	Status *Status `json:"status,omitempty"`
}

func (m MeshExternalServiceItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshExternalServiceItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshExternalServiceItem) GetType() MeshExternalServiceItemType {
	if o == nil {
		return MeshExternalServiceItemType("")
	}
	return o.Type
}

func (o *MeshExternalServiceItem) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshExternalServiceItem) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshExternalServiceItem) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshExternalServiceItem) GetSpec() MeshExternalServiceItemSpec {
	if o == nil {
		return MeshExternalServiceItemSpec{}
	}
	return o.Spec
}

func (o *MeshExternalServiceItem) GetCreationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *MeshExternalServiceItem) GetModificationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModificationTime
}

func (o *MeshExternalServiceItem) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

type MeshExternalServiceItemInput struct {
	// the type of the resource
	Type MeshExternalServiceItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshExternalService resource.
	Spec MeshExternalServiceItemSpec `json:"spec"`
}

func (m MeshExternalServiceItemInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshExternalServiceItemInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshExternalServiceItemInput) GetType() MeshExternalServiceItemType {
	if o == nil {
		return MeshExternalServiceItemType("")
	}
	return o.Type
}

func (o *MeshExternalServiceItemInput) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshExternalServiceItemInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshExternalServiceItemInput) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshExternalServiceItemInput) GetSpec() MeshExternalServiceItemSpec {
	if o == nil {
		return MeshExternalServiceItemSpec{}
	}
	return o.Spec
}
