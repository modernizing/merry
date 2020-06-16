package config

type IgsoConfig struct {
	ReporterPath string
}

const VERSION = "0.0.1"
var CmdConfig = &IgsoConfig{
	ReporterPath: "merry",
}
