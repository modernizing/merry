package cmd

import (
	"github.com/phodal/igso/cmdtest/testcase"
	"testing"
)

func Test_ShouldBuildBundleMap(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "map",
		Cmd:    "map -p ../_fixtures/bundle/all",
		Golden: "testdata/map/all.csv",
	}}
	RunTestCmd(t, tests)
}