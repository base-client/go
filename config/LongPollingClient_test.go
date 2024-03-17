package config_test

import (
	"testing"

	"github.com/heaven-chp/base-client-go/config"
)

func TestLongPollingClient(t *testing.T) {
	longPollingClientConfig, err := config.Get[config.LongPollingClient]("./LongPollingClient.config")
	if err != nil {
		t.Fatal(err)
	}

	if longPollingClientConfig.Address != ":30000" {
		t.Fatal("invalid -", longPollingClientConfig.Address)
	}

	if longPollingClientConfig.SubscriptionURI != "/subscription" {
		t.Fatal("invalid -", longPollingClientConfig.SubscriptionURI)
	}

	if longPollingClientConfig.PublishURI != "/publish" {
		t.Fatal("invalid -", longPollingClientConfig.PublishURI)
	}

	if longPollingClientConfig.Log.Level != "debug" {
		t.Fatal("invalid -", longPollingClientConfig.Log.Level)
	}

	if longPollingClientConfig.Log.Output != "file" {
		t.Fatal("invalid -", longPollingClientConfig.Log.Output)
	}

	if longPollingClientConfig.Log.File.Name != "./long-polling-client" {
		t.Fatal("invalid -", longPollingClientConfig.Log.File.Name)
	}

	if longPollingClientConfig.Log.File.ExtensionName != "log" {
		t.Fatal("invalid -", longPollingClientConfig.Log.File.ExtensionName)
	}

	if longPollingClientConfig.Log.File.AddDate {
		t.Fatal("invalid -", longPollingClientConfig.Log.File.AddDate)
	}

	if longPollingClientConfig.Log.WithCallerInfo == false {
		t.Fatal("invalid -", longPollingClientConfig.Log.WithCallerInfo)
	}
}
