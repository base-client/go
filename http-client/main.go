package main

import (
	"errors"
	"flag"
	net_http "net/http"

	"github.com/base-client/go/config"
	"github.com/base-client/go/http-client/log"
	"github.com/common-library/go/command-line/flags"
	"github.com/common-library/go/http"
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
	flagInfos := []flags.FlagInfo{
		{FlagName: "config_file", Usage: "config/http_client.config", DefaultValue: string("")},
	}

	if err := flags.Parse(flagInfos); err != nil {
		flag.Usage()
		return err
	} else if flag.NFlag() != 1 {
		flag.Usage()
		return errors.New("invalid flag")
	} else {
		return nil
	}
}

func (this *Main) setConfig() error {
	fileName := flags.Get[string]("config_file")

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
