package manifest

import (
	"fmt"
	"github.com/phodal/igso/pkg/domain"
	"github.com/phodal/igso/pkg/infrastructure/ast"
	"github.com/phodal/igso/pkg/infrastructure/bundle"
	"github.com/phodal/igso/pkg/instrastructure"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ProcessManifest(content string, filename string) dependency.IgsoManifest {
	results := ast.Analysis(content, filename)
	return results
}

func ExtractManifest(ppath string) {
	var jarFileFilter = func(path string) bool {
		return strings.HasSuffix(path, ".jar")
	}

	jarPaths := instrastructure.GetFilesByFilter(ppath, jarFileFilter)
	fmt.Println(ppath)
	for _, path := range jarPaths {
		_, content, _ := bundle.GetFileFromJar(path, "MANIFEST.MF")
		pureName := path[len(ppath)+1 : len(path)-len(".jar")]

		filePath := "_m/" + pureName + "/META-INF/"

		_ = os.MkdirAll(filepath.FromSlash(filePath), os.ModePerm)
		_ = ioutil.WriteFile(filepath.FromSlash(filePath+"MANIFEST.MF"), []byte(content), os.ModePerm)
	}
}
