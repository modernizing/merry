package cmd


import (
	. "github.com/onsi/gomega"
	"github.com/phodal/igso/cmdtest/testcase"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func Test_ShouldBuildPomXmlWithBoom(t *testing.T) {
	g := NewGomegaWithT(t)
	tests := []testcase.CmdTestCase{{
		Name:   "boom",
		Cmd:    "boom -p testdata/boom",
		Golden: "",
	}}
	RunTestCmd(t, tests)

	content, _ := ioutil.ReadFile(filepath.FromSlash("testdata/boom/pom.xml"))
	g.Expect(string(content)).To(Equal(`<project xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://maven.apache.org/POM/4.0.0"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">

    <modelVersion>4.0.0</modelVersion>
	<groupId>my-test-group</groupId>        
	<artifactId>my-test-app</artifactId>
	<version>1.0-SNAPSHOT</version>

    <dependencies>
      <dependency>
          <groupId>org.jetbrains</groupId>
          <artifactId>annotations</artifactId>
          <version>13.0</version>
      </dependency>

    </dependencies>

</project>
`))
}