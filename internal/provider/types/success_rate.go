// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SuccessRate struct {
	MinimumHosts            types.Int32   `tfsdk:"minimum_hosts"`
	RequestVolume           types.Int32   `tfsdk:"request_volume"`
	StandardDeviationFactor *MeshItemMode `tfsdk:"standard_deviation_factor"`
}
