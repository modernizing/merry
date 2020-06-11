package manifest

import (
	. "github.com/onsi/gomega"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func Test_ShouldParseManifestFile(t *testing.T) {
	g := NewGomegaWithT(t)

	file := filepath.FromSlash("../../../_fixtures/manifest/spring-core/MANIFEST.MF")
	content, _ := ioutil.ReadFile(file)
	manifest := ProcessManifest(string(content), "hello.mf")
	g.Expect(len(manifest.ExportPackage)).To(Equal(19))
	g.Expect(len(manifest.ImportPackage)).To(Equal(15))
}


func Test_ShouldSuccessScanManifestFiles(t *testing.T) {
	g := NewGomegaWithT(t)

	file := filepath.FromSlash("../../../_fixtures/manifest/deps")
	manifests := ScanByPath(file)

	g.Expect(len(manifests)).To(Equal(5))
}
