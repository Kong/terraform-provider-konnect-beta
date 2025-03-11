// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
)

type UpdateMeshRateLimitRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshRateLimit
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	MeshRateLimitItem shared.MeshRateLimitItemInput `request:"mediaType=application/json"`
}

func (o *UpdateMeshRateLimitRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *UpdateMeshRateLimitRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *UpdateMeshRateLimitRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateMeshRateLimitRequest) GetMeshRateLimitItem() shared.MeshRateLimitItemInput {
	if o == nil {
		return shared.MeshRateLimitItemInput{}
	}
	return o.MeshRateLimitItem
}

type UpdateMeshRateLimitResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Updated
	MeshRateLimitCreateOrUpdateSuccessResponse *shared.MeshRateLimitCreateOrUpdateSuccessResponse
}

func (o *UpdateMeshRateLimitResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateMeshRateLimitResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateMeshRateLimitResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateMeshRateLimitResponse) GetMeshRateLimitCreateOrUpdateSuccessResponse() *shared.MeshRateLimitCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshRateLimitCreateOrUpdateSuccessResponse
}
