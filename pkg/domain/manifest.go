package dependency

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
	Package       []JavaPackage
	KeyValues     []KeyValue
}
