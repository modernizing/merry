package xmlconv

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestShouldGetProjectName(t *testing.T) {
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
	g.Expect(result.Default).To(Equal(`dist`))
	g.Expect(result.BaseDir).To(Equal(`.`))
}

func TestShouldGetPropertyName(t *testing.T) {
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
	g.Expect(result.Property[0].Name).To(Equal(`src`))
	g.Expect(result.Property[1].Name).To(Equal(`build`))
	g.Expect(result.Property[2].Name).To(Equal(`dist`))
	g.Expect(result.Property[2].Location).To(Equal(`dist`))
}

func TestShouldGetTargetInfo(t *testing.T) {
	g := NewGomegaWithT(t)
	xml := `
<project name="testing.example" default="dist">
  <target name="test.management" depends="dist">
    <launch target="management"/>
  </target>
</project>
`

	result := XmlConvert(xml)
	g.Expect(result.Target[0].Name).To(Equal(`test.management`))
	g.Expect(result.Target[0].Depends).To(Equal(`dist`))
	g.Expect(result.Target[0].Launch.Target).To(Equal(`management`))
}

func TestShouldIdentifyPackageInfo(t *testing.T) {
	g := NewGomegaWithT(t)
	xml := `
<project>
  <modelVersion>4.0.0</modelVersion>
  <artifactId>my-test-app</artifactId>
  <groupId>my-test-group</groupId>
  <version>1.0-SNAPSHOT</version>
</project>
`

	result := XmlConvert(xml)
	g.Expect(result.ModelVersion).To(Equal(`4.0.0`))
	g.Expect(result.ArtifactId).To(Equal(`my-test-app`))
	g.Expect(result.GroupId).To(Equal(`my-test-group`))
	g.Expect(result.Version).To(Equal(`1.0-SNAPSHOT`))
}
