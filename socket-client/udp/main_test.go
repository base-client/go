package main

import (
	"flag"
	"net"
	"os"
	"testing"
	"time"

	"github.com/base-client/go/common/config"
	"github.com/common-library/go/file"
	"github.com/common-library/go/socket/udp"
)

type TestServer struct {
	Network          string
	Address          string
	Greeting         string
	PrefixOfResponse string

	server udp.Server
}

func (ts *TestServer) Start(t *testing.T) {
	ts.Network = "udp"
	ts.Address = config.Get("socket.udp.address").(string)
	ts.Greeting = "greeting"
	ts.PrefixOfResponse = "[response] "

	packetHandler := func(data []byte, addr net.Addr, conn net.PacketConn) {
		// Send greeting first
		if _, err := conn.WriteTo([]byte(ts.Greeting), addr); err != nil {
			t.Error(err)
			return
		}

		// Echo back with prefix
		response := ts.PrefixOfResponse + string(data)
		if _, err := conn.WriteTo([]byte(response), addr); err != nil {
			t.Error(err)
		}
	}

	errorHandler := func(err error) {
		t.Error(err)
	}

	if err := ts.server.Start(ts.Network, ts.Address, 1024, packetHandler, false, errorHandler); err != nil {
		t.Fatal(err)
	}

	for !ts.server.IsRunning() {
		time.Sleep(100 * time.Millisecond)
	}
}

func (ts *TestServer) Stop(t *testing.T) {
	if err := ts.server.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestMain(t *testing.T) {
	const configFile = "../../common/config/config.yaml"

	if err := config.Read(configFile); err != nil {
		t.Fatal(err)
	}
	defer file.Remove(config.Get("socket.log.file.name").(string) + "." + config.Get("socket.log.file.extensionName").(string))

	testServer := TestServer{}
	testServer.Start(t)
	defer testServer.Stop(t)

	os.Args = []string{"test", "-config-file=" + configFile}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	main()
}
