package commands

import (
	"fmt"

	"os"

	"code.cloudfoundry.org/credhub-cli/config"
	"code.cloudfoundry.org/credhub-cli/errors"
	"code.cloudfoundry.org/credhub-cli/version"
)

func PrintVersion() error {
	cfg := config.ReadConfig()

	credHubServerVersion := "Not Found. Have you targeted and authenticated against a CredHub server?"
	fmt.Println("CLI Version:", version.Version)

	if cfg.ApiURL != "" {
		credhubClient, err := initializeCredhubClient(cfg)

		if err == nil || err.Error() != errors.NewRevokedTokenError().Error() {
			_, err := credhubClient.FindAllPaths()

			if err == nil {
				version, err := credhubClient.ServerVersion()
				if err == nil {
					credHubServerVersion = version.String()
				}
			}
		}
	}

	fmt.Println("Server Version:", credHubServerVersion)

	return nil
}

func init() {
	CredHub.Version = func() {
		PrintVersion()
		os.Exit(0)
	}
}
