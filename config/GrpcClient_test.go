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

	if grpcClientConfig.LogLevel != "DEBUG" {
		t.Errorf("invalid data - LogLevel : (%s)", grpcClientConfig.LogLevel)
	}

	if grpcClientConfig.LogOutputPath != "./log/" {
		t.Errorf("invalid data - LogOutputPath : (%s)", grpcClientConfig.LogOutputPath)
	}

	if grpcClientConfig.LogFileNamePrefix != "grpc-client" {
		t.Errorf("invalid data - LogFileNamePrefix : (%s)", grpcClientConfig.LogFileNamePrefix)
	}

	if grpcClientConfig.Address != ":50051" {
		t.Errorf("invalid data - Address : (%s)", grpcClientConfig.Address)
	}

	if grpcClientConfig.Timeout != 3 {
		t.Errorf("invalid data - Timeout : (%d)", grpcClientConfig.Timeout)
	}
}
