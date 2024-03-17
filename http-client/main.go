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
	if err := this.parseFlag(); err != nil {
		return err
	} else if err := this.setConfig(); err != nil {
		return err
	} else {
		log.Initialize(this.httpClientConfig)

		return nil
	}
}

func (this *Main) parseFlag() error {
	flagInfos := []command_line_flag.FlagInfo{
		{FlagName: "config_file", Usage: "config/http_client.config", DefaultValue: string("")},
	}

	if err := command_line_flag.Parse(flagInfos); err != nil {
		return nil
	} else if flag.NFlag() != 1 {
		flag.Usage()
		return errors.New("invalid flag")
	} else {
		return nil
	}
}

func (this *Main) setConfig() error {
	fileName := command_line_flag.Get[string]("config_file")

	if httpClientConfig, err := config.Get[config.HttpClient](fileName); err != nil {
		return err
	} else {
		this.httpClientConfig = httpClientConfig
		return nil
	}
}

func (this *Main) Run() error {
	defer log.Client.Flush()

	if err := this.initialize(); err != nil {
		return err
	}

	log.Client.Info("process start")
	defer log.Client.Info("process end")

	if response, err := http.Request("http://"+this.httpClientConfig.Address+"/v1/test/id-01", net_http.MethodGet, map[string][]string{"header-1": {"value-1"}}, "", 3, "", ""); err != nil {
		return err
	} else {
		log.Client.Info("result", "response", response)

		return nil
	}
}

func main() {
	if err := (&Main{}).Run(); err != nil {
		log.Client.Error(err.Error())
		log.Client.Flush()
	}
}
