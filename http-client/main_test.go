package main

import (
	"flag"
	net_http "net/http"
	"os"
	"testing"
	"time"

	"github.com/heaven-chp/common-library-go/http"
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
	server := http.Server{}
	server.AddHandler("/v1/test/{id}", net_http.MethodGet, func(responseWriter net_http.ResponseWriter, request *net_http.Request) {
		responseWriter.WriteHeader(net_http.StatusOK)
		responseWriter.Write([]byte(`{"field_1":"value-1"}`))
	})

	middlewareFunction := func(nextHandler net_http.Handler) net_http.Handler {
		return net_http.HandlerFunc(func(responseWriter net_http.ResponseWriter, request *net_http.Request) {
			nextHandler.ServeHTTP(responseWriter, request)
		})
	}

	err := server.Start(":10000", func(err error) { t.Error(err) }, middlewareFunction)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Duration(200) * time.Millisecond)

	defer func() {
		err := server.Stop(5)
		if err != nil {
			t.Error(err)
		}
	}()

	{
		path, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}

		os.Args = []string{"test", "-config_file=" + path + "/../config/HttpClient.config"}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		main()
	}
}
