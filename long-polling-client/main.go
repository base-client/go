package main

import (
	"errors"
	"flag"
	"sync"

	"github.com/heaven-chp/base-client-go/config"
	command_line_argument "github.com/heaven-chp/common-library-go/command-line-argument"
	log "github.com/heaven-chp/common-library-go/log/file"
	long_polling "github.com/heaven-chp/common-library-go/long-polling"
)

var onceForLog sync.Once
var fileLog *log.FileLog

func log_instance() *log.FileLog {
	onceForLog.Do(func() {
		fileLog = &log.FileLog{}
	})

	return fileLog
}

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
	err := command_line_argument.Set([]command_line_argument.CommandLineArgumentInfo{
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
	return config.Parsing(&this.longPollingClientConfig, command_line_argument.Get("config_file").(string))
}

func (this *Main) initializeLog() error {
	return log_instance().Initialize(log.Setting{
		Level:           this.longPollingClientConfig.Log.Level,
		OutputPath:      this.longPollingClientConfig.Log.OutputPath,
		FileNamePrefix:  this.longPollingClientConfig.Log.FileNamePrefix,
		PrintCallerInfo: this.longPollingClientConfig.Log.PrintCallerInfo,
		ChannelSize:     this.longPollingClientConfig.Log.ChannelSize})
}

func (this *Main) finalizeLog() error {
	return log_instance().Finalize()
}

func (this *Main) subscription(category string) error {
	request := long_polling.SubscriptionRequest{Category: category, Timeout: 300, SinceTime: 1}
	response, err := long_polling.Subscription("http://"+this.longPollingClientConfig.Address+this.longPollingClientConfig.SubscriptionURI, nil, request, "", "")
	if err != nil {
		return err
	}

	log_instance().Infof("subscription response : (%#v)", response)

	return nil
}

func (this *Main) publish(category, data string) error {
	request := long_polling.PublishRequest{Category: category, Data: data}
	response, err := long_polling.Publish("http://"+this.longPollingClientConfig.Address+this.longPollingClientConfig.PublishURI, 10, nil, request, "", "")
	if err != nil {
		return err
	}

	log_instance().Infof("publish response : (%#v)", response)

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
			log_instance().Error(err)
		}
	}()

	log_instance().Info("process start")
	defer log_instance().Info("process end")

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
