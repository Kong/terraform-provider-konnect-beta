// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// MeshServiceList - List
type MeshServiceList struct {
	Items []MeshServiceItem `json:"items,omitempty"`
	// The total number of entities
	Total *float64 `json:"total,omitempty"`
	// URL to the next page
	Next *string `json:"next,omitempty"`
}

func (o *MeshServiceList) GetItems() []MeshServiceItem {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *MeshServiceList) GetTotal() *float64 {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *MeshServiceList) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}
