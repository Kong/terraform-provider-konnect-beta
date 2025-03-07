// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
)

type GetMeshTimeoutRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshTimeout
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetMeshTimeoutRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *GetMeshTimeoutRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *GetMeshTimeoutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type GetMeshTimeoutResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	MeshTimeoutItem *shared.MeshTimeoutItem
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetMeshTimeoutResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshTimeoutResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshTimeoutResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshTimeoutResponse) GetMeshTimeoutItem() *shared.MeshTimeoutItem {
	if o == nil {
		return nil
	}
	return o.MeshTimeoutItem
}

func (o *GetMeshTimeoutResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
