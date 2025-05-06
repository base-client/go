package main

import (
	"flag"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/base-client/go/common/config"
	"github.com/common-library/go/file"
	long_polling "github.com/common-library/go/long-polling"
)

type TestServer struct {
	server long_polling.Server
}

func (this *TestServer) Start(t *testing.T) {
	serverInfo := long_polling.ServerInfo{
		Address:                        config.Get("longPolling.address").(string),
		TimeoutSeconds:                 3600,
		SubscriptionURI:                config.Get("longPolling.subscriptionURI").(string),
		HandlerToRunBeforeSubscription: func(w http.ResponseWriter, r *http.Request) bool { return true },
		PublishURI:                     config.Get("longPolling.publishURI").(string),
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

func TestMain(t *testing.T) {
	const configFile = "../common/config/config.yaml"

	if err := config.Read(configFile); err != nil {
		t.Fatal(err)
	}
	defer file.Remove(config.Get("longPolling.log.file.name").(string) + "." + config.Get("longPolling.log.file.extensionName").(string))

	testServer := TestServer{}
	testServer.Start(t)
	defer testServer.Stop(t)

	os.Args = []string{"test", "-config-file=" + configFile}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	main()
}
