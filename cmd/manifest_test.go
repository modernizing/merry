package cmd

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/igso/cmdtest/testcase"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func Test_ShouldCommandManifest(t *testing.T) {
	g := NewGomegaWithT(t)
	tests := []testcase.CmdTestCase{{
		Name:   "manifest",
		Cmd:    "manifest -p ../_fixtures/manifest/deps/log4j -s",
		Golden: "",
	}}
	RunTestCmd(t, tests)

	content, _ := ioutil.ReadFile(filepath.FromSlash("dep.dot"))
	g.Expect(string(content)).To(Equal(`digraph G {

}
`))
}