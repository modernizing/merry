package dupsearch

import (
	"fmt"
	. "github.com/phodal/igso/pkg/domain"
	"github.com/phodal/igso/pkg/infrastructure"
	"github.com/phodal/igso/pkg/infrastructure/properties"
	"strings"
)

func DupSearch(root string) []MavenDependency {
	var results []MavenDependency
	var jarFileFilter = func(path string) bool {
		return strings.HasSuffix(path, ".jar")
	}

	jarPaths := infrastructure.GetFilesByFilter(root, jarFileFilter)
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

