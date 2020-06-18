package visual

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/phodal/merry/pkg/domain"
	"testing"
)

func Test_ShouldEmptyWhenConvert(t *testing.T) {
	g := NewGomegaWithT(t)

	content := ManifestFileToDContent("../../../_fixtures/visual/empty-map.json")

	g.Expect(string(content)).To(Equal(`{
	"nodes": [
		{
			"group": 1
		}
	]
}`))
}

func Test_ShouldHaveTwoImportWhenConvert(t *testing.T) {
	g := NewGomegaWithT(t)

	content := ManifestFileToDContent("../../../_fixtures/visual/import-map.json")

	var data domain.DData
	json.Unmarshal(content, &data)

	g.Expect(len(data.Nodes)).To(Equal(4))
	g.Expect(len(data.Links)).To(Equal(6))
}
