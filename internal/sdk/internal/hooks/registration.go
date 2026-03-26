package hooks

import (
	"os"
)

func initHooks(h *Hooks) {
	h.registerBeforeRequestHook(&MeshDefaultsHook{})
	h.registerBeforeRequestHook(&AuthServerClientCreateHook{})

	// Domain customization - enable usage with non-prod domains
	h.registerBeforeRequestHook(&CustomizeKongDomainHook{
		Enabled:           os.Getenv("KONG_CUSTOM_DOMAIN") != "",
		Domain:            os.Getenv("KONG_CUSTOM_DOMAIN"),
		ReplaceFullDomain: os.Getenv("KONG_CUSTOM_DOMAIN_REPLACE_FULL_DOMAIN") == "true",
	})

	// Debug hooks - dump request/response
	h.registerBeforeRequestHook(&HTTPDumpRequestHook{
		Enabled: os.Getenv("KONNECT_SDK_HTTP_DUMP_REQUEST") == "true",
	})
	h.registerAfterSuccessHook(&HTTPDumpResponseHook{
		Enabled: os.Getenv("KONNECT_SDK_HTTP_DUMP_RESPONSE") == "true",
	})
	h.registerAfterErrorHook(&HTTPDumpResponseErrorHook{
		Enabled: os.Getenv("KONNECT_SDK_HTTP_DUMP_RESPONSE_ERROR") == "true",
	})

}
