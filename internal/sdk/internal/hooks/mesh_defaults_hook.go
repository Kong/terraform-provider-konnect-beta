package hooks

import (
	shared_speakeasy "github.com/Kong/shared-speakeasy/hooks/mesh_defaults"
	"net/http"
	"regexp"
)

// MeshDefaultsHook is a struct that implements the BeforeRequestHook interface.
type MeshDefaultsHook struct{}

var konnectMatchFeatures = func(req *http.Request) bool {
	return req.URL.Path == "/v1/mesh/control-planes" && req.Method == http.MethodPost
}

var konnectMatchPolicies = func(req *http.Request) bool {
	match, err := regexp.MatchString(`^/v1/mesh/control-planes/[^/]+/api/meshes/[^/]+$`, req.URL.Path)
	if err != nil {
		return false
	}

	return match && req.Method == http.MethodPut
}

// BeforeRequest modifies the request before sending it.
func (e MeshDefaultsHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	return shared_speakeasy.BeforeRequest(konnectMatchFeatures, konnectMatchPolicies)(req)
}
