package main

import (
	"time"

	"github.com/base-client/go/common/config"
	"github.com/base-client/go/common/main_sub"
	"github.com/common-library/go/log/slog"
	"github.com/common-library/go/socket/udp"
)

func main() {
	job := func(log *slog.Log) error {
		var client udp.Client
		defer client.Close()

		if err := client.Connect("udp", config.Get("socket.udp.address").(string)); err != nil {
			return err
		}

		// UDP는 연결 없는 프로토콜이므로 클라이언트가 먼저 데이터를 보내야 함
		writeData := []byte("test")
		if _, err := client.Send(writeData); err != nil {
			return err
		} else {
			log.Info("send", "data", string(writeData))
		}

		// 서버로부터 greeting 수신
		if readData, _, err := client.Receive(1024, 5*time.Second); err != nil {
			return err
		} else {
			log.Info("receive", "data", string(readData))
		}

		// 서버로부터 response 수신
		if readData, _, err := client.Receive(1024, 5*time.Second); err != nil {
			return err
		} else {
			log.Info("receive", "data", string(readData))
		}

		return nil
	}

	if err := (&main_sub.Main{}).RunWithSlog(main_sub.Socket, job); err != nil {
		panic(err)
	}
}
