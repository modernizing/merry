package dependency


import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestShouldGetVersionByName(t *testing.T) {
	g := NewGomegaWithT(t)

	var deps []MavenDependency
	dependency := MavenDependency{
		Version:    "1.0.0",
		GroupId:    "com.phodal",
		ArtifactId: "igso",
	}
	deps = append(deps, dependency)

	result := RemoveDuplicate(deps)
	g.Expect(len(result)).To(Equal(1))
}
