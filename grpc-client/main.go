package main

import (
	"context"
	"errors"
	"flag"
	"sync"
	"time"

	"github.com/heaven-chp/base-client-go/config"
	command_line_argument "github.com/heaven-chp/common-library-go/command-line-argument"
	"github.com/heaven-chp/common-library-go/grpc"
	"github.com/heaven-chp/common-library-go/grpc/sample"
	log "github.com/heaven-chp/common-library-go/log/file"
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
	err := command_line_argument.Set([]command_line_argument.CommandLineArgumentInfo{
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
	return config.Parsing(&this.grpcClientConfig, command_line_argument.Get("config_file").(string))
}

func (this *Main) initializeLog() error {
	return log_instance().Initialize(log.Setting{
		Level:           this.grpcClientConfig.Log.Level,
		OutputPath:      this.grpcClientConfig.Log.OutputPath,
		FileNamePrefix:  this.grpcClientConfig.Log.FileNamePrefix,
		PrintCallerInfo: this.grpcClientConfig.Log.PrintCallerInfo,
		ChannelSize:     this.grpcClientConfig.Log.ChannelSize})

}

func (this *Main) finalizeLog() error {
	return log_instance().Finalize()
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

	log_instance().Infof("reply - Data1 : (%d), Data2 : (%s)", reply.Data1, reply.Data2)

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

	return this.job()
}

func main() {
	main := Main{}

	err := main.Run()
	if err != nil {
		panic(err)
	}
}
