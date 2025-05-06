package main

import (
	net_http "net/http"

	"github.com/base-client/go/common/config"
	"github.com/base-client/go/common/main_sub"
	"github.com/common-library/go/http"
	"github.com/common-library/go/log/slog"
)

func main() {
	job := func(log *slog.Log) error {
		if response, err := http.Request("http://"+config.Get("http.address").(string)+"/v1/test/id-01", net_http.MethodGet, map[string][]string{"header-1": {"value-1"}}, "", 3, "", "", nil); err != nil {
			return err
		} else {
			log.Info("result", "response", response)

			return nil
		}
	}

	if err := (&main_sub.Main{}).RunWithSlog(main_sub.Http, job); err != nil {
		panic(err)
	}
}
