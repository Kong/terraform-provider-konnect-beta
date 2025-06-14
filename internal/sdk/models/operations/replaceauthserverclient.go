// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
)

type ReplaceAuthServerClientRequest struct {
	// The auth server ID
	AuthServerID string `pathParam:"style=simple,explode=false,name=authServerId"`
	// The OAuth 2.0 client ID
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
	// Client to be replaced
	ReplaceClient shared.ReplaceClient `request:"mediaType=application/json"`
}

func (o *ReplaceAuthServerClientRequest) GetAuthServerID() string {
	if o == nil {
		return ""
	}
	return o.AuthServerID
}

func (o *ReplaceAuthServerClientRequest) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *ReplaceAuthServerClientRequest) GetReplaceClient() shared.ReplaceClient {
	if o == nil {
		return shared.ReplaceClient{}
	}
	return o.ReplaceClient
}

type ReplaceAuthServerClientResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A client
	Client *shared.Client
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Not Found
	NotFoundError *shared.NotFoundError
	// Conflict
	ConflictError *shared.ConflictError
}

func (o *ReplaceAuthServerClientResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReplaceAuthServerClientResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReplaceAuthServerClientResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ReplaceAuthServerClientResponse) GetClient() *shared.Client {
	if o == nil {
		return nil
	}
	return o.Client
}

func (o *ReplaceAuthServerClientResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *ReplaceAuthServerClientResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}

func (o *ReplaceAuthServerClientResponse) GetConflictError() *shared.ConflictError {
	if o == nil {
		return nil
	}
	return o.ConflictError
}
