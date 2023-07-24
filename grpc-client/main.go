package main

import (
	"context"
	"errors"
	"flag"
	"time"

	"github.com/heaven-chp/base-client-go/config"
	"github.com/heaven-chp/common-library-go/grpc"
	"github.com/heaven-chp/common-library-go/grpc/sample"
	"github.com/heaven-chp/common-library-go/log"
)

type Main struct {
	configFile string

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
	configFile := flag.String("config_file", "", "config file")
	flag.Parse()

	if flag.NFlag() != 1 {
		flag.Usage()
		return errors.New("invalid flag")
	}

	this.configFile = *configFile

	return nil
}

func (this *Main) initializeConfig() error {
	return config.Parsing(&this.grpcClientConfig, this.configFile)
}

func (this *Main) initializeLog() error {
	level, err := log.ToIntLevel(this.grpcClientConfig.LogLevel)
	if err != nil {
		return err
	}

	return log.Initialize(level, this.grpcClientConfig.LogOutputPath, this.grpcClientConfig.LogFileNamePrefix)
}

func (this *Main) finalizeLog() error {
	log.Flush()

	return log.Finalize()
}

func (this *Main) job() error {
	timeout := this.grpcClientConfig.Timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	connection, err := grpc.GetConnection(this.grpcClientConfig.Address)
	if err != nil {
		return err
	}
	defer connection.Close()

	client := sample.NewSampleClient(connection)

	request := sample.Request{Data1: 1, Data2: "abc"}
	reply, err := client.Func(ctx, &request)
	if err != nil {
		return err
	}

	log.Info("request - Data1 : (%d), Data2 : (%s)", request.Data1, request.Data2)
	log.Info("reply - Data1 : (%d), Data2 : (%s)", reply.Data1, reply.Data2)

	return nil
}

func (this *Main) Run() error {
	err := this.initialize()
	if err != nil {
		return err
	}
	defer this.finalize()

	log.Info("process start")
	defer log.Info("process end")

	return this.job()
}

func main() {
	main := Main{}

	err := main.Run()
	if err != nil {
		log.Error(err.Error())
	}

}
