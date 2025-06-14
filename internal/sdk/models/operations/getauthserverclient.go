// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
)

type GetAuthServerClientRequest struct {
	// The auth server ID
	AuthServerID string `pathParam:"style=simple,explode=false,name=authServerId"`
	// The OAuth 2.0 client ID
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
}

func (o *GetAuthServerClientRequest) GetAuthServerID() string {
	if o == nil {
		return ""
	}
	return o.AuthServerID
}

func (o *GetAuthServerClientRequest) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

type GetAuthServerClientResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A client
	Client *shared.Client
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetAuthServerClientResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAuthServerClientResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAuthServerClientResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAuthServerClientResponse) GetClient() *shared.Client {
	if o == nil {
		return nil
	}
	return o.Client
}

func (o *GetAuthServerClientResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
