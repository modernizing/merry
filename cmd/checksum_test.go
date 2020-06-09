package cmd

import (
	"github.com/phodal/igso/cmdtest/testcase"
	"testing"
)

func Test_ShouldOutputCount(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "checksum",
		Cmd:    "checksum -p ../LICENSE",
		Golden: "testdata/checksum.txt",
	}}
	RunTestCmd(t, tests)
}