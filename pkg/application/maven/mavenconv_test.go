package maven

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestShouldGetPropertyName(t *testing.T) {
	g := NewGomegaWithT(t)
	result := FromAnt(`
<project name="testing.example" default="dist">
  <target name="test.management" depends="dist">
    <path id="classpath">
       <pathelement path="../../../_fixtures/jar/sample/annotations-13.0.jar" />
    </path>
  </target>
</project>
`)
	g.Expect(len(result)).To(Equal(1))
}

