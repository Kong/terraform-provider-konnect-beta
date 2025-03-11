// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type PrometheusMetricsBackendConfig struct {
	Aggregate []Aggregate             `tfsdk:"aggregate"`
	Envoy     *Envoy                  `tfsdk:"envoy"`
	Path      types.String            `tfsdk:"path"`
	Port      types.Int64             `tfsdk:"port"`
	SkipMTLS  types.Bool              `tfsdk:"skip_mtls"`
	Tags      map[string]types.String `tfsdk:"tags"`
	TLS       *MeshServices           `tfsdk:"tls"`
}
