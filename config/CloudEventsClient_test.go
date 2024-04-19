package config_test

import (
	"testing"

	"github.com/base-client/go/config"
)

func TestCloudEventsClient(t *testing.T) {
	if cloudEventsClientConfig, err := config.Get[config.CloudEventsClient]("./CloudEventsClient.config"); err != nil {
		t.Fatal(err)
	} else if cloudEventsClientConfig.Address != "http://:40000" {
		t.Fatal("invalid -", cloudEventsClientConfig.Address)
	}
}
