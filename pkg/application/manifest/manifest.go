package manifest

import (
	"fmt"
	"github.com/phodal/merry/pkg/domain"
	"github.com/phodal/merry/pkg/infrastructure"
	"github.com/phodal/merry/pkg/infrastructure/ast"
	"github.com/phodal/merry/pkg/infrastructure/bundle"
	"io/ioutil"
	"strings"
)

func ProcessManifest(content string, filename string) domain.IgsoManifest {
	results := ast.Analysis(content, filename)
	return results
}

func ScanByPath(path string) []domain.IgsoManifest {
	var manifests []domain.IgsoManifest

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

func ScanByFile(path string) domain.IgsoManifest {
	_, content, err := bundle.GetFileFromJar(path, "MANIFEST.MF")
	if err != nil {
		fmt.Println(err)
		return domain.IgsoManifest{};
	}

	mani := ProcessManifest(content, "MANIFEST.MF")

	return mani
}