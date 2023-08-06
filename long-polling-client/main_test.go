package main

import (
	"flag"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/heaven-chp/base-client-go/config"
	long_polling "github.com/heaven-chp/common-library-go/long-polling"
)

type TestServer struct {
	server long_polling.Server
}

func (this *TestServer) Start(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	configFile := path + "/../config/LongPollingClient.config"

	longPollingClientConfig := config.LongPollingClient{}
	err = config.Parsing(&longPollingClientConfig, configFile)
	if err != nil {
		t.Fatal(err)
	}

	serverInfo := long_polling.ServerInfo{
		Address:                        longPollingClientConfig.Address,
		Timeout:                        3600,
		SubscriptionURI:                longPollingClientConfig.SubscriptionURI,
		HandlerToRunBeforeSubscription: func(w http.ResponseWriter, r *http.Request) bool { return true },
		PublishURI:                     longPollingClientConfig.PublishURI,
		HandlerToRunBeforePublish:      func(w http.ResponseWriter, r *http.Request) bool { return true }}

	filePersistorInfo := long_polling.FilePersistorInfo{Use: false}

	err = this.server.Start(serverInfo, filePersistorInfo, func(err error) { t.Fatal(err) })
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Duration(200) * time.Millisecond)
}

func (this *TestServer) Stop(t *testing.T) {
	err := this.server.Stop(10)
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
		configFile := path + "/../config/LongPollingClient.config"

		os.Args = []string{"test", "-config_file=" + configFile}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		main()
	}
}
