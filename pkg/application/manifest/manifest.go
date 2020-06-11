package manifest

import (
	"github.com/phodal/igso/pkg/domain"
	"github.com/phodal/igso/pkg/infrastructure/ast"
)

func ProcessManifest(content string, filename string) dependency.IgsoManifest {
	results := ast.Analysis(content, filename)
	return results
}