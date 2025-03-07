// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/internal/utils"
	"time"
)

// MeshTraceItemType - the type of the resource
type MeshTraceItemType string

const (
	MeshTraceItemTypeMeshTrace MeshTraceItemType = "MeshTrace"
)

func (e MeshTraceItemType) ToPointer() *MeshTraceItemType {
	return &e
}
func (e *MeshTraceItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MeshTrace":
		*e = MeshTraceItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTraceItemType: %v", v)
	}
}

// Datadog backend configuration.
type Datadog struct {
	// Determines if datadog service name should be split based on traffic
	// direction and destination. For example, with `splitService: true` and a
	// `backend` service that communicates with a couple of databases, you would
	// get service names like `backend_INBOUND`, `backend_OUTBOUND_db1`, and
	// `backend_OUTBOUND_db2` in Datadog.
	SplitService *bool `default:"false" json:"splitService"`
	// Address of Datadog collector, only host and port are allowed (no paths,
	// fragments etc.)
	URL string `json:"url"`
}

func (d Datadog) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Datadog) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Datadog) GetSplitService() *bool {
	if o == nil {
		return nil
	}
	return o.SplitService
}

func (o *Datadog) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

// MeshTraceItemOpenTelemetry - OpenTelemetry backend configuration.
type MeshTraceItemOpenTelemetry struct {
	// Address of OpenTelemetry collector.
	Endpoint string `json:"endpoint"`
}

func (o *MeshTraceItemOpenTelemetry) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

type MeshTraceItemSpecType string

const (
	MeshTraceItemSpecTypeZipkin        MeshTraceItemSpecType = "Zipkin"
	MeshTraceItemSpecTypeDatadog       MeshTraceItemSpecType = "Datadog"
	MeshTraceItemSpecTypeOpenTelemetry MeshTraceItemSpecType = "OpenTelemetry"
)

func (e MeshTraceItemSpecType) ToPointer() *MeshTraceItemSpecType {
	return &e
}
func (e *MeshTraceItemSpecType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Zipkin":
		fallthrough
	case "Datadog":
		fallthrough
	case "OpenTelemetry":
		*e = MeshTraceItemSpecType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTraceItemSpecType: %v", v)
	}
}

// APIVersion - Version of the API.
// https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L66
type APIVersion string

const (
	APIVersionHTTPJSON  APIVersion = "httpJson"
	APIVersionHTTPProto APIVersion = "httpProto"
)

func (e APIVersion) ToPointer() *APIVersion {
	return &e
}
func (e *APIVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "httpJson":
		fallthrough
	case "httpProto":
		*e = APIVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APIVersion: %v", v)
	}
}

// Zipkin backend configuration.
type Zipkin struct {
	// Version of the API.
	// https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L66
	APIVersion *APIVersion `default:"httpJson" json:"apiVersion"`
	// Determines whether client and server spans will share the same span
	// context.
	// https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L63
	SharedSpanContext *bool `default:"true" json:"sharedSpanContext"`
	// Generate 128bit traces.
	TraceId128bit *bool `default:"false" json:"traceId128bit"`
	// Address of Zipkin collector.
	URL string `json:"url"`
}

func (z Zipkin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(z, "", false)
}

func (z *Zipkin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &z, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Zipkin) GetAPIVersion() *APIVersion {
	if o == nil {
		return nil
	}
	return o.APIVersion
}

func (o *Zipkin) GetSharedSpanContext() *bool {
	if o == nil {
		return nil
	}
	return o.SharedSpanContext
}

func (o *Zipkin) GetTraceId128bit() *bool {
	if o == nil {
		return nil
	}
	return o.TraceId128bit
}

func (o *Zipkin) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

// MeshTraceItemBackends - Only one of zipkin, datadog or openTelemetry can be used.
type MeshTraceItemBackends struct {
	// Datadog backend configuration.
	Datadog *Datadog `json:"datadog,omitempty"`
	// OpenTelemetry backend configuration.
	OpenTelemetry *MeshTraceItemOpenTelemetry `json:"openTelemetry,omitempty"`
	Type          MeshTraceItemSpecType       `json:"type"`
	// Zipkin backend configuration.
	Zipkin *Zipkin `json:"zipkin,omitempty"`
}

