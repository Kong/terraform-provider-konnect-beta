package hooks

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"testing"
)

func TestCheckMatchingFeatures(t *testing.T) {
	type testCase struct {
		name           string
		input          string
		matchingFn     func(*http.Request) bool
		expectedOutput bool
	}

	testCases := []testCase{
		{
			name:           "konnect matches",
			input:          "/v1/mesh/control-planes",
			matchingFn:     konnectMatchFeatures,
			expectedOutput: true,
		},
		{
			name:           "konnect does not match",
			input:          "/v0/mesh/control-planes",
			matchingFn:     konnectMatchFeatures,
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			matched := tc.matchingFn(&http.Request{
				Method: http.MethodPost,
				URL: &url.URL{
					Path: tc.input,
				},
			})

			require.Equal(t, tc.expectedOutput, matched)
		})
	}
}

func TestCheckMatchingPolicies(t *testing.T) {
	type testCase struct {
		name           string
		input          string
		matchingFn     func(*http.Request) bool
		expectedOutput bool
	}

	testCases := []testCase{
		{
			name:           "konnect matches",
			input:          "/v1/mesh/control-planes/d15693a7-ed2e-4b77-a54d-9c3675c49c3f/api/meshes/default",
			matchingFn:     konnectMatchPolicies,
			expectedOutput: true,
		},
		{
			name:           "konnect does not match",
			input:          "/v1/mesh/control-planes/d15693a7-ed2e-4b77-a54d-9c3675c49c3f/api/meshtrafficpermissions/default",
			matchingFn:     konnectMatchPolicies,
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			matched := tc.matchingFn(&http.Request{
				Method: http.MethodPut,
				URL: &url.URL{
					Path: tc.input,
				},
			})

			require.Equal(t, tc.expectedOutput, matched)
		})
	}
}
