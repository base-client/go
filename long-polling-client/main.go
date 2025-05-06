package main

import (
	"github.com/base-client/go/common/config"
	"github.com/base-client/go/common/main_sub"
	"github.com/common-library/go/log/slog"
	long_polling "github.com/common-library/go/long-polling"
)

func main() {
	job := func(log *slog.Log) error {
		subscription := func(category string) error {
			address := config.Get("longPolling.address").(string)
			subscriptionURI := config.Get("longPolling.subscriptionURI").(string)

			request := long_polling.SubscriptionRequest{Category: category, TimeoutSeconds: 300, SinceTime: 1}
			if response, err := long_polling.Subscription("http://"+address+subscriptionURI, nil, request, "", "", nil); err != nil {
				return err
			} else {
				log.Info("subscription", "response", response)

				return nil
			}
		}

		publish := func(category, data string) error {
			address := config.Get("longPolling.address").(string)
			publishURI := config.Get("longPolling.publishURI").(string)

			request := long_polling.PublishRequest{Category: category, Data: data}
			if response, err := long_polling.Publish("http://"+address+publishURI, 10, nil, request, "", "", nil); err != nil {
				return err
			} else {
				log.Info("publish", "response", response)

				return nil
			}
		}

		const category = "category-1"
		const data = "data-1"
		if err := publish(category, data); err != nil {
			return err
		} else {
			return subscription(category)
		}
	}

	if err := (&main_sub.Main{}).RunWithSlog(main_sub.LongPolling, job); err != nil {
		panic(err)
	}
}
