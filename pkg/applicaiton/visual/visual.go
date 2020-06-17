package visual

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/merry/pkg/domain"
	"io/ioutil"
	"path/filepath"
)

func ManifestFileToDContent(path string) []byte {
	abs, _ := filepath.Abs(filepath.FromSlash(path))
	originContent, err := ioutil.ReadFile(abs)
	if err != nil {
		fmt.Println(err)
	}

	var manifests []domain.IgsoManifest
	_ = json.Unmarshal(originContent, &manifests)
	dData := domain.VisualFromManifest(manifests)
	dContent, err := json.Marshal(dData)
	return dContent
}
