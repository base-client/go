package config_test

import (
	"testing"

	"github.com/heaven-chp/base-client-go/config"
)

func TestLongPollingClient(t *testing.T) {
	longPollingClientConfig := config.LongPollingClient{}

	err := config.Parsing(&longPollingClientConfig, "./LongPollingClient.config")
	if err != nil {
		t.Fatal(err)
	}

	if longPollingClientConfig.Address != ":30000" {
		t.Errorf("invalid data - Address : (%s)", longPollingClientConfig.Address)
	}

	if longPollingClientConfig.SubscriptionURI != "/subscription" {
		t.Errorf("invalid data - SubscriptionURI : (%s)", longPollingClientConfig.SubscriptionURI)
	}

	if longPollingClientConfig.PublishURI != "/publish" {
		t.Errorf("invalid data - PublishURI : (%s)", longPollingClientConfig.PublishURI)
	}

	if longPollingClientConfig.Log.Level != "DEBUG" {
		t.Errorf("invalid data - Log.Level : (%s)", longPollingClientConfig.Log.Level)
	}

	if longPollingClientConfig.Log.OutputPath != "./log/" {
		t.Errorf("invalid data - Log.OutputPath : (%s)", longPollingClientConfig.Log.OutputPath)
	}

	if longPollingClientConfig.Log.FileNamePrefix != "long-polling-client" {
		t.Errorf("invalid data - Log.FileNamePrefix : (%s)", longPollingClientConfig.Log.FileNamePrefix)
	}

	if longPollingClientConfig.Log.PrintCallerInfo != true {
		t.Errorf("invalid data - Log.PrintCallerInfo : (%t)", longPollingClientConfig.Log.PrintCallerInfo)
	}

	if longPollingClientConfig.Log.ChannelSize != 1024 {
		t.Errorf("invalid data - Log.ChannelSize : (%d)", longPollingClientConfig.Log.ChannelSize)
	}
}
