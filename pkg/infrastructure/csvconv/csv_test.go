package csvconv

import (
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func TestShouldGetPropertyName(t *testing.T) {
	g := NewGomegaWithT(t)
	csvPath := `../../../_fixtures/map/map.csv`
	absPath := filepath.FromSlash(csvPath)

	csv := ParseCSV(absPath)
	g.Expect(len(csv)).To(Equal(10))
	header := csv[0]
	g.Expect(len(header)).To(Equal(5))
	g.Expect(header[0]).To(Equal("origin"))
}

