package config

import "github.com/heaven-chp/common-library-go/json"

type GrpcClient struct {
	Address string `json:"address"`
	Timeout int    `json:"timeout"`

	Log struct {
		Level           string `json:"level"`
		OutputPath      string `json:"output_path"`
		FileNamePrefix  string `json:"file_name_prefix"`
		PrintCallerInfo bool   `json:"print_caller_info"`
		ChannelSize     int    `json:"channel_size"`
	} `json:"log"`
}

func (this *GrpcClient) parsing(from interface{}) error {
	return json.ToStructFromFile(from.(string), this)
}
