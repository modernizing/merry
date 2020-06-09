package cmd


import (
	. "github.com/onsi/gomega"
	"github.com/phodal/igso/cmdtest/testcase"
	"io/ioutil"
	"testing"
)

func Test_ShouldBuildPomXmlWithBoom(t *testing.T) {
	g := NewGomegaWithT(t)
	tests := []testcase.CmdTestCase{{
		Name:   "boom",
		Cmd:    "boom -p testdata",
		Golden: "",
	}}
	RunTestCmd(t, tests)

	content, _ := ioutil.ReadFile("testdata/pom.xml")
	g.Expect(string(content)).To(Equal(`<project xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://maven.apache.org/POM/4.0.0"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">

    <modelVersion>4.0.0</modelVersion>
	<groupId>my-test-group</groupId>        
	<artifactId>my-test-app</artifactId>
	<version>1.0-SNAPSHOT</version>

    <dependencies>
      <dependency>
          <groupId>org.vaadin.example.osgi</groupId>
          <artifactId>felix-tut-8</artifactId>
          <version>1.0-SNAPSHOT</version>
      </dependency>

      <dependency>
          <groupId>org.jetbrains</groupId>
          <artifactId>annotations</artifactId>
          <version>13.0</version>
      </dependency>

    </dependencies>

</project>
`))
}