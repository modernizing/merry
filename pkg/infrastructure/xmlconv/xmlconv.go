package xmlconv

import (
	"encoding/xml"
	"fmt"
)

type Property struct {
	Name string `xml:"name,attr"`
}

type AntModel struct {
	XMLName     xml.Name   `xml:"project"`
	Name        string     `xml:"name,attr"`
	Description string     `xml:"description"`
	Property    []Property `xml:"property"`
}

func XmlConvert(str string) AntModel {
	var model AntModel
	err := xml.Unmarshal([]byte(str), &model)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(model)
	return model
}
