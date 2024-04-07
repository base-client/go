package main

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/base-client/go/config"
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
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	configFile := path + "/../config/SocketClient.config"

	socketClientConfig, err := config.Get[config.SocketClient](configFile)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Remove(socketClientConfig.Log.File.Name + "." + socketClientConfig.Log.File.ExtensionName)

	this.Network = "tcp"
	this.Address = socketClientConfig.Address
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
	testServer := TestServer{}
	testServer.Start(t)
	defer testServer.Stop(t)

	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	configFile := path + "/../config/SocketClient.config"

	if socketClientConfig, err := config.Get[config.SocketClient](configFile); err != nil {
		t.Fatal(err)
	} else {
		defer file.Remove(socketClientConfig.Log.File.Name + "." + socketClientConfig.Log.File.ExtensionName)
	}

	os.Args = []string{"test", "-config_file=" + configFile}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	main()
}
