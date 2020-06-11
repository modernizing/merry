package manifest

import (
	"github.com/phodal/igso/pkg/domain"
	"github.com/phodal/igso/pkg/infrastructure/ast"
	"github.com/phodal/igso/pkg/instrastructure"
	"strings"
)

func ProcessManifest(content string, filename string) dependency.IgsoManifest {
	results := ast.Analysis(content, filename)
	return results
}

func ScanManifest(path string) []dependency.IgsoManifest {
	var manifests []dependency.IgsoManifest

	manifestPaths := instrastructure.GetFilesByFilter(path, func(path string) bool {
		return strings.HasSuffix(path, "MANIFEST.MF")
	})

	for _, mPath := range manifestPaths {
		//fmt.Println(manifest)
		manifest := ProcessManifest(mPath, "MANIFEST.MF")
		manifests = append(manifests, manifest)
	}

	return manifests
}