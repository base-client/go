package main

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/heaven-chp/base-client-go/config"
	"github.com/heaven-chp/common-library-go/grpc"
	"github.com/heaven-chp/common-library-go/grpc/sample"
)

func TestMain1(t *testing.T) {
	os.Args = []string{"test"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	main := Main{}
	err := main.Run()
	if err.Error() != "invalid flag" {
		t.Error(err)
	}
}

func TestMain2(t *testing.T) {
	os.Args = []string{"test", "-config_file=invalid"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	main := Main{}
	err := main.Run()
	if err.Error() != "open invalid: no such file or directory" {
		t.Error(err)
	}
}

func TestMain3(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	configFile := path + "/../config/GrpcClient.config"

	grpcClientConfig := config.GrpcClient{}
	err = config.Parsing(&grpcClientConfig, configFile)
	if err != nil {
		t.Fatal(err)
	}

	server := grpc.Server{}
	go func() {
		err := server.Start(grpcClientConfig.Address, &sample.Server{})
		if err != nil {
			t.Error(err)
		}
	}()
	time.Sleep(100 * time.Millisecond)

	{
		os.Args = []string{"test", "-config_file=" + configFile}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		main()
	}

	err = server.Stop()
	if err != nil {
		t.Error(err)
	}
}
