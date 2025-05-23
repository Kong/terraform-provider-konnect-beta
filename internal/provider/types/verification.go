// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Verification struct {
	CaCert          *CaCert           `tfsdk:"ca_cert"`
	ClientCert      *CaCert           `tfsdk:"client_cert"`
	ClientKey       *CaCert           `tfsdk:"client_key"`
	Mode            types.String      `tfsdk:"mode"`
	ServerName      types.String      `tfsdk:"server_name"`
	SubjectAltNames []SubjectAltNames `tfsdk:"subject_alt_names"`
}
