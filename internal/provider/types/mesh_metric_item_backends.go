// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MeshMetricItemBackends struct {
	OpenTelemetry *OpenTelemetry `tfsdk:"open_telemetry"`
	Prometheus    *Prometheus    `tfsdk:"prometheus"`
	Type          types.String   `tfsdk:"type"`
}
