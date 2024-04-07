package config_test

import (
	"testing"

	"github.com/base-client/go/config"
)

func TestHttpClient(t *testing.T) {
	httpClientConfig, err := config.Get[config.HttpClient]("./HttpClient.config")
	if err != nil {
		t.Fatal(err)
	}

	if httpClientConfig.Address != ":10000" {
		t.Fatal("invalid -", httpClientConfig.Address)
	}

	if httpClientConfig.Log.Level != "debug" {
		t.Fatal("invalid -", httpClientConfig.Log.Level)
	}

	if httpClientConfig.Log.Output != "file" {
		t.Fatal("invalid -", httpClientConfig.Log.Output)
	}

	if httpClientConfig.Log.File.Name != "./http-client" {
		t.Fatal("invalid -", httpClientConfig.Log.File.Name)
	}

	if httpClientConfig.Log.File.ExtensionName != "log" {
		t.Fatal("invalid -", httpClientConfig.Log.File.ExtensionName)
	}

	if httpClientConfig.Log.File.AddDate {
		t.Fatal("invalid -", httpClientConfig.Log.File.AddDate)
	}

	if httpClientConfig.Log.WithCallerInfo == false {
		t.Fatal("invalid -", httpClientConfig.Log.WithCallerInfo)
	}
}
