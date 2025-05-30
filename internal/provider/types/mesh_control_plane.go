// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MeshControlPlane struct {
	CreatedAt   types.String              `tfsdk:"created_at"`
	Description types.String              `tfsdk:"description"`
	Features    []MeshControlPlaneFeature `tfsdk:"features"`
	ID          types.String              `tfsdk:"id"`
	Labels      map[string]types.String   `tfsdk:"labels"`
	Name        types.String              `tfsdk:"name"`
	UpdatedAt   types.String              `tfsdk:"updated_at"`
}
