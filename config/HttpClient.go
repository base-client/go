package config

type HttpClient struct {
	Address string `json:"address"`

	Log struct {
		Level  string `json:"level"`
		Output string `json:"output"`
		File   struct {
			Name          string `json:"name"`
			ExtensionName string `json:"extensionName"`
			AddDate       bool   `json:"addDate"`
		} `json:"file"`
		WithCallerInfo bool `json:"withCallerInfo"`
	} `json:"log"`
}
