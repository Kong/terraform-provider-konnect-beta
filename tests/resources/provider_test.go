package tests

import (
	"log"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/kong/terraform-provider-konnect-beta/internal/provider"
)

var (
	providerFactory = map[string]func() (tfprotov6.ProviderServer, error){
		"konnect-beta": providerserver.NewProtocol6WithError(provider.New("")()),
	}
)

func providerConfigFromEnv() (string, int, string) {
	serverHost := "us.api.konghq.com"
	if len(os.Getenv("E2E_SERVER_URL")) > 0 {
		serverHost = os.Getenv("E2E_SERVER_URL")
	}

	serverPort := 443
	if len(os.Getenv("E2E_SERVER_PORT")) > 0 {
		serverPortStr := os.Getenv("E2E_SERVER_PORT")
		var err error
		serverPort, err = strconv.Atoi(serverPortStr)
		if err != nil {
			log.Fatal("Invalid server port:", err)
		}
	}

	serverScheme := "https"
	if len(os.Getenv("E2E_SERVER_SCHEME")) > 0 {
		serverScheme = os.Getenv("E2E_SERVER_SCHEME")
	}

	return serverHost, serverPort, serverScheme
}
