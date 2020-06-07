package dependency

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestShouldGetProjectName(t *testing.T) {
	g := NewGomegaWithT(t)

	result := ByFileName(`pax-exam-1.1.0.jar`)
	g.Expect(result.Version).To(Equal(`1.1.0`))
}

