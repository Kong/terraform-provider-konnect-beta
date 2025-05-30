// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PortalOIDCConfig struct {
	ClaimMappings *PortalClaimMappings `tfsdk:"claim_mappings"`
	ClientID      types.String         `tfsdk:"client_id"`
	Issuer        types.String         `tfsdk:"issuer"`
	Scopes        []types.String       `tfsdk:"scopes"`
}
