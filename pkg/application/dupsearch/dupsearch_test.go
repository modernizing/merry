package dupsearch

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestShouldGetByPath(t *testing.T) {
	g := NewGomegaWithT(t)

	files := DupSearch("../../../_fixtures/bundle/sample")
	g.Expect(len(files)).To(Equal(2))
}