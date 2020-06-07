package xmlconv

import (
	"encoding/xml"
	"fmt"
)

type Property struct {
	Name     string `xml:"name,attr"`
	Location string `xml:"location,attr"`
}

type Launch struct {
	Target string `xml:"target,attr"`
}

type Target struct {
	Name    string `xml:"name,attr"`
	Depends string `xml:"depends,attr"`
	Launch  Launch `xml:"launch"`
}

type AntModel struct {
	XMLName     xml.Name   `xml:"project"`
	Name        string     `xml:"name,attr"`
	Default     string     `xml:"default,attr"`
	BaseDir     string     `xml:"basedir,attr"`
	Description string     `xml:"description"`
	Property    []Property `xml:"property"`
	Target      []Target   `xml:"target"`
}

func XmlConvert(str string) AntModel {
	var model AntModel
	err := xml.Unmarshal([]byte(str), &model)
	if err != nil {
		fmt.Print(err)
	}
	return model
}
