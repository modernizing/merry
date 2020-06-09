package maven

import (
	"fmt"
	. "github.com/phodal/igso/pkg/domain"
	"github.com/phodal/igso/pkg/infrastructure/bundle"
	"github.com/phodal/igso/pkg/infrastructure/properties"
	"github.com/phodal/igso/pkg/infrastructure/xmlconv"
	"strings"
)

func FromAnt(str string, extract bool) MavenProject {
	antModel := xmlconv.XmlConvert(str)

	project := &MavenProject{
		Version:      antModel.Version,
		ArtifactId:   antModel.ArtifactId,
		GroupId:      antModel.GroupId,
		ModelVersion: antModel.ModelVersion,
	}

	project.Dependencies = BuildDependencies(antModel, extract)
	return *project
}

func BuildDependencies(antModel xmlconv.AntModel, extract bool) []MavenDependency {
	var results []MavenDependency
	for _, target := range antModel.Target {
		if target.Path.PathElements != nil {
			for _, pathElement := range target.Path.PathElements {
				if extract {
					results = readWithUnzip(pathElement, results)
				} else {
					props, err := properties.ReadPropertiesFromZip(pathElement.Path)
					if (err == nil) && props["groupId"] != "" {
						results = append(results, MavenDependency{
							Version:    props["version"],
							GroupId:    props["groupId"],
							ArtifactId: props["artifactId"],
						})
					} else {
						fmt.Println("lost properties file: " + pathElement.Path)
					}
				}
			}
		}
	}

	return results
}

func readWithUnzip(pathElement xmlconv.PathElement, results []MavenDependency) []MavenDependency {
	filenames, _ := bundle.Unzip(pathElement.Path, "tmp")
	for _, filename := range filenames {
		if strings.HasSuffix(filename, "pom.properties") {
			props, _ := properties.ReadPropertiesFile(filename)
			results = append(results, MavenDependency{
				Version:    props["version"],
				GroupId:    props["groupId"],
				ArtifactId: props["artifactId"],
			})
		}
	}
	return results
}
