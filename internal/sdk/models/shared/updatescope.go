// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UpdateScope - Scope to be update
type UpdateScope struct {
	// The name of the scope
	Name *string `json:"name,omitempty"`
	// Description of the scope
	Description *string `json:"description,omitempty"`
	// Specifies whether the scope is included by default in access tokens without being explicitly requested by the client. If the scope is not allowed by the client, it will not be included in the access token.
	Default *bool `json:"default,omitempty"`
	// Specifies whether to include the scope in the metadata document
	IncludeInMetadata *bool `json:"include_in_metadata,omitempty"`
	// Specifies whether the scope is enabled
	Enabled *bool `json:"enabled,omitempty"`
}

func (o *UpdateScope) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateScope) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateScope) GetDefault() *bool {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *UpdateScope) GetIncludeInMetadata() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeInMetadata
}

func (o *UpdateScope) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}
