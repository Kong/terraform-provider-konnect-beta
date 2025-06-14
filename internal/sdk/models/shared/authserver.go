// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/internal/utils"
	"time"
)

// AuthServer - An auth server
type AuthServer struct {
	// The ID of the auth server
	ID string `json:"id"`
	// The name of the auth server
	Name string `json:"name"`
	// The description of the auth server
	Description string `json:"description"`
	// The recipients that the tokens are intended for. This becomes the 'aud' claim in an access token
	Audience string `json:"audience"`
	// Algorithm used in the key signing process
	SigningAlgorithm *Algorithm `default:"RS256" json:"signing_algorithm"`
	// The complete URL for the custom authorization server. This becomes the 'iss' claim in an access token.
	Issuer string `json:"issuer"`
	// The URI of the metadata document for the auth server
	MetadataURI string `json:"metadata_uri"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels map[string]*string `json:"labels"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (a AuthServer) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthServer) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AuthServer) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AuthServer) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AuthServer) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *AuthServer) GetAudience() string {
	if o == nil {
		return ""
	}
	return o.Audience
}

func (o *AuthServer) GetSigningAlgorithm() *Algorithm {
	if o == nil {
		return nil
	}
	return o.SigningAlgorithm
}

func (o *AuthServer) GetIssuer() string {
	if o == nil {
		return ""
	}
	return o.Issuer
}

func (o *AuthServer) GetMetadataURI() string {
	if o == nil {
		return ""
	}
	return o.MetadataURI
}

func (o *AuthServer) GetLabels() map[string]*string {
	if o == nil {
		return map[string]*string{}
	}
	return o.Labels
}

func (o *AuthServer) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *AuthServer) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
