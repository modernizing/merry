package infrastructure

import (
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func TestShouldGetPropertyName(t *testing.T) {
	g := NewGomegaWithT(t)
	path := `../../../_fixtures/bundle/sample/felix-tut-8-1.0-SNAPSHOT.jar`
	absPath := filepath.FromSlash(path)

	result, _ := HashFileMD5(absPath)
	g.Expect(result).To(Equal(`49228cf9d883ca8343234b895dfb3623`))
}


