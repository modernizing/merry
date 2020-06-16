package config

type TypeCocaConfig struct {
	ReporterPath string
}

const VERSION = "0.0.1"
var CocaConfig = &TypeCocaConfig{
	ReporterPath: "coca_reporter",
}
