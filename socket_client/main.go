package main

import (
	"errors"
	"flag"
	"github.com/heaven-chp/base-client-go/config"
	"github.com/heaven-chp/common-library-go/json"
	"github.com/heaven-chp/common-library-go/log"
	"github.com/heaven-chp/common-library-go/socket"
)

type Main struct {
	configFile string

	socketClientConfig config.SocketClient
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
	return json.ToStructFromFile(main.configFile, &main.socketClientConfig)
}

func (main *Main) initializeLog() error {
	level, err := log.ToIntLevel(main.socketClientConfig.LogLevel)
	if err != nil {
		return err
	}

	return log.Initialize(level, main.socketClientConfig.LogOutputPath, main.socketClientConfig.LogFileNamePrefix)
}

func (main *Main) finalizeLog() {
	defer log.Finalize()

	log.Flush()
}

func (main *Main) Run() {
	var client socket.Client
	defer client.Close()

	err := client.Connect("tcp", main.socketClientConfig.Address)
	if err != nil {
		log.Error(err.Error())
		return
	}

	readData, err := client.Read(1024)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("read : (%s)", readData)

	writeData := "test"
	_, err = client.Write(writeData)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("write : (%s)", writeData)

	readData, err = client.Read(1024)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("read : (%s)", readData)
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
