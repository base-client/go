package config_test

import (
	"testing"

	"github.com/heaven-chp/base-client-go/config"
)

func TestSocketClient(t *testing.T) {
	socketClientConfig := config.SocketClient{}

	err := config.Parsing(&socketClientConfig, "./SocketClient.config")
	if err != nil {
		t.Fatal(err)
	}

	if socketClientConfig.Address != ":20000" {
		t.Errorf("invalid data - Address : (%s)", socketClientConfig.Address)
	}

	if socketClientConfig.Log.Level != "DEBUG" {
		t.Errorf("invalid data - Log.Level : (%s)", socketClientConfig.Log.Level)
	}

	if socketClientConfig.Log.OutputPath != "./log/" {
		t.Errorf("invalid data - Log.OutputPath : (%s)", socketClientConfig.Log.OutputPath)
	}

	if socketClientConfig.Log.FileNamePrefix != "socket-client" {
		t.Errorf("invalid data - Log.FileNamePrefix : (%s)", socketClientConfig.Log.FileNamePrefix)
	}

	if socketClientConfig.Log.PrintCallerInfo != true {
		t.Errorf("invalid data - Log.PrintCallerInfo : (%t)", socketClientConfig.Log.PrintCallerInfo)
	}

	if socketClientConfig.Log.ChannelSize != 1024 {
		t.Errorf("invalid data - Log.ChannelSize : (%d)", socketClientConfig.Log.ChannelSize)
	}
}
