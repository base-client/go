package config_test

import (
	"testing"

	"github.com/base-client/go/common/config"
)

func TestRead(t *testing.T) {
	if err := config.Read("./config.yaml"); err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	answer := map[string]any{
		"cloudEvents.address": ":40000",

		"gRPC.address":                ":50051",
		"gRPC.timeout":                "3s",
		"gRPC.log.level":              "debug",
		"gRPC.log.output":             "file",
		"gRPC.log.file.name":          "./grpc-client",
		"gRPC.log.file.extensionName": "log",
		"gRPC.log.file.addDate":       false,
		"gRPC.log.withCallerInfo":     true,

		"http.address":                ":10000",
		"http.log.level":              "debug",
		"http.log.output":             "file",
		"http.log.file.name":          "./http-client",
		"http.log.file.extensionName": "log",
		"http.log.file.addDate":       false,
		"http.log.withCallerInfo":     true,

		"longPolling.address":                                   ":30000",
		"longPolling.subscriptionURI":                           "/subscription",
		"longPolling.publishURI":                                "/publish",
		"longPolling.filePersistorInfo.use":                     false,
		"longPolling.filePersistorInfo.fileName":                "./file-persistor.txt",
		"longPolling.filePersistorInfo.writeBufferSize":         250,
		"longPolling.filePersistorInfo.writeFlushPeriodSeconds": 1,
		"longPolling.log.level":                                 "debug",
		"longPolling.log.output":                                "file",
		"longPolling.log.file.name":                             "./long-polling-client",
		"longPolling.log.file.extensionName":                    "log",
		"longPolling.log.file.addDate":                          false,
		"longPolling.log.withCallerInfo":                        true,

		"socket.tcp.address":            ":20000",
		"socket.udp.address":            ":20001",
		"socket.log.level":              "debug",
		"socket.log.output":             "file",
		"socket.log.file.name":          "./socket-client",
		"socket.log.file.extensionName": "log",
		"socket.log.file.addDate":       false,
		"socket.log.withCallerInfo":     true,
	}

	if err := config.Read("./config.yaml"); err != nil {
		t.Fatal(err)
	}

	for key, value := range answer {
		if result := config.Get(key); result != value {
			t.Fatal(key, ",", value, ",", result)
		}
	}
}
