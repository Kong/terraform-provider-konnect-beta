// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type RequestBody struct {
	MaxSize     types.Int32 `tfsdk:"max_size"`
	SendRawBody types.Bool  `tfsdk:"send_raw_body"`
}
