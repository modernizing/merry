package cmd

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/igso/cmdtest/testcase"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func Test_ShouldBuildPomXmlWithDupSearch(t *testing.T) {
	g := NewGomegaWithT(t)
	tests := []testcase.CmdTestCase{{
		Name:   "dupsearch",
		Cmd:    "dupsearch -p ../_fixtures/bundle/sample",
		Golden: "",
	}}
	RunTestCmd(t, tests)

	content, _ := ioutil.ReadFile(filepath.FromSlash("../_fixtures/bundle/sample/pom.xml"))
	g.Expect(string(content)).To(Equal(`<project xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://maven.apache.org/POM/4.0.0"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">

    <modelVersion>4.0.0</modelVersion>
	<groupId>com.igso</groupId>        
	<artifactId>test</artifactId>
	<version>0.0.1</version>

    <dependencies>
      <dependency>
          <groupId>org.vaadin.example.osgi</groupId>
          <artifactId>felix-tut-8</artifactId>
          <version>1.0-SNAPSHOT</version>
      </dependency>

    </dependencies>

</project>
`))
}
