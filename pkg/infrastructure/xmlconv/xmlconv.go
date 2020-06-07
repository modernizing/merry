package xmlconv

import (
	"encoding/xml"
	"fmt"
)

type Property struct {
	Name     string `xml:"name,attr"`
	Location string `xml:"location,attr"`
}

type AntModel struct {
	XMLName     xml.Name   `xml:"project"`
	Name        string     `xml:"name,attr"`
	Default     string     `xml:"default,attr"`
	BaseDir     string     `xml:"basedir,attr"`
	Description string     `xml:"description"`
	Property    []Property `xml:"property"`
}

func XmlConvert(str string) AntModel {
	var model AntModel
	err := xml.Unmarshal([]byte(str), &model)
	if err != nil {
		fmt.Print(err)
	}
	return model
}
