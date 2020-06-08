package properties

import (
	"fmt"
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func TestShouldGetPropertyName(t *testing.T) {
	g := NewGomegaWithT(t)
	path := `../../../_fixtures/properties/sample.properties`
	absPath := filepath.FromSlash(path)

	result, _ := ReadPropertiesFile(absPath)
	g.Expect(result["version"]).To(Equal("13.0"))
	g.Expect(result["groupId"]).To(Equal("org.jetbrains"))
	g.Expect(result["artifactId"]).To(Equal("annotations"))
}

func TestShouldGetPropertyFromZip(t *testing.T) {
	g := NewGomegaWithT(t)
	path := `../../../_fixtures/jar/sample/annotations-13.0.jar`
	absPath := filepath.FromSlash(path)

	result, _ := ReadPropertiesFromZip(absPath)
	fmt.Println(result)
	g.Expect(result["version"]).To(Equal("13.0"))
	g.Expect(result["groupId"]).To(Equal("org.jetbrains"))
	g.Expect(result["artifactId"]).To(Equal("annotations"))
}
