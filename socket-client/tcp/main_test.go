package main

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/base-client/go/common/config"
	"github.com/common-library/go/file"
	"github.com/common-library/go/socket/tcp"
)

type TestServer struct {
	Network          string
	Address          string
	Greeting         string
	PrefixOfResponse string

	server tcp.Server
}

func (ts *TestServer) Start(t *testing.T) {
	ts.Network = "tcp"
	ts.Address = config.Get("socket.tcp.address").(string)
	ts.Greeting = "greeting"
	ts.PrefixOfResponse = "[response] "
	acceptSuccessFunc := func(client tcp.Client) {
		if writeLen, err := client.Write(ts.Greeting); err != nil {
			t.Fatal(err)
		} else if writeLen != len(ts.Greeting) {
			t.Fatalf("invalid write - (%d)(%d)", writeLen, len(ts.Greeting))
		}

		readData, err := client.Read(1024)
		if err != nil {
			t.Fatal(err)
		}

		writeData := ts.PrefixOfResponse + readData
		if writeLen, err := client.Write(writeData); err != nil {
			t.Fatal(err)
		} else if writeLen != len(writeData) {
			t.Fatalf("invalid write - (%d)(%d)", writeLen, len(writeData))
		}
	}

	acceptFailureFunc := func(err error) {
		t.Fatal(err)
	}

	if err := ts.server.Start(ts.Network, ts.Address, 100, acceptSuccessFunc, acceptFailureFunc); err != nil {
		t.Fatal(err)
	}

	for ts.server.GetCondition() == false {
		time.Sleep(100 * time.Millisecond)
	}
}

func (ts *TestServer) Stop(t *testing.T) {
	if err := ts.server.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestMain(t *testing.T) {
	const configFile = "../../common/config/config.yaml"

	if err := config.Read(configFile); err != nil {
		t.Fatal(err)
	}
	defer file.Remove(config.Get("socket.log.file.name").(string) + "." + config.Get("socket.log.file.extensionName").(string))

	testServer := TestServer{}
	testServer.Start(t)
	defer testServer.Stop(t)

	os.Args = []string{"test", "-config-file=" + configFile}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	main()
}
