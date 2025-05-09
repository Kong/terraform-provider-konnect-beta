// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/provider/typeconvert"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
)

func (r *PortalCustomDomainResourceModel) ToSharedCreatePortalCustomDomainRequest(ctx context.Context) (*shared.CreatePortalCustomDomainRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var hostname string
	hostname = r.Hostname.ValueString()

	var enabled bool
	enabled = r.Enabled.ValueBool()

	domainVerificationMethod := shared.PortalCustomDomainVerificationMethod(r.Ssl.DomainVerificationMethod.ValueString())
	customCertificate := new(string)
	if !r.Ssl.CustomCertificate.IsUnknown() && !r.Ssl.CustomCertificate.IsNull() {
		*customCertificate = r.Ssl.CustomCertificate.ValueString()
	} else {
		customCertificate = nil
	}
	customPrivateKey := new(string)
	if !r.Ssl.CustomPrivateKey.IsUnknown() && !r.Ssl.CustomPrivateKey.IsNull() {
		*customPrivateKey = r.Ssl.CustomPrivateKey.ValueString()
	} else {
		customPrivateKey = nil
	}
	ssl := shared.CreatePortalCustomDomainSSL{
		DomainVerificationMethod: domainVerificationMethod,
		CustomCertificate:        customCertificate,
		CustomPrivateKey:         customPrivateKey,
	}
	out := shared.CreatePortalCustomDomainRequest{
		Hostname: hostname,
		Enabled:  enabled,
		Ssl:      ssl,
	}

	return &out, diags
}

func (r *PortalCustomDomainResourceModel) ToOperationsCreatePortalCustomDomainRequest(ctx context.Context) (*operations.CreatePortalCustomDomainRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var portalID string
	portalID = r.PortalID.ValueString()

	createPortalCustomDomainRequest, createPortalCustomDomainRequestDiags := r.ToSharedCreatePortalCustomDomainRequest(ctx)
	diags.Append(createPortalCustomDomainRequestDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.CreatePortalCustomDomainRequest{
		PortalID:                        portalID,
		CreatePortalCustomDomainRequest: *createPortalCustomDomainRequest,
	}

	return &out, diags
}

func (r *PortalCustomDomainResourceModel) ToSharedUpdatePortalCustomDomainRequest(ctx context.Context) (*shared.UpdatePortalCustomDomainRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	out := shared.UpdatePortalCustomDomainRequest{
		Enabled: enabled,
	}

	return &out, diags
}

func (r *PortalCustomDomainResourceModel) ToOperationsUpdatePortalCustomDomainRequest(ctx context.Context) (*operations.UpdatePortalCustomDomainRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var portalID string
	portalID = r.PortalID.ValueString()

	updatePortalCustomDomainRequest, updatePortalCustomDomainRequestDiags := r.ToSharedUpdatePortalCustomDomainRequest(ctx)
	diags.Append(updatePortalCustomDomainRequestDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.UpdatePortalCustomDomainRequest{
		PortalID:                        portalID,
		UpdatePortalCustomDomainRequest: *updatePortalCustomDomainRequest,
	}

	return &out, diags
}

func (r *PortalCustomDomainResourceModel) ToOperationsGetPortalCustomDomainRequest(ctx context.Context) (*operations.GetPortalCustomDomainRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var portalID string
	portalID = r.PortalID.ValueString()

	out := operations.GetPortalCustomDomainRequest{
		PortalID: portalID,
	}

	return &out, diags
}

func (r *PortalCustomDomainResourceModel) ToOperationsDeletePortalCustomDomainRequest(ctx context.Context) (*operations.DeletePortalCustomDomainRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var portalID string
	portalID = r.PortalID.ValueString()

	out := operations.DeletePortalCustomDomainRequest{
		PortalID: portalID,
	}

	return &out, diags
}

func (r *PortalCustomDomainResourceModel) RefreshFromSharedPortalCustomDomain(ctx context.Context, resp *shared.PortalCustomDomain) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CnameStatus = types.StringValue(string(resp.CnameStatus))
		r.CreatedAt = types.StringValue(typeconvert.TimeToString(resp.CreatedAt))
		r.Enabled = types.BoolValue(resp.Enabled)
		r.Hostname = types.StringValue(resp.Hostname)
		r.Ssl.DomainVerificationMethod = types.StringValue(string(resp.Ssl.DomainVerificationMethod))
		r.Ssl.ExpiresAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.Ssl.ExpiresAt))
		r.Ssl.UploadedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.Ssl.UploadedAt))
		r.Ssl.ValidationErrors = make([]types.String, 0, len(resp.Ssl.ValidationErrors))
		for _, v := range resp.Ssl.ValidationErrors {
			r.Ssl.ValidationErrors = append(r.Ssl.ValidationErrors, types.StringValue(v))
		}
		r.Ssl.VerificationStatus = types.StringValue(string(resp.Ssl.VerificationStatus))
		r.UpdatedAt = types.StringValue(typeconvert.TimeToString(resp.UpdatedAt))
	}

	return diags
}
