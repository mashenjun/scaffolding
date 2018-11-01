package template

const ConfigFile =`package config

type Backend struct {
	HttpHost       string         `+"`json:\"http_host\"`"+"\n"+
`}

// GlobalConf readonly configure
var GlobalConf Backend

func Init(configPath string) error{
	panic("not implemented")
}`
