package config

import (
	"testing"

	"github.com/heaven-chp/common-library-go/json"
)

func TestGrpcClient(t *testing.T) {
	var grpcClient GrpcClient

	err := json.ToStructFromFile("./grpc_client.config", &grpcClient)
	if err != nil {
		t.Error(err)
	}

	if grpcClient.LogLevel != "DEBUG" {
		t.Errorf("invalid data - LogLevel : (%s)", grpcClient.LogLevel)
	}

	if grpcClient.LogOutputPath != "./log/" {
		t.Errorf("invalid data - LogOutputPath : (%s)", grpcClient.LogOutputPath)
	}

	if grpcClient.LogFileNamePrefix != "grpc_client" {
		t.Errorf("invalid data - LogFileNamePrefix : (%s)", grpcClient.LogFileNamePrefix)
	}

	if grpcClient.Address != "127.0.0.1:50051" {
		t.Errorf("invalid data - Address : (%s)", grpcClient.Address)
	}

	if grpcClient.Timeout != 3 {
		t.Errorf("invalid data - Timeout : (%d)", grpcClient.Timeout)
	}
}

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
