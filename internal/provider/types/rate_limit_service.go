// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RateLimitService struct {
	LimitOnServiceFail types.Bool   `tfsdk:"limit_on_service_fail"`
	Timeout            types.String `tfsdk:"timeout"`
	URL                types.String `tfsdk:"url"`
}
