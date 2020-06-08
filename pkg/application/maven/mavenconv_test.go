package maven

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestShouldGetPropertyName(t *testing.T) {
	g := NewGomegaWithT(t)
	result := FromAnt("")
	g.Expect(result).To(Equal(""))
}

