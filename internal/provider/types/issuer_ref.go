// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type IssuerRef struct {
	Group types.String `tfsdk:"group"`
	Kind  types.String `tfsdk:"kind"`
	Name  types.String `tfsdk:"name"`
}
