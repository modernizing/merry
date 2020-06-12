package infrastructure


import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestShouldGetProjectName(t *testing.T) {
	g := NewGomegaWithT(t)
	path := GetJarPath("com/phodal/hello.jar")

	g.Expect(path).To(Equal(`com/phodal/`))
}