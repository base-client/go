package log

import (
	"strings"

	"github.com/base-client/go/config"
	"github.com/common-library/go/log/slog"
)

var Client slog.Log

func Initialize(clientConfig config.SocketClient) {
	switch strings.ToLower(clientConfig.Log.Level) {
	case "trace":
		Client.SetLevel(slog.LevelTrace)
	case "debug":
		Client.SetLevel(slog.LevelDebug)
	case "info":
		Client.SetLevel(slog.LevelInfo)
	case "warn":
		Client.SetLevel(slog.LevelWarn)
	case "error":
		Client.SetLevel(slog.LevelError)
	case "fatal":
		Client.SetLevel(slog.LevelFatal)
	default:
		Client.SetLevel(slog.LevelInfo)
	}

	switch strings.ToLower(clientConfig.Log.Output) {
	case "stdout":
		Client.SetOutputToStdout()
	case "stderr":
		Client.SetOutputToStderr()
	case "file":
		name := clientConfig.Log.File.Name
		extensionName := clientConfig.Log.File.ExtensionName
		addDate := clientConfig.Log.File.AddDate

		Client.SetOutputToFile(name, extensionName, addDate)
	}

	Client.SetWithCallerInfo(clientConfig.Log.WithCallerInfo)
}
