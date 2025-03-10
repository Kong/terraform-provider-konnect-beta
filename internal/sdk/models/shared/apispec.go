// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type APISpec struct {
	// The raw content of your API specification.
	//
	Content *string `json:"content,omitempty"`
	// The type of specification being stored. This allows us to render the specification correctly.
	//
	// If this field is not set, it will be autodetected from `content`
	//
	Type *APISpecType `json:"type,omitempty"`
}

func (o *APISpec) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *APISpec) GetType() *APISpecType {
	if o == nil {
		return nil
	}
	return o.Type
}
