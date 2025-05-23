// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MeshGatewayItem struct {
	Conf      *Conf                   `tfsdk:"conf"`
	Labels    map[string]types.String `tfsdk:"labels"`
	Mesh      types.String            `tfsdk:"mesh"`
	Name      types.String            `tfsdk:"name"`
	Selectors []Selectors             `tfsdk:"selectors"`
	Tags      map[string]types.String `tfsdk:"tags"`
	Type      types.String            `tfsdk:"type"`
}
