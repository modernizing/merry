package domain

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/igso/pkg/infrastructure/csvconv"
	"path/filepath"
	"testing"
)

func Test_ShouldConvertCsvToMavenDependencies(t *testing.T) {
	g := NewGomegaWithT(t)

	absPath := filepath.FromSlash(`../../_fixtures/map/map.csv`)
	csv := csvconv.ParseCSV(absPath)

	var deps []MavenDependency
	deps = FromCSV(csv)

	g.Expect(len(deps)).To(Equal(9))
}
