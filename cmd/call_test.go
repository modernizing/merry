package cmd

import (
	"github.com/phodal/igso/cmdtest/testcase"
	"testing"
)

func Test_ShouldStartServer(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "call",
		Cmd:    "call -p .",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}
