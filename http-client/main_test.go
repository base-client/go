package main

import (
	"flag"
	net_http "net/http"
	"os"
	"testing"
	"time"

	"github.com/base-client/go/config"
	"github.com/common-library/go/file"
	"github.com/common-library/go/http"
)

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
	server := http.Server{}
	server.RegisterHandlerFunc("/v1/test/{id}", net_http.MethodGet, func(responseWriter net_http.ResponseWriter, request *net_http.Request) {
		responseWriter.WriteHeader(net_http.StatusOK)
		responseWriter.Write([]byte(`{"field_1":"value-1"}`))
	})

	middlewareFunction := func(nextHandler net_http.Handler) net_http.Handler {
		return net_http.HandlerFunc(func(responseWriter net_http.ResponseWriter, request *net_http.Request) {
			nextHandler.ServeHTTP(responseWriter, request)
		})
	}

	if err := server.Start(":10000", func(err error) { t.Error(err) }, middlewareFunction); err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Duration(200) * time.Millisecond)

	defer func() {
		if err := server.Stop(5); err != nil {
			t.Fatal(err)
		}
	}()

	{
		path, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}
		configFile := path + "/../config/HttpClient.config"

		if httpClientConfig, err := config.Get[config.HttpClient](configFile); err != nil {
			t.Fatal(err)
		} else {
			defer file.Remove(httpClientConfig.Log.File.Name + "." + httpClientConfig.Log.File.ExtensionName)
		}

		os.Args = []string{"test", "-config_file=" + configFile}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		main()
	}
}
