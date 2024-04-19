package main

import (
	"flag"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/base-client/go/config"
	"github.com/common-library/go/event/cloudevents"
)

func TestMain1(t *testing.T) {
	os.Args = []string{"test"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	if err := (&Main{}).Run(); err.Error() != "invalid flag" {
		t.Fatal(err)
	}
}

func TestMain2(t *testing.T) {
	os.Args = []string{"test", "-config-file=invalid"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	if err := (&Main{}).Run(); err.Error() != "open invalid: no such file or directory" {
		t.Fatal(err)
	}
}

func TestMain3(t *testing.T) {
	configFile := "/../config/CloudEventsClient.config"
	if path, err := os.Getwd(); err != nil {
		t.Fatal(err)
	} else {
		configFile = path + configFile
	}

	if cloudEventsClientConfig, err := config.Get[config.CloudEventsClient](configFile); err != nil {
		t.Fatal(err)
	} else {
		address := strings.Replace(cloudEventsClientConfig.Address, "http://", "", 1)
		handler := func(event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
			responseEvent := event.Clone()
			return &responseEvent, cloudevents.NewHTTPResult(http.StatusOK, "")
		}
		listenAndServeFailureFunc := func(err error) { t.Fatal(err) }

		server := cloudevents.Server{}
		if err := server.Start(address, handler, listenAndServeFailureFunc); err != nil {
			t.Fatal(err)
		}
		time.Sleep(200 * time.Millisecond)
		defer func() {
			if err := server.Stop(10); err != nil {
				t.Fatal(err)
			}
		}()
	}

	os.Args = []string{"test", "-config-file=" + configFile}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	main()
}
