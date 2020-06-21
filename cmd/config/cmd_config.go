package config

type IgsoConfig struct {
	ReporterPath string
}

const VERSION = "0.0.2"
var CmdConfig = &IgsoConfig{
	ReporterPath: "merry",
}
