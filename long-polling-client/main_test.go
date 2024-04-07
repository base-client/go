package main

import (
	"flag"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/base-client/go/config"
	"github.com/common-library/go/file"
	long_polling "github.com/common-library/go/long-polling"
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

	longPollingClientConfig, err := config.Get[config.LongPollingClient](configFile)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Remove(longPollingClientConfig.Log.File.Name + "." + longPollingClientConfig.Log.File.ExtensionName)

	serverInfo := long_polling.ServerInfo{
		Address:                        longPollingClientConfig.Address,
		Timeout:                        3600,
		SubscriptionURI:                longPollingClientConfig.SubscriptionURI,
		HandlerToRunBeforeSubscription: func(w http.ResponseWriter, r *http.Request) bool { return true },
		PublishURI:                     longPollingClientConfig.PublishURI,
		HandlerToRunBeforePublish:      func(w http.ResponseWriter, r *http.Request) bool { return true }}

	filePersistorInfo := long_polling.FilePersistorInfo{Use: false}

	if err := this.server.Start(serverInfo, filePersistorInfo, func(err error) { t.Fatal(err) }); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Duration(200) * time.Millisecond)
}

func (this *TestServer) Stop(t *testing.T) {
	if err := this.server.Stop(10); err != nil {
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
	configFile := path + "/../config/LongPollingClient.config"

	if longPollingClientConfig, err := config.Get[config.LongPollingClient](configFile); err != nil {
		t.Fatal(err)
	} else {
		defer file.Remove(longPollingClientConfig.Log.File.Name + "." + longPollingClientConfig.Log.File.ExtensionName)
	}

	os.Args = []string{"test", "-config_file=" + configFile}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	main()
}
