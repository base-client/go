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

func (main *Main) Initialize() error {
	err := main.initializeFlag()
	if err != nil {
		return err
	}

	err = main.initializeConfig()
	if err != nil {
		return err
	}

	err = main.initializeLog()
	if err != nil {
		return err
	}

	return nil
}

func (main *Main) Finalize() {
	defer main.finalizeLog()
}

func (main *Main) initializeFlag() error {
	configFile := flag.String("config_file", "", "config file")
	flag.Parse()

	if flag.NFlag() != 1 {
		flag.Usage()
		return errors.New("invalid flag")
	}

	main.configFile = *configFile

	return nil
}

func (main *Main) initializeConfig() error {
	return config.Parsing(&main.grpcClientConfig, main.configFile)
}

func (main *Main) initializeLog() error {
	level, err := log.ToIntLevel(main.grpcClientConfig.LogLevel)
	if err != nil {
		return err
	}

	return log.Initialize(level, main.grpcClientConfig.LogOutputPath, main.grpcClientConfig.LogFileNamePrefix)
}

func (main *Main) finalizeLog() {
	defer log.Finalize()

	log.Flush()
}

func (main *Main) Run() {
	timeout := main.grpcClientConfig.Timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	connection, err := grpc.GetConnection(main.grpcClientConfig.Address)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer connection.Close()

	client := sample.NewSampleClient(connection)

	request := sample.Request{Data1: 1, Data2: "abc"}
	reply, err := client.Func(ctx, &request)
	if err != nil {
		log.Error("Func fail - error : %s", err.Error())
		return
	}

	log.Info("request - Data1 : (%d), Data2 : (%s)", request.Data1, request.Data2)
	log.Info("reply - Data1 : (%d), Data2 : (%s)", reply.Data1, reply.Data2)
}

func main() {
	var main Main
	err := main.Initialize()
	if err != nil {
		log.Error("Initialize fail - error : (%s)", err.Error())
		return
	}
	defer main.Finalize()

	log.Info("process start")
	defer log.Info("process end")

	main.Run()
}
