package main

import (
	"context"
	"time"

	"github.com/base-client/go/common/config"
	"github.com/base-client/go/common/main_sub"
	"github.com/common-library/go/grpc"
	"github.com/common-library/go/grpc/sample"
	"github.com/common-library/go/log/slog"
)

func main() {
	job := func(log *slog.Log) error {
		connection, err := grpc.GetConnection(config.Get("gRPC.address").(string))
		if err != nil {
			return err
		}
		defer connection.Close()

		client := sample.NewSampleClient(connection)

		timeout := time.Duration(-1)
		if duration, err := time.ParseDuration(config.Get("gRPC.timeout").(string)); err != nil {
			return err
		} else {
			timeout = duration
		}

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if reply, err := client.Func1(ctx, &sample.Request{Data1: 1, Data2: "abc"}); err != nil {
			return err
		} else {
			log.Info("reply", "Data1", reply.Data1, "Data2", reply.Data2)

			return nil
		}
	}

	if err := (&main_sub.Main{}).RunWithSlog(main_sub.GRPC, job); err != nil {
		panic(err)
	}
}
