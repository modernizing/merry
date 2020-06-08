package maven

import (
	. "github.com/phodal/igso/pkg/domain"
)

func FromAnt(str string) []MavenDependency {
	var results []MavenDependency
	results = append(results, MavenDependency{})
	return results
}
