// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"time"
)

func (r *APIDocumentResourceModel) ToSharedCreateAPIDocumentRequest() *shared.CreateAPIDocumentRequest {
	var content string
	content = r.Content.ValueString()

	title := new(string)
	if !r.Title.IsUnknown() && !r.Title.IsNull() {
		*title = r.Title.ValueString()
	} else {
		title = nil
	}
	slug := new(string)
	if !r.Slug.IsUnknown() && !r.Slug.IsNull() {
		*slug = r.Slug.ValueString()
	} else {
		slug = nil
	}
	status := new(shared.APIDocumentStatus)
	if !r.Status.IsUnknown() && !r.Status.IsNull() {
		*status = shared.APIDocumentStatus(r.Status.ValueString())
	} else {
		status = nil
	}
	parentDocumentID := new(string)
	if !r.ParentDocumentID.IsUnknown() && !r.ParentDocumentID.IsNull() {
		*parentDocumentID = r.ParentDocumentID.ValueString()
	} else {
		parentDocumentID = nil
	}
	labels := make(map[string]*string)
	for labelsKey, labelsValue := range r.Labels {
		labelsInst := new(string)
		if !labelsValue.IsUnknown() && !labelsValue.IsNull() {
			*labelsInst = labelsValue.ValueString()
		} else {
			labelsInst = nil
		}
		labels[labelsKey] = labelsInst
	}
	out := shared.CreateAPIDocumentRequest{
		Content:          content,
		Title:            title,
		Slug:             slug,
		Status:           status,
		ParentDocumentID: parentDocumentID,
		Labels:           labels,
	}
	return &out
}

func (r *APIDocumentResourceModel) RefreshFromSharedAPIDocumentResponse(resp *shared.APIDocumentResponse) {
	if resp != nil {
		r.Content = types.StringValue(resp.Content)
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.ID = types.StringValue(resp.ID)
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String, len(resp.Labels))
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringPointerValue(value)
			}
		}
		r.ParentDocumentID = types.StringPointerValue(resp.ParentDocumentID)
		r.Slug = types.StringValue(resp.Slug)
		if resp.Status != nil {
			r.Status = types.StringValue(string(*resp.Status))
		} else {
			r.Status = types.StringNull()
		}
		r.Title = types.StringValue(resp.Title)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	}
}

func (r *APIDocumentResourceModel) ToSharedAPIDocument() *shared.APIDocument {
	content := new(string)
	if !r.Content.IsUnknown() && !r.Content.IsNull() {
		*content = r.Content.ValueString()
	} else {
		content = nil
	}
	title := new(string)
	if !r.Title.IsUnknown() && !r.Title.IsNull() {
		*title = r.Title.ValueString()
	} else {
		title = nil
	}
	slug := new(string)
	if !r.Slug.IsUnknown() && !r.Slug.IsNull() {
		*slug = r.Slug.ValueString()
	} else {
		slug = nil
	}
	status := new(shared.APIDocumentStatus)
	if !r.Status.IsUnknown() && !r.Status.IsNull() {
		*status = shared.APIDocumentStatus(r.Status.ValueString())
	} else {
		status = nil
	}
	parentDocumentID := new(string)
	if !r.ParentDocumentID.IsUnknown() && !r.ParentDocumentID.IsNull() {
		*parentDocumentID = r.ParentDocumentID.ValueString()
	} else {
		parentDocumentID = nil
	}
	labels := make(map[string]*string)
	for labelsKey, labelsValue := range r.Labels {
		labelsInst := new(string)
		if !labelsValue.IsUnknown() && !labelsValue.IsNull() {
			*labelsInst = labelsValue.ValueString()
		} else {
			labelsInst = nil
		}
		labels[labelsKey] = labelsInst
	}
	out := shared.APIDocument{
		Content:          content,
		Title:            title,
		Slug:             slug,
		Status:           status,
		ParentDocumentID: parentDocumentID,
		Labels:           labels,
	}
	return &out
}
