package infrastructure


import (
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func TestShouldGetProjectName(t *testing.T) {
	g := NewGomegaWithT(t)
	path := GetJarPath(filepath.FromSlash("com/phodal/hello.jar"))

	g.Expect(path).To(Equal(filepath.FromSlash(`com/phodal/`)))
}
