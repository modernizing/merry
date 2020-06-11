package manifest

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
	graph := BuildFullGraph(scanManifest)

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
