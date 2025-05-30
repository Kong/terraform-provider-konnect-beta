// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
)

type GetPortalCustomizationRequest struct {
	// The Portal identifier
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
}

func (o *GetPortalCustomizationRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

type GetPortalCustomizationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The current customization options for a portal.
	PortalCustomization *shared.PortalCustomization
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetPortalCustomizationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPortalCustomizationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPortalCustomizationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPortalCustomizationResponse) GetPortalCustomization() *shared.PortalCustomization {
	if o == nil {
		return nil
	}
	return o.PortalCustomization
}

func (o *GetPortalCustomizationResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *GetPortalCustomizationResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *GetPortalCustomizationResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