func (o *MeshTraceItemBackends) GetDatadog() *Datadog {
	if o == nil {
		return nil
	}
	return o.Datadog
}

func (o *MeshTraceItemBackends) GetOpenTelemetry() *MeshTraceItemOpenTelemetry {
	if o == nil {
		return nil
	}
	return o.OpenTelemetry
}

func (o *MeshTraceItemBackends) GetType() MeshTraceItemSpecType {
	if o == nil {
		return MeshTraceItemSpecType("")
	}
	return o.Type
}

func (o *MeshTraceItemBackends) GetZipkin() *Zipkin {
	if o == nil {
		return nil
	}
	return o.Zipkin
}

type ClientType string

const (
	ClientTypeInteger ClientType = "integer"
	ClientTypeStr     ClientType = "str"
)

// Client - Target percentage of requests that will be force traced if the
// 'x-client-trace-id' header is set. Mirror of client_sampling in Envoy
// https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L127-L133
// Either int or decimal represented as string.
// If not specified then the default value is 100.
type Client struct {
	Integer *int64  `queryParam:"inline"`
	Str     *string `queryParam:"inline"`

	Type ClientType
}

func CreateClientInteger(integer int64) Client {
	typ := ClientTypeInteger

	return Client{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateClientStr(str string) Client {
	typ := ClientTypeStr

	return Client{
		Str:  &str,
		Type: typ,
	}
}

func (u *Client) UnmarshalJSON(data []byte) error {

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = ClientTypeInteger
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ClientTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Client", string(data))
}

func (u Client) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type Client: all fields are null")
}

type OverallType string

const (
	OverallTypeInteger OverallType = "integer"
	OverallTypeStr     OverallType = "str"
)

// Overall - Target percentage of requests will be traced
// after all other sampling checks have been applied (client, force tracing,
// random sampling). This field functions as an upper limit on the total
// configured sampling rate. For instance, setting client to 100
// but overall to 1 will result in only 1% of client requests with
// the appropriate headers to be force traced. Mirror of
// overall_sampling in Envoy
// https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L142-L150
// Either int or decimal represented as string.
// If not specified then the default value is 100.
type Overall struct {
	Integer *int64  `queryParam:"inline"`
	Str     *string `queryParam:"inline"`

	Type OverallType
}

func CreateOverallInteger(integer int64) Overall {
	typ := OverallTypeInteger

	return Overall{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateOverallStr(str string) Overall {
	typ := OverallTypeStr

	return Overall{
		Str:  &str,
		Type: typ,
	}
}

func (u *Overall) UnmarshalJSON(data []byte) error {

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = OverallTypeInteger
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = OverallTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Overall", string(data))
}

func (u Overall) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type Overall: all fields are null")
}

type RandomType string

const (
	RandomTypeInteger RandomType = "integer"
	RandomTypeStr     RandomType = "str"
)

// Random - Target percentage of requests that will be randomly selected for trace
// generation, if not requested by the client or not forced.
// Mirror of random_sampling in Envoy
// https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L135-L140
// Either int or decimal represented as string.
// If not specified then the default value is 100.
type Random struct {
	Integer *int64  `queryParam:"inline"`
	Str     *string `queryParam:"inline"`

	Type RandomType
}

func CreateRandomInteger(integer int64) Random {
	typ := RandomTypeInteger

	return Random{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateRandomStr(str string) Random {
	typ := RandomTypeStr

	return Random{
		Str:  &str,
		Type: typ,
	}
}

func (u *Random) UnmarshalJSON(data []byte) error {

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = RandomTypeInteger
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = RandomTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Random", string(data))
}

func (u Random) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type Random: all fields are null")
}

// Sampling configuration.
// Sampling is the process by which a decision is made on whether to
// process/export a span or not.
type Sampling struct {
	// Target percentage of requests that will be force traced if the
	// 'x-client-trace-id' header is set. Mirror of client_sampling in Envoy
	// https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L127-L133
	// Either int or decimal represented as string.
	// If not specified then the default value is 100.
	Client *Client `json:"client,omitempty"`
	// Target percentage of requests will be traced
	// after all other sampling checks have been applied (client, force tracing,
	// random sampling). This field functions as an upper limit on the total
	// configured sampling rate. For instance, setting client to 100
	// but overall to 1 will result in only 1% of client requests with
	// the appropriate headers to be force traced. Mirror of
	// overall_sampling in Envoy
	// https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L142-L150
	// Either int or decimal represented as string.
	// If not specified then the default value is 100.
	Overall *Overall `json:"overall,omitempty"`
	// Target percentage of requests that will be randomly selected for trace
	// generation, if not requested by the client or not forced.
	// Mirror of random_sampling in Envoy
	// https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L135-L140
	// Either int or decimal represented as string.
	// If not specified then the default value is 100.
	Random *Random `json:"random,omitempty"`
}

func (o *Sampling) GetClient() *Client {
	if o == nil {
		return nil
	}
	return o.Client
}

func (o *Sampling) GetOverall() *Overall {
	if o == nil {
		return nil
	}
	return o.Overall
}

func (o *Sampling) GetRandom() *Random {
	if o == nil {
		return nil
	}
	return o.Random
}

// Header - Tag taken from a header.
type Header struct {
	// Default value to use if header is missing.
	// If the default is missing and there is no value the tag will not be
	// included.
	Default *string `json:"default,omitempty"`
	// Name of the header.
	Name string `json:"name"`
}

func (o *Header) GetDefault() *string {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *Header) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// Tags - Custom tags configuration.
// Only one of literal or header can be used.
type Tags struct {
	// Tag taken from a header.
	Header *Header `json:"header,omitempty"`
	// Tag taken from literal value.
	Literal *string `json:"literal,omitempty"`
	// Name of the tag.
	Name string `json:"name"`
}

func (o *Tags) GetHeader() *Header {
	if o == nil {
		return nil
	}
	return o.Header
}

func (o *Tags) GetLiteral() *string {
	if o == nil {
		return nil
	}
	return o.Literal
}

func (o *Tags) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// MeshTraceItemDefault - MeshTrace configuration.
type MeshTraceItemDefault struct {
	// A one element array of backend definition.
	// Envoy allows configuring only 1 backend, so the natural way of
	// representing that would be just one object. Unfortunately due to the
	// reasons explained in MADR 009-tracing-policy this has to be a one element
	// array for now.
	Backends []MeshTraceItemBackends `json:"backends,omitempty"`
	// Sampling configuration.
	// Sampling is the process by which a decision is made on whether to
	// process/export a span or not.
	Sampling *Sampling `json:"sampling,omitempty"`
	// Custom tags configuration. You can add custom tags to traces based on
	// headers or literal values.
	Tags []Tags `json:"tags,omitempty"`
}

func (o *MeshTraceItemDefault) GetBackends() []MeshTraceItemBackends {
	if o == nil {
		return nil
	}
	return o.Backends
}

func (o *MeshTraceItemDefault) GetSampling() *Sampling {
	if o == nil {
		return nil
	}
	return o.Sampling
}

func (o *MeshTraceItemDefault) GetTags() []Tags {
	if o == nil {
		return nil
	}
	return o.Tags
}

// MeshTraceItemKind - Kind of the referenced resource
type MeshTraceItemKind string

const (
	MeshTraceItemKindMesh                 MeshTraceItemKind = "Mesh"
	MeshTraceItemKindMeshSubset           MeshTraceItemKind = "MeshSubset"
	MeshTraceItemKindMeshGateway          MeshTraceItemKind = "MeshGateway"
	MeshTraceItemKindMeshService          MeshTraceItemKind = "MeshService"
	MeshTraceItemKindMeshExternalService  MeshTraceItemKind = "MeshExternalService"
	MeshTraceItemKindMeshMultiZoneService MeshTraceItemKind = "MeshMultiZoneService"
	MeshTraceItemKindMeshServiceSubset    MeshTraceItemKind = "MeshServiceSubset"
	MeshTraceItemKindMeshHTTPRoute        MeshTraceItemKind = "MeshHTTPRoute"
	MeshTraceItemKindDataplane            MeshTraceItemKind = "Dataplane"
)

func (e MeshTraceItemKind) ToPointer() *MeshTraceItemKind {
	return &e
}
func (e *MeshTraceItemKind) UnmarshalJSON(data []byte) error {
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
		*e = MeshTraceItemKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTraceItemKind: %v", v)
	}
}

type MeshTraceItemProxyTypes string

const (
	MeshTraceItemProxyTypesSidecar MeshTraceItemProxyTypes = "Sidecar"
	MeshTraceItemProxyTypesGateway MeshTraceItemProxyTypes = "Gateway"
)

func (e MeshTraceItemProxyTypes) ToPointer() *MeshTraceItemProxyTypes {
	return &e
}
func (e *MeshTraceItemProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshTraceItemProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshTraceItemProxyTypes: %v", v)
	}
}

// MeshTraceItemTargetRef - TargetRef is a reference to the resource the policy takes an effect on.
// The resource could be either a real store object or virtual resource
// defined inplace.
type MeshTraceItemTargetRef struct {
	// Kind of the referenced resource
	Kind MeshTraceItemKind `json:"kind"`
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
	ProxyTypes []MeshTraceItemProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshTraceItemTargetRef) GetKind() MeshTraceItemKind {
	if o == nil {
		return MeshTraceItemKind("")
	}
	return o.Kind
}

func (o *MeshTraceItemTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTraceItemTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTraceItemTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshTraceItemTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshTraceItemTargetRef) GetProxyTypes() []MeshTraceItemProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshTraceItemTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshTraceItemTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// MeshTraceItemSpec - Spec is the specification of the Kuma MeshTrace resource.
type MeshTraceItemSpec struct {
	// MeshTrace configuration.
	Default *MeshTraceItemDefault `json:"default,omitempty"`
	// TargetRef is a reference to the resource the policy takes an effect on.
	// The resource could be either a real store object or virtual resource
	// defined inplace.
	TargetRef *MeshTraceItemTargetRef `json:"targetRef,omitempty"`
}

func (o *MeshTraceItemSpec) GetDefault() *MeshTraceItemDefault {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *MeshTraceItemSpec) GetTargetRef() *MeshTraceItemTargetRef {
	if o == nil {
		return nil
	}
	return o.TargetRef
}

// MeshTraceItem - Successful response
type MeshTraceItem struct {
	// the type of the resource
	Type MeshTraceItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshTrace resource.
	Spec MeshTraceItemSpec `json:"spec"`
	// Time at which the resource was created
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// Time at which the resource was updated
	ModificationTime *time.Time `json:"modificationTime,omitempty"`
}

func (m MeshTraceItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshTraceItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshTraceItem) GetType() MeshTraceItemType {
	if o == nil {
		return MeshTraceItemType("")
	}
	return o.Type
}

func (o *MeshTraceItem) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTraceItem) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshTraceItem) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTraceItem) GetSpec() MeshTraceItemSpec {
	if o == nil {
		return MeshTraceItemSpec{}
	}
	return o.Spec
}

func (o *MeshTraceItem) GetCreationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *MeshTraceItem) GetModificationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModificationTime
}

type MeshTraceItemInput struct {
	// the type of the resource
	Type MeshTraceItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshTrace resource.
	Spec MeshTraceItemSpec `json:"spec"`
}

func (m MeshTraceItemInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshTraceItemInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshTraceItemInput) GetType() MeshTraceItemType {
	if o == nil {
		return MeshTraceItemType("")
	}
	return o.Type
}

func (o *MeshTraceItemInput) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshTraceItemInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshTraceItemInput) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshTraceItemInput) GetSpec() MeshTraceItemSpec {
	if o == nil {
		return MeshTraceItemSpec{}
	}
	return o.Spec
}
