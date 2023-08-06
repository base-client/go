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

	if httpClientConfig.Address != ":10000" {
		t.Errorf("invalid data - Address : (%s)", httpClientConfig.Address)
	}

	if httpClientConfig.Log.Level != "DEBUG" {
		t.Errorf("invalid data - Log.Level : (%s)", httpClientConfig.Log.Level)
	}

	if httpClientConfig.Log.OutputPath != "./log/" {
		t.Errorf("invalid data - Log.OutputPath : (%s)", httpClientConfig.Log.OutputPath)
	}

	if httpClientConfig.Log.FileNamePrefix != "http-client" {
		t.Errorf("invalid data - Log.FileNamePrefix : (%s)", httpClientConfig.Log.FileNamePrefix)
	}

	if httpClientConfig.Log.PrintCallerInfo != true {
		t.Errorf("invalid data - Log.PrintCallerInfo : (%t)", httpClientConfig.Log.PrintCallerInfo)
	}

	if httpClientConfig.Log.ChannelSize != 1024 {
		t.Errorf("invalid data - Log.ChannelSize : (%d)", httpClientConfig.Log.ChannelSize)
	}

}
