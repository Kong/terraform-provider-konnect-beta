// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
)

type UpdateAPISpecRequest struct {
	// The UUID API identifier
	APIID string `pathParam:"style=simple,explode=false,name=apiId"`
	// The API specification identifier
	SpecID  string         `pathParam:"style=simple,explode=false,name=specId"`
	APISpec shared.APISpec `request:"mediaType=application/json"`
}

func (o *UpdateAPISpecRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *UpdateAPISpecRequest) GetSpecID() string {
	if o == nil {
		return ""
	}
	return o.SpecID
}

func (o *UpdateAPISpecRequest) GetAPISpec() shared.APISpec {
	if o == nil {
		return shared.APISpec{}
	}
	return o.APISpec
}

type UpdateAPISpecResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// API specification (OpenAPI or AsyncAPI)
	APISpecResponse *shared.APISpecResponse
	// Bad Request
	BadRequestError *shared.BadRequestError
	// ApiUnauthorized
	UnauthorizedError *shared.UnauthorizedError
	// ApiForbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
	// Conflict - name attribute must be unique across specifications
	ConflictError *shared.ConflictError
	// Unsupported Media Type
	UnsupportedMediaTypeError *shared.UnsupportedMediaTypeError
}

func (o *UpdateAPISpecResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAPISpecResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAPISpecResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateAPISpecResponse) GetAPISpecResponse() *shared.APISpecResponse {
	if o == nil {
		return nil
	}
	return o.APISpecResponse
}

func (o *UpdateAPISpecResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *UpdateAPISpecResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *UpdateAPISpecResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *UpdateAPISpecResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}

func (o *UpdateAPISpecResponse) GetConflictError() *shared.ConflictError {
	if o == nil {
		return nil
	}
	return o.ConflictError
}

func (o *UpdateAPISpecResponse) GetUnsupportedMediaTypeError() *shared.UnsupportedMediaTypeError {
	if o == nil {
		return nil
	}
	return o.UnsupportedMediaTypeError
}
