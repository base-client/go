package main

import (
	"errors"

	"github.com/base-client/go/common/config"
	"github.com/base-client/go/common/main_sub"
	"github.com/common-library/go/event/cloudevents"
	"github.com/common-library/go/log/klog"
)

func main() {
	job := func() error {
		return nil

		event := cloudevents.NewEvent()
		event.SetID("id")
		event.SetType("type")
		event.SetSource("source")

		address := "http://" + config.Get("cloudEvents.address").(string)
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

	if err := (&main_sub.Main{}).RunWithKlog(main_sub.CloudEvents, job); err != nil {
		panic(err)
	}
}
