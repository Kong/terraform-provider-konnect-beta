package hooks

import(
"os"
)

func initHooks(h *Hooks) {
    h.registerBeforeRequestHook(&MeshDefaultsHook{})

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
