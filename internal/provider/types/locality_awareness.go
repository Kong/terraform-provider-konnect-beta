// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LocalityAwareness struct {
	CrossZone *CrossZone `tfsdk:"cross_zone"`
	Disabled  types.Bool `tfsdk:"disabled"`
	LocalZone *LocalZone `tfsdk:"local_zone"`
}
