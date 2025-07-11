// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
)

type ListAuthServerClientsRequest struct {
	// The auth server ID
	AuthServerID string `pathParam:"style=simple,explode=false,name=authServerId"`
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
}

func (o *ListAuthServerClientsRequest) GetAuthServerID() string {
	if o == nil {
		return ""
	}
	return o.AuthServerID
}

func (o *ListAuthServerClientsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListAuthServerClientsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

type ListAuthServerClientsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List of clients
	ClientList *shared.ClientList
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *ListAuthServerClientsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAuthServerClientsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAuthServerClientsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListAuthServerClientsResponse) GetClientList() *shared.ClientList {
	if o == nil {
		return nil
	}
	return o.ClientList
}

func (o *ListAuthServerClientsResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
