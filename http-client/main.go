package main

import (
	"errors"
	"flag"
	net_http "net/http"
	"sync"

	"github.com/heaven-chp/base-client-go/config"
	command_line_argument "github.com/heaven-chp/common-library-go/command-line-argument"
	"github.com/heaven-chp/common-library-go/http"
	log "github.com/heaven-chp/common-library-go/log/file"
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
	httpClientConfig config.HttpClient
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
		{FlagName: "config_file", Usage: "config/http_client.config", DefaultValue: string("")},
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
	return config.Parsing(&this.httpClientConfig, command_line_argument.Get("config_file").(string))
}

func (this *Main) initializeLog() error {
	return log_instance().Initialize(log.Setting{
		Level:           this.httpClientConfig.Log.Level,
		OutputPath:      this.httpClientConfig.Log.OutputPath,
		FileNamePrefix:  this.httpClientConfig.Log.FileNamePrefix,
		PrintCallerInfo: this.httpClientConfig.Log.PrintCallerInfo,
		ChannelSize:     this.httpClientConfig.Log.ChannelSize})

}

func (this *Main) finalizeLog() error {
	return log_instance().Finalize()
}

func (this *Main) Run() error {
	err := this.initialize()
	if err != nil {
		return err
	}
	defer this.finalize()

	log_instance().Info("process start")
	defer log_instance().Info("process end")

	response, err := http.Request("http://"+this.httpClientConfig.Address+"/v1/test/id-01", net_http.MethodGet, map[string][]string{"header-1": {"value-1"}}, "", 3, "", "")
	if err != nil {
		return err
	}

	log_instance().Infof("%#v\n", response)

	return nil
}

func main() {
	var main Main
	err := main.Run()
	if err != nil {
		panic(err)
	}
}
