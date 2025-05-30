// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PublishedStatus - Whether the resource is visible on a given portal. Defaults to false.
type PublishedStatus string

const (
	PublishedStatusPublished   PublishedStatus = "published"
	PublishedStatusUnpublished PublishedStatus = "unpublished"
)

func (e PublishedStatus) ToPointer() *PublishedStatus {
	return &e
}
func (e *PublishedStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "published":
		fallthrough
	case "unpublished":
		*e = PublishedStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PublishedStatus: %v", v)
	}
}
