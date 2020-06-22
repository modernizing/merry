package domain

type KeyValue struct {
	Key   string
	Value string
}

type JavaPackage struct {
	Name          string
	VersionInfo   string
	ExportVersion string
	StartVersion  string
	EndVersion    string
	Config        []KeyValue
}

type IgsoManifest struct {
	ExportPackage []JavaPackage
	ImportPackage []JavaPackage
	KeyValues     []KeyValue
	PackageName   string
	Version       string
}


func (m *IgsoManifest) BuildJarFileName() string {
	return m.PackageName + "_" + m.Version + ".jar"
}

