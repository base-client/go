package main

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/base-client/go/config"
	"github.com/common-library/go/file"
	"github.com/common-library/go/grpc"
	"github.com/common-library/go/grpc/sample"
)

func TestMain1(t *testing.T) {
	os.Args = []string{"test"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	if err := (&Main{}).Run(); err.Error() != "invalid flag" {
		t.Fatal(err)
	}
}

func TestMain2(t *testing.T) {
	os.Args = []string{"test", "-config_file=invalid"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	if err := (&Main{}).Run(); err.Error() != "open invalid: no such file or directory" {
		t.Fatal(err)
	}
}

func TestMain3(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	configFile := path + "/../config/GrpcClient.config"

	grpcClientConfig, err := config.Get[config.GrpcClient](configFile)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Remove(grpcClientConfig.Log.File.Name + "." + grpcClientConfig.Log.File.ExtensionName)

	server := grpc.Server{}
	go func() {
		if err := server.Start(grpcClientConfig.Address, &sample.Server{}); err != nil {
			t.Fatal(err)
		}
	}()
	time.Sleep(100 * time.Millisecond)

	os.Args = []string{"test", "-config_file=" + configFile}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	main()

	if err := server.Stop(); err != nil {
		t.Fatal(err)
	}
}
