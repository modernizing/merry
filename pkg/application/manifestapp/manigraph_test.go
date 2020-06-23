package manifestapp

import (
	. "github.com/onsi/gomega"
	"path/filepath"
	"strings"
	"testing"
)

func Test_ShouldSuccessBuildGraph(t *testing.T) {
	g := NewGomegaWithT(t)

	file := filepath.FromSlash("../../../_fixtures/manifest/deps/aspectj")
	scanManifest := ScanByPath(file)
	graph := BuildFullGraph(scanManifest, nil)

	ignores := strings.Split("", ",")
	var nodeFilter = func(key string) bool {
		for _, f := range ignores {
			if key == f {
				return true
			}
		}
		return false
	}
	g.Expect(graph.ToDot("", nodeFilter).String()).To(Equal(`graph G {

}
`))
}

func Test_ShouldSuccessIdentifyGraph(t *testing.T) {
	g := NewGomegaWithT(t)

	file := filepath.FromSlash("../../../_fixtures/manifest/deps/commons")
	scanManifest := ScanByPath(file)
	graph := BuildFullGraph(scanManifest, nil)

	dot := graph.ToDot("", func(s string) bool {
		return false
	})
	g.Expect(len(dot.Nodes.Nodes)).To(Equal(12))
	g.Expect(len(dot.SubGraphs.SubGraphs)).To(Equal(12))
}

func Test_ShouldSuccessFoundImportPackage(t *testing.T) {
	g := NewGomegaWithT(t)

	file := filepath.FromSlash("../../../_fixtures/manifest/graph/importpkg")
	scanManifest := ScanByPath(file)
	graph := BuildFullGraph(scanManifest, nil)

	dot := graph.ToDot("", func(s string) bool {
		return false
	})
	g.Expect(len(dot.Nodes.Nodes)).To(Equal(4))
	g.Expect(len(dot.SubGraphs.SubGraphs)).To(Equal(4))
}
