package domain


import (
	. "github.com/onsi/gomega"
	"testing"
)

func Test_ShouldBuildDData(t *testing.T) {
	g := NewGomegaWithT(t)


	var pkgs []JavaPackage
	pkgs = append(pkgs, JavaPackage{Name: "ognl", StartVersion: "5.0.0_201"})
	pkgs = append(pkgs, JavaPackage{Name: "ognl.factory", StartVersion: "5.0.0_201"})
	pkgs = append(pkgs, JavaPackage{Name: "factory", StartVersion: "5.0.0_201"})

	var manifests []IgsoManifest
	manifest := IgsoManifest{
		ExportPackage: nil,
		ImportPackage: pkgs,
		KeyValues:     nil,
		PackageName:   "hello.name",
		Version:       "",
	}
	manifests = append(manifests, manifest)

	result := VisualFromManifest(manifests)
	g.Expect(len(result.Links)).To(Equal(3))
	g.Expect(len(result.Nodes)).To(Equal(4))
}