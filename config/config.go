// Package config provides a struct that can store json type config file
package config

type GrpcClient struct {
	LogLevel          string `json:"log_level"`
	LogOutputPath     string `json:"log_output_path"`
	LogFileNamePrefix string `json:"log_file_name_prefix"`

	Address string `json:"address"`
	Timeout int    `json:"timeout"`
}

type SocketClient struct {
	LogLevel          string `json:"log_level"`
	LogOutputPath     string `json:"log_output_path"`
	LogFileNamePrefix string `json:"log_file_name_prefix"`

	Address string `json:"address"`
}
