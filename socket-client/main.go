package main

import (
	"errors"
	"flag"
	"sync"

	"github.com/heaven-chp/base-client-go/config"
	command_line_argument "github.com/heaven-chp/common-library-go/command-line-argument"
	log "github.com/heaven-chp/common-library-go/log/file"
	"github.com/heaven-chp/common-library-go/socket"
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
	socketClientConfig config.SocketClient
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
	return config.Parsing(&this.socketClientConfig, command_line_argument.Get("config_file").(string))
}

func (this *Main) initializeLog() error {
	return log_instance().Initialize(log.Setting{
		Level:           this.socketClientConfig.Log.Level,
		OutputPath:      this.socketClientConfig.Log.OutputPath,
		FileNamePrefix:  this.socketClientConfig.Log.FileNamePrefix,
		PrintCallerInfo: this.socketClientConfig.Log.PrintCallerInfo,
		ChannelSize:     this.socketClientConfig.Log.ChannelSize})
}

func (this *Main) finalizeLog() error {
	return log_instance().Finalize()
}

func (this *Main) job() error {
	var client socket.Client
	defer client.Close()

	err := client.Connect("tcp", this.socketClientConfig.Address)
	if err != nil {
		return err
	}

	readData, err := client.Read(1024)
	if err != nil {
		return err
	}
	log_instance().Infof("read : (%s)", readData)

	writeData := "test"
	_, err = client.Write(writeData)
	if err != nil {
		return err
	}
	log_instance().Infof("write : (%s)", writeData)

	readData, err = client.Read(1024)
	if err != nil {
		return err
	}
	log_instance().Infof("read : (%s)", readData)

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
