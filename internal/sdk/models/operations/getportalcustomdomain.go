// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
)

type GetPortalCustomDomainRequest struct {
	// The Portal identifier
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
}

func (o *GetPortalCustomDomainRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

type GetPortalCustomDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Portal custom domain
	PortalCustomDomain *shared.PortalCustomDomain
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetPortalCustomDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPortalCustomDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPortalCustomDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPortalCustomDomainResponse) GetPortalCustomDomain() *shared.PortalCustomDomain {
	if o == nil {
		return nil
	}
	return o.PortalCustomDomain
}

func (o *GetPortalCustomDomainResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *GetPortalCustomDomainResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *GetPortalCustomDomainResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
