// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HostSelection struct {
	Predicate       types.String            `tfsdk:"predicate"`
	Tags            map[string]types.String `tfsdk:"tags"`
	UpdateFrequency types.Int32             `tfsdk:"update_frequency"`
}
