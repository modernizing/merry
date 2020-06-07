package dependency

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestShouldGetVersionByName(t *testing.T) {
	g := NewGomegaWithT(t)

	result := ByFileName(`pax-exam-1.1.0.jar`)
	g.Expect(result.Version).To(Equal(`1.1.0`))
}

func TestShouldGetArtifactId(t *testing.T) {
	g := NewGomegaWithT(t)

	result := ByFileName(`pax-exam-1.1.0.jar`)
	g.Expect(result.ArtifactId).To(Equal(`pax-exam`))
}
