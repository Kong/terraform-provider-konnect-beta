// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
)

type UpdatePortalPageRequest struct {
	// The Portal identifier
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// ID of the page.
	PageID string `pathParam:"style=simple,explode=false,name=pageId"`
	// Update a page in a portal.
	UpdatePortalPageRequest shared.UpdatePortalPageRequest `request:"mediaType=application/json"`
}

func (o *UpdatePortalPageRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *UpdatePortalPageRequest) GetPageID() string {
	if o == nil {
		return ""
	}
	return o.PageID
}

func (o *UpdatePortalPageRequest) GetUpdatePortalPageRequest() shared.UpdatePortalPageRequest {
	if o == nil {
		return shared.UpdatePortalPageRequest{}
	}
	return o.UpdatePortalPageRequest
}

type UpdatePortalPageResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Details about a page in a portal.
	PortalPageResponse *shared.PortalPageResponse
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *UpdatePortalPageResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePortalPageResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePortalPageResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdatePortalPageResponse) GetPortalPageResponse() *shared.PortalPageResponse {
	if o == nil {
		return nil
	}
	return o.PortalPageResponse
}

func (o *UpdatePortalPageResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *UpdatePortalPageResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *UpdatePortalPageResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *UpdatePortalPageResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
