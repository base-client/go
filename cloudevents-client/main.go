package main

import (
	"errors"
	"flag"

	"github.com/base-client/go/config"
	"github.com/common-library/go/command-line/flags"
	"github.com/common-library/go/event/cloudevents"
	"github.com/common-library/go/log/klog"
)

type Main struct {
	cloudEventsClientConfig config.CloudEventsClient
}

func (this *Main) initialize() error {
	if err := this.parseFlag(); err != nil {
		return err
	} else if err := this.setConfig(); err != nil {
		return err
	} else {
		return nil
	}
}

func (this *Main) parseFlag() error {
	flagInfos := []flags.FlagInfo{
		{FlagName: "config-file", Usage: "config/CloudEventsClient.config", DefaultValue: string("")},
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
	fileName := flags.Get[string]("config-file")

	if cloudEventsClientConfig, err := config.Get[config.CloudEventsClient](fileName); err != nil {
		return err
	} else {
		this.cloudEventsClientConfig = cloudEventsClientConfig
		return nil
	}
}

func (this *Main) Run() error {
	if err := this.initialize(); err != nil {
		return err
	}

	event := cloudevents.NewEvent()
	event.SetID("id")
	event.SetType("type")
	event.SetSource("source")

	address := this.cloudEventsClientConfig.Address

	if client, err := cloudevents.NewHttp(address, nil, nil); err != nil {
		return err
	} else if result := client.Send(event); result.IsUndelivered() {
		return errors.New(result.Error())
	} else if event, result := client.Request(event); result.IsUndelivered() {
		return errors.New(result.Error())
	} else {
		if event != nil {
			klog.InfoS("response", "event", event.String())
		}

		return nil
	}
}

func main() {
	defer klog.Flush()

	klog.InfoS("main start")
	defer klog.InfoS("main end")

	if err := (&Main{}).Run(); err != nil {
		klog.ErrorS(err, "")
	}
}
