// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/internal/utils"
)

// CreatePortalPageRequest - Create a page in a portal.
type CreatePortalPageRequest struct {
	// The slug of a page in a portal. Is used to compute the full path /slug1/slug2/slug3.
	Slug string `json:"slug"`
	// The title of a page in a portal.
	Title *string `json:"title,omitempty"`
	// The renderable markdown content of a page in a portal.
	Content string `json:"content"`
	// Whether a page is publicly accessible to non-authenticated users.
	Visibility *string `default:"private" json:"visibility"`
	// Whether the resource is visible on a given portal. Defaults to false.
	Status      *PublishedStatus `json:"status,omitempty"`
	Description *string          `json:"description,omitempty"`
	// Pages may be rendered as a tree of files.
	//
	// Specify the `id` of another page as the `parent_page_id` to add some hierarchy to your pages.
	//
	ParentPageID *string `json:"parent_page_id,omitempty"`
}

func (c CreatePortalPageRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePortalPageRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePortalPageRequest) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *CreatePortalPageRequest) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *CreatePortalPageRequest) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *CreatePortalPageRequest) GetVisibility() *string {
	if o == nil {
		return nil
	}
	return o.Visibility
}

func (o *CreatePortalPageRequest) GetStatus() *PublishedStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *CreatePortalPageRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreatePortalPageRequest) GetParentPageID() *string {
	if o == nil {
		return nil
	}
	return o.ParentPageID
}
