package tequila


import (
	. "github.com/onsi/gomega"
	"testing"
)

func Test_ShouldGetComWhenMergeHeader(t *testing.T) {
	g := NewGomegaWithT(t)

	header := MergePackageFunc("com.phodal.chapi")
	g.Expect(header).To(Equal("com"))
}

func Test_ShouldMergeSplitWhenMergeHeader(t *testing.T) {
	g := NewGomegaWithT(t)

	header := MergePackageFunc("pkg/com.phodal.chapi")
	g.Expect(header).To(Equal("pkg"))
}

func Test_ShouldGetMainWhenHasOneSelect(t *testing.T) {
	g := NewGomegaWithT(t)

	header := MergePackageFunc("com")
	g.Expect(header).To(Equal("main"))
}

