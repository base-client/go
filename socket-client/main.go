package main

import (
	"errors"
	"flag"

	"github.com/heaven-chp/base-client-go/config"
	"github.com/heaven-chp/base-client-go/socket-client/log"
	command_line_flag "github.com/heaven-chp/common-library-go/command-line/flag"
	"github.com/heaven-chp/common-library-go/socket"
)

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
	err := command_line_flag.Parse([]command_line_flag.FlagInfo{
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
	fileName := command_line_flag.Get[string]("config_file")

	if socketClientConfig, err := config.Get[config.SocketClient](fileName); err != nil {
		return err
	} else {
		this.socketClientConfig = socketClientConfig
		return nil
	}
}

func (this *Main) initializeLog() error {
	log.Initialize(this.socketClientConfig)

	return nil
}

func (this *Main) finalizeLog() error {
	log.Client.Flush()

	return nil
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
	log.Client.Info("read", "data", readData)

	writeData := "test"
	_, err = client.Write(writeData)
	if err != nil {
		return err
	}
	log.Client.Info("write", "data", writeData)

	readData, err = client.Read(1024)
	if err != nil {
		return err
	}
	log.Client.Info("read", "data", readData)

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
