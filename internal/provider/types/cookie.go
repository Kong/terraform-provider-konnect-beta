// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Cookie struct {
	Name types.String `tfsdk:"name"`
	Path types.String `tfsdk:"path"`
	TTL  types.String `tfsdk:"ttl"`
}
