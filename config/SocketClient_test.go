package config_test

import (
	"testing"

	"github.com/heaven-chp/base-client-go/config"
)

func TestSocketClient(t *testing.T) {
	socketClient := config.SocketClient{}

	err := config.Parsing(&socketClient, "./SocketClient.config")
	if err != nil {
		t.Fatal(err)
	}

	if socketClient.LogLevel != "DEBUG" {
		t.Errorf("invalid data - LogLevel : (%s)", socketClient.LogLevel)
	}

	if socketClient.LogOutputPath != "./log/" {
		t.Errorf("invalid data - LogOutputPath : (%s)", socketClient.LogOutputPath)
	}

	if socketClient.LogFileNamePrefix != "socket-client" {
		t.Errorf("invalid data - LogFileNamePrefix : (%s)", socketClient.LogFileNamePrefix)
	}

	if socketClient.Address != ":20000" {
		t.Errorf("invalid data - Address : (%s)", socketClient.Address)
	}
}
