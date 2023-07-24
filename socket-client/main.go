package main

import (
	"errors"
	"flag"

	"github.com/heaven-chp/base-client-go/config"
	"github.com/heaven-chp/common-library-go/log"
	"github.com/heaven-chp/common-library-go/socket"
)

type Main struct {
	configFile string

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
	return config.Parsing(&this.socketClientConfig, this.configFile)
}

func (this *Main) initializeLog() error {
	level, err := log.ToIntLevel(this.socketClientConfig.LogLevel)
	if err != nil {
		return err
	}

	return log.Initialize(level, this.socketClientConfig.LogOutputPath, this.socketClientConfig.LogFileNamePrefix)
}

func (this *Main) finalizeLog() error {
	log.Flush()

	return log.Finalize()
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
	log.Info("read : (%s)", readData)

	writeData := "test"
	_, err = client.Write(writeData)
	if err != nil {
		return err
	}
	log.Info("write : (%s)", writeData)

	readData, err = client.Read(1024)
	if err != nil {
		return err
	}
	log.Info("read : (%s)", readData)

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
