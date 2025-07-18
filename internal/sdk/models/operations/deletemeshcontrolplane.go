// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
)

type DeleteMeshControlPlaneRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
}

func (o *DeleteMeshControlPlaneRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

type DeleteMeshControlPlaneResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Validation Error
	BadRequestError *shared.BadRequestError
	// Unauthorized Error
	UnauthorizedError *shared.UnauthorizedError
	// Permission denied
	ForbiddenError *shared.ForbiddenError
	// Not found
	NotFoundError *shared.NotFoundError
}

func (o *DeleteMeshControlPlaneResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteMeshControlPlaneResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteMeshControlPlaneResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteMeshControlPlaneResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *DeleteMeshControlPlaneResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *DeleteMeshControlPlaneResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *DeleteMeshControlPlaneResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
