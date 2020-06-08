package dupsearch

import (
	. "github.com/onsi/gomega"
	"strings"
	"testing"
)

func TestShouldGetPropertyName(t *testing.T) {
	g := NewGomegaWithT(t)

	var jarFileFilter = func(path string) bool {
		return strings.HasSuffix(path, ".jar")
	}

	files := GetFilesByFilter("../../../_fixtures/bundle/sample", jarFileFilter)

	g.Expect(len(files)).To(Equal(2))
}

func TestShouldGetByPath(t *testing.T) {
	g := NewGomegaWithT(t)


	files := DupSearch("../../../_fixtures/bundle/sample")
	g.Expect(len(files)).To(Equal(2))
}