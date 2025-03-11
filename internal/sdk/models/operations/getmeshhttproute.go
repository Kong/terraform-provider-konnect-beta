// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
)

type GetMeshHTTPRouteRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshHTTPRoute
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetMeshHTTPRouteRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *GetMeshHTTPRouteRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *GetMeshHTTPRouteRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type GetMeshHTTPRouteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	MeshHTTPRouteItem *shared.MeshHTTPRouteItem
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetMeshHTTPRouteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshHTTPRouteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshHTTPRouteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshHTTPRouteResponse) GetMeshHTTPRouteItem() *shared.MeshHTTPRouteItem {
	if o == nil {
		return nil
	}
	return o.MeshHTTPRouteItem
}

func (o *GetMeshHTTPRouteResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
