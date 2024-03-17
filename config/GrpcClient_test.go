package config_test

import (
	"testing"

	"github.com/heaven-chp/base-client-go/config"
)

func TestGrpcClient(t *testing.T) {
	grpcClientConfig, err := config.Get[config.GrpcClient]("./GrpcClient.config")
	if err != nil {
		t.Fatal(err)
	}

	if grpcClientConfig.Address != ":50051" {
		t.Fatal("invalid -", grpcClientConfig.Address)
	}

	if grpcClientConfig.Timeout != 3 {
		t.Fatal("invalid -", grpcClientConfig.Timeout)
	}

	if grpcClientConfig.Log.Level != "debug" {
		t.Fatal("invalid -", grpcClientConfig.Log.Level)
	}

	if grpcClientConfig.Log.Output != "file" {
		t.Fatal("invalid -", grpcClientConfig.Log.Output)
	}

	if grpcClientConfig.Log.File.Name != "./grpc-client" {
		t.Fatal("invalid -", grpcClientConfig.Log.File.Name)
	}

	if grpcClientConfig.Log.File.ExtensionName != "log" {
		t.Fatal("invalid -", grpcClientConfig.Log.File.ExtensionName)
	}

	if grpcClientConfig.Log.File.AddDate {
		t.Fatal("invalid -", grpcClientConfig.Log.File.AddDate)
	}

	if grpcClientConfig.Log.WithCallerInfo == false {
		t.Fatal("invalid -", grpcClientConfig.Log.WithCallerInfo)
	}
}
