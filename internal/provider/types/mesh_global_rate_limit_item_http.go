// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type MeshGlobalRateLimitItemHTTP struct {
	Disabled           types.Bool                                  `tfsdk:"disabled"`
	OnRateLimit        *OnRateLimit                                `tfsdk:"on_rate_limit"`
	RatelimitOnRequest []RatelimitOnRequest                        `tfsdk:"ratelimit_on_request"`
	RequestRate        *MeshGlobalRateLimitItemSpecFromRequestRate `tfsdk:"request_rate"`
}
