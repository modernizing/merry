package cmd

import (
	"github.com/phodal/merry/cmdtest/testcase"
	"testing"
)

func Test_CMD_Fix(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "fix jar package issue",
		Cmd:    "fix -p ../_fixtures/fix",
		Golden: "testdata/fix.txt",
	}}
	RunTestCmd(t, tests)
}