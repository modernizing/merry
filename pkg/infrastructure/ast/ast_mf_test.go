package ast

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestShouldGetPropertyName(t *testing.T) {
	g := NewGomegaWithT(t)
	code := `
Manifest-Version: 1.0

`
	results := Analysis(code, "hello.mf")
	g.Expect(len(results)).To(Equal(1))
	g.Expect(results[0].Key).To(Equal("Manifest-Version"))
	g.Expect(results[0].Value).To(Equal("1.0"))
}

