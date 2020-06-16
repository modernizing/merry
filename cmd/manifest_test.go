package cmd

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/merry/cmd/config"
	"github.com/phodal/merry/cmdtest/testcase"
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

	content, _ := ioutil.ReadFile(filepath.FromSlash(config.CmdConfig.ReporterPath + "/" + "dep.dot"))
	g.Expect(string(content)).To(Equal(`digraph G {
	subgraph cluster1 {
	label="main";
	rankdir=LR;
	node1 [ label="", shape=box ];

}
;

}
`))
}

func Test_ShouldVersionManifest(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "manifest",
		Cmd:    "manifest  -p ../_fixtures/bundle/sample/felix-duplicate.jar -v",
		Golden: "testdata/manifest/version.txt",
	}}
	RunTestCmd(t, tests)
}

func Test_ShouldFilterJarForManifest(t *testing.T) {
	g := NewGomegaWithT(t)
	tests := []testcase.CmdTestCase{{
		Name:   "manifest",
		Cmd:    "manifest -p ../_fixtures/manifest/deps/log4j -s",
		Golden: "",
	}}
	RunTestCmd(t, tests)

	content, _ := ioutil.ReadFile(filepath.FromSlash(config.CmdConfig.ReporterPath + "/" + "dep.dot"))
	g.Expect(string(content)).To(Equal(`digraph G {

}
`))
}
