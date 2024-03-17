package main

import (
	"errors"
	"flag"

	"github.com/heaven-chp/base-client-go/config"
	"github.com/heaven-chp/base-client-go/long-polling-client/log"
	command_line_flag "github.com/heaven-chp/common-library-go/command-line/flag"
	long_polling "github.com/heaven-chp/common-library-go/long-polling"
)

type Main struct {
	longPollingClientConfig config.LongPollingClient
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
		{FlagName: "config_file", Usage: "config/SocketClient.config", DefaultValue: string("")},
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

	if longPollingClientConfig, err := config.Get[config.LongPollingClient](fileName); err != nil {
		return err
	} else {
		this.longPollingClientConfig = longPollingClientConfig
		return nil
	}
}

func (this *Main) initializeLog() error {
	log.Initialize(this.longPollingClientConfig)

	return nil
}

func (this *Main) finalizeLog() error {
	log.Client.Flush()

	return nil
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
	err := this.initialize()
	if err != nil {
		return err
	}
	defer func() {
		err := this.finalize()
		if err != nil {
			log.Client.Error(err.Error())
		}
	}()

	log.Client.Info("process start")
	defer log.Client.Info("process end")

	const category = "category-1"
	const data = "data-1"

	err = this.publish(category, data)
	if err != nil {
		return err
	}

	err = this.subscription(category)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	main := Main{}

	err := main.Run()
	if err != nil {
		panic(err)
	}
}
