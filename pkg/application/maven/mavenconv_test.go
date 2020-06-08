package maven

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestShouldGetPropertyName(t *testing.T) {
	g := NewGomegaWithT(t)
	ant := FromAnt(`
<project name="testing.example" default="dist">
  <target name="test.management" depends="dist">
    <path id="classpath">
       <pathelement path="../../../_fixtures/jar/sample/annotations-13.0.jar" />
    </path>
  </target>
</project>
`)
	result := ant.Dependencies

	g.Expect(len(result)).To(Equal(1))
	g.Expect(result[0].ArtifactId).To(Equal("annotations"))
	g.Expect(result[0].GroupId).To(Equal("org.jetbrains"))
	g.Expect(result[0].Version).To(Equal("13.0"))
}

