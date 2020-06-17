package visual

import (
	. "github.com/onsi/gomega"
	"testing"
)

func Test_ShouldEmptyWhenConvert(t *testing.T) {
	g := NewGomegaWithT(t)

	content := ManifestFileToDContent("../../../_fixtures/visual/empty-map.json")

	g.Expect(string(content)).To(Equal(`{"nodes":[{"group":1}]}`))
}
