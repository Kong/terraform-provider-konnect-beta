// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
)

type PutMeshRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the Mesh
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	MeshItem shared.MeshItem `request:"mediaType=application/json"`
}

func (o *PutMeshRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *PutMeshRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PutMeshRequest) GetMeshItem() shared.MeshItem {
	if o == nil {
		return shared.MeshItem{}
	}
	return o.MeshItem
}

type PutMeshResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Updated
	MeshCreateOrUpdateSuccessResponse *shared.MeshCreateOrUpdateSuccessResponse
}

func (o *PutMeshResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutMeshResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutMeshResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutMeshResponse) GetMeshCreateOrUpdateSuccessResponse() *shared.MeshCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshCreateOrUpdateSuccessResponse
}
