package cmd

import (
	"github.com/phodal/merry/cmdtest"
	"github.com/phodal/merry/cmdtest/testcase"
	"testing"
)

func RunTestCmd(t *testing.T, tests []testcase.CmdTestCase) {
	cmdtest.RunTestCaseWithCmd(t, tests, NewRootCmd)
}

