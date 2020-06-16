package domain

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/merry/pkg/infrastructure/csvconv"
	"path/filepath"
	"testing"
)

func Test_ShouldConvertCsvToMavenDependencies(t *testing.T) {
	g := NewGomegaWithT(t)

	absPath := filepath.FromSlash(`../../_fixtures/map/map.csv`)
	csv := csvconv.ParseCSV(absPath)

	var deps []MavenDependency
	deps, _ = MapFromCSV(csv)

	g.Expect(len(deps)).To(Equal(10))
}

func Test_ShouldBuildDependenciesMap(t *testing.T) {
	g := NewGomegaWithT(t)

	absPath := filepath.FromSlash(`../../_fixtures/map/map.csv`)
	csv := csvconv.ParseCSV(absPath)

	_, depmap := MapFromCSV(csv)

	dependency := depmap["org.apache.commons.logging"]

	g.Expect(dependency.GroupId).To(Equal("commons-logging"))
	g.Expect(dependency.ArtifactId).To(Equal("commons-logging"))
}

