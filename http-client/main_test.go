package main

import (
	"flag"
	net_http "net/http"
	"os"
	"testing"
	"time"

	"github.com/base-client/go/common/config"
	"github.com/common-library/go/file"
	"github.com/common-library/go/http/mux"
)

func TestMain(t *testing.T) {
	server := mux.Server{}
	server.RegisterHandlerFunc(net_http.MethodGet, "/v1/test/{id}", func(responseWriter net_http.ResponseWriter, request *net_http.Request) {
		responseWriter.WriteHeader(net_http.StatusOK)
		responseWriter.Write([]byte(`{"field_1":"value-1"}`))
	})

	middlewareFunction := func(nextHandler net_http.Handler) net_http.Handler {
		return net_http.HandlerFunc(func(responseWriter net_http.ResponseWriter, request *net_http.Request) {
			nextHandler.ServeHTTP(responseWriter, request)
		})
	}
	server.Use(middlewareFunction)

	const configFile = "../common/config/config.yaml"

	if err := config.Read(configFile); err != nil {
		t.Fatal(err)
	}
	defer file.Remove(config.Get("http.log.file.name").(string) + "." + config.Get("http.log.file.extensionName").(string))

	if err := server.Start(config.Get("http.address").(string), func(err error) { t.Error(err) }); err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Duration(200) * time.Millisecond)

	defer func() {
		if err := server.Stop(5); err != nil {
			t.Fatal(err)
		}
	}()

	os.Args = []string{"test", "-config-file=" + configFile}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	main()
}
