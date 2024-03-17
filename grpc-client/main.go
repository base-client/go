package main

import (
	"context"
	"errors"
	"flag"
	"time"

	"github.com/heaven-chp/base-client-go/config"
	"github.com/heaven-chp/base-client-go/grpc-client/log"
	command_line_flag "github.com/heaven-chp/common-library-go/command-line/flag"
	"github.com/heaven-chp/common-library-go/grpc"
	"github.com/heaven-chp/common-library-go/grpc/sample"
)

type Main struct {
	grpcClientConfig config.GrpcClient
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
		{FlagName: "config_file", Usage: "config/GrpcClient.config", DefaultValue: string("")},
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

	if grpcClientConfig, err := config.Get[config.GrpcClient](fileName); err != nil {
		return err
	} else {
		this.grpcClientConfig = grpcClientConfig
		return nil
	}
}

func (this *Main) initializeLog() error {
	log.Initialize(this.grpcClientConfig)

	return nil
}

func (this *Main) finalizeLog() error {
	log.Client.Flush()

	return nil
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
	reply, err := client.Func1(ctx, &sample.Request{Data1: 1, Data2: "abc"})
	if err != nil {
		return err
	}

	log.Client.Info("reply", "Data1", reply.Data1, "Data2", reply.Data2)

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

	return this.job()
}

func main() {
	main := Main{}

	err := main.Run()
	if err != nil {
		panic(err)
	}
}
