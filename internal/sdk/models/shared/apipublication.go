// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/internal/utils"
)

// APIPublication - An API publication in a portal
type APIPublication struct {
	// Whether the application registration auto approval on this portal for the api is enabled. If set to false, fallbacks on portal's auto_approve_applications value.
	AutoApproveRegistrations *bool `json:"auto_approve_registrations,omitempty"`
	// The auth strategy the API enforces for applications in the portal.
	// Omitting this property means the portal's default application auth strategy will be used.
	// Setting to null means the API will not require application authentication.
	// DCR support for application registration is currently in development.
	//
	AuthStrategyIds []string `json:"auth_strategy_ids,omitempty"`
	// The visibility of the API in the portal.
	// Public API publications do not require authentication to view and retrieve information about them.
	// Private API publications require authentication to retrieve information about them.
	//
	Visibility *APIPublicationVisibility `default:"private" json:"visibility"`
}

func (a APIPublication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *APIPublication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *APIPublication) GetAutoApproveRegistrations() *bool {
	if o == nil {
		return nil
	}
	return o.AutoApproveRegistrations
}

func (o *APIPublication) GetAuthStrategyIds() []string {
	if o == nil {
		return nil
	}
	return o.AuthStrategyIds
}

func (o *APIPublication) GetVisibility() *APIPublicationVisibility {
	if o == nil {
		return nil
	}
	return o.Visibility
}
