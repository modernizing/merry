package xmlconv

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestNormalConvert(t *testing.T) {
	g := NewGomegaWithT(t)
	xml := `
<?xml version="1.0" encoding="UTF-8"?>
<project name="MyProject" default="dist" basedir=".">
  <description>
    simple example build file
  </description>
  <!-- set global properties for this build -->
  <property name="src" location="src"/>
  <property name="build" location="build"/>
  <property name="dist" location="dist"/>
</project>`

	result := XmlConvert(xml)
	g.Expect(result.Name).To(Equal(`MyProject`))
}
