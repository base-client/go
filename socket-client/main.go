package main

import (
	"errors"
	"flag"

	"github.com/heaven-chp/base-client-go/config"
	"github.com/heaven-chp/base-client-go/socket-client/log"
	"github.com/heaven-chp/common-library-go/command-line/flags"
	"github.com/heaven-chp/common-library-go/socket"
)

type Main struct {
	socketClientConfig config.SocketClient
}

func (this *Main) initialize() error {
	if err := this.parseFlag(); err != nil {
		return err
	} else if err := this.setConfig(); err != nil {
		return err
	} else {
		log.Initialize(this.socketClientConfig)

		return nil
	}
}

func (this *Main) parseFlag() error {
	flagInfos := []flags.FlagInfo{
		{FlagName: "config_file", Usage: "config/SocketClient.config", DefaultValue: string("")},
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

	if socketClientConfig, err := config.Get[config.SocketClient](fileName); err != nil {
		return err
	} else {
		this.socketClientConfig = socketClientConfig
		return nil
	}
}

func (this *Main) job() error {
	var client socket.Client
	defer client.Close()

	if err := client.Connect("tcp", this.socketClientConfig.Address); err != nil {
		return err
	}

	if readData, err := client.Read(1024); err != nil {
		return err
	} else {
		log.Client.Info("read", "data", readData)
	}

	writeData := "test"
	if _, err := client.Write(writeData); err != nil {
		return err
	} else {
		log.Client.Info("write", "data", writeData)
	}

	if readData, err := client.Read(1024); err != nil {
		return err
	} else {
		log.Client.Info("read", "data", readData)
	}

	return nil
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
