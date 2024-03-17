package main

import (
	"errors"
	"flag"
	net_http "net/http"

	"github.com/heaven-chp/base-client-go/config"
	"github.com/heaven-chp/base-client-go/http-client/log"
	command_line_flag "github.com/heaven-chp/common-library-go/command-line/flag"
	"github.com/heaven-chp/common-library-go/http"
)

type Main struct {
	httpClientConfig config.HttpClient
}

func (this *Main) initialize() error {
	err := this.initializeFlag()
	if err != nil {
		return err
	}

	err = this.initializeConfig()
	if err != nil {
		return err
	}

	err = this.initializeLog()
	if err != nil {
		return err
	}

	return nil
}

func (this *Main) finalize() error {
	return this.finalizeLog()
}

func (this *Main) initializeFlag() error {
	err := command_line_flag.Parse([]command_line_flag.FlagInfo{
		{FlagName: "config_file", Usage: "config/http_client.config", DefaultValue: string("")},
	})
	if err != nil {
		return nil
	}

	if flag.NFlag() != 1 {
		flag.Usage()
		return errors.New("invalid flag")
	}

	return nil
}

func (this *Main) initializeConfig() error {
	fileName := command_line_flag.Get[string]("config_file")

	if httpClientConfig, err := config.Get[config.HttpClient](fileName); err != nil {
		return err
	} else {
		this.httpClientConfig = httpClientConfig
		return nil
	}
}

func (this *Main) initializeLog() error {
	log.Initialize(this.httpClientConfig)

	return nil
}

func (this *Main) finalizeLog() error {
	log.Client.Flush()

	return nil
}

func (this *Main) Run() error {
	err := this.initialize()
	if err != nil {
		return err
	}
	defer this.finalize()

	log.Client.Info("process start")
	defer log.Client.Info("process end")

	response, err := http.Request("http://"+this.httpClientConfig.Address+"/v1/test/id-01", net_http.MethodGet, map[string][]string{"header-1": {"value-1"}}, "", 3, "", "")
	if err != nil {
		return err
	}

	log.Client.Info("result", "response", response)

	return nil
}

func main() {
	var main Main
	err := main.Run()
	if err != nil {
		panic(err)
	}
}
