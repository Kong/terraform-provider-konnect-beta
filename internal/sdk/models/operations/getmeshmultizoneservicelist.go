// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"net/http"
)

// GetMeshMultiZoneServiceListQueryParamFilter - filter by labels when multiple filters are present, they are ANDed
type GetMeshMultiZoneServiceListQueryParamFilter struct {
	Key   *string `queryParam:"name=key"`
	Value *string `queryParam:"name=value"`
}

func (o *GetMeshMultiZoneServiceListQueryParamFilter) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetMeshMultiZoneServiceListQueryParamFilter) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetMeshMultiZoneServiceListRequest struct {
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
	// offset in the list of entities
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
	// the number of items per page
	Size *int64 `default:"100" queryParam:"style=form,explode=true,name=size"`
	// filter by labels when multiple filters are present, they are ANDed
	Filter *GetMeshMultiZoneServiceListQueryParamFilter `queryParam:"style=form,explode=true,name=filter"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
}

func (g GetMeshMultiZoneServiceListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetMeshMultiZoneServiceListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetMeshMultiZoneServiceListRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

func (o *GetMeshMultiZoneServiceListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetMeshMultiZoneServiceListRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetMeshMultiZoneServiceListRequest) GetFilter() *GetMeshMultiZoneServiceListQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetMeshMultiZoneServiceListRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

type GetMeshMultiZoneServiceListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List
	MeshMultiZoneServiceList *shared.MeshMultiZoneServiceList
}

func (o *GetMeshMultiZoneServiceListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshMultiZoneServiceListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshMultiZoneServiceListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshMultiZoneServiceListResponse) GetMeshMultiZoneServiceList() *shared.MeshMultiZoneServiceList {
	if o == nil {
		return nil
	}
	return o.MeshMultiZoneServiceList
}
