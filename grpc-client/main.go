package main

import (
	"context"
	"errors"
	"flag"
	"time"

	"github.com/base-client/go/config"
	"github.com/base-client/go/grpc-client/log"
	"github.com/common-library/go/command-line/flags"
	"github.com/common-library/go/grpc"
	"github.com/common-library/go/grpc/sample"
)

type Main struct {
	grpcClientConfig config.GrpcClient
}

func (this *Main) initialize() error {
	if err := this.parseFlag(); err != nil {
		return err
	} else if err := this.setConfig(); err != nil {
		return err
	} else {
		log.Initialize(this.grpcClientConfig)

		return nil
	}
}

func (this *Main) parseFlag() error {
	flagInfos := []flags.FlagInfo{
		{FlagName: "config_file", Usage: "config/GrpcClient.config", DefaultValue: string("")},
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

	if grpcClientConfig, err := config.Get[config.GrpcClient](fileName); err != nil {
		return err
	} else {
		this.grpcClientConfig = grpcClientConfig
		return nil
	}
}

func (this *Main) job() error {
	connection, err := grpc.GetConnection(this.grpcClientConfig.Address)
	if err != nil {
		return err
	}
	defer connection.Close()

	client := sample.NewSampleClient(connection)

	timeout := this.grpcClientConfig.Timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	if reply, err := client.Func1(ctx, &sample.Request{Data1: 1, Data2: "abc"}); err != nil {
		return err
	} else {
		log.Client.Info("reply", "Data1", reply.Data1, "Data2", reply.Data2)

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

	return this.job()
}

func main() {
	if err := (&Main{}).Run(); err != nil {
		log.Client.Error(err.Error())
		log.Client.Flush()
	}
}
