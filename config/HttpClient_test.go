package config_test

import (
	"testing"

	"github.com/heaven-chp/base-client-go/config"
)

func TestHttpClient(t *testing.T) {
	httpClientConfig := config.HttpClient{}

	err := config.Parsing(&httpClientConfig, "./HttpClient.config")
	if err != nil {
		t.Fatal(err)
	}

	if httpClientConfig.LogLevel != "DEBUG" {
		t.Errorf("invalid data - LogLevel : (%s)", httpClientConfig.LogLevel)
	}

	if httpClientConfig.LogOutputPath != "./log/" {
		t.Errorf("invalid data - LogOutputPath : (%s)", httpClientConfig.LogOutputPath)
	}

	if httpClientConfig.LogFileNamePrefix != "http-client" {
		t.Errorf("invalid data - LogFileNamePrefix : (%s)", httpClientConfig.LogFileNamePrefix)
	}

	if httpClientConfig.Address != ":10000" {
		t.Errorf("invalid data - Address : (%s)", httpClientConfig.Address)
	}
}
