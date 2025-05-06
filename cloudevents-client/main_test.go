package main

import (
	"flag"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/base-client/go/common/config"
	"github.com/common-library/go/event/cloudevents"
)

func TestMain(t *testing.T) {
	const configFile = "../common/config/config.yaml"

	if err := config.Read(configFile); err != nil {
		t.Fatal(err)
	}

	address := config.Get("cloudEvents.address").(string)
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

	os.Args = []string{"test", "-config-file=" + configFile}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	main()
}
