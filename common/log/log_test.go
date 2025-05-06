package log_test

import (
	"strings"
	"testing"

	"github.com/base-client/go/common/config"
	"github.com/base-client/go/common/log"
	"github.com/common-library/go/file"
)

func Test(t *testing.T) {
	const configFile = "../config/config.yaml"

	if err := config.Read(configFile); err != nil {
		t.Fatal(err)
	}

	kind := "gRPC"
	fileName := config.Get(kind+".log.file.name").(string) + "." + config.Get(kind+".log.file.extensionName").(string)

	log.Initialize(kind)
	defer file.Remove(fileName)

	content := "test"
	log.Log.Info(content)
	log.Log.Flush()

	if data, err := file.Read(fileName); err != nil {
		t.Fatal(err)
	} else if strings.Contains(data, `"msg":"`+content+`"`) == false {
		t.Fatal("invalid :", data, ",", content)
	}
}
