package dupsearch

import (
	"fmt"
	. "github.com/phodal/igso/pkg/domain"
	"github.com/phodal/igso/pkg/infrastructure/properties"
	"os"
	"path/filepath"
	"strings"
)

func DupSearch(root string) []MavenDependency {
	var results []MavenDependency
	var jarFileFilter = func(path string) bool {
		return strings.HasSuffix(path, ".jar")
	}

	jarPaths := GetFilesByFilter(root, jarFileFilter)
	for _, path := range jarPaths {
		props, err := properties.ReadPropertiesFromZip(path)
		if (err == nil) && props["groupId"] != "" {
			results = append(results, MavenDependency{
				Version:    props["version"],
				GroupId:    props["groupId"],
				ArtifactId: props["artifactId"],
			})
		} else {
			fmt.Println("lost properties file: " + path)
		}
	}

	return results
}

func GetFilesByFilter(root string, filter func(path string) bool) []string {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filter(path) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
