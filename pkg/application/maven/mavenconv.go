package maven

import (
	. "github.com/phodal/igso/pkg/domain"
	"github.com/phodal/igso/pkg/infrastructure/bundle"
	"github.com/phodal/igso/pkg/infrastructure/xmlconv"
)

func FromAnt(str string) []MavenDependency {
	var results []MavenDependency
	antModel := xmlconv.XmlConvert(str)
	for _, target := range antModel.Target {
		if target.Path.PathElements != nil {
			for _, pathElement := range target.Path.PathElements {
				bundle.Unzip(pathElement.Path, "tmp")
				results = append(results, MavenDependency{})
			}
		}
	}

	return results
}
