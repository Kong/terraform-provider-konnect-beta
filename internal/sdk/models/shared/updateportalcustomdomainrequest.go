// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UpdatePortalCustomDomainRequest - Create a portal custom domain.
type UpdatePortalCustomDomainRequest struct {
	Enabled *bool `json:"enabled,omitempty"`
}

func (o *UpdatePortalCustomDomainRequest) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}
