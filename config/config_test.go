package config

import (
	"github.com/heaven-chp/common-library-go/json"
	"testing"
)

func TestSocketClient(t *testing.T) {
	var socketClient SocketClient

	err := json.ToStructFromFile("./socket_client.config", &socketClient)
	if err != nil {
		t.Error(err)
	}

	if socketClient.LogLevel != "DEBUG" {
		t.Errorf("invalid data - LogLevel : (%s)", socketClient.LogLevel)
	}

	if socketClient.LogOutputPath != "./log/" {
		t.Errorf("invalid data - LogOutputPath : (%s)", socketClient.LogOutputPath)
	}

	if socketClient.LogFileNamePrefix != "socket_client" {
		t.Errorf("invalid data - LogFileNamePrefix : (%s)", socketClient.LogFileNamePrefix)
	}

	if socketClient.Address != "127.0.0.1:11111" {
		t.Errorf("invalid data - Address : (%s)", socketClient.Address)
	}
}
