package manifest

import (
	"fmt"
	"github.com/phodal/igso/pkg/domain"
	"github.com/phodal/igso/pkg/infrastructure"
	"github.com/phodal/igso/pkg/infrastructure/ast"
	"github.com/phodal/igso/pkg/infrastructure/bundle"
	"io/ioutil"
	"strings"
)

func ProcessManifest(content string, filename string) dependency.IgsoManifest {
	results := ast.Analysis(content, filename)
	return results
}

func ScanByPath(path string) []dependency.IgsoManifest {
	var manifests []dependency.IgsoManifest

	manifestPaths := infrastructure.GetFilesByFilter(path, func(path string) bool {
		return strings.HasSuffix(path, "MANIFEST.MF")
	})

	for _, mPath := range manifestPaths {
		content, _ := ioutil.ReadFile(mPath)
		manifest := ProcessManifest(string(content), "MANIFEST.MF")
		manifests = append(manifests, manifest)
	}

	return manifests
}

func ScanByFile(path string) dependency.IgsoManifest {
	_, content, err := bundle.GetFileFromJar(path, "MANIFEST.MF")
	if err != nil {
		fmt.Println(err)
		return dependency.IgsoManifest{};
	}

	mani := ProcessManifest(content, "MANIFEST.MF")

	return mani
}