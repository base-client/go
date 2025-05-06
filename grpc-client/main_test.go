package main

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/base-client/go/common/config"
	"github.com/common-library/go/file"
	"github.com/common-library/go/grpc"
	"github.com/common-library/go/grpc/sample"
)

func TestMain(t *testing.T) {
	const configFile = "../common/config/config.yaml"

	if err := config.Read(configFile); err != nil {
		t.Fatal(err)
	}
	defer file.Remove(config.Get("gRPC.log.file.name").(string) + "." + config.Get("gRPC.log.file.extensionName").(string))

	server := grpc.Server{}
	go func() {
		if err := server.Start(config.Get("gRPC.address").(string), &sample.Server{}); err != nil {
			t.Fatal(err)
		}
	}()
	time.Sleep(100 * time.Millisecond)
	defer func() {
		if err := server.Stop(); err != nil {
			t.Fatal(err)
		}
	}()

	os.Args = []string{"test", "-config-file=" + configFile}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	main()
}
