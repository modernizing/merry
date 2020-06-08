package maven

import (
	. "github.com/phodal/igso/pkg/domain"
	"github.com/phodal/igso/pkg/infrastructure/bundle"
	"github.com/phodal/igso/pkg/infrastructure/properties"
	"github.com/phodal/igso/pkg/infrastructure/xmlconv"
	"strings"
)

func FromAnt(str string) MavenProject {
	antModel := xmlconv.XmlConvert(str)

	project := &MavenProject{
		Version:      antModel.Version,
		ArtifactId:   antModel.ArtifactId,
		GroupId:      antModel.GroupId,
		ModelVersion: antModel.ModelVersion,
	}
	project.Dependencies = BuildDependencies(antModel)

	return *project
}

func BuildDependencies(antModel xmlconv.AntModel) []MavenDependency {
	var results []MavenDependency
	for _, target := range antModel.Target {
		if target.Path.PathElements != nil {
			for _, pathElement := range target.Path.PathElements {
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
			}
		}
	}

	return results
}
