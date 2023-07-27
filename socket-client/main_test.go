package main

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/heaven-chp/base-client-go/config"
	"github.com/heaven-chp/common-library-go/socket"
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

	socketClientConfig := config.SocketClient{}
	err = config.Parsing(&socketClientConfig, configFile)
	if err != nil {
		t.Fatal(err)
	}

	this.Network = "tcp"
	this.Address = socketClientConfig.Address
	this.Greeting = "greeting"
	this.PrefixOfResponse = "[response] "

	acceptSuccessFunc := func(client socket.Client) {
		writeLen, err := client.Write(this.Greeting)
		if err != nil {
			t.Error(err)
		}
		if writeLen != len(this.Greeting) {
			t.Errorf("invalid write - (%d)(%d)", writeLen, len(this.Greeting))
		}

		readData, err := client.Read(1024)
		if err != nil {
			t.Error(err)
		}

		writeData := this.PrefixOfResponse + readData
		writeLen, err = client.Write(writeData)
		if err != nil {
			t.Error(err)
		}
		if writeLen != len(writeData) {
			t.Errorf("invalid write - (%d)(%d)", writeLen, len(writeData))
		}
	}

	acceptFailureFunc := func(err error) {
		t.Error(err)
	}

	err = this.server.Start(this.Network, this.Address, 100, acceptSuccessFunc, acceptFailureFunc)
	if err != nil {
		t.Fatal(err)
	}
	for this.server.GetCondition() == false {
		time.Sleep(100 * time.Millisecond)
	}
}

func (this *TestServer) Stop(t *testing.T) {
	err := this.server.Stop()
	if err != nil {
		t.Fatal(err)
	}
}

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
	testServer := TestServer{}
	testServer.Start(t)
	defer testServer.Stop(t)

	{
		path, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}
		configFile := path + "/../config/SocketClient.config"

		os.Args = []string{"test", "-config_file=" + configFile}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		main()
	}
}
