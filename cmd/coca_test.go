package cmd

import (
	"github.com/phodal/igso/cmdtest"
	"github.com/phodal/igso/cmdtest/testcase"
	"testing"
)

func RunTestCmd(t *testing.T, tests []testcase.CmdTestCase) {
	cmdtest.RunTestCaseWithCmd(t, tests, NewRootCmd)
}

