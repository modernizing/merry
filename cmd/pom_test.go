package cmd

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/merry/cmdtest/testcase"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func Test_ShouldGeneratePomFromJar(t *testing.T) {
	g := NewGomegaWithT(t)
	tests := []testcase.CmdTestCase{{
		Name:   "pom",
		Cmd:    "pom -p ../_fixtures/pom/",
		Golden: "",
	}}
	RunTestCmd(t, tests)

	content, _ := ioutil.ReadFile(filepath.FromSlash("../_fixtures/pom/felix-tut-1-1.0-20200611.070711-1.pom"))
	g.Expect(string(content)).To(Equal(`<project xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://maven.apache.org/POM/4.0.0"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">

    <modelVersion>4.0.0</modelVersion>
	<groupId>20200611.070711</groupId>        
	<artifactId></artifactId>
	<version>1</version>

    <dependencies>
      <dependency>
          <groupId>org.osgi</groupId>
          <artifactId>framework</artifactId>
          <version>1.8</version>
      </dependency>

    </dependencies>

</project>
`))
}
