package config_test

import (
	"testing"

	"github.com/heaven-chp/base-client-go/config"
)

func TestGrpcClient(t *testing.T) {
	grpcClientConfig := config.GrpcClient{}

	err := config.Parsing(&grpcClientConfig, "./GrpcClient.config")
	if err != nil {
		t.Fatal(err)
	}

	if grpcClientConfig.Address != ":50051" {
		t.Errorf("invalid data - Address : (%s)", grpcClientConfig.Address)
	}

	if grpcClientConfig.Timeout != 3 {
		t.Errorf("invalid data - Timeout : (%d)", grpcClientConfig.Timeout)
	}

	if grpcClientConfig.Log.Level != "DEBUG" {
		t.Errorf("invalid data - Log.Level : (%s)", grpcClientConfig.Log.Level)
	}

	if grpcClientConfig.Log.OutputPath != "./log/" {
		t.Errorf("invalid data - Log.OutputPath : (%s)", grpcClientConfig.Log.OutputPath)
	}

	if grpcClientConfig.Log.FileNamePrefix != "grpc-client" {
		t.Errorf("invalid data - Log.FileNamePrefix : (%s)", grpcClientConfig.Log.FileNamePrefix)
	}

	if grpcClientConfig.Log.PrintCallerInfo != true {
		t.Errorf("invalid data - Log.PrintCallerInfo : (%t)", grpcClientConfig.Log.PrintCallerInfo)
	}

	if grpcClientConfig.Log.ChannelSize != 1024 {
		t.Errorf("invalid data - Log.ChannelSize : (%d)", grpcClientConfig.Log.ChannelSize)
	}
}
