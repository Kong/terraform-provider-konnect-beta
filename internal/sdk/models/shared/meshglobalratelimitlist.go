// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// MeshGlobalRateLimitList - List
type MeshGlobalRateLimitList struct {
	Items []MeshGlobalRateLimitItem `json:"items,omitempty"`
	// The total number of entities
	Total *float64 `json:"total,omitempty"`
	// URL to the next page
	Next *string `json:"next,omitempty"`
}

func (o *MeshGlobalRateLimitList) GetItems() []MeshGlobalRateLimitItem {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *MeshGlobalRateLimitList) GetTotal() *float64 {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *MeshGlobalRateLimitList) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}
