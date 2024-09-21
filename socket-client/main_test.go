package main

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/base-client/go/common/config"
	"github.com/common-library/go/file"
	"github.com/common-library/go/socket"
)

type TestServer struct {
	Network          string
	Address          string
	Greeting         string
	PrefixOfResponse string

	server socket.Server
}

func (this *TestServer) Start(t *testing.T) {
	this.Network = "tcp"
	this.Address = config.Get("socket.address").(string)
	this.Greeting = "greeting"
	this.PrefixOfResponse = "[response] "

	acceptSuccessFunc := func(client socket.Client) {
		if writeLen, err := client.Write(this.Greeting); err != nil {
			t.Fatal(err)
		} else if writeLen != len(this.Greeting) {
			t.Fatalf("invalid write - (%d)(%d)", writeLen, len(this.Greeting))
		}

		readData, err := client.Read(1024)
		if err != nil {
			t.Fatal(err)
		}

		writeData := this.PrefixOfResponse + readData
		if writeLen, err := client.Write(writeData); err != nil {
			t.Fatal(err)
		} else if writeLen != len(writeData) {
			t.Fatalf("invalid write - (%d)(%d)", writeLen, len(writeData))
		}
	}

	acceptFailureFunc := func(err error) {
		t.Fatal(err)
	}

	if err := this.server.Start(this.Network, this.Address, 100, acceptSuccessFunc, acceptFailureFunc); err != nil {
		t.Fatal(err)
	}

	for this.server.GetCondition() == false {
		time.Sleep(100 * time.Millisecond)
	}
}

func (this *TestServer) Stop(t *testing.T) {
	if err := this.server.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestMain(t *testing.T) {
	const configFile = "../common/config/config.yaml"

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
