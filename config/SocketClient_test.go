package config_test

import (
	"testing"

	"github.com/heaven-chp/base-client-go/config"
)

func TestSocketClient(t *testing.T) {
	socketClientConfig, err := config.Get[config.SocketClient]("./SocketClient.config")
	if err != nil {
		t.Fatal(err)
	}

	if socketClientConfig.Address != ":20000" {
		t.Fatal("invalid -", socketClientConfig.Address)
	}

	if socketClientConfig.Log.Level != "debug" {
		t.Fatal("invalid -", socketClientConfig.Log.Level)
	}

	if socketClientConfig.Log.Output != "file" {
		t.Fatal("invalid -", socketClientConfig.Log.Output)
	}

	if socketClientConfig.Log.File.Name != "./socket-client" {
		t.Fatal("invalid -", socketClientConfig.Log.File.Name)
	}

	if socketClientConfig.Log.File.ExtensionName != "log" {
		t.Fatal("invalid -", socketClientConfig.Log.File.ExtensionName)
	}

	if socketClientConfig.Log.File.AddDate {
		t.Fatal("invalid -", socketClientConfig.Log.File.AddDate)
	}

	if socketClientConfig.Log.WithCallerInfo == false {
		t.Fatal("invalid -", socketClientConfig.Log.WithCallerInfo)
	}
}
