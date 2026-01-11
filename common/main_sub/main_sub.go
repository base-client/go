package main_sub

import (
	"errors"
	"flag"

	"github.com/base-client/go/common/config"
	"github.com/base-client/go/common/log"
	"github.com/common-library/go/command-line/flags"
	"github.com/common-library/go/log/klog"
	"github.com/common-library/go/log/slog"
)

type ClientKind string

const (
	CloudEvents ClientKind = "cloudEvents"
	GRPC        ClientKind = "gRPC"
	Http        ClientKind = "http"
	LongPolling ClientKind = "longPolling"
	Socket      ClientKind = "socket"
)

type Main struct {
	clientKind ClientKind
}

func (m *Main) initialize(clientKind ClientKind) error {
	if err := m.parseFlag(); err != nil {
		return err
	} else if err := config.Read(flags.Get[string]("config-file")); err != nil {
		return err
	} else {
		m.clientKind = clientKind

		if m.clientKind != CloudEvents {
			log.Initialize(string(clientKind))
		}

		return nil
	}
}

func (m *Main) parseFlag() error {
	flagInfos := []flags.FlagInfo{
		{FlagName: "config-file", Usage: "config/config.yaml", DefaultValue: string("")},
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

func (m *Main) RunWithKlog(clientKind ClientKind, job func() error) error {
	defer klog.Flush()

	if err := m.initialize(clientKind); err != nil {
		return err
	} else {
		return job()
	}
}

func (m *Main) RunWithSlog(clientKind ClientKind, job func(*slog.Log) error) error {
	defer log.Log.Flush()

	if err := m.initialize(clientKind); err != nil {
		return err
	} else {
		return job(&log.Log)
	}
}
