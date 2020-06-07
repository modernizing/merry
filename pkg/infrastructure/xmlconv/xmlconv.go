package xmlconv

import (
	"encoding/xml"
	"fmt"
)

type AntModel struct {
	XMLName xml.Name `xml:"project"`
	Name    string   `xml:"name,attr"`
}

func XmlConvert(str string) AntModel {
	var model AntModel
	err := xml.Unmarshal([]byte(str), &model)
	if err != nil {
		fmt.Print(err)
	}
	return model
}
