package xmlconv

import (
	"encoding/xml"
	"fmt"
)

type Property struct {
	Name     string `xml:"name,attr"`
	Location string `xml:"location,attr"`
}

type PathElement struct {
	Path string `xml:"path,attr"`
}

type Path struct {
	Id           string        `xml:"id,attr"`
	PathElements []PathElement `xml:"pathelement"`
}

type Launch struct {
	Target string `xml:"target,attr"`
}

type Target struct {
	Name    string `xml:"name,attr"`
	Depends string `xml:"depends,attr"`
	Launch  Launch `xml:"launch"`
	Path    Path   `xml:"path"`
}

type AntModel struct {
	XMLName      xml.Name   `xml:"project"`
	ModelVersion string     `xml:"modelVersion"`
	ArtifactId   string     `xml:"artifactId"`
	GroupId      string     `xml:"groupId"`
	Version      string     `xml:"version"`
	Name         string     `xml:"name,attr"`
	Default      string     `xml:"default,attr"`
	BaseDir      string     `xml:"basedir,attr"`
	Description  string     `xml:"description"`
	Property     []Property `xml:"property"`
	Target       []Target   `xml:"target"`
}

func XmlConvert(str string) AntModel {
	var model AntModel
	err := xml.Unmarshal([]byte(str), &model)
	if err != nil {
		fmt.Print(err)
	}
	return model
}
