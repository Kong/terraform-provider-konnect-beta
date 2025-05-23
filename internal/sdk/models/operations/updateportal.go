// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
)

type UpdatePortalRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// Update a portal's settings.
	UpdatePortal shared.UpdatePortal `request:"mediaType=application/json"`
}

func (o *UpdatePortalRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *UpdatePortalRequest) GetUpdatePortal() shared.UpdatePortal {
	if o == nil {
		return shared.UpdatePortal{}
	}
	return o.UpdatePortal
}

type UpdatePortalResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Details about a portal.
	PortalResponse *shared.PortalResponse
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
	// Conflict
	ConflictError *shared.ConflictError
}

func (o *UpdatePortalResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePortalResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePortalResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdatePortalResponse) GetPortalResponse() *shared.PortalResponse {
	if o == nil {
		return nil
	}
	return o.PortalResponse
}

func (o *UpdatePortalResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *UpdatePortalResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *UpdatePortalResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *UpdatePortalResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}

func (o *UpdatePortalResponse) GetConflictError() *shared.ConflictError {
	if o == nil {
		return nil
	}
	return o.ConflictError
}
