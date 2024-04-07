package main

import (
	"errors"
	"flag"

	"github.com/base-client/go/config"
	"github.com/base-client/go/long-polling-client/log"
	"github.com/common-library/go/command-line/flags"
	long_polling "github.com/common-library/go/long-polling"
)

type Main struct {
	longPollingClientConfig config.LongPollingClient
}

func (this *Main) initialize() error {
	if err := this.parseFlag(); err != nil {
		return err
	} else if err := this.setConfig(); err != nil {
		return err
	} else {
		log.Initialize(this.longPollingClientConfig)

		return nil
	}
}

func (this *Main) parseFlag() error {
	flagInfos := []flags.FlagInfo{
		{FlagName: "config_file", Usage: "config/SocketClient.config", DefaultValue: string("")},
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

	if longPollingClientConfig, err := config.Get[config.LongPollingClient](fileName); err != nil {
		return err
	} else {
		this.longPollingClientConfig = longPollingClientConfig
		return nil
	}
}

func (this *Main) subscription(category string) error {
	request := long_polling.SubscriptionRequest{Category: category, Timeout: 300, SinceTime: 1}
	response, err := long_polling.Subscription("http://"+this.longPollingClientConfig.Address+this.longPollingClientConfig.SubscriptionURI, nil, request, "", "")
	if err != nil {
		return err
	}

	log.Client.Info("subscription", "response", response)

	return nil
}

func (this *Main) publish(category, data string) error {
	request := long_polling.PublishRequest{Category: category, Data: data}
	response, err := long_polling.Publish("http://"+this.longPollingClientConfig.Address+this.longPollingClientConfig.PublishURI, 10, nil, request, "", "")
	if err != nil {
		return err
	}

	log.Client.Info("publish", "response", response)

	return nil
}

func (this *Main) Run() error {
	defer log.Client.Flush()

	if err := this.initialize(); err != nil {
		return err
	}

	log.Client.Info("process start")
	defer log.Client.Info("process end")

	const category = "category-1"
	const data = "data-1"

	if err := this.publish(category, data); err != nil {
		return err
	} else if err := this.subscription(category); err != nil {
		return err
	} else {
		return nil
	}
}

func main() {
	if err := (&Main{}).Run(); err != nil {
		log.Client.Error(err.Error())
		log.Client.Flush()
	}
}
