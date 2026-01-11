package main

import (
	"github.com/base-client/go/common/config"
	"github.com/base-client/go/common/main_sub"
	"github.com/common-library/go/log/slog"
	"github.com/common-library/go/socket/tcp"
)

func main() {
	job := func(log *slog.Log) error {
		var client tcp.Client
		defer client.Close()

		if err := client.Connect("tcp", config.Get("socket.tcp.address").(string)); err != nil {
			return err
		}

		if readData, err := client.Read(1024); err != nil {
			return err
		} else {
			log.Info("read", "data", readData)
		}

		writeData := "test"
		if _, err := client.Write(writeData); err != nil {
			return err
		} else {
			log.Info("write", "data", writeData)
		}

		if readData, err := client.Read(1024); err != nil {
			return err
		} else {
			log.Info("read", "data", readData)
		}

		return nil
	}

	if err := (&main_sub.Main{}).RunWithSlog(main_sub.Socket, job); err != nil {
		panic(err)
	}
}
