package main

import (
	"flag"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/heaven-chp/base-client-go/config"
	"github.com/heaven-chp/common-library-go/socket"
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
	serverJob := func(client socket.Client) {
		const greeting = "greeting"
		const prefixOfResponse = "[response] "

		writeLen, err := client.Write(greeting)
		if err != nil {
			t.Error(err)
		}
		if writeLen != len(greeting) {
			t.Errorf("invalid write - (%d)(%d)", writeLen, len(greeting))
		}

		readData, err := client.Read(1024)
		if err != nil {
			t.Error(err)
		}

		writeData := prefixOfResponse + readData
		writeLen, err = client.Write(writeData)
		if err != nil {
			t.Error(err)
		}
		if writeLen != len(writeData) {
			t.Errorf("invalid write - (%d)(%d)", writeLen, len(writeData))
		}
	}

	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	configFile := path + "/../config/SocketClient.config"

	sleep := atomic.Bool{}
	sleep.Store(true)
	server := socket.Server{}
	go func() {
		socketClientConfig := config.SocketClient{}
		err := config.Parsing(&socketClientConfig, configFile)
		if err != nil {
			t.Error(err)
		}

		err = server.Start("tcp", socketClientConfig.Address, 1024, serverJob)
		if err != nil {
			t.Error(err)
		}
		sleep.Store(false)
	}()
	for sleep.Load() && server.GetCondition() == false {
		time.Sleep(100 * time.Millisecond)
	}
	defer func() {
		err := server.Stop()
		if err != nil {
			t.Error(err)
		}
	}()

	{
		os.Args = []string{"test", "-config_file=" + configFile}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		main()
	}
}
